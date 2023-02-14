package main

import (
	"context"
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/pkg/errno"
	"douyin-user/server/user/pack"
	"douyin-user/server/user/service"
)

// UserServerImpl implements the last service interface defined in the IDL.
type UserServerImpl struct{}

// CreateUser implements the UserServerImpl interface.
func (s *UserServerImpl) CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) (resp *douyinuser.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = douyinuser.NewCreateUserResponse()
	userReq := req.GetUserReq()

	if len(userReq.GetUsername()) == 0 || len(userReq.GetPassword()) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCreateUserService(ctx).CreateUser(userReq)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserResp = &douyinuser.UserResponse{UserId: uid}
	return resp, nil
}

// CheckUser implements the UserServerImpl interface.
func (s *UserServerImpl) CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (resp *douyinuser.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = douyinuser.NewCheckUserResponse()
	userReq := req.GetUserReq()

	if len(userReq.GetUsername()) == 0 || len(userReq.GetPassword()) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(userReq)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserResp = &douyinuser.UserResponse{UserId: uid}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserInfo implements the UserServerImpl interface.
func (s *UserServerImpl) GetUserInfo(ctx context.Context, req *douyinuser.GetUserInfoRequest) (resp *douyinuser.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	resp = douyinuser.NewGetUserInfoResponse()
	userId := req.GetUserId()
	if userId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userInfo, err := service.NewGetUserInfoService(ctx).GetUserInfo(userId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.SetUser(userInfo)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
