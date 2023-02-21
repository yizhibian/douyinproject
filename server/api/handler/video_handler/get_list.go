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

// GetList implements the VideoServerImpl interface.
func GetList(ctx context.Context, c *app.RequestContext) {
	var queryVar GetListParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	req := douyinvideo.GetListRequest{
		Token:  queryVar.Token,
		UserId: queryVar.UserId,
	}
	r, err := rpc.GetList(context.Background(), &req)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
