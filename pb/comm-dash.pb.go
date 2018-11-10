// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comm-dash.proto

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

type CommDash struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               int32                `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=startDate,proto3" json:"startDate,omitempty"`
	CancelledDate        *timestamp.Timestamp `protobuf:"bytes,6,opt,name=cancelledDate,proto3" json:"cancelledDate,omitempty"`
	EstimatedDate        *timestamp.Timestamp `protobuf:"bytes,7,opt,name=estimatedDate,proto3" json:"estimatedDate,omitempty"`
	Material             *LabMaterial         `protobuf:"bytes,8,opt,name=material,proto3" json:"material,omitempty"`
	Owner                *User                `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Price                float64              `protobuf:"fixed64,10,opt,name=price,proto3" json:"price,omitempty"`
	Cas                  string               `protobuf:"bytes,11,opt,name=cas,proto3" json:"cas,omitempty"`
	Fema                 string               `protobuf:"bytes,12,opt,name=fema,proto3" json:"fema,omitempty"`
	NextAction           string               `protobuf:"bytes,13,opt,name=nextAction,proto3" json:"nextAction,omitempty"`
	TrackId              string               `protobuf:"bytes,14,opt,name=trackId,proto3" json:"trackId,omitempty"`
	Carrier              string               `protobuf:"bytes,15,opt,name=carrier,proto3" json:"carrier,omitempty"`
	Status               string               `protobuf:"bytes,16,opt,name=status,proto3" json:"status,omitempty"`
	Reason               string               `protobuf:"bytes,17,opt,name=reason,proto3" json:"reason,omitempty"`
	UserCancelled        string               `protobuf:"bytes,18,opt,name=userCancelled,proto3" json:"userCancelled,omitempty"`
	Assignee             *User                `protobuf:"bytes,19,opt,name=assignee,proto3" json:"assignee,omitempty"`
	RegulatoryStatus     string               `protobuf:"bytes,20,opt,name=regulatoryStatus,proto3" json:"regulatoryStatus,omitempty"`
	Supplier             *Supplier            `protobuf:"bytes,21,opt,name=supplier,proto3" json:"supplier,omitempty"`
	PurchaseOrder        string               `protobuf:"bytes,22,opt,name=purchaseOrder,proto3" json:"purchaseOrder,omitempty"`
	Quantity             float64              `protobuf:"fixed64,23,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PotentialKg          float64              `protobuf:"fixed64,24,opt,name=potentialKg,proto3" json:"potentialKg,omitempty"`
	Currency             *Currency            `protobuf:"bytes,25,opt,name=currency,proto3" json:"currency,omitempty"`
	ClientSpec           string               `protobuf:"bytes,26,opt,name=clientSpec,proto3" json:"clientSpec,omitempty"`
	HasPrice             bool                 `protobuf:"varint,27,opt,name=hasPrice,proto3" json:"hasPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommDash) Reset()         { *m = CommDash{} }
func (m *CommDash) String() string { return proto.CompactTextString(m) }
func (*CommDash) ProtoMessage()    {}
func (*CommDash) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_20318a8ca0d870c0, []int{0}
}
func (m *CommDash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommDash.Unmarshal(m, b)
}
func (m *CommDash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommDash.Marshal(b, m, deterministic)
}
func (dst *CommDash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommDash.Merge(dst, src)
}
func (m *CommDash) XXX_Size() int {
	return xxx_messageInfo_CommDash.Size(m)
}
func (m *CommDash) XXX_DiscardUnknown() {
	xxx_messageInfo_CommDash.DiscardUnknown(m)
}

var xxx_messageInfo_CommDash proto.InternalMessageInfo

func (m *CommDash) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CommDash) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *CommDash) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CommDash) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *CommDash) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CommDash) GetCancelledDate() *timestamp.Timestamp {
	if m != nil {
		return m.CancelledDate
	}
	return nil
}

func (m *CommDash) GetEstimatedDate() *timestamp.Timestamp {
	if m != nil {
		return m.EstimatedDate
	}
	return nil
}

func (m *CommDash) GetMaterial() *LabMaterial {
	if m != nil {
		return m.Material
	}
	return nil
}

func (m *CommDash) GetOwner() *User {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *CommDash) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CommDash) GetCas() string {
	if m != nil {
		return m.Cas
	}
	return ""
}

func (m *CommDash) GetFema() string {
	if m != nil {
		return m.Fema
	}
	return ""
}

func (m *CommDash) GetNextAction() string {
	if m != nil {
		return m.NextAction
	}
	return ""
}

func (m *CommDash) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func (m *CommDash) GetCarrier() string {
	if m != nil {
		return m.Carrier
	}
	return ""
}

func (m *CommDash) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CommDash) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *CommDash) GetUserCancelled() string {
	if m != nil {
		return m.UserCancelled
	}
	return ""
}

func (m *CommDash) GetAssignee() *User {
	if m != nil {
		return m.Assignee
	}
	return nil
}

func (m *CommDash) GetRegulatoryStatus() string {
	if m != nil {
		return m.RegulatoryStatus
	}
	return ""
}

func (m *CommDash) GetSupplier() *Supplier {
	if m != nil {
		return m.Supplier
	}
	return nil
}

func (m *CommDash) GetPurchaseOrder() string {
	if m != nil {
		return m.PurchaseOrder
	}
	return ""
}

func (m *CommDash) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *CommDash) GetPotentialKg() float64 {
	if m != nil {
		return m.PotentialKg
	}
	return 0
}

func (m *CommDash) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *CommDash) GetClientSpec() string {
	if m != nil {
		return m.ClientSpec
	}
	return ""
}

func (m *CommDash) GetHasPrice() bool {
	if m != nil {
		return m.HasPrice
	}
	return false
}

func init() {
	proto.RegisterType((*CommDash)(nil), "pb.CommDash")
}

func init() { proto.RegisterFile("comm-dash.proto", fileDescriptor_comm_dash_20318a8ca0d870c0) }

var fileDescriptor_comm_dash_20318a8ca0d870c0 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xd3, 0xa6, 0x71, 0xb7, 0x7f, 0x61, 0x29, 0x65, 0x08, 0x52, 0xb1, 0x50, 0x0f, 0x16,
	0xa8, 0xae, 0x04, 0x1c, 0x38, 0x82, 0xd2, 0x0b, 0x02, 0x04, 0x72, 0xe0, 0x01, 0xc6, 0xeb, 0x69,
	0xb2, 0xc2, 0x7f, 0xec, 0xae, 0x05, 0x79, 0x2e, 0x5e, 0x10, 0xed, 0xae, 0xed, 0x34, 0x70, 0xc8,
	0x6d, 0xbf, 0x9f, 0x19, 0x7d, 0x9a, 0x99, 0x65, 0x67, 0xa2, 0x2e, 0xcb, 0xeb, 0x1c, 0xf5, 0x2a,
	0x69, 0x54, 0x6d, 0x6a, 0x3e, 0x6a, 0xb2, 0xd9, 0xb3, 0x65, 0x5d, 0x2f, 0x0b, 0xba, 0x71, 0x4c,
	0xd6, 0xde, 0xdd, 0x18, 0x59, 0x92, 0x36, 0x58, 0x36, 0xde, 0x34, 0x3b, 0x15, 0xad, 0x52, 0x54,
	0x89, 0x75, 0x87, 0x79, 0x81, 0xd9, 0x75, 0x89, 0x86, 0x94, 0xc4, 0xa2, 0xf7, 0xe8, 0xb6, 0x69,
	0x0a, 0x49, 0xaa, 0xc3, 0xac, 0xd5, 0xfd, 0xfb, 0xf9, 0x9f, 0x09, 0x0b, 0xe7, 0x75, 0x59, 0xde,
	0xa2, 0x5e, 0xf1, 0x53, 0x36, 0x92, 0x39, 0x04, 0x51, 0x10, 0x1f, 0xa6, 0x23, 0x99, 0xf3, 0x0b,
	0x76, 0x50, 0xb5, 0x65, 0x46, 0x0a, 0x46, 0x51, 0x10, 0x8f, 0xd3, 0x0e, 0x71, 0xce, 0xf6, 0xcd,
	0xba, 0x21, 0xd8, 0x73, 0x4e, 0xf7, 0xe6, 0x6f, 0xd8, 0x44, 0x28, 0x42, 0x43, 0x39, 0xec, 0x47,
	0x41, 0x7c, 0xf4, 0x6a, 0x96, 0xf8, 0xec, 0x49, 0x9f, 0x3d, 0xf9, 0xd6, 0x67, 0x4f, 0x7b, 0x2b,
	0x7f, 0xcb, 0x0e, 0xb5, 0x41, 0x65, 0x6e, 0xd1, 0x10, 0x8c, 0x77, 0xd6, 0x6d, 0xcc, 0xfc, 0x1d,
	0x3b, 0x11, 0x58, 0x09, 0x2a, 0x0a, 0xca, 0x5d, 0xf5, 0xc1, 0xce, 0xea, 0xed, 0x02, 0xdb, 0x81,
	0xb4, 0x91, 0x76, 0x58, 0xbe, 0xc3, 0x64, 0x77, 0x87, 0xad, 0x02, 0xfe, 0x92, 0x85, 0xfd, 0xa8,
	0x21, 0x74, 0xc5, 0x67, 0x49, 0x93, 0x25, 0x9f, 0x30, 0xfb, 0xdc, 0xd1, 0xe9, 0x60, 0xe0, 0x97,
	0x6c, 0x5c, 0xff, 0xaa, 0x48, 0xc1, 0xa1, 0x73, 0x86, 0xd6, 0xf9, 0x5d, 0x93, 0x4a, 0x3d, 0xcd,
	0xcf, 0xd9, 0xb8, 0x51, 0x52, 0x10, 0xb0, 0x28, 0x88, 0x83, 0xd4, 0x03, 0x3e, 0x65, 0x7b, 0x02,
	0x35, 0x1c, 0xb9, 0x49, 0xdb, 0xa7, 0x1d, 0xfe, 0x1d, 0x95, 0x08, 0xc7, 0x7e, 0xf8, 0xf6, 0xcd,
	0x2f, 0x19, 0xab, 0xe8, 0xb7, 0x79, 0x2f, 0x8c, 0xac, 0x2b, 0x38, 0x71, 0xca, 0x3d, 0x86, 0x03,
	0x9b, 0x18, 0x85, 0xe2, 0xc7, 0x87, 0x1c, 0x4e, 0x9d, 0xd8, 0x43, 0xab, 0x08, 0x54, 0x4a, 0x92,
	0x82, 0x33, 0xaf, 0x74, 0xd0, 0x2e, 0x5f, 0x1b, 0x34, 0xad, 0x86, 0xa9, 0x13, 0x3a, 0x64, 0x79,
	0x45, 0xa8, 0xeb, 0x0a, 0x1e, 0x78, 0xde, 0x23, 0x7e, 0xc5, 0x4e, 0xec, 0x5d, 0xcd, 0xfb, 0x19,
	0x03, 0x77, 0xf2, 0x36, 0xc9, 0xaf, 0x58, 0x88, 0x5a, 0xcb, 0x65, 0x45, 0x04, 0x0f, 0xff, 0x19,
	0xc4, 0xa0, 0xf0, 0x17, 0x6c, 0xaa, 0x68, 0xd9, 0x16, 0x68, 0x6a, 0xb5, 0x5e, 0xf8, 0x14, 0xe7,
	0xae, 0xdd, 0x7f, 0x3c, 0x8f, 0x59, 0xd8, 0xdf, 0x37, 0x3c, 0x72, 0x1d, 0x8f, 0x6d, 0xc7, 0x45,
	0xc7, 0xa5, 0x83, 0x6a, 0x13, 0x36, 0xad, 0x12, 0x2b, 0xd4, 0xf4, 0x45, 0xe5, 0xa4, 0xe0, 0xc2,
	0x27, 0xdc, 0x22, 0xf9, 0x8c, 0x85, 0x3f, 0x5b, 0xac, 0x8c, 0x34, 0x6b, 0x78, 0xec, 0x56, 0x31,
	0x60, 0x1e, 0xb1, 0xa3, 0xa6, 0x36, 0x54, 0x19, 0x89, 0xc5, 0xc7, 0x25, 0x80, 0x93, 0xef, 0x53,
	0x36, 0x4d, 0xff, 0x23, 0xe1, 0xc9, 0x26, 0xcd, 0xbc, 0xe3, 0xd2, 0x41, 0xb5, 0x3b, 0x13, 0x85,
	0xa4, 0xca, 0x2c, 0x1a, 0x12, 0x30, 0xf3, 0x3b, 0xdb, 0x30, 0x36, 0xc7, 0x0a, 0xf5, 0x57, 0x77,
	0x12, 0x4f, 0xa3, 0x20, 0x0e, 0xd3, 0x01, 0x67, 0x07, 0xee, 0x36, 0x5f, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x43, 0x33, 0xba, 0x34, 0x04, 0x00, 0x00,
}
