package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"orange/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

//首页展示
func (c *MainController) Get() {
	c.Data["Website"] = OnlineUrl
	c.Ctx.SetCookie("OnlineUrl", OnlineUrl)
	c.TplNames = "index.tpl" //首页
}

// 登录方法
// @Title Logins
// @Description Logins the TbUser
// @Param			"The id you want to Logins"
// @Success 200 {object} models.TbUser
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
		c.Data["Website"] = OnlineUrl
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
			c.Data["Website"] = OnlineUrl
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

// 登录方法
// @Title LoginUser
// @Description LoginUser the TbUser
// @Param			"The id you want to LoginUser"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /LoginUser/ [post]
func (c *MainController) LoginUser() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println("接到信息为：")
	fmt.Println(jsonS)
	var v models.Userinformation
	json.Unmarshal([]byte(jsonS), &v)
	//fmt.Println(v)
	var vuser *models.Userinformation
	fmt.Println("用户名，密码：")
	fmt.Println(v.UserName, v.LoginPassword)
	vuser, err := models.GetUserinformationLogin(v.UserName, v.LoginPassword)
	if err == nil && vuser != nil {
		fmt.Println(vuser)
		c.Data["Website"] = OnlineUrl
		c.Ctx.SetCookie("username", vuser.UserName)
		c.Ctx.SetCookie("userid", strconv.Itoa(vuser.Id))
		c.Ctx.SetCookie("identityid", strconv.Itoa(vuser.IdentityId))
		c.Ctx.SetCookie("AvatarPath", vuser.AvatarPath)
		fmt.Println(vuser.AvatarPath)
		//c.TplNames = "index.tpl"
		c.Data["json"] = "OK"
	} else {
		vphoneuser, errph := models.GetUserinformationLoginPhone(v.UserName, v.LoginPassword)
		if errph == nil && vphoneuser != nil {
			fmt.Println(vphoneuser)
			c.Data["Website"] = OnlineUrl
			c.Ctx.SetCookie("username", vphoneuser.UserName)
			c.Ctx.SetCookie("userid", strconv.Itoa(vphoneuser.Id))
			c.Ctx.SetCookie("identityid", strconv.Itoa(vphoneuser.IdentityId))
			c.Ctx.SetCookie("AvatarPath", vphoneuser.AvatarPath)
			fmt.Println(vphoneuser.AvatarPath)
			//c.TplNames = "index.tpl"
			c.Data["json"] = "OK"
		} else {
			fmt.Println(err)
			c.Data["blockdiv"] = "none"
			//c.TplNames = "404.html"
			c.Data["json"] = "NO"
		}
	}
	c.ServeJson()
}

// 退出方法
// @Title OutLogins
// @Description OutLogins the TbUser
// @Param			"The id you want to OutLogins"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /OutLogins/ [get]
func (c *MainController) OutLogins() {
	c.Data["Website"] = OnlineUrl
	c.Ctx.SetCookie("username", "")
	c.Ctx.SetCookie("userid", "")
	c.Ctx.SetCookie("identityid", "")
	c.Ctx.SetCookie("AvatarPath", "")
	c.Data["json"] = "OK"
	c.ServeJson()
	//c.TplNames = "index.tpl"
}

// 注册方法
// @Title Registered
// @Description Registered the TbUser
// @Param			"The id you want to Registered"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /Registered/ [get]
func (c *MainController) Registered() {
	c.Data["Website"] = OnlineUrl
	c.TplNames = "register.html" //跳到注册页面
}

// 问答中心模块
// @Title QuestionsCenter
// @Description QuestionsCenter the TbUser
// @Param			"The id you want to QuestionsCenter"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /QuestionsCenter/ [get]
func (c *MainController) QuestionsCenter() {
	c.Data["Website"] = OnlineUrl
	wendalist, _ := models.GetQuestionaskByJingCaiCount()
	c.Data["wendacount"] = wendalist
	c.TplNames = "problem_list.html" //跳到问答中心
}

