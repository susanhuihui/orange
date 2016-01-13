package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"time"
	//"github.com/ascoders/alipay"
	"github.com/astaxie/beego"
	"orange/models"
	"os"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

//首页展示
func (c *MainController) Get() {
	c.Data["Website"] = models.OnlineUrl
	c.Ctx.SetCookie("OnlineUrl", models.OnlineUrl)
	c.TplNames = "index.tpl" //首页
}

// @Title 登录方法
// @Description 根据用户名和密码登录网站
// @Param	username form string true 用户名
// @Param	password form string true 密码
// @Success 200 {object} models.Userinformation
// @Failure 403
// @router /Logins/ [post]
func (c *MainController) Logins() {
	//c.Ctx.ResponseWriter = http.ResponseWriter 意义上等同
	var name []string = c.Ctx.Input.Request.Form["username"] //获取登录名
	textname := name[0]
	var pass []string = c.Ctx.Input.Request.Form["password"] //获取密码
	textpass := pass[0]
	c.Data["username"] = textname
	c.Data["password"] = textpass
	var vuser *models.Userinformation
	fmt.Println("用户名，密码：")
	fmt.Println(textname, textpass)
	vuser, err := models.GetUserinformationLogin(textname, textpass)
	if err == nil && vuser != nil {
		fmt.Println(vuser)
		c.Data["Website"] = models.OnlineUrl
		c.Ctx.SetCookie("username", textname)
		c.Ctx.SetCookie("userid", strconv.Itoa(vuser.Id))
		c.Ctx.SetCookie("identityid", strconv.Itoa(vuser.IdentityId))
		c.Ctx.SetCookie("AvatarPath", vuser.AvatarPath)
		fmt.Println(vuser.AvatarPath)
		c.TplNames = "index.tpl"
	} else {
		vphoneuser, errph := models.GetUserinformationLoginPhone(textname, textpass)
		if errph == nil && vphoneuser != nil {
			fmt.Println(vphoneuser)
			c.Data["Website"] = models.OnlineUrl
			c.Ctx.SetCookie("username", textname)
			c.Ctx.SetCookie("userid", strconv.Itoa(vphoneuser.Id))
			c.Ctx.SetCookie("identityid", strconv.Itoa(vphoneuser.IdentityId))
			c.Ctx.SetCookie("AvatarPath", vphoneuser.AvatarPath)
			fmt.Println(vphoneuser.AvatarPath)
			c.TplNames = "index.tpl"
		} else {
			fmt.Println(err)
			c.Data["blockdiv"] = "none"
			c.TplNames = "404.html"
		}
	}
}

// @Title 登录方法
// @Description 根据用户实体json字符串登录
// @Param json form string true json用户实体字符串
// @Success >0 {int} Id
// @Failure 0 获取用户失败
// @Failure -1 用户名昵称不存在
// @Failure -2 用户名昵称存在密码不正确
// @router /LoginUser/ [post]
func (c *MainController) LoginUser() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	//	fmt.Println("接到信息为：")
	//	fmt.Println(jsonS)
	var v models.Userinformation
	json.Unmarshal([]byte(jsonS), &v)
	var vuser *models.Userinformation
	//	fmt.Println("用户名，密码：")
	//	fmt.Println(v.UserName, v.LoginPassword)
	vuser, err := models.GetUserinformationLogin(v.UserName, v.LoginPassword)
	//fmt.Println(vuser)
	var usertrue string = "0" //0获取用户失败 -2用户名昵称存在密码不正确 -1用户名昵称不存在
	if err == nil && vuser != nil && vuser.Id > 0 {
		//fmt.Println(vuser)
		c.Data["Website"] = models.OnlineUrl
		c.Ctx.SetCookie("username", vuser.UserName)
		c.Ctx.SetCookie("userid", strconv.Itoa(vuser.Id))
		c.Ctx.SetCookie("identityid", strconv.Itoa(vuser.IdentityId))
		c.Ctx.SetCookie("AvatarPath", vuser.AvatarPath)
		c.Data["json"] = map[string]interface{}{"Id": vuser.Id, "IdentityId": vuser.IdentityId, "AvatarPath": vuser.AvatarPath}
	} else {
		vphoneuser, errph := models.GetUserinformationLoginPhone(v.UserName, v.LoginPassword)
		if errph == nil && vphoneuser != nil {
			c.Data["Website"] = models.OnlineUrl
			c.Ctx.SetCookie("username", vphoneuser.UserName)
			c.Ctx.SetCookie("userid", strconv.Itoa(vphoneuser.Id))
			c.Ctx.SetCookie("identityid", strconv.Itoa(vphoneuser.IdentityId))
			c.Ctx.SetCookie("AvatarPath", vphoneuser.AvatarPath)
			c.Data["json"] = map[string]interface{}{"Id": vphoneuser.Id, "IdentityId": vphoneuser.IdentityId, "AvatarPath": vphoneuser.AvatarPath}
		} else {
			getuserbyname, nameerr := models.GetUserinformationByUserName(v.UserName)
			if nameerr == nil && getuserbyname != nil {
				usertrue = "-2" //用户名昵称存在密码不正确
			} else if getuserbyname == nil {
				usertrue = "-1" //用户名昵称不存在
			} else if nameerr != nil {
				usertrue = "0" //获取用户失败
			}
			if usertrue != "-2" {
				getuserbyphone, phoneerr := models.GetUserinformationByPhone(v.IphoneNum)
				if phoneerr == nil && getuserbyphone != nil {
					usertrue = "-2" //用户名昵称存在密码不正确
				} else if getuserbyphone == nil {
					usertrue = "-1" //用户名昵称不存在
				} else if phoneerr != nil {
					usertrue = "0" //获取用户失败
				}
			} //map[string]interface{}{"id": id, "state": 1}
			c.Data["json"] = map[string]interface{}{"Id": usertrue, "IdentityId": 0, "AvatarPath": ""}
		}
	}
	c.ServeJson()
}

// @Title OutLogins
// @Description 退出登录方法
// @Success 200 {string} OK
// @router /OutLogins/ [get]
func (c *MainController) OutLogins() {
	c.Data["Website"] = models.OnlineUrl
	c.Ctx.SetCookie("username", "")
	c.Ctx.SetCookie("userid", "")
	c.Ctx.SetCookie("identityid", "")
	c.Ctx.SetCookie("AvatarPath", "")
	c.Data["json"] = "OK"
	c.ServeJson()
	//c.TplNames = "index.tpl"
}

// @Title Registered
// @Description 注册方法
// @router /Registered/ [get]
func (c *MainController) Registered() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "register.html" //跳到注册页面
}

// @Title QuestionsCenter
// @Description 跳页到问答中心模块
// @router /QuestionsCenter/ [get]
func (c *MainController) QuestionsCenter() {
	c.Data["Website"] = models.OnlineUrl
	wendalist, _ := models.GetQuestionaskByJingCaiCount()
	c.Data["wendacount"] = wendalist
	c.TplNames = "problem_list.html" //跳到问答中心
}

