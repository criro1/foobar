// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/keyref.proto

package schema

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

// KeyRefId  id part of key reference
type KeyRefId struct {
	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx,proto3" json:"idx,omitempty"`
	// referred key
	RefKey               []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey,proto3" json:"ref_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRefId) Reset()         { *m = KeyRefId{} }
func (m *KeyRefId) String() string { return proto.CompactTextString(m) }
func (*KeyRefId) ProtoMessage()    {}
func (*KeyRefId) Descriptor() ([]byte, []int) {
	return fileDescriptor_114d4d3fc965bf18, []int{0}
}

func (m *KeyRefId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRefId.Unmarshal(m, b)
}
func (m *KeyRefId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRefId.Marshal(b, m, deterministic)
}
func (m *KeyRefId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRefId.Merge(m, src)
}
func (m *KeyRefId) XXX_Size() int {
	return xxx_messageInfo_KeyRefId.Size(m)
}
func (m *KeyRefId) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRefId.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRefId proto.InternalMessageInfo

func (m *KeyRefId) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *KeyRefId) GetIdx() string {
	if m != nil {
		return m.Idx
	}
	return ""
}

func (m *KeyRefId) GetRefKey() []string {
	if m != nil {
		return m.RefKey
	}
	return nil
}

type KeyRef struct {
	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx,proto3" json:"idx,omitempty"`
	// referred key
	RefKey []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey,proto3" json:"ref_key,omitempty"`
	// primary key instance linked to
	PKey                 []string `protobuf:"bytes,4,rep,name=p_key,json=pKey,proto3" json:"p_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRef) Reset()         { *m = KeyRef{} }
func (m *KeyRef) String() string { return proto.CompactTextString(m) }
func (*KeyRef) ProtoMessage()    {}
func (*KeyRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_114d4d3fc965bf18, []int{1}
}

func (m *KeyRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRef.Unmarshal(m, b)
}
func (m *KeyRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRef.Marshal(b, m, deterministic)
}
func (m *KeyRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRef.Merge(m, src)
}
func (m *KeyRef) XXX_Size() int {
	return xxx_messageInfo_KeyRef.Size(m)
}
func (m *KeyRef) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRef.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRef proto.InternalMessageInfo

func (m *KeyRef) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *KeyRef) GetIdx() string {
	if m != nil {
		return m.Idx
	}
	return ""
}

func (m *KeyRef) GetRefKey() []string {
	if m != nil {
		return m.RefKey
	}
	return nil
}

func (m *KeyRef) GetPKey() []string {
	if m != nil {
		return m.PKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRefId)(nil), "cckit.state.schema.KeyRefId")
	proto.RegisterType((*KeyRef)(nil), "cckit.state.schema.KeyRef")
}

func init() {
	proto.RegisterFile("schema/keyref.proto", fileDescriptor_114d4d3fc965bf18)
}

var fileDescriptor_114d4d3fc965bf18 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x8f, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x86, 0x83, 0x20, 0xea, 0x4d, 0xa6, 0x24, 0xca, 0x48, 0x98, 0x9c, 0x5a, 0x13, 0x07, 0x07,
	0x37, 0x37, 0x43, 0x5c, 0x18, 0x5d, 0x14, 0xca, 0x55, 0x08, 0x12, 0x48, 0xa9, 0x89, 0x7d, 0x7b,
	0xc3, 0xd5, 0x47, 0x70, 0xba, 0xbb, 0x2f, 0x7f, 0xfe, 0xcb, 0x07, 0xd1, 0x28, 0x6b, 0xec, 0x0a,
	0xd1, 0xa2, 0xd5, 0xa8, 0xf8, 0xa0, 0x7b, 0xd3, 0x33, 0x26, 0x65, 0xdb, 0x18, 0x3e, 0x9a, 0xc2,
	0x20, 0x77, 0x81, 0xf4, 0x0a, 0xcb, 0x0c, 0x6d, 0x8e, 0xea, 0x52, 0xb1, 0x0d, 0x84, 0x8e, 0xc6,
	0x5e, 0xe2, 0xed, 0x56, 0xf9, 0xef, 0x62, 0x6b, 0xf0, 0x9b, 0xea, 0x13, 0xcf, 0x08, 0x4e, 0x2b,
	0xdb, 0xc2, 0x42, 0xa3, 0xba, 0xb7, 0x68, 0x63, 0x3f, 0xf1, 0xa7, 0xa8, 0x46, 0x95, 0xa1, 0x4d,
	0x1f, 0x10, 0xba, 0xba, 0x3f, 0x94, 0xb1, 0x08, 0xe6, 0x03, 0xe1, 0x80, 0x70, 0x30, 0x64, 0x68,
	0xcf, 0xfb, 0x1b, 0x7f, 0x36, 0xa6, 0x7e, 0x97, 0x5c, 0xf6, 0x9d, 0x18, 0x8f, 0x06, 0x65, 0xfd,
	0x2a, 0x4a, 0x41, 0x6e, 0x82, 0xdc, 0x84, 0x7b, 0x75, 0x72, 0xa3, 0x0c, 0xc9, 0xfe, 0xf0, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0x27, 0x7d, 0x39, 0x14, 0x01, 0x00, 0x00,
}