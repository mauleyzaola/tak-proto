// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory-record.proto

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

type InventoryRecord struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Origin               string               `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Su                   string               `protobuf:"bytes,4,opt,name=su,proto3" json:"su,omitempty"`
	Quantity             float64              `protobuf:"fixed64,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	LabBatch             *LabBatch            `protobuf:"bytes,7,opt,name=labBatch,proto3" json:"labBatch,omitempty"`
	User                 *User                `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InventoryRecord) Reset()         { *m = InventoryRecord{} }
func (m *InventoryRecord) String() string { return proto.CompactTextString(m) }
func (*InventoryRecord) ProtoMessage()    {}
func (*InventoryRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_inventory_record_ec3cf5167456ac98, []int{0}
}
func (m *InventoryRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryRecord.Unmarshal(m, b)
}
func (m *InventoryRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryRecord.Marshal(b, m, deterministic)
}
func (dst *InventoryRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryRecord.Merge(dst, src)
}
func (m *InventoryRecord) XXX_Size() int {
	return xxx_messageInfo_InventoryRecord.Size(m)
}
func (m *InventoryRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryRecord.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryRecord proto.InternalMessageInfo

func (m *InventoryRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InventoryRecord) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *InventoryRecord) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *InventoryRecord) GetSu() string {
	if m != nil {
		return m.Su
	}
	return ""
}

func (m *InventoryRecord) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *InventoryRecord) GetLabBatch() *LabBatch {
	if m != nil {
		return m.LabBatch
	}
	return nil
}

func (m *InventoryRecord) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*InventoryRecord)(nil), "pb.InventoryRecord")
}

func init() {
	proto.RegisterFile("inventory-record.proto", fileDescriptor_inventory_record_ec3cf5167456ac98)
}

var fileDescriptor_inventory_record_ec3cf5167456ac98 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x49, 0xad, 0x6b, 0x1d, 0xc5, 0x85, 0x1c, 0x96, 0x50, 0x04, 0x8b, 0xa7, 0x5e, 0x36,
	0x0b, 0xfa, 0x06, 0xde, 0x04, 0x4f, 0x41, 0x1f, 0x20, 0xd9, 0xc6, 0x1a, 0xe8, 0x36, 0x35, 0x99,
	0x08, 0xfb, 0xb0, 0xbe, 0xcb, 0xd2, 0xe9, 0x9f, 0x5b, 0xe6, 0xcb, 0xc7, 0xef, 0x37, 0x03, 0x3b,
	0xd7, 0xff, 0xd9, 0x1e, 0x7d, 0x38, 0xef, 0x83, 0x3d, 0xfa, 0xd0, 0xc8, 0x21, 0x78, 0xf4, 0x3c,
	0x1b, 0x4c, 0xf9, 0xd4, 0x7a, 0xdf, 0x76, 0xf6, 0x40, 0x89, 0x49, 0xdf, 0x07, 0x74, 0x27, 0x1b,
	0x51, 0x9f, 0x86, 0xa9, 0x54, 0x6e, 0x3b, 0x6d, 0xf6, 0x46, 0xe3, 0xf1, 0x67, 0x0e, 0x20, 0x45,
	0x1b, 0xa6, 0xf7, 0xf3, 0x3f, 0x83, 0xed, 0xfb, 0x02, 0x57, 0xc4, 0xe6, 0x0f, 0x90, 0xb9, 0x46,
	0xb0, 0x8a, 0xd5, 0xb7, 0x2a, 0x73, 0x0d, 0x97, 0x90, 0x37, 0x1a, 0xad, 0xc8, 0x2a, 0x56, 0xdf,
	0xbd, 0x94, 0x72, 0x12, 0xca, 0x45, 0x28, 0x3f, 0x17, 0xa1, 0xa2, 0x1e, 0xdf, 0xc1, 0xc6, 0x07,
	0xd7, 0xba, 0x5e, 0x5c, 0x11, 0x63, 0x9e, 0x46, 0x6e, 0x4c, 0x22, 0x9f, 0xb8, 0x31, 0xf1, 0x12,
	0x8a, 0xdf, 0xa4, 0x7b, 0x74, 0x78, 0x16, 0xd7, 0x15, 0xab, 0x99, 0x5a, 0x67, 0x5e, 0x43, 0xd1,
	0x69, 0xf3, 0x36, 0x6e, 0x2d, 0x6e, 0xc8, 0x7b, 0x2f, 0x07, 0x23, 0x3f, 0xe6, 0x4c, 0xad, 0xbf,
	0xfc, 0x11, 0xf2, 0xf1, 0x1e, 0x51, 0x50, 0xab, 0x18, 0x5b, 0x5f, 0xd1, 0x06, 0x45, 0xa9, 0xd9,
	0xd0, 0x96, 0xaf, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x7a, 0xff, 0x52, 0x42, 0x01, 0x00,
	0x00,
}