// @Title UserTeacher
// @Description 跳页到老师个人中心
// @Param tapid query int true 老师用户主键id
// @router /UserTeacher/:tapid [get]
func (c *MainController) UserTeacher() {
	c.Data["Website"] = models.OnlineUrl
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	tapidStr := c.Ctx.Input.Params[":tapid"]
	tapid, _ := strconv.Atoi(tapidStr) //获取tapid
	//基本信息展示
	var showTeacher models.UserinformationTeacher
	var err error
	showTeacher, err = models.GetUserinformationTeacher(stuuserid)
	if err == nil {
		c.Data["UserName"] = showTeacher.UserName
		c.Data["StudentUserid"] = "00000" + strconv.Itoa(showTeacher.Id)
		c.Data["AvatarPath"] = showTeacher.AvatarPath
		if showTeacher.UserSex == "" {
			c.Data["StuSex"] = ""
		} else {
			c.Data["StuSex"] = "| " + showTeacher.UserSex
		}
		if showTeacher.SchoolName == "" {
			c.Data["SchoolName"] = ""
		} else {
			c.Data["SchoolName"] = "| " + showTeacher.SchoolName
			c.Data["SchoolName2"] = showTeacher.SchoolName
		}
		c.Data["IdentityName"] = showTeacher.IdentityName
		//计算课时
		fa, _ := strconv.ParseFloat(strconv.Itoa(showTeacher.AllDate), 64)
		allhour := fmt.Sprintf("%.1f", fa/60)
		c.Data["AllDate"] = allhour
		c.Data["AllCount"] = strconv.Itoa(showTeacher.AllCount)
		c.Data["AllPerson"] = strconv.Itoa(showTeacher.AllPerson)
		c.Data["Professional"] = showTeacher.Professional
		c.Data["CourseName"] = showTeacher.CourseName
		c.Data["UserHobby"] = showTeacher.UserHobby
		c.Data["BriefIntroduction"] = showTeacher.BriefIntroduction //学习难点
	}
	zhucourse, fuerr := models.GetRemedialcoursesMain(stuuserid, 0)
	var fuzhu string = ""
	if zhucourse != nil && fuerr == nil {
		for i := 0; i < len(zhucourse); i++ {
			fuzhu += zhucourse[i].CourseName
			fuzhu += " "
		}
	}
	c.Data["fuzhuCourse"] = fuzhu

	var zijin models.Accountfunds
	zijin, _ = models.GetAccountfundsByuid(stuuserid)
	c.Data["Balance"] = zijin.Balance
	fmt.Println(zijin.Balance)
	//获取老师正在申请的金额总和
	nowtixianmoney, xerr := models.GetAmountrecordsTMcountByUid(stuuserid)
	if xerr == nil {
		c.Data["Shenqingmoney"] = nowtixianmoney
	} else {
		c.Data["Shenqingmoney"] = 0
	}

	//列表信息展示
	//1.预约课程
	yuyuelist, _ := models.GetOnlinecoursebookingByTidCount(stuuserid)
	c.Data["yuyuecount"] = yuyuelist
	//全部课程
	quanbulist, _ := models.GetOnlinecourserecordByTidCount(stuuserid)
	c.Data["quanbucount"] = quanbulist
	//我的问题
	wentilist, _ := models.GetQuestionaskByTidCount(stuuserid)
	c.Data["tiwencount"] = wentilist
	//资金-交易记录
	jiaoyilist, _ := models.GetTransactionrecordsByTidCount(stuuserid)
	c.Data["jiaoyicount"] = jiaoyilist
	//资金-提现记录
	tixianlist, _ := models.GetAmountrecordsTixianByUseridCount(stuuserid)
	c.Data["tixiancount"] = tixianlist
	//我的评价
	pingjialist, _ := models.GetOnlinecourseevaluationByTidCount(stuuserid)
	c.Data["pingjiacount"] = pingjialist
	//我的留言
	liuyanlist, _ := models.GetUsermessageByTidCount(stuuserid)
	c.Data["liuyancount"] = liuyanlist
	//我的试听
	shitinglist, _ := models.OnlineTryListenByTidCount(stuuserid)
	c.Data["listencount"] = shitinglist

	c.Data["tapNum"] = tapid        //显示第几个tap
	c.TplNames = "teachermain.html" //跳到老师个人中心
}

// @Title UserStudent
// @Description 跳页到学生个人中心
// @Param tapid query int true 学生用户主键id
// @router /UserStudent/:tapid [get]
func (c *MainController) UserStudent() {
	c.Data["Website"] = models.OnlineUrl
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	tapidStr := c.Ctx.Input.Params[":tapid"]
	tapid, _ := strconv.Atoi(tapidStr) //获取tapid
	//基本信息展示
	var showStudent models.UserinformationStudent
	var err error
	showStudent, err = models.GetUserinformationStudent(stuuserid)
	if err == nil {
		c.Data["UserName"] = showStudent.UserName
		c.Data["StudentUserid"] = "00000" + strconv.Itoa(showStudent.Id)
		//计算课时
		fa, _ := strconv.ParseFloat(strconv.Itoa(showStudent.AllDate), 64)
		allhour := fmt.Sprintf("%.1f", fa/60)
		c.Data["AllDate"] = allhour
		c.Data["AllCount"] = strconv.Itoa(showStudent.AllCount)
		c.Data["AllPerson"] = strconv.Itoa(showStudent.AllPerson)
		c.Data["IdentityName"] = showStudent.IdentityName
		if showStudent.UserSex == "" {
			c.Data["StuSex"] = ""
		} else {
			c.Data["StuSex"] = "| " + showStudent.UserSex
		}
		if showStudent.SchoolName == "" {
			c.Data["SchoolName"] = ""
		} else {
			c.Data["SchoolName"] = "| " + showStudent.SchoolName
			c.Data["SchoolName2"] = showStudent.SchoolName
		}
		c.Data["AgeName"] = showStudent.AgeName
		if showStudent.LevelYear > 0 {
			c.Data["LevelYear"] = strconv.Itoa(showStudent.LevelYear)
		} else {
			c.Data["LevelYear"] = ""
		}
		c.Data["StudyDifficult"] = showStudent.StudyDifficult //学习难点
		c.Data["AvatarPath"] = showStudent.AvatarPath
	}
	var stuAllRemed []models.RemedialcoursesMain
	var errows error
	stuAllRemed, errows = models.GetRemedialcoursesMain(stuuserid, 0)
	if stuAllRemed != nil && errows == nil {
		var sturemed string
		for i := 0; i < len(stuAllRemed); i++ {
			sturemed += stuAllRemed[i].CourseName + "  "
		}
		if sturemed != "" {
			c.Data["CourseName"] = sturemed //补习科目
		}
	}

	var zijin models.Accountfunds
	zijin, _ = models.GetAccountfundsByuid(stuuserid)
	c.Data["Balance"] = zijin.Balance
	//fmt.Println(zijin.Balance)
	dongjiezijin, _ := models.GetFrozenFundsByUserid(stuuserid)
	//fmt.Println("冻结资金总和：")
	//fmt.Println(dongjiezijin.FrozenMoney)
	c.Data["FrozenMoney"] = dongjiezijin.FrozenMoney

	//列表信息展示
	//1.预约课程
	yuyuelist, _ := models.GetOnlinecoursebookingByUidCount(stuuserid)
	c.Data["yuyuecount"] = yuyuelist
	//全部课程
	quanbulist, _ := models.GetOnlinecourserecordByUidCount(stuuserid)
	c.Data["quanbucount"] = quanbulist
	//我的问题
	wentilist, _ := models.GetQuestionaskBySidCount(stuuserid)
	c.Data["tiwencount"] = wentilist
	//资金-交易记录
	jiaoyilist, _ := models.GetTransactionrecordsBySidCount(stuuserid)
	c.Data["jiaoyicount"] = jiaoyilist
	//资金-提现记录
	tixianlist, _ := models.GetAmountrecordsTixianByUseridCount(stuuserid)
	c.Data["tixiancount"] = tixianlist
	//资金-充值记录
	chongzhilist, _ := models.GetAmountrecordsByUseridCount(0, stuuserid)
	c.Data["chongzhicount"] = chongzhilist
	//我的评价
	pingjialist, _ := models.GetOnlineCourseEvaluationBySidCount(stuuserid)
	c.Data["pingjiacount"] = pingjialist
	//我的留言
	liuyanlist, _ := models.GetUsermessageBySidCount(stuuserid)
	c.Data["liuyancount"] = liuyanlist
	beego.Debug("留言内容: ", liuyanlist)
	//我的关注
	guanzhulist, _ := models.GetRelationsByUidCount(stuuserid, "关注")
	c.Data["guanzhucount"] = guanzhulist
	tuiyuelist, _ := models.GetOnlinecoursebookingBySidNotOnCount(stuuserid)
	c.Data["tuiyuecount"] = tuiyuelist

	c.Data["tapNum"] = tapid        //显示第几个tap
	c.TplNames = "studentmain.html" //跳到学生个人中心
}

