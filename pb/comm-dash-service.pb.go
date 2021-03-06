// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comm-dash-service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommDashList struct {
	Rows                 []*CommDash `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CommDashList) Reset()         { *m = CommDashList{} }
func (m *CommDashList) String() string { return proto.CompactTextString(m) }
func (*CommDashList) ProtoMessage()    {}
func (*CommDashList) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_service_5f601c1dfd2cc331, []int{0}
}
func (m *CommDashList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommDashList.Unmarshal(m, b)
}
func (m *CommDashList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommDashList.Marshal(b, m, deterministic)
}
func (dst *CommDashList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommDashList.Merge(dst, src)
}
func (m *CommDashList) XXX_Size() int {
	return xxx_messageInfo_CommDashList.Size(m)
}
func (m *CommDashList) XXX_DiscardUnknown() {
	xxx_messageInfo_CommDashList.DiscardUnknown(m)
}

var xxx_messageInfo_CommDashList proto.InternalMessageInfo

func (m *CommDashList) GetRows() []*CommDash {
	if m != nil {
		return m.Rows
	}
	return nil
}

type CommDashParams struct {
	Suppliers            []string `protobuf:"bytes,1,rep,name=suppliers,proto3" json:"suppliers,omitempty"`
	Materials            []string `protobuf:"bytes,2,rep,name=materials,proto3" json:"materials,omitempty"`
	Status               []string `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommDashParams) Reset()         { *m = CommDashParams{} }
func (m *CommDashParams) String() string { return proto.CompactTextString(m) }
func (*CommDashParams) ProtoMessage()    {}
func (*CommDashParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_service_5f601c1dfd2cc331, []int{1}
}
func (m *CommDashParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommDashParams.Unmarshal(m, b)
}
func (m *CommDashParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommDashParams.Marshal(b, m, deterministic)
}
func (dst *CommDashParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommDashParams.Merge(dst, src)
}
func (m *CommDashParams) XXX_Size() int {
	return xxx_messageInfo_CommDashParams.Size(m)
}
func (m *CommDashParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CommDashParams.DiscardUnknown(m)
}

var xxx_messageInfo_CommDashParams proto.InternalMessageInfo

