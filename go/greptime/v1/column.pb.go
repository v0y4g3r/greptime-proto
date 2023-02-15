// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/column.proto

package v1

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

type ColumnDataType int32

const (
	ColumnDataType_BOOLEAN               ColumnDataType = 0
	ColumnDataType_INT8                  ColumnDataType = 1
	ColumnDataType_INT16                 ColumnDataType = 2
	ColumnDataType_INT32                 ColumnDataType = 3
	ColumnDataType_INT64                 ColumnDataType = 4
	ColumnDataType_UINT8                 ColumnDataType = 5
	ColumnDataType_UINT16                ColumnDataType = 6
	ColumnDataType_UINT32                ColumnDataType = 7
	ColumnDataType_UINT64                ColumnDataType = 8
	ColumnDataType_FLOAT32               ColumnDataType = 9
	ColumnDataType_FLOAT64               ColumnDataType = 10
	ColumnDataType_BINARY                ColumnDataType = 11
	ColumnDataType_STRING                ColumnDataType = 12
	ColumnDataType_DATE                  ColumnDataType = 13
	ColumnDataType_DATETIME              ColumnDataType = 14
	ColumnDataType_TIMESTAMP_SECOND      ColumnDataType = 15
	ColumnDataType_TIMESTAMP_MILLISECOND ColumnDataType = 16
	ColumnDataType_TIMESTAMP_MICROSECOND ColumnDataType = 17
	ColumnDataType_TIMESTAMP_NANOSECOND  ColumnDataType = 18
)

// Enum value maps for ColumnDataType.
var (
	ColumnDataType_name = map[int32]string{
		0:  "BOOLEAN",
		1:  "INT8",
		2:  "INT16",
		3:  "INT32",
		4:  "INT64",
		5:  "UINT8",
		6:  "UINT16",
		7:  "UINT32",
		8:  "UINT64",
		9:  "FLOAT32",
		10: "FLOAT64",
		11: "BINARY",
		12: "STRING",
		13: "DATE",
		14: "DATETIME",
		15: "TIMESTAMP_SECOND",
		16: "TIMESTAMP_MILLISECOND",
		17: "TIMESTAMP_MICROSECOND",
		18: "TIMESTAMP_NANOSECOND",
	}
	ColumnDataType_value = map[string]int32{
		"BOOLEAN":               0,
		"INT8":                  1,
		"INT16":                 2,
		"INT32":                 3,
		"INT64":                 4,
		"UINT8":                 5,
		"UINT16":                6,
		"UINT32":                7,
		"UINT64":                8,
		"FLOAT32":               9,
		"FLOAT64":               10,
		"BINARY":                11,
		"STRING":                12,
		"DATE":                  13,
		"DATETIME":              14,
		"TIMESTAMP_SECOND":      15,
		"TIMESTAMP_MILLISECOND": 16,
		"TIMESTAMP_MICROSECOND": 17,
		"TIMESTAMP_NANOSECOND":  18,
	}
)

func (x ColumnDataType) Enum() *ColumnDataType {
	p := new(ColumnDataType)
	*p = x
	return p
}

func (x ColumnDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_column_proto_enumTypes[0].Descriptor()
}

func (ColumnDataType) Type() protoreflect.EnumType {
	return &file_greptime_v1_column_proto_enumTypes[0]
}

func (x ColumnDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnDataType.Descriptor instead.
func (ColumnDataType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0}
}

type Column_SemanticType int32

const (
	Column_TAG       Column_SemanticType = 0
	Column_FIELD     Column_SemanticType = 1
	Column_TIMESTAMP Column_SemanticType = 2
)

// Enum value maps for Column_SemanticType.
var (
	Column_SemanticType_name = map[int32]string{
		0: "TAG",
		1: "FIELD",
		2: "TIMESTAMP",
	}
	Column_SemanticType_value = map[string]int32{
		"TAG":       0,
		"FIELD":     1,
		"TIMESTAMP": 2,
	}
)

func (x Column_SemanticType) Enum() *Column_SemanticType {
	p := new(Column_SemanticType)
	*p = x
	return p
}

func (x Column_SemanticType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Column_SemanticType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_column_proto_enumTypes[1].Descriptor()
}

