// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_formal_client_system.proto

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
	FormalClientSystemService_CreateFormalClientSystem_FullMethodName         = "/pb.FormalClientSystemService/createFormalClientSystem"
	FormalClientSystemService_CountFormalClientSystems_FullMethodName         = "/pb.FormalClientSystemService/countFormalClientSystems"
	FormalClientSystemService_ListFormalClientSystems_FullMethodName          = "/pb.FormalClientSystemService/listFormalClientSystems"
	FormalClientSystemService_UpdateFormalClientSystem_FullMethodName         = "/pb.FormalClientSystemService/updateFormalClientSystem"
	FormalClientSystemService_FindFormalClientSystemWithDataId_FullMethodName = "/pb.FormalClientSystemService/findFormalClientSystemWithDataId"
)

// FormalClientSystemServiceClient is the client API for FormalClientSystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FormalClientSystemServiceClient interface {
	// 创建操作系统信息
	CreateFormalClientSystem(ctx context.Context, in *CreateFormalClientSystemRequest, opts ...grpc.CallOption) (*CreateFormalClientSystemResponse, error)
	// 计算操作系统信息数量
	CountFormalClientSystems(ctx context.Context, in *CountFormalClientSystemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页操作系统信息
	ListFormalClientSystems(ctx context.Context, in *ListFormalClientSystemsRequest, opts ...grpc.CallOption) (*ListFormalClientSystemsResponse, error)
	// 修改操作系统信息
	UpdateFormalClientSystem(ctx context.Context, in *UpdateFormalClientSystemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 通过dataId查询操作系统信息
	FindFormalClientSystemWithDataId(ctx context.Context, in *FindFormalClientSystemWithDataIdRequest, opts ...grpc.CallOption) (*FindFormalClientSystemWithDataIdResponse, error)
}

type formalClientSystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFormalClientSystemServiceClient(cc grpc.ClientConnInterface) FormalClientSystemServiceClient {
	return &formalClientSystemServiceClient{cc}
}

func (c *formalClientSystemServiceClient) CreateFormalClientSystem(ctx context.Context, in *CreateFormalClientSystemRequest, opts ...grpc.CallOption) (*CreateFormalClientSystemResponse, error) {
	out := new(CreateFormalClientSystemResponse)
	err := c.cc.Invoke(ctx, FormalClientSystemService_CreateFormalClientSystem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formalClientSystemServiceClient) CountFormalClientSystems(ctx context.Context, in *CountFormalClientSystemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, FormalClientSystemService_CountFormalClientSystems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formalClientSystemServiceClient) ListFormalClientSystems(ctx context.Context, in *ListFormalClientSystemsRequest, opts ...grpc.CallOption) (*ListFormalClientSystemsResponse, error) {
	out := new(ListFormalClientSystemsResponse)
	err := c.cc.Invoke(ctx, FormalClientSystemService_ListFormalClientSystems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formalClientSystemServiceClient) UpdateFormalClientSystem(ctx context.Context, in *UpdateFormalClientSystemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, FormalClientSystemService_UpdateFormalClientSystem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formalClientSystemServiceClient) FindFormalClientSystemWithDataId(ctx context.Context, in *FindFormalClientSystemWithDataIdRequest, opts ...grpc.CallOption) (*FindFormalClientSystemWithDataIdResponse, error) {
	out := new(FindFormalClientSystemWithDataIdResponse)
	err := c.cc.Invoke(ctx, FormalClientSystemService_FindFormalClientSystemWithDataId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormalClientSystemServiceServer is the server API for FormalClientSystemService service.
// All implementations should embed UnimplementedFormalClientSystemServiceServer
// for forward compatibility
type FormalClientSystemServiceServer interface {
	// 创建操作系统信息
	CreateFormalClientSystem(context.Context, *CreateFormalClientSystemRequest) (*CreateFormalClientSystemResponse, error)
	// 计算操作系统信息数量
	CountFormalClientSystems(context.Context, *CountFormalClientSystemsRequest) (*RPCCountResponse, error)
	// 列出单页操作系统信息
	ListFormalClientSystems(context.Context, *ListFormalClientSystemsRequest) (*ListFormalClientSystemsResponse, error)
	// 修改操作系统信息
	UpdateFormalClientSystem(context.Context, *UpdateFormalClientSystemRequest) (*RPCSuccess, error)
	// 通过dataId查询操作系统信息
	FindFormalClientSystemWithDataId(context.Context, *FindFormalClientSystemWithDataIdRequest) (*FindFormalClientSystemWithDataIdResponse, error)
}

// UnimplementedFormalClientSystemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFormalClientSystemServiceServer struct {
}

func (UnimplementedFormalClientSystemServiceServer) CreateFormalClientSystem(context.Context, *CreateFormalClientSystemRequest) (*CreateFormalClientSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFormalClientSystem not implemented")
}
func (UnimplementedFormalClientSystemServiceServer) CountFormalClientSystems(context.Context, *CountFormalClientSystemsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFormalClientSystems not implemented")
}
func (UnimplementedFormalClientSystemServiceServer) ListFormalClientSystems(context.Context, *ListFormalClientSystemsRequest) (*ListFormalClientSystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFormalClientSystems not implemented")
}
func (UnimplementedFormalClientSystemServiceServer) UpdateFormalClientSystem(context.Context, *UpdateFormalClientSystemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFormalClientSystem not implemented")
}
func (UnimplementedFormalClientSystemServiceServer) FindFormalClientSystemWithDataId(context.Context, *FindFormalClientSystemWithDataIdRequest) (*FindFormalClientSystemWithDataIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFormalClientSystemWithDataId not implemented")
}

// UnsafeFormalClientSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FormalClientSystemServiceServer will
// result in compilation errors.
type UnsafeFormalClientSystemServiceServer interface {
	mustEmbedUnimplementedFormalClientSystemServiceServer()
}

func RegisterFormalClientSystemServiceServer(s grpc.ServiceRegistrar, srv FormalClientSystemServiceServer) {
	s.RegisterService(&FormalClientSystemService_ServiceDesc, srv)
}

func _FormalClientSystemService_CreateFormalClientSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFormalClientSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormalClientSystemServiceServer).CreateFormalClientSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormalClientSystemService_CreateFormalClientSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormalClientSystemServiceServer).CreateFormalClientSystem(ctx, req.(*CreateFormalClientSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormalClientSystemService_CountFormalClientSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFormalClientSystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormalClientSystemServiceServer).CountFormalClientSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormalClientSystemService_CountFormalClientSystems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormalClientSystemServiceServer).CountFormalClientSystems(ctx, req.(*CountFormalClientSystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormalClientSystemService_ListFormalClientSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFormalClientSystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormalClientSystemServiceServer).ListFormalClientSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormalClientSystemService_ListFormalClientSystems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormalClientSystemServiceServer).ListFormalClientSystems(ctx, req.(*ListFormalClientSystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormalClientSystemService_UpdateFormalClientSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFormalClientSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormalClientSystemServiceServer).UpdateFormalClientSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormalClientSystemService_UpdateFormalClientSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormalClientSystemServiceServer).UpdateFormalClientSystem(ctx, req.(*UpdateFormalClientSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormalClientSystemService_FindFormalClientSystemWithDataId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFormalClientSystemWithDataIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormalClientSystemServiceServer).FindFormalClientSystemWithDataId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormalClientSystemService_FindFormalClientSystemWithDataId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormalClientSystemServiceServer).FindFormalClientSystemWithDataId(ctx, req.(*FindFormalClientSystemWithDataIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FormalClientSystemService_ServiceDesc is the grpc.ServiceDesc for FormalClientSystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FormalClientSystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FormalClientSystemService",
	HandlerType: (*FormalClientSystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createFormalClientSystem",
			Handler:    _FormalClientSystemService_CreateFormalClientSystem_Handler,
		},
		{
			MethodName: "countFormalClientSystems",
			Handler:    _FormalClientSystemService_CountFormalClientSystems_Handler,
		},
		{
			MethodName: "listFormalClientSystems",
			Handler:    _FormalClientSystemService_ListFormalClientSystems_Handler,
		},
		{
			MethodName: "updateFormalClientSystem",
			Handler:    _FormalClientSystemService_UpdateFormalClientSystem_Handler,
		},
		{
			MethodName: "findFormalClientSystemWithDataId",
			Handler:    _FormalClientSystemService_FindFormalClientSystemWithDataId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_formal_client_system.proto",
}