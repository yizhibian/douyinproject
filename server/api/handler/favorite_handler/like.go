package favorite_handler

import (
	"context"
	"douyin-user/idl/douyin_favorite/kitex_gen/douyinfavorite"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

func LikeOrNot(ctx context.Context, c *app.RequestContext) {
	var queryVar LikeParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	id, exists := c.Get("identity")
	if !exists {
		log.Println("cant get id")
	}

	req := douyinfavorite.LikeRequest{
		VideoId:    queryVar.VideoId,
		UserId:     int64(id.(float64)),
		ActionType: queryVar.ActionType,
	}

	response, err := rpc.LikeOrNot(ctx, &req)

	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	pack.SendBaseResponse(c, nil, response)
}
