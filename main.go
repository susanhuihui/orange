package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "orange/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL) //注册mysql Driver//注册数据库连接
	//orm.RegisterDataBase("default", "mysql", "root:33c88fd2b0@tcp(127.0.0.1:3306)/onlineeducation") //本地数据库：数据库别名，driverName,对应链接字符串
	//orm.RegisterDataBase("default", "mysql", "fankunedu:33C88FD2B0@tcp(101.200.75.221:3306)/onlineeducation") //本地测试：数据库别名，driverName,对应链接字符串
	orm.RegisterDataBase("default", "mysql", "root:33C88FD2B0@tcp(127.0.0.1:3306)/onlineeducation") //发布后：数据库别名，driverName,对应链接字符串
}

func main() {
	beego.SessionOn = true
	beego.SessionProvider = "memory"
	beego.SessionGCMaxLifetime = 60
	beego.SessionName = "sessiontest"
	beego.SessionCookieLifeTime = 60
	beego.SessionAutoSetCookie = true
	beego.SessionSavePath = "/"
	beego.Run()
}
