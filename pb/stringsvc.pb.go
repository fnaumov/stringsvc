// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stringsvc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type UppercaseRequest struct {
	S                    string   `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UppercaseRequest) Reset()         { *m = UppercaseRequest{} }
func (m *UppercaseRequest) String() string { return proto.CompactTextString(m) }
func (*UppercaseRequest) ProtoMessage()    {}
func (*UppercaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{0}
}

func (m *UppercaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UppercaseRequest.Unmarshal(m, b)
}
func (m *UppercaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UppercaseRequest.Marshal(b, m, deterministic)
}
func (m *UppercaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UppercaseRequest.Merge(m, src)
}
func (m *UppercaseRequest) XXX_Size() int {
	return xxx_messageInfo_UppercaseRequest.Size(m)
}
func (m *UppercaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UppercaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UppercaseRequest proto.InternalMessageInfo

func (m *UppercaseRequest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type UppercaseResponse struct {
	V                    string   `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UppercaseResponse) Reset()         { *m = UppercaseResponse{} }
func (m *UppercaseResponse) String() string { return proto.CompactTextString(m) }
func (*UppercaseResponse) ProtoMessage()    {}
func (*UppercaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{1}
}

func (m *UppercaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UppercaseResponse.Unmarshal(m, b)
}
func (m *UppercaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UppercaseResponse.Marshal(b, m, deterministic)
}
func (m *UppercaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UppercaseResponse.Merge(m, src)
}
func (m *UppercaseResponse) XXX_Size() int {
	return xxx_messageInfo_UppercaseResponse.Size(m)
}
func (m *UppercaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UppercaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UppercaseResponse proto.InternalMessageInfo

func (m *UppercaseResponse) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func (m *UppercaseResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type CountRequest struct {
	S                    string   `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{2}
}

func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (m *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(m, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

func (m *CountRequest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type CountResponse struct {
	V                    int64    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{3}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

type AuthRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{4}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f8077f7943c5ff, []int{5}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*UppercaseRequest)(nil), "pb.UppercaseRequest")
	proto.RegisterType((*UppercaseResponse)(nil), "pb.UppercaseResponse")
	proto.RegisterType((*CountRequest)(nil), "pb.CountRequest")
	proto.RegisterType((*CountResponse)(nil), "pb.CountResponse")
	proto.RegisterType((*AuthRequest)(nil), "pb.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "pb.AuthResponse")
}

func init() { proto.RegisterFile("stringsvc.proto", fileDescriptor_02f8077f7943c5ff) }

var fileDescriptor_02f8077f7943c5ff = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x96, 0x8a, 0x1d, 0x53, 0x4c, 0x97, 0x0a, 0x25, 0x28, 0x94, 0x3d, 0x79, 0xca,
	0xc1, 0x82, 0x07, 0x6f, 0x22, 0xde, 0x25, 0xa5, 0x3f, 0x20, 0x89, 0x83, 0x06, 0x71, 0x77, 0x9d,
	0xd9, 0xac, 0x07, 0xff, 0xbc, 0x24, 0xd9, 0x84, 0x10, 0x7b, 0xcb, 0x9b, 0x79, 0xf9, 0xde, 0xee,
	0x5b, 0xb8, 0x62, 0x47, 0x95, 0x7e, 0x67, 0x5f, 0xa6, 0x96, 0x8c, 0x33, 0x72, 0x66, 0x0b, 0xb5,
	0x83, 0xf8, 0x68, 0x2d, 0x52, 0x99, 0x33, 0x66, 0xf8, 0x5d, 0x23, 0x3b, 0x19, 0x81, 0xe0, 0xad,
	0xd8, 0x89, 0xbb, 0x65, 0x26, 0x58, 0xed, 0x61, 0x3d, 0x72, 0xb0, 0x35, 0x9a, 0xb1, 0xb1, 0xf8,
	0xde, 0xe2, 0x65, 0x0c, 0x73, 0x24, 0xda, 0xce, 0x5a, 0xdd, 0x7c, 0xaa, 0x1b, 0x88, 0x9e, 0x4d,
	0xad, 0xdd, 0x69, 0xe4, 0x2d, 0xac, 0xc2, 0x76, 0x8a, 0x9b, 0x67, 0xc2, 0xab, 0x17, 0xb8, 0x7c,
	0xaa, 0xdd, 0x47, 0xff, 0x6f, 0x02, 0x17, 0x47, 0x46, 0xd2, 0xf9, 0x17, 0x06, 0xc4, 0xa0, 0x9b,
	0xdd, 0x6b, 0xce, 0xfc, 0x63, 0xe8, 0x2d, 0xc4, 0x0f, 0x5a, 0x3d, 0x40, 0xd4, 0x61, 0x42, 0xc8,
	0x06, 0x16, 0xce, 0x7c, 0xa2, 0x0e, 0x90, 0x4e, 0xfc, 0x3f, 0xfb, 0xfd, 0x2f, 0xac, 0x0e, 0x6d,
	0x53, 0x07, 0x24, 0x5f, 0x95, 0x28, 0x1f, 0x61, 0x39, 0x34, 0x20, 0x37, 0xa9, 0x2d, 0xd2, 0x69,
	0x65, 0xc9, 0xf5, 0x64, 0xda, 0x45, 0xaa, 0x33, 0x99, 0xc2, 0xa2, 0xbd, 0xaa, 0x8c, 0x1b, 0xc7,
	0xb8, 0x93, 0x64, 0x3d, 0x9a, 0xf4, 0xfe, 0xe2, 0xbc, 0x7d, 0x9a, 0xfd, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x29, 0x34, 0xdd, 0xde, 0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringServiceClient interface {
	Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type stringServiceClient struct {
	cc *grpc.ClientConn
}

func NewStringServiceClient(cc *grpc.ClientConn) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error) {
	out := new(UppercaseResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Uppercase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
type StringServiceServer interface {
	Uppercase(context.Context, *UppercaseRequest) (*UppercaseResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
}

// UnimplementedStringServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (*UnimplementedStringServiceServer) Uppercase(ctx context.Context, req *UppercaseRequest) (*UppercaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uppercase not implemented")
}
func (*UnimplementedStringServiceServer) Count(ctx context.Context, req *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}

func RegisterStringServiceServer(s *grpc.Server, srv StringServiceServer) {
	s.RegisterService(&_StringService_serviceDesc, srv)
}

func _StringService_Uppercase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UppercaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Uppercase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Uppercase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Uppercase(ctx, req.(*UppercaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Uppercase",
			Handler:    _StringService_Uppercase_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _StringService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stringsvc.proto",
}