// 老师个人中心
// @Title UserTeacher
// @Description UserTeacher the TbUser
// @Param			"The id you want to UserTeacher"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /UserTeacher/:tapid [get]
func (c *MainController) UserTeacher() {
	c.Data["Website"] = OnlineUrl
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
		c.Data["StuSex"] = showTeacher.UserSex
		c.Data["SchoolName"] = showTeacher.SchoolName
		c.Data["IdentityName"] = showTeacher.IdentityName
		c.Data["AllDate"] = strconv.Itoa(showTeacher.AllDate)
		c.Data["AllCount"] = strconv.Itoa(showTeacher.AllCount)
		c.Data["AllPerson"] = strconv.Itoa(showTeacher.AllPerson)
		c.Data["Professional"] = showTeacher.Professional
		c.Data["CourseName"] = showTeacher.CourseName
		c.Data["UserHobby"] = showTeacher.UserHobby
		c.Data["BriefIntroduction"] = showTeacher.BriefIntroduction //学习难点
	}
	zhucourse, fuerr := models.GetRemedialcoursesMain(stuuserid, 0)
	fmt.Println(zhucourse)
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
	tixianlist, _ := models.GetAmountrecordsByUseridCount(1, stuuserid)
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

// 学生个人中心
// @Title UserStudent
// @Description UserStudent the TbUser
// @Param			"The id you want to UserStudent"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /UserStudent/:tapid [get]
func (c *MainController) UserStudent() {
	c.Data["Website"] = OnlineUrl
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
		c.Data["AllDate"] = strconv.Itoa(showStudent.AllDate)
		c.Data["AllCount"] = strconv.Itoa(showStudent.AllCount)
		c.Data["AllPerson"] = strconv.Itoa(showStudent.AllPerson)
		c.Data["IdentityName"] = showStudent.IdentityName
		c.Data["StuSex"] = showStudent.UserSex
		c.Data["SchoolName"] = showStudent.SchoolName
		c.Data["AgeName"] = showStudent.AgeName
		c.Data["LevelYear"] = strconv.Itoa(showStudent.LevelYear)
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
	fmt.Println(zijin.Balance)

	dongjiezijin, _ := models.GetFrozenFundsByUserid(stuuserid)
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
	tixianlist, _ := models.GetAmountrecordsByUseridCount(1, stuuserid)
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
	//我的关注
	guanzhulist, _ := models.GetRelationsByUidCount(stuuserid, "关注")
	c.Data["guanzhucount"] = guanzhulist
	tuiyuelist, _ := models.GetOnlinecoursebookingBySidNotOnCount(stuuserid)
	c.Data["tuiyuecount"] = tuiyuelist

	c.Data["tapNum"] = tapid        //显示第几个tap
	c.TplNames = "studentmain.html" //跳到学生个人中心
}

// 学生个人中心-查看全部课程中自己给老师的一条评价
// @Title QuestionsCenter
// @Description QuestionsCenter the TbUser
// @Param			"The id you want to QuestionsCenter"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetOnLineEvaluation/:evalid [get]
func (c *MainController) GetOnLineEvaluation() {
	c.Data["Website"] = OnlineUrl
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

// 老师个人中心-查看学生给自己的一条评价
// @Title GetOnLineEvaluationTeacher
// @Description GetOnLineEvaluationTeacher the TbUser
// @Param			"The id you want to GetOnLineEvaluationTeacher"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetOnLineEvaluationTeacher/:evalid [get]
func (c *MainController) GetOnLineEvaluationTeacher() {
	c.Data["Website"] = OnlineUrl
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

// 学生个人中心-添加一条全部课程中的一条评价
// @Title QuestionsCenter
// @Description QuestionsCenter the TbUser
// @Param			"The id you want to QuestionsCenter"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /AddOnLineEvaluation/:classid [get]
func (c *MainController) AddOnLineEvaluation() {
	c.Data["Website"] = OnlineUrl
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

// 学生个人中心-查看一条预约信息
// @Title GetOnlineCourseBooking
// @Description GetOnlineCourseBooking the TbUser
// @Param			"The id you want to GetOnlineCourseBooking"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetOnlineCourseBooking/:bookid [get]
func (c *MainController) GetOnlineCourseBooking() {
	c.Data["Website"] = OnlineUrl
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)

	onbook, recerr := models.GetOnlinecoursebookingById(bookid)
	if onbook != nil && recerr == nil {
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
		c.Data["onbookid"] = onbook.Id
	}
	c.TplNames = "classmainstudents.html" //跳到
}

// 老师个人中心-查看一条预约信息
// @Title GetOnlineCourseBookingByTeacher
// @Description GetOnlineCourseBookingByTeacher the TbUser
// @Param			"The id you want to GetOnlineCourseBookingByTeacher"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetOnlineCourseBookingByTeacher/:bookid [get]
func (c *MainController) GetOnlineCourseBookingByTeacher() {
	c.Data["Website"] = OnlineUrl
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)

	onbook, recerr := models.GetOnlinecoursebookingById(bookid)
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
	}
	c.TplNames = "classmainstudents.html" //跳到
}

