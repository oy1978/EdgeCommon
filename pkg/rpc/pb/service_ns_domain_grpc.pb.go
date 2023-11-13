// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ns_domain.proto

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
	NSDomainService_CreateNSDomain_FullMethodName                   = "/pb.NSDomainService/createNSDomain"
	NSDomainService_CreateNSDomains_FullMethodName                  = "/pb.NSDomainService/createNSDomains"
	NSDomainService_UpdateNSDomain_FullMethodName                   = "/pb.NSDomainService/updateNSDomain"
	NSDomainService_UpdateNSDomainStatus_FullMethodName             = "/pb.NSDomainService/updateNSDomainStatus"
	NSDomainService_DeleteNSDomain_FullMethodName                   = "/pb.NSDomainService/deleteNSDomain"
	NSDomainService_DeleteNSDomains_FullMethodName                  = "/pb.NSDomainService/deleteNSDomains"
	NSDomainService_FindNSDomain_FullMethodName                     = "/pb.NSDomainService/findNSDomain"
	NSDomainService_FindNSDomainWithName_FullMethodName             = "/pb.NSDomainService/findNSDomainWithName"
	NSDomainService_FindVerifiedNSDomainOnCluster_FullMethodName    = "/pb.NSDomainService/findVerifiedNSDomainOnCluster"
	NSDomainService_CountAllNSDomains_FullMethodName                = "/pb.NSDomainService/countAllNSDomains"
	NSDomainService_ListNSDomains_FullMethodName                    = "/pb.NSDomainService/listNSDomains"
	NSDomainService_ListNSDomainsAfterVersion_FullMethodName        = "/pb.NSDomainService/listNSDomainsAfterVersion"
	NSDomainService_FindNSDomainTSIG_FullMethodName                 = "/pb.NSDomainService/findNSDomainTSIG"
	NSDomainService_UpdateNSDomainTSIG_FullMethodName               = "/pb.NSDomainService/updateNSDomainTSIG"
	NSDomainService_ExistNSDomains_FullMethodName                   = "/pb.NSDomainService/existNSDomains"
	NSDomainService_ExistVerifiedNSDomains_FullMethodName           = "/pb.NSDomainService/existVerifiedNSDomains"
	NSDomainService_FindNSDomainVerifyingInfo_FullMethodName        = "/pb.NSDomainService/findNSDomainVerifyingInfo"
	NSDomainService_VerifyNSDomain_FullMethodName                   = "/pb.NSDomainService/verifyNSDomain"
	NSDomainService_FindNSDomainRecordsHealthCheck_FullMethodName   = "/pb.NSDomainService/findNSDomainRecordsHealthCheck"
	NSDomainService_UpdateNSDomainRecordsHealthCheck_FullMethodName = "/pb.NSDomainService/updateNSDomainRecordsHealthCheck"
)

