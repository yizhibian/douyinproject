package main

import (
	"context"
	douyinfavorite "douyin-user/idl/douyin_favorite/kitex_gen/douyinfavorite"
)

// UserServerImpl implements the last service interface defined in the IDL.
type UserServerImpl struct{}

// Like implements the UserServerImpl interface.
func (s *UserServerImpl) Like(ctx context.Context, req *douyinfavorite.LikeRequest) (resp *douyinfavorite.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoIds implements the UserServerImpl interface.
func (s *UserServerImpl) GetVideoIds(ctx context.Context, req *douyinfavorite.GetVideoIdsRequest) (resp *douyinfavorite.GetVideoIdsResponse, err error) {
	// TODO: Your code here...
	return
}
