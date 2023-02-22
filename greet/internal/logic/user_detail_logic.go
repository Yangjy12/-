package logic

import (
	"Bytecode_Project/greet/Models"
	"context"
	"errors"

	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	resp=&types.UserDetailReply{}
	m := new(Models.User)
	get, err := l.svcCtx.Engine.Where("id= ? ", req.User_id).Get(m)
	if err!=nil {
		resp.StatusMsg="string"
		resp.StatusCode=-1
		return resp, err
	}
	if !get {
		resp.StatusMsg="\"该用户不存在\""
		resp.StatusCode=-1
		return resp, errors.New("该用户不存在")
	}
	resp.StatusMsg="查看用户信息"
	resp.StatusCode=0
	resp.User=&Models.User{
		Id: m.Id,
		Name: m.Name,
		FollowCount: m.FollowCount,
		FollowerCount: m.FollowerCount,
		IsFollow: m.IsFollow,
		Avatar: m.Avatar,
		BackgroundImage: m.BackgroundImage,
		Signature: m.Signature,
		TotalFavorited: m.TotalFavorited,
		WorkCount: m.WorkCount,
		FavoriteCount: m.FavoriteCount,
	}
	return resp,nil
}
