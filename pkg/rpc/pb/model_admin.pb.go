// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_admin.proto

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

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                            // ID
	Fullname        string         `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`                 // 全称
	Username        string         `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`                 // 用户名
	IsOn            bool           `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`                        // 是否启用
	IsSuper         bool           `protobuf:"varint,5,opt,name=isSuper,proto3" json:"isSuper,omitempty"`                  // 是否为超级用户
	CreatedAt       int64          `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`              // 创建时间
	Modules         []*AdminModule `protobuf:"bytes,7,rep,name=Modules,proto3" json:"Modules,omitempty"`                   // 有权限的模块
	OtpLogin        *Login         `protobuf:"bytes,8,opt,name=otpLogin,proto3" json:"otpLogin,omitempty"`                 // OTP认证
	CanLogin        bool           `protobuf:"varint,9,opt,name=canLogin,proto3" json:"canLogin,omitempty"`                // 是否可以登录
	HasWeakPassword bool           `protobuf:"varint,10,opt,name=hasWeakPassword,proto3" json:"hasWeakPassword,omitempty"` // 是否设置了弱密码，只有超级管理员能看到此项
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_models_model_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Admin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Admin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Admin) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *Admin) GetIsSuper() bool {
	if x != nil {
		return x.IsSuper
	}
	return false
}

func (x *Admin) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Admin) GetModules() []*AdminModule {
	if x != nil {
		return x.Modules
	}
	return nil
}

func (x *Admin) GetOtpLogin() *Login {
	if x != nil {
		return x.OtpLogin
	}
	return nil
}

func (x *Admin) GetCanLogin() bool {
	if x != nil {
		return x.CanLogin
	}
	return false
}

func (x *Admin) GetHasWeakPassword() bool {
	if x != nil {
		return x.HasWeakPassword
	}
	return false
}

var File_models_model_admin_proto protoreflect.FileDescriptor

var file_models_model_admin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x05, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x74, 0x70, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x08, 0x6f, 0x74, 0x70, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x57, 0x65, 0x61, 0x6b,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x68, 0x61, 0x73, 0x57, 0x65, 0x61, 0x6b, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_admin_proto_rawDescOnce sync.Once
	file_models_model_admin_proto_rawDescData = file_models_model_admin_proto_rawDesc
)

func file_models_model_admin_proto_rawDescGZIP() []byte {
	file_models_model_admin_proto_rawDescOnce.Do(func() {
		file_models_model_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_admin_proto_rawDescData)
	})
	return file_models_model_admin_proto_rawDescData
}

var file_models_model_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_admin_proto_goTypes = []interface{}{
	(*Admin)(nil),       // 0: pb.Admin
	(*AdminModule)(nil), // 1: pb.AdminModule
	(*Login)(nil),       // 2: pb.Login
}
var file_models_model_admin_proto_depIdxs = []int32{
	1, // 0: pb.Admin.Modules:type_name -> pb.AdminModule
	2, // 1: pb.Admin.otpLogin:type_name -> pb.Login
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_admin_proto_init() }
func file_models_model_admin_proto_init() {
	if File_models_model_admin_proto != nil {
		return
	}
	file_models_model_admin_module_proto_init()
	file_models_model_login_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
			RawDescriptor: file_models_model_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_admin_proto_goTypes,
		DependencyIndexes: file_models_model_admin_proto_depIdxs,
		MessageInfos:      file_models_model_admin_proto_msgTypes,
	}.Build()
	File_models_model_admin_proto = out.File
	file_models_model_admin_proto_rawDesc = nil
	file_models_model_admin_proto_goTypes = nil
	file_models_model_admin_proto_depIdxs = nil
}