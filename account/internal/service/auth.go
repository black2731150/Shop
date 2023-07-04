package service

import (
	v1 "account/api/account/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type authService struct {
	v1.UnimplementedAuthenticationServer
	log *log.Helper
}
