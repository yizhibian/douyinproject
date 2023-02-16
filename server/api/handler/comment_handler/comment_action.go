package handlers

import (
	"context"
	"douyin-user/idl/douyin_comment/kitex_gen/comment"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"strconv"
)

func CommentAction(ctx context.Context, c *app.RequestContext) {
	var commentVar CommentRequestParam
	if err := c.Bind(&commentVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	claims := jwt.ExtractClaims(ctx, c)
	UserId := int64(claims[constants.IdentityKey].(float64))
	//UserId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if UserId <= 0 {
		SendResponse(c, errno.UserIdErr, nil)
		return
	}
	VideoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if VideoId <= 0 {
		SendResponse(c, errno.VideoErr, nil)
		return
	}

	ActionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if ActionType != 1 && ActionType != 2 {
		SendResponse(c, errno.ActionTypeErr, nil)
		return
	}
	if ActionType == 1 {
		//发布评论
		CommentTextErr := c.Query("comment_text")
		if len(CommentTextErr) == 0 {
			SendResponse(c, errno.CommentTextErr, nil)
			return
		}
		if len(CommentTextErr) > 100 {
			SendResponse(c, errno.CommentTextTooLongErr, nil)
			return
		}
		ret, err := rpc.CommentAction(ctx, &comment.CommentRequest{
			UserId:      UserId,
			VideoId:     VideoId,
			ActionType:  ActionType,
			CommentText: c.Query("comment_text"),
			CommentId:   0,
		})
		if err != nil {
			c.JSON(consts.StatusOK, comment.CommentResponse{
				StatusCode: -1,
				StatusMsg:  "失败",
				Comment:    nil,
			})
		} else {
			c.JSON(consts.StatusOK, comment.CommentResponse{
				StatusCode: 0,
				StatusMsg:  "成功",
				Comment:    ret.Comment,
			})
		}
		//err := rpc.(context.Background(), &notedemo.CreateNoteRequest{
		//	UserId:  userID,
		//	Content: noteVar.Content, Title: noteVar.Title,
		//})
		//todo:对video_id进行校验，检查是否合法
		//commentModel := &db.Comment{
		//	VideoId: VideoId,
		//	Content: CommentTextErr,
		//	UserId:  UserId,
		//	Status:  1,
		//}

	} else {
		CommentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if CommentId < 0 {
			SendResponse(c, errno.CommentIdErr, nil)
			return
		}
		_, err := rpc.CommentAction(ctx, &comment.CommentRequest{
			UserId:      UserId,
			VideoId:     VideoId,
			ActionType:  ActionType,
			CommentText: "",
			CommentId:   CommentId,
		})
		if err != nil {
			c.JSON(consts.StatusOK, comment.CommentResponse{
				StatusCode: -1,
				StatusMsg:  "失败",
				Comment:    nil,
			})
		} else {
			c.JSON(consts.StatusOK, comment.CommentResponse{
				StatusCode: 0,
				StatusMsg:  "成功",
				Comment:    nil,
			})
		}
	}
}