// @Title OwnerUser
// @Description 跳页到管理员个人中心
// @Param tapid query int true 管理员用户主键id
// @router /OwnerUser/:tapid [get]
func (c *MainController) OwnerUser() {
	tapidStr := c.Ctx.Input.Params[":tapid"]
	c.Data["Website"] = models.OnlineUrl
	account, err := models.GetAmountrecordsAllTCount()
	if err == nil && account > 0 {
		c.Data["tixiancount"] = account
	} else {
		c.Data["tixiancount"] = 0
	}
	tuijiancount, errt := models.GetAmountrecordsAllTCount()
	if errt == nil && tuijiancount > 0 {
		c.Data["tuijiancount"] = tuijiancount
	} else {
		c.Data["tuijiancount"] = 0
	}
	allteachercount, errat := models.GetUserinformationTeacherAllCount()
	if errat == nil && allteachercount > 0 {
		c.Data["allteachercount"] = allteachercount
	} else {
		c.Data["allteachercount"] = 0
	}
	c.Data["tapNum"] = tapidStr //显示第几个tap
	c.TplNames = "adminhtml.html"
}

// @Title QuestionsCenter
// @Description 学生个人中心-查看全部课程中自己给老师的一条评价
// @Param evalid query int true 评价信息主键id
// @router /GetOnLineEvaluation/:evalid [get]
func (c *MainController) GetOnLineEvaluation() {
	c.Data["Website"] = models.OnlineUrl
	evalidStr := c.Ctx.Input.Params[":evalid"]
	evalid, _ := strconv.Atoi(evalidStr)
	c.Data["seeOradd"] = 1
	onlineev, everr := models.GetOnlinecourseevaluationById(evalid)
	if onlineev != nil && everr == nil {
		onrecord, recerr := models.GetOnlinecourserecordById(onlineev.OCRId)
		if onrecord != nil && recerr == nil {
			userinfo, usererr := models.GetUserinformationTeacher(onrecord.UserIdPassive)
			if usererr == nil {
				c.Data["AvatarPath"] = userinfo.AvatarPath
				c.Data["UserName"] = userinfo.UserName
				c.Data["SchoolName"] = userinfo.SchoolName
				c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
				c.Data["TeacherId"] = userinfo.Id
			}
		}
		//解答清晰度
		var starqingxistr string = ""
		var qingxi int = onlineev.StartClear
		for i := 0; i < 5; i++ {
			if qingxi > 0 {
				for a := 0; a < qingxi; a++ {
					starqingxistr += `<img style="height: 25px;" src="images/look.png" />`
					i += 1
				}
				qingxi = 0
			}
			if i < 5 {
				starqingxistr += `<img style="height: 25px;" src="images/looked.png" />`
			}
		}
		c.Data["StartClear"] = starqingxistr
		//解答态度
		var startaidustr string = ""
		var taidu int = onlineev.StartEfficiency
		for i := 0; i < 5; i++ {
			if taidu > 0 {
				for a := 0; a < taidu; a++ {
					startaidustr += `<img style="height: 25px;" src="images/look.png" />`
					i++
				}
				taidu = 0
			}
			if i < 5 {
				startaidustr += `<img style="height: 25px;" src="images/looked.png" />`
			}
		}
		c.Data["StartEfficiency"] = startaidustr

		c.Data["ReviewContent"] = onlineev.ReviewContent
	}
	c.TplNames = "evaluationmainstudents.html" //
}

// @Title GetOnLineEvaluationTeacher
// @Description 老师个人中心-查看学生给自己的一条评价
// @Param evalid query int true 评价信息主键id
// @router /GetOnLineEvaluationTeacher/:evalid [get]
func (c *MainController) GetOnLineEvaluationTeacher() {
	c.Data["Website"] = models.OnlineUrl
	evalidStr := c.Ctx.Input.Params[":evalid"]
	evalid, _ := strconv.Atoi(evalidStr)
	c.Data["seeOradd"] = 1
	onlineev, everr := models.GetOnlinecourseevaluationById(evalid)
	if onlineev != nil && everr == nil {
		onrecord, recerr := models.GetOnlinecourserecordById(onlineev.OCRId)
		if onrecord != nil && recerr == nil {
			userinfo, usererr := models.GetUserinformationTeacher(onrecord.UserIdActive)
			if usererr == nil {
				c.Data["AvatarPath"] = userinfo.AvatarPath
				c.Data["UserName"] = userinfo.UserName
				c.Data["SchoolName"] = userinfo.SchoolName
				c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
				c.Data["TeacherId"] = userinfo.Id
			}
		}
		//解答清晰度
		var starqingxistr string = ""
		var qingxi int = onlineev.StartClear
		for i := 0; i < 5; i++ {
			if qingxi > 0 {
				for a := 0; a < qingxi; a++ {
					starqingxistr += `<img style="height: 25px;" src="images/look.png" />`
					i += 1
				}
				qingxi = 0
			}
			if i < 5 {
				starqingxistr += `<img style="height: 25px;" src="images/looked.png" />`
			}
		}
		c.Data["StartClear"] = starqingxistr
		//解答态度
		var startaidustr string = ""
		var taidu int = onlineev.StartEfficiency
		for i := 0; i < 5; i++ {
			if taidu > 0 {
				for a := 0; a < taidu; a++ {
					startaidustr += `<img style="height: 25px;" src="images/look.png" />`
					i += 1
				}
				taidu = 0
			}
			if i < 5 {
				startaidustr += `<img style="height: 25px;" src="images/looked.png" />`
			}
		}
		c.Data["StartEfficiency"] = startaidustr

		c.Data["ReviewContent"] = onlineev.ReviewContent
	}
	c.TplNames = "evaluationmainstudents.html" //
}

