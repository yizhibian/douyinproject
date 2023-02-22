package main

import (
	comment "douyin-user/idl/douyin_comment/kitex_gen/comment/commentservice"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/tracer"
	"douyin-user/server/comment/dal"
	"douyin-user/server/comment/rpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	rpc.InitRPC()
	dal.Init()
	tracer.InitJaeger(constants.CommentServiceName)

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		log.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")

	svr := comment.NewServer(new(CommentServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
