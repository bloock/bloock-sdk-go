// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proof.proto

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

// ProofServiceClient is the client API for ProofService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProofServiceClient interface {
	GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error)
	ValidateRoot(ctx context.Context, in *ValidateRootRequest, opts ...grpc.CallOption) (*ValidateRootResponse, error)
	VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofResponse, error)
	VerifyRecords(ctx context.Context, in *VerifyRecordsRequest, opts ...grpc.CallOption) (*VerifyRecordsResponse, error)
}

type proofServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProofServiceClient(cc grpc.ClientConnInterface) ProofServiceClient {
	return &proofServiceClient{cc}
}

func (c *proofServiceClient) GetProof(ctx context.Context, in *GetProofRequest, opts ...grpc.CallOption) (*GetProofResponse, error) {
	out := new(GetProofResponse)
	err := c.cc.Invoke(ctx, "/bloock.ProofService/GetProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proofServiceClient) ValidateRoot(ctx context.Context, in *ValidateRootRequest, opts ...grpc.CallOption) (*ValidateRootResponse, error) {
	out := new(ValidateRootResponse)
	err := c.cc.Invoke(ctx, "/bloock.ProofService/ValidateRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proofServiceClient) VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofResponse, error) {
	out := new(VerifyProofResponse)
	err := c.cc.Invoke(ctx, "/bloock.ProofService/VerifyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proofServiceClient) VerifyRecords(ctx context.Context, in *VerifyRecordsRequest, opts ...grpc.CallOption) (*VerifyRecordsResponse, error) {
	out := new(VerifyRecordsResponse)
	err := c.cc.Invoke(ctx, "/bloock.ProofService/VerifyRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProofServiceServer is the server API for ProofService service.
// All implementations must embed UnimplementedProofServiceServer
// for forward compatibility
type ProofServiceServer interface {
	GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error)
	ValidateRoot(context.Context, *ValidateRootRequest) (*ValidateRootResponse, error)
	VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofResponse, error)
	VerifyRecords(context.Context, *VerifyRecordsRequest) (*VerifyRecordsResponse, error)
	mustEmbedUnimplementedProofServiceServer()
}

// UnimplementedProofServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProofServiceServer struct {
}

func (UnimplementedProofServiceServer) GetProof(context.Context, *GetProofRequest) (*GetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedProofServiceServer) ValidateRoot(context.Context, *ValidateRootRequest) (*ValidateRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateRoot not implemented")
}
func (UnimplementedProofServiceServer) VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProof not implemented")
}
func (UnimplementedProofServiceServer) VerifyRecords(context.Context, *VerifyRecordsRequest) (*VerifyRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRecords not implemented")
}
func (UnimplementedProofServiceServer) mustEmbedUnimplementedProofServiceServer() {}

// UnsafeProofServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProofServiceServer will
// result in compilation errors.
type UnsafeProofServiceServer interface {
	mustEmbedUnimplementedProofServiceServer()
}

func RegisterProofServiceServer(s grpc.ServiceRegistrar, srv ProofServiceServer) {
	s.RegisterService(&ProofService_ServiceDesc, srv)
}

func _ProofService_GetProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofServiceServer).GetProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.ProofService/GetProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofServiceServer).GetProof(ctx, req.(*GetProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProofService_ValidateRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofServiceServer).ValidateRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.ProofService/ValidateRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofServiceServer).ValidateRoot(ctx, req.(*ValidateRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProofService_VerifyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofServiceServer).VerifyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.ProofService/VerifyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofServiceServer).VerifyProof(ctx, req.(*VerifyProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProofService_VerifyRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofServiceServer).VerifyRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.ProofService/VerifyRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofServiceServer).VerifyRecords(ctx, req.(*VerifyRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProofService_ServiceDesc is the grpc.ServiceDesc for ProofService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProofService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bloock.ProofService",
	HandlerType: (*ProofServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProof",
			Handler:    _ProofService_GetProof_Handler,
		},
		{
			MethodName: "ValidateRoot",
			Handler:    _ProofService_ValidateRoot_Handler,
		},
		{
			MethodName: "VerifyProof",
			Handler:    _ProofService_VerifyProof_Handler,
		},
		{
			MethodName: "VerifyRecords",
			Handler:    _ProofService_VerifyRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proof.proto",
}
