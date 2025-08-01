// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockgcp

import (
	"bytes"
	"context"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type mockRoundTripper struct {
	grpcConnection *grpc.ClientConn
	grpcListener   net.Listener

	iamPolicies *mockIAMPolicies

	registeredServices *mockgcpregistry.Services

	services []registeredService

	server *grpc.Server
}

type registeredService struct {
	hostRegexes []*regexp.Regexp
	handler     http.Handler
	impl        mockgcpregistry.MockService
}

func (h *registeredService) MatchesHost(host string) (http.Handler, bool) {
	for _, hostRegex := range h.hostRegexes {
		if hostRegex.MatchString(host) {
			return h.handler, true
		}
	}
	return nil, false
}

type SupportsTestCommands interface {
	// RunTestCommand is a "backdoor" into our mock implementation that is useful for fault injection or faking scaling events etc
	// In our script-driven tests, we trigger this with a special `MockGCPBackdoor` object.
	RunTestCommand(ctx context.Context, service string, command string) error
}

func NewMockRoundTripperForTest(t *testing.T, k8sClient client.Client, storage storage.Storage) Interface {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	mockRoundTripper, err := NewMockRoundTripper(ctx, k8sClient, storage)
	if err != nil {
		t.Fatalf("building mockgcp: %v", err)
	}

	go func() {
		if err := mockRoundTripper.Run(ctx); err != nil {
			t.Errorf("error from grpc server: %v", err)
		}
	}()

	return mockRoundTripper
}

func (m *mockRoundTripper) ConfigureVisitor(requestURL string, visitor mockgcpregistry.NormalizingVisitor) {
	m.registeredServices.ConfigureVisitor(requestURL, visitor)
}

func (m *mockRoundTripper) Previsit(event mockgcpregistry.Event, visitor mockgcpregistry.NormalizingVisitor) {
	m.registeredServices.Previsit(event, visitor)
}

func (m *mockRoundTripper) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		m.server.Stop()
	}()
	return m.server.Serve(m.grpcListener)
}

func (m *mockRoundTripper) RunTestCommand(ctx context.Context, serviceName string, command string) error {
	for _, service := range m.services {
		if _, match := service.MatchesHost(serviceName); !match {
			continue
		}

		supportsTestCommands, ok := service.impl.(SupportsTestCommands)
		if !ok {
			return fmt.Errorf("service %T does not support test commands", service)
		}
		return supportsTestCommands.RunTestCommand(ctx, serviceName, command)
	}
	return fmt.Errorf("service %q not known", serviceName)
}

func (m *mockRoundTripper) NewGRPCConnection(ctx context.Context) *grpc.ClientConn {
	endpoint := m.grpcListener.Addr().String()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		klog.Fatalf("error dialing grpc endpoint %q: %v", endpoint, err)
	}
	return conn
}

func toHostRegex(host string) *regexp.Regexp {
	r := regexp.MustCompile(`{[^}]+}`)

	tokens := strings.Split(host, ".")
	for i, token := range tokens {
		token = r.ReplaceAllStringFunc(token, func(match string) string {
			return "[^.]*"
		})
		tokens[i] = token
	}
	return regexp.MustCompile("^" + strings.Join(tokens, `\.`) + "$")
}

func (m *mockRoundTripper) prefilterRequest(req *http.Request) error {
	if req.Body != nil {
		var requestBody bytes.Buffer
		if _, err := io.Copy(&requestBody, req.Body); err != nil {
			return fmt.Errorf("error reading request body: %w", err)
		}

		if requestBody.Len() != 0 {
			o := make(map[string]any)
			if err := json.Unmarshal(requestBody.Bytes(), &o); err != nil {
				return fmt.Errorf("parsing json: %w", err)
			}

			if err := m.modifyUpdateMask(o); err != nil {
				return err
			}

			if err := pruneNilArrays(o); err != nil {
				return err
			}

			b, err := json.Marshal(o)
			if err != nil {
				return fmt.Errorf("building json: %w", err)
			}

			req.Body = io.NopCloser(bytes.NewBuffer(b))
		}
	} else {
		// When sending a delete request for a ComputeFirewallPolicyRule resource,
		// The request URL looks like POST https://compute.googleapis.com/compute/v1/locations/global/firewallPolicies/${firewallPolicyID}/removeRule.
		// It's uncommon to use POST requests for delete operations, and a nil request body for POST method is unexpected,
		// I got the "missing form body" error. Ref: https://go.dev/src/net/http/request.go?s=41070:41129 line 1340
		// So instead of sending a nil request body, send an empty request body to ensure successful processing of the remove rule request.
		body := &bytes.Buffer{}
		req.Body = io.NopCloser(body)
	}
	return nil
}

