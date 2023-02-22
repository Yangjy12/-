package logic

import (
	"Bytecode_Project/greet/Models"
	"Bytecode_Project/greet/helper"
	"context"
	"errors"

	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserLoginReply)
	user := new(Models.User)
	//1.从数据库中查询当前用户
	get, err := l.svcCtx.Engine.Where("name= ? AND password= ?", req.Username, helper.Md5(req.Password)).Get(user)
	if err!=nil {
		resp.StatusCode=-1
		resp.StatusMsg="出错啦"
		return resp, err
	}
	if !get {
		resp.StatusCode=-1
		resp.StatusMsg="用户名或密码错误"
		return resp,errors.New("用户名或密码错误")
	}
	//2.生成token
	token, err := helper.GenerateToken(user.Id, user.Name)
	if err!=nil {
		return nil,err
	}
	resp.StatusCode=0
	resp.StatusMsg="登陆成功！"
	resp.UserID=user.Id
	resp.Token=token
	return resp,nil
}
