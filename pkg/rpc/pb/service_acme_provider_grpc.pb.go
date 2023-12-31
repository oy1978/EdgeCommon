// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_acme_provider.proto

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
	ACMEProviderService_FindAllACMEProviders_FullMethodName     = "/pb.ACMEProviderService/findAllACMEProviders"
	ACMEProviderService_FindACMEProviderWithCode_FullMethodName = "/pb.ACMEProviderService/findACMEProviderWithCode"
)

// ACMEProviderServiceClient is the client API for ACMEProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ACMEProviderServiceClient interface {
	// 查找所有的服务商
	FindAllACMEProviders(ctx context.Context, in *FindAllACMEProvidersRequest, opts ...grpc.CallOption) (*FindAllACMEProvidersResponse, error)
	// 根据代号查找服务商
	FindACMEProviderWithCode(ctx context.Context, in *FindACMEProviderWithCodeRequest, opts ...grpc.CallOption) (*FindACMEProviderWithCodeResponse, error)
}

type aCMEProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewACMEProviderServiceClient(cc grpc.ClientConnInterface) ACMEProviderServiceClient {
	return &aCMEProviderServiceClient{cc}
}

func (c *aCMEProviderServiceClient) FindAllACMEProviders(ctx context.Context, in *FindAllACMEProvidersRequest, opts ...grpc.CallOption) (*FindAllACMEProvidersResponse, error) {
	out := new(FindAllACMEProvidersResponse)
	err := c.cc.Invoke(ctx, ACMEProviderService_FindAllACMEProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCMEProviderServiceClient) FindACMEProviderWithCode(ctx context.Context, in *FindACMEProviderWithCodeRequest, opts ...grpc.CallOption) (*FindACMEProviderWithCodeResponse, error) {
	out := new(FindACMEProviderWithCodeResponse)
	err := c.cc.Invoke(ctx, ACMEProviderService_FindACMEProviderWithCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACMEProviderServiceServer is the server API for ACMEProviderService service.
// All implementations should embed UnimplementedACMEProviderServiceServer
// for forward compatibility
type ACMEProviderServiceServer interface {
	// 查找所有的服务商
	FindAllACMEProviders(context.Context, *FindAllACMEProvidersRequest) (*FindAllACMEProvidersResponse, error)
	// 根据代号查找服务商
	FindACMEProviderWithCode(context.Context, *FindACMEProviderWithCodeRequest) (*FindACMEProviderWithCodeResponse, error)
}

// UnimplementedACMEProviderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedACMEProviderServiceServer struct {
}

func (UnimplementedACMEProviderServiceServer) FindAllACMEProviders(context.Context, *FindAllACMEProvidersRequest) (*FindAllACMEProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllACMEProviders not implemented")
}
func (UnimplementedACMEProviderServiceServer) FindACMEProviderWithCode(context.Context, *FindACMEProviderWithCodeRequest) (*FindACMEProviderWithCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindACMEProviderWithCode not implemented")
}

// UnsafeACMEProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ACMEProviderServiceServer will
// result in compilation errors.
type UnsafeACMEProviderServiceServer interface {
	mustEmbedUnimplementedACMEProviderServiceServer()
}

func RegisterACMEProviderServiceServer(s grpc.ServiceRegistrar, srv ACMEProviderServiceServer) {
	s.RegisterService(&ACMEProviderService_ServiceDesc, srv)
}

func _ACMEProviderService_FindAllACMEProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllACMEProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEProviderServiceServer).FindAllACMEProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ACMEProviderService_FindAllACMEProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEProviderServiceServer).FindAllACMEProviders(ctx, req.(*FindAllACMEProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACMEProviderService_FindACMEProviderWithCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindACMEProviderWithCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACMEProviderServiceServer).FindACMEProviderWithCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ACMEProviderService_FindACMEProviderWithCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACMEProviderServiceServer).FindACMEProviderWithCode(ctx, req.(*FindACMEProviderWithCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ACMEProviderService_ServiceDesc is the grpc.ServiceDesc for ACMEProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ACMEProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ACMEProviderService",
	HandlerType: (*ACMEProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllACMEProviders",
			Handler:    _ACMEProviderService_FindAllACMEProviders_Handler,
		},
		{
			MethodName: "findACMEProviderWithCode",
			Handler:    _ACMEProviderService_FindACMEProviderWithCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_acme_provider.proto",
}
