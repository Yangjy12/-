package logic

import (
	"Bytecode_Project/greet/Models"
	"Bytecode_Project/greet/helper"
	"context"
	"log"

	"Bytecode_Project/greet/internal/svc"
	"Bytecode_Project/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	//判断用户名是否已经存在
	count, err := l.svcCtx.Engine.Where("name= ?", req.Username).Count(new(Models.User))
	if err!=nil {
		return nil, err
	}
	if count>0 {
		resp:=new(types.UserRegisterReply)
		resp.StatusCode=-1
		resp.StatusMsg="该用户名已注册，清更换用户名称"
		return resp,nil
	}
	//数据入库，开始注册信息
	user:=&Models.User{
		Name: req.Username,
		Password:helper.Md5(req.Password),
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err!=nil {
		return nil, err
	}
	log.Println("insert user row:",insert)
	//1.从数据库中查询当前用户信息，用于封装token和读取user的ID
	user2 := new(Models.User)
	_, err = l.svcCtx.Engine.Where("name= ? AND password= ?", req.Username, helper.Md5(req.Password)).Get(user2)
	if err!=nil {
		return nil, err
	}
	//2.生成token
	token, err := helper.GenerateToken(user2.Id, user2.Name)
	if err!=nil {
		return nil,err
	}
	resp2:=new(types.UserRegisterReply)
	resp2.StatusCode=0
	resp2.StatusMsg="注册成功！"
	resp2.UserID=user2.Id
	resp2.Token=token
	return resp2,nil
}
