package models

//Model层公共文件
import ()

//域名X
var OnlineUrl string = "10.10.0.2:8080" //本地测试域名

// var OnlineUrl string = "www.fankunedu.com" //发布后域名

var CreateStrVerify string = "80c8c5bb2dd7db8e8652e97f42c4b37d"

//白板服务器所在的路径
var OnlineClassUrl string = "http://meeting.fankunedu.com/bigbluebutton/api/create"

//进入路径
var OnlineInClassUrl string = "http://meeting.fankunedu.com/bigbluebutton/api/join"

var TotalMinute int = 50 //每节课总课时

var AdvanceMinutes int = 10 //提前几分钟可以进入课程

var TradingWayId int = 1 //余额充值类型值
