// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_ad_package_period.proto

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

// 创建有效期
type CreateADPackagePeriodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Unit  string `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"` // month | year
}

func (x *CreateADPackagePeriodRequest) Reset() {
	*x = CreateADPackagePeriodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateADPackagePeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateADPackagePeriodRequest) ProtoMessage() {}

func (x *CreateADPackagePeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateADPackagePeriodRequest.ProtoReflect.Descriptor instead.
func (*CreateADPackagePeriodRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{0}
}

func (x *CreateADPackagePeriodRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateADPackagePeriodRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type CreateADPackagePeriodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriodId int64 `protobuf:"varint,1,opt,name=adPackagePeriodId,proto3" json:"adPackagePeriodId,omitempty"`
}

func (x *CreateADPackagePeriodResponse) Reset() {
	*x = CreateADPackagePeriodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateADPackagePeriodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateADPackagePeriodResponse) ProtoMessage() {}

func (x *CreateADPackagePeriodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateADPackagePeriodResponse.ProtoReflect.Descriptor instead.
func (*CreateADPackagePeriodResponse) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{1}
}

func (x *CreateADPackagePeriodResponse) GetAdPackagePeriodId() int64 {
	if x != nil {
		return x.AdPackagePeriodId
	}
	return 0
}

// 修改有效期
type UpdateADPackagePeriodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriodId int64 `protobuf:"varint,1,opt,name=adPackagePeriodId,proto3" json:"adPackagePeriodId,omitempty"`
	IsOn              bool  `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateADPackagePeriodRequest) Reset() {
	*x = UpdateADPackagePeriodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateADPackagePeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateADPackagePeriodRequest) ProtoMessage() {}

func (x *UpdateADPackagePeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateADPackagePeriodRequest.ProtoReflect.Descriptor instead.
func (*UpdateADPackagePeriodRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateADPackagePeriodRequest) GetAdPackagePeriodId() int64 {
	if x != nil {
		return x.AdPackagePeriodId
	}
	return 0
}

func (x *UpdateADPackagePeriodRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 删除有效期
type DeleteADPackagePeriodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriodId int64 `protobuf:"varint,1,opt,name=adPackagePeriodId,proto3" json:"adPackagePeriodId,omitempty"`
}

func (x *DeleteADPackagePeriodRequest) Reset() {
	*x = DeleteADPackagePeriodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteADPackagePeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteADPackagePeriodRequest) ProtoMessage() {}

func (x *DeleteADPackagePeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteADPackagePeriodRequest.ProtoReflect.Descriptor instead.
func (*DeleteADPackagePeriodRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteADPackagePeriodRequest) GetAdPackagePeriodId() int64 {
	if x != nil {
		return x.AdPackagePeriodId
	}
	return 0
}

// 查找有效期
type FindADPackagePeriodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriodId int64 `protobuf:"varint,1,opt,name=adPackagePeriodId,proto3" json:"adPackagePeriodId,omitempty"`
}

func (x *FindADPackagePeriodRequest) Reset() {
	*x = FindADPackagePeriodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindADPackagePeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindADPackagePeriodRequest) ProtoMessage() {}

func (x *FindADPackagePeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindADPackagePeriodRequest.ProtoReflect.Descriptor instead.
func (*FindADPackagePeriodRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{4}
}

func (x *FindADPackagePeriodRequest) GetAdPackagePeriodId() int64 {
	if x != nil {
		return x.AdPackagePeriodId
	}
	return 0
}

type FindADPackagePeriodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriod *ADPackagePeriod `protobuf:"bytes,1,opt,name=adPackagePeriod,proto3" json:"adPackagePeriod,omitempty"`
}

func (x *FindADPackagePeriodResponse) Reset() {
	*x = FindADPackagePeriodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindADPackagePeriodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindADPackagePeriodResponse) ProtoMessage() {}

func (x *FindADPackagePeriodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindADPackagePeriodResponse.ProtoReflect.Descriptor instead.
func (*FindADPackagePeriodResponse) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{5}
}

func (x *FindADPackagePeriodResponse) GetAdPackagePeriod() *ADPackagePeriod {
	if x != nil {
		return x.AdPackagePeriod
	}
	return nil
}

// 列出所有有效期
type FindAllADPackagePeriodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllADPackagePeriodsRequest) Reset() {
	*x = FindAllADPackagePeriodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllADPackagePeriodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllADPackagePeriodsRequest) ProtoMessage() {}

func (x *FindAllADPackagePeriodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllADPackagePeriodsRequest.ProtoReflect.Descriptor instead.
func (*FindAllADPackagePeriodsRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{6}
}

type FindAllADPackagePeriodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriods []*ADPackagePeriod `protobuf:"bytes,1,rep,name=adPackagePeriods,proto3" json:"adPackagePeriods,omitempty"`
}

func (x *FindAllADPackagePeriodsResponse) Reset() {
	*x = FindAllADPackagePeriodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllADPackagePeriodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllADPackagePeriodsResponse) ProtoMessage() {}

func (x *FindAllADPackagePeriodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllADPackagePeriodsResponse.ProtoReflect.Descriptor instead.
func (*FindAllADPackagePeriodsResponse) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{7}
}

func (x *FindAllADPackagePeriodsResponse) GetAdPackagePeriods() []*ADPackagePeriod {
	if x != nil {
		return x.AdPackagePeriods
	}
	return nil
}

// 列出所有可用有效期
type FindAllAvailableADPackagePeriodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllAvailableADPackagePeriodsRequest) Reset() {
	*x = FindAllAvailableADPackagePeriodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableADPackagePeriodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableADPackagePeriodsRequest) ProtoMessage() {}

func (x *FindAllAvailableADPackagePeriodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableADPackagePeriodsRequest.ProtoReflect.Descriptor instead.
func (*FindAllAvailableADPackagePeriodsRequest) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{8}
}

type FindAllAvailableADPackagePeriodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdPackagePeriods []*ADPackagePeriod `protobuf:"bytes,1,rep,name=adPackagePeriods,proto3" json:"adPackagePeriods,omitempty"`
}

func (x *FindAllAvailableADPackagePeriodsResponse) Reset() {
	*x = FindAllAvailableADPackagePeriodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ad_package_period_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableADPackagePeriodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableADPackagePeriodsResponse) ProtoMessage() {}

func (x *FindAllAvailableADPackagePeriodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ad_package_period_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableADPackagePeriodsResponse.ProtoReflect.Descriptor instead.
func (*FindAllAvailableADPackagePeriodsResponse) Descriptor() ([]byte, []int) {
	return file_service_ad_package_period_proto_rawDescGZIP(), []int{9}
}

func (x *FindAllAvailableADPackagePeriodsResponse) GetAdPackagePeriods() []*ADPackagePeriod {
	if x != nil {
		return x.AdPackagePeriods
	}
	return nil
}

var File_service_ad_package_period_proto protoreflect.FileDescriptor

var file_service_ad_package_period_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x22, 0x4d, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x61, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x22,
	0x60, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x61, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x22, 0x4c, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x61, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x22,
	0x4a, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x11, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x1b, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x61, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0f, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x1f, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x10, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x44,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x10, 0x61,
	0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x22,
	0x29, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x28, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x44,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x10, 0x61, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x32, 0xc7, 0x04, 0x0a, 0x16, 0x41, 0x44, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x15, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x44,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1e, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62,
	0x0a, 0x17, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x7d, 0x0a, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x44, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_ad_package_period_proto_rawDescOnce sync.Once
	file_service_ad_package_period_proto_rawDescData = file_service_ad_package_period_proto_rawDesc
)

func file_service_ad_package_period_proto_rawDescGZIP() []byte {
	file_service_ad_package_period_proto_rawDescOnce.Do(func() {
		file_service_ad_package_period_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ad_package_period_proto_rawDescData)
	})
	return file_service_ad_package_period_proto_rawDescData
}

var file_service_ad_package_period_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_ad_package_period_proto_goTypes = []interface{}{
	(*CreateADPackagePeriodRequest)(nil),             // 0: pb.CreateADPackagePeriodRequest
	(*CreateADPackagePeriodResponse)(nil),            // 1: pb.CreateADPackagePeriodResponse
	(*UpdateADPackagePeriodRequest)(nil),             // 2: pb.UpdateADPackagePeriodRequest
	(*DeleteADPackagePeriodRequest)(nil),             // 3: pb.DeleteADPackagePeriodRequest
	(*FindADPackagePeriodRequest)(nil),               // 4: pb.FindADPackagePeriodRequest
	(*FindADPackagePeriodResponse)(nil),              // 5: pb.FindADPackagePeriodResponse
	(*FindAllADPackagePeriodsRequest)(nil),           // 6: pb.FindAllADPackagePeriodsRequest
	(*FindAllADPackagePeriodsResponse)(nil),          // 7: pb.FindAllADPackagePeriodsResponse
	(*FindAllAvailableADPackagePeriodsRequest)(nil),  // 8: pb.FindAllAvailableADPackagePeriodsRequest
	(*FindAllAvailableADPackagePeriodsResponse)(nil), // 9: pb.FindAllAvailableADPackagePeriodsResponse
	(*ADPackagePeriod)(nil),                          // 10: pb.ADPackagePeriod
	(*RPCSuccess)(nil),                               // 11: pb.RPCSuccess
}
var file_service_ad_package_period_proto_depIdxs = []int32{
	10, // 0: pb.FindADPackagePeriodResponse.adPackagePeriod:type_name -> pb.ADPackagePeriod
	10, // 1: pb.FindAllADPackagePeriodsResponse.adPackagePeriods:type_name -> pb.ADPackagePeriod
	10, // 2: pb.FindAllAvailableADPackagePeriodsResponse.adPackagePeriods:type_name -> pb.ADPackagePeriod
	0,  // 3: pb.ADPackagePeriodService.createADPackagePeriod:input_type -> pb.CreateADPackagePeriodRequest
	2,  // 4: pb.ADPackagePeriodService.updateADPackagePeriod:input_type -> pb.UpdateADPackagePeriodRequest
	3,  // 5: pb.ADPackagePeriodService.deleteADPackagePeriod:input_type -> pb.DeleteADPackagePeriodRequest
	4,  // 6: pb.ADPackagePeriodService.findADPackagePeriod:input_type -> pb.FindADPackagePeriodRequest
	6,  // 7: pb.ADPackagePeriodService.findAllADPackagePeriods:input_type -> pb.FindAllADPackagePeriodsRequest
	8,  // 8: pb.ADPackagePeriodService.findAllAvailableADPackagePeriods:input_type -> pb.FindAllAvailableADPackagePeriodsRequest
	1,  // 9: pb.ADPackagePeriodService.createADPackagePeriod:output_type -> pb.CreateADPackagePeriodResponse
	11, // 10: pb.ADPackagePeriodService.updateADPackagePeriod:output_type -> pb.RPCSuccess
	11, // 11: pb.ADPackagePeriodService.deleteADPackagePeriod:output_type -> pb.RPCSuccess
	5,  // 12: pb.ADPackagePeriodService.findADPackagePeriod:output_type -> pb.FindADPackagePeriodResponse
	7,  // 13: pb.ADPackagePeriodService.findAllADPackagePeriods:output_type -> pb.FindAllADPackagePeriodsResponse
	9,  // 14: pb.ADPackagePeriodService.findAllAvailableADPackagePeriods:output_type -> pb.FindAllAvailableADPackagePeriodsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_ad_package_period_proto_init() }
func file_service_ad_package_period_proto_init() {
	if File_service_ad_package_period_proto != nil {
		return
	}
	file_models_model_ad_package_period_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ad_package_period_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateADPackagePeriodRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateADPackagePeriodResponse); i {
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
		file_service_ad_package_period_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateADPackagePeriodRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteADPackagePeriodRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindADPackagePeriodRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindADPackagePeriodResponse); i {
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
		file_service_ad_package_period_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllADPackagePeriodsRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllADPackagePeriodsResponse); i {
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
		file_service_ad_package_period_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableADPackagePeriodsRequest); i {
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
		file_service_ad_package_period_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableADPackagePeriodsResponse); i {
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
			RawDescriptor: file_service_ad_package_period_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ad_package_period_proto_goTypes,
		DependencyIndexes: file_service_ad_package_period_proto_depIdxs,
		MessageInfos:      file_service_ad_package_period_proto_msgTypes,
	}.Build()
	File_service_ad_package_period_proto = out.File
	file_service_ad_package_period_proto_rawDesc = nil
	file_service_ad_package_period_proto_goTypes = nil
	file_service_ad_package_period_proto_depIdxs = nil
}