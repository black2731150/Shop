package biz

import (
	"account/internal/conf"
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

type UserRepo interface {
	FetchByUsername(ctx context.Context, username string) (user *User, err error)
	FetchByEmail(ctx context.Context, email string) (user *User, err error)
	Save(ctx context.Context, user *User) (id int64, err error)
}

type AccountUseCase struct {
	config   *conf.Bootstrap
	userRepo UserRepo
	logger   log.Helper
	authRepo AuthRepo
	// encryptService EncryptService
}

type Claims struct {
	Username string
	Email    string
	UserID   int64
	jwt.StandardClaims
}

type AuthRepo interface {
	GenerateToken(ctx context.Context, conf *conf.Bootstrap, username, email string, userID int64) (token string, err error)
	ParseToken(ctx context.Context, conf *conf.Bootstrap, token string) (claims *Claims, err error)
	SendEmailCode(ctx context.Context, email string) (success bool, err error)
	TestEmailCode(ctx context.Context, email, email_code string) (success bool, err error)
}

var (
	ErrUserNotFound *kerr.Error = kerr.New(001, "USER_NOT_FUND", "Can not find the user")
)

func NewAccountUseCase(logger log.Logger, Config *conf.Bootstrap, userRepo UserRepo, authRepo AuthRepo) *AccountUseCase {
	return &AccountUseCase{
		userRepo: userRepo,
		logger:   *log.NewHelper(logger),
		config:   Config,
		authRepo: authRepo,
	}
}

func (a *AccountUseCase) Register(ctx context.Context, username, password, email string) (err error) {
	//参数校验
	if username == "" || password == "" || email == "" {
		return fmt.Errorf("注册失败")
	}

	//判断用户是否注册过一次
	user, err := a.userRepo.FetchByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		log.Errorf("注册失败,参数[username: %s, password: %s, email: %s],err:%v", username, password, email, err)
		return fmt.Errorf("注册失败")
	}

	if user != nil {
		return fmt.Errorf("用户已经存在")
	}

	_, err = a.userRepo.Save(ctx, &User{
		Username: username,
		Password: password,
		Email:    email,
	})

	if err != nil {
		return fmt.Errorf("注册失败：%w", err)
	}

	return nil
}

func (a *AccountUseCase) TestEmail(ctx context.Context, email string) (err error) {
	//参数校验
	if email == "" {
		return fmt.Errorf("参数为空")
	}
	//判断用户是否注册过一次
	user, err := a.userRepo.FetchByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		log.Errorf("注册失败,参数[email: %s],err:%v", email, err)
		return fmt.Errorf("注册失败")
	}

	if user != nil {
		return fmt.Errorf("用户已经存在")
	}

	return nil
}

func (a *AccountUseCase) SendEmailCode(ctx context.Context, email string) (err error) {
	ok, err := a.authRepo.SendEmailCode(ctx, email)

	if ok {
		return nil
	} else {
		return err
	}
}

func (a *AccountUseCase) TestEmailCode(ctx context.Context, email, emailcode string) (err error) {
	ok, err := a.authRepo.TestEmailCode(ctx, email, emailcode)
	if ok {
		return nil
	} else {
		return err
	}
}

func (a *AccountUseCase) TestUsername(ctx context.Context, username string) (err error) {
	if username == "" {
		return fmt.Errorf("用户名为空")
	}

	for _, ch := range username {
		if ch == '@' {
			return fmt.Errorf("用户名中有不符合规范的字符%c", ch)
		}
	}

	user, err := a.userRepo.FetchByUsername(ctx, username)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		log.Errorf("检测失败,参数[email: %s],err:%v", username, err)
		return fmt.Errorf("检测失败")
	}

	if user != nil {
		return fmt.Errorf("用户已经存在")
	}

	return nil
}

func (a *AccountUseCase) LoginWithPassword(ctx context.Context, username_or_email, password string) (token string, err error) {
	token = ""

	if username_or_email == "" || password == "" {
		return token, fmt.Errorf("登录失败,参数不能为空")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	var user *User
	if emailRegex.MatchString(username_or_email) {
		user, err = a.userRepo.FetchByEmail(ctx, username_or_email)
		if err != nil && !errors.Is(err, ErrUserNotFound) {
			log.Errorf("登录失败,参数[email: %s, password: %s],err:%v", username_or_email, password, err)
			return token, fmt.Errorf("登录失败")
		}
	} else {
		user, err = a.userRepo.FetchByUsername(ctx, username_or_email)
		if err != nil && !errors.Is(err, ErrUserNotFound) {
			log.Errorf("登录失败,参数[username: %s, password: %s],err:%v", username_or_email, password, err)
			return token, fmt.Errorf("登录失败")
		}
	}

	if user == nil {
		return token, fmt.Errorf("登录失败，用户不存在")
	}

	if user.Password != password {
		return token, fmt.Errorf("登录失败,用户名或者密码错误")
	}

	token, err = a.authRepo.GenerateToken(ctx, a.config, user.Username, user.Email, user.ID)
	if err != nil {
		return token, fmt.Errorf("token生成失败")
	}
	return token, nil
}

func (a *AccountUseCase) LoginWithCode(ctx context.Context, email, emial_code string) (token string, err error) {
	token = ""

	if email == "" || emial_code == "" {
		return token, fmt.Errorf("登录失败,参数不能为空")
	}

	user, err := a.userRepo.FetchByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		log.Errorf("登录失败,参数[email: %s],err:%v", email, err)
		return token, fmt.Errorf("登录失败")
	}

	if user == nil {
		return token, fmt.Errorf("登录失败，用户不存在")
	}

	ok, err := a.authRepo.TestEmailCode(ctx, email, emial_code)
	if !ok {
		return token, err
	}

	token, err = a.authRepo.GenerateToken(ctx, a.config, user.Username, user.Email, user.ID)
	if err != nil {
		return token, fmt.Errorf("token生成失败")
	}

	return token, nil

}