func (Column_SemanticType) Type() protoreflect.EnumType {
	return &file_greptime_v1_column_proto_enumTypes[1]
}

func (x Column_SemanticType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Column_SemanticType.Descriptor instead.
func (Column_SemanticType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0, 0}
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName   string              `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	SemanticType Column_SemanticType `protobuf:"varint,2,opt,name=semantic_type,json=semanticType,proto3,enum=greptime.v1.Column_SemanticType" json:"semantic_type,omitempty"`
	// The array of non-null values in this column.
	//
	// For example: suppose there is a column "foo" that contains some int32 values (1, 2, 3, 4, 5, null, 7, 8, 9, null);
	//
	//	column:
	//	  column_name: foo
	//	  semantic_type: Tag
	//	  values: 1, 2, 3, 4, 5, 7, 8, 9
	//	  null_masks: 00100000 00000010
	Values *Column_Values `protobuf:"bytes,3,opt,name=values,proto3" json:"values,omitempty"`
	// Mask maps the positions of null values.
	// If a bit in null_mask is 1, it indicates that the column value at that position is null.
	NullMask []byte `protobuf:"bytes,4,opt,name=null_mask,json=nullMask,proto3" json:"null_mask,omitempty"`
	// Helpful in creating vector from column.
	Datatype ColumnDataType `protobuf:"varint,5,opt,name=datatype,proto3,enum=greptime.v1.ColumnDataType" json:"datatype,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_column_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_column_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0}
}

func (x *Column) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *Column) GetSemanticType() Column_SemanticType {
	if x != nil {
		return x.SemanticType
	}
	return Column_TAG
}

func (x *Column) GetValues() *Column_Values {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Column) GetNullMask() []byte {
	if x != nil {
		return x.NullMask
	}
	return nil
}

func (x *Column) GetDatatype() ColumnDataType {
	if x != nil {
		return x.Datatype
	}
	return ColumnDataType_BOOLEAN
}

type ColumnDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Datatype          ColumnDataType `protobuf:"varint,2,opt,name=datatype,proto3,enum=greptime.v1.ColumnDataType" json:"datatype,omitempty"`
	IsNullable        bool           `protobuf:"varint,3,opt,name=is_nullable,json=isNullable,proto3" json:"is_nullable,omitempty"`
	DefaultConstraint []byte         `protobuf:"bytes,4,opt,name=default_constraint,json=defaultConstraint,proto3" json:"default_constraint,omitempty"`
}

func (x *ColumnDef) Reset() {
	*x = ColumnDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_column_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnDef) ProtoMessage() {}

func (x *ColumnDef) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_column_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnDef.ProtoReflect.Descriptor instead.
func (*ColumnDef) Descriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{1}
}

func (x *ColumnDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnDef) GetDatatype() ColumnDataType {
	if x != nil {
		return x.Datatype
	}
	return ColumnDataType_BOOLEAN
}

func (x *ColumnDef) GetIsNullable() bool {
	if x != nil {
		return x.IsNullable
	}
	return false
}

func (x *ColumnDef) GetDefaultConstraint() []byte {
	if x != nil {
		return x.DefaultConstraint
	}
	return nil
}

