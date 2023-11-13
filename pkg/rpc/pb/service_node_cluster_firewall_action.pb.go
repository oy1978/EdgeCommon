// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_node_cluster_firewall_action.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建动作
type CreateNodeClusterFirewallActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64  `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EventLevel    string `protobuf:"bytes,3,opt,name=eventLevel,proto3" json:"eventLevel,omitempty"`
	Type          string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ParamsJSON    []byte `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
}

func (x *CreateNodeClusterFirewallActionRequest) Reset() {
	*x = CreateNodeClusterFirewallActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeClusterFirewallActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeClusterFirewallActionRequest) ProtoMessage() {}

func (x *CreateNodeClusterFirewallActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeClusterFirewallActionRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeClusterFirewallActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeClusterFirewallActionRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *CreateNodeClusterFirewallActionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNodeClusterFirewallActionRequest) GetEventLevel() string {
	if x != nil {
		return x.EventLevel
	}
	return ""
}

func (x *CreateNodeClusterFirewallActionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateNodeClusterFirewallActionRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

type NodeClusterFirewallActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallActionId int64 `protobuf:"varint,1,opt,name=nodeClusterFirewallActionId,proto3" json:"nodeClusterFirewallActionId,omitempty"`
}

func (x *NodeClusterFirewallActionResponse) Reset() {
	*x = NodeClusterFirewallActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeClusterFirewallActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeClusterFirewallActionResponse) ProtoMessage() {}

func (x *NodeClusterFirewallActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeClusterFirewallActionResponse.ProtoReflect.Descriptor instead.
func (*NodeClusterFirewallActionResponse) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{1}
}

func (x *NodeClusterFirewallActionResponse) GetNodeClusterFirewallActionId() int64 {
	if x != nil {
		return x.NodeClusterFirewallActionId
	}
	return 0
}

