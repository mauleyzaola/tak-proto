// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplier.proto

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

type Supplier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_supplier_392fc94801cc1cf2, []int{0}
}
func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Supplier.Unmarshal(m, b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Supplier.Marshal(b, m, deterministic)
}
func (dst *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(dst, src)
}
func (m *Supplier) XXX_Size() int {
	return xxx_messageInfo_Supplier.Size(m)
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Supplier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Supplier)(nil), "pb.Supplier")
}

func init() { proto.RegisterFile("supplier.proto", fileDescriptor_supplier_392fc94801cc1cf2) }

var fileDescriptor_supplier_392fc94801cc1cf2 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2e, 0x2d, 0x28,
	0xc8, 0xc9, 0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2,
	0xe3, 0xe2, 0x08, 0x86, 0x8a, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81,
	0x45, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0x56, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x94,
	0xb0, 0xbd, 0x4c, 0x00, 0x00, 0x00,
}