// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeFutureReservationRef{}

// ComputeFutureReservationRef defines a reference to a ComputeFutureReservation resource.
type ComputeFutureReservationRef struct {
	// A reference to a ComputeFutureReservation resource.
	// +oneOf
	Name string `json:"name,omitempty"`

	// A reference to a ComputeFutureReservation resource.
	// +oneOf
	Namespace string `json:"namespace,omitempty"`

	// A reference to a ComputeFutureReservation resource.
	// +oneOf
	External string `json:"external,omitempty"`
}

func (r *ComputeFutureReservationRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on ComputeFutureReservation reference")
	}
	if r.External != "" {
		return r.External, nil
	}
	namespace := r.Namespace
	if namespace == "" {
		namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeFutureReservationGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("error getting referenced ComputeFutureReservation %q: %w", r.Name, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return actualExternalRef, nil
}
