package data

import (
	"account/internal/biz"
	"account/internal/conf"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
)

var EmailCodeMap map[string]string = make(map[string]string, 10000)
var mu sync.Mutex

func AddToMap(key, value string, duration time.Duration) {
	mu.Lock()
	EmailCodeMap[key] = value
	mu.Unlock()

	time.AfterFunc(duration, func() {
		mu.Lock()
		delete(EmailCodeMap, key)
		mu.Unlock()
	})
}

func GetEmailCodeFromMap(key string) (string, bool) {
	mu.Lock()
	value, ok := EmailCodeMap[key]
	mu.Unlock()
	return value, ok
}

func GetRandSixCode() string {
	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < 6; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}
	return code
}

type authRepo struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewAuthRepo(conf *conf.Bootstrap, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (a *authRepo) GenerateToken(ctx context.Context, conf *conf.Bootstrap, username, email string, userID int64) (token string, err error) {
	claims := biz.Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(conf.GetAuth().GetExpireDuration().AsDuration()).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(a.conf.GetAuth().GetJwtSecret()))
	return token, err
}

func (a *authRepo) ParseToken(ctx context.Context, conf *conf.Bootstrap, token string) (claims *biz.Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &biz.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.conf.GetAuth().GetJwtSecret()), nil
	})
	// fmt.Println("tokenClaims:", tokenClaims)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*biz.Claims)
		// fmt.Println("claims:", claims)
		if ok && tokenClaims.Valid {
			// fmt.Println("return success!")
			return claims, nil
		}
	}
	// fmt.Println("err:", err)
	return nil, err
}

func (a *authRepo) SendEmailCode(ctx context.Context, email string) (success bool, err error) {
	//参数校验
	if email == "" {
		return false, fmt.Errorf("参数为空")
	}

	emailCode := GetRandSixCode()

	smtpHost := a.conf.GetAuth().GetHost()
	smtpPort := a.conf.GetAuth().GetPort()
	sender := a.conf.GetAuth().GetUsername()
	senderPassword := a.conf.GetAuth().GetPassword()

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "电商验证码")
	m.SetBody("text/plain", fmt.Sprintf("你的电商注册验证码为：%s", emailCode))

	d := gomail.NewDialer(smtpHost, int(smtpPort), sender, senderPassword)

	err = d.DialAndSend(m)
	if err != nil {
		return false, fmt.Errorf("发送验证码失败:%w", err)
	}

	AddToMap(email, emailCode, 3*60*time.Second)

	return true, nil
}
func (a *authRepo) TestEmailCode(ctx context.Context, email, email_code string) (success bool, err error) {
	if realCode, ok := GetEmailCodeFromMap(email); ok && realCode == email_code {
		return true, nil
	} else {
		return false, fmt.Errorf("验证码错误")
	}
}
