package main

import (
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo/videoserver"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/tracer"
	"douyin-user/server/video/dal"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	dal.Init()
	tracer.InitJaeger(constants.VideoServiceName)
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8989")
	svr := videoserver.NewServer(new(VideoServerImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}), // server name
		server.WithRegistry(r), // registry
		server.WithServiceAddr(addr),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
