// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/echo/v1/echo_service.proto

package bazel_example_echo_v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type Phrase struct {
	Value                string               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ProcessingTime       *timestamp.Timestamp `protobuf:"bytes,100,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Phrase) Reset()         { *m = Phrase{} }
func (m *Phrase) String() string { return proto.CompactTextString(m) }
func (*Phrase) ProtoMessage()    {}
func (*Phrase) Descriptor() ([]byte, []int) {
	return fileDescriptor_df3dbf3120bd4117, []int{0}
}

func (m *Phrase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phrase.Unmarshal(m, b)
}
func (m *Phrase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phrase.Marshal(b, m, deterministic)
}
func (m *Phrase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phrase.Merge(m, src)
}
func (m *Phrase) XXX_Size() int {
	return xxx_messageInfo_Phrase.Size(m)
}
func (m *Phrase) XXX_DiscardUnknown() {
	xxx_messageInfo_Phrase.DiscardUnknown(m)
}

var xxx_messageInfo_Phrase proto.InternalMessageInfo

func (m *Phrase) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Phrase) GetProcessingTime() *timestamp.Timestamp {
	if m != nil {
		return m.ProcessingTime
	}
	return nil
}

type EchoRequest struct {
	Phrase               *Phrase  `protobuf:"bytes,1,opt,name=phrase,proto3" json:"phrase,omitempty"`
	ValidateOnly         bool     `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	RequestId            string   `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df3dbf3120bd4117, []int{1}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetPhrase() *Phrase {
	if m != nil {
		return m.Phrase
	}
	return nil
}

func (m *EchoRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

func (m *EchoRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type EchoResponse struct {
	Phrase               *Phrase  `protobuf:"bytes,1,opt,name=phrase,proto3" json:"phrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df3dbf3120bd4117, []int{2}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetPhrase() *Phrase {
	if m != nil {
		return m.Phrase
	}
	return nil
}

