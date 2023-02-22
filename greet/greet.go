package main

import (
	"flag"
	"fmt"

	"Bytecode_Project/greet/internal/config"
	"Bytecode_Project/greet/internal/handler"
	"Bytecode_Project/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//上面的都是加载配置什么的
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)//此方法是注册路由和路由映射Handler，重点在这里

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
