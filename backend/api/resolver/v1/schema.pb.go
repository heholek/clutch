// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resolver/v1/schema.proto

package resolverv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StringField struct {
	Placeholder          string   `protobuf:"bytes,1,opt,name=placeholder,proto3" json:"placeholder,omitempty"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringField) Reset()         { *m = StringField{} }
func (m *StringField) String() string { return proto.CompactTextString(m) }
func (*StringField) ProtoMessage()    {}
func (*StringField) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{0}
}

func (m *StringField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringField.Unmarshal(m, b)
}
func (m *StringField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringField.Marshal(b, m, deterministic)
}
func (m *StringField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringField.Merge(m, src)
}
func (m *StringField) XXX_Size() int {
	return xxx_messageInfo_StringField.Size(m)
}
func (m *StringField) XXX_DiscardUnknown() {
	xxx_messageInfo_StringField.DiscardUnknown(m)
}

var xxx_messageInfo_StringField proto.InternalMessageInfo

func (m *StringField) GetPlaceholder() string {
	if m != nil {
		return m.Placeholder
	}
	return ""
}

func (m *StringField) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type Option struct {
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Option_StringValue
	Value                isOption_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{1}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type isOption_Value interface {
	isOption_Value()
}

type Option_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*Option_StringValue) isOption_Value() {}

func (m *Option) GetValue() isOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Option) GetStringValue() string {
	if x, ok := m.GetValue().(*Option_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Option) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Option_StringValue)(nil),
	}
}

type OptionField struct {
	IncludeAllOption      bool      `protobuf:"varint,1,opt,name=include_all_option,json=includeAllOption,proto3" json:"include_all_option,omitempty"`
	IncludeDynamicOptions []string  `protobuf:"bytes,2,rep,name=include_dynamic_options,json=includeDynamicOptions,proto3" json:"include_dynamic_options,omitempty"`
	Options               []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}  `json:"-"`
	XXX_unrecognized      []byte    `json:"-"`
	XXX_sizecache         int32     `json:"-"`
}

func (m *OptionField) Reset()         { *m = OptionField{} }
func (m *OptionField) String() string { return proto.CompactTextString(m) }
func (*OptionField) ProtoMessage()    {}
func (*OptionField) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{2}
}

func (m *OptionField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionField.Unmarshal(m, b)
}
func (m *OptionField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionField.Marshal(b, m, deterministic)
}
func (m *OptionField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionField.Merge(m, src)
}
func (m *OptionField) XXX_Size() int {
	return xxx_messageInfo_OptionField.Size(m)
}
func (m *OptionField) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionField.DiscardUnknown(m)
}

var xxx_messageInfo_OptionField proto.InternalMessageInfo

func (m *OptionField) GetIncludeAllOption() bool {
	if m != nil {
		return m.IncludeAllOption
	}
	return false
}

func (m *OptionField) GetIncludeDynamicOptions() []string {
	if m != nil {
		return m.IncludeDynamicOptions
	}
	return nil
}

func (m *OptionField) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type Field struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Metadata             *FieldMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{3}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetMetadata() *FieldMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type FieldMetadata struct {
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Required    bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*FieldMetadata_StringField
	//	*FieldMetadata_OptionField
	Type                 isFieldMetadata_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FieldMetadata) Reset()         { *m = FieldMetadata{} }
func (m *FieldMetadata) String() string { return proto.CompactTextString(m) }
func (*FieldMetadata) ProtoMessage()    {}
func (*FieldMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{4}
}

func (m *FieldMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMetadata.Unmarshal(m, b)
}
func (m *FieldMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMetadata.Marshal(b, m, deterministic)
}
func (m *FieldMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMetadata.Merge(m, src)
}
func (m *FieldMetadata) XXX_Size() int {
	return xxx_messageInfo_FieldMetadata.Size(m)
}
func (m *FieldMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMetadata proto.InternalMessageInfo

func (m *FieldMetadata) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *FieldMetadata) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

type isFieldMetadata_Type interface {
	isFieldMetadata_Type()
}

type FieldMetadata_StringField struct {
	StringField *StringField `protobuf:"bytes,3,opt,name=string_field,json=stringField,proto3,oneof"`
}

type FieldMetadata_OptionField struct {
	OptionField *OptionField `protobuf:"bytes,4,opt,name=option_field,json=optionField,proto3,oneof"`
}

func (*FieldMetadata_StringField) isFieldMetadata_Type() {}

func (*FieldMetadata_OptionField) isFieldMetadata_Type() {}

func (m *FieldMetadata) GetType() isFieldMetadata_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *FieldMetadata) GetStringField() *StringField {
	if x, ok := m.GetType().(*FieldMetadata_StringField); ok {
		return x.StringField
	}
	return nil
}

func (m *FieldMetadata) GetOptionField() *OptionField {
	if x, ok := m.GetType().(*FieldMetadata_OptionField); ok {
		return x.OptionField
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldMetadata_StringField)(nil),
		(*FieldMetadata_OptionField)(nil),
	}
}

type SchemaMetadata struct {
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Searchable           bool     `protobuf:"varint,2,opt,name=searchable,proto3" json:"searchable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaMetadata) Reset()         { *m = SchemaMetadata{} }
func (m *SchemaMetadata) String() string { return proto.CompactTextString(m) }
func (*SchemaMetadata) ProtoMessage()    {}
func (*SchemaMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{5}
}

func (m *SchemaMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaMetadata.Unmarshal(m, b)
}
func (m *SchemaMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaMetadata.Marshal(b, m, deterministic)
}
func (m *SchemaMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaMetadata.Merge(m, src)
}
func (m *SchemaMetadata) XXX_Size() int {
	return xxx_messageInfo_SchemaMetadata.Size(m)
}
func (m *SchemaMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaMetadata proto.InternalMessageInfo

func (m *SchemaMetadata) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SchemaMetadata) GetSearchable() bool {
	if m != nil {
		return m.Searchable
	}
	return false
}

type Schema struct {
	// The type URL of the object the schema was produced from, which becomes the 'have' type URL when submitting a
	// filled-in schema.
	TypeUrl              string          `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Metadata             *SchemaMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Fields               []*Field        `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a802f2859632ed7, []int{6}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Schema) GetMetadata() *SchemaMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Schema) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*StringField)(nil), "clutch.resolver.v1.StringField")
	proto.RegisterType((*Option)(nil), "clutch.resolver.v1.Option")
	proto.RegisterType((*OptionField)(nil), "clutch.resolver.v1.OptionField")
	proto.RegisterType((*Field)(nil), "clutch.resolver.v1.Field")
	proto.RegisterType((*FieldMetadata)(nil), "clutch.resolver.v1.FieldMetadata")
	proto.RegisterType((*SchemaMetadata)(nil), "clutch.resolver.v1.SchemaMetadata")
	proto.RegisterType((*Schema)(nil), "clutch.resolver.v1.Schema")
}

