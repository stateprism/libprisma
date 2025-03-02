// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: protoutils/extension.proto

package protoutils

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

type ExtensionType int32

const (
	ExtensionType_EMPTY     ExtensionType = 0
	ExtensionType_STRING    ExtensionType = 1
	ExtensionType_INTEGER   ExtensionType = 2
	ExtensionType_BOOLEAN   ExtensionType = 3
	ExtensionType_BYTES     ExtensionType = 4
	ExtensionType_EXTENSION ExtensionType = 5
	ExtensionType_ARRAY     ExtensionType = 6
)

// Enum value maps for ExtensionType.
var (
	ExtensionType_name = map[int32]string{
		0: "EMPTY",
		1: "STRING",
		2: "INTEGER",
		3: "BOOLEAN",
		4: "BYTES",
		5: "EXTENSION",
		6: "ARRAY",
	}
	ExtensionType_value = map[string]int32{
		"EMPTY":     0,
		"STRING":    1,
		"INTEGER":   2,
		"BOOLEAN":   3,
		"BYTES":     4,
		"EXTENSION": 5,
		"ARRAY":     6,
	}
)

func (x ExtensionType) Enum() *ExtensionType {
	p := new(ExtensionType)
	*p = x
	return p
}

func (x ExtensionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionType) Descriptor() protoreflect.EnumDescriptor {
	return file_protoutils_extension_proto_enumTypes[0].Descriptor()
}

func (ExtensionType) Type() protoreflect.EnumType {
	return &file_protoutils_extension_proto_enumTypes[0]
}

func (x ExtensionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionType.Descriptor instead.
func (ExtensionType) EnumDescriptor() ([]byte, []int) {
	return file_protoutils_extension_proto_rawDescGZIP(), []int{0}
}

type ExtensionArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Extension `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ExtensionArray) Reset() {
	*x = ExtensionArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoutils_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionArray) ProtoMessage() {}

func (x *ExtensionArray) ProtoReflect() protoreflect.Message {
	mi := &file_protoutils_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionArray.ProtoReflect.Descriptor instead.
func (*ExtensionArray) Descriptor() ([]byte, []int) {
	return file_protoutils_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionArray) GetValues() []*Extension {
	if x != nil {
		return x.Values
	}
	return nil
}

type Extensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions map[string]*Extension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsEmpty    bool                  `protobuf:"varint,2,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
}

func (x *Extensions) Reset() {
	*x = Extensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoutils_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensions) ProtoMessage() {}

func (x *Extensions) ProtoReflect() protoreflect.Message {
	mi := &file_protoutils_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensions.ProtoReflect.Descriptor instead.
func (*Extensions) Descriptor() ([]byte, []int) {
	return file_protoutils_extension_proto_rawDescGZIP(), []int{1}
}

func (x *Extensions) GetExtensions() map[string]*Extension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *Extensions) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           ExtensionType   `protobuf:"varint,1,opt,name=type,proto3,enum=ExtensionType" json:"type,omitempty"`
	StringValue    *string         `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty"`
	IntegerValue   *int64          `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof" json:"integer_value,omitempty"`
	BooleanValue   *bool           `protobuf:"varint,4,opt,name=boolean_value,json=booleanValue,proto3,oneof" json:"boolean_value,omitempty"`
	BytesValue     []byte          `protobuf:"bytes,5,opt,name=bytes_value,json=bytesValue,proto3,oneof" json:"bytes_value,omitempty"`
	ExtensionValue *Extensions     `protobuf:"bytes,6,opt,name=extension_value,json=extensionValue,proto3,oneof" json:"extension_value,omitempty"`
	ArrayValue     *ExtensionArray `protobuf:"bytes,7,opt,name=array_value,json=arrayValue,proto3,oneof" json:"array_value,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoutils_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_protoutils_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_protoutils_extension_proto_rawDescGZIP(), []int{2}
}

func (x *Extension) GetType() ExtensionType {
	if x != nil {
		return x.Type
	}
	return ExtensionType_EMPTY
}

func (x *Extension) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Extension) GetIntegerValue() int64 {
	if x != nil && x.IntegerValue != nil {
		return *x.IntegerValue
	}
	return 0
}

func (x *Extension) GetBooleanValue() bool {
	if x != nil && x.BooleanValue != nil {
		return *x.BooleanValue
	}
	return false
}

func (x *Extension) GetBytesValue() []byte {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *Extension) GetExtensionValue() *Extensions {
	if x != nil {
		return x.ExtensionValue
	}
	return nil
}

func (x *Extension) GetArrayValue() *ExtensionArray {
	if x != nil {
		return x.ArrayValue
	}
	return nil
}

var File_protoutils_extension_proto protoreflect.FileDescriptor

var file_protoutils_extension_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x22,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x49, 0x0a, 0x0f, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x03, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x03, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x04,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x05, 0x52, 0x0a, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x65, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c,
	0x45, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x04,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x06, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x72,
	0x69, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x61, 0x5f, 0x63, 0x61, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoutils_extension_proto_rawDescOnce sync.Once
	file_protoutils_extension_proto_rawDescData = file_protoutils_extension_proto_rawDesc
)

func file_protoutils_extension_proto_rawDescGZIP() []byte {
	file_protoutils_extension_proto_rawDescOnce.Do(func() {
		file_protoutils_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoutils_extension_proto_rawDescData)
	})
	return file_protoutils_extension_proto_rawDescData
}

var file_protoutils_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protoutils_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoutils_extension_proto_goTypes = []interface{}{
	(ExtensionType)(0),     // 0: ExtensionType
	(*ExtensionArray)(nil), // 1: ExtensionArray
	(*Extensions)(nil),     // 2: Extensions
	(*Extension)(nil),      // 3: Extension
	nil,                    // 4: Extensions.ExtensionsEntry
}
var file_protoutils_extension_proto_depIdxs = []int32{
	3, // 0: ExtensionArray.values:type_name -> Extension
	4, // 1: Extensions.extensions:type_name -> Extensions.ExtensionsEntry
	0, // 2: Extension.type:type_name -> ExtensionType
	2, // 3: Extension.extension_value:type_name -> Extensions
	1, // 4: Extension.array_value:type_name -> ExtensionArray
	3, // 5: Extensions.ExtensionsEntry.value:type_name -> Extension
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protoutils_extension_proto_init() }
func file_protoutils_extension_proto_init() {
	if File_protoutils_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoutils_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionArray); i {
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
		file_protoutils_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensions); i {
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
		file_protoutils_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
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
	file_protoutils_extension_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoutils_extension_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoutils_extension_proto_goTypes,
		DependencyIndexes: file_protoutils_extension_proto_depIdxs,
		EnumInfos:         file_protoutils_extension_proto_enumTypes,
		MessageInfos:      file_protoutils_extension_proto_msgTypes,
	}.Build()
	File_protoutils_extension_proto = out.File
	file_protoutils_extension_proto_rawDesc = nil
	file_protoutils_extension_proto_goTypes = nil
	file_protoutils_extension_proto_depIdxs = nil
}
