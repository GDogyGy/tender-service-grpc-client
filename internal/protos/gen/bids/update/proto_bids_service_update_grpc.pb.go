// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: internal/protos/proto/bids/proto_bids_service_update.proto

package update

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
	BidsServiceFetch_Edit_FullMethodName           = "/update.BidsServiceFetch/Edit"
	BidsServiceFetch_Rollback_FullMethodName       = "/update.BidsServiceFetch/Rollback"
	BidsServiceFetch_Status_FullMethodName         = "/update.BidsServiceFetch/Status"
	BidsServiceFetch_SubmitDecision_FullMethodName = "/update.BidsServiceFetch/SubmitDecision"
)

// BidsServiceFetchClient is the client API for BidsServiceFetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidsServiceFetchClient interface {
	Edit(ctx context.Context, in *BidsRequestEditV1, opts ...grpc.CallOption) (*BidEditV1, error)
	Rollback(ctx context.Context, in *BidsRequestRollbackV1, opts ...grpc.CallOption) (*BidEditV1, error)
	Status(ctx context.Context, in *BidsRequestStatusV1, opts ...grpc.CallOption) (*BidEditV1, error)
	SubmitDecision(ctx context.Context, in *BidsRequestSubmitDecisionV1, opts ...grpc.CallOption) (*BidEditV1, error)
}

type bidsServiceFetchClient struct {
	cc grpc.ClientConnInterface
}

func NewBidsServiceFetchClient(cc grpc.ClientConnInterface) BidsServiceFetchClient {
	return &bidsServiceFetchClient{cc}
}

func (c *bidsServiceFetchClient) Edit(ctx context.Context, in *BidsRequestEditV1, opts ...grpc.CallOption) (*BidEditV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidEditV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_Edit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) Rollback(ctx context.Context, in *BidsRequestRollbackV1, opts ...grpc.CallOption) (*BidEditV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidEditV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_Rollback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) Status(ctx context.Context, in *BidsRequestStatusV1, opts ...grpc.CallOption) (*BidEditV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidEditV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidsServiceFetchClient) SubmitDecision(ctx context.Context, in *BidsRequestSubmitDecisionV1, opts ...grpc.CallOption) (*BidEditV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidEditV1)
	err := c.cc.Invoke(ctx, BidsServiceFetch_SubmitDecision_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidsServiceFetchServer is the server API for BidsServiceFetch service.
// All implementations must embed UnimplementedBidsServiceFetchServer
// for forward compatibility.
type BidsServiceFetchServer interface {
	Edit(context.Context, *BidsRequestEditV1) (*BidEditV1, error)
	Rollback(context.Context, *BidsRequestRollbackV1) (*BidEditV1, error)
	Status(context.Context, *BidsRequestStatusV1) (*BidEditV1, error)
	SubmitDecision(context.Context, *BidsRequestSubmitDecisionV1) (*BidEditV1, error)
	mustEmbedUnimplementedBidsServiceFetchServer()
}

// UnimplementedBidsServiceFetchServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBidsServiceFetchServer struct{}

func (UnimplementedBidsServiceFetchServer) Edit(context.Context, *BidsRequestEditV1) (*BidEditV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedBidsServiceFetchServer) Rollback(context.Context, *BidsRequestRollbackV1) (*BidEditV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedBidsServiceFetchServer) Status(context.Context, *BidsRequestStatusV1) (*BidEditV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedBidsServiceFetchServer) SubmitDecision(context.Context, *BidsRequestSubmitDecisionV1) (*BidEditV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDecision not implemented")
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

func _BidsServiceFetch_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestEditV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_Edit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).Edit(ctx, req.(*BidsRequestEditV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestRollbackV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_Rollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).Rollback(ctx, req.(*BidsRequestRollbackV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestStatusV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).Status(ctx, req.(*BidsRequestStatusV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidsServiceFetch_SubmitDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidsRequestSubmitDecisionV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidsServiceFetchServer).SubmitDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BidsServiceFetch_SubmitDecision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidsServiceFetchServer).SubmitDecision(ctx, req.(*BidsRequestSubmitDecisionV1))
	}
	return interceptor(ctx, in, info, handler)
}

// BidsServiceFetch_ServiceDesc is the grpc.ServiceDesc for BidsServiceFetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidsServiceFetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "update.BidsServiceFetch",
	HandlerType: (*BidsServiceFetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Edit",
			Handler:    _BidsServiceFetch_Edit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _BidsServiceFetch_Rollback_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _BidsServiceFetch_Status_Handler,
		},
		{
			MethodName: "SubmitDecision",
			Handler:    _BidsServiceFetch_SubmitDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protos/proto/bids/proto_bids_service_update.proto",
}
