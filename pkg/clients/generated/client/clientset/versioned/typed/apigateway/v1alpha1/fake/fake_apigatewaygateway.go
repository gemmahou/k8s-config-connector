// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigateway/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIGatewayGateways implements APIGatewayGatewayInterface
type FakeAPIGatewayGateways struct {
	Fake *FakeApigatewayV1alpha1
	ns   string
}

var apigatewaygatewaysResource = schema.GroupVersionResource{Group: "apigateway.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "apigatewaygateways"}

var apigatewaygatewaysKind = schema.GroupVersionKind{Group: "apigateway.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "APIGatewayGateway"}

// Get takes name of the aPIGatewayGateway, and returns the corresponding aPIGatewayGateway object, and an error if there is any.
func (c *FakeAPIGatewayGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewaygatewaysResource, c.ns, name), &v1alpha1.APIGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIGatewayGateway), err
}

// List takes label and field selectors, and returns the list of APIGatewayGateways that match those selectors.
func (c *FakeAPIGatewayGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIGatewayGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewaygatewaysResource, apigatewaygatewaysKind, c.ns, opts), &v1alpha1.APIGatewayGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIGatewayGatewayList{ListMeta: obj.(*v1alpha1.APIGatewayGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIGatewayGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIGatewayGateways.
func (c *FakeAPIGatewayGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewaygatewaysResource, c.ns, opts))

}

// Create takes the representation of a aPIGatewayGateway and creates it.  Returns the server's representation of the aPIGatewayGateway, and an error, if there is any.
func (c *FakeAPIGatewayGateways) Create(ctx context.Context, aPIGatewayGateway *v1alpha1.APIGatewayGateway, opts v1.CreateOptions) (result *v1alpha1.APIGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewaygatewaysResource, c.ns, aPIGatewayGateway), &v1alpha1.APIGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIGatewayGateway), err
}

// Update takes the representation of a aPIGatewayGateway and updates it. Returns the server's representation of the aPIGatewayGateway, and an error, if there is any.
func (c *FakeAPIGatewayGateways) Update(ctx context.Context, aPIGatewayGateway *v1alpha1.APIGatewayGateway, opts v1.UpdateOptions) (result *v1alpha1.APIGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewaygatewaysResource, c.ns, aPIGatewayGateway), &v1alpha1.APIGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIGatewayGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIGatewayGateways) UpdateStatus(ctx context.Context, aPIGatewayGateway *v1alpha1.APIGatewayGateway, opts v1.UpdateOptions) (*v1alpha1.APIGatewayGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewaygatewaysResource, "status", c.ns, aPIGatewayGateway), &v1alpha1.APIGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIGatewayGateway), err
}

// Delete takes name of the aPIGatewayGateway and deletes it. Returns an error if one occurs.
func (c *FakeAPIGatewayGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apigatewaygatewaysResource, c.ns, name, opts), &v1alpha1.APIGatewayGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIGatewayGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewaygatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIGatewayGatewayList{})
	return err
}

// Patch applies the patch and returns the patched aPIGatewayGateway.
func (c *FakeAPIGatewayGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIGatewayGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewaygatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.APIGatewayGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIGatewayGateway), err
}