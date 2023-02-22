package svc

import (
	"Bytecode_Project/greet/Models"
	"Bytecode_Project/greet/internal/config"
	"Bytecode_Project/greet/internal/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB *redis.Client
	Auth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine:Models.Init(c.Mysql.DataSource),
		RDB:Models.InitRedis(c),
		Auth: middleware.NewAuthMiddleware().Handle,
	}
}
