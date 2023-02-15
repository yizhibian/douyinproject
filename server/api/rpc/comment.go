package rpc

import (
	"context"
	"douyin-user/idl/douyin_comment/kitex_gen/comment"
	"douyin-user/idl/douyin_comment/kitex_gen/comment/commentservice"
	"douyin-user/pkg/constants"
	"douyin-user/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}
func CommentAction(ctx context.Context, req *comment.CommentRequest) (*comment.CommentResponse, error) {
	resp, err := commentClient.Action(ctx, req)
	if err != nil {
		return nil, err
	}
	//if resp.BaseResp.StatusCode != 0 {
	//	return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	//}
	return resp, err
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	resp, err := commentClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	//if resp.BaseResp.StatusCode != 0 {
	//	return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	//}
	return resp, err
}
