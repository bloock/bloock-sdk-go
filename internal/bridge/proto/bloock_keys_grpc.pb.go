// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: bloock_keys.proto

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
	LoadLocalKey(ctx context.Context, in *LoadLocalKeyRequest, opts ...grpc.CallOption) (*LoadLocalKeyResponse, error)
	GenerateManagedKey(ctx context.Context, in *GenerateManagedKeyRequest, opts ...grpc.CallOption) (*GenerateManagedKeyResponse, error)
	LoadManagedKey(ctx context.Context, in *LoadManagedKeyRequest, opts ...grpc.CallOption) (*LoadManagedKeyResponse, error)
	GenerateLocalCertificate(ctx context.Context, in *GenerateLocalCertificateRequest, opts ...grpc.CallOption) (*GenerateLocalCertificateResponse, error)
	LoadLocalCertificate(ctx context.Context, in *LoadLocalCertificateRequest, opts ...grpc.CallOption) (*LoadLocalCertificateResponse, error)
	GenerateManagedCertificate(ctx context.Context, in *GenerateManagedCertificateRequest, opts ...grpc.CallOption) (*GenerateManagedCertificateResponse, error)
	LoadManagedCertificate(ctx context.Context, in *LoadManagedCertificateRequest, opts ...grpc.CallOption) (*LoadManagedCertificateResponse, error)
	ImportManagedCertificate(ctx context.Context, in *ImportManagedCertificateRequest, opts ...grpc.CallOption) (*ImportManagedCertificateResponse, error)
	SetupTotpAccessControl(ctx context.Context, in *SetupTotpAccessControlRequest, opts ...grpc.CallOption) (*SetupTotpAccessControlResponse, error)
	SetupSecretAccessControl(ctx context.Context, in *SetupSecretAccessControlRequest, opts ...grpc.CallOption) (*SetupSecretAccessControlResponse, error)
	RecoverTotpAccessControl(ctx context.Context, in *RecoverTotpAccessControlRequest, opts ...grpc.CallOption) (*RecoverTotpAccessControlResponse, error)
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

