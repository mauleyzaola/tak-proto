// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email-service.proto

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

type EmailRequest struct {
	Email                *Email   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_email_service_5958dc32ba35f2bd, []int{0}
}
func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (dst *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(dst, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetEmail() *Email {
	if m != nil {
		return m.Email
	}
	return nil
}

func init() {
	proto.RegisterType((*EmailRequest)(nil), "pb.EmailRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Send(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Response, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.EmailService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Send(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.EmailService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Ping(context.Context, *Request) (*Response, error)
	Send(context.Context, *EmailRequest) (*Response, error)
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EmailService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EmailService_Ping_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email-service.proto",
}

func init() { proto.RegisterFile("email-service.proto", fileDescriptor_email_service_5958dc32ba35f2bd) }

var fileDescriptor_email_service_5958dc32ba35f2bd = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x92, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x86, 0x8a, 0x49, 0xf1,
	0x16, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x40, 0xb9, 0xdc, 0x60, 0x7d, 0x10, 0x8e, 0x92, 0x3e,
	0x17, 0x8f, 0x2b, 0x88, 0x1b, 0x04, 0x51, 0x22, 0x24, 0xcf, 0xc5, 0x0a, 0x96, 0x96, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd4, 0x2b, 0x48, 0xd2, 0x83, 0x28, 0x80, 0x88, 0x1b, 0xc5, 0x42,
	0x35, 0x04, 0x43, 0xac, 0x15, 0x52, 0xe6, 0x62, 0x09, 0xc8, 0xcc, 0x4b, 0x17, 0xe2, 0x06, 0xa9,
	0x84, 0x9a, 0x22, 0xc5, 0x03, 0xe1, 0x40, 0x5c, 0xa1, 0xc4, 0x20, 0xa4, 0xc1, 0xc5, 0x12, 0x9c,
	0x9a, 0x97, 0x22, 0x24, 0x80, 0x30, 0x0e, 0xbb, 0xca, 0x24, 0x36, 0xb0, 0xb3, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x04, 0x6a, 0xf0, 0x1f, 0xdd, 0x00, 0x00, 0x00,
}
