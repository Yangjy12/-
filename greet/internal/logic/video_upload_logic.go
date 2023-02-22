package logic

import (
	"Bytecode_Project/greet/Models"
	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type VideoUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoUploadLogic {
	return &VideoUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoUploadLogic) VideoUpload(req *types.VideoUploadRequest,name string) (resp *types.VideoUploadReply, err error) {
	// todo: add your logic here and delete this line
	resp=new(types.VideoUploadReply)
	rp:=&Models.Video{
		Title: req.Title,
		PlayUrl: req.Path,
		Hash: req.Hash,
	}
	l.svcCtx.Engine.ShowSQL(true)//这样就可以看到运行时的sql语句
	_, err = l.svcCtx.Engine.Insert(rp)
	if err!=nil {
		resp.StatusCode=-1
		resp.StatusMsg="上传失败"
		return resp, err
	}
	video:=new(Models.Video)
	get, err := l.svcCtx.Engine.Where("hash=?", req.Hash).Get(video)
	if err!=nil {
		resp.StatusCode=-1
		resp.StatusMsg="上传失败"
		return resp, err
	}
	if !get {
		resp.StatusCode=-1
		resp.StatusMsg="上传失败"
		return resp, err
	}
	ur:=&Models.UserRepository{
		AuthorName: name,
		VideoID:    video.ID,
		PlayUrl:       video.PlayUrl,
		Title:         req.Title,
	}
	_, err = l.svcCtx.Engine.Insert(ur)

	resp.StatusMsg="上传成功"
	resp.StatusCode=0
	return resp,nil
}