type Column_Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I8Values            []int32   `protobuf:"varint,1,rep,packed,name=i8_values,json=i8Values,proto3" json:"i8_values,omitempty"`
	I16Values           []int32   `protobuf:"varint,2,rep,packed,name=i16_values,json=i16Values,proto3" json:"i16_values,omitempty"`
	I32Values           []int32   `protobuf:"varint,3,rep,packed,name=i32_values,json=i32Values,proto3" json:"i32_values,omitempty"`
	I64Values           []int64   `protobuf:"varint,4,rep,packed,name=i64_values,json=i64Values,proto3" json:"i64_values,omitempty"`
	U8Values            []uint32  `protobuf:"varint,5,rep,packed,name=u8_values,json=u8Values,proto3" json:"u8_values,omitempty"`
	U16Values           []uint32  `protobuf:"varint,6,rep,packed,name=u16_values,json=u16Values,proto3" json:"u16_values,omitempty"`
	U32Values           []uint32  `protobuf:"varint,7,rep,packed,name=u32_values,json=u32Values,proto3" json:"u32_values,omitempty"`
	U64Values           []uint64  `protobuf:"varint,8,rep,packed,name=u64_values,json=u64Values,proto3" json:"u64_values,omitempty"`
	F32Values           []float32 `protobuf:"fixed32,9,rep,packed,name=f32_values,json=f32Values,proto3" json:"f32_values,omitempty"`
	F64Values           []float64 `protobuf:"fixed64,10,rep,packed,name=f64_values,json=f64Values,proto3" json:"f64_values,omitempty"`
	BoolValues          []bool    `protobuf:"varint,11,rep,packed,name=bool_values,json=boolValues,proto3" json:"bool_values,omitempty"`
	BinaryValues        [][]byte  `protobuf:"bytes,12,rep,name=binary_values,json=binaryValues,proto3" json:"binary_values,omitempty"`
	StringValues        []string  `protobuf:"bytes,13,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty"`
	DateValues          []int32   `protobuf:"varint,14,rep,packed,name=date_values,json=dateValues,proto3" json:"date_values,omitempty"`
	DatetimeValues      []int64   `protobuf:"varint,15,rep,packed,name=datetime_values,json=datetimeValues,proto3" json:"datetime_values,omitempty"`
	TsSecondValues      []int64   `protobuf:"varint,16,rep,packed,name=ts_second_values,json=tsSecondValues,proto3" json:"ts_second_values,omitempty"`
	TsMillisecondValues []int64   `protobuf:"varint,17,rep,packed,name=ts_millisecond_values,json=tsMillisecondValues,proto3" json:"ts_millisecond_values,omitempty"`
	TsMicrosecondValues []int64   `protobuf:"varint,18,rep,packed,name=ts_microsecond_values,json=tsMicrosecondValues,proto3" json:"ts_microsecond_values,omitempty"`
	TsNanosecondValues  []int64   `protobuf:"varint,19,rep,packed,name=ts_nanosecond_values,json=tsNanosecondValues,proto3" json:"ts_nanosecond_values,omitempty"`
}

func (x *Column_Values) Reset() {
	*x = Column_Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_column_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column_Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column_Values) ProtoMessage() {}

func (x *Column_Values) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_column_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column_Values.ProtoReflect.Descriptor instead.
func (*Column_Values) Descriptor() ([]byte, []int) {
	return file_greptime_v1_column_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Column_Values) GetI8Values() []int32 {
	if x != nil {
		return x.I8Values
	}
	return nil
}

func (x *Column_Values) GetI16Values() []int32 {
	if x != nil {
		return x.I16Values
	}
	return nil
}

func (x *Column_Values) GetI32Values() []int32 {
	if x != nil {
		return x.I32Values
	}
	return nil
}

func (x *Column_Values) GetI64Values() []int64 {
	if x != nil {
		return x.I64Values
	}
	return nil
}

func (x *Column_Values) GetU8Values() []uint32 {
	if x != nil {
		return x.U8Values
	}
	return nil
}

func (x *Column_Values) GetU16Values() []uint32 {
	if x != nil {
		return x.U16Values
	}
	return nil
}

func (x *Column_Values) GetU32Values() []uint32 {
	if x != nil {
		return x.U32Values
	}
	return nil
}

func (x *Column_Values) GetU64Values() []uint64 {
	if x != nil {
		return x.U64Values
	}
	return nil
}

func (x *Column_Values) GetF32Values() []float32 {
	if x != nil {
		return x.F32Values
	}
	return nil
}

func (x *Column_Values) GetF64Values() []float64 {
	if x != nil {
		return x.F64Values
	}
	return nil
}

func (x *Column_Values) GetBoolValues() []bool {
	if x != nil {
		return x.BoolValues
	}
	return nil
}

func (x *Column_Values) GetBinaryValues() [][]byte {
	if x != nil {
		return x.BinaryValues
	}
	return nil
}

func (x *Column_Values) GetStringValues() []string {
	if x != nil {
		return x.StringValues
	}
	return nil
}

func (x *Column_Values) GetDateValues() []int32 {
	if x != nil {
		return x.DateValues
	}
	return nil
}

