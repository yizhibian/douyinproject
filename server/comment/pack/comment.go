package pack

import (
	"context"
	"douyin-user/idl/douyin_comment/kitex_gen/comment"
	"douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	"douyin-user/server/comment/dal/db"
	"douyin-user/server/comment/rpc"
)

func Comment(m *db.Comment, ctx context.Context) *comment.Comment {
	if m == nil {
		return nil
	}
	info, _ := rpc.GetUserInfo(ctx, &douyinuser.GetUserInfoRequest{
		UserId: m.UserId,
	})

	return &comment.Comment{
		Id: int64(m.ID),
		User: &comment.User{
			Id:            info.Id,
			FollowCount:   info.FollowerCount,
			FollowerCount: info.FollowerCount,
			Name:          info.Name,
		},
		Content:    m.Content,
		CreateDate: m.CreatedAt.String(),
	}
}

// Notes pack list of note info
func Comments(ms []*db.Comment, ctx context.Context) []*comment.Comment {

	comments := make([]*comment.Comment, 0)
	for _, m := range ms {
		if n := Comment(m, ctx); n != nil {
			comments = append(comments, n)
		}
	}
	return comments
}
