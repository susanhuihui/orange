package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "orange/docs"
	_ "orange/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL) //注册mysql Driver//注册数据库连接
	//orm.RegisterDataBase("default", "mysql", "root:33c88fd2b0@tcp(127.0.0.1:3306)/onlineeducation?charset=utf8&loc=Local", 30) //本地数据库：数据库别名，driverName,对应链接字符串
	orm.RegisterDataBase("default", "mysql", "fankunedu:33C88FD2B0@tcp(101.200.75.221:3306)/onlineeducation?charset=utf8&loc=Local", 30) //发布后：数据库别名，driverName,对应链接字符串
	// 2015/12/15
	// 李向哲修改连库字符串为本地时区解析
}

func main() {
	beego.SessionOn = true
	beego.SessionProvider = "memory"
	beego.SessionGCMaxLifetime = 60
	beego.SessionName = "sessiontest"
	beego.SessionCookieLifeTime = 60
	beego.SessionAutoSetCookie = true
	beego.SessionSavePath = "/"

	//if beego.RunMode == "dev" {
	beego.DirectoryIndex = true
	beego.StaticDir["/swagger"] = "swagger"
	//}
	beego.Run()
}
