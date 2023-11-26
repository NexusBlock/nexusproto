// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: odin/odin.proto

package odin

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

// OdinClient is the client API for Odin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OdinClient interface {
	Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error)
	StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Odin_StateSyncEventsClient, error)
	FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error)
}

type odinClient struct {
	cc grpc.ClientConnInterface
}

func NewOdinClient(cc grpc.ClientConnInterface) OdinClient {
	return &odinClient{cc}
}

func (c *odinClient) Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error) {
	out := new(SpanResponse)
	err := c.cc.Invoke(ctx, "/odin.Odin/Span", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *odinClient) StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Odin_StateSyncEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Odin_ServiceDesc.Streams[0], "/odin.Odin/StateSyncEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &odinStateSyncEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Odin_StateSyncEventsClient interface {
	Recv() (*StateSyncEventsResponse, error)
	grpc.ClientStream
}

type odinStateSyncEventsClient struct {
	grpc.ClientStream
}

func (x *odinStateSyncEventsClient) Recv() (*StateSyncEventsResponse, error) {
	m := new(StateSyncEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *odinClient) FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error) {
	out := new(FetchCheckpointResponse)
	err := c.cc.Invoke(ctx, "/odin.Odin/FetchCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *odinClient) FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error) {
	out := new(FetchCheckpointCountResponse)
	err := c.cc.Invoke(ctx, "/odin.Odin/FetchCheckpointCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OdinServer is the server API for Odin service.
// All implementations must embed UnimplementedOdinServer
// for forward compatibility
type OdinServer interface {
	Span(context.Context, *SpanRequest) (*SpanResponse, error)
	StateSyncEvents(*StateSyncEventsRequest, Odin_StateSyncEventsServer) error
	FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error)
	mustEmbedUnimplementedOdinServer()
}

// UnimplementedOdinServer must be embedded to have forward compatible implementations.
type UnimplementedOdinServer struct {
}

func (UnimplementedOdinServer) Span(context.Context, *SpanRequest) (*SpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Span not implemented")
}
func (UnimplementedOdinServer) StateSyncEvents(*StateSyncEventsRequest, Odin_StateSyncEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StateSyncEvents not implemented")
}
func (UnimplementedOdinServer) FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpoint not implemented")
}
func (UnimplementedOdinServer) FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpointCount not implemented")
}
func (UnimplementedOdinServer) mustEmbedUnimplementedOdinServer() {}

// UnsafeOdinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OdinServer will
// result in compilation errors.
type UnsafeOdinServer interface {
	mustEmbedUnimplementedOdinServer()
}

func RegisterOdinServer(s grpc.ServiceRegistrar, srv OdinServer) {
	s.RegisterService(&Odin_ServiceDesc, srv)
}

func _Odin_Span_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OdinServer).Span(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odin.Odin/Span",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OdinServer).Span(ctx, req.(*SpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Odin_StateSyncEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateSyncEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OdinServer).StateSyncEvents(m, &odinStateSyncEventsServer{stream})
}

type Odin_StateSyncEventsServer interface {
	Send(*StateSyncEventsResponse) error
	grpc.ServerStream
}

type odinStateSyncEventsServer struct {
	grpc.ServerStream
}

func (x *odinStateSyncEventsServer) Send(m *StateSyncEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Odin_FetchCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OdinServer).FetchCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odin.Odin/FetchCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OdinServer).FetchCheckpoint(ctx, req.(*FetchCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Odin_FetchCheckpointCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OdinServer).FetchCheckpointCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/odin.Odin/FetchCheckpointCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OdinServer).FetchCheckpointCount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Odin_ServiceDesc is the grpc.ServiceDesc for Odin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Odin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "odin.Odin",
	HandlerType: (*OdinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Span",
			Handler:    _Odin_Span_Handler,
		},
		{
			MethodName: "FetchCheckpoint",
			Handler:    _Odin_FetchCheckpoint_Handler,
		},
		{
			MethodName: "FetchCheckpointCount",
			Handler:    _Odin_FetchCheckpointCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StateSyncEvents",
			Handler:       _Odin_StateSyncEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "odin/odin.proto",
}
