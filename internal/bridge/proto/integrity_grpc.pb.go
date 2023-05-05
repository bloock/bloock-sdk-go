// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: integrity.proto

package proto

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

// IntegrityServiceClient is the client API for IntegrityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrityServiceClient interface {
	SendRecords(ctx context.Context, in *SendRecordsRequest, opts ...grpc.CallOption) (*SendRecordsResponse, error)
	GetAnchor(ctx context.Context, in *GetAnchorRequest, opts ...grpc.CallOption) (*GetAnchorResponse, error)
	WaitAnchor(ctx context.Context, in *WaitAnchorRequest, opts ...grpc.CallOption) (*WaitAnchorResponse, error)
	GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error)
	ValidateRoot(ctx context.Context, in *ValidateRootRequest, opts ...grpc.CallOption) (*ValidateRootResponse, error)
	VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofResponse, error)
	VerifyRecords(ctx context.Context, in *VerifyRecordsRequest, opts ...grpc.CallOption) (*VerifyRecordsResponse, error)
}

type integrityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrityServiceClient(cc grpc.ClientConnInterface) IntegrityServiceClient {
	return &integrityServiceClient{cc}
}

func (c *integrityServiceClient) SendRecords(ctx context.Context, in *SendRecordsRequest, opts ...grpc.CallOption) (*SendRecordsResponse, error) {
	out := new(SendRecordsResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/SendRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) GetAnchor(ctx context.Context, in *GetAnchorRequest, opts ...grpc.CallOption) (*GetAnchorResponse, error) {
	out := new(GetAnchorResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/GetAnchor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) WaitAnchor(ctx context.Context, in *WaitAnchorRequest, opts ...grpc.CallOption) (*WaitAnchorResponse, error) {
	out := new(WaitAnchorResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/WaitAnchor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error) {
	out := new(GetProofResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/GetProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) ValidateRoot(ctx context.Context, in *ValidateRootRequest, opts ...grpc.CallOption) (*ValidateRootResponse, error) {
	out := new(ValidateRootResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/ValidateRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofResponse, error) {
	out := new(VerifyProofResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/VerifyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrityServiceClient) VerifyRecords(ctx context.Context, in *VerifyRecordsRequest, opts ...grpc.CallOption) (*VerifyRecordsResponse, error) {
	out := new(VerifyRecordsResponse)
	err := c.cc.Invoke(ctx, "/bloock.IntegrityService/VerifyRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrityServiceServer is the server API for IntegrityService service.
// All implementations must embed UnimplementedIntegrityServiceServer
// for forward compatibility
type IntegrityServiceServer interface {
	SendRecords(context.Context, *SendRecordsRequest) (*SendRecordsResponse, error)
	GetAnchor(context.Context, *GetAnchorRequest) (*GetAnchorResponse, error)
	WaitAnchor(context.Context, *WaitAnchorRequest) (*WaitAnchorResponse, error)
	GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error)
	ValidateRoot(context.Context, *ValidateRootRequest) (*ValidateRootResponse, error)
	VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofResponse, error)
	VerifyRecords(context.Context, *VerifyRecordsRequest) (*VerifyRecordsResponse, error)
	mustEmbedUnimplementedIntegrityServiceServer()
}

// UnimplementedIntegrityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrityServiceServer struct {
}

func (UnimplementedIntegrityServiceServer) SendRecords(context.Context, *SendRecordsRequest) (*SendRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRecords not implemented")
}
func (UnimplementedIntegrityServiceServer) GetAnchor(context.Context, *GetAnchorRequest) (*GetAnchorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnchor not implemented")
}
func (UnimplementedIntegrityServiceServer) WaitAnchor(context.Context, *WaitAnchorRequest) (*WaitAnchorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitAnchor not implemented")
}
func (UnimplementedIntegrityServiceServer) GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedIntegrityServiceServer) ValidateRoot(context.Context, *ValidateRootRequest) (*ValidateRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateRoot not implemented")
}
func (UnimplementedIntegrityServiceServer) VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProof not implemented")
}
func (UnimplementedIntegrityServiceServer) VerifyRecords(context.Context, *VerifyRecordsRequest) (*VerifyRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRecords not implemented")
}
func (UnimplementedIntegrityServiceServer) mustEmbedUnimplementedIntegrityServiceServer() {}

// UnsafeIntegrityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrityServiceServer will
// result in compilation errors.
type UnsafeIntegrityServiceServer interface {
	mustEmbedUnimplementedIntegrityServiceServer()
}

func RegisterIntegrityServiceServer(s grpc.ServiceRegistrar, srv IntegrityServiceServer) {
	s.RegisterService(&IntegrityService_ServiceDesc, srv)
}

func _IntegrityService_SendRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).SendRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/SendRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).SendRecords(ctx, req.(*SendRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_GetAnchor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnchorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).GetAnchor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/GetAnchor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).GetAnchor(ctx, req.(*GetAnchorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_WaitAnchor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitAnchorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).WaitAnchor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/WaitAnchor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).WaitAnchor(ctx, req.(*WaitAnchorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_GetProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).GetProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/GetProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).GetProof(ctx, req.(*GetProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_ValidateRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).ValidateRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/ValidateRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).ValidateRoot(ctx, req.(*ValidateRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_VerifyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).VerifyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/VerifyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).VerifyProof(ctx, req.(*VerifyProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrityService_VerifyRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrityServiceServer).VerifyRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.IntegrityService/VerifyRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrityServiceServer).VerifyRecords(ctx, req.(*VerifyRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrityService_ServiceDesc is the grpc.ServiceDesc for IntegrityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bloock.IntegrityService",
	HandlerType: (*IntegrityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRecords",
			Handler:    _IntegrityService_SendRecords_Handler,
		},
		{
			MethodName: "GetAnchor",
			Handler:    _IntegrityService_GetAnchor_Handler,
		},
		{
			MethodName: "WaitAnchor",
			Handler:    _IntegrityService_WaitAnchor_Handler,
		},
		{
			MethodName: "GetProof",
			Handler:    _IntegrityService_GetProof_Handler,
		},
		{
			MethodName: "ValidateRoot",
			Handler:    _IntegrityService_ValidateRoot_Handler,
		},
		{
			MethodName: "VerifyProof",
			Handler:    _IntegrityService_VerifyProof_Handler,
		},
		{
			MethodName: "VerifyRecords",
			Handler:    _IntegrityService_VerifyRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrity.proto",
}
