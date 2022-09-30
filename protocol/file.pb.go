// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/file.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size_                int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa384d2da731e7da, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize_() int32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func init() {
	proto.RegisterType((*File)(nil), "protocol.File")
}

func init() { proto.RegisterFile("protocol/file.proto", fileDescriptor_fa384d2da731e7da) }

var fileDescriptor_fa384d2da731e7da = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0xcb, 0xcc, 0x49, 0xd5, 0x03, 0xf3, 0x84, 0x38, 0x60, 0x82,
	0x4a, 0x7a, 0x5c, 0x2c, 0x6e, 0x99, 0x39, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48, 0xac, 0x38, 0xb3, 0x2a, 0x55, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xcc, 0x76, 0xe2, 0x8a, 0x82, 0xeb, 0x4d, 0x62, 0x03, 0xb3,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0xe2, 0x22, 0x50, 0x63, 0x00, 0x00, 0x00,
}