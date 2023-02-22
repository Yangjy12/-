package logic

import (
	"Bytecode_Project/greet/define"
	"context"

	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListRequest,userId string,name string) (resp *types.VideoListReply, err error) {
	// todo: add your logic here and delete this line
	resp=new(types.VideoListReply)
	vl:=new(define.VideoList)
	err = l.svcCtx.Engine.Table("user_repository").Where("name=?", name).
		Join("LEFT", "video", "user_repository.video_i_d=video.i_d").Find(&vl)
	if err!=nil {
		return nil,err
	}
	resp.StatusMsg="查询成功"
	resp.StatusCode=0
	resp.VideoList=vl
	return
}
