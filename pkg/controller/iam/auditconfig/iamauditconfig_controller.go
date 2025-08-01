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

package auditconfig

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	kontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourcewatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "iamauditconfig-controller"

var logger = log.Log.WithName(controllerName)

func Add(mgr manager.Manager, deps *kontroller.Deps) error {
	if deps.JitterGen == nil {
		var dclML metadata.ServiceMetadataLoader
		if deps.DCLConverter != nil {
			dclML = deps.DCLConverter.MetadataLoader
		}
		deps.JitterGen = jitter.NewDefaultGenerator(deps.TFLoader, dclML)
	}

	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	reconciler, err := NewReconciler(mgr, deps.TFProvider, deps.TFLoader, deps.DCLConverter, deps.DCLConfig, immediateReconcileRequests, resourceWatcherRoutines, deps.Defaulters, deps.JitterGen)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

func NewReconciler(mgr manager.Manager, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, converter *conversion.Converter, dclConfig *mmdcl.Config, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted, defaulters []k8s.Defaulter, jg jitter.Generator) (*Reconciler, error) {
	r := Reconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		Client:                     mgr.GetClient(),
		iamClient:                  kcciamclient.New(provider, smLoader, mgr.GetClient(), converter, dclConfig).TFIAMClient,
		scheme:                     mgr.GetScheme(),
		config:                     mgr.GetConfig(),
		defaulters:                 defaulters,
		immediateReconcileRequests: immediateReconcileRequests,
		resourceWatcherRoutines:    resourceWatcherRoutines,
		jitterGen:                  jg,
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *Reconciler) error {
	obj := &iamv1beta1.IAMAuditConfig{}
	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		WatchesRawSource(source.TypedChannel(r.immediateReconcileRequests, &handler.EnqueueRequestForObject{})).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicate.UnderlyingResourceOutOfSyncPredicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	return nil
}

var _ reconcile.Reconciler = &Reconciler{}

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	client.Client
	metrics.ReconcilerMetrics
	iamClient  *kcciamclient.TFIAMClient
	scheme     *runtime.Scheme
	config     *rest.Config
	defaulters []k8s.Defaulter
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies

	jitterGen jitter.Generator
}

type reconcileContext struct {
	Reconciler     *Reconciler
	Ctx            context.Context
	NamespacedName types.NamespacedName
}

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
	logger.Info("Starting reconcile", "resource", request.NamespacedName)
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, k8s.ReconcileDeadline)
	defer cancel()
	r.RecordReconcileWorkers(ctx, iamv1beta1.IAMAuditConfigGVK)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, iamv1beta1.IAMAuditConfigGVK, request.Namespace, request.Name, startTime, &err)

	var auditConfig iamv1beta1.IAMAuditConfig
	if err := r.Get(context.TODO(), request.NamespacedName, &auditConfig); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("resource not found in API server; finishing reconcile", "resource", request.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	// r.Get() overrides the TypeMeta to empty value, so need to configure it
	// after r.Get().
	auditConfig.SetGroupVersionKind(iamv1beta1.IAMAuditConfigGVK)
	if err := r.handleDefaults(ctx, &auditConfig); err != nil {
		return reconcile.Result{}, fmt.Errorf("error handling default values for IAM policy '%v': %w", k8s.GetNamespacedName(&auditConfig), err)
	}
	reconcileContext := &reconcileContext{
		Reconciler:     r,
		Ctx:            ctx,
		NamespacedName: request.NamespacedName,
	}
	requeue, err := reconcileContext.doReconcile(&auditConfig)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod, err := r.jitterGen.JitteredReenqueue(iamv1beta1.IAMPolicyGVK, &auditConfig)
	if err != nil {
		return reconcile.Result{}, err
	}
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *Reconciler) handleDefaults(ctx context.Context, auditConfig *iamv1beta1.IAMAuditConfig) error {
	changeCount := 0
	for _, defaulter := range r.defaulters {
		changed, err := defaulter.ApplyDefaults(ctx, k8s.ReconcilerTypeIAMAuditConfig, auditConfig)
		if err != nil {
			return err
		}
		if changed {
			changeCount++
		}
	}
	if changeCount > 0 {
		if err := r.Update(ctx, auditConfig); err != nil {
			return fmt.Errorf("applying update after setting defaults: %w", err)
		}
	}
	return nil
}

