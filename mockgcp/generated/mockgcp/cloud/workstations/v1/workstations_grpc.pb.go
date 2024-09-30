// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/workstations/v1/workstations.proto

package workstationspb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkstationsClient is the client API for Workstations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkstationsClient interface {
	// Returns the requested workstation cluster.
	GetWorkstationCluster(ctx context.Context, in *GetWorkstationClusterRequest, opts ...grpc.CallOption) (*WorkstationCluster, error)
	// Returns all workstation clusters in the specified location.
	ListWorkstationClusters(ctx context.Context, in *ListWorkstationClustersRequest, opts ...grpc.CallOption) (*ListWorkstationClustersResponse, error)
	// Creates a new workstation cluster.
	CreateWorkstationCluster(ctx context.Context, in *CreateWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an existing workstation cluster.
	UpdateWorkstationCluster(ctx context.Context, in *UpdateWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes the specified workstation cluster.
	DeleteWorkstationCluster(ctx context.Context, in *DeleteWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns the requested workstation configuration.
	GetWorkstationConfig(ctx context.Context, in *GetWorkstationConfigRequest, opts ...grpc.CallOption) (*WorkstationConfig, error)
	// Returns all workstation configurations in the specified cluster.
	ListWorkstationConfigs(ctx context.Context, in *ListWorkstationConfigsRequest, opts ...grpc.CallOption) (*ListWorkstationConfigsResponse, error)
	// Returns all workstation configurations in the specified cluster on which
	// the caller has the "workstations.workstation.create" permission.
	ListUsableWorkstationConfigs(ctx context.Context, in *ListUsableWorkstationConfigsRequest, opts ...grpc.CallOption) (*ListUsableWorkstationConfigsResponse, error)
	// Creates a new workstation configuration.
	CreateWorkstationConfig(ctx context.Context, in *CreateWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an existing workstation configuration.
	UpdateWorkstationConfig(ctx context.Context, in *UpdateWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes the specified workstation configuration.
	DeleteWorkstationConfig(ctx context.Context, in *DeleteWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns the requested workstation.
	GetWorkstation(ctx context.Context, in *GetWorkstationRequest, opts ...grpc.CallOption) (*Workstation, error)
	// Returns all Workstations using the specified workstation configuration.
	ListWorkstations(ctx context.Context, in *ListWorkstationsRequest, opts ...grpc.CallOption) (*ListWorkstationsResponse, error)
	// Returns all workstations using the specified workstation configuration
	// on which the caller has the "workstations.workstations.use" permission.
	ListUsableWorkstations(ctx context.Context, in *ListUsableWorkstationsRequest, opts ...grpc.CallOption) (*ListUsableWorkstationsResponse, error)
	// Creates a new workstation.
	CreateWorkstation(ctx context.Context, in *CreateWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an existing workstation.
	UpdateWorkstation(ctx context.Context, in *UpdateWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes the specified workstation.
	DeleteWorkstation(ctx context.Context, in *DeleteWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Starts running a workstation so that users can connect to it.
	StartWorkstation(ctx context.Context, in *StartWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Stops running a workstation, reducing costs.
	StopWorkstation(ctx context.Context, in *StopWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns a short-lived credential that can be used to send authenticated and
	// authorized traffic to a workstation.
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
}

type workstationsClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkstationsClient(cc grpc.ClientConnInterface) WorkstationsClient {
	return &workstationsClient{cc}
}

func (c *workstationsClient) GetWorkstationCluster(ctx context.Context, in *GetWorkstationClusterRequest, opts ...grpc.CallOption) (*WorkstationCluster, error) {
	out := new(WorkstationCluster)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstationCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) ListWorkstationClusters(ctx context.Context, in *ListWorkstationClustersRequest, opts ...grpc.CallOption) (*ListWorkstationClustersResponse, error) {
	out := new(ListWorkstationClustersResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstationClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) CreateWorkstationCluster(ctx context.Context, in *CreateWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstationCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) UpdateWorkstationCluster(ctx context.Context, in *UpdateWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstationCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) DeleteWorkstationCluster(ctx context.Context, in *DeleteWorkstationClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstationCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) GetWorkstationConfig(ctx context.Context, in *GetWorkstationConfigRequest, opts ...grpc.CallOption) (*WorkstationConfig, error) {
	out := new(WorkstationConfig)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstationConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) ListWorkstationConfigs(ctx context.Context, in *ListWorkstationConfigsRequest, opts ...grpc.CallOption) (*ListWorkstationConfigsResponse, error) {
	out := new(ListWorkstationConfigsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstationConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) ListUsableWorkstationConfigs(ctx context.Context, in *ListUsableWorkstationConfigsRequest, opts ...grpc.CallOption) (*ListUsableWorkstationConfigsResponse, error) {
	out := new(ListUsableWorkstationConfigsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/ListUsableWorkstationConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) CreateWorkstationConfig(ctx context.Context, in *CreateWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstationConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) UpdateWorkstationConfig(ctx context.Context, in *UpdateWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstationConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) DeleteWorkstationConfig(ctx context.Context, in *DeleteWorkstationConfigRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstationConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) GetWorkstation(ctx context.Context, in *GetWorkstationRequest, opts ...grpc.CallOption) (*Workstation, error) {
	out := new(Workstation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) ListWorkstations(ctx context.Context, in *ListWorkstationsRequest, opts ...grpc.CallOption) (*ListWorkstationsResponse, error) {
	out := new(ListWorkstationsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) ListUsableWorkstations(ctx context.Context, in *ListUsableWorkstationsRequest, opts ...grpc.CallOption) (*ListUsableWorkstationsResponse, error) {
	out := new(ListUsableWorkstationsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/ListUsableWorkstations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) CreateWorkstation(ctx context.Context, in *CreateWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) UpdateWorkstation(ctx context.Context, in *UpdateWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) DeleteWorkstation(ctx context.Context, in *DeleteWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) StartWorkstation(ctx context.Context, in *StartWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/StartWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) StopWorkstation(ctx context.Context, in *StopWorkstationRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/StopWorkstation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workstationsClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workstations.v1.Workstations/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkstationsServer is the server API for Workstations service.
// All implementations must embed UnimplementedWorkstationsServer
// for forward compatibility
type WorkstationsServer interface {
	// Returns the requested workstation cluster.
	GetWorkstationCluster(context.Context, *GetWorkstationClusterRequest) (*WorkstationCluster, error)
	// Returns all workstation clusters in the specified location.
	ListWorkstationClusters(context.Context, *ListWorkstationClustersRequest) (*ListWorkstationClustersResponse, error)
	// Creates a new workstation cluster.
	CreateWorkstationCluster(context.Context, *CreateWorkstationClusterRequest) (*longrunningpb.Operation, error)
	// Updates an existing workstation cluster.
	UpdateWorkstationCluster(context.Context, *UpdateWorkstationClusterRequest) (*longrunningpb.Operation, error)
	// Deletes the specified workstation cluster.
	DeleteWorkstationCluster(context.Context, *DeleteWorkstationClusterRequest) (*longrunningpb.Operation, error)
	// Returns the requested workstation configuration.
	GetWorkstationConfig(context.Context, *GetWorkstationConfigRequest) (*WorkstationConfig, error)
	// Returns all workstation configurations in the specified cluster.
	ListWorkstationConfigs(context.Context, *ListWorkstationConfigsRequest) (*ListWorkstationConfigsResponse, error)
	// Returns all workstation configurations in the specified cluster on which
	// the caller has the "workstations.workstation.create" permission.
	ListUsableWorkstationConfigs(context.Context, *ListUsableWorkstationConfigsRequest) (*ListUsableWorkstationConfigsResponse, error)
	// Creates a new workstation configuration.
	CreateWorkstationConfig(context.Context, *CreateWorkstationConfigRequest) (*longrunningpb.Operation, error)
	// Updates an existing workstation configuration.
	UpdateWorkstationConfig(context.Context, *UpdateWorkstationConfigRequest) (*longrunningpb.Operation, error)
	// Deletes the specified workstation configuration.
	DeleteWorkstationConfig(context.Context, *DeleteWorkstationConfigRequest) (*longrunningpb.Operation, error)
	// Returns the requested workstation.
	GetWorkstation(context.Context, *GetWorkstationRequest) (*Workstation, error)
	// Returns all Workstations using the specified workstation configuration.
	ListWorkstations(context.Context, *ListWorkstationsRequest) (*ListWorkstationsResponse, error)
	// Returns all workstations using the specified workstation configuration
	// on which the caller has the "workstations.workstations.use" permission.
	ListUsableWorkstations(context.Context, *ListUsableWorkstationsRequest) (*ListUsableWorkstationsResponse, error)
	// Creates a new workstation.
	CreateWorkstation(context.Context, *CreateWorkstationRequest) (*longrunningpb.Operation, error)
	// Updates an existing workstation.
	UpdateWorkstation(context.Context, *UpdateWorkstationRequest) (*longrunningpb.Operation, error)
	// Deletes the specified workstation.
	DeleteWorkstation(context.Context, *DeleteWorkstationRequest) (*longrunningpb.Operation, error)
	// Starts running a workstation so that users can connect to it.
	StartWorkstation(context.Context, *StartWorkstationRequest) (*longrunningpb.Operation, error)
	// Stops running a workstation, reducing costs.
	StopWorkstation(context.Context, *StopWorkstationRequest) (*longrunningpb.Operation, error)
	// Returns a short-lived credential that can be used to send authenticated and
	// authorized traffic to a workstation.
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	mustEmbedUnimplementedWorkstationsServer()
}

// UnimplementedWorkstationsServer must be embedded to have forward compatible implementations.
type UnimplementedWorkstationsServer struct {
}

func (UnimplementedWorkstationsServer) GetWorkstationCluster(context.Context, *GetWorkstationClusterRequest) (*WorkstationCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkstationCluster not implemented")
}
func (UnimplementedWorkstationsServer) ListWorkstationClusters(context.Context, *ListWorkstationClustersRequest) (*ListWorkstationClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkstationClusters not implemented")
}
func (UnimplementedWorkstationsServer) CreateWorkstationCluster(context.Context, *CreateWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkstationCluster not implemented")
}
func (UnimplementedWorkstationsServer) UpdateWorkstationCluster(context.Context, *UpdateWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkstationCluster not implemented")
}
func (UnimplementedWorkstationsServer) DeleteWorkstationCluster(context.Context, *DeleteWorkstationClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkstationCluster not implemented")
}
func (UnimplementedWorkstationsServer) GetWorkstationConfig(context.Context, *GetWorkstationConfigRequest) (*WorkstationConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkstationConfig not implemented")
}
func (UnimplementedWorkstationsServer) ListWorkstationConfigs(context.Context, *ListWorkstationConfigsRequest) (*ListWorkstationConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkstationConfigs not implemented")
}
func (UnimplementedWorkstationsServer) ListUsableWorkstationConfigs(context.Context, *ListUsableWorkstationConfigsRequest) (*ListUsableWorkstationConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsableWorkstationConfigs not implemented")
}
func (UnimplementedWorkstationsServer) CreateWorkstationConfig(context.Context, *CreateWorkstationConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkstationConfig not implemented")
}
func (UnimplementedWorkstationsServer) UpdateWorkstationConfig(context.Context, *UpdateWorkstationConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkstationConfig not implemented")
}
func (UnimplementedWorkstationsServer) DeleteWorkstationConfig(context.Context, *DeleteWorkstationConfigRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkstationConfig not implemented")
}
func (UnimplementedWorkstationsServer) GetWorkstation(context.Context, *GetWorkstationRequest) (*Workstation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) ListWorkstations(context.Context, *ListWorkstationsRequest) (*ListWorkstationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkstations not implemented")
}
func (UnimplementedWorkstationsServer) ListUsableWorkstations(context.Context, *ListUsableWorkstationsRequest) (*ListUsableWorkstationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsableWorkstations not implemented")
}
func (UnimplementedWorkstationsServer) CreateWorkstation(context.Context, *CreateWorkstationRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) UpdateWorkstation(context.Context, *UpdateWorkstationRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) DeleteWorkstation(context.Context, *DeleteWorkstationRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) StartWorkstation(context.Context, *StartWorkstationRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) StopWorkstation(context.Context, *StopWorkstationRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWorkstation not implemented")
}
func (UnimplementedWorkstationsServer) GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAccessToken not implemented")
}
func (UnimplementedWorkstationsServer) mustEmbedUnimplementedWorkstationsServer() {}

// UnsafeWorkstationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkstationsServer will
// result in compilation errors.
type UnsafeWorkstationsServer interface {
	mustEmbedUnimplementedWorkstationsServer()
}

func RegisterWorkstationsServer(s grpc.ServiceRegistrar, srv WorkstationsServer) {
	s.RegisterService(&Workstations_ServiceDesc, srv)
}

func _Workstations_GetWorkstationCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkstationClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).GetWorkstationCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstationCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).GetWorkstationCluster(ctx, req.(*GetWorkstationClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_ListWorkstationClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkstationClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).ListWorkstationClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstationClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).ListWorkstationClusters(ctx, req.(*ListWorkstationClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_CreateWorkstationCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkstationClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).CreateWorkstationCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstationCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).CreateWorkstationCluster(ctx, req.(*CreateWorkstationClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_UpdateWorkstationCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkstationClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).UpdateWorkstationCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstationCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).UpdateWorkstationCluster(ctx, req.(*UpdateWorkstationClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_DeleteWorkstationCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkstationClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).DeleteWorkstationCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstationCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).DeleteWorkstationCluster(ctx, req.(*DeleteWorkstationClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_GetWorkstationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkstationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).GetWorkstationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstationConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).GetWorkstationConfig(ctx, req.(*GetWorkstationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_ListWorkstationConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkstationConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).ListWorkstationConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstationConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).ListWorkstationConfigs(ctx, req.(*ListWorkstationConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_ListUsableWorkstationConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsableWorkstationConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).ListUsableWorkstationConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/ListUsableWorkstationConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).ListUsableWorkstationConfigs(ctx, req.(*ListUsableWorkstationConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_CreateWorkstationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkstationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).CreateWorkstationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstationConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).CreateWorkstationConfig(ctx, req.(*CreateWorkstationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_UpdateWorkstationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkstationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).UpdateWorkstationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstationConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).UpdateWorkstationConfig(ctx, req.(*UpdateWorkstationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_DeleteWorkstationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkstationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).DeleteWorkstationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstationConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).DeleteWorkstationConfig(ctx, req.(*DeleteWorkstationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_GetWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).GetWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/GetWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).GetWorkstation(ctx, req.(*GetWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_ListWorkstations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkstationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).ListWorkstations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/ListWorkstations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).ListWorkstations(ctx, req.(*ListWorkstationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_ListUsableWorkstations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsableWorkstationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).ListUsableWorkstations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/ListUsableWorkstations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).ListUsableWorkstations(ctx, req.(*ListUsableWorkstationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_CreateWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).CreateWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/CreateWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).CreateWorkstation(ctx, req.(*CreateWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_UpdateWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).UpdateWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/UpdateWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).UpdateWorkstation(ctx, req.(*UpdateWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_DeleteWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).DeleteWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/DeleteWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).DeleteWorkstation(ctx, req.(*DeleteWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_StartWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).StartWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/StartWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).StartWorkstation(ctx, req.(*StartWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_StopWorkstation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopWorkstationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).StopWorkstation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/StopWorkstation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).StopWorkstation(ctx, req.(*StopWorkstationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workstations_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationsServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workstations.v1.Workstations/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationsServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Workstations_ServiceDesc is the grpc.ServiceDesc for Workstations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workstations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.workstations.v1.Workstations",
	HandlerType: (*WorkstationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkstationCluster",
			Handler:    _Workstations_GetWorkstationCluster_Handler,
		},
		{
			MethodName: "ListWorkstationClusters",
			Handler:    _Workstations_ListWorkstationClusters_Handler,
		},
		{
			MethodName: "CreateWorkstationCluster",
			Handler:    _Workstations_CreateWorkstationCluster_Handler,
		},
		{
			MethodName: "UpdateWorkstationCluster",
			Handler:    _Workstations_UpdateWorkstationCluster_Handler,
		},
		{
			MethodName: "DeleteWorkstationCluster",
			Handler:    _Workstations_DeleteWorkstationCluster_Handler,
		},
		{
			MethodName: "GetWorkstationConfig",
			Handler:    _Workstations_GetWorkstationConfig_Handler,
		},
		{
			MethodName: "ListWorkstationConfigs",
			Handler:    _Workstations_ListWorkstationConfigs_Handler,
		},
		{
			MethodName: "ListUsableWorkstationConfigs",
			Handler:    _Workstations_ListUsableWorkstationConfigs_Handler,
		},
		{
			MethodName: "CreateWorkstationConfig",
			Handler:    _Workstations_CreateWorkstationConfig_Handler,
		},
		{
			MethodName: "UpdateWorkstationConfig",
			Handler:    _Workstations_UpdateWorkstationConfig_Handler,
		},
		{
			MethodName: "DeleteWorkstationConfig",
			Handler:    _Workstations_DeleteWorkstationConfig_Handler,
		},
		{
			MethodName: "GetWorkstation",
			Handler:    _Workstations_GetWorkstation_Handler,
		},
		{
			MethodName: "ListWorkstations",
			Handler:    _Workstations_ListWorkstations_Handler,
		},
		{
			MethodName: "ListUsableWorkstations",
			Handler:    _Workstations_ListUsableWorkstations_Handler,
		},
		{
			MethodName: "CreateWorkstation",
			Handler:    _Workstations_CreateWorkstation_Handler,
		},
		{
			MethodName: "UpdateWorkstation",
			Handler:    _Workstations_UpdateWorkstation_Handler,
		},
		{
			MethodName: "DeleteWorkstation",
			Handler:    _Workstations_DeleteWorkstation_Handler,
		},
		{
			MethodName: "StartWorkstation",
			Handler:    _Workstations_StartWorkstation_Handler,
		},
		{
			MethodName: "StopWorkstation",
			Handler:    _Workstations_StopWorkstation_Handler,
		},
		{
			MethodName: "GenerateAccessToken",
			Handler:    _Workstations_GenerateAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/workstations/v1/workstations.proto",
}