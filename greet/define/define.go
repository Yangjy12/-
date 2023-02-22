package define

import (
	"Bytecode_Project/greet/Models"
	"github.com/dgrijalva/jwt-go"
	"os"
)

//封装一下视频流接口的返回响应内容

type FeedStateModel struct {
	NextTime   *int64  `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"` // 返回状态描述
	VideoList  []Models.Video `json:"video_list"` // 视频列表
}

//为了生成token，封装使用到的UserClaim

type UserClaim struct {
	Id int64
	Name string
	jwt.StandardClaims
}

//为了发布列表的返回

type VideoList struct {
	Author        Models.User   `json:"author"`        // 视频作者信息
	CommentCount  int64  `json:"comment_count"` // 视频的评论总数
	CoverURL      string `json:"cover_url"`     // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"`// 视频的点赞总数
	ID            int64  `json:"id"`            // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`   // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`      // 视频播放地址
	Title         string `json:"title"`         // 视频标题
}

//jwt钥匙

var JwtKey="Bytecode_Project-key"

//TencentSecreKey 腾讯云对象存储
var TencentSecreKey = os.Getenv("TencentSecreKey")
var TencentSecretID = os.Getenv("TencentSecretID")
//cos传输桶
var CosBucket="https://bytecode-1315418729.cos.ap-nanjing.myqcloud.com"