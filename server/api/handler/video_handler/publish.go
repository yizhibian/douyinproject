package video_handler

import (
	"context"
	"douyin-user/idl/douyin_video/kitex_gen/douyinvideo"
	"douyin-user/pkg/errno"
	"douyin-user/server/api/pack"
	"douyin-user/server/api/rpc"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	data, err2 := c.FormFile("data")
	title := c.PostForm("title")
	if err2 != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err2), nil)
		return
	}
	_, err := uploadVideo(data, title)
	if err != nil {
		fmt.Println("上传失败==")
		return
	}
	videoUrl := getUrl(title + ".mp4")
	coverTitle := strconv.FormatInt(time.Now().Unix(), 10) + "." + title + ".jpg"
	filePath := "/opt/covers/" + coverTitle
	cmd := exec.Command(
		"ffmpeg", "-i", videoUrl.String(),
		"-ss", "00:00:02", "-vframes:v", "1",
		filePath,
	)
	if err := cmd.Run(); err != nil {
		log.Printf("Video cover generation failed")
		return
	}
	uploadCover(filePath, coverTitle)
	coverUrl := getUrl(coverTitle)
	id, exists := c.Get("identity")
	if !exists {
		log.Println("cant get id")
	}
	log.Printf("id======%#v\n", id)
	v := douyinvideo.PublishRequest{
		AuthorId: int64(id.(float64)),
		Title:    title,
		PlayUrl:  videoUrl.String(),
		CoverUrl: coverUrl.String(),
	}
	r, err := rpc.Publish(ctx, &v)
	if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(consts.StatusOK, r)
	return
}
