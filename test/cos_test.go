package test

import (
	"Bytecode_Project/greet/define"
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestCos(t *testing.T) {
	u, _ := url.Parse("https://bytecode-1315418729.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: define.TencentSecreKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	name := "抖音视频/test.mp4"

	_,_,err:=c.Object.Upload(context.Background(),name,"./video/VID_20230115_213122.mp4",nil)
	if err!=nil {
		panic(err)
	}
}
func TestFilerUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://bytecode-1315418729.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: define.TencentSecreKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	name := "抖音视频/test.mp4"

	file, err2 := os.ReadFile("./video/VID_20230115_213122.mp4")
	if err2!=nil {
		return
	}
	_,err:=c.Object.Put(context.Background(),name,bytes.NewReader(file),nil)
	if err!=nil {
		panic(err)
	}
}
