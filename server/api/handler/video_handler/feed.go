package video_handler

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
)

// Feed implements the VideoServerImpl interface.
func Feed(ctx context.Context, c *app.RequestContext) {
	var queryVar FeedParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	log.Info(queryVar)
	req := douyinvideo.FeedRequest{
		LatestTime: queryVar.LatestTime,
		Token:      queryVar.Token,
	}
	r, err := rpc.Feed(context.Background(), &req)
	log.Info(r)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
