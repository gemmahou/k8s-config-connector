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

package mockcompute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"time"

	pbv1beta "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1beta"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
)

type computeOperationsV1Beta struct {
	storage storage.Storage
}

func (s *computeOperationsV1Beta) zonalOperationFQN(projectID string, zone string, name string) string {
	return "projects/" + projectID + "/zones/" + zone + "/operations/" + name
}

func (s *computeOperationsV1Beta) startLRO0(ctx context.Context, op *pbv1beta.Operation, fqn string, callback func() (proto.Message, error)) (*pbv1beta.Operation, error) {
	log := klog.FromContext(ctx)

	now := time.Now()
	nanos := now.UnixNano()

	if op == nil {
		op = &pbv1beta.Operation{}
	}

	op.StartTime = PtrTo(formatTime(now))
	op.InsertTime = PtrTo(formatTime(now))
	op.Id = PtrTo(uint64(nanos))

	// Specific to ComputeFirewallPolicy
	// Remove targetId and targetLink when status is RUNNING to match realGCP operation
	// ref: https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/2800/commits/32fdacd53d59c36626fce16f2b0125a8a455f3d6#r1783642429
	targetId := op.TargetId
	targetLink := op.TargetLink
	if op.OperationType != nil && *op.OperationType == "createFirewallPolicy" {
		op.TargetId = nil
		op.TargetLink = nil
	}

	if op.Progress == nil {
		op.Progress = PtrTo(int32(0))
	}

	if op.Status == nil {
		op.Status = PtrTo(pbv1beta.Operation_RUNNING)
	}

	op.Kind = PtrTo("compute#operation")
	op.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))

	log.Info("storing operation", "fqn", fqn)
	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, err
	}

	go func() {
		result, err := callback()
		finished := &pbv1beta.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		finished.Progress = PtrTo(int32(100))
		finished.Status = PtrTo(pbv1beta.Operation_DONE)
		finished.EndTime = PtrTo(formatTime(time.Now()))

		// Specific to ComputeFirewallPolicy
		// Add targetId and targetLink back when status is DONE to match realGCP operation
		if op.OperationType != nil && *op.OperationType == "createFirewallPolicy" {
			finished.TargetId = targetId
			finished.TargetLink = targetLink
		}

		if err != nil {
			code := status.Code(err)
			message := err.Error()

			finished.Error = &pbv1beta.Error{
				Errors: []*pbv1beta.Errors{
					{
						Code:    PtrTo(code.String()),
						Message: PtrTo(message),
					},
				},
			}
			klog.Warningf("TODO: more fully handle LRO error %v", err)
		} else {
			// The LRO result does not appear to be returned in the operation
			klog.V(4).Infof("LRO result: %+v", result)
		}
		if err := s.storage.Update(ctx, fqn, finished); err != nil {
			klog.Warningf("error updating LRO: %v", err)
			return
		}
	}()

	return op, nil
}

func (s *computeOperationsV1Beta) startZonalLRO(ctx context.Context, projectID string, zone string, op *pbv1beta.Operation, callback func() (proto.Message, error)) (*pbv1beta.Operation, error) {
	now := time.Now()
	millis := now.UnixMilli()
	id := uuid.NewUUID()

	name := fmt.Sprintf("operation-%d-%s", millis, id)
	fqn := s.zonalOperationFQN(projectID, zone, name)

	op.Name = PtrTo(name)
	op.Zone = PtrTo(BuildComputeSelfLink(ctx, "projects/"+projectID+"/zones/"+zone))
	return s.startLRO0(ctx, op, fqn, callback)
}

func (s *computeOperationsV1Beta) getOperation(ctx context.Context, fqn string) (*pbv1beta.Operation, error) {
	op := &pbv1beta.Operation{}
	if err := s.storage.Get(ctx, fqn, op); err != nil {
		return nil, err
	}

	return op, nil
}
