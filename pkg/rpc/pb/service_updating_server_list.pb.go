// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_updating_server_list.proto

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

// 查找要更新的服务配置
type FindUpdatingServerListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastId int64 `protobuf:"varint,1,opt,name=lastId,proto3" json:"lastId,omitempty"` // 上一次读取的列表ID
}

func (x *FindUpdatingServerListsRequest) Reset() {
	*x = FindUpdatingServerListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_updating_server_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUpdatingServerListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUpdatingServerListsRequest) ProtoMessage() {}

func (x *FindUpdatingServerListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_updating_server_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUpdatingServerListsRequest.ProtoReflect.Descriptor instead.
func (*FindUpdatingServerListsRequest) Descriptor() ([]byte, []int) {
	return file_service_updating_server_list_proto_rawDescGZIP(), []int{0}
}

func (x *FindUpdatingServerListsRequest) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

type FindUpdatingServerListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServersJSON []byte `protobuf:"bytes,1,opt,name=serversJSON,proto3" json:"serversJSON,omitempty"` // 服务配置列表
	MaxId       int64  `protobuf:"varint,2,opt,name=maxId,proto3" json:"maxId,omitempty"`            // 最大的一个列表ID
}

func (x *FindUpdatingServerListsResponse) Reset() {
	*x = FindUpdatingServerListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_updating_server_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUpdatingServerListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUpdatingServerListsResponse) ProtoMessage() {}

func (x *FindUpdatingServerListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_updating_server_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUpdatingServerListsResponse.ProtoReflect.Descriptor instead.
func (*FindUpdatingServerListsResponse) Descriptor() ([]byte, []int) {
	return file_service_updating_server_list_proto_rawDescGZIP(), []int{1}
}

func (x *FindUpdatingServerListsResponse) GetServersJSON() []byte {
	if x != nil {
		return x.ServersJSON
	}
	return nil
}

func (x *FindUpdatingServerListsResponse) GetMaxId() int64 {
	if x != nil {
		return x.MaxId
	}
	return 0
}

var File_service_updating_server_list_proto protoreflect.FileDescriptor

var file_service_updating_server_list_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x59, 0x0a, 0x1f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x78, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x32, 0x7f, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x17, 0x66, 0x69,
	0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_updating_server_list_proto_rawDescOnce sync.Once
	file_service_updating_server_list_proto_rawDescData = file_service_updating_server_list_proto_rawDesc
)

func file_service_updating_server_list_proto_rawDescGZIP() []byte {
	file_service_updating_server_list_proto_rawDescOnce.Do(func() {
		file_service_updating_server_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_updating_server_list_proto_rawDescData)
	})
	return file_service_updating_server_list_proto_rawDescData
}

var file_service_updating_server_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_updating_server_list_proto_goTypes = []interface{}{
	(*FindUpdatingServerListsRequest)(nil),  // 0: pb.FindUpdatingServerListsRequest
	(*FindUpdatingServerListsResponse)(nil), // 1: pb.FindUpdatingServerListsResponse
}
var file_service_updating_server_list_proto_depIdxs = []int32{
	0, // 0: pb.UpdatingServerListService.findUpdatingServerLists:input_type -> pb.FindUpdatingServerListsRequest
	1, // 1: pb.UpdatingServerListService.findUpdatingServerLists:output_type -> pb.FindUpdatingServerListsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_updating_server_list_proto_init() }
func file_service_updating_server_list_proto_init() {
	if File_service_updating_server_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_updating_server_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUpdatingServerListsRequest); i {
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
		file_service_updating_server_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUpdatingServerListsResponse); i {
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
			RawDescriptor: file_service_updating_server_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_updating_server_list_proto_goTypes,
		DependencyIndexes: file_service_updating_server_list_proto_depIdxs,
		MessageInfos:      file_service_updating_server_list_proto_msgTypes,
	}.Build()
	File_service_updating_server_list_proto = out.File
	file_service_updating_server_list_proto_rawDesc = nil
	file_service_updating_server_list_proto_goTypes = nil
	file_service_updating_server_list_proto_depIdxs = nil
}
