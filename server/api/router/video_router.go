package router

import (
	"douyin-user/server/api/handler/video_handler"
	"douyin-user/server/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// UserRegister registers user module routers.
func VideoRegister(r *server.Hertz) {
	feedG := r.Group("/douyin/feed")
	//feedG.Use(mw.JwtMiddleware.MiddlewareFunc())
	feedG.GET("/", video_handler.Feed) //视频流接口

	pG := r.Group("/douyin/publish")
	pG.Use(mw.JwtMiddleware.MiddlewareFunc())
	pG.GET("/list/", video_handler.GetList)    //获取发布视频列表
	pG.POST("/action/", video_handler.Publish) //投稿
}