func (r *reconcileContext) doReconcile(auditConfig *iamv1beta1.IAMAuditConfig) (requeue bool, err error) {
	defer execution.RecoverWithInternalError(&err)

	cc, ccc, err := kccstate.FetchLiveKCCState(r.Ctx, r.Reconciler.Client, r.NamespacedName)
	if err != nil {
		return true, err
	}

	am := resourceactuation.DecideActuationMode(cc, ccc)
	switch am {
	case v1beta1.Reconciling:
		logger.V(2).Info("Actuating a resource as actuation mode is \"Reconciling\"", "resource", r.NamespacedName)
	case v1beta1.Paused:
		logger.Info("Skipping actuation of resource as actuation mode is \"Paused\"", "resource", r.NamespacedName)

		// add finalizers for deletion defender to make sure we don't delete cloud provider resources when uninstalling
		if auditConfig.GetDeletionTimestamp().IsZero() {
			k8s.EnsureFinalizers(auditConfig, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName)
		}

		return false, nil
	default:
		return false, fmt.Errorf("unknown actuation mode %v", am)
	}

	if !auditConfig.DeletionTimestamp.IsZero() {
		if !k8s.HasFinalizer(auditConfig, k8s.ControllerFinalizerName) {
			// Resource has no controller finalizer; no finalization necessary
			return false, nil
		}
		if k8s.HasFinalizer(auditConfig, k8s.DeletionDefenderFinalizerName) {
			// Deletion defender has not yet been finalized; requeuing
			logger.Info("deletion defender has not yet been finalized; requeuing", "resource", k8s.GetNamespacedName(auditConfig))
			return true, nil
		}
		if !k8s.HasAbandonAnnotation(auditConfig) {
			if err := r.Reconciler.iamClient.DeleteAuditConfig(r.Ctx, auditConfig); err != nil {
				if !errors.Is(err, kcciamclient.ErrNotFound) && !k8s.IsReferenceNotFoundError(err) {
					if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
						logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
						resource, err := ToK8sResource(auditConfig)
						if err != nil {
							return false, fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling unresolvable dependencies event: %w", err)
						}
						// Requeue resource for reconciliation with exponential backoff applied
						return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, unwrappedErr)
					}
					return false, r.handleDeleteFailed(auditConfig, err)
				}
			}
		}
		return false, r.handleDeleted(auditConfig)
	}
	if _, err := r.Reconciler.iamClient.GetAuditConfig(r.Ctx, auditConfig); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
			return r.handleUnresolvableDeps(auditConfig, unwrappedErr)
		}
		if !errors.Is(err, kcciamclient.ErrNotFound) {
			return false, r.handleUpdateFailed(auditConfig, err)
		}
	}
	if !k8s.EnsureFinalizers(auditConfig, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName) {
		if err := r.update(auditConfig); err != nil {
			return false, r.handleUpdateFailed(auditConfig, err)
		}
	}
	if _, err := r.Reconciler.iamClient.SetAuditConfig(r.Ctx, auditConfig); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
			return r.handleUnresolvableDeps(auditConfig, unwrappedErr)
		}
		return false, r.handleUpdateFailed(auditConfig, fmt.Errorf("error setting audit config: %w", err))
	}
	if isAPIServerUpdateRequired(auditConfig) {
		return false, r.handleUpToDate(auditConfig)
	}
	return false, nil
}

func (r *reconcileContext) update(auditConfig *iamv1beta1.IAMAuditConfig) error {
	if err := r.Reconciler.Client.Update(r.Ctx, auditConfig); err != nil {
		return fmt.Errorf("error updating '%v' in API server: %w", r.NamespacedName, err)
	}
	return nil
}

func (r *reconcileContext) handleUpToDate(auditConfig *iamv1beta1.IAMAuditConfig) error {
	resource, err := ToK8sResource(auditConfig)
	if err != nil {
		return fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	}
	return r.Reconciler.HandleUpToDate(r.Ctx, resource)
}

