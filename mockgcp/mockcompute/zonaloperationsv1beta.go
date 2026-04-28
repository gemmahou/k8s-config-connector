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
	"time"

	pbv1beta "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1beta"
)

type ZonalOperationsV1Beta struct {
	*MockService
	pbv1beta.UnimplementedZoneOperationsServer
}

func (s *ZonalOperationsV1Beta) Get(ctx context.Context, req *pbv1beta.GetZoneOperationRequest) (*pbv1beta.Operation, error) {
	fqn := s.computeOperations.zonalOperationFQN(req.Project, req.Zone, req.Operation)
	lro, err := s.computeOperationsV1Beta.getOperation(ctx, fqn)
	if err != nil {
		return nil, err
	}

	return lro, nil
}

func (s *ZonalOperationsV1Beta) Wait(ctx context.Context, req *pbv1beta.WaitZoneOperationRequest) (*pbv1beta.Operation, error) {
	fqn := s.computeOperations.zonalOperationFQN(req.Project, req.Zone, req.Operation)

	deadline := 2 * time.Minute
	timeoutAt := time.Now().Add(deadline)
	for {
		lro, err := s.computeOperationsV1Beta.getOperation(ctx, fqn)
		if err != nil {
			return nil, err
		}
		switch ValueOf(lro.Status) {
		case pbv1beta.Operation_DONE:
			return lro, nil
		}

		if time.Now().After(timeoutAt) {
			return lro, nil
		}

		time.Sleep(2 * time.Second)
	}
}
