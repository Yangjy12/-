package test

import (
	"Bytecode_Project/greet/Models"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

// 测试xorm的连接

func TestXorm(t *testing.T){
	//创建数据库引擎，连接到byte_code_project数据库
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/byte_code_project?charset=utf8")
	if err!=nil {
		t.Fatal(err)
	}
	//创建User用户对象
	//TableUser:=make([]*Models.User,0)
	//v := new(Models.Video)
	ur := new(Models.UserRepository)
	//m := new(Models.User)
	//创建User的数据表
	//err = engine.CreateTables(m)
	err = engine.CreateTables(ur)
	if err!=nil {
		t.Fatal(err)
	}
	/*err = engine.Find(&TableUser)//这里查出来的是地址
	if err!=nil {
		t.Fatal(err)
	}
	//将以json格式输出
	bytes, err:= json.Marshal(TableUser)//转换成BYTE数组
	if err!=nil {
		t.Fatal(err)
	}
	dst := new(bytes2.Buffer)
	err = json.Indent(dst, bytes, "", "  ") //转换为buffer
	if err!=nil {
		t.Fatal(err)
	}
	println(dst.String())*/
}
