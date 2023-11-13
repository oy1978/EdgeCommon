// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_http_page.proto

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
	HTTPPageService_CreateHTTPPage_FullMethodName            = "/pb.HTTPPageService/createHTTPPage"
	HTTPPageService_UpdateHTTPPage_FullMethodName            = "/pb.HTTPPageService/updateHTTPPage"
	HTTPPageService_FindEnabledHTTPPageConfig_FullMethodName = "/pb.HTTPPageService/findEnabledHTTPPageConfig"
)

// HTTPPageServiceClient is the client API for HTTPPageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPPageServiceClient interface {
	// 创建Page
	CreateHTTPPage(ctx context.Context, in *CreateHTTPPageRequest, opts ...grpc.CallOption) (*CreateHTTPPageResponse, error)
	// 修改Page
	UpdateHTTPPage(ctx context.Context, in *UpdateHTTPPageRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个Page配置
	FindEnabledHTTPPageConfig(ctx context.Context, in *FindEnabledHTTPPageConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPPageConfigResponse, error)
}

type hTTPPageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPPageServiceClient(cc grpc.ClientConnInterface) HTTPPageServiceClient {
	return &hTTPPageServiceClient{cc}
}

func (c *hTTPPageServiceClient) CreateHTTPPage(ctx context.Context, in *CreateHTTPPageRequest, opts ...grpc.CallOption) (*CreateHTTPPageResponse, error) {
	out := new(CreateHTTPPageResponse)
	err := c.cc.Invoke(ctx, HTTPPageService_CreateHTTPPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPPageServiceClient) UpdateHTTPPage(ctx context.Context, in *UpdateHTTPPageRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPPageService_UpdateHTTPPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPPageServiceClient) FindEnabledHTTPPageConfig(ctx context.Context, in *FindEnabledHTTPPageConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPPageConfigResponse, error) {
	out := new(FindEnabledHTTPPageConfigResponse)
	err := c.cc.Invoke(ctx, HTTPPageService_FindEnabledHTTPPageConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPPageServiceServer is the server API for HTTPPageService service.
// All implementations should embed UnimplementedHTTPPageServiceServer
// for forward compatibility
type HTTPPageServiceServer interface {
	// 创建Page
	CreateHTTPPage(context.Context, *CreateHTTPPageRequest) (*CreateHTTPPageResponse, error)
	// 修改Page
	UpdateHTTPPage(context.Context, *UpdateHTTPPageRequest) (*RPCSuccess, error)
	// 查找单个Page配置
	FindEnabledHTTPPageConfig(context.Context, *FindEnabledHTTPPageConfigRequest) (*FindEnabledHTTPPageConfigResponse, error)
}

// UnimplementedHTTPPageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPPageServiceServer struct {
}

func (UnimplementedHTTPPageServiceServer) CreateHTTPPage(context.Context, *CreateHTTPPageRequest) (*CreateHTTPPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPPage not implemented")
}
func (UnimplementedHTTPPageServiceServer) UpdateHTTPPage(context.Context, *UpdateHTTPPageRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPPage not implemented")
}
func (UnimplementedHTTPPageServiceServer) FindEnabledHTTPPageConfig(context.Context, *FindEnabledHTTPPageConfigRequest) (*FindEnabledHTTPPageConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPPageConfig not implemented")
}

// UnsafeHTTPPageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPPageServiceServer will
// result in compilation errors.
type UnsafeHTTPPageServiceServer interface {
	mustEmbedUnimplementedHTTPPageServiceServer()
}

func RegisterHTTPPageServiceServer(s grpc.ServiceRegistrar, srv HTTPPageServiceServer) {
	s.RegisterService(&HTTPPageService_ServiceDesc, srv)
}

func _HTTPPageService_CreateHTTPPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).CreateHTTPPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPPageService_CreateHTTPPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).CreateHTTPPage(ctx, req.(*CreateHTTPPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPPageService_UpdateHTTPPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).UpdateHTTPPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPPageService_UpdateHTTPPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).UpdateHTTPPage(ctx, req.(*UpdateHTTPPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPPageService_FindEnabledHTTPPageConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPPageConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPPageServiceServer).FindEnabledHTTPPageConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPPageService_FindEnabledHTTPPageConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPPageServiceServer).FindEnabledHTTPPageConfig(ctx, req.(*FindEnabledHTTPPageConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPPageService_ServiceDesc is the grpc.ServiceDesc for HTTPPageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPPageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPPageService",
	HandlerType: (*HTTPPageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPPage",
			Handler:    _HTTPPageService_CreateHTTPPage_Handler,
		},
		{
			MethodName: "updateHTTPPage",
			Handler:    _HTTPPageService_UpdateHTTPPage_Handler,
		},
		{
			MethodName: "findEnabledHTTPPageConfig",
			Handler:    _HTTPPageService_FindEnabledHTTPPageConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_page.proto",
}