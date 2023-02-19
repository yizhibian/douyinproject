package main

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/video/pack"
	"douyin-user/server/video/service"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"time"
)

// VideoServerImpl implements the last service interface defined in the IDL.
type VideoServerImpl struct{}

// Publish implements the VideoServerImpl interface.
func (s *VideoServerImpl) Publish(ctx context.Context, req *douyinvideo.PublishRequest) (resp *douyinvideo.PublishResponse, err error) {
	resp = douyinvideo.NewPublishResponse()
	title := req.GetTitle()
	id := ctx.Value("PublishUserId")
	data := req.Data
	if len(data) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	service.NewPublishService(ctx).Publish(title, data, id.(int64))
	return resp, nil
}

// GetList implements the VideoServerImpl interface.
func (s *VideoServerImpl) GetList(ctx context.Context, req *douyinvideo.GetListRequest) (resp *douyinvideo.GetListResponse, err error) {
	resp = douyinvideo.NewGetListResponse()
	userId := req.GetUserId()
	if userId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	list := service.NewGetListService(ctx).GetList(userId)
	resp.SetVideoList(list)
	return resp, nil
}

// Feed implements the VideoServerImpl interface.
func (s *VideoServerImpl) Feed(ctx context.Context, req *douyinvideo.FeedRequest) (resp *douyinvideo.FeedResponse, err error) {
	resp = douyinvideo.NewFeedResponse()
	latestTime := req.GetLatestTime()
	log.Info(latestTime)
	list := service.NewFeedService(ctx).Feed(latestTime)
	resp.SetVideoList(list)
	resp.SetNextTime(time.Now().Unix())
	return
}
