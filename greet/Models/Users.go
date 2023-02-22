package Models

import "time"

//用户信息类

type User struct {
	Id            int64  `json:"id,omitempty"` // 用户id
	Name          string `json:"name,omitempty"` // 用户名称
	Password	  string`json:"password"`
	Email 		  string `json:"email"`
	Avatar          string `json:"avatar"`          // 用户头像
	BackgroundImage string `json:"background_image"`// 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`  // 喜欢数
	FollowCount     int64  `json:"follow_count"`    // 关注总数
	FollowerCount   int64  `json:"follower_count"`  // 粉丝总数
	IsFollow        bool   `json:"is_follow"`       // true-已关注，false-未关注
	Signature       string `json:"signature"`       // 个人简介
	TotalFavorited  string `json:"total_favorited"` // 获赞数量
	WorkCount       int64  `json:"work_count"`      // 作品数
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

//用户信息所对应的表名

func (table User) TableName() string {
	return "user"
}