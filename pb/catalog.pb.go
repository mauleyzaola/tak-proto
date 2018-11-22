// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

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

type Catalog struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,100,opt,name=name,proto3" json:"name,omitempty"`
	Date1                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date1,proto3" json:"date1,omitempty"`
	Date2                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date2,proto3" json:"date2,omitempty"`
	Date3                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date3,proto3" json:"date3,omitempty"`
	Date4                *timestamp.Timestamp `protobuf:"bytes,5,opt,name=date4,proto3" json:"date4,omitempty"`
	String1              string               `protobuf:"bytes,6,opt,name=string1,proto3" json:"string1,omitempty"`
	String2              string               `protobuf:"bytes,7,opt,name=string2,proto3" json:"string2,omitempty"`
	String3              string               `protobuf:"bytes,8,opt,name=string3,proto3" json:"string3,omitempty"`
	String4              string               `protobuf:"bytes,9,opt,name=string4,proto3" json:"string4,omitempty"`
	Number1              float64              `protobuf:"fixed64,10,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float64              `protobuf:"fixed64,11,opt,name=number2,proto3" json:"number2,omitempty"`
	Number3              float64              `protobuf:"fixed64,12,opt,name=number3,proto3" json:"number3,omitempty"`
	Number4              float64              `protobuf:"fixed64,13,opt,name=number4,proto3" json:"number4,omitempty"`
	Int1                 int64                `protobuf:"varint,14,opt,name=int1,proto3" json:"int1,omitempty"`
	Int2                 int64                `protobuf:"varint,15,opt,name=int2,proto3" json:"int2,omitempty"`
	Int3                 int64                `protobuf:"varint,16,opt,name=int3,proto3" json:"int3,omitempty"`
	Int4                 int64                `protobuf:"varint,17,opt,name=int4,proto3" json:"int4,omitempty"`
	Bool1                bool                 `protobuf:"varint,18,opt,name=bool1,proto3" json:"bool1,omitempty"`
	Bool2                bool                 `protobuf:"varint,19,opt,name=bool2,proto3" json:"bool2,omitempty"`
	Bool3                bool                 `protobuf:"varint,20,opt,name=bool3,proto3" json:"bool3,omitempty"`
	Bool4                bool                 `protobuf:"varint,21,opt,name=bool4,proto3" json:"bool4,omitempty"`
	String5              string               `protobuf:"bytes,22,opt,name=string5,proto3" json:"string5,omitempty"`
	String6              string               `protobuf:"bytes,23,opt,name=string6,proto3" json:"string6,omitempty"`
	String7              string               `protobuf:"bytes,24,opt,name=string7,proto3" json:"string7,omitempty"`
	String8              string               `protobuf:"bytes,25,opt,name=string8,proto3" json:"string8,omitempty"`
	String9              string               `protobuf:"bytes,26,opt,name=string9,proto3" json:"string9,omitempty"`
	String10             string               `protobuf:"bytes,27,opt,name=string10,proto3" json:"string10,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Catalog) Reset()         { *m = Catalog{} }
func (m *Catalog) String() string { return proto.CompactTextString(m) }
func (*Catalog) ProtoMessage()    {}
func (*Catalog) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_949b2a4f4a36785a, []int{0}
}
func (m *Catalog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Catalog.Unmarshal(m, b)
}
func (m *Catalog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Catalog.Marshal(b, m, deterministic)
}
func (dst *Catalog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Catalog.Merge(dst, src)
}
func (m *Catalog) XXX_Size() int {
	return xxx_messageInfo_Catalog.Size(m)
}
func (m *Catalog) XXX_DiscardUnknown() {
	xxx_messageInfo_Catalog.DiscardUnknown(m)
}

var xxx_messageInfo_Catalog proto.InternalMessageInfo

func (m *Catalog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Catalog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Catalog) GetDate1() *timestamp.Timestamp {
	if m != nil {
		return m.Date1
	}
	return nil
}

func (m *Catalog) GetDate2() *timestamp.Timestamp {
	if m != nil {
		return m.Date2
	}
	return nil
}

func (m *Catalog) GetDate3() *timestamp.Timestamp {
	if m != nil {
		return m.Date3
	}
	return nil
}

func (m *Catalog) GetDate4() *timestamp.Timestamp {
	if m != nil {
		return m.Date4
	}
	return nil
}

func (m *Catalog) GetString1() string {
	if m != nil {
		return m.String1
	}
	return ""
}

func (m *Catalog) GetString2() string {
	if m != nil {
		return m.String2
	}
	return ""
}

func (m *Catalog) GetString3() string {
	if m != nil {
		return m.String3
	}
	return ""
}

func (m *Catalog) GetString4() string {
	if m != nil {
		return m.String4
	}
	return ""
}

func (m *Catalog) GetNumber1() float64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Catalog) GetNumber2() float64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

func (m *Catalog) GetNumber3() float64 {
	if m != nil {
		return m.Number3
	}
	return 0
}

func (m *Catalog) GetNumber4() float64 {
	if m != nil {
		return m.Number4
	}
	return 0
}

func (m *Catalog) GetInt1() int64 {
	if m != nil {
		return m.Int1
	}
	return 0
}

func (m *Catalog) GetInt2() int64 {
	if m != nil {
		return m.Int2
	}
	return 0
}

func (m *Catalog) GetInt3() int64 {
	if m != nil {
		return m.Int3
	}
	return 0
}

func (m *Catalog) GetInt4() int64 {
	if m != nil {
		return m.Int4
	}
	return 0
}

func (m *Catalog) GetBool1() bool {
	if m != nil {
		return m.Bool1
	}
	return false
}

