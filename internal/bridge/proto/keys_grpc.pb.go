// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: keys.proto

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

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyServiceClient interface {
	GenerateLocalKey(ctx context.Context, in *GenerateLocalKeyRequest, opts ...grpc.CallOption) (*GenerateLocalKeyResponse, error)
	GenerateManagedKey(ctx context.Context, in *GenerateManagedKeyRequest, opts ...grpc.CallOption) (*GenerateManagedKeyResponse, error)
	LoadLocalKey(ctx context.Context, in *LoadLocalKeyRequest, opts ...grpc.CallOption) (*LoadLocalKeyResponse, error)
	LoadManagedKey(ctx context.Context, in *LoadManagedKeyRequest, opts ...grpc.CallOption) (*LoadManagedKeyResponse, error)
}

type keyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyServiceClient(cc grpc.ClientConnInterface) KeyServiceClient {
	return &keyServiceClient{cc}
}

func (c *keyServiceClient) GenerateLocalKey(ctx context.Context, in *GenerateLocalKeyRequest, opts ...grpc.CallOption) (*GenerateLocalKeyResponse, error) {
	out := new(GenerateLocalKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/GenerateLocalKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GenerateManagedKey(ctx context.Context, in *GenerateManagedKeyRequest, opts ...grpc.CallOption) (*GenerateManagedKeyResponse, error) {
	out := new(GenerateManagedKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/GenerateManagedKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) LoadLocalKey(ctx context.Context, in *LoadLocalKeyRequest, opts ...grpc.CallOption) (*LoadLocalKeyResponse, error) {
	out := new(LoadLocalKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadLocalKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) LoadManagedKey(ctx context.Context, in *LoadManagedKeyRequest, opts ...grpc.CallOption) (*LoadManagedKeyResponse, error) {
	out := new(LoadManagedKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadManagedKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyServiceServer is the server API for KeyService service.
// All implementations must embed UnimplementedKeyServiceServer
// for forward compatibility
type KeyServiceServer interface {
	GenerateLocalKey(context.Context, *GenerateLocalKeyRequest) (*GenerateLocalKeyResponse, error)
	GenerateManagedKey(context.Context, *GenerateManagedKeyRequest) (*GenerateManagedKeyResponse, error)
	LoadLocalKey(context.Context, *LoadLocalKeyRequest) (*LoadLocalKeyResponse, error)
	LoadManagedKey(context.Context, *LoadManagedKeyRequest) (*LoadManagedKeyResponse, error)
	mustEmbedUnimplementedKeyServiceServer()
}

// UnimplementedKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyServiceServer struct {
}

func (UnimplementedKeyServiceServer) GenerateLocalKey(context.Context, *GenerateLocalKeyRequest) (*GenerateLocalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLocalKey not implemented")
}
func (UnimplementedKeyServiceServer) GenerateManagedKey(context.Context, *GenerateManagedKeyRequest) (*GenerateManagedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateManagedKey not implemented")
}
func (UnimplementedKeyServiceServer) LoadLocalKey(context.Context, *LoadLocalKeyRequest) (*LoadLocalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadLocalKey not implemented")
}
func (UnimplementedKeyServiceServer) LoadManagedKey(context.Context, *LoadManagedKeyRequest) (*LoadManagedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadManagedKey not implemented")
}
func (UnimplementedKeyServiceServer) mustEmbedUnimplementedKeyServiceServer() {}

// UnsafeKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyServiceServer will
// result in compilation errors.
type UnsafeKeyServiceServer interface {
	mustEmbedUnimplementedKeyServiceServer()
}

func RegisterKeyServiceServer(s grpc.ServiceRegistrar, srv KeyServiceServer) {
	s.RegisterService(&KeyService_ServiceDesc, srv)
}

func _KeyService_GenerateLocalKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLocalKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateLocalKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/GenerateLocalKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateLocalKey(ctx, req.(*GenerateLocalKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GenerateManagedKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateManagedKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateManagedKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/GenerateManagedKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateManagedKey(ctx, req.(*GenerateManagedKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_LoadLocalKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadLocalKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).LoadLocalKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/LoadLocalKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).LoadLocalKey(ctx, req.(*LoadLocalKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_LoadManagedKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadManagedKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).LoadManagedKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/LoadManagedKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).LoadManagedKey(ctx, req.(*LoadManagedKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyService_ServiceDesc is the grpc.ServiceDesc for KeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bloock.KeyService",
	HandlerType: (*KeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateLocalKey",
			Handler:    _KeyService_GenerateLocalKey_Handler,
		},
		{
			MethodName: "GenerateManagedKey",
			Handler:    _KeyService_GenerateManagedKey_Handler,
		},
		{
			MethodName: "LoadLocalKey",
			Handler:    _KeyService_LoadLocalKey_Handler,
		},
		{
			MethodName: "LoadManagedKey",
			Handler:    _KeyService_LoadManagedKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keys.proto",
}
