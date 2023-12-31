// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_client_agent_ip.proto

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
	ClientAgentIPService_CreateClientAgentIPs_FullMethodName      = "/pb.ClientAgentIPService/createClientAgentIPs"
	ClientAgentIPService_ListClientAgentIPsAfterId_FullMethodName = "/pb.ClientAgentIPService/listClientAgentIPsAfterId"
)

// ClientAgentIPServiceClient is the client API for ClientAgentIPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientAgentIPServiceClient interface {
	// 创建一组IP
	CreateClientAgentIPs(ctx context.Context, in *CreateClientAgentIPsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询最新的IP
	ListClientAgentIPsAfterId(ctx context.Context, in *ListClientAgentIPsAfterIdRequest, opts ...grpc.CallOption) (*ListClientAgentIPsAfterIdResponse, error)
}

type clientAgentIPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientAgentIPServiceClient(cc grpc.ClientConnInterface) ClientAgentIPServiceClient {
	return &clientAgentIPServiceClient{cc}
}

func (c *clientAgentIPServiceClient) CreateClientAgentIPs(ctx context.Context, in *CreateClientAgentIPsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ClientAgentIPService_CreateClientAgentIPs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAgentIPServiceClient) ListClientAgentIPsAfterId(ctx context.Context, in *ListClientAgentIPsAfterIdRequest, opts ...grpc.CallOption) (*ListClientAgentIPsAfterIdResponse, error) {
	out := new(ListClientAgentIPsAfterIdResponse)
	err := c.cc.Invoke(ctx, ClientAgentIPService_ListClientAgentIPsAfterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAgentIPServiceServer is the server API for ClientAgentIPService service.
// All implementations should embed UnimplementedClientAgentIPServiceServer
// for forward compatibility
type ClientAgentIPServiceServer interface {
	// 创建一组IP
	CreateClientAgentIPs(context.Context, *CreateClientAgentIPsRequest) (*RPCSuccess, error)
	// 查询最新的IP
	ListClientAgentIPsAfterId(context.Context, *ListClientAgentIPsAfterIdRequest) (*ListClientAgentIPsAfterIdResponse, error)
}

// UnimplementedClientAgentIPServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClientAgentIPServiceServer struct {
}

func (UnimplementedClientAgentIPServiceServer) CreateClientAgentIPs(context.Context, *CreateClientAgentIPsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientAgentIPs not implemented")
}
func (UnimplementedClientAgentIPServiceServer) ListClientAgentIPsAfterId(context.Context, *ListClientAgentIPsAfterIdRequest) (*ListClientAgentIPsAfterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientAgentIPsAfterId not implemented")
}

// UnsafeClientAgentIPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientAgentIPServiceServer will
// result in compilation errors.
type UnsafeClientAgentIPServiceServer interface {
	mustEmbedUnimplementedClientAgentIPServiceServer()
}

func RegisterClientAgentIPServiceServer(s grpc.ServiceRegistrar, srv ClientAgentIPServiceServer) {
	s.RegisterService(&ClientAgentIPService_ServiceDesc, srv)
}

func _ClientAgentIPService_CreateClientAgentIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientAgentIPsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAgentIPServiceServer).CreateClientAgentIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAgentIPService_CreateClientAgentIPs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAgentIPServiceServer).CreateClientAgentIPs(ctx, req.(*CreateClientAgentIPsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAgentIPService_ListClientAgentIPsAfterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientAgentIPsAfterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAgentIPServiceServer).ListClientAgentIPsAfterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAgentIPService_ListClientAgentIPsAfterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAgentIPServiceServer).ListClientAgentIPsAfterId(ctx, req.(*ListClientAgentIPsAfterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientAgentIPService_ServiceDesc is the grpc.ServiceDesc for ClientAgentIPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientAgentIPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClientAgentIPService",
	HandlerType: (*ClientAgentIPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createClientAgentIPs",
			Handler:    _ClientAgentIPService_CreateClientAgentIPs_Handler,
		},
		{
			MethodName: "listClientAgentIPsAfterId",
			Handler:    _ClientAgentIPService_ListClientAgentIPsAfterId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_client_agent_ip.proto",
}
