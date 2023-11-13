// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_ns_user_plan.proto

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

// 创建用户套餐
type CreateNSUserPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NsPlanId   int64  `protobuf:"varint,2,opt,name=nsPlanId,proto3" json:"nsPlanId,omitempty"`
	DayFrom    string `protobuf:"bytes,3,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`       // YYYYMMDD
	DayTo      string `protobuf:"bytes,4,opt,name=dayTo,proto3" json:"dayTo,omitempty"`           // YYYYMMDD
	PeriodUnit string `protobuf:"bytes,5,opt,name=periodUnit,proto3" json:"periodUnit,omitempty"` // yearly|monthly
}

func (x *CreateNSUserPlanRequest) Reset() {
	*x = CreateNSUserPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSUserPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSUserPlanRequest) ProtoMessage() {}

func (x *CreateNSUserPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSUserPlanRequest.ProtoReflect.Descriptor instead.
func (*CreateNSUserPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNSUserPlanRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateNSUserPlanRequest) GetNsPlanId() int64 {
	if x != nil {
		return x.NsPlanId
	}
	return 0
}

func (x *CreateNSUserPlanRequest) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *CreateNSUserPlanRequest) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *CreateNSUserPlanRequest) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

type CreateNSUserPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsUserPlanId int64 `protobuf:"varint,1,opt,name=nsUserPlanId,proto3" json:"nsUserPlanId,omitempty"`
}

func (x *CreateNSUserPlanResponse) Reset() {
	*x = CreateNSUserPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSUserPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSUserPlanResponse) ProtoMessage() {}

func (x *CreateNSUserPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSUserPlanResponse.ProtoReflect.Descriptor instead.
func (*CreateNSUserPlanResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNSUserPlanResponse) GetNsUserPlanId() int64 {
	if x != nil {
		return x.NsUserPlanId
	}
	return 0
}

// 修改用户套餐
type UpdateNSUserPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsUserPlanId int64  `protobuf:"varint,1,opt,name=nsUserPlanId,proto3" json:"nsUserPlanId,omitempty"`
	NsPlanId     int64  `protobuf:"varint,2,opt,name=nsPlanId,proto3" json:"nsPlanId,omitempty"`
	DayFrom      string `protobuf:"bytes,3,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`       // YYYYMMDD
	DayTo        string `protobuf:"bytes,4,opt,name=dayTo,proto3" json:"dayTo,omitempty"`           // YYYYMMDD
	PeriodUnit   string `protobuf:"bytes,5,opt,name=periodUnit,proto3" json:"periodUnit,omitempty"` // yearly|monthly
}

func (x *UpdateNSUserPlanRequest) Reset() {
	*x = UpdateNSUserPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNSUserPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNSUserPlanRequest) ProtoMessage() {}

func (x *UpdateNSUserPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNSUserPlanRequest.ProtoReflect.Descriptor instead.
func (*UpdateNSUserPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNSUserPlanRequest) GetNsUserPlanId() int64 {
	if x != nil {
		return x.NsUserPlanId
	}
	return 0
}

func (x *UpdateNSUserPlanRequest) GetNsPlanId() int64 {
	if x != nil {
		return x.NsPlanId
	}
	return 0
}

func (x *UpdateNSUserPlanRequest) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *UpdateNSUserPlanRequest) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *UpdateNSUserPlanRequest) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

// 删除用户套餐
type DeleteNSUserPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsUserPlanId int64 `protobuf:"varint,1,opt,name=nsUserPlanId,proto3" json:"nsUserPlanId,omitempty"`
}

func (x *DeleteNSUserPlanRequest) Reset() {
	*x = DeleteNSUserPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNSUserPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNSUserPlanRequest) ProtoMessage() {}

func (x *DeleteNSUserPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNSUserPlanRequest.ProtoReflect.Descriptor instead.
func (*DeleteNSUserPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNSUserPlanRequest) GetNsUserPlanId() int64 {
	if x != nil {
		return x.NsUserPlanId
	}
	return 0
}

// 读取用户套餐
type FindNSUserPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"` // 和 nsUserPlanId 二选一
	NsUserPlanId int64 `protobuf:"varint,2,opt,name=nsUserPlanId,proto3" json:"nsUserPlanId,omitempty"`
}

