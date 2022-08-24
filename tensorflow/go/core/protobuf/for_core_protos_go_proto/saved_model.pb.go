// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: tensorflow/core/protobuf/saved_model.proto

package for_core_protos_go_proto

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

// SavedModel is the high level serialization format for TensorFlow Models.
// See [todo: doc links, similar to session_bundle] for more information.
type SavedModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The schema version of the SavedModel instance. Used for versioning when
	// making future changes to the specification/implementation. Initial value
	// at release will be 1.
	SavedModelSchemaVersion int64 `protobuf:"varint,1,opt,name=saved_model_schema_version,json=savedModelSchemaVersion,proto3" json:"saved_model_schema_version,omitempty"`
	// One or more MetaGraphs.
	MetaGraphs []*MetaGraphDef `protobuf:"bytes,2,rep,name=meta_graphs,json=metaGraphs,proto3" json:"meta_graphs,omitempty"`
}

func (x *SavedModel) Reset() {
	*x = SavedModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_saved_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedModel) ProtoMessage() {}

func (x *SavedModel) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_saved_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedModel.ProtoReflect.Descriptor instead.
func (*SavedModel) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_saved_model_proto_rawDescGZIP(), []int{0}
}

func (x *SavedModel) GetSavedModelSchemaVersion() int64 {
	if x != nil {
		return x.SavedModelSchemaVersion
	}
	return 0
}

func (x *SavedModel) GetMetaGraphs() []*MetaGraphDef {
	if x != nil {
		return x.MetaGraphs
	}
	return nil
}

var File_tensorflow_core_protobuf_saved_model_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_saved_model_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x73, 0x61, 0x76, 0x65, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x61, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x42, 0x88, 0x01, 0x0a, 0x18, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x10, 0x53, 0x61, 0x76, 0x65, 0x64, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_saved_model_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_saved_model_proto_rawDescData = file_tensorflow_core_protobuf_saved_model_proto_rawDesc
)

func file_tensorflow_core_protobuf_saved_model_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_saved_model_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_saved_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_saved_model_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_saved_model_proto_rawDescData
}

var file_tensorflow_core_protobuf_saved_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_saved_model_proto_goTypes = []interface{}{
	(*SavedModel)(nil),   // 0: tensorflow.SavedModel
	(*MetaGraphDef)(nil), // 1: tensorflow.MetaGraphDef
}
var file_tensorflow_core_protobuf_saved_model_proto_depIdxs = []int32{
	1, // 0: tensorflow.SavedModel.meta_graphs:type_name -> tensorflow.MetaGraphDef
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_saved_model_proto_init() }
func file_tensorflow_core_protobuf_saved_model_proto_init() {
	if File_tensorflow_core_protobuf_saved_model_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_meta_graph_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_saved_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedModel); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_saved_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_saved_model_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_saved_model_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_saved_model_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_saved_model_proto = out.File
	file_tensorflow_core_protobuf_saved_model_proto_rawDesc = nil
	file_tensorflow_core_protobuf_saved_model_proto_goTypes = nil
	file_tensorflow_core_protobuf_saved_model_proto_depIdxs = nil
}
