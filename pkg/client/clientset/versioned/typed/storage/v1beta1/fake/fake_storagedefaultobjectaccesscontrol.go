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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageDefaultObjectAccessControls implements StorageDefaultObjectAccessControlInterface
type FakeStorageDefaultObjectAccessControls struct {
	Fake *FakeStorageV1beta1
	ns   string
}

var storagedefaultobjectaccesscontrolsResource = schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagedefaultobjectaccesscontrols"}

var storagedefaultobjectaccesscontrolsKind = schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageDefaultObjectAccessControl"}

// Get takes name of the storageDefaultObjectAccessControl, and returns the corresponding storageDefaultObjectAccessControl object, and an error if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagedefaultobjectaccesscontrolsResource, c.ns, name), &v1beta1.StorageDefaultObjectAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageDefaultObjectAccessControl), err
}

// List takes label and field selectors, and returns the list of StorageDefaultObjectAccessControls that match those selectors.
func (c *FakeStorageDefaultObjectAccessControls) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StorageDefaultObjectAccessControlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagedefaultobjectaccesscontrolsResource, storagedefaultobjectaccesscontrolsKind, c.ns, opts), &v1beta1.StorageDefaultObjectAccessControlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.StorageDefaultObjectAccessControlList{ListMeta: obj.(*v1beta1.StorageDefaultObjectAccessControlList).ListMeta}
	for _, item := range obj.(*v1beta1.StorageDefaultObjectAccessControlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageDefaultObjectAccessControls.
func (c *FakeStorageDefaultObjectAccessControls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagedefaultobjectaccesscontrolsResource, c.ns, opts))

}

// Create takes the representation of a storageDefaultObjectAccessControl and creates it.  Returns the server's representation of the storageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Create(ctx context.Context, storageDefaultObjectAccessControl *v1beta1.StorageDefaultObjectAccessControl, opts v1.CreateOptions) (result *v1beta1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagedefaultobjectaccesscontrolsResource, c.ns, storageDefaultObjectAccessControl), &v1beta1.StorageDefaultObjectAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageDefaultObjectAccessControl), err
}

// Update takes the representation of a storageDefaultObjectAccessControl and updates it. Returns the server's representation of the storageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Update(ctx context.Context, storageDefaultObjectAccessControl *v1beta1.StorageDefaultObjectAccessControl, opts v1.UpdateOptions) (result *v1beta1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagedefaultobjectaccesscontrolsResource, c.ns, storageDefaultObjectAccessControl), &v1beta1.StorageDefaultObjectAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageDefaultObjectAccessControl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageDefaultObjectAccessControls) UpdateStatus(ctx context.Context, storageDefaultObjectAccessControl *v1beta1.StorageDefaultObjectAccessControl, opts v1.UpdateOptions) (*v1beta1.StorageDefaultObjectAccessControl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagedefaultobjectaccesscontrolsResource, "status", c.ns, storageDefaultObjectAccessControl), &v1beta1.StorageDefaultObjectAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageDefaultObjectAccessControl), err
}

// Delete takes name of the storageDefaultObjectAccessControl and deletes it. Returns an error if one occurs.
func (c *FakeStorageDefaultObjectAccessControls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagedefaultobjectaccesscontrolsResource, c.ns, name), &v1beta1.StorageDefaultObjectAccessControl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageDefaultObjectAccessControls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagedefaultobjectaccesscontrolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.StorageDefaultObjectAccessControlList{})
	return err
}

// Patch applies the patch and returns the patched storageDefaultObjectAccessControl.
func (c *FakeStorageDefaultObjectAccessControls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagedefaultobjectaccesscontrolsResource, c.ns, name, pt, data, subresources...), &v1beta1.StorageDefaultObjectAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageDefaultObjectAccessControl), err
}