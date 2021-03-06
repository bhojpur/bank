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

// BankUIClient is the client API for BankUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankUIClient interface {
	// ListBranchSpecs returns a list of Branch that can be started through the UI.
	ListBranchSpecs(ctx context.Context, in *ListBranchSpecsRequest, opts ...grpc.CallOption) (BankUI_ListBranchSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type bankUIClient struct {
	cc grpc.ClientConnInterface
}

func NewBankUIClient(cc grpc.ClientConnInterface) BankUIClient {
	return &bankUIClient{cc}
}

func (c *bankUIClient) ListBranchSpecs(ctx context.Context, in *ListBranchSpecsRequest, opts ...grpc.CallOption) (BankUI_ListBranchSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankUI_ServiceDesc.Streams[0], "/v1.BankUI/ListBranchSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &bankUIListBranchSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BankUI_ListBranchSpecsClient interface {
	Recv() (*ListBranchSpecsResponse, error)
	grpc.ClientStream
}

type bankUIListBranchSpecsClient struct {
	grpc.ClientStream
}

func (x *bankUIListBranchSpecsClient) Recv() (*ListBranchSpecsResponse, error) {
	m := new(ListBranchSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.BankUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankUIServer is the server API for BankUI service.
// All implementations must embed UnimplementedBankUIServer
// for forward compatibility
type BankUIServer interface {
	// ListBranchSpecs returns a list of Branch that can be started through the UI.
	ListBranchSpecs(*ListBranchSpecsRequest, BankUI_ListBranchSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedBankUIServer()
}

// UnimplementedBankUIServer must be embedded to have forward compatible implementations.
type UnimplementedBankUIServer struct {
}

func (UnimplementedBankUIServer) ListBranchSpecs(*ListBranchSpecsRequest, BankUI_ListBranchSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBranchSpecs not implemented")
}
func (UnimplementedBankUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedBankUIServer) mustEmbedUnimplementedBankUIServer() {}

// UnsafeBankUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankUIServer will
// result in compilation errors.
type UnsafeBankUIServer interface {
	mustEmbedUnimplementedBankUIServer()
}

func RegisterBankUIServer(s grpc.ServiceRegistrar, srv BankUIServer) {
	s.RegisterService(&BankUI_ServiceDesc, srv)
}

func _BankUI_ListBranchSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBranchSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankUIServer).ListBranchSpecs(m, &bankUIListBranchSpecsServer{stream})
}

type BankUI_ListBranchSpecsServer interface {
	Send(*ListBranchSpecsResponse) error
	grpc.ServerStream
}

type bankUIListBranchSpecsServer struct {
	grpc.ServerStream
}

func (x *bankUIListBranchSpecsServer) Send(m *ListBranchSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BankUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.BankUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankUI_ServiceDesc is the grpc.ServiceDesc for BankUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.BankUI",
	HandlerType: (*BankUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _BankUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBranchSpecs",
			Handler:       _BankUI_ListBranchSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bank-ui.proto",
}
