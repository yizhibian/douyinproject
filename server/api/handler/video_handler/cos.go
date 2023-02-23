package video_handler

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
)

// BaseURL 访问各 API 所需的基础 URL
type BaseURL struct {
	// 访问 bucket, object 相关 API 的基础 URL（不包含 path 部分）: https://video-1312658769.cos.ap-guangzhou.myqcloud.com
	BucketURL *url.URL
	// 访问 service API 的基础 URL（不包含 path 部分）: https://cos.ap-guangzhou.myqcloud.com
	ServiceURL *url.URL
	// 访问 Batch API 的基础 URL （不包含 path 部分）: https://<UIN>.cos-control.<Region>.myqcloud.com
	BatchURL *url.URL
	// 访问 CI 的基础 URL （不包含 path 部分）: https://examplebucket-1250000000.ci.<Region>.myqcloud.com
	CIURL *url.URL
}

var secretID = "AKIDWq1zjEXbaO2QbV8Yspa4KL6brqZ1C6f0"
var secretKey = "B7oT4ScXDess36uT9sZ4VkocAUcWBakp"

func getCOSClient() *cos.Client {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://video-1312658769.cos.ap-guangzhou.myqcloud.com")
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: secretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return client
}

func getUrl(key string) *url.URL {
	client := getCOSClient()
	objectURL := client.Object.GetObjectURL(key)
	return objectURL
}

func uploadVideo(data *multipart.FileHeader, title string) (*http.Response, error) {
	file, _ := data.Open()
	client := getCOSClient()
	defer file.Close()
	resp, err := client.Object.Put(context.Background(), title+".mp4", file, nil)
	return resp.Response, err
}

func uploadCover(filepath string, title string) (*cos.Response, error) {
	client := getCOSClient()
	_, resp, err := client.Object.Upload(
		context.Background(), title, filepath, nil,
	)
	if err != nil {
		panic(err)
	}
	return resp, err
}
