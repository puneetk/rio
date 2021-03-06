// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/sds.proto

package envoy_service_discovery_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2a4da2e99d9a3e6, []int{0}
}

func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (m *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(m, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.discovery.v2.SdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/sds.proto", fileDescriptor_f2a4da2e99d9a3e6)
}

var fileDescriptor_f2a4da2e99d9a3e6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x4d, 0xc5, 0x2a, 0x47, 0x05, 0x09, 0x68, 0x35, 0xd4, 0x5a, 0x6a, 0x91, 0xe2, 0x70,
	0x91, 0x74, 0xeb, 0x58, 0x8a, 0x73, 0x31, 0xe0, 0xaa, 0x67, 0xf2, 0xd1, 0x1e, 0x34, 0xb9, 0x78,
	0x77, 0x39, 0xcd, 0x26, 0x4e, 0xe2, 0xe2, 0x20, 0x88, 0x0f, 0xe0, 0x53, 0xf8, 0x04, 0xae, 0xe2,
	0x2b, 0x38, 0x3a, 0xe9, 0x2e, 0xd2, 0x5c, 0xda, 0x92, 0x4a, 0x9d, 0x9c, 0x7f, 0xff, 0xff, 0xef,
	0xbe, 0xe3, 0xfb, 0x50, 0x03, 0x42, 0xc5, 0x12, 0x5b, 0x00, 0x57, 0xd4, 0x03, 0xdb, 0xa7, 0xc2,
	0x63, 0x0a, 0x78, 0x62, 0x2b, 0xc7, 0x16, 0xbe, 0xc0, 0x11, 0x67, 0x92, 0x99, 0x56, 0x9a, 0xc2,
	0x59, 0x0a, 0x4f, 0x52, 0x58, 0x39, 0x56, 0x45, 0x1b, 0x48, 0x44, 0x47, 0x9d, 0x29, 0x4a, 0x9b,
	0x56, 0xa5, 0xcf, 0x58, 0x7f, 0x08, 0x29, 0x26, 0x61, 0xc8, 0x24, 0x91, 0x94, 0x85, 0x99, 0xd7,
	0xaa, 0x65, 0xdd, 0x29, 0xb0, 0x39, 0x08, 0x16, 0x73, 0x0f, 0xb2, 0x44, 0x35, 0xf6, 0x23, 0x92,
	0x0b, 0x04, 0xb4, 0xcf, 0x89, 0x1c, 0xf3, 0xed, 0x5f, 0x5c, 0x48, 0x22, 0xe3, 0xec, 0x81, 0x3a,
	0x42, 0x2b, 0xae, 0x2f, 0xba, 0x71, 0x10, 0x24, 0xce, 0x47, 0x01, 0x6d, 0xb8, 0xe0, 0x71, 0x90,
	0xdd, 0xf1, 0x90, 0xae, 0xfe, 0x90, 0x79, 0x82, 0x4a, 0x5d, 0x18, 0x4a, 0xa2, 0xb1, 0x30, 0x77,
	0xb1, 0xfe, 0x30, 0x89, 0x28, 0x56, 0x0e, 0x4e, 0xd9, 0xa4, 0x74, 0x04, 0xe7, 0x31, 0x08, 0x69,
	0x35, 0xfe, 0x0e, 0x89, 0x88, 0x85, 0x02, 0xea, 0x0b, 0x4d, 0xe3, 0xc0, 0x30, 0x8f, 0xd1, 0xaa,
	0x2b, 0x39, 0x90, 0x60, 0xfc, 0x42, 0x75, 0xa6, 0x3c, 0x2b, 0xdf, 0x99, 0xcb, 0x73, 0xde, 0x0b,
	0x54, 0x3a, 0x04, 0xe9, 0x0d, 0xfe, 0x4d, 0xbb, 0x77, 0xfd, 0xf6, 0x7e, 0x5f, 0x28, 0xd7, 0xd7,
	0x73, 0xcb, 0x6c, 0x0b, 0xed, 0x4f, 0xe1, 0x62, 0xdb, 0xd8, 0xb7, 0x6a, 0xb7, 0x4f, 0x0f, 0x5f,
	0xcb, 0x16, 0xda, 0xcc, 0xf9, 0x48, 0x2c, 0x07, 0x58, 0xcf, 0xd2, 0x39, 0xfd, 0x7c, 0xfc, 0xbe,
	0x5b, 0xda, 0x32, 0xcb, 0xf9, 0xdb, 0xd1, 0x26, 0xac, 0x5a, 0xcf, 0x57, 0x2f, 0xaf, 0xc5, 0xc2,
	0x9a, 0x81, 0x9a, 0x94, 0xe9, 0xa9, 0x22, 0xce, 0x2e, 0x13, 0x3c, 0xff, 0xd4, 0x3a, 0xa3, 0x5d,
	0xf6, 0x46, 0x7b, 0xed, 0x19, 0x37, 0x86, 0x71, 0x56, 0x4c, 0x77, 0xdc, 0xfa, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x32, 0x5d, 0xb1, 0xc7, 0xc4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error)
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.SecretDiscoveryService/DeltaSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceDeltaSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_DeltaSecretsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceDeltaSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	DeltaSecrets(SecretDiscoveryService_DeltaSecretsServer) error
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedSecretDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretDiscoveryServiceServer struct {
}

func (*UnimplementedSecretDiscoveryServiceServer) DeltaSecrets(srv SecretDiscoveryService_DeltaSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) StreamSecrets(srv SecretDiscoveryService_StreamSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) FetchSecrets(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSecrets not implemented")
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_DeltaSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).DeltaSecrets(&secretDiscoveryServiceDeltaSecretsServer{stream})
}

type SecretDiscoveryService_DeltaSecretsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceDeltaSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaSecrets",
			Handler:       _SecretDiscoveryService_DeltaSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/sds.proto",
}
