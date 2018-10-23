// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lab-material.proto

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

type LabMaterial struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// double stock = 3;
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	LastSupplier         *Supplier            `protobuf:"bytes,5,opt,name=lastSupplier,proto3" json:"lastSupplier,omitempty"`
	LastExpires          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=lastExpires,proto3" json:"lastExpires,omitempty"`
	LastLocation         *Location            `protobuf:"bytes,7,opt,name=lastLocation,proto3" json:"lastLocation,omitempty"`
	Bases                string               `protobuf:"bytes,8,opt,name=bases,proto3" json:"bases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LabMaterial) Reset()         { *m = LabMaterial{} }
func (m *LabMaterial) String() string { return proto.CompactTextString(m) }
func (*LabMaterial) ProtoMessage()    {}
func (*LabMaterial) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_material_c4416058c2371816, []int{0}
}
func (m *LabMaterial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabMaterial.Unmarshal(m, b)
}
func (m *LabMaterial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabMaterial.Marshal(b, m, deterministic)
}
func (dst *LabMaterial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabMaterial.Merge(dst, src)
}
func (m *LabMaterial) XXX_Size() int {
	return xxx_messageInfo_LabMaterial.Size(m)
}
func (m *LabMaterial) XXX_DiscardUnknown() {
	xxx_messageInfo_LabMaterial.DiscardUnknown(m)
}

var xxx_messageInfo_LabMaterial proto.InternalMessageInfo

func (m *LabMaterial) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LabMaterial) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LabMaterial) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabMaterial) GetLastSupplier() *Supplier {
	if m != nil {
		return m.LastSupplier
	}
	return nil
}

func (m *LabMaterial) GetLastExpires() *timestamp.Timestamp {
	if m != nil {
		return m.LastExpires
	}
	return nil
}

func (m *LabMaterial) GetLastLocation() *Location {
	if m != nil {
		return m.LastLocation
	}
	return nil
}

func (m *LabMaterial) GetBases() string {
	if m != nil {
		return m.Bases
	}
	return ""
}

func init() {
	proto.RegisterType((*LabMaterial)(nil), "pb.LabMaterial")
}

func init() { proto.RegisterFile("lab-material.proto", fileDescriptor_lab_material_c4416058c2371816) }

var fileDescriptor_lab_material_c4416058c2371816 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4f, 0x03, 0x21,
	0x14, 0x84, 0xb3, 0x9b, 0xb6, 0x2a, 0xdb, 0xf4, 0xf0, 0xe2, 0x81, 0xec, 0xc5, 0xc6, 0x53, 0x2f,
	0x52, 0xa3, 0x57, 0xaf, 0xde, 0xea, 0x05, 0xfd, 0x03, 0x8f, 0xee, 0xb3, 0x21, 0x81, 0x42, 0x80,
	0x26, 0xfe, 0x79, 0x13, 0xb3, 0xb0, 0xc4, 0xee, 0x6d, 0xe6, 0xcb, 0xc0, 0xbc, 0x61, 0x60, 0x50,
	0x3d, 0x59, 0x4c, 0x14, 0x34, 0x1a, 0xe1, 0x83, 0x4b, 0x0e, 0x5a, 0xaf, 0xfa, 0x87, 0x93, 0x73,
	0x27, 0x43, 0xfb, 0x4c, 0xd4, 0xe5, 0x7b, 0x9f, 0xb4, 0xa5, 0x98, 0xd0, 0xfa, 0x12, 0xea, 0x37,
	0xf1, 0xe2, 0xbd, 0xd1, 0x14, 0xaa, 0x37, 0xee, 0x88, 0x49, 0xbb, 0x73, 0xf1, 0x8f, 0xbf, 0x0d,
	0xeb, 0x0e, 0xa8, 0x3e, 0xa6, 0xaf, 0x61, 0xc3, 0x5a, 0x3d, 0xf0, 0x66, 0xdb, 0xec, 0xee, 0x64,
	0xab, 0x07, 0x00, 0xb6, 0x38, 0xba, 0x81, 0x78, 0x9b, 0x49, 0xd6, 0x23, 0x3b, 0xa3, 0x25, 0xbe,
	0x28, 0x6c, 0xd4, 0xf0, 0xcc, 0xd6, 0x06, 0x63, 0xfa, 0x9c, 0xda, 0xf8, 0x72, 0xdb, 0xec, 0xba,
	0x97, 0xb5, 0xf0, 0x4a, 0x54, 0x26, 0x67, 0x09, 0x78, 0x63, 0xdd, 0xe8, 0xdf, 0x7f, 0xbc, 0x0e,
	0x14, 0xf9, 0x2a, 0x3f, 0xe8, 0x45, 0x19, 0x24, 0xea, 0x20, 0xf1, 0x55, 0x07, 0xc9, 0xeb, 0x78,
	0xed, 0x3b, 0x4c, 0x6b, 0xf8, 0xcd, 0x7f, 0x5f, 0x65, 0x72, 0x96, 0x80, 0x7b, 0xb6, 0x54, 0x18,
	0x29, 0xf2, 0xdb, 0x7c, 0x76, 0x31, 0x6a, 0x95, 0x8b, 0x5e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1c, 0x36, 0xb8, 0x16, 0x61, 0x01, 0x00, 0x00,
}