func (x *FindNSUserPlanRequest) Reset() {
	*x = FindNSUserPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSUserPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSUserPlanRequest) ProtoMessage() {}

func (x *FindNSUserPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSUserPlanRequest.ProtoReflect.Descriptor instead.
func (*FindNSUserPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{4}
}

func (x *FindNSUserPlanRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindNSUserPlanRequest) GetNsUserPlanId() int64 {
	if x != nil {
		return x.NsUserPlanId
	}
	return 0
}

type FindNSUserPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsUserPlan *NSUserPlan `protobuf:"bytes,1,opt,name=nsUserPlan,proto3" json:"nsUserPlan,omitempty"`
}

func (x *FindNSUserPlanResponse) Reset() {
	*x = FindNSUserPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSUserPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSUserPlanResponse) ProtoMessage() {}

func (x *FindNSUserPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSUserPlanResponse.ProtoReflect.Descriptor instead.
func (*FindNSUserPlanResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{5}
}

func (x *FindNSUserPlanResponse) GetNsUserPlan() *NSUserPlan {
	if x != nil {
		return x.NsUserPlan
	}
	return nil
}

// 计算用户套餐数量
type CountNSUserPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NsPlanId   int64  `protobuf:"varint,2,opt,name=nsPlanId,proto3" json:"nsPlanId,omitempty"`
	PeriodUnit string `protobuf:"bytes,3,opt,name=periodUnit,proto3" json:"periodUnit,omitempty"`
	IsExpired  bool   `protobuf:"varint,4,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	ExpireDays int32  `protobuf:"varint,5,opt,name=expireDays,proto3" json:"expireDays,omitempty"`
}

func (x *CountNSUserPlansRequest) Reset() {
	*x = CountNSUserPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountNSUserPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountNSUserPlansRequest) ProtoMessage() {}

func (x *CountNSUserPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountNSUserPlansRequest.ProtoReflect.Descriptor instead.
func (*CountNSUserPlansRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{6}
}

func (x *CountNSUserPlansRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountNSUserPlansRequest) GetNsPlanId() int64 {
	if x != nil {
		return x.NsPlanId
	}
	return 0
}

func (x *CountNSUserPlansRequest) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

func (x *CountNSUserPlansRequest) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

func (x *CountNSUserPlansRequest) GetExpireDays() int32 {
	if x != nil {
		return x.ExpireDays
	}
	return 0
}

// 列出单页套餐
type ListNSUserPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NsPlanId   int64  `protobuf:"varint,2,opt,name=nsPlanId,proto3" json:"nsPlanId,omitempty"`
	PeriodUnit string `protobuf:"bytes,3,opt,name=periodUnit,proto3" json:"periodUnit,omitempty"`
	IsExpired  bool   `protobuf:"varint,4,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	ExpireDays int32  `protobuf:"varint,5,opt,name=expireDays,proto3" json:"expireDays,omitempty"`
	Offset     int64  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Size       int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListNSUserPlansRequest) Reset() {
	*x = ListNSUserPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNSUserPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNSUserPlansRequest) ProtoMessage() {}

