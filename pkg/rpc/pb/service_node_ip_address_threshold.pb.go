// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_node_ip_address_threshold.proto

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

// 创建阈值
type CreateNodeIPAddressThresholdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId int64  `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
	ItemsJSON       []byte `protobuf:"bytes,2,opt,name=itemsJSON,proto3" json:"itemsJSON,omitempty"`
	ActionsJSON     []byte `protobuf:"bytes,3,opt,name=actionsJSON,proto3" json:"actionsJSON,omitempty"`
}

func (x *CreateNodeIPAddressThresholdRequest) Reset() {
	*x = CreateNodeIPAddressThresholdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeIPAddressThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeIPAddressThresholdRequest) ProtoMessage() {}

func (x *CreateNodeIPAddressThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeIPAddressThresholdRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeIPAddressThresholdRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeIPAddressThresholdRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

func (x *CreateNodeIPAddressThresholdRequest) GetItemsJSON() []byte {
	if x != nil {
		return x.ItemsJSON
	}
	return nil
}

func (x *CreateNodeIPAddressThresholdRequest) GetActionsJSON() []byte {
	if x != nil {
		return x.ActionsJSON
	}
	return nil
}

type CreateNodeIPAddressThresholdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressThresholdId int64 `protobuf:"varint,1,opt,name=nodeIPAddressThresholdId,proto3" json:"nodeIPAddressThresholdId,omitempty"`
}

func (x *CreateNodeIPAddressThresholdResponse) Reset() {
	*x = CreateNodeIPAddressThresholdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeIPAddressThresholdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeIPAddressThresholdResponse) ProtoMessage() {}

func (x *CreateNodeIPAddressThresholdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeIPAddressThresholdResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeIPAddressThresholdResponse) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNodeIPAddressThresholdResponse) GetNodeIPAddressThresholdId() int64 {
	if x != nil {
		return x.NodeIPAddressThresholdId
	}
	return 0
}

// 修改阈值
type UpdateNodeIPAddressThresholdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressThresholdId int64  `protobuf:"varint,1,opt,name=nodeIPAddressThresholdId,proto3" json:"nodeIPAddressThresholdId,omitempty"`
	ItemsJSON                []byte `protobuf:"bytes,2,opt,name=itemsJSON,proto3" json:"itemsJSON,omitempty"`
	ActionsJSON              []byte `protobuf:"bytes,3,opt,name=actionsJSON,proto3" json:"actionsJSON,omitempty"`
}

func (x *UpdateNodeIPAddressThresholdRequest) Reset() {
	*x = UpdateNodeIPAddressThresholdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeIPAddressThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeIPAddressThresholdRequest) ProtoMessage() {}

func (x *UpdateNodeIPAddressThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeIPAddressThresholdRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeIPAddressThresholdRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNodeIPAddressThresholdRequest) GetNodeIPAddressThresholdId() int64 {
	if x != nil {
		return x.NodeIPAddressThresholdId
	}
	return 0
}

func (x *UpdateNodeIPAddressThresholdRequest) GetItemsJSON() []byte {
	if x != nil {
		return x.ItemsJSON
	}
	return nil
}

func (x *UpdateNodeIPAddressThresholdRequest) GetActionsJSON() []byte {
	if x != nil {
		return x.ActionsJSON
	}
	return nil
}

// 删除阈值
type DeleteNodeIPAddressThresholdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressThresholdId int64 `protobuf:"varint,1,opt,name=nodeIPAddressThresholdId,proto3" json:"nodeIPAddressThresholdId,omitempty"`
}

func (x *DeleteNodeIPAddressThresholdRequest) Reset() {
	*x = DeleteNodeIPAddressThresholdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeIPAddressThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeIPAddressThresholdRequest) ProtoMessage() {}

func (x *DeleteNodeIPAddressThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeIPAddressThresholdRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeIPAddressThresholdRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNodeIPAddressThresholdRequest) GetNodeIPAddressThresholdId() int64 {
	if x != nil {
		return x.NodeIPAddressThresholdId
	}
	return 0
}

// 查找IP的所有阈值
type FindAllEnabledNodeIPAddressThresholdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId int64 `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
}

func (x *FindAllEnabledNodeIPAddressThresholdsRequest) Reset() {
	*x = FindAllEnabledNodeIPAddressThresholdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeIPAddressThresholdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeIPAddressThresholdsRequest) ProtoMessage() {}