// @Title QuestionsCenter
// @Description 学生个人中心-添加一条全部课程中的一条评价
// @Param classid query int true 在线课程主键id
// @router /AddOnLineEvaluation/:classid [get]
func (c *MainController) AddOnLineEvaluation() {
	c.Data["Website"] = models.OnlineUrl
	classidStr := c.Ctx.Input.Params[":classid"]
	classid, _ := strconv.Atoi(classidStr)
	c.Data["seeOradd"] = 0
	onrecord, recerr := models.GetOnlinecourserecordById(classid)
	if onrecord != nil && recerr == nil {
		userinfo, usererr := models.GetUserinformationTeacher(onrecord.UserIdPassive)
		if usererr == nil {
			c.Data["AvatarPath"] = userinfo.AvatarPath
			c.Data["UserName"] = userinfo.UserName
			c.Data["SchoolName"] = userinfo.SchoolName
			c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
			c.Data["TeacherId"] = userinfo.Id
			c.Data["classid"] = classid
		}
	}
	c.TplNames = "evaluationmainstudents.html" //跳到
}

// @Title GetOnlineCourseBooking
// @Description 老师和学生个人中心-查看一条预约信息
// @Param bookid query int true 预约课程信息主键id
// @router /GetOnlineCourseBooking/:bookid [get]
func (c *MainController) GetOnlineCourseBooking() {
	c.Data["Website"] = models.OnlineUrl
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)

	onbook, recerr := models.GetOnlinecoursebookingById(bookid)
	if onbook != nil && recerr == nil {
		//获取当前用户身份id
		identityid, _ := strconv.Atoi(c.Ctx.GetCookie("identityid"))
		if identityid != 0 && identityid == 1 { //老师查看
			userinfo, usererr := models.GetUserinformationTeacher(onbook.UserIdActive)
			if usererr == nil {
				c.Data["AvatarPath"] = userinfo.AvatarPath
				c.Data["UserName"] = userinfo.UserName
				c.Data["SchoolName"] = userinfo.SchoolName
				c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
				c.Data["TeacherId"] = userinfo.Id
			}
			c.Data["StartTime"] = onbook.StartTime
			c.Data["EndTime"] = onbook.EndTime
			c.Data["onbookid"] = onbook.Id
			c.Data["AppointMessage"] = onbook.AppointMessage
		} else if identityid >= 2 && identityid <= 3 { //学生查看
			userinfo, usererr := models.GetUserinformationTeacher(onbook.UserIdPassive)
			if usererr == nil {
				c.Data["AvatarPath"] = userinfo.AvatarPath
				c.Data["UserName"] = userinfo.UserName
				c.Data["SchoolName"] = userinfo.SchoolName
				c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
				c.Data["TeacherId"] = userinfo.Id
			}
			c.Data["StartTime"] = onbook.StartTime
			c.Data["EndTime"] = onbook.EndTime
			c.Data["AppointMessage"] = onbook.AppointMessage
			c.Data["onbookid"] = onbook.Id
		}
	}
	c.TplNames = "classmainstudents.html" //跳到
}

// @Title GetOnlineCourseBookingByTeacher
// @Description GetOnlineCourseBookingByTeacher the Userinformation
// @Param			"The id you want to GetOnlineCourseBookingByTeacher"
// @Success 200 {object} models.Userinformation
// @Failure 403
// @router /GetOnlineCourseBookingByTeacher/:bookid [get]
func (c *MainController) GetOnlineCourseBookingByTeacher() {
	c.Data["Website"] = models.OnlineUrl
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)

	onbook, recerr := models.GetOnlinecoursebookingById(bookid)
	fmt.Println(onbook)
	if onbook != nil && recerr == nil {
		userinfo, usererr := models.GetUserinformationTeacher(onbook.UserIdActive)
		if usererr == nil {
			c.Data["AvatarPath"] = userinfo.AvatarPath
			c.Data["UserName"] = userinfo.UserName
			c.Data["SchoolName"] = userinfo.SchoolName
			c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
			c.Data["TeacherId"] = userinfo.Id
		}
		c.Data["StartTime"] = onbook.StartTime
		c.Data["EndTime"] = onbook.EndTime
		c.Data["onbookid"] = onbook.Id
		c.Data["AppointMessage"] = onbook.AppointMessage
	}
	c.TplNames = "classmainstudents.html" //跳到
}

// @Title GetUserMessageList
// @Description 学生个人中心-查看一条留下下的所有留言与回复
// @Param msgid query int true 留言信息主键id
// @router /GetUserMessageList/:msgid [get]
func (c *MainController) GetUserMessageList() {
	c.Data["Website"] = models.OnlineUrl
	msgidStr := c.Ctx.Input.Params[":msgid"]
	msgid, _ := strconv.Atoi(msgidStr)

	messagelist, mserr := models.GetUsermessageByMessageId(msgid, msgid)
	if messagelist != nil && mserr == nil {
		userinfo, useerr := models.GetUserinformationTeacher(messagelist[0].PassiveUserId) ///根据被留言的老师主键查询一条老师信息
		if useerr == nil {
			c.Data["AvatarPath"] = userinfo.AvatarPath
			c.Data["UserName"] = userinfo.UserName
			c.Data["SchoolName"] = userinfo.SchoolName
			c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
			c.Data["TeacherId"] = userinfo.Id
		}
		var htmltext string = ``
		var studId int = messagelist[0].ActiveUserId //第一条留言学生id
		for i := 0; i < len(messagelist); i++ {
			if messagelist[i].ActiveUserId == studId {
				htmltext += `<div class="row messtext"><div class="col-sm-1"></div>
							<div class="col-sm-8" style="color:red;">
							留言（` + messagelist[i].ActiveName + `）: ` + messagelist[i].Contents + `
							</div></div>`
			} else {
				htmltext += `<div class="row messtext"><div class="col-sm-1"></div>
							<div class="col-sm-8" style="color:green;">
							回复（` + messagelist[i].ActiveName + `）: ` + messagelist[i].Contents + `
							</div></div>`
			}
		}
		c.Data["messid"] = messagelist[0].Id               //第一条留言的主键
		c.Data["studentid"] = messagelist[0].ActiveUserId  //第一条留言学生的主键
		c.Data["teacherid"] = messagelist[0].PassiveUserId //第一条留言老师的主键
		c.Data["stuTeaContents"] = htmltext
	}
	c.TplNames = "messagemain.html" //跳到
}

