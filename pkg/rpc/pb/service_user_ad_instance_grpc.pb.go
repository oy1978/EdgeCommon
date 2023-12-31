// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_ad_instance.proto

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
	UserADInstanceService_CreateUserADInstance_FullMethodName        = "/pb.UserADInstanceService/createUserADInstance"
	UserADInstanceService_BuyUserADInstance_FullMethodName           = "/pb.UserADInstanceService/buyUserADInstance"
	UserADInstanceService_CountUserADInstances_FullMethodName        = "/pb.UserADInstanceService/countUserADInstances"
	UserADInstanceService_ListUserADInstances_FullMethodName         = "/pb.UserADInstanceService/listUserADInstances"
	UserADInstanceService_FindUserADInstance_FullMethodName          = "/pb.UserADInstanceService/findUserADInstance"
	UserADInstanceService_DeleteUserADInstance_FullMethodName        = "/pb.UserADInstanceService/deleteUserADInstance"
	UserADInstanceService_RenewUserADInstance_FullMethodName         = "/pb.UserADInstanceService/renewUserADInstance"
	UserADInstanceService_UpdateUserADInstanceObjects_FullMethodName = "/pb.UserADInstanceService/updateUserADInstanceObjects"
)

// UserADInstanceServiceClient is the client API for UserADInstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserADInstanceServiceClient interface {
	// 创建用户高防实例
	CreateUserADInstance(ctx context.Context, in *CreateUserADInstanceRequest, opts ...grpc.CallOption) (*CreateUserADInstanceResponse, error)
	// 购买用户高防实例
	BuyUserADInstance(ctx context.Context, in *BuyUserADInstanceRequest, opts ...grpc.CallOption) (*BuyUserADInstanceResponse, error)
	// 计算用户高防实例数量
	CountUserADInstances(ctx context.Context, in *CountUserADInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页用户高防实例
	ListUserADInstances(ctx context.Context, in *ListUserADInstancesRequest, opts ...grpc.CallOption) (*ListUserADInstancesResponse, error)
	// 查找单个用户高防实例
	FindUserADInstance(ctx context.Context, in *FindUserADInstanceRequest, opts ...grpc.CallOption) (*FindUserADInstanceResponse, error)
	// 删除用户高防实例
	DeleteUserADInstance(ctx context.Context, in *DeleteUserADInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 续期用户高防实例
	RenewUserADInstance(ctx context.Context, in *RenewUserADInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改实例防护对象
	UpdateUserADInstanceObjects(ctx context.Context, in *UpdateUserADInstanceObjectsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type userADInstanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserADInstanceServiceClient(cc grpc.ClientConnInterface) UserADInstanceServiceClient {
	return &userADInstanceServiceClient{cc}
}

func (c *userADInstanceServiceClient) CreateUserADInstance(ctx context.Context, in *CreateUserADInstanceRequest, opts ...grpc.CallOption) (*CreateUserADInstanceResponse, error) {
	out := new(CreateUserADInstanceResponse)
	err := c.cc.Invoke(ctx, UserADInstanceService_CreateUserADInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) BuyUserADInstance(ctx context.Context, in *BuyUserADInstanceRequest, opts ...grpc.CallOption) (*BuyUserADInstanceResponse, error) {
	out := new(BuyUserADInstanceResponse)
	err := c.cc.Invoke(ctx, UserADInstanceService_BuyUserADInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) CountUserADInstances(ctx context.Context, in *CountUserADInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, UserADInstanceService_CountUserADInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) ListUserADInstances(ctx context.Context, in *ListUserADInstancesRequest, opts ...grpc.CallOption) (*ListUserADInstancesResponse, error) {
	out := new(ListUserADInstancesResponse)
	err := c.cc.Invoke(ctx, UserADInstanceService_ListUserADInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) FindUserADInstance(ctx context.Context, in *FindUserADInstanceRequest, opts ...grpc.CallOption) (*FindUserADInstanceResponse, error) {
	out := new(FindUserADInstanceResponse)
	err := c.cc.Invoke(ctx, UserADInstanceService_FindUserADInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) DeleteUserADInstance(ctx context.Context, in *DeleteUserADInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserADInstanceService_DeleteUserADInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) RenewUserADInstance(ctx context.Context, in *RenewUserADInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserADInstanceService_RenewUserADInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userADInstanceServiceClient) UpdateUserADInstanceObjects(ctx context.Context, in *UpdateUserADInstanceObjectsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserADInstanceService_UpdateUserADInstanceObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserADInstanceServiceServer is the server API for UserADInstanceService service.
// All implementations should embed UnimplementedUserADInstanceServiceServer
// for forward compatibility
type UserADInstanceServiceServer interface {
	// 创建用户高防实例
	CreateUserADInstance(context.Context, *CreateUserADInstanceRequest) (*CreateUserADInstanceResponse, error)
	// 购买用户高防实例
	BuyUserADInstance(context.Context, *BuyUserADInstanceRequest) (*BuyUserADInstanceResponse, error)
	// 计算用户高防实例数量
	CountUserADInstances(context.Context, *CountUserADInstancesRequest) (*RPCCountResponse, error)
	// 列出单页用户高防实例
	ListUserADInstances(context.Context, *ListUserADInstancesRequest) (*ListUserADInstancesResponse, error)
	// 查找单个用户高防实例
	FindUserADInstance(context.Context, *FindUserADInstanceRequest) (*FindUserADInstanceResponse, error)
	// 删除用户高防实例
	DeleteUserADInstance(context.Context, *DeleteUserADInstanceRequest) (*RPCSuccess, error)
	// 续期用户高防实例
	RenewUserADInstance(context.Context, *RenewUserADInstanceRequest) (*RPCSuccess, error)
	// 修改实例防护对象
	UpdateUserADInstanceObjects(context.Context, *UpdateUserADInstanceObjectsRequest) (*RPCSuccess, error)
}

// UnimplementedUserADInstanceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserADInstanceServiceServer struct {
}

func (UnimplementedUserADInstanceServiceServer) CreateUserADInstance(context.Context, *CreateUserADInstanceRequest) (*CreateUserADInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserADInstance not implemented")
}
func (UnimplementedUserADInstanceServiceServer) BuyUserADInstance(context.Context, *BuyUserADInstanceRequest) (*BuyUserADInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyUserADInstance not implemented")
}
func (UnimplementedUserADInstanceServiceServer) CountUserADInstances(context.Context, *CountUserADInstancesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUserADInstances not implemented")
}
func (UnimplementedUserADInstanceServiceServer) ListUserADInstances(context.Context, *ListUserADInstancesRequest) (*ListUserADInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserADInstances not implemented")
}
func (UnimplementedUserADInstanceServiceServer) FindUserADInstance(context.Context, *FindUserADInstanceRequest) (*FindUserADInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserADInstance not implemented")
}
func (UnimplementedUserADInstanceServiceServer) DeleteUserADInstance(context.Context, *DeleteUserADInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserADInstance not implemented")
}
func (UnimplementedUserADInstanceServiceServer) RenewUserADInstance(context.Context, *RenewUserADInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewUserADInstance not implemented")
}
func (UnimplementedUserADInstanceServiceServer) UpdateUserADInstanceObjects(context.Context, *UpdateUserADInstanceObjectsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserADInstanceObjects not implemented")
}

// UnsafeUserADInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserADInstanceServiceServer will
// result in compilation errors.
type UnsafeUserADInstanceServiceServer interface {
	mustEmbedUnimplementedUserADInstanceServiceServer()
}

func RegisterUserADInstanceServiceServer(s grpc.ServiceRegistrar, srv UserADInstanceServiceServer) {
	s.RegisterService(&UserADInstanceService_ServiceDesc, srv)
}

func _UserADInstanceService_CreateUserADInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserADInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).CreateUserADInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_CreateUserADInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).CreateUserADInstance(ctx, req.(*CreateUserADInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_BuyUserADInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyUserADInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).BuyUserADInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_BuyUserADInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).BuyUserADInstance(ctx, req.(*BuyUserADInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_CountUserADInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUserADInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).CountUserADInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_CountUserADInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).CountUserADInstances(ctx, req.(*CountUserADInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_ListUserADInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserADInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).ListUserADInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_ListUserADInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).ListUserADInstances(ctx, req.(*ListUserADInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_FindUserADInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserADInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).FindUserADInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_FindUserADInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).FindUserADInstance(ctx, req.(*FindUserADInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_DeleteUserADInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserADInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).DeleteUserADInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_DeleteUserADInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).DeleteUserADInstance(ctx, req.(*DeleteUserADInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_RenewUserADInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewUserADInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).RenewUserADInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_RenewUserADInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).RenewUserADInstance(ctx, req.(*RenewUserADInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserADInstanceService_UpdateUserADInstanceObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserADInstanceObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserADInstanceServiceServer).UpdateUserADInstanceObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserADInstanceService_UpdateUserADInstanceObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserADInstanceServiceServer).UpdateUserADInstanceObjects(ctx, req.(*UpdateUserADInstanceObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserADInstanceService_ServiceDesc is the grpc.ServiceDesc for UserADInstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserADInstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserADInstanceService",
	HandlerType: (*UserADInstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUserADInstance",
			Handler:    _UserADInstanceService_CreateUserADInstance_Handler,
		},
		{
			MethodName: "buyUserADInstance",
			Handler:    _UserADInstanceService_BuyUserADInstance_Handler,
		},
		{
			MethodName: "countUserADInstances",
			Handler:    _UserADInstanceService_CountUserADInstances_Handler,
		},
		{
			MethodName: "listUserADInstances",
			Handler:    _UserADInstanceService_ListUserADInstances_Handler,
		},
		{
			MethodName: "findUserADInstance",
			Handler:    _UserADInstanceService_FindUserADInstance_Handler,
		},
		{
			MethodName: "deleteUserADInstance",
			Handler:    _UserADInstanceService_DeleteUserADInstance_Handler,
		},
		{
			MethodName: "renewUserADInstance",
			Handler:    _UserADInstanceService_RenewUserADInstance_Handler,
		},
		{
			MethodName: "updateUserADInstanceObjects",
			Handler:    _UserADInstanceService_UpdateUserADInstanceObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_ad_instance.proto",
}