func init() {
	proto.RegisterFile("resolver/v1/schema.proto", fileDescriptor_0a802f2859632ed7)
}

var fileDescriptor_0a802f2859632ed7 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x4d, 0xea, 0xb8, 0x63, 0x17, 0xa1, 0x95, 0x10, 0x6e, 0x0f, 0x90, 0xba, 0x97, 0x1c,
	0x90, 0xab, 0x14, 0xc4, 0x0d, 0x24, 0xa2, 0x08, 0x71, 0x01, 0x24, 0x07, 0x90, 0xe8, 0xc5, 0xda,
	0xda, 0x53, 0x62, 0x69, 0xfc, 0xc1, 0x7a, 0x6d, 0x29, 0x7f, 0x84, 0xff, 0xc0, 0xdf, 0xe2, 0x97,
	0x20, 0xef, 0xae, 0x13, 0x57, 0x4d, 0x25, 0x7a, 0xf3, 0xce, 0x7b, 0xfb, 0xf6, 0xcd, 0xbc, 0x31,
	0xf8, 0x02, 0xeb, 0x92, 0x5a, 0x14, 0x17, 0xed, 0xfc, 0xa2, 0x4e, 0xd6, 0x98, 0xf3, 0xb0, 0x12,
	0xa5, 0x2c, 0x19, 0x4b, 0xa8, 0x91, 0xc9, 0x3a, 0xec, 0x09, 0x61, 0x3b, 0x0f, 0xbe, 0x82, 0xbb,
	0x92, 0x22, 0x2b, 0x7e, 0x7e, 0xc8, 0x90, 0x52, 0x36, 0x05, 0xb7, 0x22, 0x9e, 0xe0, 0xba, 0xa4,
	0x14, 0x85, 0x6f, 0x4d, 0xad, 0xd9, 0x51, 0x34, 0x2c, 0xb1, 0x73, 0x38, 0x4e, 0xf1, 0x86, 0x37,
	0x24, 0xe3, 0x96, 0x53, 0x83, 0xfe, 0x81, 0xe2, 0x78, 0xa6, 0xf8, 0xbd, 0xab, 0x05, 0x3f, 0xc0,
	0xfe, 0x52, 0xc9, 0xac, 0x2c, 0xd8, 0x19, 0x78, 0x69, 0x56, 0x57, 0xc4, 0x37, 0x71, 0xc1, 0x73,
	0xec, 0x15, 0x4d, 0xed, 0x33, 0xcf, 0x91, 0x9d, 0x83, 0x57, 0x2b, 0x0b, 0x43, 0xc1, 0x8f, 0x8f,
	0x22, 0x57, 0x57, 0x95, 0xe2, 0x62, 0x02, 0x87, 0x0a, 0x0d, 0xfe, 0x58, 0xe0, 0x6a, 0x6d, 0xed,
	0xf8, 0x25, 0xb0, 0xac, 0x48, 0xa8, 0x49, 0x31, 0xe6, 0x44, 0x71, 0xa9, 0x20, 0xf5, 0x8c, 0x13,
	0x3d, 0x31, 0xc8, 0x7b, 0x22, 0x63, 0xe7, 0x0d, 0x3c, 0xeb, 0xd9, 0xe9, 0xa6, 0xe0, 0x79, 0x96,
	0x98, 0x1b, 0xb5, 0x7f, 0x30, 0x1d, 0xcd, 0x8e, 0xa2, 0xa7, 0x06, 0x5e, 0x6a, 0x54, 0x5f, 0xab,
	0xd9, 0x6b, 0x98, 0xf4, 0xbc, 0xd1, 0x74, 0x34, 0x73, 0x2f, 0x4f, 0xc3, 0xbb, 0xc3, 0x0c, 0x35,
	0x3b, 0xea, 0xa9, 0xc1, 0x15, 0x1c, 0x6a, 0x93, 0x0c, 0xc6, 0x83, 0xee, 0xd5, 0x37, 0x7b, 0x0b,
	0x4e, 0x8e, 0x92, 0xa7, 0x5c, 0x72, 0xd5, 0xb2, 0x7b, 0x79, 0xb6, 0x4f, 0x53, 0x09, 0x7c, 0x32,
	0xc4, 0x68, 0x7b, 0x25, 0xf8, 0x6b, 0xc1, 0xf1, 0x2d, 0xec, 0x7f, 0x46, 0x7d, 0x0a, 0x8e, 0xc0,
	0x5f, 0x4d, 0x26, 0x30, 0x55, 0x6f, 0x3a, 0xd1, 0xf6, 0xcc, 0x96, 0xdb, 0x18, 0x6e, 0x3a, 0x59,
	0x7f, 0xa4, 0x3c, 0xbd, 0xd8, 0xe7, 0x69, 0xb0, 0x31, 0xbb, 0x9c, 0x74, 0xa7, 0x4b, 0xf0, 0x74,
	0xf7, 0x46, 0x65, 0x7c, 0xbf, 0xca, 0x20, 0xc5, 0x4e, 0xa5, 0xdc, 0x1d, 0x17, 0x36, 0x8c, 0xe5,
	0xa6, 0xc2, 0x60, 0x05, 0x8f, 0x57, 0x6a, 0x83, 0x1f, 0xd2, 0xe4, 0x73, 0x80, 0x1a, 0xb9, 0x48,
	0xd6, 0xfc, 0x9a, 0xd0, 0xb4, 0x39, 0xa8, 0x04, 0xbf, 0x2d, 0xb0, 0xb5, 0x2a, 0x3b, 0x01, 0xa7,
	0x7b, 0x27, 0x6e, 0x04, 0x19, 0xa5, 0x49, 0x77, 0xfe, 0x26, 0x88, 0xbd, 0xbb, 0x13, 0x4f, 0xb0,
	0x77, 0x14, 0xb7, 0xec, 0xed, 0xf2, 0x61, 0x73, 0xb0, 0xd5, 0x04, 0xfa, 0x85, 0x39, 0xb9, 0x37,
	0xdc, 0xc8, 0x10, 0x17, 0xde, 0x15, 0xf4, 0x60, 0x3b, 0xbf, 0xb6, 0xd5, 0x4f, 0xfb, 0xea, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x75, 0xbc, 0x9d, 0xd0, 0x03, 0x00, 0x00,
}