// modifyUpdateMask fixes up the updateMask parameter, which is a proto FieldMask.
// Technically, when transported over JSON it should be passed as json fields (displayName),
// and when transported over proto is should be passed as proto fields (display_name).
// However, because GCP APIs seem to accept display_name or displayName over JSON.
// If we don't map display_name => displayName, the proto validation will reject it.
// e.g. https://github.com/grpc-ecosystem/grpc-gateway/issues/2239
func (m *mockRoundTripper) modifyUpdateMask(o map[string]any) error {
	for k, v := range o {
		switch k {
		case "updateMask":
			vString := v.(string)
			tokens := strings.Split(vString, ",")
			for i, token := range tokens {
				switch token {
				case "display_name":
					tokens[i] = "displayName"
				case "content_type":
					tokens[i] = "contentType"
				}
			}
			o[k] = strings.Join(tokens, ",")
		}
	}

	return nil
}

// pruneNilArrays replaces [nil] => []
// For some reason terraform sends [nil], which is not really valid
func pruneNilArrays(o map[string]any) error {
	for k, v := range o {
		if v == nil {
			continue
		}
		switch v := v.(type) {
		case map[string]any:
			if err := pruneNilArrays(v); err != nil {
				return err
			}
		case []any:
			if len(v) == 1 && v[0] == nil {
				o[k] = []any{}
			}
		case string, int64, bool, float64:
			// ignore
		default:
			return fmt.Errorf("unhandled type %T", v)
		}
	}

	return nil
}

// roundTripIAMPolicy serves the IAM policy verbs (e.g. :getIamPolicy)
// These are implemented on most resources, and rather than mock them
// per-resource, we implement them once here.
func (m *mockRoundTripper) roundTripIAMPolicy(req *http.Request) (*http.Response, error) {
	requestPath := req.URL.Path

	lastColon := strings.LastIndex(requestPath, ":")
	verb := requestPath[lastColon+1:]

	requestPath = strings.TrimSuffix(requestPath, ":"+verb)

	switch verb {
	case "getIamPolicy":
		if req.Method == "GET" || req.Method == "POST" {
			resourcePath := req.URL.Host + requestPath
			return m.iamPolicies.serveGetIAMPolicy(resourcePath)
		} else {
			response := &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "method not supported",
				Body:       io.NopCloser(strings.NewReader("{}")),
			}
			return response, nil
		}

	case "setIamPolicy":
		if req.Method == "POST" {
			resourcePath := req.URL.Host + requestPath
			return m.iamPolicies.serveSetIAMPolicy(resourcePath, req)
		} else {
			response := &http.Response{
				StatusCode: http.StatusMethodNotAllowed,
				Status:     "method not supported",
				Body:       io.NopCloser(strings.NewReader("{}")),
			}
			return response, nil
		}

	default:
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Body:       io.NopCloser(strings.NewReader("{}")),
		}, nil
	}
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	klog.Infof("mockgcp request: %v %v", req.Method, req.URL)

	requestPath := req.URL.Path
	if strings.HasSuffix(requestPath, ":getIamPolicy") || strings.HasSuffix(requestPath, ":setIamPolicy") {
		return m.roundTripIAMPolicy(req)
	}

	var mux http.Handler
	for _, service := range m.services {
		m, found := service.MatchesHost(req.Host)
		if found {
			mux = m
			break
		}
	}
	if mux != nil {
		if err := m.prefilterRequest(req); err != nil {
			return nil, err
		}

		var body bytes.Buffer
		w := &bufferedResponseWriter{body: &body, header: make(http.Header)}
		mux.ServeHTTP(w, req)
		response := &http.Response{}
		response.Body = io.NopCloser(&body)
		response.Header = w.header
		if w.statusCode == 0 {
			w.statusCode = 200
		}
		klog.Infof("mockgcp response: %v %v => %d", req.Method, req.URL, w.statusCode)
		response.Status = fmt.Sprintf("%d %s", w.statusCode, http.StatusText(w.statusCode))
		response.StatusCode = w.statusCode
		return response, nil
	}

	request := fmt.Sprintf("%s %s", req.Method, req.URL)
	body := make(map[string]interface{})

	response := &http.Response{
		StatusCode: 403,
		Status:     "403 mockRoundTripper injecting fake response for unknown service " + req.Host,
	}

	if request == "GET https://openidconnect.googleapis.com/v1/userinfo?alt=json" {
		body["email"] = "test@example.com"
		response.StatusCode = 200
		response.Status = "200 OK"
	} else {
		klog.Errorf("host name %q not known.  "+
			"Please verify the ExpectedHost in service.go and retry.", req.Host)
	}

	if len(body) != 0 {
		j, err := json.Marshal(body)
		if err != nil {
			panic("json.Marshal failed")
		}

		log.Printf("response: %d %s", response.StatusCode, string(j))

		response.Body = io.NopCloser(bytes.NewReader(j))
	} else {
		log.Printf("response: %d %s", response.StatusCode, "-")
	}

	return response, nil
}

// bufferedResponseWriter implements http.ResponseWriter and stores the response.
type bufferedResponseWriter struct {
	statusCode int
	body       io.Writer
	header     http.Header
}

var _ http.ResponseWriter = &bufferedResponseWriter{}

// Header implements http.ResponseWriter
func (w *bufferedResponseWriter) Header() http.Header {
	return w.header
}

// Write implements http.ResponseWriter
func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.body.Write(b)
}

// WriteHeader implements http.ResponseWriter
func (w *bufferedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
