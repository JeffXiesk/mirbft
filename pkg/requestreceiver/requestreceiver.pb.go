// Code generated by protoc-gen-go. DO NOT EDIT.
// source: requestreceiver/requestreceiver.proto

package requestreceiver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	messagepb "github.com/hyperledger-labs/mirbft/pkg/pb/messagepb"
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

type ByeBye struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByeBye) Reset()         { *m = ByeBye{} }
func (m *ByeBye) String() string { return proto.CompactTextString(m) }
func (*ByeBye) ProtoMessage()    {}
func (*ByeBye) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab513c9b3159a7d, []int{0}
}

func (m *ByeBye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByeBye.Unmarshal(m, b)
}
func (m *ByeBye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByeBye.Marshal(b, m, deterministic)
}
func (m *ByeBye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByeBye.Merge(m, src)
}
func (m *ByeBye) XXX_Size() int {
	return xxx_messageInfo_ByeBye.Size(m)
}
func (m *ByeBye) XXX_DiscardUnknown() {
	xxx_messageInfo_ByeBye.DiscardUnknown(m)
}

var xxx_messageInfo_ByeBye proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ByeBye)(nil), "requestreceiver.ByeBye")
}

func init() {
	proto.RegisterFile("requestreceiver/requestreceiver.proto", fileDescriptor_eab513c9b3159a7d)
}

var fileDescriptor_eab513c9b3159a7d = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x29, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0x2d, 0xd2, 0x47, 0xe3, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x09, 0x4b, 0x49, 0xe6, 0xa6, 0x16, 0x17, 0x27, 0xa6,
	0xa7, 0x16, 0x24, 0xe9, 0xc3, 0x59, 0x10, 0xb5, 0x4a, 0x1c, 0x5c, 0x6c, 0x4e, 0x95, 0xa9, 0x4e,
	0x95, 0xa9, 0x46, 0x5e, 0x5c, 0xfc, 0x41, 0x10, 0x7d, 0x41, 0x50, 0x7d, 0x42, 0xe6, 0x5c, 0x6c,
	0x3e, 0x99, 0xc5, 0x25, 0xa9, 0x79, 0x42, 0x42, 0x7a, 0x08, 0x8d, 0x50, 0x55, 0x52, 0xe2, 0x7a,
	0xe8, 0xd6, 0x43, 0x4c, 0xd2, 0x60, 0x74, 0xb2, 0x88, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xa8, 0x2c, 0x48, 0x2d, 0xca, 0x49, 0x4d, 0x49, 0x4f, 0x2d,
	0xd2, 0xcd, 0x49, 0x4c, 0x2a, 0xd6, 0xcf, 0xcd, 0x2c, 0x4a, 0x4a, 0x2b, 0xd1, 0x2f, 0xc8, 0x4e,
	0x47, 0xf7, 0x41, 0x12, 0x1b, 0xd8, 0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xe4,
	0x27, 0x65, 0xeb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RequestReceiverClient is the client API for RequestReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestReceiverClient interface {
	Listen(ctx context.Context, opts ...grpc.CallOption) (RequestReceiver_ListenClient, error)
}

type requestReceiverClient struct {
	cc *grpc.ClientConn
}

func NewRequestReceiverClient(cc *grpc.ClientConn) RequestReceiverClient {
	return &requestReceiverClient{cc}
}

func (c *requestReceiverClient) Listen(ctx context.Context, opts ...grpc.CallOption) (RequestReceiver_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RequestReceiver_serviceDesc.Streams[0], "/requestreceiver.RequestReceiver/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestReceiverListenClient{stream}
	return x, nil
}

type RequestReceiver_ListenClient interface {
	Send(*messagepb.Request) error
	CloseAndRecv() (*ByeBye, error)
	grpc.ClientStream
}

type requestReceiverListenClient struct {
	grpc.ClientStream
}

func (x *requestReceiverListenClient) Send(m *messagepb.Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *requestReceiverListenClient) CloseAndRecv() (*ByeBye, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ByeBye)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RequestReceiverServer is the server API for RequestReceiver service.
type RequestReceiverServer interface {
	Listen(RequestReceiver_ListenServer) error
}

// UnimplementedRequestReceiverServer can be embedded to have forward compatible implementations.
type UnimplementedRequestReceiverServer struct {
}

func (*UnimplementedRequestReceiverServer) Listen(srv RequestReceiver_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}

func RegisterRequestReceiverServer(s *grpc.Server, srv RequestReceiverServer) {
	s.RegisterService(&_RequestReceiver_serviceDesc, srv)
}

func _RequestReceiver_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RequestReceiverServer).Listen(&requestReceiverListenServer{stream})
}

type RequestReceiver_ListenServer interface {
	SendAndClose(*ByeBye) error
	Recv() (*messagepb.Request, error)
	grpc.ServerStream
}

type requestReceiverListenServer struct {
	grpc.ServerStream
}

func (x *requestReceiverListenServer) SendAndClose(m *ByeBye) error {
	return x.ServerStream.SendMsg(m)
}

func (x *requestReceiverListenServer) Recv() (*messagepb.Request, error) {
	m := new(messagepb.Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RequestReceiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "requestreceiver.RequestReceiver",
	HandlerType: (*RequestReceiverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _RequestReceiver_Listen_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "requestreceiver/requestreceiver.proto",
}