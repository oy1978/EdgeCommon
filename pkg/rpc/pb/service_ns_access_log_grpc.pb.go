// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ns_access_log.proto

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
	NSAccessLogService_CreateNSAccessLogs_FullMethodName = "/pb.NSAccessLogService/createNSAccessLogs"
	NSAccessLogService_ListNSAccessLogs_FullMethodName   = "/pb.NSAccessLogService/listNSAccessLogs"
	NSAccessLogService_FindNSAccessLog_FullMethodName    = "/pb.NSAccessLogService/findNSAccessLog"
)

// NSAccessLogServiceClient is the client API for NSAccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSAccessLogServiceClient interface {
	// 创建访问日志
	CreateNSAccessLogs(ctx context.Context, in *CreateNSAccessLogsRequest, opts ...grpc.CallOption) (*CreateNSAccessLogsResponse, error)
	// 列出单页访问日志
	ListNSAccessLogs(ctx context.Context, in *ListNSAccessLogsRequest, opts ...grpc.CallOption) (*ListNSAccessLogsResponse, error)
	// 查找单个日志
	FindNSAccessLog(ctx context.Context, in *FindNSAccessLogRequest, opts ...grpc.CallOption) (*FindNSAccessLogResponse, error)
}

type nSAccessLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSAccessLogServiceClient(cc grpc.ClientConnInterface) NSAccessLogServiceClient {
	return &nSAccessLogServiceClient{cc}
}

func (c *nSAccessLogServiceClient) CreateNSAccessLogs(ctx context.Context, in *CreateNSAccessLogsRequest, opts ...grpc.CallOption) (*CreateNSAccessLogsResponse, error) {
	out := new(CreateNSAccessLogsResponse)
	err := c.cc.Invoke(ctx, NSAccessLogService_CreateNSAccessLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAccessLogServiceClient) ListNSAccessLogs(ctx context.Context, in *ListNSAccessLogsRequest, opts ...grpc.CallOption) (*ListNSAccessLogsResponse, error) {
	out := new(ListNSAccessLogsResponse)
	err := c.cc.Invoke(ctx, NSAccessLogService_ListNSAccessLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAccessLogServiceClient) FindNSAccessLog(ctx context.Context, in *FindNSAccessLogRequest, opts ...grpc.CallOption) (*FindNSAccessLogResponse, error) {
	out := new(FindNSAccessLogResponse)
	err := c.cc.Invoke(ctx, NSAccessLogService_FindNSAccessLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSAccessLogServiceServer is the server API for NSAccessLogService service.
// All implementations should embed UnimplementedNSAccessLogServiceServer
// for forward compatibility
type NSAccessLogServiceServer interface {
	// 创建访问日志
	CreateNSAccessLogs(context.Context, *CreateNSAccessLogsRequest) (*CreateNSAccessLogsResponse, error)
	// 列出单页访问日志
	ListNSAccessLogs(context.Context, *ListNSAccessLogsRequest) (*ListNSAccessLogsResponse, error)
	// 查找单个日志
	FindNSAccessLog(context.Context, *FindNSAccessLogRequest) (*FindNSAccessLogResponse, error)
}

// UnimplementedNSAccessLogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNSAccessLogServiceServer struct {
}

func (UnimplementedNSAccessLogServiceServer) CreateNSAccessLogs(context.Context, *CreateNSAccessLogsRequest) (*CreateNSAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNSAccessLogs not implemented")
}
func (UnimplementedNSAccessLogServiceServer) ListNSAccessLogs(context.Context, *ListNSAccessLogsRequest) (*ListNSAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNSAccessLogs not implemented")
}
func (UnimplementedNSAccessLogServiceServer) FindNSAccessLog(context.Context, *FindNSAccessLogRequest) (*FindNSAccessLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSAccessLog not implemented")
}

// UnsafeNSAccessLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSAccessLogServiceServer will
// result in compilation errors.
type UnsafeNSAccessLogServiceServer interface {
	mustEmbedUnimplementedNSAccessLogServiceServer()
}

func RegisterNSAccessLogServiceServer(s grpc.ServiceRegistrar, srv NSAccessLogServiceServer) {
	s.RegisterService(&NSAccessLogService_ServiceDesc, srv)
}

func _NSAccessLogService_CreateNSAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNSAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAccessLogServiceServer).CreateNSAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSAccessLogService_CreateNSAccessLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAccessLogServiceServer).CreateNSAccessLogs(ctx, req.(*CreateNSAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAccessLogService_ListNSAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNSAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAccessLogServiceServer).ListNSAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSAccessLogService_ListNSAccessLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAccessLogServiceServer).ListNSAccessLogs(ctx, req.(*ListNSAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAccessLogService_FindNSAccessLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSAccessLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAccessLogServiceServer).FindNSAccessLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSAccessLogService_FindNSAccessLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAccessLogServiceServer).FindNSAccessLog(ctx, req.(*FindNSAccessLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSAccessLogService_ServiceDesc is the grpc.ServiceDesc for NSAccessLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSAccessLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSAccessLogService",
	HandlerType: (*NSAccessLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNSAccessLogs",
			Handler:    _NSAccessLogService_CreateNSAccessLogs_Handler,
		},
		{
			MethodName: "listNSAccessLogs",
			Handler:    _NSAccessLogService_ListNSAccessLogs_Handler,
		},
		{
			MethodName: "findNSAccessLog",
			Handler:    _NSAccessLogService_FindNSAccessLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns_access_log.proto",
}
