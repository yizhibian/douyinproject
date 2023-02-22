package video_handler

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

// Publish implements the VideoServerImpl interface.
func Publish(ctx context.Context, c *app.RequestContext) {
	var queryVar PublishParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//file, err2 := c.FormFile("data")
	//if err2 != nil {
	//	pack.SendBaseResponse(c, errno.ConvertErr(err2), nil)
	//	return
	//}
	req := douyinvideo.PublishRequest{
		Data:  nil,
		Token: queryVar.Token,
		Title: queryVar.Title,
	}

	//var id interface{}
	//if id, exists := c.Get("identity"); exists {
	//	fmt.Println("the id is ", id)
	//}
	id, exists := c.Get("identity")
	if !exists {
		log.Println("cant get id")
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//id := int64(claims[constants.IdentityKey].(float64))
	//log.Printf("id======%#v\n", id)
	//d := 2
	sonCtx := context.WithValue(ctx, "PublishUserId", id)
	r, err := rpc.Publish(sonCtx, &req)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
