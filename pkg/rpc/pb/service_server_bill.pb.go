// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_server_bill.proto

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

// 查询服务账单数量
type CountAllServerBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Month  string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *CountAllServerBillsRequest) Reset() {
	*x = CountAllServerBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllServerBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllServerBillsRequest) ProtoMessage() {}

func (x *CountAllServerBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllServerBillsRequest.ProtoReflect.Descriptor instead.
func (*CountAllServerBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_bill_proto_rawDescGZIP(), []int{0}
}

func (x *CountAllServerBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountAllServerBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

// 查询服务账单列表
type ListServerBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Month  string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListServerBillsRequest) Reset() {
	*x = ListServerBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServerBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerBillsRequest) ProtoMessage() {}

func (x *ListServerBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerBillsRequest.ProtoReflect.Descriptor instead.
func (*ListServerBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_bill_proto_rawDescGZIP(), []int{1}
}

func (x *ListServerBillsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListServerBillsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *ListServerBillsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListServerBillsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListServerBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerBills []*ServerBill `protobuf:"bytes,1,rep,name=serverBills,proto3" json:"serverBills,omitempty"`
}

func (x *ListServerBillsResponse) Reset() {
	*x = ListServerBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServerBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerBillsResponse) ProtoMessage() {}

func (x *ListServerBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_bill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerBillsResponse.ProtoReflect.Descriptor instead.
func (*ListServerBillsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_bill_proto_rawDescGZIP(), []int{2}
}

func (x *ListServerBillsResponse) GetServerBills() []*ServerBill {
	if x != nil {
		return x.ServerBills
	}
	return nil
}

var File_service_server_bill_proto protoreflect.FileDescriptor

var file_service_server_bill_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x1a, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x72, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x32, 0xac, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42,
	0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x6c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_bill_proto_rawDescOnce sync.Once
	file_service_server_bill_proto_rawDescData = file_service_server_bill_proto_rawDesc
)

func file_service_server_bill_proto_rawDescGZIP() []byte {
	file_service_server_bill_proto_rawDescOnce.Do(func() {
		file_service_server_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_bill_proto_rawDescData)
	})
	return file_service_server_bill_proto_rawDescData
}

var file_service_server_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_server_bill_proto_goTypes = []interface{}{
	(*CountAllServerBillsRequest)(nil), // 0: pb.CountAllServerBillsRequest
	(*ListServerBillsRequest)(nil),     // 1: pb.ListServerBillsRequest
	(*ListServerBillsResponse)(nil),    // 2: pb.ListServerBillsResponse
	(*ServerBill)(nil),                 // 3: pb.ServerBill
	(*RPCCountResponse)(nil),           // 4: pb.RPCCountResponse
}
var file_service_server_bill_proto_depIdxs = []int32{
	3, // 0: pb.ListServerBillsResponse.serverBills:type_name -> pb.ServerBill
	0, // 1: pb.ServerBillService.countAllServerBills:input_type -> pb.CountAllServerBillsRequest
	1, // 2: pb.ServerBillService.listServerBills:input_type -> pb.ListServerBillsRequest
	4, // 3: pb.ServerBillService.countAllServerBills:output_type -> pb.RPCCountResponse
	2, // 4: pb.ServerBillService.listServerBills:output_type -> pb.ListServerBillsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_server_bill_proto_init() }
func file_service_server_bill_proto_init() {
	if File_service_server_bill_proto != nil {
		return
	}
	file_models_model_server_bill_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllServerBillsRequest); i {
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
		file_service_server_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServerBillsRequest); i {
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
		file_service_server_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServerBillsResponse); i {
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
			RawDescriptor: file_service_server_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_bill_proto_goTypes,
		DependencyIndexes: file_service_server_bill_proto_depIdxs,
		MessageInfos:      file_service_server_bill_proto_msgTypes,
	}.Build()
	File_service_server_bill_proto = out.File
	file_service_server_bill_proto_rawDesc = nil
	file_service_server_bill_proto_goTypes = nil
	file_service_server_bill_proto_depIdxs = nil
}