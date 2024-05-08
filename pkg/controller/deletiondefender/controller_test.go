// Copyright 2024 Google LLC
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

package deletiondefender_test

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/deletiondefender"
	"reflect"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/google/go-cmp/cmp"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr manager.Manager

	fakeCRD            = newTestKindCRD()
	fakeCRDForDeletion = newTestKindCRDForDeletion()
)

// Resource is being deleted, remove the deletion defender finalizer and allow the controller to delete the underlying
// resource on GCP.
func TestReconcile_DeleteResource(t *testing.T) {
	t.Parallel()
	testID := testvariable.NewUniqueID()
	client := mgr.GetClient()
	testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
	testcontroller.EnsureNamespaceExistsT(t, client, testID)

	resourceNN := types.NamespacedName{
		Namespace: testID,
		Name:      testID,
	}
	resource := newTestKindUnstructured(resourceNN)
	resource.SetDeletionTimestamp(&metav1.Time{Time: time.Now()})
	test.EnsureObjectExists(t, resource, client)

	reconciler, err := deletiondefender.NewReconciler(mgr, fakeCRD)
	if err != nil {
		t.Fatal(fmt.Errorf("error creating reconciler: %w", err))
	}
	res, err := reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: resourceNN})
	if err != nil {
		t.Fatal(fmt.Errorf("unexpected error during reconciliation: %w", err))
	}
	emptyResult := reconcile.Result{}
	if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
	}

	found, err := hasDeletionDefenderFinalizer(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's deletion defender finalizer: %w", err))
	}
	if found {
		t.Fatalf("unexpected deletion defender finalizer. It should be removed while deleting the resource.")
	}
}

// Resource is being deleted, and its CRD is being deleted, remove both KCC finalizers and set the resource to abandon
func TestReconcile_DeleteResourceAndDeleteCRD(t *testing.T) {
	t.Parallel()
	testID := testvariable.NewUniqueID()
	client := mgr.GetClient()
	testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
	testcontroller.EnsureNamespaceExistsT(t, client, testID)

	resourceNN := types.NamespacedName{
		Namespace: testID,
		Name:      testID,
	}
	resource := newTestKindUnstructured(resourceNN)
	test.EnsureObjectExists(t, resource, client)

	reconciler, err := deletiondefender.NewReconciler(mgr, fakeCRDForDeletion)
	if err != nil {
		t.Fatal(fmt.Errorf("error creating reconciler: %w", err))
	}
	res, err := reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: resourceNN})
	if err != nil {
		t.Fatal(fmt.Errorf("unexpected error during reconciliation: %w", err))
	}
	emptyResult := reconcile.Result{}
	if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
	}

	found, err := hasDeletionDefenderFinalizer(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's deletion defender finalizer: %w", err))
	}
	if found {
		t.Fatalf("unexpected deletion defender finalizer. It should be removed while deleting the resource.")
	}

	found, err = hasControllerFinalizer(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's controller finalizer: %w", err))
	}
	if found {
		t.Fatalf("unexpected controller finalizer. It should be removed while deleting the resource.")
	}

	annotation, found, err := getCurrentDeletionPolicyAnnotation(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's deletion policy annotation: %w", err))
	}
	if !found {
		t.Fatalf("got nil deletion policy annotation for resource")
	}
	if gotAnnotation, wantAnnotation := annotation, "abandon"; gotAnnotation != wantAnnotation {
		t.Fatalf("got condition with reason '%v' for resource, want condition with reason '%v'", gotAnnotation, wantAnnotation)
	}
}

func newTestKindCRD() *apiextensions.CustomResourceDefinition {
	crd := test.CRDForGVK(metav1.GroupVersionKind{
		Group:   "test.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TestKind",
	})
	// Enable the status subresource for this CRD. This is needed to allow
	// UpdateStatus() calls to work on custom resources belonging to this CRD
	// on the API server.
	crd.Spec.Versions[0].Subresources = &apiextensions.CustomResourceSubresources{
		Status: &apiextensions.CustomResourceSubresourceStatus{},
	}

	return crd
}

func newTestKindCRDForDeletion() *apiextensions.CustomResourceDefinition {
	crd := test.CRDForGVK(metav1.GroupVersionKind{
		Group:   "test.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TestKind",
	})
	// Enable the status subresource for this CRD. This is needed to allow
	// UpdateStatus() calls to work on custom resources belonging to this CRD
	// on the API server.
	crd.Spec.Versions[0].Subresources = &apiextensions.CustomResourceSubresources{
		Status: &apiextensions.CustomResourceSubresourceStatus{},
	}
	// deletion timestamp not working
	//now := metav1.Now()
	//crd.SetDeletionTimestamp(&now)
	return crd
}

func newTestKindUnstructured(nn types.NamespacedName) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": fmt.Sprintf("%v/%v", fakeCRD.Spec.Group, k8s.GetVersionFromCRD(fakeCRD)),
			"kind":       fakeCRD.Spec.Names.Kind,
			"metadata": map[string]interface{}{
				"namespace":  nn.Namespace,
				"name":       nn.Name,
				"finalizers": []string{"cnrm.cloud.google.com/finalizer", "cnrm.cloud.google.com/deletion-defender"},
				// deletion timestamp
			},
		},
	}
}

func getCurrentDeletionPolicyAnnotation(ctx context.Context, c client.Client, u *unstructured.Unstructured) (annotation string, found bool, err error) {
	nn := k8s.GetNamespacedName(u)
	unstruct := &unstructured.Unstructured{}
	unstruct.SetGroupVersionKind(u.GroupVersionKind())
	if err := c.Get(ctx, nn, unstruct); err != nil {
		return "", false, fmt.Errorf("error getting resource from API server: %w", err)
	}
	resource, err := k8s.NewResource(unstruct)
	if err != nil {
		return "", false, fmt.Errorf("error marhsalling unstruct to k8s resource: %w", err)
	}
	annotation, found = k8s.GetAnnotation(k8s.DeletionPolicyAnnotation, resource)
	return annotation, found, nil
}

func hasDeletionDefenderFinalizer(ctx context.Context, c client.Client, u *unstructured.Unstructured) (found bool, err error) {
	nn := k8s.GetNamespacedName(u)
	unstruct := &unstructured.Unstructured{}
	unstruct.SetGroupVersionKind(u.GroupVersionKind())
	if err := c.Get(ctx, nn, unstruct); err != nil {
		return false, fmt.Errorf("error getting resource from API server: %w", err)
	}
	resource, err := k8s.NewResource(unstruct)
	if err != nil {
		return false, fmt.Errorf("error marhsalling unstruct to k8s resource: %w", err)
	}
	found = k8s.HasFinalizer(resource, k8s.DeletionDefenderFinalizerName)
	return found, nil
}

func hasControllerFinalizer(ctx context.Context, c client.Client, u *unstructured.Unstructured) (found bool, err error) {
	nn := k8s.GetNamespacedName(u)
	unstruct := &unstructured.Unstructured{}
	unstruct.SetGroupVersionKind(u.GroupVersionKind())
	if err := c.Get(ctx, nn, unstruct); err != nil {
		return false, fmt.Errorf("error getting resource from API server: %w", err)
	}
	resource, err := k8s.NewResource(unstruct)
	if err != nil {
		return false, fmt.Errorf("error marhsalling unstruct to k8s resource: %w", err)
	}
	found = k8s.HasFinalizer(resource, k8s.ControllerFinalizerName)
	return found, nil
}

func TestMain(m *testing.M) {
	testmain.ForUnitTestsWithCRDs(m, []*apiextensions.CustomResourceDefinition{fakeCRD}, &mgr)
}
