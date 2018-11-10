// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type File struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string               `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	User                 *User                `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Length               int32                `protobuf:"varint,5,opt,name=length,proto3" json:"length,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	MimeType             string               `protobuf:"bytes,7,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	Extension            string               `protobuf:"bytes,8,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_2119741b6756252c, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *File) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *File) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *File) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *File) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *File) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "pb.File")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_file_2119741b6756252c) }

var fileDescriptor_file_2119741b6756252c = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x98, 0xa6, 0xe9, 0x08, 0x5e, 0xcc, 0x85, 0x0c, 0xa1, 0x60, 0xf0, 0x2a, 0x57,
	0x5b, 0x50, 0x9f, 0xc1, 0x07, 0x08, 0xf5, 0x01, 0x12, 0x33, 0x8d, 0x0b, 0xd9, 0x1f, 0x36, 0x5b,
	0xd0, 0x37, 0xf6, 0x31, 0xa4, 0x93, 0xc6, 0xde, 0xcd, 0xf9, 0xf6, 0xec, 0xe1, 0x03, 0x38, 0x99,
	0x89, 0x75, 0x88, 0x3e, 0x79, 0xcc, 0x42, 0x5f, 0x3d, 0x8d, 0xde, 0x8f, 0x13, 0x1f, 0x84, 0xf4,
	0xe7, 0xd3, 0x21, 0x19, 0xcb, 0x73, 0xea, 0x6c, 0x58, 0x4a, 0x15, 0x9c, 0x67, 0x8e, 0xcb, 0xfd,
	0xfc, 0xab, 0x20, 0x7f, 0x37, 0x13, 0xe3, 0x03, 0x64, 0x66, 0x20, 0x55, 0xab, 0x66, 0xd7, 0x66,
	0x66, 0xc0, 0x47, 0x28, 0x42, 0x17, 0xd9, 0x25, 0xca, 0x84, 0x5d, 0x13, 0x22, 0xe4, 0xae, 0xb3,
	0x4c, 0x77, 0x42, 0xe5, 0xc6, 0x3d, 0xe4, 0x97, 0x49, 0xca, 0x6b, 0xd5, 0xdc, 0xbf, 0x94, 0x3a,
	0xf4, 0xfa, 0x63, 0xe6, 0xd8, 0x0a, 0xbd, 0x2c, 0x4d, 0xec, 0xc6, 0xf4, 0x45, 0x9b, 0x5a, 0x35,
	0x9b, 0xf6, 0x9a, 0xf0, 0x0d, 0xb6, 0x9f, 0x91, 0xbb, 0xc4, 0x03, 0x15, 0xf2, 0xb1, 0xd2, 0x8b,
	0xb9, 0x5e, 0xcd, 0xf5, 0x71, 0x35, 0x6f, 0xd7, 0x2a, 0x56, 0x50, 0x5a, 0x63, 0xf9, 0xf8, 0x13,
	0x98, 0xb6, 0xe2, 0xf0, 0x9f, 0x71, 0x0f, 0x3b, 0xfe, 0x4e, 0xec, 0x66, 0xe3, 0x1d, 0x95, 0xf2,
	0x78, 0x03, 0x7d, 0x21, 0xb3, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x70, 0x08, 0x77,
	0x30, 0x01, 0x00, 0x00,
}