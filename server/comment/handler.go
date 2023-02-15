package main

import (
	"context"
	"douyin-user/idl/douyin_comment/kitex_gen/comment"
	"douyin-user/server/comment/service"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Action implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Action(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentResponse)
	resp, err = service.Action(ctx, req)
	return resp, err
}

// List implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) List(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentListResponse)
	resp, err = service.List(ctx, req)
	return resp, err
}
