package main

import (
	"context"
	"douyin-user/idl/douyin_favorite/kitex_gen/douyinfavorite"
	"douyin-user/pkg/errno"
	"douyin-user/server/favorite/pack"
	"douyin-user/server/favorite/service"
)

// UserServerImpl implements the last service interface defined in the IDL.
type UserServerImpl struct{}

// Like implements the UserServerImpl interface.
func (s *UserServerImpl) Like(ctx context.Context, req *douyinfavorite.LikeRequest) (resp *douyinfavorite.LikeResponse, err error) {
	// TODO: Your code here...

	resp = douyinfavorite.NewLikeResponse()
	v := req.GetActionType()

	switch v {

	case 1:
		err := service.NewLikeService(ctx).Like(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		break
	case 2:
		err := service.NewLikeService(ctx).NotLike(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		break
	default:
		pack.BuildBaseResp(errno.WrongActionTypeErr)
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// GetVideoIds implements the UserServerImpl interface.
func (s *UserServerImpl) GetVideoIds(ctx context.Context, req *douyinfavorite.GetVideoIdsRequest) (resp *douyinfavorite.GetVideoIdsResponse, err error) {
	// TODO: Your code here...
	return
}
