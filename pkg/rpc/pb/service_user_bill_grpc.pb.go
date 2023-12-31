// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_bill.proto

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
	UserBillService_GenerateAllUserBills_FullMethodName = "/pb.UserBillService/generateAllUserBills"
	UserBillService_CountAllUserBills_FullMethodName    = "/pb.UserBillService/countAllUserBills"
	UserBillService_ListUserBills_FullMethodName        = "/pb.UserBillService/listUserBills"
	UserBillService_FindUserBill_FullMethodName         = "/pb.UserBillService/findUserBill"
	UserBillService_PayUserBill_FullMethodName          = "/pb.UserBillService/payUserBill"
	UserBillService_SumUserUnpaidBills_FullMethodName   = "/pb.UserBillService/sumUserUnpaidBills"
)

// UserBillServiceClient is the client API for UserBillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserBillServiceClient interface {
	// 手工生成账单
	GenerateAllUserBills(ctx context.Context, in *GenerateAllUserBillsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算所有账单数量
	CountAllUserBills(ctx context.Context, in *CountAllUserBillsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页账单
	ListUserBills(ctx context.Context, in *ListUserBillsRequest, opts ...grpc.CallOption) (*ListUserBillsResponse, error)
	// 查找账单信息
	FindUserBill(ctx context.Context, in *FindUserBillRequest, opts ...grpc.CallOption) (*FindUserBillResponse, error)
	// 支付账单
	PayUserBill(ctx context.Context, in *PayUserBillRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算用户所有未支付账单总额
	SumUserUnpaidBills(ctx context.Context, in *SumUserUnpaidBillsRequest, opts ...grpc.CallOption) (*SumUserUnpaidBillsResponse, error)
}

type userBillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBillServiceClient(cc grpc.ClientConnInterface) UserBillServiceClient {
	return &userBillServiceClient{cc}
}

func (c *userBillServiceClient) GenerateAllUserBills(ctx context.Context, in *GenerateAllUserBillsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserBillService_GenerateAllUserBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) CountAllUserBills(ctx context.Context, in *CountAllUserBillsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, UserBillService_CountAllUserBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) ListUserBills(ctx context.Context, in *ListUserBillsRequest, opts ...grpc.CallOption) (*ListUserBillsResponse, error) {
	out := new(ListUserBillsResponse)
	err := c.cc.Invoke(ctx, UserBillService_ListUserBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) FindUserBill(ctx context.Context, in *FindUserBillRequest, opts ...grpc.CallOption) (*FindUserBillResponse, error) {
	out := new(FindUserBillResponse)
	err := c.cc.Invoke(ctx, UserBillService_FindUserBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) PayUserBill(ctx context.Context, in *PayUserBillRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserBillService_PayUserBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBillServiceClient) SumUserUnpaidBills(ctx context.Context, in *SumUserUnpaidBillsRequest, opts ...grpc.CallOption) (*SumUserUnpaidBillsResponse, error) {
	out := new(SumUserUnpaidBillsResponse)
	err := c.cc.Invoke(ctx, UserBillService_SumUserUnpaidBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBillServiceServer is the server API for UserBillService service.
// All implementations should embed UnimplementedUserBillServiceServer
// for forward compatibility
type UserBillServiceServer interface {
	// 手工生成账单
	GenerateAllUserBills(context.Context, *GenerateAllUserBillsRequest) (*RPCSuccess, error)
	// 计算所有账单数量
	CountAllUserBills(context.Context, *CountAllUserBillsRequest) (*RPCCountResponse, error)
	// 列出单页账单
	ListUserBills(context.Context, *ListUserBillsRequest) (*ListUserBillsResponse, error)
	// 查找账单信息
	FindUserBill(context.Context, *FindUserBillRequest) (*FindUserBillResponse, error)
	// 支付账单
	PayUserBill(context.Context, *PayUserBillRequest) (*RPCSuccess, error)
	// 计算用户所有未支付账单总额
	SumUserUnpaidBills(context.Context, *SumUserUnpaidBillsRequest) (*SumUserUnpaidBillsResponse, error)
}

// UnimplementedUserBillServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserBillServiceServer struct {
}

func (UnimplementedUserBillServiceServer) GenerateAllUserBills(context.Context, *GenerateAllUserBillsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAllUserBills not implemented")
}
func (UnimplementedUserBillServiceServer) CountAllUserBills(context.Context, *CountAllUserBillsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllUserBills not implemented")
}
func (UnimplementedUserBillServiceServer) ListUserBills(context.Context, *ListUserBillsRequest) (*ListUserBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserBills not implemented")
}
func (UnimplementedUserBillServiceServer) FindUserBill(context.Context, *FindUserBillRequest) (*FindUserBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserBill not implemented")
}
func (UnimplementedUserBillServiceServer) PayUserBill(context.Context, *PayUserBillRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayUserBill not implemented")
}
func (UnimplementedUserBillServiceServer) SumUserUnpaidBills(context.Context, *SumUserUnpaidBillsRequest) (*SumUserUnpaidBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumUserUnpaidBills not implemented")
}

// UnsafeUserBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserBillServiceServer will
// result in compilation errors.
type UnsafeUserBillServiceServer interface {
	mustEmbedUnimplementedUserBillServiceServer()
}

func RegisterUserBillServiceServer(s grpc.ServiceRegistrar, srv UserBillServiceServer) {
	s.RegisterService(&UserBillService_ServiceDesc, srv)
}

func _UserBillService_GenerateAllUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAllUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).GenerateAllUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_GenerateAllUserBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).GenerateAllUserBills(ctx, req.(*GenerateAllUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_CountAllUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).CountAllUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_CountAllUserBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).CountAllUserBills(ctx, req.(*CountAllUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_ListUserBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).ListUserBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_ListUserBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).ListUserBills(ctx, req.(*ListUserBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_FindUserBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).FindUserBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_FindUserBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).FindUserBill(ctx, req.(*FindUserBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_PayUserBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayUserBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).PayUserBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_PayUserBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).PayUserBill(ctx, req.(*PayUserBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBillService_SumUserUnpaidBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumUserUnpaidBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBillServiceServer).SumUserUnpaidBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBillService_SumUserUnpaidBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBillServiceServer).SumUserUnpaidBills(ctx, req.(*SumUserUnpaidBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserBillService_ServiceDesc is the grpc.ServiceDesc for UserBillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserBillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserBillService",
	HandlerType: (*UserBillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateAllUserBills",
			Handler:    _UserBillService_GenerateAllUserBills_Handler,
		},
		{
			MethodName: "countAllUserBills",
			Handler:    _UserBillService_CountAllUserBills_Handler,
		},
		{
			MethodName: "listUserBills",
			Handler:    _UserBillService_ListUserBills_Handler,
		},
		{
			MethodName: "findUserBill",
			Handler:    _UserBillService_FindUserBill_Handler,
		},
		{
			MethodName: "payUserBill",
			Handler:    _UserBillService_PayUserBill_Handler,
		},
		{
			MethodName: "sumUserUnpaidBills",
			Handler:    _UserBillService_SumUserUnpaidBills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_bill.proto",
}
