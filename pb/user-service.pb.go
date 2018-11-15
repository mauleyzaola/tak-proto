// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-service.proto

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

type UserListResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_service_c195da2bfd154828, []int{0}
}
func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (dst *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(dst, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserListResponse)(nil), "pb.UserListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Load(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	LoadPublic(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserListResponse, error)
	Validate(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	ChangePassword(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Load(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserService/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoadPublic(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserService/LoadPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Remove(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Validate(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Ping(context.Context, *Request) (*Response, error)
	Load(context.Context, *User) (*User, error)
	LoadPublic(context.Context, *User) (*User, error)
	Create(context.Context, *User) (*Response, error)
	Update(context.Context, *User) (*Response, error)
	Remove(context.Context, *User) (*Response, error)
	List(context.Context, *User) (*UserListResponse, error)
	Validate(context.Context, *User) (*Response, error)
	ChangePassword(context.Context, *User) (*Response, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Load(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoadPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoadPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/LoadPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoadPublic(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Remove(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Validate(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UserService_Ping_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _UserService_Load_Handler,
		},
		{
			MethodName: "LoadPublic",
			Handler:    _UserService_LoadPublic_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _UserService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _UserService_Validate_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service.proto",
}

func init() { proto.RegisterFile("user-service.proto", fileDescriptor_user_service_c195da2bfd154828) }

var fileDescriptor_user_service_c195da2bfd154828 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x63, 0x8d, 0xa1, 0x4c, 0xb4, 0xc8, 0xe0, 0x41, 0x82, 0x88, 0xac, 0x1e, 0x8a, 0x60,
	0x0e, 0xf1, 0x11, 0x7a, 0xed, 0x21, 0x44, 0xea, 0x7d, 0xd3, 0x0c, 0x75, 0xa1, 0x66, 0xd7, 0x99,
	0xa4, 0xbe, 0xac, 0x0f, 0x23, 0xeb, 0xc6, 0x12, 0x84, 0xda, 0xdb, 0xff, 0xcf, 0x7c, 0x33, 0x3f,
	0xcc, 0x00, 0xf6, 0x42, 0xfc, 0x24, 0xc4, 0x3b, 0xb3, 0xa6, 0xdc, 0xb1, 0xed, 0x2c, 0x4e, 0x5c,
	0x9d, 0xcd, 0x98, 0xc4, 0xd9, 0x56, 0x86, 0x5a, 0x76, 0xc1, 0xf4, 0xd1, 0x93, 0x74, 0x83, 0x05,
	0x3f, 0x16, 0xb4, 0x2a, 0xe0, 0x72, 0x25, 0xc4, 0x4b, 0x23, 0x5d, 0x35, 0x0c, 0xe1, 0x2d, 0x9c,
	0x79, 0x42, 0xae, 0x4f, 0xee, 0x4e, 0xe7, 0x69, 0x31, 0xcd, 0x5d, 0x9d, 0x7b, 0xa8, 0x0a, 0xe5,
	0xe2, 0x6b, 0x02, 0xa9, 0xf7, 0x2f, 0x21, 0x18, 0xef, 0x21, 0x2e, 0x4d, 0xbb, 0xc1, 0xd4, 0x83,
	0x55, 0x88, 0xca, 0xce, 0x83, 0x09, 0x2b, 0x55, 0x84, 0x37, 0x10, 0x2f, 0xad, 0x6e, 0x70, 0xbf,
	0x2d, 0xdb, 0x2b, 0x15, 0xa1, 0x02, 0xf0, 0xdd, 0xb2, 0xaf, 0xb7, 0x66, 0x7d, 0x90, 0x49, 0x16,
	0x4c, 0xba, 0xa3, 0x51, 0xff, 0x6f, 0x8a, 0x82, 0x64, 0xe5, 0x9a, 0xa3, 0x4c, 0x45, 0xef, 0x76,
	0xf7, 0x1f, 0x33, 0x87, 0xd8, 0x9f, 0x64, 0x44, 0x5c, 0xfd, 0xaa, 0xf1, 0xa9, 0x54, 0x84, 0x0f,
	0x30, 0x7d, 0xd5, 0x5b, 0x73, 0x24, 0xf3, 0x11, 0x66, 0x8b, 0x37, 0xdd, 0x6e, 0xa8, 0xd4, 0x22,
	0x9f, 0x96, 0x9b, 0xc3, 0x6c, 0x9d, 0xfc, 0x7c, 0xe6, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x31,
	0xf7, 0x30, 0x85, 0xde, 0x01, 0x00, 0x00,
}
