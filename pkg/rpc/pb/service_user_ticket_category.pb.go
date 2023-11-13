// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_user_ticket_category.proto

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

// 创建分类
type CreateUserTicketCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateUserTicketCategoryRequest) Reset() {
	*x = CreateUserTicketCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTicketCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTicketCategoryRequest) ProtoMessage() {}

func (x *CreateUserTicketCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTicketCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateUserTicketCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserTicketCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateUserTicketCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategoryId int64 `protobuf:"varint,1,opt,name=userTicketCategoryId,proto3" json:"userTicketCategoryId,omitempty"`
}

func (x *CreateUserTicketCategoryResponse) Reset() {
	*x = CreateUserTicketCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTicketCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTicketCategoryResponse) ProtoMessage() {}

func (x *CreateUserTicketCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTicketCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateUserTicketCategoryResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserTicketCategoryResponse) GetUserTicketCategoryId() int64 {
	if x != nil {
		return x.UserTicketCategoryId
	}
	return 0
}

// 修改分类
type UpdateUserTicketCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategoryId int64  `protobuf:"varint,1,opt,name=userTicketCategoryId,proto3" json:"userTicketCategoryId,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn                 bool   `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateUserTicketCategoryRequest) Reset() {
	*x = UpdateUserTicketCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserTicketCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserTicketCategoryRequest) ProtoMessage() {}

func (x *UpdateUserTicketCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserTicketCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserTicketCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserTicketCategoryRequest) GetUserTicketCategoryId() int64 {
	if x != nil {
		return x.UserTicketCategoryId
	}
	return 0
}

func (x *UpdateUserTicketCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserTicketCategoryRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 删除分类
type DeleteUserTicketCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategoryId int64 `protobuf:"varint,1,opt,name=userTicketCategoryId,proto3" json:"userTicketCategoryId,omitempty"`
}

func (x *DeleteUserTicketCategoryRequest) Reset() {
	*x = DeleteUserTicketCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTicketCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTicketCategoryRequest) ProtoMessage() {}

func (x *DeleteUserTicketCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTicketCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTicketCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserTicketCategoryRequest) GetUserTicketCategoryId() int64 {
	if x != nil {
		return x.UserTicketCategoryId
	}
	return 0
}

// 查找所有分类
type FindAllUserTicketCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllUserTicketCategoriesRequest) Reset() {
	*x = FindAllUserTicketCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllUserTicketCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllUserTicketCategoriesRequest) ProtoMessage() {}

func (x *FindAllUserTicketCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllUserTicketCategoriesRequest.ProtoReflect.Descriptor instead.
func (*FindAllUserTicketCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{4}
}

type FindAllUserTicketCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategories []*UserTicketCategory `protobuf:"bytes,1,rep,name=userTicketCategories,proto3" json:"userTicketCategories,omitempty"`
}

func (x *FindAllUserTicketCategoriesResponse) Reset() {
	*x = FindAllUserTicketCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllUserTicketCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllUserTicketCategoriesResponse) ProtoMessage() {}

func (x *FindAllUserTicketCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllUserTicketCategoriesResponse.ProtoReflect.Descriptor instead.
func (*FindAllUserTicketCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllUserTicketCategoriesResponse) GetUserTicketCategories() []*UserTicketCategory {
	if x != nil {
		return x.UserTicketCategories
	}
	return nil
}

// 查找所有启用中的分类
type FindAllAvailableUserTicketCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllAvailableUserTicketCategoriesRequest) Reset() {
	*x = FindAllAvailableUserTicketCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableUserTicketCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableUserTicketCategoriesRequest) ProtoMessage() {}

func (x *FindAllAvailableUserTicketCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableUserTicketCategoriesRequest.ProtoReflect.Descriptor instead.
func (*FindAllAvailableUserTicketCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{6}
}

type FindAllAvailableUserTicketCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategories []*UserTicketCategory `protobuf:"bytes,1,rep,name=userTicketCategories,proto3" json:"userTicketCategories,omitempty"`
}

