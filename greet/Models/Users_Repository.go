package Models

import "time"

//用户视频类

type UserRepository struct {
	Id      int64  `json:"id,omitempty"` //视频唯一标识
	AuthorName  string `json:"author_name"` //视频作者信息
	VideoID int64 `json:"video_id"`
	PlayUrl string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"` // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` // 视频点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"` // 视频评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"` //  true-已点赞，false-未点赞
	Title         string `json:"title"`         // 视频标题
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

//用户视频类所对应的信息

func (table UserRepository) TableName() string  {
	return "user_repository"
}