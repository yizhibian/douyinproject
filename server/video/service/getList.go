package service

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/server/video/dal/db"
	"douyin-user/server/video/pack"
	"log"
)

type GetListService struct {
	ctx context.Context
}

func NewGetListService(ctx context.Context) *GetListService {
	return &GetListService{
		ctx: ctx,
	}
}

func (s *GetListService) GetList(userId int64) []*douyinvideo.Video {
	var vs []*db.Video
	db.DB.Where("author_id = ?", userId).Find(&vs)
	log.Printf("list:%#v\n", vs)
	return pack.VideoInfos(vs)
}