func (x *FindAllAvailableUserTicketCategoriesResponse) Reset() {
	*x = FindAllAvailableUserTicketCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableUserTicketCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableUserTicketCategoriesResponse) ProtoMessage() {}

func (x *FindAllAvailableUserTicketCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableUserTicketCategoriesResponse.ProtoReflect.Descriptor instead.
func (*FindAllAvailableUserTicketCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{7}
}

func (x *FindAllAvailableUserTicketCategoriesResponse) GetUserTicketCategories() []*UserTicketCategory {
	if x != nil {
		return x.UserTicketCategories
	}
	return nil
}

// 查询单个分类
type FindUserTicketCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategoryId int64 `protobuf:"varint,1,opt,name=userTicketCategoryId,proto3" json:"userTicketCategoryId,omitempty"`
}

func (x *FindUserTicketCategoryRequest) Reset() {
	*x = FindUserTicketCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserTicketCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserTicketCategoryRequest) ProtoMessage() {}

func (x *FindUserTicketCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserTicketCategoryRequest.ProtoReflect.Descriptor instead.
func (*FindUserTicketCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{8}
}

func (x *FindUserTicketCategoryRequest) GetUserTicketCategoryId() int64 {
	if x != nil {
		return x.UserTicketCategoryId
	}
	return 0
}

type FindUserTicketCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketCategory *UserTicketCategory `protobuf:"bytes,1,opt,name=userTicketCategory,proto3" json:"userTicketCategory,omitempty"`
}

func (x *FindUserTicketCategoryResponse) Reset() {
	*x = FindUserTicketCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_category_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserTicketCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserTicketCategoryResponse) ProtoMessage() {}

func (x *FindUserTicketCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_category_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserTicketCategoryResponse.ProtoReflect.Descriptor instead.
func (*FindUserTicketCategoryResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_category_proto_rawDescGZIP(), []int{9}
}

func (x *FindUserTicketCategoryResponse) GetUserTicketCategory() *UserTicketCategory {
	if x != nil {
		return x.UserTicketCategory
	}
	return nil
}

var File_service_user_ticket_category_proto protoreflect.FileDescriptor

var file_service_user_ticket_category_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x1f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x1f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x55, 0x0a, 0x1f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x22, 0x24, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x71, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x2b, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x2c, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x1e, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x12,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x32, 0x81, 0x05, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x65, 0x0a, 0x18, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x18, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4f, 0x0a, 0x18, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6e, 0x0a, 0x1b, 0x66,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x24,
	0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_ticket_category_proto_rawDescOnce sync.Once
	file_service_user_ticket_category_proto_rawDescData = file_service_user_ticket_category_proto_rawDesc
)

func file_service_user_ticket_category_proto_rawDescGZIP() []byte {
	file_service_user_ticket_category_proto_rawDescOnce.Do(func() {
		file_service_user_ticket_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_ticket_category_proto_rawDescData)
	})
	return file_service_user_ticket_category_proto_rawDescData
}

