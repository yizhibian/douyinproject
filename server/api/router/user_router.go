package router

import (
	"douyin-user/server/api/handler/user_handler"
	"douyin-user/server/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// UserRegister registers user module routers.
func UserRegister(r *server.Hertz) {

	user1 := r.Group("/douyin/user")
	user1.POST("/login", mw.JwtMiddleware.LoginHandler)
	user1.POST("/register", user_handler.Register)

	//在这之前都是没有登陆拦截的
	//使用这个后即加入登陆中间件 需携带token访问
	user1.Use(mw.JwtMiddleware.MiddlewareFunc())
	user1.GET("/", user_handler.GetUserInfo)
}
