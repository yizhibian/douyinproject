package pack

import (
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type FeedResponse struct {
	StatusCode int32               `json:"status_code"`
	StatusMsg  string              `json:"status_msg,omitempty"`
	VideoList  []douyinvideo.Video `json:"video_list,omitempty"`
	NextTime   int64               `json:"next_time,omitempty"`
}

// SendUserInfoResponse pack user register response
func SendFeedResponse(c *app.RequestContext, err error, userInfo *douyinuser.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserInfoResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		User:    *userInfo,
	})
}
