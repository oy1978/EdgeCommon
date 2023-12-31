// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_http_location.proto

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
	HTTPLocationService_CreateHTTPLocation_FullMethodName                        = "/pb.HTTPLocationService/createHTTPLocation"
	HTTPLocationService_UpdateHTTPLocation_FullMethodName                        = "/pb.HTTPLocationService/updateHTTPLocation"
	HTTPLocationService_FindEnabledHTTPLocationConfig_FullMethodName             = "/pb.HTTPLocationService/findEnabledHTTPLocationConfig"
	HTTPLocationService_DeleteHTTPLocation_FullMethodName                        = "/pb.HTTPLocationService/deleteHTTPLocation"
	HTTPLocationService_FindAndInitHTTPLocationReverseProxyConfig_FullMethodName = "/pb.HTTPLocationService/findAndInitHTTPLocationReverseProxyConfig"
	HTTPLocationService_FindAndInitHTTPLocationWebConfig_FullMethodName          = "/pb.HTTPLocationService/findAndInitHTTPLocationWebConfig"
	HTTPLocationService_UpdateHTTPLocationReverseProxy_FullMethodName            = "/pb.HTTPLocationService/updateHTTPLocationReverseProxy"
)

// HTTPLocationServiceClient is the client API for HTTPLocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPLocationServiceClient interface {
	// 创建路径规则
	CreateHTTPLocation(ctx context.Context, in *CreateHTTPLocationRequest, opts ...grpc.CallOption) (*CreateHTTPLocationResponse, error)
	// 修改路径规则
	UpdateHTTPLocation(ctx context.Context, in *UpdateHTTPLocationRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找路径规则配置
	FindEnabledHTTPLocationConfig(ctx context.Context, in *FindEnabledHTTPLocationConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPLocationConfigResponse, error)
	// 删除路径规则
	DeleteHTTPLocation(ctx context.Context, in *DeleteHTTPLocationRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找反向代理设置
	FindAndInitHTTPLocationReverseProxyConfig(ctx context.Context, in *FindAndInitHTTPLocationReverseProxyConfigRequest, opts ...grpc.CallOption) (*FindAndInitHTTPLocationReverseProxyConfigResponse, error)
	// 初始化Web设置
	FindAndInitHTTPLocationWebConfig(ctx context.Context, in *FindAndInitHTTPLocationWebConfigRequest, opts ...grpc.CallOption) (*FindAndInitHTTPLocationWebConfigResponse, error)
	// 修改反向代理设置
	UpdateHTTPLocationReverseProxy(ctx context.Context, in *UpdateHTTPLocationReverseProxyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPLocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPLocationServiceClient(cc grpc.ClientConnInterface) HTTPLocationServiceClient {
	return &hTTPLocationServiceClient{cc}
}

func (c *hTTPLocationServiceClient) CreateHTTPLocation(ctx context.Context, in *CreateHTTPLocationRequest, opts ...grpc.CallOption) (*CreateHTTPLocationResponse, error) {
	out := new(CreateHTTPLocationResponse)
	err := c.cc.Invoke(ctx, HTTPLocationService_CreateHTTPLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) UpdateHTTPLocation(ctx context.Context, in *UpdateHTTPLocationRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPLocationService_UpdateHTTPLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) FindEnabledHTTPLocationConfig(ctx context.Context, in *FindEnabledHTTPLocationConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPLocationConfigResponse, error) {
	out := new(FindEnabledHTTPLocationConfigResponse)
	err := c.cc.Invoke(ctx, HTTPLocationService_FindEnabledHTTPLocationConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) DeleteHTTPLocation(ctx context.Context, in *DeleteHTTPLocationRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPLocationService_DeleteHTTPLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) FindAndInitHTTPLocationReverseProxyConfig(ctx context.Context, in *FindAndInitHTTPLocationReverseProxyConfigRequest, opts ...grpc.CallOption) (*FindAndInitHTTPLocationReverseProxyConfigResponse, error) {
	out := new(FindAndInitHTTPLocationReverseProxyConfigResponse)
	err := c.cc.Invoke(ctx, HTTPLocationService_FindAndInitHTTPLocationReverseProxyConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) FindAndInitHTTPLocationWebConfig(ctx context.Context, in *FindAndInitHTTPLocationWebConfigRequest, opts ...grpc.CallOption) (*FindAndInitHTTPLocationWebConfigResponse, error) {
	out := new(FindAndInitHTTPLocationWebConfigResponse)
	err := c.cc.Invoke(ctx, HTTPLocationService_FindAndInitHTTPLocationWebConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPLocationServiceClient) UpdateHTTPLocationReverseProxy(ctx context.Context, in *UpdateHTTPLocationReverseProxyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPLocationService_UpdateHTTPLocationReverseProxy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPLocationServiceServer is the server API for HTTPLocationService service.
// All implementations should embed UnimplementedHTTPLocationServiceServer
// for forward compatibility
type HTTPLocationServiceServer interface {
	// 创建路径规则
	CreateHTTPLocation(context.Context, *CreateHTTPLocationRequest) (*CreateHTTPLocationResponse, error)
	// 修改路径规则
	UpdateHTTPLocation(context.Context, *UpdateHTTPLocationRequest) (*RPCSuccess, error)
	// 查找路径规则配置
	FindEnabledHTTPLocationConfig(context.Context, *FindEnabledHTTPLocationConfigRequest) (*FindEnabledHTTPLocationConfigResponse, error)
	// 删除路径规则
	DeleteHTTPLocation(context.Context, *DeleteHTTPLocationRequest) (*RPCSuccess, error)
	// 查找反向代理设置
	FindAndInitHTTPLocationReverseProxyConfig(context.Context, *FindAndInitHTTPLocationReverseProxyConfigRequest) (*FindAndInitHTTPLocationReverseProxyConfigResponse, error)
	// 初始化Web设置
	FindAndInitHTTPLocationWebConfig(context.Context, *FindAndInitHTTPLocationWebConfigRequest) (*FindAndInitHTTPLocationWebConfigResponse, error)
	// 修改反向代理设置
	UpdateHTTPLocationReverseProxy(context.Context, *UpdateHTTPLocationReverseProxyRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPLocationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPLocationServiceServer struct {
}

func (UnimplementedHTTPLocationServiceServer) CreateHTTPLocation(context.Context, *CreateHTTPLocationRequest) (*CreateHTTPLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPLocation not implemented")
}
func (UnimplementedHTTPLocationServiceServer) UpdateHTTPLocation(context.Context, *UpdateHTTPLocationRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPLocation not implemented")
}
func (UnimplementedHTTPLocationServiceServer) FindEnabledHTTPLocationConfig(context.Context, *FindEnabledHTTPLocationConfigRequest) (*FindEnabledHTTPLocationConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPLocationConfig not implemented")
}
func (UnimplementedHTTPLocationServiceServer) DeleteHTTPLocation(context.Context, *DeleteHTTPLocationRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHTTPLocation not implemented")
}
func (UnimplementedHTTPLocationServiceServer) FindAndInitHTTPLocationReverseProxyConfig(context.Context, *FindAndInitHTTPLocationReverseProxyConfigRequest) (*FindAndInitHTTPLocationReverseProxyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAndInitHTTPLocationReverseProxyConfig not implemented")
}
func (UnimplementedHTTPLocationServiceServer) FindAndInitHTTPLocationWebConfig(context.Context, *FindAndInitHTTPLocationWebConfigRequest) (*FindAndInitHTTPLocationWebConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAndInitHTTPLocationWebConfig not implemented")
}
func (UnimplementedHTTPLocationServiceServer) UpdateHTTPLocationReverseProxy(context.Context, *UpdateHTTPLocationReverseProxyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPLocationReverseProxy not implemented")
}

// UnsafeHTTPLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPLocationServiceServer will
// result in compilation errors.
type UnsafeHTTPLocationServiceServer interface {
	mustEmbedUnimplementedHTTPLocationServiceServer()
}

func RegisterHTTPLocationServiceServer(s grpc.ServiceRegistrar, srv HTTPLocationServiceServer) {
	s.RegisterService(&HTTPLocationService_ServiceDesc, srv)
}

func _HTTPLocationService_CreateHTTPLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).CreateHTTPLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_CreateHTTPLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).CreateHTTPLocation(ctx, req.(*CreateHTTPLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_UpdateHTTPLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).UpdateHTTPLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_UpdateHTTPLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).UpdateHTTPLocation(ctx, req.(*UpdateHTTPLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_FindEnabledHTTPLocationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPLocationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).FindEnabledHTTPLocationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_FindEnabledHTTPLocationConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).FindEnabledHTTPLocationConfig(ctx, req.(*FindEnabledHTTPLocationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_DeleteHTTPLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHTTPLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).DeleteHTTPLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_DeleteHTTPLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).DeleteHTTPLocation(ctx, req.(*DeleteHTTPLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_FindAndInitHTTPLocationReverseProxyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAndInitHTTPLocationReverseProxyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).FindAndInitHTTPLocationReverseProxyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_FindAndInitHTTPLocationReverseProxyConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).FindAndInitHTTPLocationReverseProxyConfig(ctx, req.(*FindAndInitHTTPLocationReverseProxyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_FindAndInitHTTPLocationWebConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAndInitHTTPLocationWebConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).FindAndInitHTTPLocationWebConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_FindAndInitHTTPLocationWebConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).FindAndInitHTTPLocationWebConfig(ctx, req.(*FindAndInitHTTPLocationWebConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPLocationService_UpdateHTTPLocationReverseProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPLocationReverseProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPLocationServiceServer).UpdateHTTPLocationReverseProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPLocationService_UpdateHTTPLocationReverseProxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPLocationServiceServer).UpdateHTTPLocationReverseProxy(ctx, req.(*UpdateHTTPLocationReverseProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPLocationService_ServiceDesc is the grpc.ServiceDesc for HTTPLocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPLocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPLocationService",
	HandlerType: (*HTTPLocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPLocation",
			Handler:    _HTTPLocationService_CreateHTTPLocation_Handler,
		},
		{
			MethodName: "updateHTTPLocation",
			Handler:    _HTTPLocationService_UpdateHTTPLocation_Handler,
		},
		{
			MethodName: "findEnabledHTTPLocationConfig",
			Handler:    _HTTPLocationService_FindEnabledHTTPLocationConfig_Handler,
		},
		{
			MethodName: "deleteHTTPLocation",
			Handler:    _HTTPLocationService_DeleteHTTPLocation_Handler,
		},
		{
			MethodName: "findAndInitHTTPLocationReverseProxyConfig",
			Handler:    _HTTPLocationService_FindAndInitHTTPLocationReverseProxyConfig_Handler,
		},
		{
			MethodName: "findAndInitHTTPLocationWebConfig",
			Handler:    _HTTPLocationService_FindAndInitHTTPLocationWebConfig_Handler,
		},
		{
			MethodName: "updateHTTPLocationReverseProxy",
			Handler:    _HTTPLocationService_UpdateHTTPLocationReverseProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_location.proto",
}
