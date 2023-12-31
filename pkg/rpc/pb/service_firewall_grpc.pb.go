// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_firewall.proto

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
	FirewallService_ComposeFirewallGlobalBoard_FullMethodName = "/pb.FirewallService/composeFirewallGlobalBoard"
	FirewallService_NotifyHTTPFirewallEvent_FullMethodName    = "/pb.FirewallService/notifyHTTPFirewallEvent"
	FirewallService_CountFirewallDailyBlocks_FullMethodName   = "/pb.FirewallService/countFirewallDailyBlocks"
)

// FirewallServiceClient is the client API for FirewallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirewallServiceClient interface {
	// 组合看板数据
	ComposeFirewallGlobalBoard(ctx context.Context, in *ComposeFirewallGlobalBoardRequest, opts ...grpc.CallOption) (*ComposeFirewallGlobalBoardResponse, error)
	// 发送告警(notify)消息
	NotifyHTTPFirewallEvent(ctx context.Context, in *NotifyHTTPFirewallEventRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取当前Block动作次数
	CountFirewallDailyBlocks(ctx context.Context, in *CountFirewallDailyBlocksRequest, opts ...grpc.CallOption) (*CountFirewallDailyBlocksResponse, error)
}

type firewallServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirewallServiceClient(cc grpc.ClientConnInterface) FirewallServiceClient {
	return &firewallServiceClient{cc}
}

func (c *firewallServiceClient) ComposeFirewallGlobalBoard(ctx context.Context, in *ComposeFirewallGlobalBoardRequest, opts ...grpc.CallOption) (*ComposeFirewallGlobalBoardResponse, error) {
	out := new(ComposeFirewallGlobalBoardResponse)
	err := c.cc.Invoke(ctx, FirewallService_ComposeFirewallGlobalBoard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firewallServiceClient) NotifyHTTPFirewallEvent(ctx context.Context, in *NotifyHTTPFirewallEventRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, FirewallService_NotifyHTTPFirewallEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firewallServiceClient) CountFirewallDailyBlocks(ctx context.Context, in *CountFirewallDailyBlocksRequest, opts ...grpc.CallOption) (*CountFirewallDailyBlocksResponse, error) {
	out := new(CountFirewallDailyBlocksResponse)
	err := c.cc.Invoke(ctx, FirewallService_CountFirewallDailyBlocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirewallServiceServer is the server API for FirewallService service.
// All implementations should embed UnimplementedFirewallServiceServer
// for forward compatibility
type FirewallServiceServer interface {
	// 组合看板数据
	ComposeFirewallGlobalBoard(context.Context, *ComposeFirewallGlobalBoardRequest) (*ComposeFirewallGlobalBoardResponse, error)
	// 发送告警(notify)消息
	NotifyHTTPFirewallEvent(context.Context, *NotifyHTTPFirewallEventRequest) (*RPCSuccess, error)
	// 读取当前Block动作次数
	CountFirewallDailyBlocks(context.Context, *CountFirewallDailyBlocksRequest) (*CountFirewallDailyBlocksResponse, error)
}

// UnimplementedFirewallServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFirewallServiceServer struct {
}

func (UnimplementedFirewallServiceServer) ComposeFirewallGlobalBoard(context.Context, *ComposeFirewallGlobalBoardRequest) (*ComposeFirewallGlobalBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeFirewallGlobalBoard not implemented")
}
func (UnimplementedFirewallServiceServer) NotifyHTTPFirewallEvent(context.Context, *NotifyHTTPFirewallEventRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyHTTPFirewallEvent not implemented")
}
func (UnimplementedFirewallServiceServer) CountFirewallDailyBlocks(context.Context, *CountFirewallDailyBlocksRequest) (*CountFirewallDailyBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFirewallDailyBlocks not implemented")
}

// UnsafeFirewallServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirewallServiceServer will
// result in compilation errors.
type UnsafeFirewallServiceServer interface {
	mustEmbedUnimplementedFirewallServiceServer()
}

func RegisterFirewallServiceServer(s grpc.ServiceRegistrar, srv FirewallServiceServer) {
	s.RegisterService(&FirewallService_ServiceDesc, srv)
}

func _FirewallService_ComposeFirewallGlobalBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeFirewallGlobalBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirewallServiceServer).ComposeFirewallGlobalBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirewallService_ComposeFirewallGlobalBoard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirewallServiceServer).ComposeFirewallGlobalBoard(ctx, req.(*ComposeFirewallGlobalBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirewallService_NotifyHTTPFirewallEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyHTTPFirewallEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirewallServiceServer).NotifyHTTPFirewallEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirewallService_NotifyHTTPFirewallEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirewallServiceServer).NotifyHTTPFirewallEvent(ctx, req.(*NotifyHTTPFirewallEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirewallService_CountFirewallDailyBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFirewallDailyBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirewallServiceServer).CountFirewallDailyBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirewallService_CountFirewallDailyBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirewallServiceServer).CountFirewallDailyBlocks(ctx, req.(*CountFirewallDailyBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FirewallService_ServiceDesc is the grpc.ServiceDesc for FirewallService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FirewallService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FirewallService",
	HandlerType: (*FirewallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "composeFirewallGlobalBoard",
			Handler:    _FirewallService_ComposeFirewallGlobalBoard_Handler,
		},
		{
			MethodName: "notifyHTTPFirewallEvent",
			Handler:    _FirewallService_NotifyHTTPFirewallEvent_Handler,
		},
		{
			MethodName: "countFirewallDailyBlocks",
			Handler:    _FirewallService_CountFirewallDailyBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_firewall.proto",
}