// @Title GetUserMessageListTeacher
// @Description 老师个人中心-查看一条留下下的所有留言与回复
// @Param msgid path int true 留言信息主键id
// @router /GetUserMessageListTeacher/:msgid [get]
func (c *MainController) GetUserMessageListTeacher() {
	c.Data["Website"] = models.OnlineUrl
	msgidStr := c.Ctx.Input.Params[":msgid"]
	msgid, _ := strconv.Atoi(msgidStr)

	messagelist, mserr := models.GetUsermessageByMessageId(msgid, msgid)
	if messagelist != nil && mserr == nil {
		userinfo, useerr := models.GetUserinformationTeacher(messagelist[0].ActiveUserId) ///根据被留言的老师主键查询一条老师信息
		if useerr == nil {
			c.Data["AvatarPath"] = userinfo.AvatarPath
			c.Data["UserName"] = userinfo.UserName
			c.Data["SchoolName"] = userinfo.SchoolName
			c.Data["Id"] = "00000" + strconv.Itoa(userinfo.Id)
			c.Data["TeacherId"] = userinfo.Id
		}
		var htmltext string = ``
		var studId int = messagelist[0].ActiveUserId //第一条留言学生id
		for i := 0; i < len(messagelist); i++ {
			if messagelist[i].ActiveUserId == studId {
				htmltext += `<div class="row messtext"><div class="col-sm-1"></div>
							<div class="col-sm-8" style="color:red;">
							留言（` + messagelist[i].ActiveName + `）: ` + messagelist[i].Contents + `
							</div></div>`
			} else {
				htmltext += `<div class="row messtext"><div class="col-sm-1"></div>
							<div class="col-sm-8" style="color:green;">
							回复（` + messagelist[i].ActiveName + `）: ` + messagelist[i].Contents + `
							</div></div>`
			}
		}
		c.Data["messid"] = messagelist[0].Id               //第一条留言的主键
		c.Data["studentid"] = messagelist[0].ActiveUserId  //第一条留言学生的主键
		c.Data["teacherid"] = messagelist[0].PassiveUserId //第一条留言老师的主键
		c.Data["stuTeaContents"] = htmltext
	}
	c.TplNames = "messagemain.html" //跳到
}

// @Title UpdateStudent
// @Description 学生个人中心-编辑个人信息
// @router /UpdateStudent/ [get]
func (c *MainController) UpdateStudent() {
	fmt.Println("woshiget")
	c.Data["Website"] = models.OnlineUrl
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	fmt.Println(stuuserid)
	userinfo, usererr := models.GetUserinformationStudent(stuuserid)
	fmt.Println(userinfo.UserName)
	if usererr == nil {
		c.Data["AvatarPath"] = userinfo.AvatarPath
		c.Data["UserName"] = userinfo.UserName
		c.Data["UserSex"] = userinfo.UserSex
		c.Data["SchoolName"] = userinfo.SchoolName
		c.Data["AgeName"] = userinfo.AgeName
		if userinfo.LevelYear > 0 {
			c.Data["LevelYear"] = userinfo.LevelYear
		} else {
			c.Data["LevelYear"] = ""
		}
		c.Data["Mailbox"] = userinfo.Mailbox
		c.Data["ParentMailbox"] = userinfo.ParentMailbox
		c.Data["IphoneNum"] = userinfo.IphoneNum
		c.Data["SchoolAgeId"] = userinfo.SchoolAgeId
		c.Data["StudyDifficult"] = userinfo.StudyDifficult
		c.Data["SchoolId"] = userinfo.SchoolId
		c.Data["Mailbox"] = userinfo.Mailbox
	}
	userclass, clerr := models.GetRemedialcoursesMain(stuuserid, 0)
	var userlistclass string = ""
	var userclassstr string = "" //名称集合
	if userclass != nil && clerr == nil {
		for i := 0; i < len(userclass); i++ {
			userlistclass += strconv.Itoa(userclass[i].CoursesId) + ","
			userclassstr += userclass[i].CourseName + "  "
		}
	}
	c.Data["userlistclass"] = userlistclass
	c.Data["userclassstr"] = userclassstr

	c.TplNames = "personal.html" //跳到
}

// @Title UpdateStudent
// @Description 学生个人中心-编辑个人信息
// @Param userinformation form models.Userinformation true 用户信息
// @router /UpdateStudent/ [post]
func (c *MainController) UpdateStudent2() {
	c.Data["Website"] = models.OnlineUrl
	if c.Ctx.Input.Request.Method == "GET" {
		c.TplNames = "personal.html" //跳到
	}
	request := c.Ctx.Request
	_, imgstr := models.GetImganddata2(request, Headurl)
	fmt.Println("图片路径为：")
	//fmt.Println(jsons)
	//	fmt.Println(imgstr)
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	var name []string = c.Ctx.Input.Request.Form["txtUserName"]        //
	var sex []string = c.Ctx.Input.Request.Form["selUserSex"]          //
	var age []string = c.Ctx.Input.Request.Form["selage"]              //
	var school []string = c.Ctx.Input.Request.Form["txtnowschoolname"] //
	var year []string = c.Ctx.Input.Request.Form["txtruxueyear"]       //
	var mail []string = c.Ctx.Input.Request.Form["txtmail"]            //
	var pmail []string = c.Ctx.Input.Request.Form["txtuseremal"]       //
	var class []string = c.Ctx.Input.Request.Form["stucheckclass"]     //
	fmt.Println(class)
	fmt.Println("ddd")
	var dif []string = c.Ctx.Input.Request.Form["txtnandian"] //
	userinfo, usererr := models.GetUserinformationById(stuuserid)
	fmt.Println("ddd2")
	if usererr == nil && userinfo.Id > 0 {
		userinfo.UserName = name[0]
		userinfo.UserSex = sex[0]
		fmt.Println("ddd3")
		if age != nil {
			userinfo.SchoolAgeId, _ = strconv.Atoi(age[0])
		}
		userinfo.SchoolName = school[0]
		fmt.Println("ddd4")
		userinfo.LevelYear, _ = strconv.Atoi(year[0])
		userinfo.Mailbox = mail[0]
		userinfo.ParentMailbox = pmail[0]
		fmt.Println("ddd5")
		userinfo.StudyDifficult = dif[0]
		fmt.Println("ddd6")
		fmt.Println(imgstr)
		if imgstr != "" {
			userinfo.AvatarPath = imgstr
			c.Ctx.SetCookie("username", name[0])
			c.Ctx.SetCookie("AvatarPath", imgstr)
		}
		fmt.Println("ddd3")
		upresulterr := models.UpdateUserinformationById(userinfo)
		err := SetUserClassList(userinfo.Id, class)
		if upresulterr == nil && err == nil {
			c.Data["json"] = map[string]interface{}{"state": 1} //修改成功
		} else {
			c.Data["json"] = map[string]interface{}{"state": 0} //修改失败
		}
	}
	c.Redirect("/orange/Main/UpdateStudent/", 302)
}

