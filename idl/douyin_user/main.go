package main

import (
	douyinuser "douyin-user/idl/douyin_user/kitex_gen/douyinuser/userserver"
	"log"
)

func main() {
	svr := douyinuser.NewServer(new(UserServerImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