// NSDomainServiceClient is the client API for NSDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSDomainServiceClient interface {
	// 创建单个域名
	CreateNSDomain(ctx context.Context, in *CreateNSDomainRequest, opts ...grpc.CallOption) (*CreateNSDomainResponse, error)
	// 批量创建域名
	CreateNSDomains(ctx context.Context, in *CreateNSDomainsRequest, opts ...grpc.CallOption) (*CreateNSDomainsResponse, error)
	// 修改域名
	UpdateNSDomain(ctx context.Context, in *UpdateNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改域名状态
	UpdateNSDomainStatus(ctx context.Context, in *UpdateNSDomainStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除域名
	DeleteNSDomain(ctx context.Context, in *DeleteNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 批量删除域名
	DeleteNSDomains(ctx context.Context, in *DeleteNSDomainsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个域名
	FindNSDomain(ctx context.Context, in *FindNSDomainRequest, opts ...grpc.CallOption) (*FindNSDomainResponse, error)
	// 根据域名名称查找域名
	FindNSDomainWithName(ctx context.Context, in *FindNSDomainWithNameRequest, opts ...grpc.CallOption) (*FindNSDomainWithNameResponse, error)
	// 根据域名名称查找集群中的已验证域名
	FindVerifiedNSDomainOnCluster(ctx context.Context, in *FindVerifiedNSDomainOnClusterRequest, opts ...grpc.CallOption) (*FindVerifiedNSDomainOnClusterResponse, error)
	// 计算域名数量
	CountAllNSDomains(ctx context.Context, in *CountAllNSDomainsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页域名
	ListNSDomains(ctx context.Context, in *ListNSDomainsRequest, opts ...grpc.CallOption) (*ListNSDomainsResponse, error)
	// 根据版本列出一组域名
	ListNSDomainsAfterVersion(ctx context.Context, in *ListNSDomainsAfterVersionRequest, opts ...grpc.CallOption) (*ListNSDomainsAfterVersionResponse, error)
	// 查找TSIG配置
	FindNSDomainTSIG(ctx context.Context, in *FindNSDomainTSIGRequest, opts ...grpc.CallOption) (*FindNSDomainTSIGResponse, error)
	// 修改TSIG配置
	UpdateNSDomainTSIG(ctx context.Context, in *UpdateNSDomainTSIGRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 检查一组域名是否在用户账户中存在
	ExistNSDomains(ctx context.Context, in *ExistNSDomainsRequest, opts ...grpc.CallOption) (*ExistNSDomainsResponse, error)
	// 检查一组域名是否已通过验证
	ExistVerifiedNSDomains(ctx context.Context, in *ExistVerifiedNSDomainsRequest, opts ...grpc.CallOption) (*ExistVerifiedNSDomainsResponse, error)
	// 获取域名验证信息
	FindNSDomainVerifyingInfo(ctx context.Context, in *FindNSDomainVerifyingInfoRequest, opts ...grpc.CallOption) (*FindNSDomainVerifyingInfoResponse, error)
	// 验证域名信息
	VerifyNSDomain(ctx context.Context, in *VerifyNSDomainRequest, opts ...grpc.CallOption) (*VerifyNSDomainResponse, error)
	// 查询记录健康检查全局设置
	FindNSDomainRecordsHealthCheck(ctx context.Context, in *FindNSDomainRecordsHealthCheckRequest, opts ...grpc.CallOption) (*FindNSDomainRecordsHealthCheckResponse, error)
	// 修改记录健康检查全局设置
	UpdateNSDomainRecordsHealthCheck(ctx context.Context, in *UpdateNSDomainRecordsHealthCheckRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type nSDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSDomainServiceClient(cc grpc.ClientConnInterface) NSDomainServiceClient {
	return &nSDomainServiceClient{cc}
}

func (c *nSDomainServiceClient) CreateNSDomain(ctx context.Context, in *CreateNSDomainRequest, opts ...grpc.CallOption) (*CreateNSDomainResponse, error) {
	out := new(CreateNSDomainResponse)
	err := c.cc.Invoke(ctx, NSDomainService_CreateNSDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) CreateNSDomains(ctx context.Context, in *CreateNSDomainsRequest, opts ...grpc.CallOption) (*CreateNSDomainsResponse, error) {
	out := new(CreateNSDomainsResponse)
	err := c.cc.Invoke(ctx, NSDomainService_CreateNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) UpdateNSDomain(ctx context.Context, in *UpdateNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_UpdateNSDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) UpdateNSDomainStatus(ctx context.Context, in *UpdateNSDomainStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_UpdateNSDomainStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) DeleteNSDomain(ctx context.Context, in *DeleteNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_DeleteNSDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) DeleteNSDomains(ctx context.Context, in *DeleteNSDomainsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_DeleteNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindNSDomain(ctx context.Context, in *FindNSDomainRequest, opts ...grpc.CallOption) (*FindNSDomainResponse, error) {
	out := new(FindNSDomainResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindNSDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindNSDomainWithName(ctx context.Context, in *FindNSDomainWithNameRequest, opts ...grpc.CallOption) (*FindNSDomainWithNameResponse, error) {
	out := new(FindNSDomainWithNameResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindNSDomainWithName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindVerifiedNSDomainOnCluster(ctx context.Context, in *FindVerifiedNSDomainOnClusterRequest, opts ...grpc.CallOption) (*FindVerifiedNSDomainOnClusterResponse, error) {
	out := new(FindVerifiedNSDomainOnClusterResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindVerifiedNSDomainOnCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) CountAllNSDomains(ctx context.Context, in *CountAllNSDomainsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, NSDomainService_CountAllNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) ListNSDomains(ctx context.Context, in *ListNSDomainsRequest, opts ...grpc.CallOption) (*ListNSDomainsResponse, error) {
	out := new(ListNSDomainsResponse)
	err := c.cc.Invoke(ctx, NSDomainService_ListNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) ListNSDomainsAfterVersion(ctx context.Context, in *ListNSDomainsAfterVersionRequest, opts ...grpc.CallOption) (*ListNSDomainsAfterVersionResponse, error) {
	out := new(ListNSDomainsAfterVersionResponse)
	err := c.cc.Invoke(ctx, NSDomainService_ListNSDomainsAfterVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindNSDomainTSIG(ctx context.Context, in *FindNSDomainTSIGRequest, opts ...grpc.CallOption) (*FindNSDomainTSIGResponse, error) {
	out := new(FindNSDomainTSIGResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindNSDomainTSIG_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) UpdateNSDomainTSIG(ctx context.Context, in *UpdateNSDomainTSIGRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_UpdateNSDomainTSIG_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) ExistNSDomains(ctx context.Context, in *ExistNSDomainsRequest, opts ...grpc.CallOption) (*ExistNSDomainsResponse, error) {
	out := new(ExistNSDomainsResponse)
	err := c.cc.Invoke(ctx, NSDomainService_ExistNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) ExistVerifiedNSDomains(ctx context.Context, in *ExistVerifiedNSDomainsRequest, opts ...grpc.CallOption) (*ExistVerifiedNSDomainsResponse, error) {
	out := new(ExistVerifiedNSDomainsResponse)
	err := c.cc.Invoke(ctx, NSDomainService_ExistVerifiedNSDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindNSDomainVerifyingInfo(ctx context.Context, in *FindNSDomainVerifyingInfoRequest, opts ...grpc.CallOption) (*FindNSDomainVerifyingInfoResponse, error) {
	out := new(FindNSDomainVerifyingInfoResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindNSDomainVerifyingInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) VerifyNSDomain(ctx context.Context, in *VerifyNSDomainRequest, opts ...grpc.CallOption) (*VerifyNSDomainResponse, error) {
	out := new(VerifyNSDomainResponse)
	err := c.cc.Invoke(ctx, NSDomainService_VerifyNSDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) FindNSDomainRecordsHealthCheck(ctx context.Context, in *FindNSDomainRecordsHealthCheckRequest, opts ...grpc.CallOption) (*FindNSDomainRecordsHealthCheckResponse, error) {
	out := new(FindNSDomainRecordsHealthCheckResponse)
	err := c.cc.Invoke(ctx, NSDomainService_FindNSDomainRecordsHealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainServiceClient) UpdateNSDomainRecordsHealthCheck(ctx context.Context, in *UpdateNSDomainRecordsHealthCheckRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainService_UpdateNSDomainRecordsHealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSDomainServiceServer is the server API for NSDomainService service.
// All implementations should embed UnimplementedNSDomainServiceServer
// for forward compatibility
type NSDomainServiceServer interface {
	// 创建单个域名
	CreateNSDomain(context.Context, *CreateNSDomainRequest) (*CreateNSDomainResponse, error)
	// 批量创建域名
	CreateNSDomains(context.Context, *CreateNSDomainsRequest) (*CreateNSDomainsResponse, error)
	// 修改域名
	UpdateNSDomain(context.Context, *UpdateNSDomainRequest) (*RPCSuccess, error)
	// 修改域名状态
	UpdateNSDomainStatus(context.Context, *UpdateNSDomainStatusRequest) (*RPCSuccess, error)
	// 删除域名
	DeleteNSDomain(context.Context, *DeleteNSDomainRequest) (*RPCSuccess, error)
	// 批量删除域名
	DeleteNSDomains(context.Context, *DeleteNSDomainsRequest) (*RPCSuccess, error)
	// 查找单个域名
	FindNSDomain(context.Context, *FindNSDomainRequest) (*FindNSDomainResponse, error)
	// 根据域名名称查找域名
	FindNSDomainWithName(context.Context, *FindNSDomainWithNameRequest) (*FindNSDomainWithNameResponse, error)
	// 根据域名名称查找集群中的已验证域名
	FindVerifiedNSDomainOnCluster(context.Context, *FindVerifiedNSDomainOnClusterRequest) (*FindVerifiedNSDomainOnClusterResponse, error)
	// 计算域名数量
	CountAllNSDomains(context.Context, *CountAllNSDomainsRequest) (*RPCCountResponse, error)
	// 列出单页域名
	ListNSDomains(context.Context, *ListNSDomainsRequest) (*ListNSDomainsResponse, error)
	// 根据版本列出一组域名
	ListNSDomainsAfterVersion(context.Context, *ListNSDomainsAfterVersionRequest) (*ListNSDomainsAfterVersionResponse, error)
	// 查找TSIG配置
	FindNSDomainTSIG(context.Context, *FindNSDomainTSIGRequest) (*FindNSDomainTSIGResponse, error)
	// 修改TSIG配置
	UpdateNSDomainTSIG(context.Context, *UpdateNSDomainTSIGRequest) (*RPCSuccess, error)
	// 检查一组域名是否在用户账户中存在
	ExistNSDomains(context.Context, *ExistNSDomainsRequest) (*ExistNSDomainsResponse, error)
	// 检查一组域名是否已通过验证
	ExistVerifiedNSDomains(context.Context, *ExistVerifiedNSDomainsRequest) (*ExistVerifiedNSDomainsResponse, error)
	// 获取域名验证信息
	FindNSDomainVerifyingInfo(context.Context, *FindNSDomainVerifyingInfoRequest) (*FindNSDomainVerifyingInfoResponse, error)
	// 验证域名信息
	VerifyNSDomain(context.Context, *VerifyNSDomainRequest) (*VerifyNSDomainResponse, error)
	// 查询记录健康检查全局设置
	FindNSDomainRecordsHealthCheck(context.Context, *FindNSDomainRecordsHealthCheckRequest) (*FindNSDomainRecordsHealthCheckResponse, error)
	// 修改记录健康检查全局设置
	UpdateNSDomainRecordsHealthCheck(context.Context, *UpdateNSDomainRecordsHealthCheckRequest) (*RPCSuccess, error)
}

// UnimplementedNSDomainServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNSDomainServiceServer struct {
}

func (UnimplementedNSDomainServiceServer) CreateNSDomain(context.Context, *CreateNSDomainRequest) (*CreateNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNSDomain not implemented")
}
func (UnimplementedNSDomainServiceServer) CreateNSDomains(context.Context, *CreateNSDomainsRequest) (*CreateNSDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) UpdateNSDomain(context.Context, *UpdateNSDomainRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNSDomain not implemented")
}
func (UnimplementedNSDomainServiceServer) UpdateNSDomainStatus(context.Context, *UpdateNSDomainStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNSDomainStatus not implemented")
}
func (UnimplementedNSDomainServiceServer) DeleteNSDomain(context.Context, *DeleteNSDomainRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNSDomain not implemented")
}
func (UnimplementedNSDomainServiceServer) DeleteNSDomains(context.Context, *DeleteNSDomainsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) FindNSDomain(context.Context, *FindNSDomainRequest) (*FindNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomain not implemented")
}
func (UnimplementedNSDomainServiceServer) FindNSDomainWithName(context.Context, *FindNSDomainWithNameRequest) (*FindNSDomainWithNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomainWithName not implemented")
}
func (UnimplementedNSDomainServiceServer) FindVerifiedNSDomainOnCluster(context.Context, *FindVerifiedNSDomainOnClusterRequest) (*FindVerifiedNSDomainOnClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVerifiedNSDomainOnCluster not implemented")
}
func (UnimplementedNSDomainServiceServer) CountAllNSDomains(context.Context, *CountAllNSDomainsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) ListNSDomains(context.Context, *ListNSDomainsRequest) (*ListNSDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) ListNSDomainsAfterVersion(context.Context, *ListNSDomainsAfterVersionRequest) (*ListNSDomainsAfterVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNSDomainsAfterVersion not implemented")
}
func (UnimplementedNSDomainServiceServer) FindNSDomainTSIG(context.Context, *FindNSDomainTSIGRequest) (*FindNSDomainTSIGResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomainTSIG not implemented")
}
func (UnimplementedNSDomainServiceServer) UpdateNSDomainTSIG(context.Context, *UpdateNSDomainTSIGRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNSDomainTSIG not implemented")
}
func (UnimplementedNSDomainServiceServer) ExistNSDomains(context.Context, *ExistNSDomainsRequest) (*ExistNSDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) ExistVerifiedNSDomains(context.Context, *ExistVerifiedNSDomainsRequest) (*ExistVerifiedNSDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistVerifiedNSDomains not implemented")
}
func (UnimplementedNSDomainServiceServer) FindNSDomainVerifyingInfo(context.Context, *FindNSDomainVerifyingInfoRequest) (*FindNSDomainVerifyingInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomainVerifyingInfo not implemented")
}
func (UnimplementedNSDomainServiceServer) VerifyNSDomain(context.Context, *VerifyNSDomainRequest) (*VerifyNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyNSDomain not implemented")
}
func (UnimplementedNSDomainServiceServer) FindNSDomainRecordsHealthCheck(context.Context, *FindNSDomainRecordsHealthCheckRequest) (*FindNSDomainRecordsHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomainRecordsHealthCheck not implemented")
}
func (UnimplementedNSDomainServiceServer) UpdateNSDomainRecordsHealthCheck(context.Context, *UpdateNSDomainRecordsHealthCheckRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNSDomainRecordsHealthCheck not implemented")
}

// UnsafeNSDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSDomainServiceServer will
// result in compilation errors.
type UnsafeNSDomainServiceServer interface {
	mustEmbedUnimplementedNSDomainServiceServer()
}

func RegisterNSDomainServiceServer(s grpc.ServiceRegistrar, srv NSDomainServiceServer) {
	s.RegisterService(&NSDomainService_ServiceDesc, srv)
}

func _NSDomainService_CreateNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).CreateNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_CreateNSDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).CreateNSDomain(ctx, req.(*CreateNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_CreateNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).CreateNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_CreateNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).CreateNSDomains(ctx, req.(*CreateNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_UpdateNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).UpdateNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_UpdateNSDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).UpdateNSDomain(ctx, req.(*UpdateNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_UpdateNSDomainStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNSDomainStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).UpdateNSDomainStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_UpdateNSDomainStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).UpdateNSDomainStatus(ctx, req.(*UpdateNSDomainStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_DeleteNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).DeleteNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_DeleteNSDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).DeleteNSDomain(ctx, req.(*DeleteNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_DeleteNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).DeleteNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_DeleteNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).DeleteNSDomains(ctx, req.(*DeleteNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindNSDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindNSDomain(ctx, req.(*FindNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindNSDomainWithName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainWithNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindNSDomainWithName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindNSDomainWithName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindNSDomainWithName(ctx, req.(*FindNSDomainWithNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindVerifiedNSDomainOnCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVerifiedNSDomainOnClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindVerifiedNSDomainOnCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindVerifiedNSDomainOnCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindVerifiedNSDomainOnCluster(ctx, req.(*FindVerifiedNSDomainOnClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_CountAllNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).CountAllNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_CountAllNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).CountAllNSDomains(ctx, req.(*CountAllNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_ListNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).ListNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_ListNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).ListNSDomains(ctx, req.(*ListNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_ListNSDomainsAfterVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNSDomainsAfterVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).ListNSDomainsAfterVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_ListNSDomainsAfterVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).ListNSDomainsAfterVersion(ctx, req.(*ListNSDomainsAfterVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindNSDomainTSIG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainTSIGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindNSDomainTSIG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindNSDomainTSIG_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindNSDomainTSIG(ctx, req.(*FindNSDomainTSIGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_UpdateNSDomainTSIG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNSDomainTSIGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).UpdateNSDomainTSIG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_UpdateNSDomainTSIG_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).UpdateNSDomainTSIG(ctx, req.(*UpdateNSDomainTSIGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_ExistNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).ExistNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_ExistNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).ExistNSDomains(ctx, req.(*ExistNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_ExistVerifiedNSDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistVerifiedNSDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).ExistVerifiedNSDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_ExistVerifiedNSDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).ExistVerifiedNSDomains(ctx, req.(*ExistVerifiedNSDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindNSDomainVerifyingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainVerifyingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindNSDomainVerifyingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindNSDomainVerifyingInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindNSDomainVerifyingInfo(ctx, req.(*FindNSDomainVerifyingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_VerifyNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).VerifyNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_VerifyNSDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).VerifyNSDomain(ctx, req.(*VerifyNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_FindNSDomainRecordsHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainRecordsHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).FindNSDomainRecordsHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_FindNSDomainRecordsHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).FindNSDomainRecordsHealthCheck(ctx, req.(*FindNSDomainRecordsHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainService_UpdateNSDomainRecordsHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNSDomainRecordsHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainServiceServer).UpdateNSDomainRecordsHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainService_UpdateNSDomainRecordsHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainServiceServer).UpdateNSDomainRecordsHealthCheck(ctx, req.(*UpdateNSDomainRecordsHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSDomainService_ServiceDesc is the grpc.ServiceDesc for NSDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSDomainService",
	HandlerType: (*NSDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNSDomain",
			Handler:    _NSDomainService_CreateNSDomain_Handler,
		},
		{
			MethodName: "createNSDomains",
			Handler:    _NSDomainService_CreateNSDomains_Handler,
		},
		{
			MethodName: "updateNSDomain",
			Handler:    _NSDomainService_UpdateNSDomain_Handler,
		},
		{
			MethodName: "updateNSDomainStatus",
			Handler:    _NSDomainService_UpdateNSDomainStatus_Handler,
		},
		{
			MethodName: "deleteNSDomain",
			Handler:    _NSDomainService_DeleteNSDomain_Handler,
		},
		{
			MethodName: "deleteNSDomains",
			Handler:    _NSDomainService_DeleteNSDomains_Handler,
		},
		{
			MethodName: "findNSDomain",
			Handler:    _NSDomainService_FindNSDomain_Handler,
		},
		{
			MethodName: "findNSDomainWithName",
			Handler:    _NSDomainService_FindNSDomainWithName_Handler,
		},
		{
			MethodName: "findVerifiedNSDomainOnCluster",
			Handler:    _NSDomainService_FindVerifiedNSDomainOnCluster_Handler,
		},
		{
			MethodName: "countAllNSDomains",
			Handler:    _NSDomainService_CountAllNSDomains_Handler,
		},
		{
			MethodName: "listNSDomains",
			Handler:    _NSDomainService_ListNSDomains_Handler,
		},
		{
			MethodName: "listNSDomainsAfterVersion",
			Handler:    _NSDomainService_ListNSDomainsAfterVersion_Handler,
		},
		{
			MethodName: "findNSDomainTSIG",
			Handler:    _NSDomainService_FindNSDomainTSIG_Handler,
		},
		{
			MethodName: "updateNSDomainTSIG",
			Handler:    _NSDomainService_UpdateNSDomainTSIG_Handler,
		},
		{
			MethodName: "existNSDomains",
			Handler:    _NSDomainService_ExistNSDomains_Handler,
		},
		{
			MethodName: "existVerifiedNSDomains",
			Handler:    _NSDomainService_ExistVerifiedNSDomains_Handler,
		},
		{
			MethodName: "findNSDomainVerifyingInfo",
			Handler:    _NSDomainService_FindNSDomainVerifyingInfo_Handler,
		},
		{
			MethodName: "verifyNSDomain",
			Handler:    _NSDomainService_VerifyNSDomain_Handler,
		},
		{
			MethodName: "findNSDomainRecordsHealthCheck",
			Handler:    _NSDomainService_FindNSDomainRecordsHealthCheck_Handler,
		},
		{
			MethodName: "updateNSDomainRecordsHealthCheck",
			Handler:    _NSDomainService_UpdateNSDomainRecordsHealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns_domain.proto",
}
