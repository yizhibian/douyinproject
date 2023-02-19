package video_handler

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Publish implements the VideoServerImpl interface.
func Publish(ctx context.Context, c *app.RequestContext) {
	var queryVar PublishParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	req := douyinvideo.PublishRequest{
		Data:  queryVar.Data,
		Token: queryVar.Token,
		Title: queryVar.Title,
	}
	var id interface{}
	if id, exists := c.Get("identity"); exists {
		fmt.Println("the id is ", id)
	}
	sonCtx := context.WithValue(ctx, "PublishUserId", id.(int64))
	r, err := rpc.Publish(sonCtx, &req)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
