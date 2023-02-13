package service

import (
	"context"
	"douyin-user/idl/kitex_test/kitex_gen/douyinuser"
	"douyin-user/server/user/dal/db"
	"douyin-user/server/user/pack"
)

type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService return userInfo
func NewGetUserInfoService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) GetUserInfo(userId int64) (*douyinuser.User, error) {

	user, err := db.GetUser(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	info := pack.UserInfo(user)
	return info, nil
}
