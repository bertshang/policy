// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package policy_user_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	DeletedAt            string   `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,7,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "policy.user.service.User")
	proto.RegisterType((*Request)(nil), "policy.user.service.Request")
	proto.RegisterType((*Response)(nil), "policy.user.service.Response")
	proto.RegisterType((*Error)(nil), "policy.user.service.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x8b, 0xdb, 0x40,
	0x10, 0xc5, 0xa3, 0xbf, 0xb6, 0xc6, 0xc4, 0xc5, 0xe4, 0x0f, 0x1b, 0x13, 0x83, 0x11, 0x04, 0xd2,
	0x44, 0x06, 0xa7, 0x4e, 0x21, 0x92, 0xe0, 0x5e, 0x49, 0x6a, 0x23, 0x6b, 0xa7, 0x10, 0x91, 0xb4,
	0xca, 0xee, 0x2a, 0xc7, 0x7d, 0x99, 0xfb, 0x6a, 0xf7, 0x35, 0xae, 0x3c, 0x76, 0x24, 0x1f, 0x57,
	0xd8, 0x77, 0xc5, 0x35, 0xf6, 0xce, 0xef, 0xbd, 0x7d, 0x2c, 0x6f, 0x04, 0xef, 0x7a, 0xad, 0xac,
	0xda, 0x0e, 0x86, 0x34, 0xff, 0x64, 0x3c, 0xe3, 0x9b, 0x5e, 0x35, 0x75, 0x75, 0x9d, 0x31, 0x32,
	0xa4, 0xff, 0xd7, 0x15, 0xa5, 0x77, 0x1e, 0x84, 0x7f, 0x0c, 0x69, 0x5c, 0x82, 0x5f, 0x4b, 0xe1,
	0x6d, 0xbc, 0xcf, 0x49, 0xe1, 0xd7, 0x12, 0x11, 0xc2, 0xae, 0x6c, 0x49, 0xf8, 0x4c, 0xf8, 0x8c,
	0x6f, 0x21, 0xa2, 0xb6, 0xac, 0x1b, 0x11, 0x30, 0x1c, 0x07, 0x5c, 0xc1, 0xbc, 0x2f, 0x8d, 0xb9,
	0x52, 0x5a, 0x8a, 0x90, 0x85, 0x87, 0x19, 0xdf, 0x43, 0x6c, 0x6c, 0x69, 0x07, 0x23, 0x22, 0x56,
	0xa6, 0x09, 0xd7, 0x00, 0x92, 0x1a, 0xb2, 0x24, 0x0f, 0xa5, 0x15, 0x31, 0x6b, 0xc9, 0x44, 0x72,
	0x8b, 0x9f, 0x60, 0xa9, 0xa9, 0xa5, 0xf6, 0x48, 0xfa, 0x60, 0xd5, 0x5f, 0xea, 0xc4, 0x8c, 0x2d,
	0xaf, 0x4f, 0xf4, 0xb7, 0x83, 0x2e, 0xa5, 0xd2, 0x54, 0x4e, 0x29, 0xf3, 0x31, 0x65, 0x22, 0xb9,
	0x75, 0xf2, 0xd0, 0xcb, 0x93, 0x9c, 0x8c, 0xf2, 0x44, 0x72, 0x9b, 0x26, 0x30, 0x2b, 0xe8, 0xdf,
	0x40, 0xc6, 0xa6, 0x37, 0x1e, 0xcc, 0x0b, 0x32, 0xbd, 0xea, 0x0c, 0xe1, 0x17, 0x08, 0x5d, 0x45,
	0xdc, 0xc5, 0x62, 0xf7, 0x21, 0x3b, 0x53, 0x5b, 0xe6, 0x2a, 0x2b, 0xd8, 0x86, 0x5b, 0x88, 0xdc,
	0xbf, 0x11, 0xfe, 0x26, 0x78, 0xda, 0x3f, 0xfa, 0x70, 0x07, 0x31, 0x69, 0xad, 0xb4, 0x11, 0x01,
	0xdf, 0x58, 0x9d, 0xbd, 0xf1, 0xd3, 0x59, 0x8a, 0xc9, 0x99, 0x7e, 0x83, 0x88, 0x81, 0x5b, 0x4b,
	0xa5, 0x24, 0xf1, 0xe3, 0xa2, 0x82, 0xcf, 0xb8, 0x81, 0x85, 0x24, 0x53, 0xe9, 0xba, 0xb7, 0xb5,
	0xea, 0xa6, 0x8d, 0x3d, 0x46, 0xbb, 0x5b, 0x0f, 0x16, 0xee, 0x09, 0xbf, 0xc6, 0x70, 0xfc, 0x01,
	0xf1, 0x77, 0xae, 0x09, 0x2f, 0x3f, 0x77, 0xb5, 0x3e, 0x2b, 0x9d, 0x6a, 0x4a, 0x5f, 0x61, 0x0e,
	0xc1, 0x9e, 0xec, 0x8b, 0x22, 0xf6, 0x10, 0xef, 0xc9, 0xe6, 0x4d, 0x83, 0x1f, 0x2f, 0x58, 0x79,
	0x41, 0xcf, 0x06, 0x1d, 0x63, 0xfe, 0xc6, 0xbf, 0xde, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xe2,
	0xe3, 0xf3, 0xfc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "policy.user.service"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}
