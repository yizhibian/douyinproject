package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Register registers user module routers.
func Register(r *server.Hertz) {

	UserRegister(r)
	CommentRegister(r)
}
