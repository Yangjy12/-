package Models

import "time"

//视频类

type Video struct {
	ID int64 `json:"id"`
	Hash      string `json:"hash"`
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

//视频类所对应的信息

func (table Video) TableName() string  {
	return "video"
}