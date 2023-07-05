package data

import (
	"account/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("user"),
	}
}

func (u *userRepo) FetchByUsername(ctx context.Context, username string) (user *biz.User, err error) {
	user = &biz.User{}
	u.table.WithContext(ctx).First(user, "username = ?", username)
	if user.ID == 0 {
		return nil, biz.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepo) FetchByEmail(ctx context.Context, email string) (user *biz.User, err error) {
	user = &biz.User{}
	u.table.WithContext(ctx).First(user, "email = ?", email)
	if user.ID == 0 {
		return nil, biz.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (id int64, err error) {
	result := u.table.WithContext(ctx).Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *userRepo) Delete(ctx context.Context, user *biz.User) (err error) {
	result := u.table.WithContext(ctx).Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
