// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: internal/protos/proto/bids/proto_bids_service_fetch.proto

package fetch

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
	BidsServiceFetch_FetchListByTender_FullMethodName = "/fetch.BidsServiceFetch/FetchListByTender"
	BidsServiceFetch_FetchListByUser_FullMethodName   = "/fetch.BidsServiceFetch/FetchListByUser"
	BidsServiceFetch_FetchStatus_FullMethodName       = "/fetch.BidsServiceFetch/FetchStatus"
	BidsServiceFetch_FetchReviews_FullMethodName      = "/fetch.BidsServiceFetch/FetchReviews"
)

// BidsServiceFetchClient is the client API for BidsServiceFetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidsServiceFetchClient interface {
	FetchListByTender(ctx context.Context, in *BidsRequestFetchListV1, opts ...grpc.CallOption) (*ResponseBidsV1, error)
	FetchListByUser(ctx context.Context, in *BidsRequestFetchListByUserV1, opts ...grpc.CallOption) (*ResponseBidsV1, error)
	FetchStatus(ctx context.Context, in *BidsRequestFetchStatusV1, opts ...grpc.CallOption) (*BidsResponseFetchStatusV1, error)
	FetchReviews(ctx context.Context, in *BidsRequestFetchReviewsV1, opts ...grpc.CallOption) (*BidsResponseFetchReviewsV1, error)
}

type bidsServiceFetchClient struct {
	cc grpc.ClientConnInterface
}

func NewBidsServiceFetchClient(cc grpc.ClientConnInterface) BidsServiceFetchClient {
	return &bidsServiceFetchClient{cc}
}

func (c *bidsServiceFetchClient) FetchListByTender(ctx context.Context, in *BidsRequestFetchListV1, opts ...grpc.CallOption) (*ResponseBidsV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBidsV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_FetchListByTender_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) FetchListByUser(ctx context.Context, in *BidsRequestFetchListByUserV1, opts ...grpc.CallOption) (*ResponseBidsV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBidsV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_FetchListByUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) FetchStatus(ctx context.Context, in *BidsRequestFetchStatusV1, opts ...grpc.CallOption) (*BidsResponseFetchStatusV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidsResponseFetchStatusV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_FetchStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) FetchReviews(ctx context.Context, in *BidsRequestFetchReviewsV1, opts ...grpc.CallOption) (*BidsResponseFetchReviewsV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidsResponseFetchReviewsV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_FetchReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidsServiceFetchServer is the server API for BidsServiceFetch service.
// All implementations must embed UnimplementedBidsServiceFetchServer
// for forward compatibility.
type BidsServiceFetchServer interface {
	FetchListByTender(context.Context, *BidsRequestFetchListV1) (*ResponseBidsV1, error)
	FetchListByUser(context.Context, *BidsRequestFetchListByUserV1) (*ResponseBidsV1, error)
	FetchStatus(context.Context, *BidsRequestFetchStatusV1) (*BidsResponseFetchStatusV1, error)
	FetchReviews(context.Context, *BidsRequestFetchReviewsV1) (*BidsResponseFetchReviewsV1, error)
	mustEmbedUnimplementedBidsServiceFetchServer()
}

// UnimplementedBidsServiceFetchServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBidsServiceFetchServer struct{}

func (UnimplementedBidsServiceFetchServer) FetchListByTender(context.Context, *BidsRequestFetchListV1) (*ResponseBidsV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListByTender not implemented")
}
func (UnimplementedBidsServiceFetchServer) FetchListByUser(context.Context, *BidsRequestFetchListByUserV1) (*ResponseBidsV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListByUser not implemented")
}
func (UnimplementedBidsServiceFetchServer) FetchStatus(context.Context, *BidsRequestFetchStatusV1) (*BidsResponseFetchStatusV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchStatus not implemented")
}
func (UnimplementedBidsServiceFetchServer) FetchReviews(context.Context, *BidsRequestFetchReviewsV1) (*BidsResponseFetchReviewsV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchReviews not implemented")
}
func (UnimplementedBidsServiceFetchServer) mustEmbedUnimplementedBidsServiceFetchServer() {}
func (UnimplementedBidsServiceFetchServer) testEmbeddedByValue()                          {}

// UnsafeBidsServiceFetchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidsServiceFetchServer will
// result in compilation errors.
type UnsafeBidsServiceFetchServer interface {
	mustEmbedUnimplementedBidsServiceFetchServer()
}

func RegisterBidsServiceFetchServer(s grpc.ServiceRegistrar, srv BidsServiceFetchServer) {
	// If the following call pancis, it indicates UnimplementedBidsServiceFetchServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BidsServiceFetch_ServiceDesc, srv)
}

func _BidsServiceFetch_FetchListByTender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestFetchListV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).FetchListByTender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_FetchListByTender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).FetchListByTender(ctx, req.(*BidsRequestFetchListV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_FetchListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestFetchListByUserV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).FetchListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_FetchListByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).FetchListByUser(ctx, req.(*BidsRequestFetchListByUserV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_FetchStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestFetchStatusV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).FetchStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_FetchStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).FetchStatus(ctx, req.(*BidsRequestFetchStatusV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_FetchReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestFetchReviewsV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).FetchReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_FetchReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).FetchReviews(ctx, req.(*BidsRequestFetchReviewsV1))
	}
	return interceptor(ctx, in, info, handler)
}

// BidsServiceFetch_ServiceDesc is the grpc.ServiceDesc for BidsServiceFetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidsServiceFetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetch.BidsServiceFetch",
	HandlerType: (*BidsServiceFetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListByTender",
			Handler:    _BidsServiceFetch_FetchListByTender_Handler,
		},
		{
			MethodName: "FetchListByUser",
			Handler:    _BidsServiceFetch_FetchListByUser_Handler,
		},
		{
			MethodName: "FetchStatus",
			Handler:    _BidsServiceFetch_FetchStatus_Handler,
		},
		{
			MethodName: "FetchReviews",
			Handler:    _BidsServiceFetch_FetchReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protos/proto/bids/proto_bids_service_fetch.proto",
}
