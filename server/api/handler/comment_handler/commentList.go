package handlers

import (
	"context"
	"douyin-user/idl/douyin_comment/kitex_gen/comment"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"strconv"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	//UserId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//if UserId <= 0 {
	//	SendResponse(c, errno.UserIdErr, nil)
	//	return
	//}
	VideoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if VideoId <= 0 {
		SendResponse(c, errno.VideoErr, nil)
		return
	}
	ret, err := rpc.CommentList(ctx, &comment.CommentListRequest{
		UserId:  0,
		VideoId: VideoId,
	})
	if err != nil {
		c.JSON(consts.StatusOK, comment.CommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "失败",
			CommentList: nil,
		})
	} else {
		c.JSON(consts.StatusOK, comment.CommentListResponse{
			StatusCode:  0,
			StatusMsg:   "成功",
			CommentList: ret.CommentList,
		})
	}
}
