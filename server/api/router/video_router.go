package router

import (
	"douyin-user/server/api/handler/video_handler"
	_ "douyin-user/server/api/handler/video_handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// UserRegister registers user module routers.
func VideoRegister(r *server.Hertz) {

	feedG := r.Group("/feed")
	feedG.GET("/", video_handler.Feed) //视频流接口
	pG := r.Group("/publish")
	pG.GET("/list", video_handler.GetList)    //获取发布视频列表
	pG.POST("/action", video_handler.Publish) //投稿
}
