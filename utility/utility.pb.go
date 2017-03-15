// Code generated by protoc-gen-go.
// source: utility.proto
// DO NOT EDIT!

/*
Package utility is a generated protocol buffer package.

It is generated from these files:
	utility.proto

It has these top-level messages:
	Request
	Response
*/
package utility

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

// Request sent to the server
type Request struct {
	DelayMilliSeconds int32 `protobuf:"varint,1,opt,name=DelayMilliSeconds,json=delayMilliSeconds" json:"DelayMilliSeconds,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetDelayMilliSeconds() int32 {
	if m != nil {
		return m.DelayMilliSeconds
	}
	return 0
}

// Response from the Server
type Response struct {
	PingResponse string `protobuf:"bytes,1,opt,name=PingResponse,json=pingResponse" json:"PingResponse,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetPingResponse() string {
	if m != nil {
		return m.PingResponse
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "utility.Request")
	proto.RegisterType((*Response)(nil), "utility.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Utility service

type UtilityClient interface {
	// Makes a Ping call
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type utilityClient struct {
	cc *grpc.ClientConn
}

func NewUtilityClient(cc *grpc.ClientConn) UtilityClient {
	return &utilityClient{cc}
}

func (c *utilityClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/utility.Utility/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Utility service

type UtilityServer interface {
	// Makes a Ping call
	Ping(context.Context, *Request) (*Response, error)
}

func RegisterUtilityServer(s *grpc.Server, srv UtilityServer) {
	s.RegisterService(&_Utility_serviceDesc, srv)
}

func _Utility_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utility.Utility/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Utility_serviceDesc = grpc.ServiceDesc{
	ServiceName: "utility.Utility",
	HandlerType: (*UtilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Utility_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utility.proto",
}

func init() { proto.RegisterFile("utility.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2d, 0xc9, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xcc, 0xb9,
	0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x74, 0xb8, 0x04, 0x5d, 0x52, 0x73, 0x12,
	0x2b, 0x7d, 0x33, 0x73, 0x72, 0x32, 0x83, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x8a, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0x04, 0x53, 0xd0, 0x25, 0x94, 0xf4, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x94, 0xb8, 0x78, 0x02, 0x32, 0xf3, 0xd2, 0x61, 0x7c, 0xb0, 0x26,
	0xce, 0x20, 0x9e, 0x02, 0x24, 0x31, 0x23, 0x0b, 0x2e, 0xf6, 0x50, 0x88, 0x9d, 0x42, 0xba, 0x5c,
	0x2c, 0x20, 0xe5, 0x42, 0x02, 0x7a, 0x30, 0x47, 0x41, 0x9d, 0x20, 0x25, 0x88, 0x24, 0x02, 0xd1,
	0xa7, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x47, 0x99, 0xb0,
	0x29, 0xc3, 0x00, 0x00, 0x00,
}