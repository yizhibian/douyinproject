package router

import (
	"douyin-user/server/api/handler/favorite_handler"
	"douyin-user/server/api/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func FavoriteRegister(r *server.Hertz) {

	favorite := r.Group("/douyin/favorite")
	favorite.Use(mw.JwtMiddleware.MiddlewareFunc())
	favorite.POST("/action/", favorite_handler.LikeOrNot)

}