// 修改动作
type UpdateNodeClusterFirewallActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallActionId int64  `protobuf:"varint,1,opt,name=nodeClusterFirewallActionId,proto3" json:"nodeClusterFirewallActionId,omitempty"`
	Name                        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EventLevel                  string `protobuf:"bytes,3,opt,name=eventLevel,proto3" json:"eventLevel,omitempty"`
	Type                        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ParamsJSON                  []byte `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
}

func (x *UpdateNodeClusterFirewallActionRequest) Reset() {
	*x = UpdateNodeClusterFirewallActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeClusterFirewallActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeClusterFirewallActionRequest) ProtoMessage() {}

func (x *UpdateNodeClusterFirewallActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeClusterFirewallActionRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeClusterFirewallActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNodeClusterFirewallActionRequest) GetNodeClusterFirewallActionId() int64 {
	if x != nil {
		return x.NodeClusterFirewallActionId
	}
	return 0
}

func (x *UpdateNodeClusterFirewallActionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNodeClusterFirewallActionRequest) GetEventLevel() string {
	if x != nil {
		return x.EventLevel
	}
	return ""
}

func (x *UpdateNodeClusterFirewallActionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateNodeClusterFirewallActionRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

// 删除动作
type DeleteNodeClusterFirewallActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallActionId int64 `protobuf:"varint,1,opt,name=nodeClusterFirewallActionId,proto3" json:"nodeClusterFirewallActionId,omitempty"`
}

func (x *DeleteNodeClusterFirewallActionRequest) Reset() {
	*x = DeleteNodeClusterFirewallActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeClusterFirewallActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeClusterFirewallActionRequest) ProtoMessage() {}

func (x *DeleteNodeClusterFirewallActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeClusterFirewallActionRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeClusterFirewallActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNodeClusterFirewallActionRequest) GetNodeClusterFirewallActionId() int64 {
	if x != nil {
		return x.NodeClusterFirewallActionId
	}
	return 0
}

// 查询集群的所有动作
type FindAllEnabledNodeClusterFirewallActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
}

func (x *FindAllEnabledNodeClusterFirewallActionsRequest) Reset() {
	*x = FindAllEnabledNodeClusterFirewallActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeClusterFirewallActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeClusterFirewallActionsRequest) ProtoMessage() {}

func (x *FindAllEnabledNodeClusterFirewallActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeClusterFirewallActionsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeClusterFirewallActionsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllEnabledNodeClusterFirewallActionsRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

type FindAllEnabledNodeClusterFirewallActionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallActions []*NodeClusterFirewallAction `protobuf:"bytes,1,rep,name=nodeClusterFirewallActions,proto3" json:"nodeClusterFirewallActions,omitempty"`
}

func (x *FindAllEnabledNodeClusterFirewallActionsResponse) Reset() {
	*x = FindAllEnabledNodeClusterFirewallActionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeClusterFirewallActionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeClusterFirewallActionsResponse) ProtoMessage() {}

func (x *FindAllEnabledNodeClusterFirewallActionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeClusterFirewallActionsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeClusterFirewallActionsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllEnabledNodeClusterFirewallActionsResponse) GetNodeClusterFirewallActions() []*NodeClusterFirewallAction {
	if x != nil {
		return x.NodeClusterFirewallActions
	}
	return nil
}

// 查询单个动作
type FindEnabledNodeClusterFirewallActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallActionId int64 `protobuf:"varint,1,opt,name=nodeClusterFirewallActionId,proto3" json:"nodeClusterFirewallActionId,omitempty"`
}

func (x *FindEnabledNodeClusterFirewallActionRequest) Reset() {
	*x = FindEnabledNodeClusterFirewallActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledNodeClusterFirewallActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledNodeClusterFirewallActionRequest) ProtoMessage() {}

func (x *FindEnabledNodeClusterFirewallActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledNodeClusterFirewallActionRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledNodeClusterFirewallActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledNodeClusterFirewallActionRequest) GetNodeClusterFirewallActionId() int64 {
	if x != nil {
		return x.NodeClusterFirewallActionId
	}
	return 0
}

type FindEnabledNodeClusterFirewallActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterFirewallAction *NodeClusterFirewallAction `protobuf:"bytes,1,opt,name=nodeClusterFirewallAction,proto3" json:"nodeClusterFirewallAction,omitempty"`
}

func (x *FindEnabledNodeClusterFirewallActionResponse) Reset() {
	*x = FindEnabledNodeClusterFirewallActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledNodeClusterFirewallActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledNodeClusterFirewallActionResponse) ProtoMessage() {}

func (x *FindEnabledNodeClusterFirewallActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledNodeClusterFirewallActionResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledNodeClusterFirewallActionResponse) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledNodeClusterFirewallActionResponse) GetNodeClusterFirewallAction() *NodeClusterFirewallAction {
	if x != nil {
		return x.NodeClusterFirewallAction
	}
	return nil
}

// 计算动作数量
type CountAllEnabledNodeClusterFirewallActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
}

func (x *CountAllEnabledNodeClusterFirewallActionsRequest) Reset() {
	*x = CountAllEnabledNodeClusterFirewallActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_firewall_action_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledNodeClusterFirewallActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledNodeClusterFirewallActionsRequest) ProtoMessage() {}

func (x *CountAllEnabledNodeClusterFirewallActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_firewall_action_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledNodeClusterFirewallActionsRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledNodeClusterFirewallActionsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_firewall_action_proto_rawDescGZIP(), []int{8}
}

func (x *CountAllEnabledNodeClusterFirewallActionsRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

var File_service_node_cluster_firewall_action_proto protoreflect.FileDescriptor

var file_service_node_cluster_firewall_action_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a,
	0x26, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x65, 0x0a, 0x21, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x6e, 0x6f,
	0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xd2, 0x01, 0x0a,
	0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x6e, 0x6f,
	0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x22, 0x6a, 0x0a, 0x26, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1b, 0x6e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a,
	0x2f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x30, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x1a, 0x6e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1a,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6f, 0x0a, 0x2b, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1b, 0x6e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x2c,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x19,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x30, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x32, 0xf3, 0x05, 0x0a, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x1f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d,
	0x0a, 0x1f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x5d, 0x0a,
	0x1f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x95, 0x01, 0x0a,
	0x28, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x24, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x77, 0x0a, 0x29, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_cluster_firewall_action_proto_rawDescOnce sync.Once
	file_service_node_cluster_firewall_action_proto_rawDescData = file_service_node_cluster_firewall_action_proto_rawDesc
)

func file_service_node_cluster_firewall_action_proto_rawDescGZIP() []byte {
	file_service_node_cluster_firewall_action_proto_rawDescOnce.Do(func() {
		file_service_node_cluster_firewall_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_cluster_firewall_action_proto_rawDescData)
	})
	return file_service_node_cluster_firewall_action_proto_rawDescData
}

var file_service_node_cluster_firewall_action_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_node_cluster_firewall_action_proto_goTypes = []interface{}{
	(*CreateNodeClusterFirewallActionRequest)(nil),           // 0: pb.CreateNodeClusterFirewallActionRequest
	(*NodeClusterFirewallActionResponse)(nil),                // 1: pb.NodeClusterFirewallActionResponse
	(*UpdateNodeClusterFirewallActionRequest)(nil),           // 2: pb.UpdateNodeClusterFirewallActionRequest
	(*DeleteNodeClusterFirewallActionRequest)(nil),           // 3: pb.DeleteNodeClusterFirewallActionRequest
	(*FindAllEnabledNodeClusterFirewallActionsRequest)(nil),  // 4: pb.FindAllEnabledNodeClusterFirewallActionsRequest
	(*FindAllEnabledNodeClusterFirewallActionsResponse)(nil), // 5: pb.FindAllEnabledNodeClusterFirewallActionsResponse
	(*FindEnabledNodeClusterFirewallActionRequest)(nil),      // 6: pb.FindEnabledNodeClusterFirewallActionRequest
	(*FindEnabledNodeClusterFirewallActionResponse)(nil),     // 7: pb.FindEnabledNodeClusterFirewallActionResponse
	(*CountAllEnabledNodeClusterFirewallActionsRequest)(nil), // 8: pb.CountAllEnabledNodeClusterFirewallActionsRequest
	(*NodeClusterFirewallAction)(nil),                        // 9: pb.NodeClusterFirewallAction
	(*RPCSuccess)(nil),                                       // 10: pb.RPCSuccess
	(*RPCCountResponse)(nil),                                 // 11: pb.RPCCountResponse
}
var file_service_node_cluster_firewall_action_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllEnabledNodeClusterFirewallActionsResponse.nodeClusterFirewallActions:type_name -> pb.NodeClusterFirewallAction
	9,  // 1: pb.FindEnabledNodeClusterFirewallActionResponse.nodeClusterFirewallAction:type_name -> pb.NodeClusterFirewallAction
	0,  // 2: pb.NodeClusterFirewallActionService.createNodeClusterFirewallAction:input_type -> pb.CreateNodeClusterFirewallActionRequest
	2,  // 3: pb.NodeClusterFirewallActionService.updateNodeClusterFirewallAction:input_type -> pb.UpdateNodeClusterFirewallActionRequest
	3,  // 4: pb.NodeClusterFirewallActionService.deleteNodeClusterFirewallAction:input_type -> pb.DeleteNodeClusterFirewallActionRequest
	4,  // 5: pb.NodeClusterFirewallActionService.findAllEnabledNodeClusterFirewallActions:input_type -> pb.FindAllEnabledNodeClusterFirewallActionsRequest
	6,  // 6: pb.NodeClusterFirewallActionService.findEnabledNodeClusterFirewallAction:input_type -> pb.FindEnabledNodeClusterFirewallActionRequest
	8,  // 7: pb.NodeClusterFirewallActionService.countAllEnabledNodeClusterFirewallActions:input_type -> pb.CountAllEnabledNodeClusterFirewallActionsRequest
	1,  // 8: pb.NodeClusterFirewallActionService.createNodeClusterFirewallAction:output_type -> pb.NodeClusterFirewallActionResponse
	10, // 9: pb.NodeClusterFirewallActionService.updateNodeClusterFirewallAction:output_type -> pb.RPCSuccess
	10, // 10: pb.NodeClusterFirewallActionService.deleteNodeClusterFirewallAction:output_type -> pb.RPCSuccess
	5,  // 11: pb.NodeClusterFirewallActionService.findAllEnabledNodeClusterFirewallActions:output_type -> pb.FindAllEnabledNodeClusterFirewallActionsResponse
	7,  // 12: pb.NodeClusterFirewallActionService.findEnabledNodeClusterFirewallAction:output_type -> pb.FindEnabledNodeClusterFirewallActionResponse
	11, // 13: pb.NodeClusterFirewallActionService.countAllEnabledNodeClusterFirewallActions:output_type -> pb.RPCCountResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_node_cluster_firewall_action_proto_init() }
func file_service_node_cluster_firewall_action_proto_init() {
	if File_service_node_cluster_firewall_action_proto != nil {
		return
	}
	file_models_model_node_cluster_firewall_action_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_cluster_firewall_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeClusterFirewallActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeClusterFirewallActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeClusterFirewallActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeClusterFirewallActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeClusterFirewallActionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeClusterFirewallActionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledNodeClusterFirewallActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledNodeClusterFirewallActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_cluster_firewall_action_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledNodeClusterFirewallActionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_node_cluster_firewall_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_cluster_firewall_action_proto_goTypes,
		DependencyIndexes: file_service_node_cluster_firewall_action_proto_depIdxs,
		MessageInfos:      file_service_node_cluster_firewall_action_proto_msgTypes,
	}.Build()
	File_service_node_cluster_firewall_action_proto = out.File
	file_service_node_cluster_firewall_action_proto_rawDesc = nil
	file_service_node_cluster_firewall_action_proto_goTypes = nil
	file_service_node_cluster_firewall_action_proto_depIdxs = nil
}