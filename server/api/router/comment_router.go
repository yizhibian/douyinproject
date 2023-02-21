package router

import (
	handlers "douyin-user/server/api/handler/comment_handler"
	"douyin-user/server/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// UserRegister registers user module routers.
func CommentRegister(r *server.Hertz) {
	comment1 := r.Group("/douyin/comment")
	comment1.Use(mw.JwtMiddleware.MiddlewareFunc())
	comment1.POST("/action", handlers.CommentAction)
	comment1.GET("/list", handlers.CommentList)
}
