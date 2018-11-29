// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

type Message struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject              string               `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	User                 *User                `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Body                 string               `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Result               bool                 `protobuf:"varint,6,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_807098d448874aa8, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Message) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_807098d448874aa8) }

var fileDescriptor_message_807098d448874aa8 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xcf, 0x4a, 0xc6, 0x30,
	0x10, 0xc4, 0x49, 0x8c, 0xfd, 0x3e, 0x57, 0xf4, 0xb0, 0x07, 0x09, 0x45, 0xb0, 0x78, 0xea, 0x29,
	0x05, 0x7d, 0x0e, 0x2f, 0x41, 0x1f, 0xa0, 0x31, 0x6b, 0xa9, 0xb4, 0x24, 0xe4, 0xcf, 0xc1, 0x87,
	0xf2, 0x1d, 0xc5, 0xcd, 0xd7, 0xdb, 0xcc, 0xec, 0x30, 0xfb, 0x83, 0xbb, 0x9d, 0x72, 0x9e, 0x17,
	0x32, 0x31, 0x85, 0x12, 0x50, 0x46, 0xd7, 0x3f, 0x2d, 0x21, 0x2c, 0x1b, 0x4d, 0x9c, 0xb8, 0xfa,
	0x35, 0x95, 0x75, 0xa7, 0x5c, 0xe6, 0x3d, 0xb6, 0x52, 0x0f, 0x35, 0x53, 0x6a, 0xfa, 0xf9, 0x57,
	0xc0, 0xe9, 0xad, 0x4d, 0xe0, 0x3d, 0xc8, 0xd5, 0x6b, 0x31, 0x88, 0xf1, 0xc6, 0xca, 0xd5, 0xa3,
	0x86, 0x53, 0xae, 0xee, 0x9b, 0x3e, 0x8b, 0x96, 0x1c, 0x1e, 0x16, 0x1f, 0x41, 0xfd, 0x6f, 0xe8,
	0xab, 0x41, 0x8c, 0xb7, 0x2f, 0x67, 0x13, 0x9d, 0xf9, 0xc8, 0x94, 0x2c, 0xa7, 0x68, 0x40, 0xf9,
	0xb9, 0x90, 0x56, 0x7c, 0xed, 0x4d, 0xe3, 0x31, 0x07, 0x8f, 0x79, 0x3f, 0x78, 0x2c, 0xf7, 0x10,
	0x41, 0xb9, 0xe0, 0x7f, 0xf4, 0x35, 0x3f, 0x61, 0x8d, 0x0f, 0xd0, 0x25, 0xca, 0x75, 0x2b, 0xba,
	0x1b, 0xc4, 0x78, 0xb6, 0x17, 0xe7, 0x3a, 0x5e, 0x79, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x91, 0x23, 0xba, 0xf8, 0x00, 0x00, 0x00,
}