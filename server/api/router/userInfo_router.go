package router

import (
	"douyin-user/server/api/handler/user_handler"
	"douyin-user/server/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// UserInfoRegister registers user module routers.
func UserInfoRegister(r *server.Hertz) {
	user1 := r.Group("/usertest")
	user1.POST("/login", mw.JwtMiddleware.LoginHandler)
	user1.POST("/register", user_handler.Register)

	user1.Use(mw.JwtMiddleware.MiddlewareFunc())
	user1.GET("/", user_handler.GetUserInfo)
}
