package service

import (
	"context"
	"douyin-user/idl/douyin_favorite/kitex_gen/douyinfavorite"
	"douyin-user/pkg/errno"
	"douyin-user/server/favorite/dal/db"
)

type LikeService struct {
	ctx context.Context
}

// NewLikeService new CheckUserService
func NewLikeService(ctx context.Context) *LikeService {
	return &LikeService{
		ctx: ctx,
	}
}

// Like favorite a video
func (s *LikeService) Like(req *douyinfavorite.LikeRequest) error {

	videoId := req.VideoId
	userId := req.GetUserId()
	favorite := db.Favorite{
		VideoId: videoId,
		UserId:  userId,
	}
	getFavorite, err := db.GetFavorite(s.ctx, &favorite)

	if err != nil {
		return err
	}

	if getFavorite.UserId != 0 {
		return errno.FavoriteErr
	}

	err = db.CreateFavorite(s.ctx, &favorite)

	return err

}

// NotLike favorite a video
func (s *LikeService) NotLike(req *douyinfavorite.LikeRequest) error {

	videoId := req.VideoId
	userId := req.GetUserId()
	favorite := db.Favorite{
		VideoId: videoId,
		UserId:  userId,
	}
	getFavorite, err := db.GetFavorite(s.ctx, &favorite)

	if err != nil {
		return err
	}

	if getFavorite.UserId == 0 {
		return errno.FavoriteErr
	}

	err = db.DelFavorite(s.ctx, &favorite)

	return err

}
