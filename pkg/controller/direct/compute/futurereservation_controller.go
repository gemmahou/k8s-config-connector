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

// +tool:controller
// proto.service: google.cloud.compute.v1.FutureReservations
// proto.message: google.cloud.compute.v1.FutureReservation
// crd.type: ComputeFutureReservation
// crd.version: v1alpha1

package compute

import (
	"reflect"
	"sort"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"context"
	"fmt"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeFutureReservationGVK, NewFutureReservationModel)
}

func NewFutureReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &futureReservationModel{config: config}, nil
}

var _ directbase.Model = &futureReservationModel{}

type futureReservationModel struct {
	config *config.ControllerConfig
}

func (m *futureReservationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeFutureReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFutureReservationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	futureReservationClient, err := gcpClient.newFutureReservationsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &FutureReservationAdapter{
		gcpClient: futureReservationClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *futureReservationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FutureReservationAdapter struct {
	gcpClient *compute.FutureReservationsClient
	id        *v1alpha1.FutureReservationIdentity
	desired   *krm.ComputeFutureReservation
	actual    *computepb.FutureReservation
	reader    client.Reader
}

var _ directbase.Adapter = &FutureReservationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FutureReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FutureReservation", "name", a.id)

	req := &computepb.GetFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FutureReservation %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeFutureReservationSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ID())

	req := &computepb.InsertFutureReservationRequest{
		Project:                   a.id.Parent().ProjectID,
		Zone:                      a.id.Parent().Location,
		FutureReservationResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FutureReservation %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute FutureReservation %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute FutureReservation in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting FutureReservation %s: %w", a.id, err)
	}

	status := &krm.ComputeFutureReservationStatus{}
	status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeFutureReservationSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ID())

	updateMask := fieldmaskpb.FieldMask{}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// GCP API does not have AutoDeleteAutoCreatedReservations field in the response body.
	// If it's false, it will not be populated; If it's true, AutoCreatedReservationsDeleteTime will be populated.
	if resource.GetAutoDeleteAutoCreatedReservations() == false {
		if a.actual.AutoCreatedReservationsDeleteTime != nil {
			report.AddField("auto_delete_auto_created_reservations", "true", resource.GetAutoDeleteAutoCreatedReservations())
			updateMask.Paths = append(updateMask.Paths, "auto_delete_auto_created_reservations")
		}
	} else {
		if a.actual.AutoCreatedReservationsDeleteTime == nil {
			report.AddField("auto_delete_auto_created_reservations", "false", resource.GetAutoDeleteAutoCreatedReservations())
			updateMask.Paths = append(updateMask.Paths, "auto_delete_auto_created_reservations")
		}
	}

	// handle AutoCreatedReservationsDeleteTime calculated by duration
	if resource.AutoCreatedReservationsDuration != nil {
		start, err := time.Parse(time.RFC3339Nano, direct.ValueOf(resource.TimeWindow.StartTime))
		if err != nil {
			return fmt.Errorf("invalid start_time: %w", err)
		}
		duration := time.Duration(0)
		if resource.AutoCreatedReservationsDuration.Seconds != nil {
			duration = duration + time.Duration(direct.ValueOf(resource.AutoCreatedReservationsDuration.Seconds))*time.Second
		}
		if resource.AutoCreatedReservationsDuration.Nanos != nil {
			duration = duration + time.Duration(direct.ValueOf(resource.AutoCreatedReservationsDuration.Nanos))*time.Nanosecond
		}
		end := start.Add(duration)
		if a.actual.AutoCreatedReservationsDeleteTime == nil {
			report.AddField("auto_created_reservations_duration", resource.AutoCreatedReservationsDuration, nil)
			updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_duration")
		} else if actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetAutoCreatedReservationsDeleteTime()); err == nil {
			calculatedDelete := end
			if !calculatedDelete.Equal(actualDelete) {
				report.AddField("auto_created_reservations_duration", resource.AutoCreatedReservationsDuration, nil)
				updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_duration")
			}
		} else {
			return fmt.Errorf("invalid auto_created_reservations_delete_time: %w", err)
		}
	}

	if resource.AutoCreatedReservationsDeleteTime != nil {
		actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetAutoCreatedReservationsDeleteTime())
		if err != nil {
			return fmt.Errorf("invalid auto_created_reservations_delete_time: %w", err)
		}
		desiredDelete, err := time.Parse(time.RFC3339Nano, resource.GetAutoCreatedReservationsDeleteTime())
		if err != nil {
			return fmt.Errorf("invalid auto_created_reservations_delete_time: %w", err)
		}
		if !desiredDelete.Equal(actualDelete) {
			report.AddField("auto_created_reservations_delete_time", resource.AutoCreatedReservationsDeleteTime, a.actual.AutoCreatedReservationsDeleteTime)
			updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_delete_time")
		}
	}

	if !reflect.DeepEqual(resource.TimeWindow, a.actual.TimeWindow) {
		if resource.TimeWindow != nil && a.actual.TimeWindow != nil {
			if !reflect.DeepEqual(resource.TimeWindow.StartTime, a.actual.TimeWindow.StartTime) {
				report.AddField("time_window.start_time", resource.TimeWindow.StartTime, a.actual.TimeWindow.StartTime)
				updateMask.Paths = append(updateMask.Paths, "time_window.start_time")
			}
		}
	}

	// handle endtime calculated by duration
	if resource.TimeWindow != nil && resource.TimeWindow.Duration != nil {
		start, err := time.Parse(time.RFC3339Nano, direct.ValueOf(resource.TimeWindow.StartTime))
		if err != nil {
			return fmt.Errorf("invalid start_time: %w", err)
		}
		duration := time.Duration(0)
		if resource.TimeWindow.Duration.Seconds != nil {
			duration = duration + time.Duration(direct.ValueOf(resource.TimeWindow.Duration.Seconds))*time.Second
		}
		if resource.TimeWindow.Duration.Nanos != nil {
			duration = duration + time.Duration(direct.ValueOf(resource.TimeWindow.Duration.Nanos))*time.Nanosecond
		}
		end := start.Add(duration)
		if a.actual.TimeWindow.EndTime == nil {
			report.AddField("auto_created_reservations_duration", resource.TimeWindow.EndTime, nil)
			updateMask.Paths = append(updateMask.Paths, "time_window.duration")
		} else if actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetTimeWindow().GetEndTime()); err == nil {
			calculatedDelete := end
			if !calculatedDelete.Equal(actualDelete) {
				report.AddField("time_window.duration", end, a.actual.TimeWindow.EndTime)
				updateMask.Paths = append(updateMask.Paths, "time_window.duration")
			}
		} else {
			return fmt.Errorf("invalid time_window.end_time: %w", err)
		}
	}

	if resource.TimeWindow.EndTime != nil {
		actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetTimeWindow().GetEndTime())
		if err != nil {
			return fmt.Errorf("invalid time_window.end_time: %w", err)
		}
		desiredDelete, err := time.Parse(time.RFC3339Nano, resource.GetTimeWindow().GetEndTime())
		if err != nil {
			return fmt.Errorf("invalid time_window.end_time: %w", err)
		}
		if !desiredDelete.Equal(actualDelete) {
			report.AddField("time_window.end_time", resource.TimeWindow.EndTime, a.actual.TimeWindow.EndTime)
			updateMask.Paths = append(updateMask.Paths, "time_window.end_time")
		}
	}

	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		report.AddField("description", resource.Description, a.actual.Description)
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if !reflect.DeepEqual(resource.SpecificSkuProperties, a.actual.SpecificSkuProperties) {
		if resource.SpecificSkuProperties != nil && a.actual.SpecificSkuProperties != nil {
			if !reflect.DeepEqual(resource.SpecificSkuProperties.TotalCount, a.actual.SpecificSkuProperties.TotalCount) {
				report.AddField("specific_sku_properties.total_count", resource.SpecificSkuProperties.TotalCount, a.actual.SpecificSkuProperties.TotalCount)
				updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.total_count")
			}
			if !reflect.DeepEqual(resource.SpecificSkuProperties.InstanceProperties, a.actual.SpecificSkuProperties.InstanceProperties) {
				if resource.SpecificSkuProperties.InstanceProperties != nil && a.actual.SpecificSkuProperties.GetInstanceProperties() != nil {
					if !reflect.DeepEqual(resource.SpecificSkuProperties.InstanceProperties.MachineType, a.actual.SpecificSkuProperties.InstanceProperties.MachineType) {
						report.AddField("specific_sku_properties.instance_properties.machine_type", resource.SpecificSkuProperties.InstanceProperties.MachineType, a.actual.SpecificSkuProperties.InstanceProperties.MachineType)
						updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.instance_properties.machine_type")
					}
					if !reflect.DeepEqual(resource.SpecificSkuProperties.InstanceProperties.MinCpuPlatform, a.actual.SpecificSkuProperties.InstanceProperties.MinCpuPlatform) {
						report.AddField("specific_sku_properties.instance_properties.min_cpu_platform", resource.SpecificSkuProperties.InstanceProperties.MinCpuPlatform, a.actual.SpecificSkuProperties.InstanceProperties.MinCpuPlatform)
						updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.instance_properties.min_cpu_platform")
					}
				}
			}
		}
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ComputeFutureReservationStatus{}
		status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		if a.desired.Status.ExternalRef == nil {
			// If it is the first reconciliation after switching to direct controller,
			// or is an acquisition with updates, then fill out the ExternalRef.
			status.ExternalRef = direct.LazyPtr(a.id.String())
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	sort.Strings(updateMask.Paths)
	log.Info("*****", "paths", updateMask.Paths, "resource", resource)

	req := &computepb.UpdateFutureReservationRequest{
		Project:                   a.id.Parent().ProjectID,
		Zone:                      a.id.Parent().Location,
		FutureReservation:         a.id.ID(),
		FutureReservationResource: resource,
		UpdateMask:                direct.PtrTo(strings.Join(updateMask.Paths, ",")),
	}
	op, err := a.gcpClient.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("updating compute FutureReservation %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of compute FutureReservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated compute FutureReservation", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFutureReservation %s: %w", a.id, err)
	}

	status := &krm.ComputeFutureReservationStatus{}
	status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		// If it is the first reconciliation after switching to direct controller,
		// or is an acquisition with updates, then fill out the ExternalRef.
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FutureReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeFutureReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeFutureReservationSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeFutureReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FutureReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FutureReservation", "name", a.id)

	req := &computepb.DeleteFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute FutureReservation %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute FutureReservation", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute FutureReservation %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *FutureReservationAdapter) get(ctx context.Context) (*computepb.FutureReservation, error) {
	getReq := &computepb.GetFutureReservationRequest{
		Project:           a.id.Parent().ProjectID,
		Zone:              a.id.Parent().Location,
		FutureReservation: a.id.ID(),
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeFutureReservation %s: %w", a.id, err)
	}
	return resource, nil
}
