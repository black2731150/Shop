// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v4.23.3
// source: account/v1/account.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAccountDeleteAccounter = "/account.v1.Account/DeleteAccounter"
const OperationAccountLoginWithCode = "/account.v1.Account/LoginWithCode"
const OperationAccountLoginWithPassword = "/account.v1.Account/LoginWithPassword"
const OperationAccountRegister = "/account.v1.Account/Register"
const OperationAccountSendEmailCode = "/account.v1.Account/SendEmailCode"
const OperationAccountTestEmail = "/account.v1.Account/TestEmail"
const OperationAccountTestEmailCode = "/account.v1.Account/TestEmailCode"
const OperationAccountTestUsername = "/account.v1.Account/TestUsername"

type AccountHTTPServer interface {
	DeleteAccounter(context.Context, *DeleteAccounterRequest) (*DeleteAccounterResponse, error)
	LoginWithCode(context.Context, *LoginWithCodeRequest) (*LoginWithCodeRsponse, error)
	LoginWithPassword(context.Context, *LoginWithPasswordRequest) (*LoginWithPasswordResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	SendEmailCode(context.Context, *SendEmailCodeRequest) (*SendEmailCodeResponse, error)
	TestEmail(context.Context, *TestEmailRequest) (*TestEmailResponse, error)
	TestEmailCode(context.Context, *TestEmailCodeRequest) (*TestEmailCodeResponse, error)
	TestUsername(context.Context, *TestUsernameRequest) (*TestUsernameResponse, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("account/register", _Account_Register0_HTTP_Handler(srv))
	r.POST("account/test_email", _Account_TestEmail0_HTTP_Handler(srv))
	r.POST("account/send_email_code", _Account_SendEmailCode0_HTTP_Handler(srv))
	r.POST("account/test_email_code", _Account_TestEmailCode0_HTTP_Handler(srv))
	r.POST("account/test_username", _Account_TestUsername0_HTTP_Handler(srv))
	r.POST("account/login_with_password", _Account_LoginWithPassword0_HTTP_Handler(srv))
	r.POST("account/login_with_code", _Account_LoginWithCode0_HTTP_Handler(srv))
	r.POST("account/delete_accounter", _Account_DeleteAccounter0_HTTP_Handler(srv))
}

func _Account_Register0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_TestEmail0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountTestEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TestEmail(ctx, req.(*TestEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestEmailResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_SendEmailCode0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendEmailCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountSendEmailCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendEmailCode(ctx, req.(*SendEmailCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendEmailCodeResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_TestEmailCode0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestEmailCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountTestEmailCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TestEmailCode(ctx, req.(*TestEmailCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestEmailCodeResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_TestUsername0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestUsernameRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountTestUsername)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TestUsername(ctx, req.(*TestUsernameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestUsernameResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_LoginWithPassword0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountLoginWithPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithPassword(ctx, req.(*LoginWithPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginWithPasswordResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_LoginWithCode0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountLoginWithCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithCode(ctx, req.(*LoginWithCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginWithCodeRsponse)
		return ctx.Result(200, reply)
	}
}

func _Account_DeleteAccounter0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAccounterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountDeleteAccounter)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAccounter(ctx, req.(*DeleteAccounterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAccounterResponse)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	DeleteAccounter(ctx context.Context, req *DeleteAccounterRequest, opts ...http.CallOption) (rsp *DeleteAccounterResponse, err error)
	LoginWithCode(ctx context.Context, req *LoginWithCodeRequest, opts ...http.CallOption) (rsp *LoginWithCodeRsponse, err error)
	LoginWithPassword(ctx context.Context, req *LoginWithPasswordRequest, opts ...http.CallOption) (rsp *LoginWithPasswordResponse, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResponse, err error)
	SendEmailCode(ctx context.Context, req *SendEmailCodeRequest, opts ...http.CallOption) (rsp *SendEmailCodeResponse, err error)
	TestEmail(ctx context.Context, req *TestEmailRequest, opts ...http.CallOption) (rsp *TestEmailResponse, err error)
	TestEmailCode(ctx context.Context, req *TestEmailCodeRequest, opts ...http.CallOption) (rsp *TestEmailCodeResponse, err error)
	TestUsername(ctx context.Context, req *TestUsernameRequest, opts ...http.CallOption) (rsp *TestUsernameResponse, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) DeleteAccounter(ctx context.Context, in *DeleteAccounterRequest, opts ...http.CallOption) (*DeleteAccounterResponse, error) {
	var out DeleteAccounterResponse
	pattern := "account/delete_accounter"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountDeleteAccounter))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) LoginWithCode(ctx context.Context, in *LoginWithCodeRequest, opts ...http.CallOption) (*LoginWithCodeRsponse, error) {
	var out LoginWithCodeRsponse
	pattern := "account/login_with_code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountLoginWithCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) LoginWithPassword(ctx context.Context, in *LoginWithPasswordRequest, opts ...http.CallOption) (*LoginWithPasswordResponse, error) {
	var out LoginWithPasswordResponse
	pattern := "account/login_with_password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountLoginWithPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResponse, error) {
	var out RegisterResponse
	pattern := "account/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...http.CallOption) (*SendEmailCodeResponse, error) {
	var out SendEmailCodeResponse
	pattern := "account/send_email_code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountSendEmailCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) TestEmail(ctx context.Context, in *TestEmailRequest, opts ...http.CallOption) (*TestEmailResponse, error) {
	var out TestEmailResponse
	pattern := "account/test_email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountTestEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) TestEmailCode(ctx context.Context, in *TestEmailCodeRequest, opts ...http.CallOption) (*TestEmailCodeResponse, error) {
	var out TestEmailCodeResponse
	pattern := "account/test_email_code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountTestEmailCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) TestUsername(ctx context.Context, in *TestUsernameRequest, opts ...http.CallOption) (*TestUsernameResponse, error) {
	var out TestUsernameResponse
	pattern := "account/test_username"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountTestUsername))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