var file_service_user_ticket_category_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_user_ticket_category_proto_goTypes = []interface{}{
	(*CreateUserTicketCategoryRequest)(nil),              // 0: pb.CreateUserTicketCategoryRequest
	(*CreateUserTicketCategoryResponse)(nil),             // 1: pb.CreateUserTicketCategoryResponse
	(*UpdateUserTicketCategoryRequest)(nil),              // 2: pb.UpdateUserTicketCategoryRequest
	(*DeleteUserTicketCategoryRequest)(nil),              // 3: pb.DeleteUserTicketCategoryRequest
	(*FindAllUserTicketCategoriesRequest)(nil),           // 4: pb.FindAllUserTicketCategoriesRequest
	(*FindAllUserTicketCategoriesResponse)(nil),          // 5: pb.FindAllUserTicketCategoriesResponse
	(*FindAllAvailableUserTicketCategoriesRequest)(nil),  // 6: pb.FindAllAvailableUserTicketCategoriesRequest
	(*FindAllAvailableUserTicketCategoriesResponse)(nil), // 7: pb.FindAllAvailableUserTicketCategoriesResponse
	(*FindUserTicketCategoryRequest)(nil),                // 8: pb.FindUserTicketCategoryRequest
	(*FindUserTicketCategoryResponse)(nil),               // 9: pb.FindUserTicketCategoryResponse
	(*UserTicketCategory)(nil),                           // 10: pb.UserTicketCategory
	(*RPCSuccess)(nil),                                   // 11: pb.RPCSuccess
}
var file_service_user_ticket_category_proto_depIdxs = []int32{
	10, // 0: pb.FindAllUserTicketCategoriesResponse.userTicketCategories:type_name -> pb.UserTicketCategory
	10, // 1: pb.FindAllAvailableUserTicketCategoriesResponse.userTicketCategories:type_name -> pb.UserTicketCategory
	10, // 2: pb.FindUserTicketCategoryResponse.userTicketCategory:type_name -> pb.UserTicketCategory
	0,  // 3: pb.UserTicketCategoryService.createUserTicketCategory:input_type -> pb.CreateUserTicketCategoryRequest
	2,  // 4: pb.UserTicketCategoryService.updateUserTicketCategory:input_type -> pb.UpdateUserTicketCategoryRequest
	3,  // 5: pb.UserTicketCategoryService.deleteUserTicketCategory:input_type -> pb.DeleteUserTicketCategoryRequest
	4,  // 6: pb.UserTicketCategoryService.findAllUserTicketCategories:input_type -> pb.FindAllUserTicketCategoriesRequest
	6,  // 7: pb.UserTicketCategoryService.findAllAvailableUserTicketCategories:input_type -> pb.FindAllAvailableUserTicketCategoriesRequest
	8,  // 8: pb.UserTicketCategoryService.findUserTicketCategory:input_type -> pb.FindUserTicketCategoryRequest
	1,  // 9: pb.UserTicketCategoryService.createUserTicketCategory:output_type -> pb.CreateUserTicketCategoryResponse
	11, // 10: pb.UserTicketCategoryService.updateUserTicketCategory:output_type -> pb.RPCSuccess
	11, // 11: pb.UserTicketCategoryService.deleteUserTicketCategory:output_type -> pb.RPCSuccess
	5,  // 12: pb.UserTicketCategoryService.findAllUserTicketCategories:output_type -> pb.FindAllUserTicketCategoriesResponse
	7,  // 13: pb.UserTicketCategoryService.findAllAvailableUserTicketCategories:output_type -> pb.FindAllAvailableUserTicketCategoriesResponse
	9,  // 14: pb.UserTicketCategoryService.findUserTicketCategory:output_type -> pb.FindUserTicketCategoryResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_user_ticket_category_proto_init() }
func file_service_user_ticket_category_proto_init() {
	if File_service_user_ticket_category_proto != nil {
		return
	}
	file_models_model_user_ticket_category_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_ticket_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTicketCategoryRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTicketCategoryResponse); i {
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
		file_service_user_ticket_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserTicketCategoryRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTicketCategoryRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllUserTicketCategoriesRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllUserTicketCategoriesResponse); i {
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
		file_service_user_ticket_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableUserTicketCategoriesRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableUserTicketCategoriesResponse); i {
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
		file_service_user_ticket_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserTicketCategoryRequest); i {
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
		file_service_user_ticket_category_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserTicketCategoryResponse); i {
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
			RawDescriptor: file_service_user_ticket_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_ticket_category_proto_goTypes,
		DependencyIndexes: file_service_user_ticket_category_proto_depIdxs,
		MessageInfos:      file_service_user_ticket_category_proto_msgTypes,
	}.Build()
	File_service_user_ticket_category_proto = out.File
	file_service_user_ticket_category_proto_rawDesc = nil
	file_service_user_ticket_category_proto_goTypes = nil
	file_service_user_ticket_category_proto_depIdxs = nil
}