// @Title UpdateTeacher
// @Description 跳页到老师个人中心-编辑个人信息
// @router /UpdateTeacher/ [get]
func (c *MainController) UpdateTeacher() {
	c.Data["Website"] = models.OnlineUrl
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	fmt.Println(stuuserid)
	userinfo, usererr := models.GetUserinformationTeacher(stuuserid)
	fmt.Println(userinfo.UserName)
	if usererr == nil {
		c.Data["AvatarPath"] = userinfo.AvatarPath
		c.Data["UserName"] = userinfo.UserName
		c.Data["UserSex"] = userinfo.UserSex
		c.Data["SchoolName"] = userinfo.SchoolName
		c.Data["SchoolId"] = userinfo.SchoolId
		c.Data["HighSchoolName"] = userinfo.HighSchool
		c.Data["SeniorLocation"] = userinfo.SeniorLocation //高中学校市区id
		c.Data["Professional"] = userinfo.Professional     //专业
		c.Data["DegreeName"] = userinfo.DegreeName         //学位
		c.Data["UserDegree"] = userinfo.UserDegree
		c.Data["GradeName"] = userinfo.GradeName       //所教年级
		c.Data["GradeId"] = userinfo.GradeId           //所教年级
		c.Data["CourseName"] = userinfo.CourseName     //主辅导课
		c.Data["CourseNameId"] = userinfo.CourseNameId //主辅导课程id
		if userinfo.LevelYear > 0 {
			c.Data["LevelYear"] = userinfo.LevelYear
		} else {
			c.Data["LevelYear"] = ""
		}
		c.Data["Mailbox"] = userinfo.Mailbox
		c.Data["IphoneNum"] = userinfo.IphoneNum
		c.Data["BriefIntroduction"] = userinfo.BriefIntroduction
		c.Data["UserHobby"] = userinfo.UserHobby
		c.Data["Mailbox"] = userinfo.Mailbox
	}
	userclass, clerr := models.GetRemedialcoursesMain(stuuserid, 0)
	var userlistclass string = "" //主键集合
	var userclassstr string = ""  //名称集合
	if userclass != nil && clerr == nil {
		for i := 0; i < len(userclass); i++ {
			userlistclass += strconv.Itoa(userclass[i].CoursesId) + ","
			userclassstr += userclass[i].CourseName + "  "
		}
	}
	c.Data["userlistclass"] = userlistclass
	c.Data["userclassstr"] = userclassstr

	//补习学龄段
	var strage string = ""
	if userinfo.SchoolAgeIdT != "" {
		var ageliststr string = userinfo.SchoolAgeIdT
		ageidlist := strings.Split(ageliststr, ",")
		for i := 0; i < len(ageidlist); i++ {
			ageid, _ := strconv.Atoi(ageidlist[i])
			schoolagemodel, _ := models.GetSchoolagesById(ageid)
			strage = strage + schoolagemodel.AgeName + " "
		}
	}
	c.Data["AgeNames"] = strage
	c.Data["schoolagelist"] = userinfo.SchoolAgeIdT
	c.TplNames = "personalteacher.html" //跳到
}

// @Title UpdateTeacher
// @Description 跳页到老师个人中心-编辑个人信息
// @router /UpdateTeacher/ [post]
func (c *MainController) UpdateTeacher2() {
	c.Data["Website"] = models.OnlineUrl
	if c.Ctx.Input.Request.Method == "GET" {
		c.TplNames = "personalteacher.html" //跳到
	}
	request := c.Ctx.Request
	jsons, imgstr := models.GetImganddata2(request, Headurl)
	fmt.Println("图片路径为：")
	fmt.Println(jsons)
	fmt.Println(imgstr)
	userid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	var name []string = c.Ctx.Input.Request.Form["txtUserName"] //
	var sex []string = c.Ctx.Input.Request.Form["selUserSex"]   //
	var cityid []string = c.Ctx.Input.Request.Form["selCitys"]  //高中所在市区主键id
	var highschoolname []string = c.Ctx.Input.Request.Form["txthighschool"]
	var nowschoolname []string = c.Ctx.Input.Request.Form["txtnowschoolname"]
	var year []string = c.Ctx.Input.Request.Form["txtruxueyear"]
	var mail []string = c.Ctx.Input.Request.Form["txtusermail"]
	var selDegree []string = c.Ctx.Input.Request.Form["selDegree"]
	var txtProfessional []string = c.Ctx.Input.Request.Form["txtProfessional"] //
	var age []string = c.Ctx.Input.Request.Form["ageTeacher"]                  //补习学龄段
	var selcourse []string = c.Ctx.Input.Request.Form["selcourse"]             //主辅导科目
	var selckclass []string = c.Ctx.Input.Request.Form["selckclass"]           //辅辅导课程
	var txtjianjie []string = c.Ctx.Input.Request.Form["txtjianjie"]
	var txtaihao []string = c.Ctx.Input.Request.Form["txtaihao"]
	var xuelingduan string = ""
	for i := 0; i < len(age); i++ {
		if i == (len(age) - 1) {
			xuelingduan = xuelingduan + age[i]
		} else {
			xuelingduan = xuelingduan + age[i] + ","
		}
	}

	userinfo, usererr := models.GetUserinformationById(userid)
	if usererr == nil && userinfo.Id > 0 {
		userinfo.UserName = name[0]
		userinfo.UserSex = sex[0]
		userinfo.SeniorLocation, _ = strconv.Atoi(cityid[0])
		userinfo.HighSchool = highschoolname[0]
		userinfo.SchoolName = nowschoolname[0]
		userinfo.LevelYear, _ = strconv.Atoi(year[0])
		userinfo.Mailbox = mail[0]
		userinfo.UserDegree, _ = strconv.Atoi(selDegree[0])
		userinfo.Professional = txtProfessional[0]
		userinfo.BriefIntroduction = txtjianjie[0]
		userinfo.UserHobby = txtaihao[0]
		userinfo.SchoolAgeIdT = xuelingduan
		if imgstr != "" {
			userinfo.AvatarPath = imgstr
			c.Ctx.SetCookie("username", name[0])
			c.Ctx.SetCookie("AvatarPath", imgstr)
		}
		upresulterr := models.UpdateUserinformationById(userinfo)
		ismainid, _ := strconv.Atoi(selcourse[0])
		err := SetUserClassListTeacher(userinfo.Id, selckclass, ismainid)
		if upresulterr == nil && err == nil {
			c.Data["json"] = map[string]interface{}{"state": 1} //修改成功
		} else {
			c.Data["json"] = map[string]interface{}{"state": 0} //修改失败
		}
	}
	c.Redirect("/orange/Main/UpdateTeacher/", 302)
}

// @Title RetrievePassword
// @Description 跳页到登录-找回密码
// @router /RetrievePassword/ [get]
func (c *MainController) RetrievePassword() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "password.html" //跳到
}

// @Title AboutMe
// @Description 跳页到关于我们
// @router /AboutMe/:tapid [get]
func (c *MainController) AboutMe() {
	c.Data["Website"] = models.OnlineUrl
	tapid := c.Ctx.Input.Params[":tapid"]
	c.Data["NowTapid"] = tapid
	c.TplNames = "aboutme.html" //跳到关于我们
}

