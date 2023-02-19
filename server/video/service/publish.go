package service

import (
	"bytes"
	"context"
	"douyin-user/server/video/dal/db"
	"fmt"
	"log"
	"os/exec"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *GetListService {
	return &GetListService{
		ctx: ctx,
	}
}

func (s *GetListService) Publish(title string, data []byte, id int64) string {
	StatusMsg := "ok"
	reader := bytes.NewReader(data)
	resp, err := uploadVideo(reader, title)
	videoUrl := getUrl(title)
	coverTitle := title + ".jpg"
	filePath := "/opt/covers/" + coverTitle

	if err != nil {
		fmt.Println(resp)
		panic(err)
	}
	cmd := exec.Command(
		"ffmpeg", "-i", videoUrl.String(),
		"-ss", "00:00:02", "-vframes:v", "1",
		filePath,
	)
	if err := cmd.Run(); err != nil {
		log.Fatalln("Video cover generation failed")
	}
	resp2, _ := uploadCover(filePath, coverTitle)
	fmt.Println(resp2)
	coverUrl := getUrl(coverTitle)
	v := db.Video{
		AuthorId:      id,
		PlayUrl:       videoUrl.String(),
		CoverUrl:      coverUrl.String(),
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    true,
		Title:         title,
	}
	insertVideo(&v)
	StatusMsg = title + " uploaded successfully"
	return StatusMsg
}

func insertVideo(v *db.Video) {
	result := db.DB.Create(&v)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println(v.Id)
}
