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

	//	var meetroomlist []bbb4go.MeetingRoom
	//	meetingroom1 := bbb4go.MeetingRoom{}
	//	meetingroom1.MeetingID_ = "asdf"
	//	meetingroom1.ModeratorPW_ = "1111"
	//	meetingroom2 := bbb4go.MeetingRoom{}
	//	meetingroom2.MeetingID_ = "asdf"
	//	meetingroom2.ModeratorPW_ = "1111"
	//	meetroomlist = append(meetroomlist, meetingroom1)
	//	meetroomlist = append(meetroomlist, meetingroom2)
	//	fmt.Println(meetroomlist)
	//	meetingroom1.End()
	//	meetingroom2.End()
	beego.Run()
}
