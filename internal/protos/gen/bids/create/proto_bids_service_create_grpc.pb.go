// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: internal/protos/proto/bids/proto_bids_service_create.proto

package create

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
	BidsServiceCreate_Create_FullMethodName   = "/create.BidsServiceCreate/Create"
	BidsServiceCreate_Feedback_FullMethodName = "/create.BidsServiceCreate/Feedback"
)

// BidsServiceCreateClient is the client API for BidsServiceCreate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidsServiceCreateClient interface {
	Create(ctx context.Context, in *BidRequestCreateV1, opts ...grpc.CallOption) (*BidV1, error)
	Feedback(ctx context.Context, in *BidRequestFeedbackV1, opts ...grpc.CallOption) (*BidFeedbackV1, error)
}

type bidsServiceCreateClient struct {
	cc grpc.ClientConnInterface
}

func NewBidsServiceCreateClient(cc grpc.ClientConnInterface) BidsServiceCreateClient {
	return &bidsServiceCreateClient{cc}
}

func (c *bidsServiceCreateClient) Create(ctx context.Context, in *BidRequestCreateV1, opts ...grpc.CallOption) (*BidV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidV1)
	err := c.cc.Invoke(ctx, BidsServiceCreate_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceCreateClient) Feedback(ctx context.Context, in *BidRequestFeedbackV1, opts ...grpc.CallOption) (*BidFeedbackV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidFeedbackV1)
	err := c.cc.Invoke(ctx, BidsServiceCreate_Feedback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidsServiceCreateServer is the server API for BidsServiceCreate service.
// All implementations must embed UnimplementedBidsServiceCreateServer
// for forward compatibility.
type BidsServiceCreateServer interface {
	Create(context.Context, *BidRequestCreateV1) (*BidV1, error)
	Feedback(context.Context, *BidRequestFeedbackV1) (*BidFeedbackV1, error)
	mustEmbedUnimplementedBidsServiceCreateServer()
}

// UnimplementedBidsServiceCreateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBidsServiceCreateServer struct{}

func (UnimplementedBidsServiceCreateServer) Create(context.Context, *BidRequestCreateV1) (*BidV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBidsServiceCreateServer) Feedback(context.Context, *BidRequestFeedbackV1) (*BidFeedbackV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feedback not implemented")
}
func (UnimplementedBidsServiceCreateServer) mustEmbedUnimplementedBidsServiceCreateServer() {}
func (UnimplementedBidsServiceCreateServer) testEmbeddedByValue()                           {}

// UnsafeBidsServiceCreateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidsServiceCreateServer will
// result in compilation errors.
type UnsafeBidsServiceCreateServer interface {
	mustEmbedUnimplementedBidsServiceCreateServer()
}

func RegisterBidsServiceCreateServer(s grpc.ServiceRegistrar, srv BidsServiceCreateServer) {
	// If the following call pancis, it indicates UnimplementedBidsServiceCreateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BidsServiceCreate_ServiceDesc, srv)
}

func _BidsServiceCreate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequestCreateV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceCreateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceCreate_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceCreateServer).Create(ctx, req.(*BidRequestCreateV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceCreate_Feedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequestFeedbackV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceCreateServer).Feedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceCreate_Feedback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceCreateServer).Feedback(ctx, req.(*BidRequestFeedbackV1))
	}
	return interceptor(ctx, in, info, handler)
}

// BidsServiceCreate_ServiceDesc is the grpc.ServiceDesc for BidsServiceCreate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidsServiceCreate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "create.BidsServiceCreate",
	HandlerType: (*BidsServiceCreateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BidsServiceCreate_Create_Handler,
		},
		{
			MethodName: "Feedback",
			Handler:    _BidsServiceCreate_Feedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protos/proto/bids/proto_bids_service_create.proto",
}
