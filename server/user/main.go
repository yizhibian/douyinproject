package main

import (
	douyinuser "douyin-user/idl/douyin_user/kitex_gen/douyinuser/userserver"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/middleware"
	"douyin-user/server/user/dal"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
	//tracer.InitJaeger(constants.UserServiceName)
}

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := douyinuser.NewServer(new(UserServerImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()), // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err)
	}
}
