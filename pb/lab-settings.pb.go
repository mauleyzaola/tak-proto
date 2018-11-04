// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lab-settings.proto

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

type LabSettings struct {
	BaseDomain           string   `protobuf:"bytes,1,opt,name=baseDomain,proto3" json:"baseDomain,omitempty"`
	DaysToExpire         int32    `protobuf:"varint,2,opt,name=daysToExpire,proto3" json:"daysToExpire,omitempty"`
	MailKey              string   `protobuf:"bytes,3,opt,name=mailKey,proto3" json:"mailKey,omitempty"`
	NotifyAreas          []string `protobuf:"bytes,4,rep,name=notifyAreas,proto3" json:"notifyAreas,omitempty"`
	SenderAddress        string   `protobuf:"bytes,5,opt,name=senderAddress,proto3" json:"senderAddress,omitempty"`
	GroupLaboratory      string   `protobuf:"bytes,7,opt,name=groupLaboratory,proto3" json:"groupLaboratory,omitempty"`
	GroupPurchasing      string   `protobuf:"bytes,8,opt,name=groupPurchasing,proto3" json:"groupPurchasing,omitempty"`
	GroupQA              string   `protobuf:"bytes,9,opt,name=groupQA,proto3" json:"groupQA,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabSettings) Reset()         { *m = LabSettings{} }
func (m *LabSettings) String() string { return proto.CompactTextString(m) }
func (*LabSettings) ProtoMessage()    {}
func (*LabSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_settings_02db60d7b19c2c82, []int{0}
}
func (m *LabSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabSettings.Unmarshal(m, b)
}
func (m *LabSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabSettings.Marshal(b, m, deterministic)
}
func (dst *LabSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabSettings.Merge(dst, src)
}
func (m *LabSettings) XXX_Size() int {
	return xxx_messageInfo_LabSettings.Size(m)
}
func (m *LabSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_LabSettings.DiscardUnknown(m)
}

var xxx_messageInfo_LabSettings proto.InternalMessageInfo

func (m *LabSettings) GetBaseDomain() string {
	if m != nil {
		return m.BaseDomain
	}
	return ""
}

func (m *LabSettings) GetDaysToExpire() int32 {
	if m != nil {
		return m.DaysToExpire
	}
	return 0
}

func (m *LabSettings) GetMailKey() string {
	if m != nil {
		return m.MailKey
	}
	return ""
}

func (m *LabSettings) GetNotifyAreas() []string {
	if m != nil {
		return m.NotifyAreas
	}
	return nil
}

func (m *LabSettings) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *LabSettings) GetGroupLaboratory() string {
	if m != nil {
		return m.GroupLaboratory
	}
	return ""
}

func (m *LabSettings) GetGroupPurchasing() string {
	if m != nil {
		return m.GroupPurchasing
	}
	return ""
}

func (m *LabSettings) GetGroupQA() string {
	if m != nil {
		return m.GroupQA
	}
	return ""
}

func init() {
	proto.RegisterType((*LabSettings)(nil), "pb.LabSettings")
}

func init() { proto.RegisterFile("lab-settings.proto", fileDescriptor_lab_settings_02db60d7b19c2c82) }

var fileDescriptor_lab_settings_02db60d7b19c2c82 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x60, 0x92, 0x5a, 0x6b, 0xa6, 0x8a, 0xb0, 0xa7, 0x39, 0x49, 0x28, 0x1e, 0x72, 0xd1,
	0x8b, 0x4f, 0x10, 0xd0, 0x93, 0x3d, 0x68, 0xf4, 0x05, 0x66, 0xcd, 0x1a, 0x17, 0xda, 0x9d, 0x30,
	0xb3, 0x05, 0xf7, 0x71, 0x7c, 0x53, 0x71, 0xad, 0x9a, 0x7a, 0x9c, 0x6f, 0x7f, 0x7e, 0x7e, 0x16,
	0xcc, 0x86, 0xec, 0x95, 0xba, 0x18, 0x7d, 0x18, 0xf4, 0x7a, 0x14, 0x8e, 0x6c, 0xca, 0xd1, 0xae,
	0x3e, 0x4a, 0x58, 0xae, 0xc9, 0x3e, 0xed, 0x5f, 0xcc, 0x05, 0x80, 0x25, 0x75, 0xb7, 0xbc, 0x25,
	0x1f, 0xb0, 0xa8, 0x8b, 0xa6, 0xea, 0x26, 0x62, 0x56, 0x70, 0xda, 0x53, 0xd2, 0x67, 0xbe, 0x7b,
	0x1f, 0xbd, 0x38, 0x2c, 0xeb, 0xa2, 0x99, 0x77, 0x07, 0x66, 0x10, 0x16, 0x5b, 0xf2, 0x9b, 0x7b,
	0x97, 0x70, 0x96, 0x0b, 0x7e, 0x4e, 0x53, 0xc3, 0x32, 0x70, 0xf4, 0xaf, 0xa9, 0x15, 0x47, 0x8a,
	0x47, 0xf5, 0xac, 0xa9, 0xba, 0x29, 0x99, 0x4b, 0x38, 0x53, 0x17, 0x7a, 0x27, 0x6d, 0xdf, 0x8b,
	0x53, 0xc5, 0x79, 0x6e, 0x38, 0x44, 0xd3, 0xc0, 0xf9, 0x20, 0xbc, 0x1b, 0xd7, 0x64, 0x59, 0x28,
	0xb2, 0x24, 0x5c, 0xe4, 0xdc, 0x7f, 0xfe, 0x4d, 0x3e, 0xec, 0xe4, 0xe5, 0x8d, 0xd4, 0x87, 0x01,
	0x4f, 0x26, 0xc9, 0x3f, 0xfe, 0x5a, 0x9d, 0xe9, 0xb1, 0xc5, 0xea, 0x7b, 0xf5, 0xfe, 0xb4, 0xc7,
	0xf9, 0xbb, 0x6e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xda, 0x45, 0x4d, 0x44, 0x01, 0x00,
	0x00,
}
