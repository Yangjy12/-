package Models
//连接数据库方法init（）
import (
	"Bytecode_Project/greet/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

//将方法赋值给Engine变量，方便之后调用。

/*var Engine=Init()
func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/byte_code_project?charset=utf8")
	if err!=nil {
		log.Println("engine Error:",err)//log包的println会自动return
	}
	return engine
}*/

//将数据库进行抽取
var ctx = context.Background()
func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql",dataSource )
	if err!=nil {
		log.Println("engine Error:",err)//log包的println会自动return
	}
	return engine
}
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}