// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_metric_stat.proto

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

// 统计数据
type MetricStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash        string       `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	ServerId    int64        `protobuf:"varint,3,opt,name=serverId,proto3" json:"serverId,omitempty"`
	ItemId      int64        `protobuf:"varint,4,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Keys        []string     `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`
	Value       float32      `protobuf:"fixed32,6,opt,name=value,proto3" json:"value,omitempty"`
	Time        string       `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Version     int32        `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	NodeCluster *NodeCluster `protobuf:"bytes,30,opt,name=nodeCluster,proto3" json:"nodeCluster,omitempty"`
	Node        *Node        `protobuf:"bytes,31,opt,name=node,proto3" json:"node,omitempty"`
	Server      *Server      `protobuf:"bytes,32,opt,name=server,proto3" json:"server,omitempty"`
	SumCount    int64        `protobuf:"varint,40,opt,name=sumCount,proto3" json:"sumCount,omitempty"`
	SumTotal    float32      `protobuf:"fixed32,41,opt,name=sumTotal,proto3" json:"sumTotal,omitempty"`
}

func (x *MetricStat) Reset() {
	*x = MetricStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_metric_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStat) ProtoMessage() {}

func (x *MetricStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_metric_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStat.ProtoReflect.Descriptor instead.
func (*MetricStat) Descriptor() ([]byte, []int) {
	return file_models_model_metric_stat_proto_rawDescGZIP(), []int{0}
}

func (x *MetricStat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetricStat) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *MetricStat) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *MetricStat) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *MetricStat) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *MetricStat) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *MetricStat) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *MetricStat) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MetricStat) GetNodeCluster() *NodeCluster {
	if x != nil {
		return x.NodeCluster
	}
	return nil
}

func (x *MetricStat) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *MetricStat) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *MetricStat) GetSumCount() int64 {
	if x != nil {
		return x.SumCount
	}
	return 0
}

func (x *MetricStat) GetSumTotal() float32 {
	if x != nil {
		return x.SumTotal
	}
	return 0
}

// 需要上传的统计数据
type UploadingMetricStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash  string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Keys  []string `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	Value float32  `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UploadingMetricStat) Reset() {
	*x = UploadingMetricStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_metric_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadingMetricStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadingMetricStat) ProtoMessage() {}

func (x *UploadingMetricStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_metric_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadingMetricStat.ProtoReflect.Descriptor instead.
func (*UploadingMetricStat) Descriptor() ([]byte, []int) {
	return file_models_model_metric_stat_proto_rawDescGZIP(), []int{1}
}

func (x *UploadingMetricStat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadingMetricStat) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *UploadingMetricStat) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *UploadingMetricStat) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_models_model_metric_stat_proto protoreflect.FileDescriptor

var file_models_model_metric_stat_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x0a, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x75, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x75, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x6d,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x29, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x6d,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x63, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_metric_stat_proto_rawDescOnce sync.Once
	file_models_model_metric_stat_proto_rawDescData = file_models_model_metric_stat_proto_rawDesc
)

func file_models_model_metric_stat_proto_rawDescGZIP() []byte {
	file_models_model_metric_stat_proto_rawDescOnce.Do(func() {
		file_models_model_metric_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_metric_stat_proto_rawDescData)
	})
	return file_models_model_metric_stat_proto_rawDescData
}

var file_models_model_metric_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_model_metric_stat_proto_goTypes = []interface{}{
	(*MetricStat)(nil),          // 0: pb.MetricStat
	(*UploadingMetricStat)(nil), // 1: pb.UploadingMetricStat
	(*NodeCluster)(nil),         // 2: pb.NodeCluster
	(*Node)(nil),                // 3: pb.Node
	(*Server)(nil),              // 4: pb.Server
}
var file_models_model_metric_stat_proto_depIdxs = []int32{
	2, // 0: pb.MetricStat.nodeCluster:type_name -> pb.NodeCluster
	3, // 1: pb.MetricStat.node:type_name -> pb.Node
	4, // 2: pb.MetricStat.server:type_name -> pb.Server
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_metric_stat_proto_init() }
func file_models_model_metric_stat_proto_init() {
	if File_models_model_metric_stat_proto != nil {
		return
	}
	file_models_model_node_cluster_proto_init()
	file_models_model_node_proto_init()
	file_models_model_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_metric_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricStat); i {
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
		file_models_model_metric_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadingMetricStat); i {
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
			RawDescriptor: file_models_model_metric_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_metric_stat_proto_goTypes,
		DependencyIndexes: file_models_model_metric_stat_proto_depIdxs,
		MessageInfos:      file_models_model_metric_stat_proto_msgTypes,
	}.Build()
	File_models_model_metric_stat_proto = out.File
	file_models_model_metric_stat_proto_rawDesc = nil
	file_models_model_metric_stat_proto_goTypes = nil
	file_models_model_metric_stat_proto_depIdxs = nil
}
