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
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pbv1beta "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type FutureReservationsV1Beta struct {
	*MockService
	pbv1beta.UnimplementedFutureReservationsServer
}

func (s *FutureReservationsV1Beta) Get(ctx context.Context, req *pbv1beta.GetFutureReservationRequest) (*pbv1beta.FutureReservation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *FutureReservationsV1Beta) Insert(ctx context.Context, req *pbv1beta.InsertFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservationResource().GetName()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetFutureReservationResource()).(*pbv1beta.FutureReservation)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	obj.Kind = PtrTo("compute#futureReservation")
	if obj.Status == nil {
		obj.Status = &pbv1beta.FutureReservationStatus{}
	}
	if obj.PlanningStatus == nil {
		obj.PlanningStatus = PtrTo("DRAFT")
	}
	obj.Status.ProcurementStatus = PtrTo("DRAFTING")
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, name.String()))
	if obj.SpecificReservationRequired == nil {
		obj.SpecificReservationRequired = PtrTo(false)
	}
	if obj.GetAutoDeleteAutoCreatedReservations() == false {
		obj.AutoDeleteAutoCreatedReservations = nil
	}
	obj.Zone = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/beta/projects/%s/zones/%s", name.Project.ID, name.Zone))

	if obj.GetShareSettings() != nil && obj.GetShareSettings().GetShareType() == "SPECIFIC_PROJECTS" {
		if len(obj.GetShareSettings().GetProjectMap()) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "project_map is required when share_type is SPECIFIC_PROJECTS")
		}
		obj, err = s.convertProjectNumber(ctx, obj)
		if err != nil {
			return nil, err
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperationsV1Beta.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1Beta) Update(ctx context.Context, req *pbv1beta.UpdateFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	proto.Merge(obj, req.GetFutureReservationResource())
	if obj.Status != nil {
		obj.Status.ExistingMatchingUsageInfo = &pbv1beta.FutureReservationStatusExistingMatchingUsageInfo{
			Count:     PtrTo(int64(0)),
			Timestamp: PtrTo(s.nowString()),
		}
	}
	if obj.TimeWindow != nil && obj.TimeWindow.Duration != nil {
		obj.TimeWindow.Duration = nil
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperationsV1Beta.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1Beta) Cancel(ctx context.Context, req *pbv1beta.CancelFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update status to CANCELLED or similar
	if obj.Status == nil {
		obj.Status = &pbv1beta.FutureReservationStatus{}
	}
	obj.Status.ProcurementStatus = PtrTo("CANCELLED")

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("cancel"),
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperationsV1Beta.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1Beta) Delete(ctx context.Context, req *pbv1beta.DeleteFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/futureReservations/" + req.GetFutureReservation()
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperationsV1Beta.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1Beta) AggregatedList(ctx context.Context, req *pbv1beta.AggregatedListFutureReservationsRequest) (*pbv1beta.FutureReservationsAggregatedListResponse, error) {
	response := &pbv1beta.FutureReservationsAggregatedListResponse{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#futureReservationAggregatedList")
	response.SelfLink = PtrTo(BuildComputeSelfLink(ctx, "projects/"+req.GetProject()+"/aggregated/futureReservations"))

	response.Items = make(map[string]*pbv1beta.FutureReservationsScopedList)

	findKind := (&pbv1beta.FutureReservation{}).ProtoReflect().Descriptor()
	prefix := "projects/" + req.GetProject() + "/zones/"
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		fr := obj.(*pbv1beta.FutureReservation)

		selfLink := fr.GetSelfLink()
		tokens := strings.Split(selfLink, "/")
		zone := ""
		for i, token := range tokens {
			if token == "zones" && i+1 < len(tokens) {
				zone = "zones/" + tokens[i+1]
				break
			}
		}

		if zone != "" {
			if response.Items[zone] == nil {
				response.Items[zone] = &pbv1beta.FutureReservationsScopedList{}
			}
			response.Items[zone].FutureReservations = append(response.Items[zone].FutureReservations, fr)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type futureReservationName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *futureReservationName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/futureReservations/" + n.Name
}

func (s *MockService) parseFutureReservationName(name string) (*futureReservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "futureReservations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &futureReservationName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *FutureReservationsV1Beta) convertProjectNumber(ctx context.Context, obj *pbv1beta.FutureReservation) (*pbv1beta.FutureReservation, error) {
	projectMap := obj.ShareSettings.ProjectMap
	if projectMap == nil {
		return nil, nil
	}
	projects := obj.ShareSettings.Projects
	if projects == nil {
		projects = []string{}
	}
	newMap := make(map[string]*pbv1beta.ShareSettingsProjectConfig)
	for idOrNumber, config := range projectMap {
		project, err := s.Projects.GetProjectByIDOrNumber(idOrNumber)
		if err != nil {
			return nil, err
		}
		projectNumber := strconv.FormatInt(project.Number, 10)
		newConfig := proto.Clone(config).(*pbv1beta.ShareSettingsProjectConfig)
		newConfig.ProjectId = &projectNumber
		newMap[projectNumber] = newConfig
		projects = append(projects, projectNumber)
	}
	obj.ShareSettings.ProjectMap = newMap
	obj.ShareSettings.Projects = projects
	return obj, nil
}