func (c *keyServiceClient) LoadLocalKey(ctx context.Context, in *LoadLocalKeyRequest, opts ...grpc.CallOption) (*LoadLocalKeyResponse, error) {
	out := new(LoadLocalKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadLocalKey", in, out, opts...)
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

func (c *keyServiceClient) LoadManagedKey(ctx context.Context, in *LoadManagedKeyRequest, opts ...grpc.CallOption) (*LoadManagedKeyResponse, error) {
	out := new(LoadManagedKeyResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadManagedKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GenerateLocalCertificate(ctx context.Context, in *GenerateLocalCertificateRequest, opts ...grpc.CallOption) (*GenerateLocalCertificateResponse, error) {
	out := new(GenerateLocalCertificateResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/GenerateLocalCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) LoadLocalCertificate(ctx context.Context, in *LoadLocalCertificateRequest, opts ...grpc.CallOption) (*LoadLocalCertificateResponse, error) {
	out := new(LoadLocalCertificateResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadLocalCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GenerateManagedCertificate(ctx context.Context, in *GenerateManagedCertificateRequest, opts ...grpc.CallOption) (*GenerateManagedCertificateResponse, error) {
	out := new(GenerateManagedCertificateResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/GenerateManagedCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) LoadManagedCertificate(ctx context.Context, in *LoadManagedCertificateRequest, opts ...grpc.CallOption) (*LoadManagedCertificateResponse, error) {
	out := new(LoadManagedCertificateResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/LoadManagedCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ImportManagedCertificate(ctx context.Context, in *ImportManagedCertificateRequest, opts ...grpc.CallOption) (*ImportManagedCertificateResponse, error) {
	out := new(ImportManagedCertificateResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/ImportManagedCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) SetupTotpAccessControl(ctx context.Context, in *SetupTotpAccessControlRequest, opts ...grpc.CallOption) (*SetupTotpAccessControlResponse, error) {
	out := new(SetupTotpAccessControlResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/SetupTotpAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) SetupSecretAccessControl(ctx context.Context, in *SetupSecretAccessControlRequest, opts ...grpc.CallOption) (*SetupSecretAccessControlResponse, error) {
	out := new(SetupSecretAccessControlResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/SetupSecretAccessControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) RecoverTotpAccessControl(ctx context.Context, in *RecoverTotpAccessControlRequest, opts ...grpc.CallOption) (*RecoverTotpAccessControlResponse, error) {
	out := new(RecoverTotpAccessControlResponse)
	err := c.cc.Invoke(ctx, "/bloock.KeyService/RecoverTotpAccessControl", in, out, opts...)
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
	LoadLocalKey(context.Context, *LoadLocalKeyRequest) (*LoadLocalKeyResponse, error)
	GenerateManagedKey(context.Context, *GenerateManagedKeyRequest) (*GenerateManagedKeyResponse, error)
	LoadManagedKey(context.Context, *LoadManagedKeyRequest) (*LoadManagedKeyResponse, error)
	GenerateLocalCertificate(context.Context, *GenerateLocalCertificateRequest) (*GenerateLocalCertificateResponse, error)
	LoadLocalCertificate(context.Context, *LoadLocalCertificateRequest) (*LoadLocalCertificateResponse, error)
	GenerateManagedCertificate(context.Context, *GenerateManagedCertificateRequest) (*GenerateManagedCertificateResponse, error)
	LoadManagedCertificate(context.Context, *LoadManagedCertificateRequest) (*LoadManagedCertificateResponse, error)
	ImportManagedCertificate(context.Context, *ImportManagedCertificateRequest) (*ImportManagedCertificateResponse, error)
	SetupTotpAccessControl(context.Context, *SetupTotpAccessControlRequest) (*SetupTotpAccessControlResponse, error)
	SetupSecretAccessControl(context.Context, *SetupSecretAccessControlRequest) (*SetupSecretAccessControlResponse, error)
	RecoverTotpAccessControl(context.Context, *RecoverTotpAccessControlRequest) (*RecoverTotpAccessControlResponse, error)
	mustEmbedUnimplementedKeyServiceServer()
}

// UnimplementedKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyServiceServer struct {
}

func (UnimplementedKeyServiceServer) GenerateLocalKey(context.Context, *GenerateLocalKeyRequest) (*GenerateLocalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLocalKey not implemented")
}
func (UnimplementedKeyServiceServer) LoadLocalKey(context.Context, *LoadLocalKeyRequest) (*LoadLocalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadLocalKey not implemented")
}
func (UnimplementedKeyServiceServer) GenerateManagedKey(context.Context, *GenerateManagedKeyRequest) (*GenerateManagedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateManagedKey not implemented")
}
func (UnimplementedKeyServiceServer) LoadManagedKey(context.Context, *LoadManagedKeyRequest) (*LoadManagedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadManagedKey not implemented")
}
func (UnimplementedKeyServiceServer) GenerateLocalCertificate(context.Context, *GenerateLocalCertificateRequest) (*GenerateLocalCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLocalCertificate not implemented")
}
func (UnimplementedKeyServiceServer) LoadLocalCertificate(context.Context, *LoadLocalCertificateRequest) (*LoadLocalCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadLocalCertificate not implemented")
}
func (UnimplementedKeyServiceServer) GenerateManagedCertificate(context.Context, *GenerateManagedCertificateRequest) (*GenerateManagedCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateManagedCertificate not implemented")
}
func (UnimplementedKeyServiceServer) LoadManagedCertificate(context.Context, *LoadManagedCertificateRequest) (*LoadManagedCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadManagedCertificate not implemented")
}
func (UnimplementedKeyServiceServer) ImportManagedCertificate(context.Context, *ImportManagedCertificateRequest) (*ImportManagedCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportManagedCertificate not implemented")
}
func (UnimplementedKeyServiceServer) SetupTotpAccessControl(context.Context, *SetupTotpAccessControlRequest) (*SetupTotpAccessControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupTotpAccessControl not implemented")
}
func (UnimplementedKeyServiceServer) SetupSecretAccessControl(context.Context, *SetupSecretAccessControlRequest) (*SetupSecretAccessControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupSecretAccessControl not implemented")
}
func (UnimplementedKeyServiceServer) RecoverTotpAccessControl(context.Context, *RecoverTotpAccessControlRequest) (*RecoverTotpAccessControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTotpAccessControl not implemented")
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

func _KeyService_GenerateLocalCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLocalCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateLocalCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/GenerateLocalCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateLocalCertificate(ctx, req.(*GenerateLocalCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_LoadLocalCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadLocalCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).LoadLocalCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/LoadLocalCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).LoadLocalCertificate(ctx, req.(*LoadLocalCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GenerateManagedCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateManagedCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateManagedCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/GenerateManagedCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateManagedCertificate(ctx, req.(*GenerateManagedCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_LoadManagedCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadManagedCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).LoadManagedCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/LoadManagedCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).LoadManagedCertificate(ctx, req.(*LoadManagedCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ImportManagedCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportManagedCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ImportManagedCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/ImportManagedCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ImportManagedCertificate(ctx, req.(*ImportManagedCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_SetupTotpAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupTotpAccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).SetupTotpAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/SetupTotpAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).SetupTotpAccessControl(ctx, req.(*SetupTotpAccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_SetupSecretAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupSecretAccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).SetupSecretAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/SetupSecretAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).SetupSecretAccessControl(ctx, req.(*SetupSecretAccessControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_RecoverTotpAccessControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverTotpAccessControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).RecoverTotpAccessControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bloock.KeyService/RecoverTotpAccessControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).RecoverTotpAccessControl(ctx, req.(*RecoverTotpAccessControlRequest))
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
			MethodName: "LoadLocalKey",
			Handler:    _KeyService_LoadLocalKey_Handler,
		},
		{
			MethodName: "GenerateManagedKey",
			Handler:    _KeyService_GenerateManagedKey_Handler,
		},
		{
			MethodName: "LoadManagedKey",
			Handler:    _KeyService_LoadManagedKey_Handler,
		},
		{
			MethodName: "GenerateLocalCertificate",
			Handler:    _KeyService_GenerateLocalCertificate_Handler,
		},
		{
			MethodName: "LoadLocalCertificate",
			Handler:    _KeyService_LoadLocalCertificate_Handler,
		},
		{
			MethodName: "GenerateManagedCertificate",
			Handler:    _KeyService_GenerateManagedCertificate_Handler,
		},
		{
			MethodName: "LoadManagedCertificate",
			Handler:    _KeyService_LoadManagedCertificate_Handler,
		},
		{
			MethodName: "ImportManagedCertificate",
			Handler:    _KeyService_ImportManagedCertificate_Handler,
		},
		{
			MethodName: "SetupTotpAccessControl",
			Handler:    _KeyService_SetupTotpAccessControl_Handler,
		},
		{
			MethodName: "SetupSecretAccessControl",
			Handler:    _KeyService_SetupSecretAccessControl_Handler,
		},
		{
			MethodName: "RecoverTotpAccessControl",
			Handler:    _KeyService_RecoverTotpAccessControl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bloock_keys.proto",
}
