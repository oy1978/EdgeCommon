// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_user_bill.proto

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

// 手工生成账单
type GenerateAllUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	Day   string `protobuf:"bytes,2,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *GenerateAllUserBillsRequest) Reset() {
	*x = GenerateAllUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAllUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAllUserBillsRequest) ProtoMessage() {}

func (x *GenerateAllUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAllUserBillsRequest.ProtoReflect.Descriptor instead.
func (*GenerateAllUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateAllUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *GenerateAllUserBillsRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

// 计算所有账单数量
type CountAllUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaidFlag           int32  `protobuf:"varint,1,opt,name=paidFlag,proto3" json:"paidFlag,omitempty"`                     // 可选，0|1|-1
	UserId             int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`                         // 可选，用户ID
	Month              string `protobuf:"bytes,3,opt,name=month,proto3" json:"month,omitempty"`                            // 可选，月份
	TrafficRelated     bool   `protobuf:"varint,4,opt,name=trafficRelated,proto3" json:"trafficRelated,omitempty"`         // 可选，是否为流量带宽相关
	MinDailyBillDays   int32  `protobuf:"varint,5,opt,name=minDailyBillDays,proto3" json:"minDailyBillDays,omitempty"`     // 可选，按日计费账单生成最小天数
	MinMonthlyBillDays int32  `protobuf:"varint,6,opt,name=minMonthlyBillDays,proto3" json:"minMonthlyBillDays,omitempty"` // 可选，按月计费账单生成最小天数
}

func (x *CountAllUserBillsRequest) Reset() {
	*x = CountAllUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllUserBillsRequest) ProtoMessage() {}

func (x *CountAllUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllUserBillsRequest.ProtoReflect.Descriptor instead.
func (*CountAllUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{1}
}

func (x *CountAllUserBillsRequest) GetPaidFlag() int32 {
	if x != nil {
		return x.PaidFlag
	}
	return 0
}

func (x *CountAllUserBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountAllUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *CountAllUserBillsRequest) GetTrafficRelated() bool {
	if x != nil {
		return x.TrafficRelated
	}
	return false
}

func (x *CountAllUserBillsRequest) GetMinDailyBillDays() int32 {
	if x != nil {
		return x.MinDailyBillDays
	}
	return 0
}

func (x *CountAllUserBillsRequest) GetMinMonthlyBillDays() int32 {
	if x != nil {
		return x.MinMonthlyBillDays
	}
	return 0
}

// 列出单页账单
type ListUserBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaidFlag int32  `protobuf:"varint,1,opt,name=paidFlag,proto3" json:"paidFlag,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Month    string `protobuf:"bytes,5,opt,name=month,proto3" json:"month,omitempty"`
	Offset   int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUserBillsRequest) Reset() {
	*x = ListUserBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBillsRequest) ProtoMessage() {}

func (x *ListUserBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBillsRequest.ProtoReflect.Descriptor instead.
func (*ListUserBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserBillsRequest) GetPaidFlag() int32 {
	if x != nil {
		return x.PaidFlag
	}
	return 0
}

func (x *ListUserBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListUserBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *ListUserBillsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserBillsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUserBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBills []*UserBill `protobuf:"bytes,1,rep,name=userBills,proto3" json:"userBills,omitempty"`
}

func (x *ListUserBillsResponse) Reset() {
	*x = ListUserBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserBillsResponse) ProtoMessage() {}

func (x *ListUserBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserBillsResponse.ProtoReflect.Descriptor instead.
func (*ListUserBillsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserBillsResponse) GetUserBills() []*UserBill {
	if x != nil {
		return x.UserBills
	}
	return nil
}

// 查找账单信息
type FindUserBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBillId int64  `protobuf:"varint,1,opt,name=userBillId,proto3" json:"userBillId,omitempty"` // ID，和单号二选一
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`              // 单号
}

func (x *FindUserBillRequest) Reset() {
	*x = FindUserBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserBillRequest) ProtoMessage() {}

func (x *FindUserBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserBillRequest.ProtoReflect.Descriptor instead.
func (*FindUserBillRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{4}
}

func (x *FindUserBillRequest) GetUserBillId() int64 {
	if x != nil {
		return x.UserBillId
	}
	return 0
}

func (x *FindUserBillRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type FindUserBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBill *UserBill `protobuf:"bytes,1,opt,name=userBill,proto3" json:"userBill,omitempty"`
}

func (x *FindUserBillResponse) Reset() {
	*x = FindUserBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserBillResponse) ProtoMessage() {}

func (x *FindUserBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserBillResponse.ProtoReflect.Descriptor instead.
func (*FindUserBillResponse) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{5}
}

func (x *FindUserBillResponse) GetUserBill() *UserBill {
	if x != nil {
		return x.UserBill
	}
	return nil
}

// 支付账单
type PayUserBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBillId int64 `protobuf:"varint,1,opt,name=userBillId,proto3" json:"userBillId,omitempty"`
}

func (x *PayUserBillRequest) Reset() {
	*x = PayUserBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayUserBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayUserBillRequest) ProtoMessage() {}

func (x *PayUserBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayUserBillRequest.ProtoReflect.Descriptor instead.
func (*PayUserBillRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{6}
}

func (x *PayUserBillRequest) GetUserBillId() int64 {
	if x != nil {
		return x.UserBillId
	}
	return 0
}

// 计算用户所有未支付账单总额
type SumUserUnpaidBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *SumUserUnpaidBillsRequest) Reset() {
	*x = SumUserUnpaidBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumUserUnpaidBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumUserUnpaidBillsRequest) ProtoMessage() {}

