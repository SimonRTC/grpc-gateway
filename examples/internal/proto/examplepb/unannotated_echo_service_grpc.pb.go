// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/unannotated_echo_service.proto

// Unannotated Echo Service
// Similar to echo_service.proto but without annotations. See
// unannotated_echo_service.yaml for the equivalent of the annotations in
// gRPC API configuration format.
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UnannotatedEchoService_Echo_FullMethodName       = "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/Echo"
	UnannotatedEchoService_EchoBody_FullMethodName   = "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoBody"
	UnannotatedEchoService_EchoDelete_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoDelete"
	UnannotatedEchoService_EchoNested_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService/EchoNested"
)

// UnannotatedEchoServiceClient is the client API for UnannotatedEchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Echo service responds to incoming echo requests.
type UnannotatedEchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoNested method receives a simple message and returns it.
	EchoNested(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
}

type unannotatedEchoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnannotatedEchoServiceClient(cc grpc.ClientConnInterface) UnannotatedEchoServiceClient {
	return &unannotatedEchoServiceClient{cc}
}

func (c *unannotatedEchoServiceClient) Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, UnannotatedEchoService_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, UnannotatedEchoService_EchoBody_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, UnannotatedEchoService_EchoDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoNested(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, UnannotatedEchoService_EchoNested_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnannotatedEchoServiceServer is the server API for UnannotatedEchoService service.
// All implementations should embed UnimplementedUnannotatedEchoServiceServer
// for forward compatibility.
//
// Echo service responds to incoming echo requests.
type UnannotatedEchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoNested method receives a simple message and returns it.
	EchoNested(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
}

// UnimplementedUnannotatedEchoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUnannotatedEchoServiceServer struct{}

func (UnimplementedUnannotatedEchoServiceServer) Echo(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedUnannotatedEchoServiceServer) EchoBody(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (UnimplementedUnannotatedEchoServiceServer) EchoDelete(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}
func (UnimplementedUnannotatedEchoServiceServer) EchoNested(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoNested not implemented")
}
func (UnimplementedUnannotatedEchoServiceServer) testEmbeddedByValue() {}

// UnsafeUnannotatedEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnannotatedEchoServiceServer will
// result in compilation errors.
type UnsafeUnannotatedEchoServiceServer interface {
	mustEmbedUnimplementedUnannotatedEchoServiceServer()
}

func RegisterUnannotatedEchoServiceServer(s grpc.ServiceRegistrar, srv UnannotatedEchoServiceServer) {
	// If the following call pancis, it indicates UnimplementedUnannotatedEchoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UnannotatedEchoService_ServiceDesc, srv)
}

func _UnannotatedEchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnannotatedEchoService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnannotatedEchoService_EchoBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnannotatedEchoService_EchoDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoNested_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoNested(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnannotatedEchoService_EchoNested_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoNested(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// UnannotatedEchoService_ServiceDesc is the grpc.ServiceDesc for UnannotatedEchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnannotatedEchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.UnannotatedEchoService",
	HandlerType: (*UnannotatedEchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _UnannotatedEchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _UnannotatedEchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _UnannotatedEchoService_EchoDelete_Handler,
		},
		{
			MethodName: "EchoNested",
			Handler:    _UnannotatedEchoService_EchoNested_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/unannotated_echo_service.proto",
}
