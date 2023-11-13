// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_formal_client_system.proto

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

type FormalClientSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Codes  []string `protobuf:"bytes,3,rep,name=codes,proto3" json:"codes,omitempty"`
	DataId string   `protobuf:"bytes,4,opt,name=dataId,proto3" json:"dataId,omitempty"`
	State  int32    `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *FormalClientSystem) Reset() {
	*x = FormalClientSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_formal_client_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormalClientSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormalClientSystem) ProtoMessage() {}

func (x *FormalClientSystem) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_formal_client_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormalClientSystem.ProtoReflect.Descriptor instead.
func (*FormalClientSystem) Descriptor() ([]byte, []int) {
	return file_models_model_formal_client_system_proto_rawDescGZIP(), []int{0}
}

func (x *FormalClientSystem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FormalClientSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FormalClientSystem) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *FormalClientSystem) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

func (x *FormalClientSystem) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_models_model_formal_client_system_proto protoreflect.FileDescriptor

var file_models_model_formal_client_system_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x7c, 0x0a,
	0x12, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_formal_client_system_proto_rawDescOnce sync.Once
	file_models_model_formal_client_system_proto_rawDescData = file_models_model_formal_client_system_proto_rawDesc
)

func file_models_model_formal_client_system_proto_rawDescGZIP() []byte {
	file_models_model_formal_client_system_proto_rawDescOnce.Do(func() {
		file_models_model_formal_client_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_formal_client_system_proto_rawDescData)
	})
	return file_models_model_formal_client_system_proto_rawDescData
}

var file_models_model_formal_client_system_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_formal_client_system_proto_goTypes = []interface{}{
	(*FormalClientSystem)(nil), // 0: pb.FormalClientSystem
}
var file_models_model_formal_client_system_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_formal_client_system_proto_init() }
func file_models_model_formal_client_system_proto_init() {
	if File_models_model_formal_client_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_formal_client_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormalClientSystem); i {
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
			RawDescriptor: file_models_model_formal_client_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_formal_client_system_proto_goTypes,
		DependencyIndexes: file_models_model_formal_client_system_proto_depIdxs,
		MessageInfos:      file_models_model_formal_client_system_proto_msgTypes,
	}.Build()
	File_models_model_formal_client_system_proto = out.File
	file_models_model_formal_client_system_proto_rawDesc = nil
	file_models_model_formal_client_system_proto_goTypes = nil
	file_models_model_formal_client_system_proto_depIdxs = nil
}
