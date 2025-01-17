// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: opts.proto

package ent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gen  *bool   `protobuf:"varint,1,opt,name=gen" json:"gen,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_opts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_opts_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetGen() bool {
	if x != nil && x.Gen != nil {
		return *x.Gen
	}
	return false
}

func (x *Schema) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Optional   *bool             `protobuf:"varint,1,opt,name=optional" json:"optional,omitempty"`
	Nillable   *bool             `protobuf:"varint,2,opt,name=nillable" json:"nillable,omitempty"`
	Unique     *bool             `protobuf:"varint,3,opt,name=unique" json:"unique,omitempty"`
	Sensitive  *bool             `protobuf:"varint,4,opt,name=sensitive" json:"sensitive,omitempty"`
	Immutable  *bool             `protobuf:"varint,5,opt,name=immutable" json:"immutable,omitempty"`
	Comment    *string           `protobuf:"bytes,6,opt,name=comment" json:"comment,omitempty"`
	StructTag  *string           `protobuf:"bytes,7,opt,name=struct_tag,json=structTag" json:"struct_tag,omitempty"`
	StorageKey *string           `protobuf:"bytes,8,opt,name=storage_key,json=storageKey" json:"storage_key,omitempty"`
	SchemaType map[string]string `protobuf:"bytes,9,rep,name=schema_type,json=schemaType" json:"schema_type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_opts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_opts_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetOptional() bool {
	if x != nil && x.Optional != nil {
		return *x.Optional
	}
	return false
}

func (x *Field) GetNillable() bool {
	if x != nil && x.Nillable != nil {
		return *x.Nillable
	}
	return false
}

func (x *Field) GetUnique() bool {
	if x != nil && x.Unique != nil {
		return *x.Unique
	}
	return false
}

func (x *Field) GetSensitive() bool {
	if x != nil && x.Sensitive != nil {
		return *x.Sensitive
	}
	return false
}

func (x *Field) GetImmutable() bool {
	if x != nil && x.Immutable != nil {
		return *x.Immutable
	}
	return false
}

func (x *Field) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *Field) GetStructTag() string {
	if x != nil && x.StructTag != nil {
		return *x.StructTag
	}
	return ""
}

func (x *Field) GetStorageKey() string {
	if x != nil && x.StorageKey != nil {
		return *x.StorageKey
	}
	return ""
}

func (x *Field) GetSchemaType() map[string]string {
	if x != nil {
		return x.SchemaType
	}
	return nil
}

type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unique     *bool            `protobuf:"varint,1,opt,name=unique" json:"unique,omitempty"`
	Ref        *string          `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	Required   *bool            `protobuf:"varint,3,opt,name=required" json:"required,omitempty"`
	Field      *string          `protobuf:"bytes,4,opt,name=field" json:"field,omitempty"`
	StorageKey *Edge_StorageKey `protobuf:"bytes,5,opt,name=storage_key,json=storageKey" json:"storage_key,omitempty"`
	StructTag  *string          `protobuf:"bytes,6,opt,name=struct_tag,json=structTag" json:"struct_tag,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_opts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_opts_proto_rawDescGZIP(), []int{2}
}

func (x *Edge) GetUnique() bool {
	if x != nil && x.Unique != nil {
		return *x.Unique
	}
	return false
}

func (x *Edge) GetRef() string {
	if x != nil && x.Ref != nil {
		return *x.Ref
	}
	return ""
}

func (x *Edge) GetRequired() bool {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return false
}

func (x *Edge) GetField() string {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return ""
}

func (x *Edge) GetStorageKey() *Edge_StorageKey {
	if x != nil {
		return x.StorageKey
	}
	return nil
}

func (x *Edge) GetStructTag() string {
	if x != nil && x.StructTag != nil {
		return *x.StructTag
	}
	return ""
}

type Edge_StorageKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table   *string  `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Columns []string `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
}

func (x *Edge_StorageKey) Reset() {
	*x = Edge_StorageKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge_StorageKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge_StorageKey) ProtoMessage() {}