func (x *Column_Values) GetDatetimeValues() []int64 {
	if x != nil {
		return x.DatetimeValues
	}
	return nil
}

func (x *Column_Values) GetTsSecondValues() []int64 {
	if x != nil {
		return x.TsSecondValues
	}
	return nil
}

func (x *Column_Values) GetTsMillisecondValues() []int64 {
	if x != nil {
		return x.TsMillisecondValues
	}
	return nil
}

func (x *Column_Values) GetTsMicrosecondValues() []int64 {
	if x != nil {
		return x.TsMicrosecondValues
	}
	return nil
}

func (x *Column_Values) GetTsNanosecondValues() []int64 {
	if x != nil {
		return x.TsNanosecondValues
	}
	return nil
}

var File_greptime_v1_column_proto protoreflect.FileDescriptor

var file_greptime_v1_column_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xe3, 0x07, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e,
	0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x37, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x1a, 0xb3, 0x05, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x38, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x38, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x31, 0x36, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x09, 0x69, 0x31, 0x36, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x69, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x36,
	0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09,
	0x69, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x38, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x38,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x31, 0x36, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x31, 0x36, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x75, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x02, 0x52, 0x09, 0x66, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x66, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0e, 0x74, 0x73, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x15, 0x74, 0x73, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x13, 0x74, 0x73, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x13, 0x74, 0x73, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x73, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x74, 0x73, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0c, 0x53, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x41,
	0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x02, 0x22, 0xa8, 0x01,
	0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x37, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6e,
	0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2a, 0xa7, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x54, 0x38,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36,
	0x34, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x05, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34,
	0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x33, 0x32, 0x10, 0x09, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44,
	0x10, 0x0f, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f,
	0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x10, 0x12, 0x19, 0x0a,
	0x15, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f,
	0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x11, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x4d, 0x45,
	0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4e, 0x41, 0x4e, 0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44,
	0x10, 0x12, 0x42, 0x50, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x07, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_column_proto_rawDescOnce sync.Once
	file_greptime_v1_column_proto_rawDescData = file_greptime_v1_column_proto_rawDesc
)

func file_greptime_v1_column_proto_rawDescGZIP() []byte {
	file_greptime_v1_column_proto_rawDescOnce.Do(func() {
		file_greptime_v1_column_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_column_proto_rawDescData)
	})
	return file_greptime_v1_column_proto_rawDescData
}

var file_greptime_v1_column_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_greptime_v1_column_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_greptime_v1_column_proto_goTypes = []interface{}{
	(ColumnDataType)(0),      // 0: greptime.v1.ColumnDataType
	(Column_SemanticType)(0), // 1: greptime.v1.Column.SemanticType
	(*Column)(nil),           // 2: greptime.v1.Column
	(*ColumnDef)(nil),        // 3: greptime.v1.ColumnDef
	(*Column_Values)(nil),    // 4: greptime.v1.Column.Values
}
var file_greptime_v1_column_proto_depIdxs = []int32{
	1, // 0: greptime.v1.Column.semantic_type:type_name -> greptime.v1.Column.SemanticType
	4, // 1: greptime.v1.Column.values:type_name -> greptime.v1.Column.Values
	0, // 2: greptime.v1.Column.datatype:type_name -> greptime.v1.ColumnDataType
	0, // 3: greptime.v1.ColumnDef.datatype:type_name -> greptime.v1.ColumnDataType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_greptime_v1_column_proto_init() }
func file_greptime_v1_column_proto_init() {
	if File_greptime_v1_column_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_column_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_greptime_v1_column_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnDef); i {
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
		file_greptime_v1_column_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column_Values); i {
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
			RawDescriptor: file_greptime_v1_column_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greptime_v1_column_proto_goTypes,
		DependencyIndexes: file_greptime_v1_column_proto_depIdxs,
		EnumInfos:         file_greptime_v1_column_proto_enumTypes,
		MessageInfos:      file_greptime_v1_column_proto_msgTypes,
	}.Build()
	File_greptime_v1_column_proto = out.File
	file_greptime_v1_column_proto_rawDesc = nil
	file_greptime_v1_column_proto_goTypes = nil
	file_greptime_v1_column_proto_depIdxs = nil
}