func (x *FindAllEnabledNodeIPAddressThresholdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeIPAddressThresholdsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeIPAddressThresholdsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllEnabledNodeIPAddressThresholdsRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

type FindAllEnabledNodeIPAddressThresholdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressThresholds []*NodeIPAddressThreshold `protobuf:"bytes,1,rep,name=nodeIPAddressThresholds,proto3" json:"nodeIPAddressThresholds,omitempty"`
}

func (x *FindAllEnabledNodeIPAddressThresholdsResponse) Reset() {
	*x = FindAllEnabledNodeIPAddressThresholdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledNodeIPAddressThresholdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledNodeIPAddressThresholdsResponse) ProtoMessage() {}

func (x *FindAllEnabledNodeIPAddressThresholdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledNodeIPAddressThresholdsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledNodeIPAddressThresholdsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllEnabledNodeIPAddressThresholdsResponse) GetNodeIPAddressThresholds() []*NodeIPAddressThreshold {
	if x != nil {
		return x.NodeIPAddressThresholds
	}
	return nil
}

// 计算IP阈值的数量
type CountAllEnabledNodeIPAddressThresholdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId int64 `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
}

func (x *CountAllEnabledNodeIPAddressThresholdsRequest) Reset() {
	*x = CountAllEnabledNodeIPAddressThresholdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledNodeIPAddressThresholdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledNodeIPAddressThresholdsRequest) ProtoMessage() {}

func (x *CountAllEnabledNodeIPAddressThresholdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledNodeIPAddressThresholdsRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledNodeIPAddressThresholdsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{6}
}

func (x *CountAllEnabledNodeIPAddressThresholdsRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

// 批量更新阈值
type UpdateAllNodeIPAddressThresholdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId             int64  `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
	NodeIPAddressThresholdsJSON []byte `protobuf:"bytes,2,opt,name=nodeIPAddressThresholdsJSON,proto3" json:"nodeIPAddressThresholdsJSON,omitempty"`
}

func (x *UpdateAllNodeIPAddressThresholdsRequest) Reset() {
	*x = UpdateAllNodeIPAddressThresholdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_threshold_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAllNodeIPAddressThresholdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAllNodeIPAddressThresholdsRequest) ProtoMessage() {}

func (x *UpdateAllNodeIPAddressThresholdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_threshold_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAllNodeIPAddressThresholdsRequest.ProtoReflect.Descriptor instead.
func (*UpdateAllNodeIPAddressThresholdsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_threshold_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAllNodeIPAddressThresholdsRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

func (x *UpdateAllNodeIPAddressThresholdsRequest) GetNodeIPAddressThresholdsJSON() []byte {
	if x != nil {
		return x.NodeIPAddressThresholdsJSON
	}
	return nil
}

var File_service_node_ip_address_threshold_proto protoreflect.FileDescriptor

var file_service_node_ip_address_threshold_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x2c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x62, 0x0a, 0x24, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a,
	0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x22, 0x61, 0x0a, 0x23, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x2c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x85, 0x01,
	0x0a, 0x2d, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x17, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x17, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x73, 0x22, 0x59, 0x0a, 0x2d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x22, 0x95, 0x01, 0x0a, 0x27, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1b, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0xa7, 0x05, 0x0a, 0x1d, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x1c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x27, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x57, 0x0a, 0x1c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x8c, 0x01, 0x0a, 0x25, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x30, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71,
	0x0a, 0x26, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5f, 0x0a, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_node_ip_address_threshold_proto_rawDescOnce sync.Once
	file_service_node_ip_address_threshold_proto_rawDescData = file_service_node_ip_address_threshold_proto_rawDesc
)

func file_service_node_ip_address_threshold_proto_rawDescGZIP() []byte {
	file_service_node_ip_address_threshold_proto_rawDescOnce.Do(func() {
		file_service_node_ip_address_threshold_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_ip_address_threshold_proto_rawDescData)
	})
	return file_service_node_ip_address_threshold_proto_rawDescData
}

