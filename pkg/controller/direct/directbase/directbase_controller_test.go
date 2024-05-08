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

package directbase

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/google/go-cmp/cmp"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr   manager.Manager
	gvk   schema.GroupVersionKind
	model Model

	opts                       = Deps{JitterGenerator: nil}
	immediateReconcileRequests = make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines    = semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
)

func TestReconcile_DeleteDefender(t *testing.T) {
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

	reconciler, err := NewReconciler(mgr, immediateReconcileRequests, resourceWatcherRoutines, gvk, model, opts.JitterGenerator)
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

func newTestKindUnstructured(nn types.NamespacedName) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1beta1/test.cnrm.cloud.google.com",
			"kind":       "TestKind",
			"metadata": map[string]interface{}{
				"namespace": nn.Namespace,
				"name":      nn.Name,
				// deletion timestamp
			},
		},
	}
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
