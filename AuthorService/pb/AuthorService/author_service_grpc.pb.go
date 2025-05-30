// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/AuthorService/author_service.proto

package AuthorService

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
	Author_RegisterUserAsAuthor_FullMethodName = "/authorpb.Author/RegisterUserAsAuthor"
	Author_GetAuthorIdByUserId_FullMethodName  = "/authorpb.Author/GetAuthorIdByUserId"
)

// AuthorClient is the client API for Author service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorClient interface {
	RegisterUserAsAuthor(ctx context.Context, in *RegisterUserAsAuthorRequest, opts ...grpc.CallOption) (*RegisterUserAsAuthorResponse, error)
	GetAuthorIdByUserId(ctx context.Context, in *GetAuthorIdByUserIdRequest, opts ...grpc.CallOption) (*GetAuthorIdByUserIdResponse, error)
}

type authorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorClient(cc grpc.ClientConnInterface) AuthorClient {
	return &authorClient{cc}
}

func (c *authorClient) RegisterUserAsAuthor(ctx context.Context, in *RegisterUserAsAuthorRequest, opts ...grpc.CallOption) (*RegisterUserAsAuthorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserAsAuthorResponse)
	err := c.cc.Invoke(ctx, Author_RegisterUserAsAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorClient) GetAuthorIdByUserId(ctx context.Context, in *GetAuthorIdByUserIdRequest, opts ...grpc.CallOption) (*GetAuthorIdByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuthorIdByUserIdResponse)
	err := c.cc.Invoke(ctx, Author_GetAuthorIdByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServer is the server API for Author service.
// All implementations must embed UnimplementedAuthorServer
// for forward compatibility.
type AuthorServer interface {
	RegisterUserAsAuthor(context.Context, *RegisterUserAsAuthorRequest) (*RegisterUserAsAuthorResponse, error)
	GetAuthorIdByUserId(context.Context, *GetAuthorIdByUserIdRequest) (*GetAuthorIdByUserIdResponse, error)
	mustEmbedUnimplementedAuthorServer()
}

// UnimplementedAuthorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthorServer struct{}

func (UnimplementedAuthorServer) RegisterUserAsAuthor(context.Context, *RegisterUserAsAuthorRequest) (*RegisterUserAsAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserAsAuthor not implemented")
}
func (UnimplementedAuthorServer) GetAuthorIdByUserId(context.Context, *GetAuthorIdByUserIdRequest) (*GetAuthorIdByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorIdByUserId not implemented")
}
func (UnimplementedAuthorServer) mustEmbedUnimplementedAuthorServer() {}
func (UnimplementedAuthorServer) testEmbeddedByValue()                {}

// UnsafeAuthorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServer will
// result in compilation errors.
type UnsafeAuthorServer interface {
	mustEmbedUnimplementedAuthorServer()
}

func RegisterAuthorServer(s grpc.ServiceRegistrar, srv AuthorServer) {
	// If the following call pancis, it indicates UnimplementedAuthorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Author_ServiceDesc, srv)
}

func _Author_RegisterUserAsAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserAsAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).RegisterUserAsAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Author_RegisterUserAsAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).RegisterUserAsAuthor(ctx, req.(*RegisterUserAsAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Author_GetAuthorIdByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorIdByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).GetAuthorIdByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Author_GetAuthorIdByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).GetAuthorIdByUserId(ctx, req.(*GetAuthorIdByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Author_ServiceDesc is the grpc.ServiceDesc for Author service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Author_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorpb.Author",
	HandlerType: (*AuthorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUserAsAuthor",
			Handler:    _Author_RegisterUserAsAuthor_Handler,
		},
		{
			MethodName: "GetAuthorIdByUserId",
			Handler:    _Author_GetAuthorIdByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/AuthorService/author_service.proto",
}
