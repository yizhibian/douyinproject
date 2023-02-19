package service

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/server/video/dal/db"
	"douyin-user/server/video/pack"
	"log"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(latestTime string) []*douyinvideo.Video {
	var vs []*db.Video
	db.DB.Order("created_at desc").Limit(5).Find(&vs)
	log.Printf("list:%#v\n", vs)
	return pack.VideoInfos(vs)
}
