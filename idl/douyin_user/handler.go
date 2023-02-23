package main

import (
	"context"
	douyinuser "douyin-user/idl/douyin_user/kitex_gen/douyinuser"
)

// UserServerImpl implements the last service interface defined in the IDL.
type UserServerImpl struct{}

// CreateUser implements the UserServerImpl interface.
func (s *UserServerImpl) CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) (resp *douyinuser.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServerImpl interface.
func (s *UserServerImpl) CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (resp *douyinuser.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServerImpl interface.
func (s *UserServerImpl) GetUserInfo(ctx context.Context, req *douyinuser.GetUserInfoRequest) (resp *douyinuser.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
