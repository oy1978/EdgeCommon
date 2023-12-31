// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_server_region_country_monthly_stat.proto

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
	ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_FullMethodName = "/pb.ServerRegionCountryMonthlyStatService/findTopServerRegionCountryMonthlyStats"
)

// ServerRegionCountryMonthlyStatServiceClient is the client API for ServerRegionCountryMonthlyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerRegionCountryMonthlyStatServiceClient interface {
	// 查找前N个地区
	FindTopServerRegionCountryMonthlyStats(ctx context.Context, in *FindTopServerRegionCountryMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCountryMonthlyStatsResponse, error)
}

type serverRegionCountryMonthlyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerRegionCountryMonthlyStatServiceClient(cc grpc.ClientConnInterface) ServerRegionCountryMonthlyStatServiceClient {
	return &serverRegionCountryMonthlyStatServiceClient{cc}
}

func (c *serverRegionCountryMonthlyStatServiceClient) FindTopServerRegionCountryMonthlyStats(ctx context.Context, in *FindTopServerRegionCountryMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCountryMonthlyStatsResponse, error) {
	out := new(FindTopServerRegionCountryMonthlyStatsResponse)
	err := c.cc.Invoke(ctx, ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerRegionCountryMonthlyStatServiceServer is the server API for ServerRegionCountryMonthlyStatService service.
// All implementations should embed UnimplementedServerRegionCountryMonthlyStatServiceServer
// for forward compatibility
type ServerRegionCountryMonthlyStatServiceServer interface {
	// 查找前N个地区
	FindTopServerRegionCountryMonthlyStats(context.Context, *FindTopServerRegionCountryMonthlyStatsRequest) (*FindTopServerRegionCountryMonthlyStatsResponse, error)
}

// UnimplementedServerRegionCountryMonthlyStatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServerRegionCountryMonthlyStatServiceServer struct {
}

func (UnimplementedServerRegionCountryMonthlyStatServiceServer) FindTopServerRegionCountryMonthlyStats(context.Context, *FindTopServerRegionCountryMonthlyStatsRequest) (*FindTopServerRegionCountryMonthlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTopServerRegionCountryMonthlyStats not implemented")
}

// UnsafeServerRegionCountryMonthlyStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerRegionCountryMonthlyStatServiceServer will
// result in compilation errors.
type UnsafeServerRegionCountryMonthlyStatServiceServer interface {
	mustEmbedUnimplementedServerRegionCountryMonthlyStatServiceServer()
}

func RegisterServerRegionCountryMonthlyStatServiceServer(s grpc.ServiceRegistrar, srv ServerRegionCountryMonthlyStatServiceServer) {
	s.RegisterService(&ServerRegionCountryMonthlyStatService_ServiceDesc, srv)
}

func _ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTopServerRegionCountryMonthlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRegionCountryMonthlyStatServiceServer).FindTopServerRegionCountryMonthlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRegionCountryMonthlyStatServiceServer).FindTopServerRegionCountryMonthlyStats(ctx, req.(*FindTopServerRegionCountryMonthlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerRegionCountryMonthlyStatService_ServiceDesc is the grpc.ServiceDesc for ServerRegionCountryMonthlyStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerRegionCountryMonthlyStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerRegionCountryMonthlyStatService",
	HandlerType: (*ServerRegionCountryMonthlyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findTopServerRegionCountryMonthlyStats",
			Handler:    _ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_region_country_monthly_stat.proto",
}
