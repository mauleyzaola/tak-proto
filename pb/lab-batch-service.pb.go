// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lab-batch-service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type BatchCodeRequest struct {
	// repeated LabBatch batches = 1;
	Batch                *LabBatch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Next                 int32     `protobuf:"varint,3,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BatchCodeRequest) Reset()         { *m = BatchCodeRequest{} }
func (m *BatchCodeRequest) String() string { return proto.CompactTextString(m) }
func (*BatchCodeRequest) ProtoMessage()    {}
func (*BatchCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_c3b2101b0282a3f8, []int{0}
}
func (m *BatchCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchCodeRequest.Unmarshal(m, b)
}
func (m *BatchCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchCodeRequest.Marshal(b, m, deterministic)
}
func (dst *BatchCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchCodeRequest.Merge(dst, src)
}
func (m *BatchCodeRequest) XXX_Size() int {
	return xxx_messageInfo_BatchCodeRequest.Size(m)
}
func (m *BatchCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchCodeRequest proto.InternalMessageInfo

func (m *BatchCodeRequest) GetBatch() *LabBatch {
	if m != nil {
		return m.Batch
	}
	return nil
}

func (m *BatchCodeRequest) GetNext() int32 {
	if m != nil {
		return m.Next
	}
	return 0
}

type LabBatchList struct {
	Batches              []*LabBatch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LabBatchList) Reset()         { *m = LabBatchList{} }
func (m *LabBatchList) String() string { return proto.CompactTextString(m) }
func (*LabBatchList) ProtoMessage()    {}
func (*LabBatchList) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_c3b2101b0282a3f8, []int{1}
}
func (m *LabBatchList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabBatchList.Unmarshal(m, b)
}
func (m *LabBatchList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabBatchList.Marshal(b, m, deterministic)
}
func (dst *LabBatchList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabBatchList.Merge(dst, src)
}
func (m *LabBatchList) XXX_Size() int {
	return xxx_messageInfo_LabBatchList.Size(m)
}
func (m *LabBatchList) XXX_DiscardUnknown() {
	xxx_messageInfo_LabBatchList.DiscardUnknown(m)
}

var xxx_messageInfo_LabBatchList proto.InternalMessageInfo

func (m *LabBatchList) GetBatches() []*LabBatch {
	if m != nil {
		return m.Batches
	}
	return nil
}

type GridBatchRequest struct {
	Grid                 *GridInput   `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Params               *BatchParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GridBatchRequest) Reset()         { *m = GridBatchRequest{} }
func (m *GridBatchRequest) String() string { return proto.CompactTextString(m) }
func (*GridBatchRequest) ProtoMessage()    {}
func (*GridBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_c3b2101b0282a3f8, []int{2}
}
func (m *GridBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridBatchRequest.Unmarshal(m, b)
}
func (m *GridBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridBatchRequest.Marshal(b, m, deterministic)
}
func (dst *GridBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridBatchRequest.Merge(dst, src)
}
func (m *GridBatchRequest) XXX_Size() int {
	return xxx_messageInfo_GridBatchRequest.Size(m)
}
func (m *GridBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GridBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GridBatchRequest proto.InternalMessageInfo

func (m *GridBatchRequest) GetGrid() *GridInput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridBatchRequest) GetParams() *BatchParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type GridBatch struct {
	Grid                 *GridOutput `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Rows                 []*LabBatch `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GridBatch) Reset()         { *m = GridBatch{} }
func (m *GridBatch) String() string { return proto.CompactTextString(m) }
func (*GridBatch) ProtoMessage()    {}
func (*GridBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_c3b2101b0282a3f8, []int{3}
}
func (m *GridBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridBatch.Unmarshal(m, b)
}
func (m *GridBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridBatch.Marshal(b, m, deterministic)
}
func (dst *GridBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridBatch.Merge(dst, src)
}
func (m *GridBatch) XXX_Size() int {
	return xxx_messageInfo_GridBatch.Size(m)
}
func (m *GridBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_GridBatch.DiscardUnknown(m)
}

var xxx_messageInfo_GridBatch proto.InternalMessageInfo

func (m *GridBatch) GetGrid() *GridOutput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridBatch) GetRows() []*LabBatch {
	if m != nil {
		return m.Rows
	}
	return nil
}

