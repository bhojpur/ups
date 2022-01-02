// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// UpsUIClient is the client API for UpsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpsUIClient interface {
	// ListInverterSpecs returns a list of Inverter(s) that can be started through the UI.
	ListInverterSpecs(ctx context.Context, in *ListInverterSpecsRequest, opts ...grpc.CallOption) (UpsUI_ListInverterSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type upsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewUpsUIClient(cc grpc.ClientConnInterface) UpsUIClient {
	return &upsUIClient{cc}
}

func (c *upsUIClient) ListInverterSpecs(ctx context.Context, in *ListInverterSpecsRequest, opts ...grpc.CallOption) (UpsUI_ListInverterSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UpsUI_ServiceDesc.Streams[0], "/v1.UpsUI/ListInverterSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &upsUIListInverterSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpsUI_ListInverterSpecsClient interface {
	Recv() (*ListInverterSpecsResponse, error)
	grpc.ClientStream
}

type upsUIListInverterSpecsClient struct {
	grpc.ClientStream
}

func (x *upsUIListInverterSpecsClient) Recv() (*ListInverterSpecsResponse, error) {
	m := new(ListInverterSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *upsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.UpsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpsUIServer is the server API for UpsUI service.
// All implementations must embed UnimplementedUpsUIServer
// for forward compatibility
type UpsUIServer interface {
	// ListInverterSpecs returns a list of Inverter(s) that can be started through the UI.
	ListInverterSpecs(*ListInverterSpecsRequest, UpsUI_ListInverterSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedUpsUIServer()
}

// UnimplementedUpsUIServer must be embedded to have forward compatible implementations.
type UnimplementedUpsUIServer struct {
}

func (UnimplementedUpsUIServer) ListInverterSpecs(*ListInverterSpecsRequest, UpsUI_ListInverterSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListInverterSpecs not implemented")
}
func (UnimplementedUpsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedUpsUIServer) mustEmbedUnimplementedUpsUIServer() {}

// UnsafeUpsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpsUIServer will
// result in compilation errors.
type UnsafeUpsUIServer interface {
	mustEmbedUnimplementedUpsUIServer()
}

func RegisterUpsUIServer(s grpc.ServiceRegistrar, srv UpsUIServer) {
	s.RegisterService(&UpsUI_ServiceDesc, srv)
}

func _UpsUI_ListInverterSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInverterSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpsUIServer).ListInverterSpecs(m, &upsUIListInverterSpecsServer{stream})
}

type UpsUI_ListInverterSpecsServer interface {
	Send(*ListInverterSpecsResponse) error
	grpc.ServerStream
}

type upsUIListInverterSpecsServer struct {
	grpc.ServerStream
}

func (x *upsUIListInverterSpecsServer) Send(m *ListInverterSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UpsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpsUI_ServiceDesc is the grpc.ServiceDesc for UpsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UpsUI",
	HandlerType: (*UpsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _UpsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListInverterSpecs",
			Handler:       _UpsUI_ListInverterSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ups-ui.proto",
}
