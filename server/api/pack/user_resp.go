package pack

import (
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_message"`
	Id      int64  `json:"user_id"`
	Token   string `json:"token"`
}

type UserInfoResponse struct {
	Code    int64           `json:"status_code"`
	Message string          `json:"status_message"`
	User    douyinuser.User `json:"user"`
}

// SendUserResponse pack user register response
func SendUserResponse(c *app.RequestContext, err error, resp *douyinuser.UserResponse) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Id:      resp.GetUserId(),
		Token:   resp.GetToken(),
	})
}

// SendUserInfoResponse pack user register response
func SendUserInfoResponse(c *app.RequestContext, err error, userInfo *douyinuser.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserInfoResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		User:    *userInfo,
	})
}
