// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_dns_task.proto

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
	DNSTaskService_ExistsDNSTasks_FullMethodName       = "/pb.DNSTaskService/existsDNSTasks"
	DNSTaskService_FindAllDoingDNSTasks_FullMethodName = "/pb.DNSTaskService/findAllDoingDNSTasks"
	DNSTaskService_DeleteDNSTask_FullMethodName        = "/pb.DNSTaskService/deleteDNSTask"
	DNSTaskService_DeleteAllDNSTasks_FullMethodName    = "/pb.DNSTaskService/deleteAllDNSTasks"
)

// DNSTaskServiceClient is the client API for DNSTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSTaskServiceClient interface {
	// 检查是否有正在执行的任务
	ExistsDNSTasks(ctx context.Context, in *ExistsDNSTasksRequest, opts ...grpc.CallOption) (*ExistsDNSTasksResponse, error)
	// 查找正在执行的所有任务
	FindAllDoingDNSTasks(ctx context.Context, in *FindAllDoingDNSTasksRequest, opts ...grpc.CallOption) (*FindAllDoingDNSTasksResponse, error)
	// 删除任务
	DeleteDNSTask(ctx context.Context, in *DeleteDNSTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除所有同步任务
	DeleteAllDNSTasks(ctx context.Context, in *DeleteAllDNSTasksRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type dNSTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSTaskServiceClient(cc grpc.ClientConnInterface) DNSTaskServiceClient {
	return &dNSTaskServiceClient{cc}
}

func (c *dNSTaskServiceClient) ExistsDNSTasks(ctx context.Context, in *ExistsDNSTasksRequest, opts ...grpc.CallOption) (*ExistsDNSTasksResponse, error) {
	out := new(ExistsDNSTasksResponse)
	err := c.cc.Invoke(ctx, DNSTaskService_ExistsDNSTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSTaskServiceClient) FindAllDoingDNSTasks(ctx context.Context, in *FindAllDoingDNSTasksRequest, opts ...grpc.CallOption) (*FindAllDoingDNSTasksResponse, error) {
	out := new(FindAllDoingDNSTasksResponse)
	err := c.cc.Invoke(ctx, DNSTaskService_FindAllDoingDNSTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSTaskServiceClient) DeleteDNSTask(ctx context.Context, in *DeleteDNSTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DNSTaskService_DeleteDNSTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSTaskServiceClient) DeleteAllDNSTasks(ctx context.Context, in *DeleteAllDNSTasksRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DNSTaskService_DeleteAllDNSTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSTaskServiceServer is the server API for DNSTaskService service.
// All implementations should embed UnimplementedDNSTaskServiceServer
// for forward compatibility
type DNSTaskServiceServer interface {
	// 检查是否有正在执行的任务
	ExistsDNSTasks(context.Context, *ExistsDNSTasksRequest) (*ExistsDNSTasksResponse, error)
	// 查找正在执行的所有任务
	FindAllDoingDNSTasks(context.Context, *FindAllDoingDNSTasksRequest) (*FindAllDoingDNSTasksResponse, error)
	// 删除任务
	DeleteDNSTask(context.Context, *DeleteDNSTaskRequest) (*RPCSuccess, error)
	// 删除所有同步任务
	DeleteAllDNSTasks(context.Context, *DeleteAllDNSTasksRequest) (*RPCSuccess, error)
}

// UnimplementedDNSTaskServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDNSTaskServiceServer struct {
}

func (UnimplementedDNSTaskServiceServer) ExistsDNSTasks(context.Context, *ExistsDNSTasksRequest) (*ExistsDNSTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsDNSTasks not implemented")
}
func (UnimplementedDNSTaskServiceServer) FindAllDoingDNSTasks(context.Context, *FindAllDoingDNSTasksRequest) (*FindAllDoingDNSTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDoingDNSTasks not implemented")
}
func (UnimplementedDNSTaskServiceServer) DeleteDNSTask(context.Context, *DeleteDNSTaskRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDNSTask not implemented")
}
func (UnimplementedDNSTaskServiceServer) DeleteAllDNSTasks(context.Context, *DeleteAllDNSTasksRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllDNSTasks not implemented")
}

// UnsafeDNSTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSTaskServiceServer will
// result in compilation errors.
type UnsafeDNSTaskServiceServer interface {
	mustEmbedUnimplementedDNSTaskServiceServer()
}

func RegisterDNSTaskServiceServer(s grpc.ServiceRegistrar, srv DNSTaskServiceServer) {
	s.RegisterService(&DNSTaskService_ServiceDesc, srv)
}

func _DNSTaskService_ExistsDNSTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsDNSTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).ExistsDNSTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSTaskService_ExistsDNSTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).ExistsDNSTasks(ctx, req.(*ExistsDNSTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSTaskService_FindAllDoingDNSTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDoingDNSTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).FindAllDoingDNSTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSTaskService_FindAllDoingDNSTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).FindAllDoingDNSTasks(ctx, req.(*FindAllDoingDNSTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSTaskService_DeleteDNSTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDNSTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).DeleteDNSTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSTaskService_DeleteDNSTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).DeleteDNSTask(ctx, req.(*DeleteDNSTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSTaskService_DeleteAllDNSTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllDNSTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSTaskServiceServer).DeleteAllDNSTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSTaskService_DeleteAllDNSTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSTaskServiceServer).DeleteAllDNSTasks(ctx, req.(*DeleteAllDNSTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSTaskService_ServiceDesc is the grpc.ServiceDesc for DNSTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNSTaskService",
	HandlerType: (*DNSTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "existsDNSTasks",
			Handler:    _DNSTaskService_ExistsDNSTasks_Handler,
		},
		{
			MethodName: "findAllDoingDNSTasks",
			Handler:    _DNSTaskService_FindAllDoingDNSTasks_Handler,
		},
		{
			MethodName: "deleteDNSTask",
			Handler:    _DNSTaskService_DeleteDNSTask_Handler,
		},
		{
			MethodName: "deleteAllDNSTasks",
			Handler:    _DNSTaskService_DeleteAllDNSTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dns_task.proto",
}
