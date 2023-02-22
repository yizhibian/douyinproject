package mw

import (
	"context"
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/handler/user_handler"
	"douyin-user/server/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	jwtTool "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJwt() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		//加入负载数据
		//其中data是这个 Authenticator 执行后的resp
		//然后往链里加这个负载数据
		//claims := token.Claims.(jwt.MapClaims)
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		//携带token访问时会执行到这里
		//这里结合Payload一起使用
		//其实就是拿出在Payload中封装的数据 然后放在上下文中便于取出使用吧
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[constants.IdentityKey]
		},
		//需改
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

			parse, _ := jwtTool.Parse(token, func(t *jwtTool.Token) (interface{}, error) {
				return []byte(constants.SecretKey), nil
			})

			mapClaims := parse.Claims.(jwtTool.MapClaims)
			userId := mapClaims["id"]

			c.JSON(consts.StatusOK, map[string]interface{}{
				"code":   errno.SuccessCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
				"id":     userId,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar user_handler.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			request := &douyinuser.CheckUserRequest{
				UserReq: &douyinuser.UserRequest{
					Username: loginVar.UserName,
					Password: loginVar.PassWord,
				},
			}

			return rpc.CheckUser(context.Background(), request)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt,form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
