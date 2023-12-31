// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_node_cluster_metric_item.proto

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
	NodeClusterMetricItemService_EnableNodeClusterMetricItem_FullMethodName         = "/pb.NodeClusterMetricItemService/enableNodeClusterMetricItem"
	NodeClusterMetricItemService_DisableNodeClusterMetricItem_FullMethodName        = "/pb.NodeClusterMetricItemService/disableNodeClusterMetricItem"
	NodeClusterMetricItemService_FindAllNodeClusterMetricItems_FullMethodName       = "/pb.NodeClusterMetricItemService/findAllNodeClusterMetricItems"
	NodeClusterMetricItemService_ExistsNodeClusterMetricItem_FullMethodName         = "/pb.NodeClusterMetricItemService/existsNodeClusterMetricItem"
	NodeClusterMetricItemService_FindAllNodeClustersWithMetricItemId_FullMethodName = "/pb.NodeClusterMetricItemService/findAllNodeClustersWithMetricItemId"
)

// NodeClusterMetricItemServiceClient is the client API for NodeClusterMetricItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClusterMetricItemServiceClient interface {
	// 启用某个指标
	EnableNodeClusterMetricItem(ctx context.Context, in *EnableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 禁用某个指标
	DisableNodeClusterMetricItem(ctx context.Context, in *DisableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找集群中所有指标
	FindAllNodeClusterMetricItems(ctx context.Context, in *FindAllNodeClusterMetricItemsRequest, opts ...grpc.CallOption) (*FindAllNodeClusterMetricItemsResponse, error)
	// 检查是否已添加某个指标
	ExistsNodeClusterMetricItem(ctx context.Context, in *ExistsNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCExists, error)
	// 查找使用指标的集群
	FindAllNodeClustersWithMetricItemId(ctx context.Context, in *FindAllNodeClustersWithMetricItemIdRequest, opts ...grpc.CallOption) (*FindAllNodeClustersWithMetricItemIdResponse, error)
}

type nodeClusterMetricItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClusterMetricItemServiceClient(cc grpc.ClientConnInterface) NodeClusterMetricItemServiceClient {
	return &nodeClusterMetricItemServiceClient{cc}
}

func (c *nodeClusterMetricItemServiceClient) EnableNodeClusterMetricItem(ctx context.Context, in *EnableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodeClusterMetricItemService_EnableNodeClusterMetricItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) DisableNodeClusterMetricItem(ctx context.Context, in *DisableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodeClusterMetricItemService_DisableNodeClusterMetricItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) FindAllNodeClusterMetricItems(ctx context.Context, in *FindAllNodeClusterMetricItemsRequest, opts ...grpc.CallOption) (*FindAllNodeClusterMetricItemsResponse, error) {
	out := new(FindAllNodeClusterMetricItemsResponse)
	err := c.cc.Invoke(ctx, NodeClusterMetricItemService_FindAllNodeClusterMetricItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) ExistsNodeClusterMetricItem(ctx context.Context, in *ExistsNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCExists, error) {
	out := new(RPCExists)
	err := c.cc.Invoke(ctx, NodeClusterMetricItemService_ExistsNodeClusterMetricItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) FindAllNodeClustersWithMetricItemId(ctx context.Context, in *FindAllNodeClustersWithMetricItemIdRequest, opts ...grpc.CallOption) (*FindAllNodeClustersWithMetricItemIdResponse, error) {
	out := new(FindAllNodeClustersWithMetricItemIdResponse)
	err := c.cc.Invoke(ctx, NodeClusterMetricItemService_FindAllNodeClustersWithMetricItemId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeClusterMetricItemServiceServer is the server API for NodeClusterMetricItemService service.
// All implementations should embed UnimplementedNodeClusterMetricItemServiceServer
// for forward compatibility
type NodeClusterMetricItemServiceServer interface {
	// 启用某个指标
	EnableNodeClusterMetricItem(context.Context, *EnableNodeClusterMetricItemRequest) (*RPCSuccess, error)
	// 禁用某个指标
	DisableNodeClusterMetricItem(context.Context, *DisableNodeClusterMetricItemRequest) (*RPCSuccess, error)
	// 查找集群中所有指标
	FindAllNodeClusterMetricItems(context.Context, *FindAllNodeClusterMetricItemsRequest) (*FindAllNodeClusterMetricItemsResponse, error)
	// 检查是否已添加某个指标
	ExistsNodeClusterMetricItem(context.Context, *ExistsNodeClusterMetricItemRequest) (*RPCExists, error)
	// 查找使用指标的集群
	FindAllNodeClustersWithMetricItemId(context.Context, *FindAllNodeClustersWithMetricItemIdRequest) (*FindAllNodeClustersWithMetricItemIdResponse, error)
}

// UnimplementedNodeClusterMetricItemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNodeClusterMetricItemServiceServer struct {
}

func (UnimplementedNodeClusterMetricItemServiceServer) EnableNodeClusterMetricItem(context.Context, *EnableNodeClusterMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableNodeClusterMetricItem not implemented")
}
func (UnimplementedNodeClusterMetricItemServiceServer) DisableNodeClusterMetricItem(context.Context, *DisableNodeClusterMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableNodeClusterMetricItem not implemented")
}
func (UnimplementedNodeClusterMetricItemServiceServer) FindAllNodeClusterMetricItems(context.Context, *FindAllNodeClusterMetricItemsRequest) (*FindAllNodeClusterMetricItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllNodeClusterMetricItems not implemented")
}
func (UnimplementedNodeClusterMetricItemServiceServer) ExistsNodeClusterMetricItem(context.Context, *ExistsNodeClusterMetricItemRequest) (*RPCExists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsNodeClusterMetricItem not implemented")
}
func (UnimplementedNodeClusterMetricItemServiceServer) FindAllNodeClustersWithMetricItemId(context.Context, *FindAllNodeClustersWithMetricItemIdRequest) (*FindAllNodeClustersWithMetricItemIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllNodeClustersWithMetricItemId not implemented")
}

// UnsafeNodeClusterMetricItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeClusterMetricItemServiceServer will
// result in compilation errors.
type UnsafeNodeClusterMetricItemServiceServer interface {
	mustEmbedUnimplementedNodeClusterMetricItemServiceServer()
}

func RegisterNodeClusterMetricItemServiceServer(s grpc.ServiceRegistrar, srv NodeClusterMetricItemServiceServer) {
	s.RegisterService(&NodeClusterMetricItemService_ServiceDesc, srv)
}

func _NodeClusterMetricItemService_EnableNodeClusterMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableNodeClusterMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).EnableNodeClusterMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeClusterMetricItemService_EnableNodeClusterMetricItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).EnableNodeClusterMetricItem(ctx, req.(*EnableNodeClusterMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_DisableNodeClusterMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableNodeClusterMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).DisableNodeClusterMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeClusterMetricItemService_DisableNodeClusterMetricItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).DisableNodeClusterMetricItem(ctx, req.(*DisableNodeClusterMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_FindAllNodeClusterMetricItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllNodeClusterMetricItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClusterMetricItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeClusterMetricItemService_FindAllNodeClusterMetricItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClusterMetricItems(ctx, req.(*FindAllNodeClusterMetricItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_ExistsNodeClusterMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsNodeClusterMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).ExistsNodeClusterMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeClusterMetricItemService_ExistsNodeClusterMetricItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).ExistsNodeClusterMetricItem(ctx, req.(*ExistsNodeClusterMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_FindAllNodeClustersWithMetricItemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllNodeClustersWithMetricItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClustersWithMetricItemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeClusterMetricItemService_FindAllNodeClustersWithMetricItemId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClustersWithMetricItemId(ctx, req.(*FindAllNodeClustersWithMetricItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeClusterMetricItemService_ServiceDesc is the grpc.ServiceDesc for NodeClusterMetricItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeClusterMetricItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeClusterMetricItemService",
	HandlerType: (*NodeClusterMetricItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "enableNodeClusterMetricItem",
			Handler:    _NodeClusterMetricItemService_EnableNodeClusterMetricItem_Handler,
		},
		{
			MethodName: "disableNodeClusterMetricItem",
			Handler:    _NodeClusterMetricItemService_DisableNodeClusterMetricItem_Handler,
		},
		{
			MethodName: "findAllNodeClusterMetricItems",
			Handler:    _NodeClusterMetricItemService_FindAllNodeClusterMetricItems_Handler,
		},
		{
			MethodName: "existsNodeClusterMetricItem",
			Handler:    _NodeClusterMetricItemService_ExistsNodeClusterMetricItem_Handler,
		},
		{
			MethodName: "findAllNodeClustersWithMetricItemId",
			Handler:    _NodeClusterMetricItemService_FindAllNodeClustersWithMetricItemId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_cluster_metric_item.proto",
}
