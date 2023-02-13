// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package user_handler

import (
	"context"
	"douyin-user/idl/kitex_test/kitex_gen/douyinuser"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Register register user info
func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		pack.SendBaseResponse(c, errno.ParamErr, nil)
		return
	}

	userReq := douyinuser.NewUserRequest()
	userReq.SetPassword(registerVar.PassWord)
	userReq.SetUsername(registerVar.UserName)
	uid, err := rpc.CreateUser(context.Background(), &douyinuser.CreateUserRequest{
		UserReq: userReq,
	})
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	expire := time.Now().Add(time.Hour)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()
	tokenString, err := token.SignedString([]byte(constants.SecretKey))

	resp := douyinuser.NewUserResponse()
	resp.SetUserId(uid)
	resp.SetToken(tokenString)

	pack.SendUserResponse(c, errno.Success, resp)
}