func (x *ListNSUserPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNSUserPlansRequest.ProtoReflect.Descriptor instead.
func (*ListNSUserPlansRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{7}
}

func (x *ListNSUserPlansRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListNSUserPlansRequest) GetNsPlanId() int64 {
	if x != nil {
		return x.NsPlanId
	}
	return 0
}

func (x *ListNSUserPlansRequest) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

func (x *ListNSUserPlansRequest) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

func (x *ListNSUserPlansRequest) GetExpireDays() int32 {
	if x != nil {
		return x.ExpireDays
	}
	return 0
}

func (x *ListNSUserPlansRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListNSUserPlansRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListNSUserPlansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsUserPlans []*NSUserPlan `protobuf:"bytes,1,rep,name=nsUserPlans,proto3" json:"nsUserPlans,omitempty"`
}

func (x *ListNSUserPlansResponse) Reset() {
	*x = ListNSUserPlansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNSUserPlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNSUserPlansResponse) ProtoMessage() {}

func (x *ListNSUserPlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNSUserPlansResponse.ProtoReflect.Descriptor instead.
func (*ListNSUserPlansResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{8}
}

func (x *ListNSUserPlansResponse) GetNsUserPlans() []*NSUserPlan {
	if x != nil {
		return x.NsUserPlans
	}
	return nil
}

// 使用余额购买用户套餐
type BuyNSUserPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PlanId int64  `protobuf:"varint,2,opt,name=planId,proto3" json:"planId,omitempty"`
	Period string `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *BuyNSUserPlanRequest) Reset() {
	*x = BuyNSUserPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyNSUserPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyNSUserPlanRequest) ProtoMessage() {}

func (x *BuyNSUserPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyNSUserPlanRequest.ProtoReflect.Descriptor instead.
func (*BuyNSUserPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{9}
}

func (x *BuyNSUserPlanRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BuyNSUserPlanRequest) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *BuyNSUserPlanRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

type BuyNSUserPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPlanId int64 `protobuf:"varint,1,opt,name=userPlanId,proto3" json:"userPlanId,omitempty"`
}

func (x *BuyNSUserPlanResponse) Reset() {
	*x = BuyNSUserPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_user_plan_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyNSUserPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyNSUserPlanResponse) ProtoMessage() {}

func (x *BuyNSUserPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_user_plan_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyNSUserPlanResponse.ProtoReflect.Descriptor instead.
func (*BuyNSUserPlanResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_user_plan_proto_rawDescGZIP(), []int{10}
}

func (x *BuyNSUserPlanResponse) GetUserPlanId() int64 {
	if x != nil {
		return x.UserPlanId
	}
	return 0
}

var File_service_ns_user_plan_proto protoreflect.FileDescriptor

var file_service_ns_user_plan_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6e, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x3e, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6e, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x3d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x4e,
	0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x16,
	0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x0a, 0x6e, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x73,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x73,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61,
	0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x61, 0x79, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x73, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x73, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x79, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6e, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x0b, 0x6e,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x22, 0x5e, 0x0a, 0x14, 0x42, 0x75,
	0x79, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x37, 0x0a, 0x15, 0x42, 0x75,
	0x79, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x32, 0x86, 0x04, 0x0a, 0x11, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c,
	0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x66, 0x69,
	0x6e, 0x64, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x53, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x6c, 0x69,
	0x73, 0x74, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x4e, 0x53, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x79,
	0x4e, 0x53, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x79, 0x4e, 0x53, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ns_user_plan_proto_rawDescOnce sync.Once
	file_service_ns_user_plan_proto_rawDescData = file_service_ns_user_plan_proto_rawDesc
)

func file_service_ns_user_plan_proto_rawDescGZIP() []byte {
	file_service_ns_user_plan_proto_rawDescOnce.Do(func() {
		file_service_ns_user_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ns_user_plan_proto_rawDescData)
	})
	return file_service_ns_user_plan_proto_rawDescData
}

var file_service_ns_user_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_ns_user_plan_proto_goTypes = []interface{}{
	(*CreateNSUserPlanRequest)(nil),  // 0: pb.CreateNSUserPlanRequest
	(*CreateNSUserPlanResponse)(nil), // 1: pb.CreateNSUserPlanResponse
	(*UpdateNSUserPlanRequest)(nil),  // 2: pb.UpdateNSUserPlanRequest
	(*DeleteNSUserPlanRequest)(nil),  // 3: pb.DeleteNSUserPlanRequest
	(*FindNSUserPlanRequest)(nil),    // 4: pb.FindNSUserPlanRequest
	(*FindNSUserPlanResponse)(nil),   // 5: pb.FindNSUserPlanResponse
	(*CountNSUserPlansRequest)(nil),  // 6: pb.CountNSUserPlansRequest
	(*ListNSUserPlansRequest)(nil),   // 7: pb.ListNSUserPlansRequest
	(*ListNSUserPlansResponse)(nil),  // 8: pb.ListNSUserPlansResponse
	(*BuyNSUserPlanRequest)(nil),     // 9: pb.BuyNSUserPlanRequest
	(*BuyNSUserPlanResponse)(nil),    // 10: pb.BuyNSUserPlanResponse
	(*NSUserPlan)(nil),               // 11: pb.NSUserPlan
	(*RPCSuccess)(nil),               // 12: pb.RPCSuccess
	(*RPCCountResponse)(nil),         // 13: pb.RPCCountResponse
}
var file_service_ns_user_plan_proto_depIdxs = []int32{
	11, // 0: pb.FindNSUserPlanResponse.nsUserPlan:type_name -> pb.NSUserPlan
	11, // 1: pb.ListNSUserPlansResponse.nsUserPlans:type_name -> pb.NSUserPlan
	0,  // 2: pb.NSUserPlanService.createNSUserPlan:input_type -> pb.CreateNSUserPlanRequest
	2,  // 3: pb.NSUserPlanService.updateNSUserPlan:input_type -> pb.UpdateNSUserPlanRequest
	3,  // 4: pb.NSUserPlanService.deleteNSUserPlan:input_type -> pb.DeleteNSUserPlanRequest
	4,  // 5: pb.NSUserPlanService.findNSUserPlan:input_type -> pb.FindNSUserPlanRequest
	6,  // 6: pb.NSUserPlanService.countNSUserPlans:input_type -> pb.CountNSUserPlansRequest
	7,  // 7: pb.NSUserPlanService.listNSUserPlans:input_type -> pb.ListNSUserPlansRequest
	9,  // 8: pb.NSUserPlanService.buyNSUserPlan:input_type -> pb.BuyNSUserPlanRequest
	1,  // 9: pb.NSUserPlanService.createNSUserPlan:output_type -> pb.CreateNSUserPlanResponse
	12, // 10: pb.NSUserPlanService.updateNSUserPlan:output_type -> pb.RPCSuccess
	12, // 11: pb.NSUserPlanService.deleteNSUserPlan:output_type -> pb.RPCSuccess
	5,  // 12: pb.NSUserPlanService.findNSUserPlan:output_type -> pb.FindNSUserPlanResponse
	13, // 13: pb.NSUserPlanService.countNSUserPlans:output_type -> pb.RPCCountResponse
	8,  // 14: pb.NSUserPlanService.listNSUserPlans:output_type -> pb.ListNSUserPlansResponse
	10, // 15: pb.NSUserPlanService.buyNSUserPlan:output_type -> pb.BuyNSUserPlanResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_ns_user_plan_proto_init() }
func file_service_ns_user_plan_proto_init() {
	if File_service_ns_user_plan_proto != nil {
		return
	}
	file_models_model_ns_user_plan_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ns_user_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSUserPlanRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSUserPlanResponse); i {
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
		file_service_ns_user_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNSUserPlanRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNSUserPlanRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSUserPlanRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSUserPlanResponse); i {
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
		file_service_ns_user_plan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountNSUserPlansRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNSUserPlansRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNSUserPlansResponse); i {
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
		file_service_ns_user_plan_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyNSUserPlanRequest); i {
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
		file_service_ns_user_plan_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyNSUserPlanResponse); i {
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
			RawDescriptor: file_service_ns_user_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ns_user_plan_proto_goTypes,
		DependencyIndexes: file_service_ns_user_plan_proto_depIdxs,
		MessageInfos:      file_service_ns_user_plan_proto_msgTypes,
	}.Build()
	File_service_ns_user_plan_proto = out.File
	file_service_ns_user_plan_proto_rawDesc = nil
	file_service_ns_user_plan_proto_goTypes = nil
	file_service_ns_user_plan_proto_depIdxs = nil
}