func (m *CommDashParams) GetSuppliers() []string {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

func (m *CommDashParams) GetMaterials() []string {
	if m != nil {
		return m.Materials
	}
	return nil
}

func (m *CommDashParams) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type GridCommDashRequest struct {
	Grid                 *GridInput      `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Params               *CommDashParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GridCommDashRequest) Reset()         { *m = GridCommDashRequest{} }
func (m *GridCommDashRequest) String() string { return proto.CompactTextString(m) }
func (*GridCommDashRequest) ProtoMessage()    {}
func (*GridCommDashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_service_5f601c1dfd2cc331, []int{2}
}
func (m *GridCommDashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridCommDashRequest.Unmarshal(m, b)
}
func (m *GridCommDashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridCommDashRequest.Marshal(b, m, deterministic)
}
func (dst *GridCommDashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridCommDashRequest.Merge(dst, src)
}
func (m *GridCommDashRequest) XXX_Size() int {
	return xxx_messageInfo_GridCommDashRequest.Size(m)
}
func (m *GridCommDashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GridCommDashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GridCommDashRequest proto.InternalMessageInfo

func (m *GridCommDashRequest) GetGrid() *GridInput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridCommDashRequest) GetParams() *CommDashParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type GridCommDash struct {
	Grid                 *GridOutput `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Rows                 []*CommDash `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GridCommDash) Reset()         { *m = GridCommDash{} }
func (m *GridCommDash) String() string { return proto.CompactTextString(m) }
func (*GridCommDash) ProtoMessage()    {}
func (*GridCommDash) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_service_5f601c1dfd2cc331, []int{3}
}
func (m *GridCommDash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridCommDash.Unmarshal(m, b)
}
func (m *GridCommDash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridCommDash.Marshal(b, m, deterministic)
}
func (dst *GridCommDash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridCommDash.Merge(dst, src)
}
func (m *GridCommDash) XXX_Size() int {
	return xxx_messageInfo_GridCommDash.Size(m)
}
func (m *GridCommDash) XXX_DiscardUnknown() {
	xxx_messageInfo_GridCommDash.DiscardUnknown(m)
}

var xxx_messageInfo_GridCommDash proto.InternalMessageInfo

func (m *GridCommDash) GetGrid() *GridOutput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridCommDash) GetRows() []*CommDash {
	if m != nil {
		return m.Rows
	}
	return nil
}

type CommDashCancelRequest struct {
	Dash                 *CommDash `protobuf:"bytes,1,opt,name=dash,proto3" json:"dash,omitempty"`
	User                 *User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Reason               string    `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CommDashCancelRequest) Reset()         { *m = CommDashCancelRequest{} }
func (m *CommDashCancelRequest) String() string { return proto.CompactTextString(m) }
func (*CommDashCancelRequest) ProtoMessage()    {}
func (*CommDashCancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_dash_service_5f601c1dfd2cc331, []int{4}
}
func (m *CommDashCancelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommDashCancelRequest.Unmarshal(m, b)
}
func (m *CommDashCancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommDashCancelRequest.Marshal(b, m, deterministic)
}
func (dst *CommDashCancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommDashCancelRequest.Merge(dst, src)
}
func (m *CommDashCancelRequest) XXX_Size() int {
	return xxx_messageInfo_CommDashCancelRequest.Size(m)
}
func (m *CommDashCancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommDashCancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommDashCancelRequest proto.InternalMessageInfo

func (m *CommDashCancelRequest) GetDash() *CommDash {
	if m != nil {
		return m.Dash
	}
	return nil
}

func (m *CommDashCancelRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CommDashCancelRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*CommDashList)(nil), "pb.CommDashList")
	proto.RegisterType((*CommDashParams)(nil), "pb.CommDashParams")
	proto.RegisterType((*GridCommDashRequest)(nil), "pb.GridCommDashRequest")
	proto.RegisterType((*GridCommDash)(nil), "pb.GridCommDash")
	proto.RegisterType((*CommDashCancelRequest)(nil), "pb.CommDashCancelRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommDashServiceClient is the client API for CommDashService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommDashServiceClient interface {
	Validate(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	Load(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error)
	StartProgress(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error)
	StopProgress(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error)
	Cancel(ctx context.Context, in *CommDashCancelRequest, opts ...grpc.CallOption) (*CommDash, error)
	Reopen(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error)
	Close(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error)
	NextNumber(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	Grid(ctx context.Context, in *GridCommDashRequest, opts ...grpc.CallOption) (*GridCommDash, error)
	DefaultStatus(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *CommDashParams, opts ...grpc.CallOption) (*CommDashList, error)
}

type commDashServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommDashServiceClient(cc *grpc.ClientConn) CommDashServiceClient {
	return &commDashServiceClient{cc}
}

func (c *commDashServiceClient) Validate(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Create(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Update(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Remove(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Load(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) StartProgress(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/StartProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) StopProgress(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/StopProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Cancel(ctx context.Context, in *CommDashCancelRequest, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Reopen(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Reopen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Close(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*CommDash, error) {
	out := new(CommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) NextNumber(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/NextNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) Grid(ctx context.Context, in *GridCommDashRequest, opts ...grpc.CallOption) (*GridCommDash, error) {
	out := new(GridCommDash)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/Grid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) DefaultStatus(ctx context.Context, in *CommDash, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/DefaultStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commDashServiceClient) List(ctx context.Context, in *CommDashParams, opts ...grpc.CallOption) (*CommDashList, error) {
	out := new(CommDashList)
	err := c.cc.Invoke(ctx, "/pb.CommDashService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommDashServiceServer is the server API for CommDashService service.
type CommDashServiceServer interface {
	Validate(context.Context, *CommDash) (*Response, error)
	Create(context.Context, *CommDash) (*Response, error)
	Update(context.Context, *CommDash) (*Response, error)
	Remove(context.Context, *CommDash) (*Response, error)
	Load(context.Context, *CommDash) (*CommDash, error)
	StartProgress(context.Context, *CommDash) (*CommDash, error)
	StopProgress(context.Context, *CommDash) (*CommDash, error)
	Cancel(context.Context, *CommDashCancelRequest) (*CommDash, error)
	Reopen(context.Context, *CommDash) (*CommDash, error)
	Close(context.Context, *CommDash) (*CommDash, error)
	NextNumber(context.Context, *CommDash) (*Response, error)
	Grid(context.Context, *GridCommDashRequest) (*GridCommDash, error)
	DefaultStatus(context.Context, *CommDash) (*Response, error)
	List(context.Context, *CommDashParams) (*CommDashList, error)
}

func RegisterCommDashServiceServer(s *grpc.Server, srv CommDashServiceServer) {
	s.RegisterService(&_CommDashService_serviceDesc, srv)
}

func _CommDashService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Validate(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Create(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Update(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Remove(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Load(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_StartProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).StartProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/StartProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).StartProgress(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_StopProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).StopProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/StopProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).StopProgress(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDashCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Cancel(ctx, req.(*CommDashCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Reopen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Reopen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Reopen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Reopen(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Close(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_NextNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).NextNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/NextNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).NextNumber(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_Grid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GridCommDashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).Grid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/Grid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).Grid(ctx, req.(*GridCommDashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_DefaultStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).DefaultStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/DefaultStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).DefaultStatus(ctx, req.(*CommDash))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommDashService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommDashParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommDashServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommDashService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommDashServiceServer).List(ctx, req.(*CommDashParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommDashService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CommDashService",
	HandlerType: (*CommDashServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _CommDashService_Validate_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CommDashService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CommDashService_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _CommDashService_Remove_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _CommDashService_Load_Handler,
		},
		{
			MethodName: "StartProgress",
			Handler:    _CommDashService_StartProgress_Handler,
		},
		{
			MethodName: "StopProgress",
			Handler:    _CommDashService_StopProgress_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _CommDashService_Cancel_Handler,
		},
		{
			MethodName: "Reopen",
			Handler:    _CommDashService_Reopen_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _CommDashService_Close_Handler,
		},
		{
			MethodName: "NextNumber",
			Handler:    _CommDashService_NextNumber_Handler,
		},
		{
			MethodName: "Grid",
			Handler:    _CommDashService_Grid_Handler,
		},
		{
			MethodName: "DefaultStatus",
			Handler:    _CommDashService_DefaultStatus_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CommDashService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comm-dash-service.proto",
}

func init() {
	proto.RegisterFile("comm-dash-service.proto", fileDescriptor_comm_dash_service_5f601c1dfd2cc331)
}

var fileDescriptor_comm_dash_service_5f601c1dfd2cc331 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x6f, 0x6f, 0xd3, 0x3e,
	0x10, 0xee, 0x9f, 0xfc, 0xa2, 0xed, 0xd6, 0x75, 0x3f, 0x19, 0xc1, 0x42, 0xb5, 0x17, 0x21, 0x02,
	0x54, 0x4d, 0xb4, 0x42, 0xed, 0x47, 0xe8, 0x24, 0x84, 0x34, 0x8d, 0x29, 0x65, 0xbc, 0x77, 0x9b,
	0x63, 0x8b, 0x94, 0xc4, 0xc6, 0xe7, 0x0c, 0x3e, 0x3b, 0xaf, 0x90, 0xed, 0xb8, 0x6b, 0xcb, 0x50,
	0xfa, 0xce, 0x77, 0xcf, 0xe3, 0xe7, 0xce, 0x77, 0xe7, 0x83, 0xf3, 0xb5, 0x28, 0xcb, 0x49, 0xc6,
	0xe9, 0x61, 0x42, 0xa8, 0x1e, 0xf3, 0x35, 0x4e, 0xa5, 0x12, 0x5a, 0xb0, 0x9e, 0x5c, 0x8d, 0xce,
	0x36, 0xa0, 0x73, 0x8e, 0xe0, 0x5e, 0xe5, 0x59, 0x73, 0x1e, 0x2a, 0x24, 0x29, 0x2a, 0x42, 0x8f,
	0xd5, 0x84, 0xca, 0x9d, 0x93, 0x8f, 0x30, 0x58, 0x88, 0xb2, 0xbc, 0xe2, 0xf4, 0x70, 0x9d, 0x93,
	0x66, 0x31, 0x04, 0x4a, 0xfc, 0xa4, 0xa8, 0x1b, 0xf7, 0xc7, 0x27, 0xb3, 0xc1, 0x54, 0xae, 0xa6,
	0x1e, 0x4f, 0x2d, 0x92, 0x64, 0x30, 0xf4, 0x9e, 0x5b, 0xae, 0x78, 0x49, 0xec, 0x02, 0x8e, 0xa9,
	0x96, 0xb2, 0xc8, 0x51, 0xb9, 0x8b, 0xc7, 0xe9, 0x93, 0xc3, 0xa0, 0x25, 0xd7, 0xa8, 0x72, 0x5e,
	0x50, 0xd4, 0x73, 0xe8, 0xc6, 0xc1, 0x5e, 0x41, 0x48, 0x9a, 0xeb, 0x9a, 0xa2, 0xbe, 0x85, 0x1a,
	0x2b, 0xc9, 0xe0, 0xc5, 0x27, 0x95, 0x67, 0x9b, 0xd8, 0xf8, 0xa3, 0x46, 0xd2, 0xec, 0x0d, 0x04,
	0xe6, 0x61, 0x51, 0x37, 0xee, 0x8e, 0x4f, 0x66, 0xa7, 0x26, 0x3d, 0x43, 0xfb, 0x5c, 0xc9, 0x5a,
	0xa7, 0x16, 0x62, 0x97, 0x10, 0x4a, 0x9b, 0x57, 0xd4, 0xb3, 0x24, 0xb6, 0xfd, 0x06, 0x97, 0x71,
	0xda, 0x30, 0x92, 0xaf, 0x30, 0xd8, 0x8e, 0xc2, 0x92, 0x1d, 0xf9, 0xa1, 0x97, 0xff, 0x52, 0xeb,
	0x27, 0x7d, 0x5f, 0xa1, 0xde, 0x3f, 0x2b, 0x24, 0xe0, 0xa5, 0xf7, 0x2c, 0x78, 0xb5, 0xc6, 0xc2,
	0x67, 0x1f, 0x43, 0x60, 0x5a, 0xd4, 0xc8, 0xef, 0x5d, 0x35, 0x08, 0xbb, 0x80, 0xc0, 0x34, 0xa7,
	0x49, 0xfd, 0xc8, 0x30, 0xee, 0x08, 0x55, 0x6a, 0xbd, 0xa6, 0x58, 0x0a, 0x39, 0x89, 0x2a, 0xea,
	0xc7, 0x5d, 0x53, 0x2c, 0x67, 0xcd, 0x7e, 0x07, 0x70, 0xe6, 0x85, 0x96, 0x6e, 0x36, 0xd8, 0x18,
	0x8e, 0xbe, 0xf1, 0x22, 0xcf, 0xb8, 0x46, 0xb6, 0x13, 0x69, 0x64, 0xad, 0xb4, 0x19, 0x89, 0xa4,
	0xc3, 0xde, 0x43, 0xb8, 0x50, 0x78, 0x10, 0xef, 0x4e, 0x1e, 0xa6, 0x97, 0x62, 0x29, 0x1e, 0xdb,
	0x78, 0x6f, 0x21, 0xb8, 0x16, 0x3c, 0x7b, 0x8e, 0xe5, 0xad, 0xa4, 0xc3, 0x26, 0x70, 0xba, 0xd4,
	0x5c, 0xe9, 0x5b, 0x25, 0xee, 0x15, 0x12, 0xb5, 0xd0, 0x3f, 0xc0, 0x60, 0xa9, 0x85, 0x3c, 0x90,
	0x3d, 0x87, 0xd0, 0x75, 0x88, 0xbd, 0xde, 0x46, 0x76, 0xba, 0xf6, 0xd7, 0x25, 0xfb, 0x3e, 0x21,
	0xb1, 0x6a, 0x11, 0x7f, 0x07, 0xff, 0x2d, 0x0a, 0x41, 0xd8, 0x42, 0xbb, 0x04, 0xb8, 0xc1, 0x5f,
	0xfa, 0xa6, 0x2e, 0x57, 0xa8, 0x5a, 0x4a, 0x36, 0x87, 0xc0, 0xcc, 0x23, 0x3b, 0xf7, 0x93, 0xb9,
	0xf7, 0x3f, 0x46, 0xff, 0xef, 0x03, 0xae, 0x82, 0x57, 0xf8, 0x9d, 0xd7, 0x85, 0x5e, 0xda, 0xbf,
	0xd5, 0x12, 0x63, 0x0a, 0x81, 0xdd, 0x04, 0xcf, 0xfc, 0x1b, 0x27, 0xbf, 0xbd, 0x2f, 0x92, 0xce,
	0x2a, 0xb4, 0x8b, 0x64, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xec, 0xe7, 0xcf, 0xa2, 0xa0, 0x04,
	0x00, 0x00,
}
