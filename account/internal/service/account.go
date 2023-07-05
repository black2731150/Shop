package service

import (
	v1 "account/api/account/v1"
	"account/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type accountService struct {
	v1.UnimplementedAccountServer
	log *log.Helper
	auc *biz.AccountUseCase
}

func NewAccountUseCase(logger log.Logger, auc *biz.AccountUseCase) v1.AccountServer {
	return &accountService{
		log: log.NewHelper(logger),
		auc: auc,
	}
}

func (a *accountService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	err := a.auc.Register(ctx, request.GetUsername(), request.GetPassword(), request.GetEmail())
	if err != nil {
		return nil, errors.New(500, "注册失败", err.Error())
	}
	return &v1.RegisterResponse{
		Success: true,
	}, nil
}

func (a *accountService) TestEmail(ctx context.Context, request *v1.TestEmailRequest) (*v1.TestEmailResponse, error) {
	err := a.auc.TestEmail(ctx, request.GetEmail())
	if err != nil {
		return nil, errors.New(500, "邮箱验证失败", err.Error())
	}
	return &v1.TestEmailResponse{
		Success: true,
	}, nil
}

func (a *accountService) SendEmailCode(ctx context.Context, request *v1.SendEmailCodeRequest) (*v1.SendEmailCodeResponse, error) {
	err := a.auc.SendEmailCode(ctx, request.GetEmail())
	if err != nil {
		return nil, errors.New(500, "发送验证码失败", err.Error())
	}
	return &v1.SendEmailCodeResponse{
		Success: true,
	}, nil
}

func (a *accountService) TestEmailCode(ctx context.Context, request *v1.TestEmailCodeRequest) (*v1.TestEmailCodeResponse, error) {
	err := a.auc.TestEmailCode(ctx, request.GetEmail(), request.GetEmailCode())
	if err != nil {
		return nil, errors.New(500, "验证码不正确", err.Error())
	}
	return &v1.TestEmailCodeResponse{
		Success: true,
	}, nil
}

func (a *accountService) TestUsername(ctx context.Context, request *v1.TestUsernameRequest) (*v1.TestUsernameResponse, error) {
	err := a.auc.TestUsername(ctx, request.GetUsername())
	if err != nil {
		return nil, errors.New(500, "用户名不符合规范", err.Error())
	}
	return &v1.TestUsernameResponse{
		Success: true,
	}, nil
}

func (a *accountService) LoginWithPassword(ctx context.Context, request *v1.LoginWithPasswordRequest) (*v1.LoginWithPasswordResponse, error) {
	token, err := a.auc.LoginWithPassword(ctx, request.GetUsernameOrEmail(), request.Password)
	if err != nil {
		return nil, errors.New(500, "登录失败", err.Error())
	}
	return &v1.LoginWithPasswordResponse{
		Success: true,
		Token:   token,
	}, nil
}

func (a *accountService) LoginWithCode(ctx context.Context, request *v1.LoginWithCodeRequest) (*v1.LoginWithCodeRsponse, error) {
	token, err := a.auc.LoginWithCode(ctx, request.Email, request.EmailCode)
	if err != nil {
		return nil, errors.New(500, "登录失败", err.Error())
	}
	return &v1.LoginWithCodeRsponse{
		Success: true,
		Token:   token,
	}, nil
}

func (a *accountService) DeleteAccounter(ctx context.Context, request *v1.DeleteAccounterRequest) (*v1.DeleteAccounterResponse, error) {
	err := a.auc.DeleteAccounter(ctx, request.GetEmail())
	if err != nil {
		return nil, errors.New(500, "登录失败", err.Error())
	}
	return &v1.DeleteAccounterResponse{
		Success: true,
	}, nil
}
