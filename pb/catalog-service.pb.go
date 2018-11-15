// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog-service.proto

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

type GridCatalogRequest struct {
	Grid                 *GridInput `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Params               *Catalog   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GridCatalogRequest) Reset()         { *m = GridCatalogRequest{} }
func (m *GridCatalogRequest) String() string { return proto.CompactTextString(m) }
func (*GridCatalogRequest) ProtoMessage()    {}
func (*GridCatalogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_service_9169113f1f5f2efa, []int{0}
}
func (m *GridCatalogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridCatalogRequest.Unmarshal(m, b)
}
func (m *GridCatalogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridCatalogRequest.Marshal(b, m, deterministic)
}
func (dst *GridCatalogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridCatalogRequest.Merge(dst, src)
}
func (m *GridCatalogRequest) XXX_Size() int {
	return xxx_messageInfo_GridCatalogRequest.Size(m)
}
func (m *GridCatalogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GridCatalogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GridCatalogRequest proto.InternalMessageInfo

func (m *GridCatalogRequest) GetGrid() *GridInput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridCatalogRequest) GetParams() *Catalog {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GridCatalogRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GridCatalogResponse struct {
	Grid                 *GridOutput `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Rows                 []*Catalog  `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	Name                 string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GridCatalogResponse) Reset()         { *m = GridCatalogResponse{} }
func (m *GridCatalogResponse) String() string { return proto.CompactTextString(m) }
func (*GridCatalogResponse) ProtoMessage()    {}
func (*GridCatalogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_service_9169113f1f5f2efa, []int{1}
}
func (m *GridCatalogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GridCatalogResponse.Unmarshal(m, b)
}
func (m *GridCatalogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GridCatalogResponse.Marshal(b, m, deterministic)
}
func (dst *GridCatalogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GridCatalogResponse.Merge(dst, src)
}
func (m *GridCatalogResponse) XXX_Size() int {
	return xxx_messageInfo_GridCatalogResponse.Size(m)
}
func (m *GridCatalogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GridCatalogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GridCatalogResponse proto.InternalMessageInfo

func (m *GridCatalogResponse) GetGrid() *GridOutput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *GridCatalogResponse) GetRows() []*Catalog {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *GridCatalogResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CatalogResponse struct {
	Grid                 *GridOutput `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	Rows                 []*Catalog  `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CatalogResponse) Reset()         { *m = CatalogResponse{} }
func (m *CatalogResponse) String() string { return proto.CompactTextString(m) }
func (*CatalogResponse) ProtoMessage()    {}
func (*CatalogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_catalog_service_9169113f1f5f2efa, []int{2}
}
func (m *CatalogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogResponse.Unmarshal(m, b)
}
func (m *CatalogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogResponse.Marshal(b, m, deterministic)
}
func (dst *CatalogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogResponse.Merge(dst, src)
}
func (m *CatalogResponse) XXX_Size() int {
	return xxx_messageInfo_CatalogResponse.Size(m)
}
func (m *CatalogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogResponse proto.InternalMessageInfo

func (m *CatalogResponse) GetGrid() *GridOutput {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *CatalogResponse) GetRows() []*Catalog {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*GridCatalogRequest)(nil), "pb.GridCatalogRequest")
	proto.RegisterType((*GridCatalogResponse)(nil), "pb.GridCatalogResponse")
	proto.RegisterType((*CatalogResponse)(nil), "pb.CatalogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	Validate(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error)
	Load(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Catalog, error)
	Create(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*CatalogResponse, error)
	Grid(ctx context.Context, in *GridCatalogRequest, opts ...grpc.CallOption) (*GridCatalogResponse, error)
	Ping(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) Validate(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Load(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Catalog, error) {
	out := new(Catalog)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Create(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Update(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Remove(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) List(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*CatalogResponse, error) {
	out := new(CatalogResponse)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Grid(ctx context.Context, in *GridCatalogRequest, opts ...grpc.CallOption) (*GridCatalogResponse, error) {
	out := new(GridCatalogResponse)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Grid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) Ping(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CatalogService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	Validate(context.Context, *Catalog) (*Response, error)
	Load(context.Context, *Catalog) (*Catalog, error)
	Create(context.Context, *Catalog) (*Response, error)
	Update(context.Context, *Catalog) (*Response, error)
	Remove(context.Context, *Catalog) (*Response, error)
	List(context.Context, *Catalog) (*CatalogResponse, error)
	Grid(context.Context, *GridCatalogRequest) (*GridCatalogResponse, error)
	Ping(context.Context, *Catalog) (*Response, error)
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Validate(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Load(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Create(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Update(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Remove(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).List(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Grid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GridCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Grid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Grid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Grid(ctx, req.(*GridCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatalogService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).Ping(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _CatalogService_Validate_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _CatalogService_Load_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CatalogService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CatalogService_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _CatalogService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CatalogService_List_Handler,
		},
		{
			MethodName: "Grid",
			Handler:    _CatalogService_Grid_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CatalogService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog-service.proto",
}

func init() {
	proto.RegisterFile("catalog-service.proto", fileDescriptor_catalog_service_9169113f1f5f2efa)
}

var fileDescriptor_catalog_service_9169113f1f5f2efa = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0x2e, 0x14, 0x7d, 0x73, 0x15, 0x32, 0xd4, 0xd2, 0x8b, 0xb3, 0x63, 0x38, 0x04,
	0x7b, 0x98, 0x27, 0xcf, 0x3b, 0x88, 0x30, 0x50, 0x2a, 0xee, 0x9e, 0xae, 0xa1, 0x04, 0xd6, 0x26,
	0x26, 0xe9, 0xfc, 0x6f, 0xfd, 0x5b, 0x24, 0xfd, 0xe1, 0x5a, 0x15, 0xe7, 0x65, 0xb7, 0xe4, 0xbd,
	0xcf, 0x37, 0x9f, 0xd7, 0x34, 0x70, 0xb6, 0x26, 0x9a, 0x6c, 0x78, 0x7a, 0xab, 0xa8, 0xdc, 0xb2,
	0x35, 0x0d, 0x85, 0xe4, 0x9a, 0x63, 0x5b, 0xc4, 0xfe, 0xb0, 0x6e, 0x55, 0x25, 0x1f, 0x52, 0xc9,
	0x92, 0x7a, 0xed, 0x4a, 0xaa, 0x04, 0xcf, 0x55, 0x8d, 0x07, 0x02, 0xf0, 0x83, 0x64, 0xc9, 0xa2,
	0x0a, 0x44, 0xf4, 0xad, 0xa0, 0x4a, 0xe3, 0x2b, 0x40, 0x26, 0xe3, 0x59, 0x63, 0x6b, 0x36, 0x98,
	0x0f, 0x43, 0x11, 0x87, 0x86, 0x7a, 0xcc, 0x45, 0xa1, 0xa3, 0xb2, 0x85, 0x27, 0xe0, 0x08, 0x22,
	0x49, 0xa6, 0x3c, 0xbb, 0x84, 0x06, 0x06, 0x6a, 0x8e, 0xa9, 0x5b, 0x18, 0x03, 0xca, 0x49, 0x46,
	0xbd, 0xfe, 0xd8, 0x9a, 0x1d, 0x47, 0xe5, 0x3a, 0xc8, 0x61, 0xd4, 0x31, 0x56, 0xe3, 0xe0, 0xa0,
	0xa3, 0x74, 0x1b, 0xe5, 0x53, 0xa1, 0x77, 0xce, 0x4b, 0x40, 0x92, 0xbf, 0x1b, 0x63, 0xff, 0xbb,
	0xb1, 0x6c, 0xfc, 0xea, 0x5b, 0xc1, 0xe9, 0x21, 0x5c, 0xf3, 0x0f, 0x1b, 0xdc, 0xba, 0xf2, 0x52,
	0xfd, 0x01, 0x7c, 0x0d, 0x47, 0x2b, 0xb2, 0x61, 0x09, 0xd1, 0x14, 0xb7, 0x13, 0xfe, 0x89, 0xd9,
	0x34, 0xfa, 0xa0, 0x67, 0x06, 0x58, 0x72, 0x92, 0x74, 0xa1, 0xf6, 0x26, 0xe8, 0xe1, 0x29, 0x38,
	0x0b, 0x49, 0xf7, 0x1e, 0x35, 0x05, 0xe7, 0x55, 0x24, 0xff, 0xc1, 0x22, 0x9a, 0xf1, 0xed, 0x1e,
	0xec, 0x06, 0xd0, 0x92, 0x29, 0xdd, 0x85, 0x46, 0xed, 0x8f, 0xdf, 0xb1, 0xf7, 0x80, 0xcc, 0xad,
	0xe1, 0xf3, 0xe6, 0xfe, 0xba, 0x8f, 0xc8, 0xbf, 0xf8, 0x51, 0xff, 0x8a, 0x4e, 0x00, 0x3d, 0xb3,
	0x3c, 0xfd, 0x73, 0x96, 0xd8, 0x29, 0x5f, 0xe8, 0xdd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0xc4, 0xfc, 0x19, 0xe9, 0x02, 0x00, 0x00,
}
