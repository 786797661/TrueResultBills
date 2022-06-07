// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: billsResult/v1/billsResult.proto

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

// BillsResultClient is the client API for BillsResult service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillsResultClient interface {
	// Sends a greeting
	GetTrueResult(ctx context.Context, in *GetTrueResultRequest, opts ...grpc.CallOption) (*GetTrueResultReply, error)
}

type billsResultClient struct {
	cc grpc.ClientConnInterface
}

func NewBillsResultClient(cc grpc.ClientConnInterface) BillsResultClient {
	return &billsResultClient{cc}
}

func (c *billsResultClient) GetTrueResult(ctx context.Context, in *GetTrueResultRequest, opts ...grpc.CallOption) (*GetTrueResultReply, error) {
	out := new(GetTrueResultReply)
	err := c.cc.Invoke(ctx, "/billsResult.v1.BillsResult/GetTrueResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillsResultServer is the server API for BillsResult service.
// All implementations must embed UnimplementedBillsResultServer
// for forward compatibility
type BillsResultServer interface {
	// Sends a greeting
	GetTrueResult(context.Context, *GetTrueResultRequest) (*GetTrueResultReply, error)
	mustEmbedUnimplementedBillsResultServer()
}

// UnimplementedBillsResultServer must be embedded to have forward compatible implementations.
type UnimplementedBillsResultServer struct {
}

func (UnimplementedBillsResultServer) GetTrueResult(context.Context, *GetTrueResultRequest) (*GetTrueResultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrueResult not implemented")
}
func (UnimplementedBillsResultServer) mustEmbedUnimplementedBillsResultServer() {}

// UnsafeBillsResultServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillsResultServer will
// result in compilation errors.
type UnsafeBillsResultServer interface {
	mustEmbedUnimplementedBillsResultServer()
}

func RegisterBillsResultServer(s grpc.ServiceRegistrar, srv BillsResultServer) {
	s.RegisterService(&BillsResult_ServiceDesc, srv)
}

func _BillsResult_GetTrueResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrueResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillsResultServer).GetTrueResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billsResult.v1.BillsResult/GetTrueResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillsResultServer).GetTrueResult(ctx, req.(*GetTrueResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillsResult_ServiceDesc is the grpc.ServiceDesc for BillsResult service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillsResult_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billsResult.v1.BillsResult",
	HandlerType: (*BillsResultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrueResult",
			Handler:    _BillsResult_GetTrueResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billsResult/v1/billsResult.proto",
}