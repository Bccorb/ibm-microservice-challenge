// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/gts/gts.proto

package gts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GlobalTradeSystemClient is the client API for GlobalTradeSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalTradeSystemClient interface {
	// Get all avaliable matching trades
	GetTradeList(ctx context.Context, opts ...grpc.CallOption) (GlobalTradeSystem_GetTradeListClient, error)
}

type globalTradeSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalTradeSystemClient(cc grpc.ClientConnInterface) GlobalTradeSystemClient {
	return &globalTradeSystemClient{cc}
}

func (c *globalTradeSystemClient) GetTradeList(ctx context.Context, opts ...grpc.CallOption) (GlobalTradeSystem_GetTradeListClient, error) {
	stream, err := c.cc.NewStream(ctx, &GlobalTradeSystem_ServiceDesc.Streams[0], "/globalTradeSystem.GlobalTradeSystem/GetTradeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalTradeSystemGetTradeListClient{stream}
	return x, nil
}

type GlobalTradeSystem_GetTradeListClient interface {
	Send(*TradeRequest) error
	Recv() (*TradeResponse, error)
	grpc.ClientStream
}

type globalTradeSystemGetTradeListClient struct {
	grpc.ClientStream
}

func (x *globalTradeSystemGetTradeListClient) Send(m *TradeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *globalTradeSystemGetTradeListClient) Recv() (*TradeResponse, error) {
	m := new(TradeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalTradeSystemServer is the server API for GlobalTradeSystem service.
// All implementations must embed UnimplementedGlobalTradeSystemServer
// for forward compatibility
type GlobalTradeSystemServer interface {
	// Get all avaliable matching trades
	GetTradeList(GlobalTradeSystem_GetTradeListServer) error
	mustEmbedUnimplementedGlobalTradeSystemServer()
}

// UnimplementedGlobalTradeSystemServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalTradeSystemServer struct {
}

func (UnimplementedGlobalTradeSystemServer) GetTradeList(GlobalTradeSystem_GetTradeListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTradeList not implemented")
}
func (UnimplementedGlobalTradeSystemServer) mustEmbedUnimplementedGlobalTradeSystemServer() {}

// UnsafeGlobalTradeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalTradeSystemServer will
// result in compilation errors.
type UnsafeGlobalTradeSystemServer interface {
	mustEmbedUnimplementedGlobalTradeSystemServer()
}

func RegisterGlobalTradeSystemServer(s grpc.ServiceRegistrar, srv GlobalTradeSystemServer) {
	s.RegisterService(&GlobalTradeSystem_ServiceDesc, srv)
}

func _GlobalTradeSystem_GetTradeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GlobalTradeSystemServer).GetTradeList(&globalTradeSystemGetTradeListServer{stream})
}

type GlobalTradeSystem_GetTradeListServer interface {
	Send(*TradeResponse) error
	Recv() (*TradeRequest, error)
	grpc.ServerStream
}

type globalTradeSystemGetTradeListServer struct {
	grpc.ServerStream
}

func (x *globalTradeSystemGetTradeListServer) Send(m *TradeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *globalTradeSystemGetTradeListServer) Recv() (*TradeRequest, error) {
	m := new(TradeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalTradeSystem_ServiceDesc is the grpc.ServiceDesc for GlobalTradeSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalTradeSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "globalTradeSystem.GlobalTradeSystem",
	HandlerType: (*GlobalTradeSystemServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTradeList",
			Handler:       _GlobalTradeSystem_GetTradeList_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/gts/gts.proto",
}