func (r *reconcileContext) handleUpdateFailed(auditConfig *iamv1beta1.IAMAuditConfig, origErr error) error {
	resource, err := ToK8sResource(auditConfig)
	if err != nil {
		logger.Error(err, "error converting IAMAuditConfig to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(auditConfig), "event", k8s.UpdateFailed)
		return fmt.Errorf("Update call failed: %w", origErr)
	}
	return r.Reconciler.HandleUpdateFailed(r.Ctx, resource, origErr)
}

func (r *reconcileContext) handleDeleted(auditConfig *iamv1beta1.IAMAuditConfig) error {
	resource, err := ToK8sResource(auditConfig)
	if err != nil {
		return fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling %v event: %w", k8s.Deleted, err)
	}
	return r.Reconciler.HandleDeleted(r.Ctx, resource)
}

func (r *reconcileContext) handleDeleteFailed(auditConfig *iamv1beta1.IAMAuditConfig, origErr error) error {
	resource, err := ToK8sResource(auditConfig)
	if err != nil {
		logger.Error(err, "error converting IAMAuditConfig to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(auditConfig), "event", k8s.DeleteFailed)
		return fmt.Errorf(k8s.DeleteFailedMessageTmpl, origErr)
	}
	return r.Reconciler.HandleDeleteFailed(r.Ctx, resource, origErr)
}

func (r *Reconciler) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *reconcileContext) handleUnresolvableDeps(auditConfig *iamv1beta1.IAMAuditConfig, origErr error) (requeue bool, err error) {
	resource, err := ToK8sResource(auditConfig)
	if err != nil {
		return false, fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling unresolvable dependencies event: %w", err)
	}
	refGVK, refNN, ok := lifecyclehandler.CausedByUnreadyOrNonexistentResourceRefs(origErr)
	if !ok || !r.Reconciler.supportsImmediateReconciliations() {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
	}
	// Check that the number of active resource watches
	// does not exceed the controller's cap. If the
	// capacity is not exceeded, The number of active
	// resource watches is incremented by one and a watch
	// is started
	if !r.Reconciler.resourceWatcherRoutines.TryAcquire(1) {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
	}
	// Create a logger for ResourceWatcher that contains info
	// about the referencing resource. This is done since the
	// messages logged by ResourceWatcher only include the
	// information of the resource it is watching by default.
	watcherLogger := logger.WithValues(
		"referencingResource", resource.GetNamespacedName(),
		"referencingResourceGVK", resource.GroupVersionKind())
	watcher, err := resourcewatcher.New(r.Reconciler.config, watcherLogger)
	if err != nil {
		return false, r.Reconciler.HandleUpdateFailed(r.Ctx, resource, fmt.Errorf("error initializing new resourcewatcher: %w", err))
	}

	logger := logger.WithValues(
		"resource", resource.GetNamespacedName(),
		"resourceGVK", resource.GroupVersionKind(),
		"reference", refNN,
		"referenceGVK", refGVK)
	go func() {
		// Decrement the count of active resource watches after
		// the watch finishes
		defer r.Reconciler.resourceWatcherRoutines.Release(1)
		timeoutPeriod := r.Reconciler.jitterGen.WatchJitteredTimeout()
		ctx, cancel := context.WithTimeout(context.TODO(), timeoutPeriod)
		defer cancel()
		logger.Info("starting wait with timeout on resource's reference", "timeout", timeoutPeriod)
		if err := watcher.WaitForResourceToBeReadyOrDeleted(ctx, refNN, refGVK); err != nil {
			logger.Error(err, "error while waiting for resource's reference to be ready")
			return
		}
		logger.Info("enqueuing resource for immediate reconciliation now that its reference is ready")
		r.Reconciler.enqueueForImmediateReconciliation(resource.GetNamespacedName())
	}()

	// Do not requeue resource for immediate reconciliation. Wait for either
	// the next periodic reconciliation or for the referenced resource to be ready (which
	// triggers a reconciliation), whichever comes first.
	return false, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
}

// enqueueForImmediateReconciliation enqueues the given resource for immediate
// reconciliation. Note that this function only takes in the name and namespace
// of the resource and not its GVK since the controller instance that this
// reconcile instance belongs to can only reconcile resources of one GVK.
func (r *Reconciler) enqueueForImmediateReconciliation(resourceNN types.NamespacedName) {
	genEvent := event.GenericEvent{}
	genEvent.Object = &unstructured.Unstructured{}
	genEvent.Object.SetNamespace(resourceNN.Namespace)
	genEvent.Object.SetName(resourceNN.Name)
	r.immediateReconcileRequests <- genEvent
}

func isAPIServerUpdateRequired(auditConfig *iamv1beta1.IAMAuditConfig) bool {
	// TODO: even in the event of an actual update to GCP, this function will
	// return false because the condition comparison doesn't account for time.
	conditions := []condition.Condition{
		k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
	}
	if !k8s.ConditionSlicesEqual(auditConfig.Status.Conditions, conditions) {
		return true
	}
	if auditConfig.Status.ObservedGeneration != auditConfig.GetGeneration() {
		return true
	}
	return false
}

func ToK8sResource(auditConfig *iamv1beta1.IAMAuditConfig) (*k8s.Resource, error) {
	kcciamclient.SetGVK(auditConfig)
	resource := k8s.Resource{}
	if err := util.Marshal(auditConfig, &resource); err != nil {
		return nil, fmt.Errorf("error marshalling IAMAuditConfig to k8s resource: %w", err)
	}
	return &resource, nil
}