func (x *Edge_StorageKey) ProtoReflect() protoreflect.Message {
	mi := &file_opts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge_StorageKey.ProtoReflect.Descriptor instead.
func (*Edge_StorageKey) Descriptor() ([]byte, []int) {
	return file_opts_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Edge_StorageKey) GetTable() string {
	if x != nil && x.Table != nil {
		return *x.Table
	}
	return ""
}

func (x *Edge_StorageKey) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

var file_opts_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         150119,
		Name:          "ent.schema",
		Tag:           "bytes,150119,opt,name=schema",
		Filename:      "opts.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Field)(nil),
		Field:         150119,
		Name:          "ent.field",
		Tag:           "bytes,150119,opt,name=field",
		Filename:      "opts.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Edge)(nil),
		Field:         150120,
		Name:          "ent.edge",
		Tag:           "bytes,150120,opt,name=edge",
		Filename:      "opts.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ent.Schema schema = 150119;
	E_Schema = &file_opts_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional ent.Field field = 150119;
	E_Field = &file_opts_proto_extTypes[1]
	// optional ent.Edge edge = 150120;
	E_Edge = &file_opts_proto_extTypes[2]
)

var File_opts_proto protoreflect.FileDescriptor

var file_opts_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x6e,
	0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x67, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xe9, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x69, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x61,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xf6, 0x01, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x61, 0x67, 0x1a, 0x3c, 0x0a, 0x0a, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x3a, 0x46, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xe7, 0x94, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x3a, 0x41, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe7, 0x94, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x3a, 0x3e, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe8, 0x94, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x04, 0x65,
	0x64, 0x67, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x74,
}

var (
	file_opts_proto_rawDescOnce sync.Once
	file_opts_proto_rawDescData = file_opts_proto_rawDesc
)

func file_opts_proto_rawDescGZIP() []byte {
	file_opts_proto_rawDescOnce.Do(func() {
		file_opts_proto_rawDescData = protoimpl.X.CompressGZIP(file_opts_proto_rawDescData)
	})
	return file_opts_proto_rawDescData
}

var file_opts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_opts_proto_goTypes = []interface{}{
	(*Schema)(nil),                      // 0: ent.Schema
	(*Field)(nil),                       // 1: ent.Field
	(*Edge)(nil),                        // 2: ent.Edge
	nil,                                 // 3: ent.Field.SchemaTypeEntry
	(*Edge_StorageKey)(nil),             // 4: ent.Edge.StorageKey
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 6: google.protobuf.FieldOptions
}
var file_opts_proto_depIdxs = []int32{
	3, // 0: ent.Field.schema_type:type_name -> ent.Field.SchemaTypeEntry
	4, // 1: ent.Edge.storage_key:type_name -> ent.Edge.StorageKey
	5, // 2: ent.schema:extendee -> google.protobuf.MessageOptions
	6, // 3: ent.field:extendee -> google.protobuf.FieldOptions
	6, // 4: ent.edge:extendee -> google.protobuf.FieldOptions
	0, // 5: ent.schema:type_name -> ent.Schema
	1, // 6: ent.field:type_name -> ent.Field
	2, // 7: ent.edge:type_name -> ent.Edge
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	5, // [5:8] is the sub-list for extension type_name
	2, // [2:5] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_opts_proto_init() }
func file_opts_proto_init() {
	if File_opts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
		file_opts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_opts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
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
		file_opts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge_StorageKey); i {
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
			RawDescriptor: file_opts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_opts_proto_goTypes,
		DependencyIndexes: file_opts_proto_depIdxs,
		MessageInfos:      file_opts_proto_msgTypes,
		ExtensionInfos:    file_opts_proto_extTypes,
	}.Build()
	File_opts_proto = out.File
	file_opts_proto_rawDesc = nil
	file_opts_proto_goTypes = nil
	file_opts_proto_depIdxs = nil
}
