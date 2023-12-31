// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ad_network.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ADNetworkService_CreateADNetwork_FullMethodName            = "/pb.ADNetworkService/createADNetwork"
	ADNetworkService_UpdateADNetwork_FullMethodName            = "/pb.ADNetworkService/updateADNetwork"
	ADNetworkService_FindADNetwork_FullMethodName              = "/pb.ADNetworkService/findADNetwork"
	ADNetworkService_FindAllADNetworks_FullMethodName          = "/pb.ADNetworkService/findAllADNetworks"
	ADNetworkService_FindAllAvailableADNetworks_FullMethodName = "/pb.ADNetworkService/findAllAvailableADNetworks"
	ADNetworkService_DeleteADNetwork_FullMethodName            = "/pb.ADNetworkService/deleteADNetwork"
)

// ADNetworkServiceClient is the client API for ADNetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ADNetworkServiceClient interface {
	// 创建线路
	CreateADNetwork(ctx context.Context, in *CreateADNetworkRequest, opts ...grpc.CallOption) (*CreateADNetworkResponse, error)
	// 修改线路
	UpdateADNetwork(ctx context.Context, in *UpdateADNetworkRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个线路
	FindADNetwork(ctx context.Context, in *FindADNetworkRequest, opts ...grpc.CallOption) (*FindADNetworkResponse, error)
	// 列出所有线路
	FindAllADNetworks(ctx context.Context, in *FindAllADNetworkRequest, opts ...grpc.CallOption) (*FindAllADNetworkResponse, error)
	// 列出所有可用的线路
	FindAllAvailableADNetworks(ctx context.Context, in *FindAllAvailableADNetworksRequest, opts ...grpc.CallOption) (*FindAllAvailableADNetworksResponse, error)
	// 删除线路
	DeleteADNetwork(ctx context.Context, in *DeleteADNetworkRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type aDNetworkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewADNetworkServiceClient(cc grpc.ClientConnInterface) ADNetworkServiceClient {
	return &aDNetworkServiceClient{cc}
}

func (c *aDNetworkServiceClient) CreateADNetwork(ctx context.Context, in *CreateADNetworkRequest, opts ...grpc.CallOption) (*CreateADNetworkResponse, error) {
	out := new(CreateADNetworkResponse)
	err := c.cc.Invoke(ctx, ADNetworkService_CreateADNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDNetworkServiceClient) UpdateADNetwork(ctx context.Context, in *UpdateADNetworkRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ADNetworkService_UpdateADNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDNetworkServiceClient) FindADNetwork(ctx context.Context, in *FindADNetworkRequest, opts ...grpc.CallOption) (*FindADNetworkResponse, error) {
	out := new(FindADNetworkResponse)
	err := c.cc.Invoke(ctx, ADNetworkService_FindADNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDNetworkServiceClient) FindAllADNetworks(ctx context.Context, in *FindAllADNetworkRequest, opts ...grpc.CallOption) (*FindAllADNetworkResponse, error) {
	out := new(FindAllADNetworkResponse)
	err := c.cc.Invoke(ctx, ADNetworkService_FindAllADNetworks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDNetworkServiceClient) FindAllAvailableADNetworks(ctx context.Context, in *FindAllAvailableADNetworksRequest, opts ...grpc.CallOption) (*FindAllAvailableADNetworksResponse, error) {
	out := new(FindAllAvailableADNetworksResponse)
	err := c.cc.Invoke(ctx, ADNetworkService_FindAllAvailableADNetworks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDNetworkServiceClient) DeleteADNetwork(ctx context.Context, in *DeleteADNetworkRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ADNetworkService_DeleteADNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ADNetworkServiceServer is the server API for ADNetworkService service.
// All implementations should embed UnimplementedADNetworkServiceServer
// for forward compatibility
type ADNetworkServiceServer interface {
	// 创建线路
	CreateADNetwork(context.Context, *CreateADNetworkRequest) (*CreateADNetworkResponse, error)
	// 修改线路
	UpdateADNetwork(context.Context, *UpdateADNetworkRequest) (*RPCSuccess, error)
	// 查找单个线路
	FindADNetwork(context.Context, *FindADNetworkRequest) (*FindADNetworkResponse, error)
	// 列出所有线路
	FindAllADNetworks(context.Context, *FindAllADNetworkRequest) (*FindAllADNetworkResponse, error)
	// 列出所有可用的线路
	FindAllAvailableADNetworks(context.Context, *FindAllAvailableADNetworksRequest) (*FindAllAvailableADNetworksResponse, error)
	// 删除线路
	DeleteADNetwork(context.Context, *DeleteADNetworkRequest) (*RPCSuccess, error)
}

// UnimplementedADNetworkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedADNetworkServiceServer struct {
}

func (UnimplementedADNetworkServiceServer) CreateADNetwork(context.Context, *CreateADNetworkRequest) (*CreateADNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateADNetwork not implemented")
}
func (UnimplementedADNetworkServiceServer) UpdateADNetwork(context.Context, *UpdateADNetworkRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateADNetwork not implemented")
}
func (UnimplementedADNetworkServiceServer) FindADNetwork(context.Context, *FindADNetworkRequest) (*FindADNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindADNetwork not implemented")
}
func (UnimplementedADNetworkServiceServer) FindAllADNetworks(context.Context, *FindAllADNetworkRequest) (*FindAllADNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllADNetworks not implemented")
}
func (UnimplementedADNetworkServiceServer) FindAllAvailableADNetworks(context.Context, *FindAllAvailableADNetworksRequest) (*FindAllAvailableADNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAvailableADNetworks not implemented")
}
func (UnimplementedADNetworkServiceServer) DeleteADNetwork(context.Context, *DeleteADNetworkRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteADNetwork not implemented")
}

// UnsafeADNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ADNetworkServiceServer will
// result in compilation errors.
type UnsafeADNetworkServiceServer interface {
	mustEmbedUnimplementedADNetworkServiceServer()
}

func RegisterADNetworkServiceServer(s grpc.ServiceRegistrar, srv ADNetworkServiceServer) {
	s.RegisterService(&ADNetworkService_ServiceDesc, srv)
}

func _ADNetworkService_CreateADNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateADNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).CreateADNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_CreateADNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).CreateADNetwork(ctx, req.(*CreateADNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADNetworkService_UpdateADNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateADNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).UpdateADNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_UpdateADNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).UpdateADNetwork(ctx, req.(*UpdateADNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADNetworkService_FindADNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindADNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).FindADNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_FindADNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).FindADNetwork(ctx, req.(*FindADNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADNetworkService_FindAllADNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllADNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).FindAllADNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_FindAllADNetworks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).FindAllADNetworks(ctx, req.(*FindAllADNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADNetworkService_FindAllAvailableADNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAvailableADNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).FindAllAvailableADNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_FindAllAvailableADNetworks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).FindAllAvailableADNetworks(ctx, req.(*FindAllAvailableADNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADNetworkService_DeleteADNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteADNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADNetworkServiceServer).DeleteADNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADNetworkService_DeleteADNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADNetworkServiceServer).DeleteADNetwork(ctx, req.(*DeleteADNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ADNetworkService_ServiceDesc is the grpc.ServiceDesc for ADNetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ADNetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ADNetworkService",
	HandlerType: (*ADNetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createADNetwork",
			Handler:    _ADNetworkService_CreateADNetwork_Handler,
		},
		{
			MethodName: "updateADNetwork",
			Handler:    _ADNetworkService_UpdateADNetwork_Handler,
		},
		{
			MethodName: "findADNetwork",
			Handler:    _ADNetworkService_FindADNetwork_Handler,
		},
		{
			MethodName: "findAllADNetworks",
			Handler:    _ADNetworkService_FindAllADNetworks_Handler,
		},
		{
			MethodName: "findAllAvailableADNetworks",
			Handler:    _ADNetworkService_FindAllAvailableADNetworks_Handler,
		},
		{
			MethodName: "deleteADNetwork",
			Handler:    _ADNetworkService_DeleteADNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ad_network.proto",
}
