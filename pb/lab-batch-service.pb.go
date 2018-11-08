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
	Batches              []*LabBatch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
	Batch                *LabBatch   `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchCodeRequest) Reset()         { *m = BatchCodeRequest{} }
func (m *BatchCodeRequest) String() string { return proto.CompactTextString(m) }
func (*BatchCodeRequest) ProtoMessage()    {}
func (*BatchCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4, []int{0}
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

func (m *BatchCodeRequest) GetBatches() []*LabBatch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *BatchCodeRequest) GetBatch() *LabBatch {
	if m != nil {
		return m.Batch
	}
	return nil
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
	return fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4, []int{1}
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
	return fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4, []int{2}
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
	return fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4, []int{3}
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
	return fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4, []int{4}
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lab-batch-service.proto",
}

func init() {
	proto.RegisterFile("lab-batch-service.proto", fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4)
}

var fileDescriptor_lab_batch_service_4a9c4cf43f8a47c4 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xb6, 0x34, 0xb4, 0xb7, 0x9f, 0xb2, 0x26, 0x11, 0x45, 0x48, 0x94, 0x30, 0x95, 0xbe,
	0x34, 0x93, 0x0a, 0xda, 0x03, 0x8f, 0x9b, 0xc6, 0x04, 0xaa, 0xc4, 0x48, 0x81, 0x47, 0x24, 0x27,
	0xf1, 0x4a, 0x44, 0x1b, 0x1b, 0xc7, 0x81, 0xfe, 0x15, 0x7e, 0x21, 0x7f, 0x03, 0xd9, 0x4e, 0xd2,
	0xa4, 0x2d, 0xca, 0xde, 0x7c, 0xee, 0x39, 0xf7, 0xf8, 0xda, 0xbe, 0xbe, 0xf0, 0x74, 0x8d, 0x83,
	0x59, 0x80, 0x45, 0xf8, 0x7d, 0x96, 0x12, 0xfe, 0x2b, 0x0e, 0x89, 0xc7, 0x38, 0x15, 0x14, 0x9d,
	0xb2, 0xc0, 0x19, 0x96, 0xa4, 0x0e, 0x3a, 0x7d, 0x4e, 0x7e, 0x66, 0x24, 0x15, 0x39, 0x1c, 0x70,
	0x92, 0x32, 0x9a, 0xa4, 0x79, 0x8e, 0x03, 0x2b, 0x1e, 0x47, 0xf9, 0xfa, 0xf9, 0x8a, 0xd2, 0xd5,
	0x9a, 0x5c, 0x2a, 0x14, 0x64, 0x0f, 0x97, 0x22, 0xde, 0x90, 0x54, 0xe0, 0x0d, 0xd3, 0x02, 0xf7,
	0x1b, 0x8c, 0xae, 0xa5, 0xf5, 0x0d, 0x8d, 0x88, 0xaf, 0x6d, 0xd1, 0x04, 0x9e, 0xa8, 0xed, 0x48,
	0x6a, 0x1b, 0xe3, 0xb3, 0x69, 0x77, 0xde, 0xf3, 0x58, 0xe0, 0x2d, 0x70, 0xa0, 0x94, 0x7e, 0x41,
	0x22, 0x17, 0x5a, 0x6a, 0x69, 0x9f, 0x8e, 0x8d, 0x03, 0x95, 0xa6, 0xdc, 0x2b, 0xe8, 0x15, 0xa1,
	0x45, 0xfc, 0x78, 0x6f, 0x59, 0xd7, 0x1d, 0x8f, 0x23, 0x1d, 0xcd, 0xeb, 0x7a, 0x01, 0xa6, 0x3c,
	0x9a, 0x6d, 0xa8, 0xed, 0xfa, 0x32, 0x51, 0x6a, 0xde, 0x27, 0x2c, 0x13, 0xbe, 0xa2, 0xd0, 0x2b,
	0xb0, 0x18, 0xe6, 0x78, 0x93, 0xe6, 0x35, 0x0d, 0xa5, 0x48, 0x99, 0xdc, 0xab, 0xb0, 0x9f, 0xd3,
	0xee, 0x27, 0xe8, 0x94, 0xfe, 0xc8, 0xad, 0x19, 0x0f, 0x0a, 0xe3, 0x8f, 0x99, 0xd8, 0x39, 0x8f,
	0xc1, 0xe4, 0xf4, 0xb7, 0xf4, 0x3d, 0xac, 0x5a, 0x31, 0xee, 0x5f, 0x13, 0xba, 0x95, 0xad, 0xd0,
	0x39, 0xb4, 0x42, 0x1a, 0xa9, 0x83, 0x1a, 0xd3, 0x8e, 0xaf, 0x01, 0x7a, 0x06, 0x9d, 0x35, 0x0d,
	0xb1, 0x88, 0x69, 0xa2, 0xcd, 0x3a, 0xfe, 0x2e, 0x80, 0xde, 0xc1, 0x88, 0x6c, 0x59, 0xcc, 0x15,
	0x5c, 0x0a, 0xcc, 0x45, 0x6a, 0x9f, 0xa9, 0xaa, 0x1c, 0x4f, 0x3f, 0xa5, 0x57, 0x3c, 0xa5, 0xf7,
	0xb9, 0x78, 0x4a, 0xff, 0x20, 0x07, 0x5d, 0xc3, 0x60, 0x17, 0xbb, 0x4d, 0xa2, 0xd4, 0x36, 0x1b,
	0x5d, 0xf6, 0x32, 0xd0, 0x5b, 0x80, 0x08, 0x0b, 0x92, 0x57, 0xd1, 0x6a, 0xcc, 0xaf, 0xa8, 0xd1,
	0x15, 0xb4, 0x25, 0x52, 0x3b, 0x5b, 0x8d, 0x99, 0xa5, 0x56, 0xde, 0x4e, 0x9a, 0x31, 0xb6, 0x8e,
	0x09, 0x4f, 0xed, 0x8e, 0xbe, 0x9d, 0x32, 0x80, 0x3e, 0x00, 0xda, 0xe0, 0x24, 0x7b, 0xc0, 0xa1,
	0xc8, 0x38, 0x89, 0xf2, 0xca, 0xa0, 0xd1, 0xff, 0x48, 0x96, 0xbc, 0xe9, 0x6a, 0x54, 0x55, 0xda,
	0x6d, 0xbe, 0xe9, 0xfd, 0x1c, 0x74, 0x01, 0xfd, 0xa2, 0x40, 0xf5, 0xf8, 0x76, 0x4f, 0xbd, 0x76,
	0x3d, 0x28, 0x7b, 0x41, 0x7f, 0x95, 0xbe, 0xee, 0x05, 0x05, 0x90, 0x03, 0xed, 0x4d, 0x9c, 0x2c,
	0x05, 0x0d, 0x7f, 0xd8, 0x83, 0xb1, 0x31, 0x35, 0xfc, 0x12, 0x2b, 0x0e, 0x6f, 0x35, 0x37, 0xcc,
	0xb9, 0x1c, 0xcf, 0xff, 0x9c, 0xc1, 0xb0, 0x68, 0xbe, 0xa5, 0x9e, 0x17, 0xe8, 0x25, 0x98, 0xf7,
	0x71, 0xb2, 0x42, 0x5d, 0xd9, 0x99, 0xf9, 0x8f, 0x71, 0x7a, 0x1a, 0xe8, 0xf1, 0xe0, 0x9e, 0xa0,
	0x0b, 0x30, 0x17, 0x14, 0x47, 0xa8, 0xd6, 0xbe, 0x4e, 0x0d, 0xb9, 0x27, 0x68, 0x02, 0xd6, 0x0d,
	0x27, 0x58, 0x90, 0x63, 0xba, 0x8a, 0xdb, 0x04, 0xac, 0x2f, 0x2c, 0x6a, 0xd6, 0x4d, 0xc1, 0x54,
	0x7f, 0xbf, 0xae, 0x1a, 0x55, 0x91, 0xe4, 0x95, 0xb2, 0xfd, 0x15, 0xaf, 0xe3, 0x47, 0x78, 0xbe,
	0x81, 0xde, 0x1d, 0x49, 0xca, 0xd1, 0x85, 0xce, 0xcb, 0x8f, 0x5e, 0x99, 0x64, 0x07, 0x27, 0x9b,
	0x81, 0x29, 0x3f, 0xb6, 0x56, 0xef, 0xcf, 0x17, 0xa7, 0x5f, 0x8b, 0xba, 0x27, 0x68, 0x0e, 0xd6,
	0xed, 0x96, 0x51, 0x2e, 0xfe, 0x93, 0x70, 0xe4, 0x08, 0x81, 0xa5, 0xba, 0xe6, 0xf5, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0x52, 0x62, 0x45, 0xd3, 0x05, 0x00, 0x00,
}