// 学生个人中心-查看一条留下下的所有留言与回复
// @Title GetUserMessageList
// @Description GetUserMessageList the TbUser
// @Param			"The id you want to GetUserMessageList"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetUserMessageList/:msgid [get]
func (c *MainController) GetUserMessageList() {
	c.Data["Website"] = OnlineUrl
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

// 老师个人中心-查看一条留下下的所有留言与回复
// @Title GetUserMessageListTeacher
// @Description GetUserMessageListTeacher the TbUser
// @Param			"The id you want to GetUserMessageListTeacher"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /GetUserMessageListTeacher/:msgid [get]
func (c *MainController) GetUserMessageListTeacher() {
	c.Data["Website"] = OnlineUrl
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

// 学生个人中心-编辑个人信息
// @Title UpdateStudent
// @Description UpdateStudent the TbUser
// @Param			"The id you want to UpdateStudent"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /UpdateStudent/ [get]
func (c *MainController) UpdateStudent() {
	c.Data["Website"] = OnlineUrl
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
		c.Data["LevelYear"] = userinfo.LevelYear
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

// 老师个人中心-编辑个人信息
// @Title UpdateTeacher
// @Description UpdateTeacher the TbUser
// @Param			"The id you want to UpdateTeacher"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /UpdateTeacher/ [get]
func (c *MainController) UpdateTeacher() {
	c.Data["Website"] = OnlineUrl
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
		c.Data["HighSchoolName"] = userinfo.HighSchoolName
		c.Data["SeniorLocation"] = userinfo.SeniorLocation //高中学校id
		c.Data["Professional"] = userinfo.Professional     //专业
		c.Data["DegreeName"] = userinfo.DegreeName         //学位
		c.Data["UserDegree"] = userinfo.UserDegree
		c.Data["GradeName"] = userinfo.GradeName       //所教年级
		c.Data["GradeId"] = userinfo.GradeId           //所教年级
		c.Data["CourseName"] = userinfo.CourseName     //主辅导课
		c.Data["CourseNameId"] = userinfo.CourseNameId //主辅导课程id
		c.Data["LevelYear"] = userinfo.LevelYear
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

	c.TplNames = "personalteacher.html" //跳到
}

// 登录-找回密码
// @Title RetrievePassword
// @Description RetrievePassword the TbUser
// @Param			"The id you want to RetrievePassword"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /RetrievePassword/ [get]
func (c *MainController) RetrievePassword() {
	c.Data["Website"] = OnlineUrl
	c.TplNames = "password.html" //跳到
}

// 关于我们
// @Title AboutMe
// @Description AboutMe the TbUser
// @Param			"The id you want to AboutMe"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /AboutMe/ [get]
func (c *MainController) AboutMe() {
	c.Data["Website"] = OnlineUrl
	c.TplNames = "introduction.html" //跳到关于我们
}

// 学习流程，教学流程
// @Title TechnologicalProcess
// @Description TechnologicalProcess the TbUser
// @Param			"The id you want to TechnologicalProcess"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TechnologicalProcess/ [get]
func (c *MainController) TechnologicalProcess() {
	c.Data["Website"] = OnlineUrl
	c.TplNames = "process.html" //跳到流程页面
}