func (x *SumUserUnpaidBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumUserUnpaidBillsRequest.ProtoReflect.Descriptor instead.
func (*SumUserUnpaidBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{7}
}

func (x *SumUserUnpaidBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SumUserUnpaidBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SumUserUnpaidBillsResponse) Reset() {
	*x = SumUserUnpaidBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_bill_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumUserUnpaidBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumUserUnpaidBillsResponse) ProtoMessage() {}

func (x *SumUserUnpaidBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_bill_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumUserUnpaidBillsResponse.ProtoReflect.Descriptor instead.
func (*SumUserUnpaidBillsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_bill_proto_rawDescGZIP(), []int{8}
}

func (x *SumUserUnpaidBillsResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_service_user_bill_proto protoreflect.FileDescriptor

var file_service_user_bill_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x1b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0xe8, 0x01,
	0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10,
	0x6d, 0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x61, 0x79, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x42, 0x69, 0x6c, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x42, 0x69, 0x6c, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x69, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x13,
	0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x22, 0x34, 0x0a, 0x12, 0x50, 0x61, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x22,
	0x33, 0x0a, 0x19, 0x53, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64,
	0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x1a, 0x53, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x6e, 0x70, 0x61, 0x69, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb8, 0x03, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c,
	0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x53, 0x0a, 0x12, 0x73, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e, 0x70, 0x61, 0x69,
	0x64, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_bill_proto_rawDescOnce sync.Once
	file_service_user_bill_proto_rawDescData = file_service_user_bill_proto_rawDesc
)

func file_service_user_bill_proto_rawDescGZIP() []byte {
	file_service_user_bill_proto_rawDescOnce.Do(func() {
		file_service_user_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_bill_proto_rawDescData)
	})
	return file_service_user_bill_proto_rawDescData
}

var file_service_user_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_user_bill_proto_goTypes = []interface{}{
	(*GenerateAllUserBillsRequest)(nil), // 0: pb.GenerateAllUserBillsRequest
	(*CountAllUserBillsRequest)(nil),    // 1: pb.CountAllUserBillsRequest
	(*ListUserBillsRequest)(nil),        // 2: pb.ListUserBillsRequest
	(*ListUserBillsResponse)(nil),       // 3: pb.ListUserBillsResponse
	(*FindUserBillRequest)(nil),         // 4: pb.FindUserBillRequest
	(*FindUserBillResponse)(nil),        // 5: pb.FindUserBillResponse
	(*PayUserBillRequest)(nil),          // 6: pb.PayUserBillRequest
	(*SumUserUnpaidBillsRequest)(nil),   // 7: pb.SumUserUnpaidBillsRequest
	(*SumUserUnpaidBillsResponse)(nil),  // 8: pb.SumUserUnpaidBillsResponse
	(*UserBill)(nil),                    // 9: pb.UserBill
	(*RPCSuccess)(nil),                  // 10: pb.RPCSuccess
	(*RPCCountResponse)(nil),            // 11: pb.RPCCountResponse
}
var file_service_user_bill_proto_depIdxs = []int32{
	9,  // 0: pb.ListUserBillsResponse.userBills:type_name -> pb.UserBill
	9,  // 1: pb.FindUserBillResponse.userBill:type_name -> pb.UserBill
	0,  // 2: pb.UserBillService.generateAllUserBills:input_type -> pb.GenerateAllUserBillsRequest
	1,  // 3: pb.UserBillService.countAllUserBills:input_type -> pb.CountAllUserBillsRequest
	2,  // 4: pb.UserBillService.listUserBills:input_type -> pb.ListUserBillsRequest
	4,  // 5: pb.UserBillService.findUserBill:input_type -> pb.FindUserBillRequest
	6,  // 6: pb.UserBillService.payUserBill:input_type -> pb.PayUserBillRequest
	7,  // 7: pb.UserBillService.sumUserUnpaidBills:input_type -> pb.SumUserUnpaidBillsRequest
	10, // 8: pb.UserBillService.generateAllUserBills:output_type -> pb.RPCSuccess
	11, // 9: pb.UserBillService.countAllUserBills:output_type -> pb.RPCCountResponse
	3,  // 10: pb.UserBillService.listUserBills:output_type -> pb.ListUserBillsResponse
	5,  // 11: pb.UserBillService.findUserBill:output_type -> pb.FindUserBillResponse
	10, // 12: pb.UserBillService.payUserBill:output_type -> pb.RPCSuccess
	8,  // 13: pb.UserBillService.sumUserUnpaidBills:output_type -> pb.SumUserUnpaidBillsResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_user_bill_proto_init() }
func file_service_user_bill_proto_init() {
	if File_service_user_bill_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_user_bill_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAllUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserBillsResponse); i {
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
		file_service_user_bill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserBillRequest); i {
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
		file_service_user_bill_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserBillResponse); i {
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
		file_service_user_bill_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayUserBillRequest); i {
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
		file_service_user_bill_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumUserUnpaidBillsRequest); i {
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
		file_service_user_bill_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumUserUnpaidBillsResponse); i {
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
			RawDescriptor: file_service_user_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_bill_proto_goTypes,
		DependencyIndexes: file_service_user_bill_proto_depIdxs,
		MessageInfos:      file_service_user_bill_proto_msgTypes,
	}.Build()
	File_service_user_bill_proto = out.File
	file_service_user_bill_proto_rawDesc = nil
	file_service_user_bill_proto_goTypes = nil
	file_service_user_bill_proto_depIdxs = nil
}
