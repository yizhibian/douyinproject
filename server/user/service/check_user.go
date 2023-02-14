package service

import (
	"context"
	"crypto/md5"
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/pkg/errno"
	"douyin-user/server/user/dal/db"
	"fmt"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *douyinuser.UserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.GetPassword()); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.GetUsername()
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return u.Id, nil
}