func (m *Catalog) GetBool2() bool {
	if m != nil {
		return m.Bool2
	}
	return false
}

func (m *Catalog) GetBool3() bool {
	if m != nil {
		return m.Bool3
	}
	return false
}

func (m *Catalog) GetBool4() bool {
	if m != nil {
		return m.Bool4
	}
	return false
}

func (m *Catalog) GetString5() string {
	if m != nil {
		return m.String5
	}
	return ""
}

func (m *Catalog) GetString6() string {
	if m != nil {
		return m.String6
	}
	return ""
}

func (m *Catalog) GetString7() string {
	if m != nil {
		return m.String7
	}
	return ""
}

func (m *Catalog) GetString8() string {
	if m != nil {
		return m.String8
	}
	return ""
}

func (m *Catalog) GetString9() string {
	if m != nil {
		return m.String9
	}
	return ""
}

func (m *Catalog) GetString10() string {
	if m != nil {
		return m.String10
	}
	return ""
}

func init() {
	proto.RegisterType((*Catalog)(nil), "pb.Catalog")
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_catalog_949b2a4f4a36785a) }

var fileDescriptor_catalog_949b2a4f4a36785a = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x6e, 0xa3, 0x30,
	0x18, 0x85, 0x05, 0xb9, 0x3b, 0x93, 0xcc, 0x8c, 0x27, 0xd3, 0x9e, 0xd2, 0x45, 0x51, 0x57, 0xac,
	0x48, 0xb0, 0x9d, 0xdb, 0xba, 0x6f, 0x80, 0xfa, 0x02, 0x50, 0x28, 0x42, 0x0a, 0x10, 0x25, 0xce,
	0x53, 0xf4, 0xa5, 0xab, 0x1a, 0x8c, 0xbd, 0xcc, 0xce, 0xff, 0xf7, 0xf9, 0xe8, 0x08, 0xff, 0x90,
	0xc5, 0x47, 0x22, 0x93, 0x53, 0x53, 0x84, 0xe7, 0x4b, 0x23, 0x1b, 0xea, 0x9e, 0x53, 0xef, 0xa5,
	0x68, 0x9a, 0xe2, 0x94, 0xaf, 0x15, 0x49, 0x6f, 0x9f, 0x6b, 0x59, 0x56, 0xf9, 0x55, 0x26, 0xd5,
	0xb9, 0xbd, 0xf4, 0xfa, 0x35, 0x26, 0x93, 0xb7, 0x36, 0x46, 0x97, 0xc4, 0x2d, 0x33, 0x38, 0xbe,
	0x13, 0xcc, 0x62, 0xb7, 0xcc, 0x28, 0x25, 0xc3, 0x3a, 0xa9, 0x72, 0x64, 0x8a, 0xa8, 0x33, 0xdd,
	0x90, 0x51, 0x96, 0xc8, 0x3c, 0x82, 0xeb, 0x3b, 0xc1, 0x9c, 0x79, 0x61, 0x5b, 0x10, 0xea, 0x82,
	0xf0, 0x5d, 0x17, 0xc4, 0xed, 0x45, 0x9d, 0x60, 0x18, 0xdc, 0x97, 0x60, 0x3a, 0xc1, 0x31, 0xbc,
	0x2f, 0xc1, 0x75, 0x42, 0x60, 0x74, 0x5f, 0x42, 0x50, 0x90, 0xc9, 0x55, 0x5e, 0xca, 0xba, 0x88,
	0x30, 0x56, 0x9f, 0xa7, 0x47, 0x63, 0x18, 0x26, 0xb6, 0x61, 0xc6, 0x70, 0x4c, 0x6d, 0xc3, 0x8d,
	0x11, 0x98, 0xd9, 0x46, 0xf5, 0xd4, 0xb7, 0x2a, 0xcd, 0x2f, 0x11, 0x88, 0xef, 0x04, 0x4e, 0xac,
	0x47, 0x63, 0x18, 0xe6, 0xb6, 0x61, 0xc6, 0x70, 0xfc, 0xb2, 0x0d, 0x37, 0x46, 0x60, 0x61, 0x1b,
	0xf1, 0xb3, 0xab, 0xb2, 0x96, 0x11, 0x96, 0xbe, 0x13, 0x0c, 0x62, 0x75, 0xee, 0x18, 0xc3, 0xef,
	0x9e, 0xb1, 0x8e, 0x71, 0xfc, 0xe9, 0x19, 0xef, 0x98, 0xc0, 0xdf, 0x9e, 0x09, 0xba, 0x22, 0xa3,
	0xb4, 0x69, 0x4e, 0x11, 0xa8, 0xef, 0x04, 0xd3, 0xb8, 0x1d, 0x34, 0x65, 0xf8, 0x67, 0x28, 0xd3,
	0x94, 0x63, 0x65, 0x28, 0xd7, 0x54, 0xe0, 0xbf, 0xa1, 0xd6, 0xbb, 0x6f, 0xf1, 0x60, 0xbf, 0xd4,
	0xd6, 0x98, 0x1d, 0x1e, 0x6d, 0xb3, 0x33, 0x66, 0x0f, 0xd8, 0x66, 0x6f, 0xcc, 0x01, 0x4f, 0xb6,
	0x39, 0x18, 0x73, 0x84, 0x67, 0x9b, 0x23, 0xf5, 0xc8, 0xb4, 0x5b, 0xf5, 0x06, 0xcf, 0x4a, 0xf5,
	0x73, 0x3a, 0x56, 0x3f, 0x0c, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xe8, 0xd6, 0xe4, 0x4a,
	0x03, 0x00, 0x00,
}