// @Title TechnologicalProcess
// @Description 跳页到老师注册进入页面
// @router /TechnoRegister/ [get]
func (c *MainController) TechnoRegister() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "teacherregister.html" //
}

// @Title RegisteredLinefwtk
// @Description 跳页到学生注册服务条款页面
// @router /RegisteredLinefwtk/ [get]
func (c *MainController) RegisteredLinefwtk() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "linefwtk.html" //
}

// @Title PayMentUser
// @Description 跳到支付页面
// @Param  money query float true 要充值的钱
// @router /PayMentUser/:money [get]
func (c *MainController) PayMentUser() {
	moneyStr := c.Ctx.Input.Params[":money"]
	money, _ := strconv.ParseFloat(moneyStr, 10)
	//添加一条 账户充值记录
	var addaccountpay models.Amountrecords
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	fmt.Println(stuuserid)
	nowaccuser, _ := models.GetAccountfundsByuid(stuuserid) //根据用户主键id查询用户账户信息
	fmt.Println(nowaccuser)
	addaccountpay.UserId = stuuserid
	addaccountpay.RecordMoney = money
	addaccountpay.Balance = nowaccuser.Balance + money //充值后余额显示
	addaccountpay.RecordType = 0                       //充值
	addaccountpay.RecordTime = time.Now()              //操作日期
	addaccountpay.TradingWayId = 2                     //支付宝支付
	addaccountpay.IsComplete = 0                       //0：未完成，1：已完成
	addint, adderr := models.AddAmountrecords(&addaccountpay)
	fmt.Println(addint)
	if adderr == nil {

	}
	r := models.Request{
		NotifyUrl:   `http://www.fankunedu.com/orange/Main/PayEndNotify/`, // 付款后异步通知页面
		ReturnUrl:   `http://www.fankunedu.com/orange/Main/PayEnd/`,       // 付款后返回页面
		OutTradeNo:  strconv.FormatInt(int64(addint), 10),                 // 充值订单号
		SellerEmail: `1113911608@qq.com`,                                  // 支付宝用户名
		Service:     `create_direct_pay_by_user`,                          // 不可改
		PaymentType: `1`,                                                  // 不可改
		Subject:     `账户充值`,                                               // 商品名称
		TotalFee:    money,                                                // 价格
	}

	cc := models.Config{
		Partner: `2088911257813375`,                 // 支付宝合作者身份 ID
		Key:     `scnf70tnzygvjkdp259w2z2h2e0mhrrc`, // 支付宝交易安全校验码
	}
	fromstr := models.NewPage(cc, r, os.Stdout)
	fmt.Println(fromstr)
	c.Data["subContent"] = fromstr
	c.TplNames = "payment.html"
}

type Result struct {
	// 状态
	Status int
	// 本网站订单号
	OrderNo string
	// 支付宝交易号
	TradeNo string
	// 买家支付宝账号
	BuyerEmail string
	// 错误提示
	Message string
}

// @Title PayEnd
// @Description 支付同步回调地址
// @router /PayEnd/ [get]
func (c *MainController) PayEnd() {
	c.Data["Website"] = models.OnlineUrl
	result := Return(&c.Controller)
	if result.Status == 1 {
		fmt.Println("成功")
		//获取唯一标识
		inid := result.OrderNo //本网站订单号
		inidt, _ := strconv.Atoi(inid)
		if inidt > 0 {
			accountpay, _ := models.GetAmountrecordsById(inidt)
			zhifumoney := strconv.FormatFloat(accountpay.RecordMoney, 'f', -1, 64)
			if accountpay.IsComplete == 0 { //订单是否已处理，否处理未处理的订单
				accountpay.IsComplete = 1
				upresulterr := models.UpdateAmountrecordsById(accountpay) //更新提现记录
				fmt.Println(upresulterr)
				//用户账户添加金额
				accountuser, _ := models.GetAccountfundsByuid(accountpay.UserId)
				accountuser.Balance = accountuser.Balance + accountpay.RecordMoney
				upaccerr := models.UpdateAccountfundsById(&accountuser) //更新账户余额
				if upaccerr == nil {
					c.Data["resultStr"] = "支付成功：" + zhifumoney + "元。"
				}
			} else {
				c.Data["resultStr"] = "支付成功：" + zhifumoney + "元。"
			}
		}

	} else {
		c.Data["resultStr"] = result.Message
	}

	fmt.Println(result)
	c.TplNames = "moneyover.html" //
}

// @Title PayEndNotify
// @Description 支付异步回调地址
// @router /PayEndNotify/ [post]
func (c *MainController) PayEndNotify() {
	result := Notify(&c.Controller)
	if result.Status == 1 {
		//获取唯一标识
		inid := result.OrderNo //本网站订单号
		inidt, _ := strconv.Atoi(inid)
		accountpay, _ := models.GetAmountrecordsById(inidt)
		//var caozuomoney = strconv.FormatInt(int64(accountpay.RecordMoney), 10)
		zhifumoney := strconv.FormatFloat(accountpay.RecordMoney, 'f', -1, 64)
		if accountpay.IsComplete == 0 { //订单是否已处理，否处理未处理的订单
			accountpay.IsComplete = 1
			upresulterr := models.UpdateAmountrecordsById(accountpay) //更新提现记录
			fmt.Println(upresulterr)
			//用户账户添加金额
			accountuser, _ := models.GetAccountfundsByuid(accountpay.UserId)
			accountuser.Balance = accountuser.Balance + accountpay.RecordMoney
			upaccerr := models.UpdateAccountfundsById(&accountuser) //更新账户余额
			if upaccerr == nil {
				c.Data["resultStr"] = "支付成功：" + zhifumoney + "元。"
			}
		} else {
			c.Data["resultStr"] = "支付成功：" + zhifumoney + "元。"
		}
	} else {
		c.Data["resultStr"] = result.Message
	}

	fmt.Println(result)
	c.TplNames = "moneyover.html" //
}

