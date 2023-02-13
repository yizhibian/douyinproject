package user_handler

import (
	"context"
	"douyin-user/idl/kitex_test/kitex_gen/douyinuser"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	//jwtTool "github.com/golang-jwt/jwt/v4"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var queryVar UserInfoParam
	if err := c.Bind(&queryVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if queryVar.Id == 0 {
		pack.SendBaseResponse(c, errno.ParamErr, nil)
		return
	}

	req := douyinuser.GetUserInfoRequest{
		UserId: queryVar.Id,
	}
	userInfo, err := rpc.GetUserInfo(context.Background(), &req)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if value, exists := c.Get("identity"); exists {
		fmt.Println("the id is ", value)
	}

	if userInfo.GetId() == 0 {
		pack.SendBaseResponse(c, errno.NewErrNo(errno.NilValueErrCode, "user doesnt exit"), nil)
		return
	}

	pack.SendUserInfoResponse(c, errno.Success, userInfo)
}