type BatchParams struct {
	Codes            string               `protobuf:"bytes,1,opt,name=codes,proto3" json:"codes,omitempty"`
	Locations        []string             `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
	ExpirationStarts *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expirationStarts,proto3" json:"expirationStarts,omitempty"`
	ExpirationEnds   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expirationEnds,proto3" json:"expirationEnds,omitempty"`
	DateStarts       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=dateStarts,proto3" json:"dateStarts,omitempty"`
	DateEnds         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=dateEnds,proto3" json:"dateEnds,omitempty"`
	// double minStock = 7;
	// double maxStock = 8;
	Suppliers            []string             `protobuf:"bytes,9,rep,name=suppliers,proto3" json:"suppliers,omitempty"`
	ManufacturedStarts   *timestamp.Timestamp `protobuf:"bytes,10,opt,name=manufacturedStarts,proto3" json:"manufacturedStarts,omitempty"`
	ManufacturedEnds     *timestamp.Timestamp `protobuf:"bytes,11,opt,name=manufacturedEnds,proto3" json:"manufacturedEnds,omitempty"`
	SupplierBatch        string               `protobuf:"bytes,12,opt,name=supplierBatch,proto3" json:"supplierBatch,omitempty"`
	Batch                string               `protobuf:"bytes,13,opt,name=batch,proto3" json:"batch,omitempty"`
	MinStock             float64              `protobuf:"fixed64,14,opt,name=minStock,proto3" json:"minStock,omitempty"`
	MaxStock             float64              `protobuf:"fixed64,15,opt,name=maxStock,proto3" json:"maxStock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BatchParams) Reset()         { *m = BatchParams{} }
func (m *BatchParams) String() string { return proto.CompactTextString(m) }
func (*BatchParams) ProtoMessage()    {}
func (*BatchParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_c3b2101b0282a3f8, []int{4}
}
func (m *BatchParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchParams.Unmarshal(m, b)
}
func (m *BatchParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchParams.Marshal(b, m, deterministic)
}
func (dst *BatchParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchParams.Merge(dst, src)
}
func (m *BatchParams) XXX_Size() int {
	return xxx_messageInfo_BatchParams.Size(m)
}
func (m *BatchParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchParams.DiscardUnknown(m)
}

var xxx_messageInfo_BatchParams proto.InternalMessageInfo

func (m *BatchParams) GetCodes() string {
	if m != nil {
		return m.Codes
	}
	return ""
}

func (m *BatchParams) GetLocations() []string {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *BatchParams) GetExpirationStarts() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationStarts
	}
	return nil
}

func (m *BatchParams) GetExpirationEnds() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationEnds
	}
	return nil
}

func (m *BatchParams) GetDateStarts() *timestamp.Timestamp {
	if m != nil {
		return m.DateStarts
	}
	return nil
}

func (m *BatchParams) GetDateEnds() *timestamp.Timestamp {
	if m != nil {
		return m.DateEnds
	}
	return nil
}

func (m *BatchParams) GetSuppliers() []string {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

func (m *BatchParams) GetManufacturedStarts() *timestamp.Timestamp {
	if m != nil {
		return m.ManufacturedStarts
	}
	return nil
}

func (m *BatchParams) GetManufacturedEnds() *timestamp.Timestamp {
	if m != nil {
		return m.ManufacturedEnds
	}
	return nil
}

func (m *BatchParams) GetSupplierBatch() string {
	if m != nil {
		return m.SupplierBatch
	}
	return ""
}

func (m *BatchParams) GetBatch() string {
	if m != nil {
		return m.Batch
	}
	return ""
}

func (m *BatchParams) GetMinStock() float64 {
	if m != nil {
		return m.MinStock
	}
	return 0
}

func (m *BatchParams) GetMaxStock() float64 {
	if m != nil {
		return m.MaxStock
	}
	return 0
}

func init() {
	proto.RegisterType((*BatchCodeRequest)(nil), "pb.BatchCodeRequest")
	proto.RegisterType((*LabBatchList)(nil), "pb.LabBatchList")
	proto.RegisterType((*GridBatchRequest)(nil), "pb.GridBatchRequest")
	proto.RegisterType((*GridBatch)(nil), "pb.GridBatch")
	proto.RegisterType((*BatchParams)(nil), "pb.BatchParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LabBatchServiceClient is the client API for LabBatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabBatchServiceClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Load(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*LabBatch, error)
	Create(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*LabBatchList, error)
	Validate(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error)
	GenBatchCode(ctx context.Context, in *BatchCodeRequest, opts ...grpc.CallOption) (*LabBatch, error)
	Grid(ctx context.Context, in *GridBatchRequest, opts ...grpc.CallOption) (*GridBatch, error)
	Export(ctx context.Context, in *GridBatchRequest, opts ...grpc.CallOption) (*LabBatchList, error)
	NextBatchNumber(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error)
}

type labBatchServiceClient struct {
	cc *grpc.ClientConn
}

func NewLabBatchServiceClient(cc *grpc.ClientConn) LabBatchServiceClient {
	return &labBatchServiceClient{cc}
}

func (c *labBatchServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Load(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*LabBatch, error) {
	out := new(LabBatch)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Create(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Update(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) List(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*LabBatchList, error) {
	out := new(LabBatchList)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Validate(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) GenBatchCode(ctx context.Context, in *BatchCodeRequest, opts ...grpc.CallOption) (*LabBatch, error) {
	out := new(LabBatch)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/GenBatchCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Grid(ctx context.Context, in *GridBatchRequest, opts ...grpc.CallOption) (*GridBatch, error) {
	out := new(GridBatch)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Grid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) Export(ctx context.Context, in *GridBatchRequest, opts ...grpc.CallOption) (*LabBatchList, error) {
	out := new(LabBatchList)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labBatchServiceClient) NextBatchNumber(ctx context.Context, in *LabBatch, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.LabBatchService/NextBatchNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabBatchServiceServer is the server API for LabBatchService service.
type LabBatchServiceServer interface {
	Ping(context.Context, *Request) (*Response, error)
	Load(context.Context, *LabBatch) (*LabBatch, error)
	Create(context.Context, *LabBatch) (*Response, error)
	Update(context.Context, *LabBatch) (*Response, error)
	List(context.Context, *LabBatch) (*LabBatchList, error)
	Validate(context.Context, *LabBatch) (*Response, error)
	GenBatchCode(context.Context, *BatchCodeRequest) (*LabBatch, error)
	Grid(context.Context, *GridBatchRequest) (*GridBatch, error)
	Export(context.Context, *GridBatchRequest) (*LabBatchList, error)
	NextBatchNumber(context.Context, *LabBatch) (*Response, error)
}

func RegisterLabBatchServiceServer(s *grpc.Server, srv LabBatchServiceServer) {
	s.RegisterService(&_LabBatchService_serviceDesc, srv)
}

func _LabBatchService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Load(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Create(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Update(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).List(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Validate(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_GenBatchCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).GenBatchCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/GenBatchCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).GenBatchCode(ctx, req.(*BatchCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Grid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GridBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Grid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Grid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Grid(ctx, req.(*GridBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GridBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).Export(ctx, req.(*GridBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabBatchService_NextBatchNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabBatchServiceServer).NextBatchNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LabBatchService/NextBatchNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabBatchServiceServer).NextBatchNumber(ctx, req.(*LabBatch))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabBatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LabBatchService",
	HandlerType: (*LabBatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _LabBatchService_Ping_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _LabBatchService_Load_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LabBatchService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LabBatchService_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LabBatchService_List_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _LabBatchService_Validate_Handler,
		},
		{
			MethodName: "GenBatchCode",
			Handler:    _LabBatchService_GenBatchCode_Handler,
		},
		{
			MethodName: "Grid",
			Handler:    _LabBatchService_Grid_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _LabBatchService_Export_Handler,
		},
		{
			MethodName: "NextBatchNumber",
			Handler:    _LabBatchService_NextBatchNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lab-batch-service.proto",
}

func init() {
	proto.RegisterFile("lab-batch-service.proto", fileDescriptor_lab_batch_service_c3b2101b0282a3f8)
}

var fileDescriptor_lab_batch_service_c3b2101b0282a3f8 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd4, 0x3c,
	0x10, 0x6d, 0xda, 0x6c, 0xbe, 0xee, 0xec, 0xaf, 0xac, 0x4a, 0x5f, 0x14, 0x21, 0xb1, 0x84, 0xaa,
	0xec, 0x4d, 0x53, 0x69, 0x41, 0xbd, 0xe0, 0xb2, 0x55, 0xa9, 0xa8, 0xaa, 0x52, 0x52, 0xe0, 0x12,
	0xc9, 0x49, 0xdc, 0x25, 0x62, 0x13, 0x1b, 0xdb, 0x81, 0x7d, 0x46, 0x5e, 0x84, 0xd7, 0x40, 0xb6,
	0x93, 0x34, 0xe9, 0x2e, 0x84, 0xbb, 0xcc, 0x39, 0x67, 0x4e, 0xc6, 0xf6, 0xcc, 0xc0, 0xff, 0x2b,
	0x1c, 0x1d, 0x47, 0x58, 0xc6, 0x5f, 0x8e, 0x05, 0xe1, 0xdf, 0xd3, 0x98, 0x04, 0x8c, 0x53, 0x49,
	0xd1, 0x2e, 0x8b, 0xbc, 0x49, 0x4d, 0x1a, 0xd0, 0x1b, 0x71, 0xf2, 0xad, 0x20, 0x42, 0x96, 0xe1,
	0x98, 0x13, 0xc1, 0x68, 0x2e, 0xca, 0x1c, 0x0f, 0x96, 0x3c, 0x4d, 0xca, 0xef, 0xa7, 0x4b, 0x4a,
	0x97, 0x2b, 0x72, 0xa2, 0xa3, 0xa8, 0xb8, 0x3f, 0x91, 0x69, 0x46, 0x84, 0xc4, 0x19, 0x33, 0x02,
	0xff, 0x0a, 0xa6, 0x67, 0xca, 0xfa, 0x9c, 0x26, 0x24, 0x34, 0xb6, 0xc8, 0x87, 0x9e, 0xfe, 0x9d,
	0xbb, 0x3b, 0xb3, 0xe6, 0x83, 0xc5, 0x30, 0x60, 0x51, 0x70, 0x8d, 0x23, 0xad, 0x0b, 0x0d, 0x85,
	0x10, 0xd8, 0x39, 0x59, 0x4b, 0x77, 0x6f, 0x66, 0xcd, 0x7b, 0xa1, 0xfe, 0xf6, 0x4f, 0x61, 0x58,
	0xc9, 0xae, 0x53, 0x21, 0xd1, 0x11, 0xfc, 0xa7, 0xc5, 0x44, 0xb8, 0xd6, 0x6c, 0x6f, 0xc3, 0xa9,
	0x22, 0xfd, 0xcf, 0x30, 0xbd, 0xe4, 0x69, 0x62, 0xd0, 0xb2, 0x86, 0x67, 0x60, 0xab, 0x63, 0xb8,
	0x96, 0x2e, 0x61, 0xa4, 0x12, 0x95, 0xe6, 0x6d, 0xce, 0x0a, 0x19, 0x6a, 0x0a, 0xbd, 0x00, 0x87,
	0x61, 0x8e, 0x33, 0x51, 0xd6, 0x39, 0x51, 0x22, 0x6d, 0x72, 0xab, 0xe1, 0xb0, 0xa4, 0xfd, 0xf7,
	0xd0, 0xaf, 0xfd, 0x91, 0xdf, 0x32, 0x1e, 0x57, 0xc6, 0xef, 0x0a, 0xf9, 0xe0, 0x3c, 0x03, 0x9b,
	0xd3, 0x1f, 0xca, 0x77, 0xb3, 0x6a, 0xcd, 0xf8, 0xbf, 0x6c, 0x18, 0x34, 0x7e, 0x85, 0x0e, 0xa0,
	0x17, 0xd3, 0x44, 0x1f, 0xd4, 0x9a, 0xf7, 0x43, 0x13, 0xa0, 0x27, 0xd0, 0x5f, 0xd1, 0x18, 0xcb,
	0x94, 0xe6, 0xc6, 0xac, 0x1f, 0x3e, 0x00, 0xe8, 0x0d, 0x4c, 0xc9, 0x9a, 0xa5, 0x5c, 0x87, 0x77,
	0x12, 0x73, 0x29, 0xf4, 0x75, 0x0e, 0x16, 0x5e, 0x60, 0x9e, 0x2d, 0xa8, 0x9e, 0x2d, 0xf8, 0x50,
	0x3d, 0x5b, 0xb8, 0x91, 0x83, 0xce, 0x60, 0xfc, 0x80, 0x5d, 0xe4, 0x89, 0x70, 0xed, 0x4e, 0x97,
	0x47, 0x19, 0xe8, 0x35, 0x40, 0x82, 0x25, 0x29, 0xab, 0xe8, 0x75, 0xe6, 0x37, 0xd4, 0xe8, 0x14,
	0xf6, 0x55, 0xa4, 0xff, 0xec, 0x74, 0x66, 0xd6, 0x5a, 0x75, 0x3b, 0xa2, 0x60, 0x6c, 0x95, 0x12,
	0x2e, 0xdc, 0xbe, 0xb9, 0x9d, 0x1a, 0x40, 0x57, 0x80, 0x32, 0x9c, 0x17, 0xf7, 0x38, 0x96, 0x05,
	0x27, 0x49, 0x59, 0x19, 0x74, 0xfa, 0x6f, 0xc9, 0x52, 0x37, 0xdd, 0x44, 0x75, 0xa5, 0x83, 0xee,
	0x9b, 0x7e, 0x9c, 0x83, 0x0e, 0x61, 0x54, 0x15, 0xa8, 0x1f, 0xdf, 0x1d, 0xea, 0xd7, 0x6e, 0x83,
	0xaa, 0x17, 0xcc, 0xf8, 0x8c, 0x4c, 0x2f, 0x98, 0x81, 0xf1, 0x60, 0x3f, 0x4b, 0xf3, 0x3b, 0x49,
	0xe3, 0xaf, 0xee, 0x78, 0x66, 0xcd, 0xad, 0xb0, 0x8e, 0x35, 0x87, 0xd7, 0x86, 0x9b, 0x94, 0x5c,
	0x19, 0x2f, 0x7e, 0xee, 0xc1, 0xa4, 0x6a, 0xbe, 0x3b, 0xb3, 0x1b, 0xd0, 0x73, 0xb0, 0x6f, 0xd3,
	0x7c, 0x89, 0x06, 0xaa, 0x33, 0xcb, 0x89, 0xf1, 0x86, 0x26, 0x30, 0xab, 0xc0, 0xdf, 0x41, 0x87,
	0x60, 0x5f, 0x53, 0x9c, 0xa0, 0x56, 0xfb, 0x7a, 0xad, 0xc8, 0xdf, 0x41, 0x47, 0xe0, 0x9c, 0x73,
	0x82, 0x25, 0xd9, 0xa6, 0x6b, 0xb8, 0x1d, 0x81, 0xf3, 0x91, 0x25, 0xdd, 0xba, 0x39, 0xd8, 0x7a,
	0xf6, 0xdb, 0xaa, 0x69, 0x33, 0x52, 0xbc, 0x56, 0xee, 0x7f, 0xc2, 0xab, 0xf4, 0x1f, 0x3c, 0x5f,
	0xc1, 0xf0, 0x92, 0xe4, 0xf5, 0x9a, 0x42, 0x07, 0xf5, 0xa0, 0x37, 0xb6, 0xd6, 0xc6, 0xc9, 0x8e,
	0xc1, 0x56, 0x83, 0x6d, 0xd4, 0x8f, 0xf7, 0x8b, 0x37, 0x6a, 0xa1, 0xfe, 0x0e, 0x5a, 0x80, 0x73,
	0xb1, 0x66, 0x94, 0xcb, 0x3f, 0x24, 0x6c, 0x3b, 0xc2, 0x09, 0x4c, 0x6e, 0xc8, 0x5a, 0x6a, 0xe8,
	0xa6, 0xc8, 0x22, 0xc2, 0xff, 0x7e, 0x92, 0xc8, 0xd1, 0x6d, 0xf6, 0xf2, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x55, 0xd6, 0xcb, 0xa2, 0xf0, 0x05, 0x00, 0x00,
}
