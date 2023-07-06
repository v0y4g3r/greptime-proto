// Copyright 2023 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/ddl.proto

package meta

import (
	v1 "github.com/GreptimeTeam/greptime-proto/go/greptime/v1"
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

type DdlTaskType int32

const (
	DdlTaskType_Create DdlTaskType = 0
	DdlTaskType_Drop   DdlTaskType = 1
)

// Enum value maps for DdlTaskType.
var (
	DdlTaskType_name = map[int32]string{
		0: "Create",
		1: "Drop",
	}
	DdlTaskType_value = map[string]int32{
		"Create": 0,
		"Drop":   1,
	}
)

func (x DdlTaskType) Enum() *DdlTaskType {
	p := new(DdlTaskType)
	*p = x
	return p
}

func (x DdlTaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DdlTaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_ddl_proto_enumTypes[0].Descriptor()
}

func (DdlTaskType) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_ddl_proto_enumTypes[0]
}

func (x DdlTaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DdlTaskType.Descriptor instead.
func (DdlTaskType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{0}
}

type CreateTableTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTable *v1.CreateTableExpr `protobuf:"bytes,1,opt,name=create_table,json=createTable,proto3" json:"create_table,omitempty"`
	Partitions  []*Partition        `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
	TableInfo   []byte              `protobuf:"bytes,3,opt,name=table_info,json=tableInfo,proto3" json:"table_info,omitempty"`
}

func (x *CreateTableTask) Reset() {
	*x = CreateTableTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableTask) ProtoMessage() {}

func (x *CreateTableTask) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableTask.ProtoReflect.Descriptor instead.
func (*CreateTableTask) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTableTask) GetCreateTable() *v1.CreateTableExpr {
	if x != nil {
		return x.CreateTable
	}
	return nil
}

func (x *CreateTableTask) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

func (x *CreateTableTask) GetTableInfo() []byte {
	if x != nil {
		return x.TableInfo
	}
	return nil
}

type DropTableTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DropTable *v1.DropTableExpr `protobuf:"bytes,1,opt,name=drop_table,json=dropTable,proto3" json:"drop_table,omitempty"`
}

func (x *DropTableTask) Reset() {
	*x = DropTableTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropTableTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropTableTask) ProtoMessage() {}

func (x *DropTableTask) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropTableTask.ProtoReflect.Descriptor instead.
func (*DropTableTask) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{1}
}

func (x *DropTableTask) GetDropTable() *v1.DropTableExpr {
	if x != nil {
		return x.DropTable
	}
	return nil
}

type AlterTableTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlterTable *v1.AlterExpr `protobuf:"bytes,1,opt,name=alter_table,json=alterTable,proto3" json:"alter_table,omitempty"`
}

func (x *AlterTableTask) Reset() {
	*x = AlterTableTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterTableTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterTableTask) ProtoMessage() {}

func (x *AlterTableTask) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterTableTask.ProtoReflect.Descriptor instead.
func (*AlterTableTask) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{2}
}

func (x *AlterTableTask) GetAlterTable() *v1.AlterExpr {
	if x != nil {
		return x.AlterTable
	}
	return nil
}

type SubmitDdlTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Types that are assignable to Task:
	//
	//	*SubmitDdlTaskRequest_CreateTableTask
	//	*SubmitDdlTaskRequest_DropTableTask
	//	*SubmitDdlTaskRequest_AlterTableTask
	Task isSubmitDdlTaskRequest_Task `protobuf_oneof:"task"`
}

func (x *SubmitDdlTaskRequest) Reset() {
	*x = SubmitDdlTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDdlTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDdlTaskRequest) ProtoMessage() {}

func (x *SubmitDdlTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDdlTaskRequest.ProtoReflect.Descriptor instead.
func (*SubmitDdlTaskRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitDdlTaskRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (m *SubmitDdlTaskRequest) GetTask() isSubmitDdlTaskRequest_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (x *SubmitDdlTaskRequest) GetCreateTableTask() *CreateTableTask {
	if x, ok := x.GetTask().(*SubmitDdlTaskRequest_CreateTableTask); ok {
		return x.CreateTableTask
	}
	return nil
}

func (x *SubmitDdlTaskRequest) GetDropTableTask() *DropTableTask {
	if x, ok := x.GetTask().(*SubmitDdlTaskRequest_DropTableTask); ok {
		return x.DropTableTask
	}
	return nil
}

func (x *SubmitDdlTaskRequest) GetAlterTableTask() *AlterTableTask {
	if x, ok := x.GetTask().(*SubmitDdlTaskRequest_AlterTableTask); ok {
		return x.AlterTableTask
	}
	return nil
}

type isSubmitDdlTaskRequest_Task interface {
	isSubmitDdlTaskRequest_Task()
}

type SubmitDdlTaskRequest_CreateTableTask struct {
	CreateTableTask *CreateTableTask `protobuf:"bytes,2,opt,name=create_table_task,json=createTableTask,proto3,oneof"`
}

type SubmitDdlTaskRequest_DropTableTask struct {
	DropTableTask *DropTableTask `protobuf:"bytes,3,opt,name=drop_table_task,json=dropTableTask,proto3,oneof"`
}

type SubmitDdlTaskRequest_AlterTableTask struct {
	AlterTableTask *AlterTableTask `protobuf:"bytes,4,opt,name=alter_table_task,json=alterTableTask,proto3,oneof"`
}

func (*SubmitDdlTaskRequest_CreateTableTask) isSubmitDdlTaskRequest_Task() {}

func (*SubmitDdlTaskRequest_DropTableTask) isSubmitDdlTaskRequest_Task() {}

func (*SubmitDdlTaskRequest_AlterTableTask) isSubmitDdlTaskRequest_Task() {}

type SubmitDdlTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Key is the identifier for the ddl task.
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Returns if table created.
	TableId *v1.TableId `protobuf:"bytes,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
}

