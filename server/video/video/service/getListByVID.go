package service

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/server/video/dal/db"
	"douyin-user/server/video/pack"
	"log"
)

type GetListByVIDService struct {
	ctx context.Context
}

func NewGetListByVIDService(ctx context.Context) *GetListByVIDService {
	return &GetListByVIDService{
		ctx: ctx,
	}
}

func (s *GetListByVIDService) GetListByVID(videoIds []int64) []*douyinvideo.Video {
	var vs []*db.Video
	db.DB.Where(videoIds).Find(&vs)
	log.Printf("list:%#v\n", vs)
	return pack.VideoInfos(vs)
}