/* 被动接收支付宝同步跳转的页面 */
func Return(contro *beego.Controller) *Result {
	// 列举全部传参
	type Params struct {
		//Body        string `form:"body" json:"body"`                 // 描述
		BuyerEmail  string `form:"buyer_email" json:"buyer_email"`   // 买家账号
		BuyerId     string `form:"buyer_id" json:"buyer_id"`         // 买家ID
		Exterface   string `form:"exterface" json:"exterface"`       // 接口名称
		IsSuccess   string `form:"is_success" json:"is_success"`     // 交易是否成功
		NotifyId    string `form:"notify_id" json:"notify_id"`       // 通知校验id
		NotifyTime  string `form:"notify_time" json:"notify_time"`   // 校验时间
		NotifyType  string `form:"notify_type" json:"notify_type"`   // 校验类型
		OutTradeNo  string `form:"out_trade_no" json:"out_trade_no"` // 在网站中唯一id
		PaymentType uint8  `form:"payment_type" json:"payment_type"` // 支付类型
		SellerEmail string `form:"seller_email" json:"seller_email"` // 卖家账号
		SellerId    string `form:"seller_id" json:"seller_id"`       // 卖家id
		Subject     string `form:"subject" json:"subject"`           // 商品名称
		TotalFee    string `form:"total_fee" json:"total_fee"`       // 总价
		TradeNo     string `form:"trade_no" json:"trade_no"`         // 支付宝交易号
		TradeStatus string `form:"trade_status" json:"trade_status"` // 交易状态 TRADE_FINISHED或TRADE_SUCCESS表示交易成功
		Sign        string `form:"sign" json:"sign"`                 // 签名
		SignType    string `form:"sign_type" json:"sign_type"`       // 签名类型
	}

	// 实例化参数
	param := &Params{}

	// 结果
	result := &Result{}

	// 解析表单内容，失败返回错误代码-3
	if err := contro.ParseForm(param); err != nil {
		result.Status = -3
		result.Message = "解析表单失败"
		return result
	}
	// 如果最基本的网站交易号为空，返回错误代码-1
	if param.OutTradeNo == "" { //不存在交易号
		result.Status = -1
		result.Message = "站交易号为空"
		return result
	} else {
		// 生成签名
		//signs := Sign(param)
		signs := models.Sign(param)
		log.Println(signs)
		log.Println(param.Sign)
		//log.Println(param.Body)
		// 对比签名是否相同
		if signs == param.Sign { //只有相同才说明该订单成功了
			// 判断订单是否已完成
			if param.TradeStatus == "TRADE_FINISHED" || param.TradeStatus == "TRADE_SUCCESS" { //交易成功
				result.Status = 1
				result.OrderNo = param.OutTradeNo
				result.TradeNo = param.TradeNo
				result.BuyerEmail = param.BuyerEmail
				return result
			} else { // 交易未完成，返回错误代码-4
				result.Status = -4
				result.Message = "交易未完成"
				return result
			}
		} else { // 签名认证失败，返回错误代码-2
			result.Status = -2
			result.Message = "签名认证失败"
			return result
		}
	}

	// 位置错误类型-5
	result.Status = -5
	result.Message = "位置错误"
	return result
}

/* 被动接收支付宝异步通知 */
func Notify(contro *beego.Controller) *Result {
	// 从body里读取参数，用&切割
	//postArray := strings.Split(string(contro.Ctx.Input.CopyBody()), "&")
	var postArray []string
	//var jsonS string
	for k, v := range contro.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		fmt.Println(k)
		fmt.Println(v[0])
		postArray = append(postArray, k+"="+v[0])

	}
	// 实例化url
	urls := &url.Values{}

	// 保存传参的sign
	var paramSign string
	var sign string

	// 如果字符串中包含sec_id说明是手机端的异步通知
	if strings.Index(string(contro.Ctx.Input.CopyBody()), `alipay.wap.trade.create.direct`) == -1 { // 快捷支付
		//for _, v := range postArray {
		for i := 0; i < len(postArray); i++ {
			detail := strings.Split(postArray[i], "=")

			// 使用=切割字符串 去除sign和sign_type
			if detail[0] == "sign" || detail[0] == "sign_type" {
				if detail[0] == "sign" {
					paramSign = detail[1]
				}
				continue
			} else {
				urls.Add(detail[0], detail[1])
			}
		}

		// url解码
		urlDecode, _ := url.QueryUnescape(urls.Encode())
		sign, _ = url.QueryUnescape(urlDecode)
	} else { // 手机网页支付
		// 手机字符串加密顺序
		mobileOrder := []string{"service", "v", "sec_id", "notify_data"}
		for _, v := range mobileOrder {
			//for _, value := range postArray {
			for i := 0; i < len(postArray); i++ {
				fmt.Println(postArray[i])
				detail := strings.Split(postArray[i], "=")
				// 保存sign
				if detail[0] == "sign" {
					paramSign = detail[1]
				} else {
					// 如果满足当前v
					if detail[0] == v {
						if sign == "" {
							sign = detail[0] + "=" + detail[1]
						} else {
							sign += "&" + detail[0] + "=" + detail[1]
						}
					}
				}
			}
		}
		sign, _ = url.QueryUnescape(sign)

		// 获取<trade_status></trade_status>之间的request_token
		re, _ := regexp.Compile("\\<trade_status[\\S\\s]+?\\</trade_status>")
		rt := re.FindAllString(sign, 1)
		trade_status := strings.Replace(rt[0], "<trade_status>", "", -1)
		trade_status = strings.Replace(trade_status, "</trade_status>", "", -1)
		urls.Add("trade_status", trade_status)

		// 获取<out_trade_no></out_trade_no>之间的request_token
		re, _ = regexp.Compile("\\<out_trade_no[\\S\\s]+?\\</out_trade_no>")
		rt = re.FindAllString(sign, 1)
		out_trade_no := strings.Replace(rt[0], "<out_trade_no>", "", -1)
		out_trade_no = strings.Replace(out_trade_no, "</out_trade_no>", "", -1)
		urls.Add("out_trade_no", out_trade_no)

		// 获取<buyer_email></buyer_email>之间的request_token
		re, _ = regexp.Compile("\\<buyer_email[\\S\\s]+?\\</buyer_email>")
		rt = re.FindAllString(sign, 1)
		buyer_email := strings.Replace(rt[0], "<buyer_email>", "", -1)
		buyer_email = strings.Replace(buyer_email, "</buyer_email>", "", -1)
		urls.Add("buyer_email", buyer_email)

		// 获取<trade_no></trade_no>之间的request_token
		re, _ = regexp.Compile("\\<trade_no[\\S\\s]+?\\</trade_no>")
		rt = re.FindAllString(sign, 1)
		trade_no := strings.Replace(rt[0], "<trade_no>", "", -1)
		trade_no = strings.Replace(trade_no, "</trade_no>", "", -1)
		urls.Add("trade_no", trade_no)
	}
	// 追加密钥
	sign += "scnf70tnzygvjkdp259w2z2h2e0mhrrc" //AlipayKey合作者私钥

	// 返回参数
	result := &Result{}

	// md5加密
	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))
	fmt.Println("输出签名：")
	fmt.Println(paramSign)
	fmt.Println(sign)
	if paramSign == sign { // 传进的签名等于计算出的签名，说明请求合法
		// 判断订单是否已完成
		if urls.Get("trade_status") == "TRADE_FINISHED" || urls.Get("trade_status") == "TRADE_SUCCESS" { //交易成功
			contro.Ctx.WriteString("success")
			result.Status = 1
			result.OrderNo = urls.Get("out_trade_no")
			result.TradeNo = urls.Get("trade_no")
			result.BuyerEmail = urls.Get("buyer_email")
			return result
		} else {
			contro.Ctx.WriteString("error")
		}
	} else {
		contro.Ctx.WriteString("error")
		// 签名不符，错误代码-1
		result.Status = -1
		result.Message = "签名不符"
		return result
	}
	// 未知错误-2
	result.Status = -2
	result.Message = "未知错误"
	return result
}
