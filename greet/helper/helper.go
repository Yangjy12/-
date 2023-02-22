package helper

import (
	"Bytecode_Project/greet/define"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path"
)
//数据库密码加密，使用md5

func Md5(s string) string  {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
func GetUUID() string {
	return uuid.NewV4().String()
}
//生成token方法

func GenerateToken(id int64,name string) (string,error) {
	uc:=define.UserClaim{
		Id:id,
		Name: name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc) //生成token
	//给token进行加密
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err!=nil {
		return "", err
	}
	return tokenString,nil
}

//CosUpload 文件上传到腾讯云
func CosUpload(r *http.Request) (string,error) {
		u, _ := url.Parse(define.CosBucket)
		b := &cos.BaseURL{BucketURL: u}
		c := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  define.TencentSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
				SecretKey: define.TencentSecreKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			},
		})

	file, fileHeader, err := r.FormFile("data")
	name := "抖音视频/"+GetUUID()+path.Ext(fileHeader.Filename)
	_, err = c.Object.Put(context.Background(), name, file,nil)
	if err != nil {
		panic(err)
	}
	return define.CosBucket+"/"+name,nil
}
//AnalyzeToken Token解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc:=new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil,err
	}
	if !claims .Valid{
		return uc,errors.New("token is invalid")
	}
	return uc,err
}