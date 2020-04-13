// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-grpc-example/5-both_stream_rpc/proto/both_stream.proto

package proto

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

// 定义发送请求信息
type SimpleRequest struct {
	// 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
	// 参数类型 参数名 标识号(不可重复)
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92111b4df4e5dc8, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// 定义响应信息
type SimpleResponse struct {
	// 定义接收的参数
	// 参数类型 参数名 标识号(不可重复)
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92111b4df4e5dc8, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SimpleResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// 定义流式请求信息
type StreamRequest struct {
	//流式请求参数
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92111b4df4e5dc8, []int{2}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

// 定义流式响应信息
type StreamResponse struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e92111b4df4e5dc8, []int{3}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "proto.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "proto.SimpleResponse")
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.StreamResponse")
}

func init() {
	proto.RegisterFile("go-grpc-example/5-both_stream_rpc/proto/both_stream.proto", fileDescriptor_e92111b4df4e5dc8)
}

var fileDescriptor_e92111b4df4e5dc8 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x62, 0x82, 0x0e, 0xa4, 0x87, 0xa5, 0x4a, 0xe9, 0x49, 0xd6, 0x4b, 0x40, 0xd2,
	0x88, 0xa2, 0xa0, 0x47, 0x7d, 0x83, 0xf4, 0x01, 0xca, 0x36, 0x1d, 0x6a, 0xa1, 0xcd, 0xac, 0xbb,
	0x9b, 0xea, 0x03, 0xf8, 0xe0, 0x92, 0xd9, 0x8d, 0x6d, 0x3c, 0x65, 0xfe, 0x3f, 0xdf, 0xbf, 0xff,
	0x0c, 0xbc, 0x6c, 0xa8, 0xd8, 0x58, 0x53, 0x17, 0xf8, 0xad, 0xf7, 0x66, 0x87, 0xe5, 0x53, 0xb1,
	0x22, 0xff, 0xb1, 0x74, 0xde, 0xa2, 0xde, 0x2f, 0xad, 0xa9, 0x4b, 0x63, 0xc9, 0x53, 0x79, 0xe2,
	0xce, 0xd9, 0x91, 0x09, 0x7f, 0xd4, 0x2d, 0x64, 0x8b, 0x6d, 0x17, 0xad, 0xf0, 0xb3, 0x45, 0xe7,
	0xa5, 0x84, 0xf3, 0xb5, 0xf6, 0x7a, 0x2a, 0x6e, 0x44, 0x7e, 0x59, 0xf1, 0xac, 0x5e, 0x61, 0xdc,
	0x43, 0xce, 0x50, 0xe3, 0xb0, 0xa3, 0x6a, 0x5a, 0x23, 0x53, 0x49, 0xc5, 0xb3, 0x9c, 0x40, 0x72,
	0xd0, 0xbb, 0x16, 0xa7, 0x67, 0x1c, 0x0d, 0x42, 0xdd, 0x41, 0xb6, 0xe0, 0xde, 0xbe, 0x60, 0x06,
	0x17, 0x3c, 0x6c, 0xa9, 0x89, 0x25, 0x7f, 0x5a, 0xe5, 0x30, 0xee, 0xe1, 0x58, 0x74, 0x0d, 0xa9,
	0x6e, 0xdc, 0x17, 0xda, 0xc8, 0x46, 0xf5, 0xf0, 0x23, 0x20, 0x0d, 0xa8, 0x7c, 0x86, 0xa4, 0xa2,
	0xd6, 0xa3, 0x9c, 0x84, 0xd3, 0xe6, 0x83, 0x83, 0x66, 0x57, 0xff, 0xdc, 0xf0, 0xb0, 0x1a, 0xc9,
	0x37, 0xc8, 0xde, 0xa9, 0x39, 0xa0, 0x75, 0xba, 0x2b, 0x77, 0xc7, 0xfc, 0xe9, 0xbe, 0xc7, 0xfc,
	0x60, 0x31, 0x35, 0xca, 0xc5, 0xbd, 0x58, 0xa5, 0xfc, 0xef, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xec, 0xd8, 0x83, 0xe8, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// 双向流式rpc，同时在请求参数前和响应参数前加上stream
	Conversations(ctx context.Context, opts ...grpc.CallOption) (Stream_ConversationsClient, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.Stream/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClient) Conversations(ctx context.Context, opts ...grpc.CallOption) (Stream_ConversationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/proto.Stream/Conversations", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamConversationsClient{stream}
	return x, nil
}

type Stream_ConversationsClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamConversationsClient struct {
	grpc.ClientStream
}

func (x *streamConversationsClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamConversationsClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// 双向流式rpc，同时在请求参数前和响应参数前加上stream
	Conversations(Stream_ConversationsServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Route(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (*UnimplementedStreamServer) Conversations(srv Stream_ConversationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Conversations not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Stream/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Conversations_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).Conversations(&streamConversationsServer{stream})
}

type Stream_ConversationsServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamConversationsServer struct {
	grpc.ServerStream
}

func (x *streamConversationsServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamConversationsServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _Stream_Route_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Conversations",
			Handler:       _Stream_Conversations_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "go-grpc-example/5-both_stream_rpc/proto/both_stream.proto",
}
