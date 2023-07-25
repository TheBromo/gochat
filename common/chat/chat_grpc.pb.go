// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0--rc1
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Register(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Login, error)
	Unregister(ctx context.Context, in *Logout, opts ...grpc.CallOption) (*Logout, error)
	HandleMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_HandleMessageClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Register(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Unregister(ctx context.Context, in *Logout, opts ...grpc.CallOption) (*Logout, error) {
	out := new(Logout)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) HandleMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_HandleMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chat.ChatService/HandleMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceHandleMessageClient{stream}
	return x, nil
}

type ChatService_HandleMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type chatServiceHandleMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceHandleMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceHandleMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Register(context.Context, *Login) (*Login, error)
	Unregister(context.Context, *Logout) (*Logout, error)
	HandleMessage(ChatService_HandleMessageServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Register(context.Context, *Login) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChatServiceServer) Unregister(context.Context, *Logout) (*Logout, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedChatServiceServer) HandleMessage(ChatService_HandleMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Register(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Logout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Unregister(ctx, req.(*Logout))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_HandleMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).HandleMessage(&chatServiceHandleMessageServer{stream})
}

type ChatService_HandleMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceHandleMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceHandleMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceHandleMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ChatService_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _ChatService_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleMessage",
			Handler:       _ChatService_HandleMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

// TerminalAppServiceClient is the client API for TerminalAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TerminalAppServiceClient interface {
	HandleMessage(ctx context.Context, opts ...grpc.CallOption) (TerminalAppService_HandleMessageClient, error)
}

type terminalAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTerminalAppServiceClient(cc grpc.ClientConnInterface) TerminalAppServiceClient {
	return &terminalAppServiceClient{cc}
}

func (c *terminalAppServiceClient) HandleMessage(ctx context.Context, opts ...grpc.CallOption) (TerminalAppService_HandleMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &TerminalAppService_ServiceDesc.Streams[0], "/chat.TerminalAppService/HandleMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &terminalAppServiceHandleMessageClient{stream}
	return x, nil
}

type TerminalAppService_HandleMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type terminalAppServiceHandleMessageClient struct {
	grpc.ClientStream
}

func (x *terminalAppServiceHandleMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *terminalAppServiceHandleMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TerminalAppServiceServer is the server API for TerminalAppService service.
// All implementations must embed UnimplementedTerminalAppServiceServer
// for forward compatibility
type TerminalAppServiceServer interface {
	HandleMessage(TerminalAppService_HandleMessageServer) error
	mustEmbedUnimplementedTerminalAppServiceServer()
}

// UnimplementedTerminalAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTerminalAppServiceServer struct {
}

func (UnimplementedTerminalAppServiceServer) HandleMessage(TerminalAppService_HandleMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleMessage not implemented")
}
func (UnimplementedTerminalAppServiceServer) mustEmbedUnimplementedTerminalAppServiceServer() {}

// UnsafeTerminalAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TerminalAppServiceServer will
// result in compilation errors.
type UnsafeTerminalAppServiceServer interface {
	mustEmbedUnimplementedTerminalAppServiceServer()
}

func RegisterTerminalAppServiceServer(s grpc.ServiceRegistrar, srv TerminalAppServiceServer) {
	s.RegisterService(&TerminalAppService_ServiceDesc, srv)
}

func _TerminalAppService_HandleMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TerminalAppServiceServer).HandleMessage(&terminalAppServiceHandleMessageServer{stream})
}

type TerminalAppService_HandleMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type terminalAppServiceHandleMessageServer struct {
	grpc.ServerStream
}

func (x *terminalAppServiceHandleMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *terminalAppServiceHandleMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TerminalAppService_ServiceDesc is the grpc.ServiceDesc for TerminalAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TerminalAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.TerminalAppService",
	HandlerType: (*TerminalAppServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleMessage",
			Handler:       _TerminalAppService_HandleMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
