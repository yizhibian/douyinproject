package main

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/video/dal/db"
	"douyin-user/server/video/service"
	"log"
	"strconv"
	"time"
)

// VideoServerImpl implements the last service interface defined in the IDL.
type VideoServerImpl struct{}

// Publish implements the VideoServerImpl interface.
func (s *VideoServerImpl) Publish(ctx context.Context, req *douyinvideo.PublishRequest) (resp *douyinvideo.PublishResponse, err error) {
	resp = douyinvideo.NewPublishResponse()
	preq := db.Video{
		AuthorId:      req.AuthorId,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    true,
		Title:         req.Title,
	}
	service.NewPublishService(ctx).Publish(&preq)
	return resp, nil
}

// GetList implements the VideoServerImpl interface.
func (s *VideoServerImpl) GetList(ctx context.Context, req *douyinvideo.GetListRequest) (resp *douyinvideo.GetListResponse, err error) {
	resp = douyinvideo.NewGetListResponse()
	userId := req.GetUserId()
	if userId == 0 {
		resp.StatusCode = errno.ParamErr.ErrCode
		resp.StatusMsg = errno.ParamErr.ErrMsg
		return resp, nil
	}
	list := service.NewGetListService(ctx).GetList(userId)
	resp.SetVideoList(list)
	return resp, nil
}

// Feed implements the VideoServerImpl interface.
func (s *VideoServerImpl) Feed(ctx context.Context, req *douyinvideo.FeedRequest) (resp *douyinvideo.FeedResponse, err error) {
	resp = douyinvideo.NewFeedResponse()
	latestTime, _ := strconv.ParseInt(req.GetLatestTime(), 10, 64)
	timeStr := time.Unix(latestTime, 0).Format("2006-01-02 15:04:05")
	log.Printf("from:%#v\n", timeStr)
	list := service.NewFeedService(ctx).Feed(timeStr)
	resp.SetVideoList(list)
	log.Printf("setlist:%#v\n", list)
	resp.SetNextTime(time.Now().Unix())
	return resp, nil
}

// Feed implements the VideoServerImpl interface.
func (s *VideoServerImpl) GetListByVID(ctx context.Context, req *douyinvideo.GetListByVIDRequest) (resp *douyinvideo.GetListByVIDRespose, err error) {
	resp = douyinvideo.NewGetListByVIDRespose()
	list := service.NewGetListByVIDService(ctx).GetListByVID(req.VideoId)
	resp.SetVideoList(list)
	log.Printf("setlist:%#v\n", list)
	return resp, nil
}
