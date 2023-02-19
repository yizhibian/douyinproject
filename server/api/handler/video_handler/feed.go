package video_handler

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed implements the VideoServerImpl interface.
func Feed(ctx context.Context, c *app.RequestContext) {
	var queryVar FeedParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	req := douyinvideo.FeedRequest{
		LatestTime: queryVar.LatestTime,
		Token:      queryVar.Token,
	}
	r, err := rpc.Feed(context.Background(), &req)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
