// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: account/v1/account.proto

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

const (
	Account_Register_FullMethodName          = "/account.v1.Account/Register"
	Account_TestEmail_FullMethodName         = "/account.v1.Account/TestEmail"
	Account_SendEmailCode_FullMethodName     = "/account.v1.Account/SendEmailCode"
	Account_TestEmailCode_FullMethodName     = "/account.v1.Account/TestEmailCode"
	Account_TestUsername_FullMethodName      = "/account.v1.Account/TestUsername"
	Account_LoginWithPassword_FullMethodName = "/account.v1.Account/LoginWithPassword"
	Account_LoginWithCode_FullMethodName     = "/account.v1.Account/LoginWithCode"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	TestEmail(ctx context.Context, in *TestEmailRequest, opts ...grpc.CallOption) (*TestEmailResponse, error)
	SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SendEmailCodeResponse, error)
	TestEmailCode(ctx context.Context, in *TestEmailCodeRequest, opts ...grpc.CallOption) (*TestEmailCodeResponse, error)
	TestUsername(ctx context.Context, in *TestUsernameRequest, opts ...grpc.CallOption) (*TestUsernameResponse, error)
	LoginWithPassword(ctx context.Context, in *LoginWithPasswordRequest, opts ...grpc.CallOption) (*LoginWithPasswordResponse, error)
	LoginWithCode(ctx context.Context, in *LoginWithCodeRequest, opts ...grpc.CallOption) (*LoginWithCodeRsponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Account_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) TestEmail(ctx context.Context, in *TestEmailRequest, opts ...grpc.CallOption) (*TestEmailResponse, error) {
	out := new(TestEmailResponse)
	err := c.cc.Invoke(ctx, Account_TestEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SendEmailCodeResponse, error) {
	out := new(SendEmailCodeResponse)
	err := c.cc.Invoke(ctx, Account_SendEmailCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) TestEmailCode(ctx context.Context, in *TestEmailCodeRequest, opts ...grpc.CallOption) (*TestEmailCodeResponse, error) {
	out := new(TestEmailCodeResponse)
	err := c.cc.Invoke(ctx, Account_TestEmailCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) TestUsername(ctx context.Context, in *TestUsernameRequest, opts ...grpc.CallOption) (*TestUsernameResponse, error) {
	out := new(TestUsernameResponse)
	err := c.cc.Invoke(ctx, Account_TestUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginWithPassword(ctx context.Context, in *LoginWithPasswordRequest, opts ...grpc.CallOption) (*LoginWithPasswordResponse, error) {
	out := new(LoginWithPasswordResponse)
	err := c.cc.Invoke(ctx, Account_LoginWithPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) LoginWithCode(ctx context.Context, in *LoginWithCodeRequest, opts ...grpc.CallOption) (*LoginWithCodeRsponse, error) {
	out := new(LoginWithCodeRsponse)
	err := c.cc.Invoke(ctx, Account_LoginWithCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	TestEmail(context.Context, *TestEmailRequest) (*TestEmailResponse, error)
	SendEmailCode(context.Context, *SendEmailCodeRequest) (*SendEmailCodeResponse, error)
	TestEmailCode(context.Context, *TestEmailCodeRequest) (*TestEmailCodeResponse, error)
	TestUsername(context.Context, *TestUsernameRequest) (*TestUsernameResponse, error)
	LoginWithPassword(context.Context, *LoginWithPasswordRequest) (*LoginWithPasswordResponse, error)
	LoginWithCode(context.Context, *LoginWithCodeRequest) (*LoginWithCodeRsponse, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAccountServer) TestEmail(context.Context, *TestEmailRequest) (*TestEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestEmail not implemented")
}
func (UnimplementedAccountServer) SendEmailCode(context.Context, *SendEmailCodeRequest) (*SendEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedAccountServer) TestEmailCode(context.Context, *TestEmailCodeRequest) (*TestEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestEmailCode not implemented")
}
func (UnimplementedAccountServer) TestUsername(context.Context, *TestUsernameRequest) (*TestUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestUsername not implemented")
}
func (UnimplementedAccountServer) LoginWithPassword(context.Context, *LoginWithPasswordRequest) (*LoginWithPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithPassword not implemented")
}
func (UnimplementedAccountServer) LoginWithCode(context.Context, *LoginWithCodeRequest) (*LoginWithCodeRsponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithCode not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_TestEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).TestEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_TestEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).TestEmail(ctx, req.(*TestEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_SendEmailCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SendEmailCode(ctx, req.(*SendEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_TestEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).TestEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_TestEmailCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).TestEmailCode(ctx, req.(*TestEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_TestUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).TestUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_TestUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).TestUsername(ctx, req.(*TestUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginWithPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginWithPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_LoginWithPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginWithPassword(ctx, req.(*LoginWithPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_LoginWithCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginWithCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_LoginWithCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginWithCode(ctx, req.(*LoginWithCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Account_Register_Handler,
		},
		{
			MethodName: "TestEmail",
			Handler:    _Account_TestEmail_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _Account_SendEmailCode_Handler,
		},
		{
			MethodName: "TestEmailCode",
			Handler:    _Account_TestEmailCode_Handler,
		},
		{
			MethodName: "TestUsername",
			Handler:    _Account_TestUsername_Handler,
		},
		{
			MethodName: "LoginWithPassword",
			Handler:    _Account_LoginWithPassword_Handler,
		},
		{
			MethodName: "LoginWithCode",
			Handler:    _Account_LoginWithCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1/account.proto",
}

const (
	Authentication_TestToken_FullMethodName = "/account.v1.Authentication/TestToken"
)

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	TestToken(ctx context.Context, in *TestTokenRequest, opts ...grpc.CallOption) (*TestTokenResponse, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) TestToken(ctx context.Context, in *TestTokenRequest, opts ...grpc.CallOption) (*TestTokenResponse, error) {
	out := new(TestTokenResponse)
	err := c.cc.Invoke(ctx, Authentication_TestToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	TestToken(context.Context, *TestTokenRequest) (*TestTokenResponse, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) TestToken(context.Context, *TestTokenRequest) (*TestTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_TestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).TestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_TestToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).TestToken(ctx, req.(*TestTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestToken",
			Handler:    _Authentication_TestToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1/account.proto",
}
