// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package token

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error)
	DescribeToken(ctx context.Context, in *DescribeTokenRequest, opts ...grpc.CallOption) (*Token, error)
	ChangeNamespace(ctx context.Context, in *ChangeNamespaceRequest, opts ...grpc.CallOption) (*Token, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.token.TokenService/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.token.TokenService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DescribeToken(ctx context.Context, in *DescribeTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.token.TokenService/DescribeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ChangeNamespace(ctx context.Context, in *ChangeNamespaceRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.token.TokenService/ChangeNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
	ChangeNamespace(context.Context, *ChangeNamespaceRequest) (*Token, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) IssueToken(context.Context, *IssueTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedTokenServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedTokenServiceServer) DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeToken not implemented")
}
func (UnimplementedTokenServiceServer) ChangeNamespace(context.Context, *ChangeNamespaceRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNamespace not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.token.TokenService/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.token.TokenService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DescribeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DescribeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.token.TokenService/DescribeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DescribeToken(ctx, req.(*DescribeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ChangeNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ChangeNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.token.TokenService/ChangeNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ChangeNamespace(ctx, req.(*ChangeNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.token.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _TokenService_IssueToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _TokenService_ValidateToken_Handler,
		},
		{
			MethodName: "DescribeToken",
			Handler:    _TokenService_DescribeToken_Handler,
		},
		{
			MethodName: "ChangeNamespace",
			Handler:    _TokenService_ChangeNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/token/pb/service.proto",
}
