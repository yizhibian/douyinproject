package service

import (
	"context"
	"crypto/md5"
	"douyin-user/idl/kitex_test/kitex_gen/douyinuser"
	"douyin-user/pkg/errno"
	"douyin-user/server/user/dal/db"
	"fmt"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *douyinuser.UserRequest) (int64, error) {

	users, err := db.QueryUser(s.ctx, req.GetUsername())
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.GetPassword()); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	user := db.User{
		UserName: req.GetUsername(),
		Password: passWord,
	}
	err = db.CreateUser(s.ctx, &user)
	if err != nil {
		return 0, err
	}
	return user.Id, err
}
