package command

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReverseServiceClient is the client API for ReverseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReverseServiceClient interface {
	AddBridge(ctx context.Context, in *AddBridgeRequest, opts ...grpc.CallOption) (*AddBridgeResponse, error)
	RemoveBridge(ctx context.Context, in *RemoveBridgeRequest, opts ...grpc.CallOption) (*RemoveBridgeResponse, error)
	AddPortal(ctx context.Context, in *AddPortalRequest, opts ...grpc.CallOption) (*AddPortalResponse, error)
	RemovePortal(ctx context.Context, in *RemovePortalRequest, opts ...grpc.CallOption) (*RemovePortalResponse, error)
}

type reverseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReverseServiceClient(cc grpc.ClientConnInterface) ReverseServiceClient {
	return &reverseServiceClient{cc}
}

func (c *reverseServiceClient) AddBridge(ctx context.Context, in *AddBridgeRequest, opts ...grpc.CallOption) (*AddBridgeResponse, error) {
	out := new(AddBridgeResponse)
	err := c.cc.Invoke(ctx, "/xray.app.reverse.command.ReverseService/AddBridge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseServiceClient) RemoveBridge(ctx context.Context, in *RemoveBridgeRequest, opts ...grpc.CallOption) (*RemoveBridgeResponse, error) {
	out := new(RemoveBridgeResponse)
	err := c.cc.Invoke(ctx, "/xray.app.reverse.command.ReverseService/RemoveBridge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseServiceClient) AddPortal(ctx context.Context, in *AddPortalRequest, opts ...grpc.CallOption) (*AddPortalResponse, error) {
	out := new(AddPortalResponse)
	err := c.cc.Invoke(ctx, "/xray.app.reverse.command.ReverseService/AddPortal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reverseServiceClient) RemovePortal(ctx context.Context, in *RemovePortalRequest, opts ...grpc.CallOption) (*RemovePortalResponse, error) {
	out := new(RemovePortalResponse)
	err := c.cc.Invoke(ctx, "/xray.app.reverse.command.ReverseService/RemovePortal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseServiceServer is the server API for ReverseService service.
type ReverseServiceServer interface {
	AddBridge(context.Context, *AddBridgeRequest) (*AddBridgeResponse, error)
	RemoveBridge(context.Context, *RemoveBridgeRequest) (*RemoveBridgeResponse, error)
	AddPortal(context.Context, *AddPortalRequest) (*AddPortalResponse, error)
	RemovePortal(context.Context, *RemovePortalRequest) (*RemovePortalResponse, error)
}

// UnimplementedReverseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReverseServiceServer struct {
}

func (*UnimplementedReverseServiceServer) AddBridge(context.Context, *AddBridgeRequest) (*AddBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBridge not implemented")
}
func (*UnimplementedReverseServiceServer) RemoveBridge(context.Context, *RemoveBridgeRequest) (*RemoveBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBridge not implemented")
}
func (*UnimplementedReverseServiceServer) AddPortal(context.Context, *AddPortalRequest) (*AddPortalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPortal not implemented")
}
func (*UnimplementedReverseServiceServer) RemovePortal(context.Context, *RemovePortalRequest) (*RemovePortalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePortal not implemented")
}

func RegisterReverseServiceServer(s *grpc.Server, srv ReverseServiceServer) {
	s.RegisterService(&ReverseService_ServiceDesc, srv)
}

func _ReverseService_AddBridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBridgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServiceServer).AddBridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xray.app.reverse.command.ReverseService/AddBridge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServiceServer).AddBridge(ctx, req.(*AddBridgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReverseService_RemoveBridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBridgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServiceServer).RemoveBridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xray.app.reverse.command.ReverseService/RemoveBridge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServiceServer).RemoveBridge(ctx, req.(*RemoveBridgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReverseService_AddPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServiceServer).AddPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xray.app.reverse.command.ReverseService/AddPortal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServiceServer).AddPortal(ctx, req.(*AddPortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReverseService_RemovePortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServiceServer).RemovePortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xray.app.reverse.command.ReverseService/RemovePortal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServiceServer).RemovePortal(ctx, req.(*RemovePortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var ReverseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xray.app.reverse.command.ReverseService",
	HandlerType: (*ReverseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBridge",
			Handler:    _ReverseService_AddBridge_Handler,
		},
		{
			MethodName: "RemoveBridge",
			Handler:    _ReverseService_RemoveBridge_Handler,
		},
		{
			MethodName: "AddPortal",
			Handler:    _ReverseService_AddPortal_Handler,
		},
		{
			MethodName: "RemovePortal",
			Handler:    _ReverseService_RemovePortal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/reverse/command/command.proto",
}
