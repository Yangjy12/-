package logic

import (
	"Bytecode_Project/greet/Models"
	"context"

	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	// todo: add your logic here and delete this line
	resp=new(types.FeedResponse)
	v:=new(Models.Video)
	get, err := l.svcCtx.Engine.Table("video").Select("*").Get(v)
	if !get {
		return nil,err
	}
	if err!=nil {
		return nil, err
	}
	resp.StatusMsg="请观看视频"
	resp.VideoList=v
	resp.StatusCode=0
	return
}
