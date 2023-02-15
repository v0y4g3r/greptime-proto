// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/cluster.proto

package meta

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

type BatchGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Keys   [][]byte       `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *BatchGetRequest) Reset() {
	*x = BatchGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetRequest) ProtoMessage() {}

func (x *BatchGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetRequest.ProtoReflect.Descriptor instead.
func (*BatchGetRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *BatchGetRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchGetRequest) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}

type BatchGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Kvs    []*KeyValue     `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (x *BatchGetResponse) Reset() {
	*x = BatchGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetResponse) ProtoMessage() {}

func (x *BatchGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetResponse.ProtoReflect.Descriptor instead.
func (*BatchGetResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *BatchGetResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchGetResponse) GetKvs() []*KeyValue {
	if x != nil {
		return x.Kvs
	}
	return nil
}

var File_greptime_v1_meta_cluster_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_cluster_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x1a, 0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5e, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22,
	0x7a, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x03, 0x6b, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x32, 0xa6, 0x01, 0x0a, 0x07,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_cluster_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_cluster_proto_rawDescData = file_greptime_v1_meta_cluster_proto_rawDesc
)

func file_greptime_v1_meta_cluster_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_cluster_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_cluster_proto_rawDescData)
	})
	return file_greptime_v1_meta_cluster_proto_rawDescData
}

var file_greptime_v1_meta_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greptime_v1_meta_cluster_proto_goTypes = []interface{}{
	(*BatchGetRequest)(nil),  // 0: greptime.v1.meta.BatchGetRequest
	(*BatchGetResponse)(nil), // 1: greptime.v1.meta.BatchGetResponse
	(*RequestHeader)(nil),    // 2: greptime.v1.meta.RequestHeader
	(*ResponseHeader)(nil),   // 3: greptime.v1.meta.ResponseHeader
	(*KeyValue)(nil),         // 4: greptime.v1.meta.KeyValue
	(*RangeRequest)(nil),     // 5: greptime.v1.meta.RangeRequest
	(*RangeResponse)(nil),    // 6: greptime.v1.meta.RangeResponse
}
var file_greptime_v1_meta_cluster_proto_depIdxs = []int32{
	2, // 0: greptime.v1.meta.BatchGetRequest.header:type_name -> greptime.v1.meta.RequestHeader
	3, // 1: greptime.v1.meta.BatchGetResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	4, // 2: greptime.v1.meta.BatchGetResponse.kvs:type_name -> greptime.v1.meta.KeyValue
	0, // 3: greptime.v1.meta.Cluster.BatchGet:input_type -> greptime.v1.meta.BatchGetRequest
	5, // 4: greptime.v1.meta.Cluster.Range:input_type -> greptime.v1.meta.RangeRequest
	1, // 5: greptime.v1.meta.Cluster.BatchGet:output_type -> greptime.v1.meta.BatchGetResponse
	6, // 6: greptime.v1.meta.Cluster.Range:output_type -> greptime.v1.meta.RangeResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_cluster_proto_init() }
func file_greptime_v1_meta_cluster_proto_init() {
	if File_greptime_v1_meta_cluster_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetRequest); i {
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
		file_greptime_v1_meta_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetResponse); i {
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
			RawDescriptor: file_greptime_v1_meta_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_cluster_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_cluster_proto_depIdxs,
		MessageInfos:      file_greptime_v1_meta_cluster_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_cluster_proto = out.File
	file_greptime_v1_meta_cluster_proto_rawDesc = nil
	file_greptime_v1_meta_cluster_proto_goTypes = nil
	file_greptime_v1_meta_cluster_proto_depIdxs = nil
}