func init() {
	proto.RegisterType((*Phrase)(nil), "bazel_example.echo.v1.Phrase")
	proto.RegisterType((*EchoRequest)(nil), "bazel_example.echo.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "bazel_example.echo.v1.EchoResponse")
}

func init() { proto.RegisterFile("api/echo/v1/echo_service.proto", fileDescriptor_df3dbf3120bd4117) }

var fileDescriptor_df3dbf3120bd4117 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x8e, 0xd4, 0x30,
	0x14, 0x85, 0x95, 0x01, 0x46, 0xbb, 0x9e, 0x01, 0x84, 0x11, 0x10, 0x22, 0x16, 0x46, 0x59, 0x8a,
	0x14, 0x8b, 0xa3, 0x0d, 0xda, 0x86, 0x72, 0xd1, 0x14, 0x54, 0xa0, 0x40, 0x87, 0x20, 0x72, 0x92,
	0x3b, 0x13, 0x4b, 0x1e, 0x3b, 0xd8, 0x4e, 0x34, 0x3f, 0x42, 0x42, 0x54, 0xf4, 0x3c, 0x1a, 0xaf,
	0xc0, 0x53, 0x50, 0xa1, 0xc4, 0x31, 0xa4, 0x00, 0x0a, 0xaa, 0xc4, 0xf7, 0x9e, 0xa3, 0xf3, 0x1d,
	0xcb, 0xe8, 0x21, 0xad, 0x59, 0x0c, 0x45, 0x25, 0xe3, 0xf6, 0xbc, 0xff, 0x66, 0x1a, 0x54, 0xcb,
	0x0a, 0x20, 0xb5, 0x92, 0x46, 0xe2, 0x3b, 0x39, 0xdd, 0x03, 0xcf, 0x60, 0x4b, 0x37, 0x35, 0x07,
	0xd2, 0x29, 0x48, 0x7b, 0x1e, 0x3c, 0x58, 0x4b, 0xb9, 0xe6, 0x10, 0x77, 0x6e, 0x2a, 0x84, 0x34,
	0xd4, 0x30, 0x29, 0xb4, 0x35, 0x05, 0xf7, 0x46, 0xdb, 0x82, 0x33, 0x10, 0x66, 0x58, 0x3c, 0x1a,
	0x2d, 0x56, 0x0c, 0x78, 0x99, 0xe5, 0x50, 0xd1, 0x96, 0x49, 0x35, 0x08, 0xee, 0x8f, 0x04, 0x0a,
	0xb4, 0x6c, 0x94, 0x23, 0xf9, 0xe5, 0xed, 0x4f, 0x79, 0xb3, 0x8a, 0x0d, 0xdb, 0x80, 0x36, 0x74,
	0x53, 0xbb, 0xd4, 0x96, 0x72, 0x56, 0x52, 0x03, 0xb1, 0xfb, 0xb1, 0x8b, 0xf0, 0x80, 0xa6, 0xaf,
	0x2a, 0x45, 0x35, 0xe0, 0x33, 0x74, 0xad, 0xa5, 0xbc, 0x01, 0xdf, 0x5b, 0x78, 0xd1, 0xf1, 0xe5,
	0xdd, 0x1f, 0x97, 0xb7, 0xd5, 0xad, 0xe8, 0xd3, 0x24, 0x99, 0xbf, 0x7f, 0x4b, 0x9f, 0xec, 0xdf,
	0x1d, 0x92, 0xb3, 0x8b, 0x8f, 0x8f, 0x53, 0x2b, 0xc2, 0xcf, 0xd1, 0xcd, 0x5a, 0xc9, 0x02, 0xb4,
	0x66, 0x62, 0x9d, 0x75, 0x71, 0x7e, 0xb9, 0xf0, 0xa2, 0x59, 0x12, 0x10, 0xcb, 0x42, 0x1c, 0x0b,
	0x79, 0xe3, 0x58, 0xd2, 0x1b, 0xbf, 0x2d, 0xdd, 0x30, 0xfc, 0xe2, 0xa1, 0xd9, 0xb2, 0xa8, 0x64,
	0x0a, 0x1f, 0x1a, 0xd0, 0x06, 0x5f, 0xa0, 0x69, 0xdd, 0xc3, 0xf4, 0x0c, 0xb3, 0xe4, 0x84, 0xfc,
	0xf1, 0x86, 0x89, 0x25, 0x4e, 0x07, 0x31, 0x3e, 0x45, 0xd7, 0x5d, 0xab, 0x4c, 0x0a, 0xbe, 0xf3,
	0x27, 0x0b, 0x2f, 0x3a, 0x4a, 0xe7, 0x6e, 0xf8, 0x52, 0xf0, 0x1d, 0x3e, 0x41, 0x48, 0xd9, 0x98,
	0x8c, 0x95, 0xfe, 0x95, 0xae, 0x63, 0x7a, 0x3c, 0x4c, 0x5e, 0x94, 0xe1, 0x12, 0xcd, 0x2d, 0x89,
	0xae, 0xa5, 0xd0, 0xf0, 0x9f, 0x28, 0xc9, 0xd6, 0x16, 0x7a, 0x6d, 0xdf, 0x09, 0x66, 0xe8, 0x6a,
	0x77, 0xc4, 0xe1, 0x5f, 0xdc, 0xa3, 0xf2, 0xc1, 0xe9, 0x3f, 0x35, 0x16, 0x2b, 0xf4, 0x3f, 0x7f,
	0xfb, 0xfe, 0x75, 0x82, 0xc3, 0x23, 0xf7, 0x24, 0x9f, 0x0d, 0xc9, 0xf9, 0xb4, 0xbf, 0xef, 0xa7,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xa5, 0xfd, 0x4e, 0xb5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/bazel_example.echo.v1.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bazel_example.echo.v1.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bazel_example.echo.v1.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/echo/v1/echo_service.proto",
}