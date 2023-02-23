package pack

import (
	"douyin-user/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendBaseResponse pack base response
func SendBaseResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
