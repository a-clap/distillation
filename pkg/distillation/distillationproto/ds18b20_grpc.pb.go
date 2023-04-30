// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pkg/distillation/distillationproto/ds18b20.proto

package distillationproto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DSClient is the client API for DS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DSClient interface {
	DSGet(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DSConfigs, error)
	DSConfigure(ctx context.Context, in *DSConfig, opts ...grpc.CallOption) (*DSConfig, error)
	DSGetTemperatures(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DSTemperatures, error)
}

type dSClient struct {
	cc grpc.ClientConnInterface
}

func NewDSClient(cc grpc.ClientConnInterface) DSClient {
	return &dSClient{cc}
}

func (c *dSClient) DSGet(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DSConfigs, error) {
	out := new(DSConfigs)
	err := c.cc.Invoke(ctx, "/distillationproto.DS/DSGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSClient) DSConfigure(ctx context.Context, in *DSConfig, opts ...grpc.CallOption) (*DSConfig, error) {
	out := new(DSConfig)
	err := c.cc.Invoke(ctx, "/distillationproto.DS/DSConfigure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSClient) DSGetTemperatures(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DSTemperatures, error) {
	out := new(DSTemperatures)
	err := c.cc.Invoke(ctx, "/distillationproto.DS/DSGetTemperatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DSServer is the server API for DS service.
// All implementations must embed UnimplementedDSServer
// for forward compatibility
type DSServer interface {
	DSGet(context.Context, *empty.Empty) (*DSConfigs, error)
	DSConfigure(context.Context, *DSConfig) (*DSConfig, error)
	DSGetTemperatures(context.Context, *empty.Empty) (*DSTemperatures, error)
	mustEmbedUnimplementedDSServer()
}

// UnimplementedDSServer must be embedded to have forward compatible implementations.
type UnimplementedDSServer struct {
}

func (UnimplementedDSServer) DSGet(context.Context, *empty.Empty) (*DSConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DSGet not implemented")
}
func (UnimplementedDSServer) DSConfigure(context.Context, *DSConfig) (*DSConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DSConfigure not implemented")
}
func (UnimplementedDSServer) DSGetTemperatures(context.Context, *empty.Empty) (*DSTemperatures, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DSGetTemperatures not implemented")
}
func (UnimplementedDSServer) mustEmbedUnimplementedDSServer() {}

// UnsafeDSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DSServer will
// result in compilation errors.
type UnsafeDSServer interface {
	mustEmbedUnimplementedDSServer()
}

func RegisterDSServer(s grpc.ServiceRegistrar, srv DSServer) {
	s.RegisterService(&DS_ServiceDesc, srv)
}

func _DS_DSGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSServer).DSGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distillationproto.DS/DSGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSServer).DSGet(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DS_DSConfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DSConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSServer).DSConfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distillationproto.DS/DSConfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSServer).DSConfigure(ctx, req.(*DSConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DS_DSGetTemperatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSServer).DSGetTemperatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distillationproto.DS/DSGetTemperatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSServer).DSGetTemperatures(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DS_ServiceDesc is the grpc.ServiceDesc for DS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distillationproto.DS",
	HandlerType: (*DSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DSGet",
			Handler:    _DS_DSGet_Handler,
		},
		{
			MethodName: "DSConfigure",
			Handler:    _DS_DSConfigure_Handler,
		},
		{
			MethodName: "DSGetTemperatures",
			Handler:    _DS_DSGetTemperatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/distillation/distillationproto/ds18b20.proto",
}
