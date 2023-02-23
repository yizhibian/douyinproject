package service

import (
	"context"
	"douyin-user/server/video/dal/db"
)

type AddLikeService struct {
	ctx context.Context
}

func NewAddLikeService(ctx context.Context) *AddLikeService {
	return &AddLikeService{
		ctx: ctx,
	}
}

func (s *AddLikeService) AddLike(video_id int64) {
	var vs *db.Video
	db.DB.Where("video_id = ?", video_id).Find(&vs)
	db.DB.Model(&db.Video{}).Where("video_id = ?", video_id).Update("favorite_count", vs.FavoriteCount+1)
}