func (x *SubmitDdlTaskResponse) Reset() {
	*x = SubmitDdlTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDdlTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDdlTaskResponse) ProtoMessage() {}

func (x *SubmitDdlTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDdlTaskResponse.ProtoReflect.Descriptor instead.
func (*SubmitDdlTaskResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitDdlTaskResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubmitDdlTaskResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SubmitDdlTaskResponse) GetTableId() *v1.TableId {
	if x != nil {
		return x.TableId
	}
	return nil
}

var File_greptime_v1_meta_ddl_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_ddl_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1d,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x0d, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x45, 0x78, 0x70, 0x72, 0x52, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x49, 0x0a, 0x0e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x37, 0x0a, 0x0b, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0a,
	0x61, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x14, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x11,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x49, 0x0a,
	0x0f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x72, 0x6f, 0x70, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x4c, 0x0a, 0x10, 0x61, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x94,
	0x01, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x07, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x2a, 0x23, 0x0a, 0x0b, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x10, 0x01, 0x32, 0x6b, 0x0a, 0x07, 0x44, 0x64,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x60, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44,
	0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_ddl_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_ddl_proto_rawDescData = file_greptime_v1_meta_ddl_proto_rawDesc
)

func file_greptime_v1_meta_ddl_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_ddl_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_ddl_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_ddl_proto_rawDescData)
	})
	return file_greptime_v1_meta_ddl_proto_rawDescData
}

var file_greptime_v1_meta_ddl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_meta_ddl_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_greptime_v1_meta_ddl_proto_goTypes = []interface{}{
	(DdlTaskType)(0),              // 0: greptime.v1.meta.DdlTaskType
	(*CreateTableTask)(nil),       // 1: greptime.v1.meta.CreateTableTask
	(*DropTableTask)(nil),         // 2: greptime.v1.meta.DropTableTask
	(*AlterTableTask)(nil),        // 3: greptime.v1.meta.AlterTableTask
	(*SubmitDdlTaskRequest)(nil),  // 4: greptime.v1.meta.SubmitDdlTaskRequest
	(*SubmitDdlTaskResponse)(nil), // 5: greptime.v1.meta.SubmitDdlTaskResponse
	(*v1.CreateTableExpr)(nil),    // 6: greptime.v1.CreateTableExpr
	(*Partition)(nil),             // 7: greptime.v1.meta.Partition
	(*v1.DropTableExpr)(nil),      // 8: greptime.v1.DropTableExpr
	(*v1.AlterExpr)(nil),          // 9: greptime.v1.AlterExpr
	(*RequestHeader)(nil),         // 10: greptime.v1.meta.RequestHeader
	(*ResponseHeader)(nil),        // 11: greptime.v1.meta.ResponseHeader
	(*v1.TableId)(nil),            // 12: greptime.v1.TableId
}
var file_greptime_v1_meta_ddl_proto_depIdxs = []int32{
	6,  // 0: greptime.v1.meta.CreateTableTask.create_table:type_name -> greptime.v1.CreateTableExpr
	7,  // 1: greptime.v1.meta.CreateTableTask.partitions:type_name -> greptime.v1.meta.Partition
	8,  // 2: greptime.v1.meta.DropTableTask.drop_table:type_name -> greptime.v1.DropTableExpr
	9,  // 3: greptime.v1.meta.AlterTableTask.alter_table:type_name -> greptime.v1.AlterExpr
	10, // 4: greptime.v1.meta.SubmitDdlTaskRequest.header:type_name -> greptime.v1.meta.RequestHeader
	1,  // 5: greptime.v1.meta.SubmitDdlTaskRequest.create_table_task:type_name -> greptime.v1.meta.CreateTableTask
	2,  // 6: greptime.v1.meta.SubmitDdlTaskRequest.drop_table_task:type_name -> greptime.v1.meta.DropTableTask
	3,  // 7: greptime.v1.meta.SubmitDdlTaskRequest.alter_table_task:type_name -> greptime.v1.meta.AlterTableTask
	11, // 8: greptime.v1.meta.SubmitDdlTaskResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	12, // 9: greptime.v1.meta.SubmitDdlTaskResponse.table_id:type_name -> greptime.v1.TableId
	4,  // 10: greptime.v1.meta.DdlTask.SubmitDdlTask:input_type -> greptime.v1.meta.SubmitDdlTaskRequest
	5,  // 11: greptime.v1.meta.DdlTask.SubmitDdlTask:output_type -> greptime.v1.meta.SubmitDdlTaskResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_ddl_proto_init() }
func file_greptime_v1_meta_ddl_proto_init() {
	if File_greptime_v1_meta_ddl_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_ddl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableTask); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropTableTask); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterTableTask); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDdlTaskRequest); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDdlTaskResponse); i {
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
	file_greptime_v1_meta_ddl_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SubmitDdlTaskRequest_CreateTableTask)(nil),
		(*SubmitDdlTaskRequest_DropTableTask)(nil),
		(*SubmitDdlTaskRequest_AlterTableTask)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greptime_v1_meta_ddl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_ddl_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_ddl_proto_depIdxs,
		EnumInfos:         file_greptime_v1_meta_ddl_proto_enumTypes,
		MessageInfos:      file_greptime_v1_meta_ddl_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_ddl_proto = out.File
	file_greptime_v1_meta_ddl_proto_rawDesc = nil
	file_greptime_v1_meta_ddl_proto_goTypes = nil
	file_greptime_v1_meta_ddl_proto_depIdxs = nil
}
