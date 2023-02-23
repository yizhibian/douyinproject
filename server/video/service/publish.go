package service

import (
	"context"
	"douyin-user/server/video/dal/db"
	"fmt"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *GetListService {
	return &GetListService{
		ctx: ctx,
	}
}

func (s *GetListService) Publish(v *db.Video) string {
	insertVideo(v)
	StatusMsg := " uploaded successfully"
	return StatusMsg
}

func insertVideo(v *db.Video) {
	result := db.DB.Create(&v)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println(v.Id)
}
