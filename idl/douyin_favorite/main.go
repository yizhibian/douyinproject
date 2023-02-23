package main

import (
	douyinfavorite "douyin-user/idl/douyin_favorite/kitex_gen/douyinfavorite/userserver"
	"log"
)

func main() {
	svr := douyinfavorite.NewServer(new(UserServerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