var file_service_node_ip_address_threshold_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_node_ip_address_threshold_proto_goTypes = []interface{}{
	(*CreateNodeIPAddressThresholdRequest)(nil),           // 0: pb.CreateNodeIPAddressThresholdRequest
	(*CreateNodeIPAddressThresholdResponse)(nil),          // 1: pb.CreateNodeIPAddressThresholdResponse
	(*UpdateNodeIPAddressThresholdRequest)(nil),           // 2: pb.UpdateNodeIPAddressThresholdRequest
	(*DeleteNodeIPAddressThresholdRequest)(nil),           // 3: pb.DeleteNodeIPAddressThresholdRequest
	(*FindAllEnabledNodeIPAddressThresholdsRequest)(nil),  // 4: pb.FindAllEnabledNodeIPAddressThresholdsRequest
	(*FindAllEnabledNodeIPAddressThresholdsResponse)(nil), // 5: pb.FindAllEnabledNodeIPAddressThresholdsResponse
	(*CountAllEnabledNodeIPAddressThresholdsRequest)(nil), // 6: pb.CountAllEnabledNodeIPAddressThresholdsRequest
	(*UpdateAllNodeIPAddressThresholdsRequest)(nil),       // 7: pb.UpdateAllNodeIPAddressThresholdsRequest
	(*NodeIPAddressThreshold)(nil),                        // 8: pb.NodeIPAddressThreshold
	(*RPCSuccess)(nil),                                    // 9: pb.RPCSuccess
	(*RPCCountResponse)(nil),                              // 10: pb.RPCCountResponse
}
var file_service_node_ip_address_threshold_proto_depIdxs = []int32{
	8,  // 0: pb.FindAllEnabledNodeIPAddressThresholdsResponse.nodeIPAddressThresholds:type_name -> pb.NodeIPAddressThreshold
	0,  // 1: pb.NodeIPAddressThresholdService.createNodeIPAddressThreshold:input_type -> pb.CreateNodeIPAddressThresholdRequest
	2,  // 2: pb.NodeIPAddressThresholdService.updateNodeIPAddressThreshold:input_type -> pb.UpdateNodeIPAddressThresholdRequest
	3,  // 3: pb.NodeIPAddressThresholdService.deleteNodeIPAddressThreshold:input_type -> pb.DeleteNodeIPAddressThresholdRequest
	4,  // 4: pb.NodeIPAddressThresholdService.findAllEnabledNodeIPAddressThresholds:input_type -> pb.FindAllEnabledNodeIPAddressThresholdsRequest
	6,  // 5: pb.NodeIPAddressThresholdService.countAllEnabledNodeIPAddressThresholds:input_type -> pb.CountAllEnabledNodeIPAddressThresholdsRequest
	7,  // 6: pb.NodeIPAddressThresholdService.updateAllNodeIPAddressThresholds:input_type -> pb.UpdateAllNodeIPAddressThresholdsRequest
	1,  // 7: pb.NodeIPAddressThresholdService.createNodeIPAddressThreshold:output_type -> pb.CreateNodeIPAddressThresholdResponse
	9,  // 8: pb.NodeIPAddressThresholdService.updateNodeIPAddressThreshold:output_type -> pb.RPCSuccess
	9,  // 9: pb.NodeIPAddressThresholdService.deleteNodeIPAddressThreshold:output_type -> pb.RPCSuccess
	5,  // 10: pb.NodeIPAddressThresholdService.findAllEnabledNodeIPAddressThresholds:output_type -> pb.FindAllEnabledNodeIPAddressThresholdsResponse
	10, // 11: pb.NodeIPAddressThresholdService.countAllEnabledNodeIPAddressThresholds:output_type -> pb.RPCCountResponse
	9,  // 12: pb.NodeIPAddressThresholdService.updateAllNodeIPAddressThresholds:output_type -> pb.RPCSuccess
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_node_ip_address_threshold_proto_init() }
func file_service_node_ip_address_threshold_proto_init() {
	if File_service_node_ip_address_threshold_proto != nil {
		return
	}
	file_models_model_node_ip_address_threshold_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_ip_address_threshold_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeIPAddressThresholdRequest); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeIPAddressThresholdResponse); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeIPAddressThresholdRequest); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeIPAddressThresholdRequest); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeIPAddressThresholdsRequest); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledNodeIPAddressThresholdsResponse); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledNodeIPAddressThresholdsRequest); i {
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
		file_service_node_ip_address_threshold_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAllNodeIPAddressThresholdsRequest); i {
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
			RawDescriptor: file_service_node_ip_address_threshold_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_ip_address_threshold_proto_goTypes,
		DependencyIndexes: file_service_node_ip_address_threshold_proto_depIdxs,
		MessageInfos:      file_service_node_ip_address_threshold_proto_msgTypes,
	}.Build()
	File_service_node_ip_address_threshold_proto = out.File
	file_service_node_ip_address_threshold_proto_rawDesc = nil
	file_service_node_ip_address_threshold_proto_goTypes = nil
	file_service_node_ip_address_threshold_proto_depIdxs = nil
}