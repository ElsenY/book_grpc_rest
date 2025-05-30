// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/BookService/book_service.proto

package BookService

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
	Book_InsertBook_FullMethodName       = "/bookpb.Book/InsertBook"
	Book_BorrowBook_FullMethodName       = "/bookpb.Book/BorrowBook"
	Book_ReturnBook_FullMethodName       = "/bookpb.Book/ReturnBook"
	Book_GetBookIdByTitle_FullMethodName = "/bookpb.Book/GetBookIdByTitle"
	Book_RecommendBook_FullMethodName    = "/bookpb.Book/RecommendBook"
	Book_SearchBook_FullMethodName       = "/bookpb.Book/SearchBook"
	Book_EditBookStock_FullMethodName    = "/bookpb.Book/EditBookStock"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	InsertBook(ctx context.Context, in *InsertBookRequest, opts ...grpc.CallOption) (*InsertBookResponse, error)
	BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error)
	ReturnBook(ctx context.Context, in *ReturnBookRequest, opts ...grpc.CallOption) (*ReturnBookResponse, error)
	GetBookIdByTitle(ctx context.Context, in *GetBookIdByTitleRequest, opts ...grpc.CallOption) (*GetBookIdByTitleResponse, error)
	RecommendBook(ctx context.Context, in *RecommendBookRequest, opts ...grpc.CallOption) (*RecommendBookResponse, error)
	SearchBook(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error)
	EditBookStock(ctx context.Context, in *EditBookStockRequest, opts ...grpc.CallOption) (*EditBookStockResponse, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) InsertBook(ctx context.Context, in *InsertBookRequest, opts ...grpc.CallOption) (*InsertBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertBookResponse)
	err := c.cc.Invoke(ctx, Book_InsertBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowBookResponse)
	err := c.cc.Invoke(ctx, Book_BorrowBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) ReturnBook(ctx context.Context, in *ReturnBookRequest, opts ...grpc.CallOption) (*ReturnBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReturnBookResponse)
	err := c.cc.Invoke(ctx, Book_ReturnBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookIdByTitle(ctx context.Context, in *GetBookIdByTitleRequest, opts ...grpc.CallOption) (*GetBookIdByTitleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookIdByTitleResponse)
	err := c.cc.Invoke(ctx, Book_GetBookIdByTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) RecommendBook(ctx context.Context, in *RecommendBookRequest, opts ...grpc.CallOption) (*RecommendBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecommendBookResponse)
	err := c.cc.Invoke(ctx, Book_RecommendBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) SearchBook(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchBookResponse)
	err := c.cc.Invoke(ctx, Book_SearchBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) EditBookStock(ctx context.Context, in *EditBookStockRequest, opts ...grpc.CallOption) (*EditBookStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditBookStockResponse)
	err := c.cc.Invoke(ctx, Book_EditBookStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility.
type BookServer interface {
	InsertBook(context.Context, *InsertBookRequest) (*InsertBookResponse, error)
	BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error)
	ReturnBook(context.Context, *ReturnBookRequest) (*ReturnBookResponse, error)
	GetBookIdByTitle(context.Context, *GetBookIdByTitleRequest) (*GetBookIdByTitleResponse, error)
	RecommendBook(context.Context, *RecommendBookRequest) (*RecommendBookResponse, error)
	SearchBook(context.Context, *SearchBookRequest) (*SearchBookResponse, error)
	EditBookStock(context.Context, *EditBookStockRequest) (*EditBookStockResponse, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServer struct{}

func (UnimplementedBookServer) InsertBook(context.Context, *InsertBookRequest) (*InsertBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
}
func (UnimplementedBookServer) BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowBook not implemented")
}
func (UnimplementedBookServer) ReturnBook(context.Context, *ReturnBookRequest) (*ReturnBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnBook not implemented")
}
func (UnimplementedBookServer) GetBookIdByTitle(context.Context, *GetBookIdByTitleRequest) (*GetBookIdByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookIdByTitle not implemented")
}
func (UnimplementedBookServer) RecommendBook(context.Context, *RecommendBookRequest) (*RecommendBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendBook not implemented")
}
func (UnimplementedBookServer) SearchBook(context.Context, *SearchBookRequest) (*SearchBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBook not implemented")
}
func (UnimplementedBookServer) EditBookStock(context.Context, *EditBookStockRequest) (*EditBookStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditBookStock not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}
func (UnimplementedBookServer) testEmbeddedByValue()              {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	// If the following call pancis, it indicates UnimplementedBookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_InsertBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).InsertBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_InsertBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).InsertBook(ctx, req.(*InsertBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_BorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).BorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_BorrowBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).BorrowBook(ctx, req.(*BorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_ReturnBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).ReturnBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_ReturnBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).ReturnBook(ctx, req.(*ReturnBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookIdByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookIdByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookIdByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookIdByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookIdByTitle(ctx, req.(*GetBookIdByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_RecommendBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).RecommendBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_RecommendBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).RecommendBook(ctx, req.(*RecommendBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_SearchBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).SearchBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_SearchBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).SearchBook(ctx, req.(*SearchBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_EditBookStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditBookStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).EditBookStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_EditBookStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).EditBookStock(ctx, req.(*EditBookStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookpb.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBook",
			Handler:    _Book_InsertBook_Handler,
		},
		{
			MethodName: "BorrowBook",
			Handler:    _Book_BorrowBook_Handler,
		},
		{
			MethodName: "ReturnBook",
			Handler:    _Book_ReturnBook_Handler,
		},
		{
			MethodName: "GetBookIdByTitle",
			Handler:    _Book_GetBookIdByTitle_Handler,
		},
		{
			MethodName: "RecommendBook",
			Handler:    _Book_RecommendBook_Handler,
		},
		{
			MethodName: "SearchBook",
			Handler:    _Book_SearchBook_Handler,
		},
		{
			MethodName: "EditBookStock",
			Handler:    _Book_EditBookStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/BookService/book_service.proto",
}
