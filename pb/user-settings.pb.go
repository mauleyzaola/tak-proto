// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-settings.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserSettings struct {
	NotifyAllMaterials   bool     `protobuf:"varint,1,opt,name=notifyAllMaterials,proto3" json:"notifyAllMaterials,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSettings) Reset()         { *m = UserSettings{} }
func (m *UserSettings) String() string { return proto.CompactTextString(m) }
func (*UserSettings) ProtoMessage()    {}
func (*UserSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_settings_57f7e9a5fd8b7b29, []int{0}
}
func (m *UserSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettings.Unmarshal(m, b)
}
func (m *UserSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettings.Marshal(b, m, deterministic)
}
func (dst *UserSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettings.Merge(dst, src)
}
func (m *UserSettings) XXX_Size() int {
	return xxx_messageInfo_UserSettings.Size(m)
}
func (m *UserSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettings.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettings proto.InternalMessageInfo

func (m *UserSettings) GetNotifyAllMaterials() bool {
	if m != nil {
		return m.NotifyAllMaterials
	}
	return false
}

func init() {
	proto.RegisterType((*UserSettings)(nil), "pb.UserSettings")
}

func init() { proto.RegisterFile("user-settings.proto", fileDescriptor_user_settings_57f7e9a5fd8b7b29) }

var fileDescriptor_user_settings_57f7e9a5fd8b7b29 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2d, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x52, 0xb2, 0xe3, 0xe2, 0x09, 0x2d, 0x4e, 0x2d, 0x0a, 0x86, 0xca, 0x08, 0xe9,
	0x71, 0x09, 0xe5, 0xe5, 0x97, 0x64, 0xa6, 0x55, 0x3a, 0xe6, 0xe4, 0xf8, 0x26, 0x96, 0xa4, 0x16,
	0x65, 0x26, 0xe6, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x61, 0x91, 0x49, 0x62, 0x03,
	0x1b, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0x02, 0x0c, 0x02, 0x61, 0x00, 0x00, 0x00,
}
