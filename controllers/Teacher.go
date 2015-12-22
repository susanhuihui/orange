package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"orange/models"
	"strconv"
	"strings"
)

type TeacherController struct {
	beego.Controller
}

//展示
func (c *TeacherController) Get() {
	c.Data["Website"] = models.OnlineUrl
	c.Data["allUserscount"] = 200

	c.TplNames = "teacher_Model.html" //
}

//展示
//func (c *TeacherController) Get() {
//	c.Data["Website"] = OnlineUrl
//	c.TplNames = "teacher.html" //
//}

// 老师
// @Title TeacherList
// @Description TeacherList the TbUser
// @Param			"The id you want to TeacherList"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherList/:seltype [get]
func (c *TeacherController) TeacherList() {
	c.Data["Website"] = models.OnlineUrl
	allsel := c.Ctx.Input.Params[":seltype"]
	sellist := strings.Split(allsel, "&") //总长度为8
	for i := 0; i < len(sellist); i++ {
		fmt.Println(sellist[i])
	}
	if sellist != nil {
		seltypestr := sellist[0] //c.Ctx.Input.Params[":seltype"]
		c.Data["seltype"] = sellist[0]
		newnianji := ""
		newkecheng := ""
		newjibie := ""
		newshengfen := ""
		newshiqu := ""
		nianji := sellist[1] //c.Ctx.Input.Params[":nianji"]
		c.Data["nianji"] = sellist[1]
		fmt.Println(nianji)
		if nianji != "" {
			//根据学龄段名称查询此学龄段主键id
			schoolage, errage := models.GetSchoolagesByName(nianji)
			if errage == nil && schoolage != nil {
				newnianji = `%` + strconv.Itoa(schoolage.Id) + `%`
			} else {
				newnianji = `%%`
			}
		} else if nianji == "" {
			newnianji = `%%`
		}
		kecheng := sellist[2] //c.Ctx.Input.Params[":kecheng"]
		c.Data["kecheng"] = sellist[2]
		if kecheng != "" {
			newkecheng = kecheng
		} else if kecheng == "" {
			newkecheng = `%%`
		}
		jibie := sellist[3] //c.Ctx.Input.Params[":jibie"]
		c.Data["jibie"] = sellist[3]
		if jibie != "" {
			newjibie = jibie
		} else if jibie == "" {
			newjibie = `%%`
		}
		shengfen := sellist[4] //c.Ctx.Input.Params[":shengfen"]
		c.Data["shengfen"] = sellist[4]
		if shengfen != "" {
			newshengfen = shengfen
		} else if shengfen == "" {
			newshengfen = `%%`
		}
		shiqu := sellist[5] //c.Ctx.Input.Params[":shiqu"]
		c.Data["shiqu"] = sellist[5]
		if shiqu != "" {
			newshiqu = shiqu
		} else if shiqu == "" {
			newshiqu = `%%`
		}
		seltype, _ := strconv.Atoi(seltypestr)
		//fmt.Println(seltypestr + "," + newjibie + "," + newkecheng + "," + newnianji + "," + newshengfen + "," + newshiqu)

		page := sellist[6]              //c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
		size := sellist[7]              //c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
		pages, _ := strconv.Atoi(page)  //传来的页数
		rows, _ := strconv.Atoi(size)   //传来的显示行数
		truepages := (pages - 1) * rows //计算舍弃多少行
		limit := rows                   //显示行数
		offset := truepages             //舍弃行数	//新加--------结束--------
		v, err := models.GetUserinformationAllTeacher2(seltype, newnianji, newkecheng, newjibie, newshengfen, newshiqu, offset, limit)
		vCount, _ := models.GetUserinformationAllTeacherCount2(seltype, newnianji, newkecheng, newjibie, newshengfen, newshiqu)
		c.Data["allUserscount"] = vCount
		//fmt.Println(v)
		var allteacherstr string = ""
		if err == nil && v != nil {
			//identUser, _ := strconv.Atoi(c.Ctx.GetCookie("identityid")) //身份
			var btnstr string = ""
			var btnhou string = ""
			for i := 0; i < len(v); i++ {
				var onlinestate int = v[i].OnlineState
				var onlinestr string = `<span style="font-size:8px;color:#666;">（离线）</span>`
				if onlinestate > 0 {
					onlinestr = `<span style="font-size:7px;">（在线）</span>`
					btnstr = `<a  onclick="trylisten(` + strconv.Itoa(v[i].Id) + `)"`
					btnhou = ` class="btn btn-warning messabtn" style="width: 100%;">立即试听</a>`
				} else {
					btnstr = `<a  onclick="setMessagetoT(` + strconv.Itoa(v[i].Id) + `)"`
					btnhou = `" class="btn btn-warning messabtn" style="width: 100%;">给他留言</a>`
				}
				//fmt.Println(v[i].UserName)
				var items string = `<div class="row" style="border-bottom: 1px solid #F1F1F1">
									<div class=" col-sm-2">
									<div class="teamainleft">
									<a href="http://` + models.OnlineUrl + `/orange/Teacher/TeacherInformation/` + strconv.Itoa(v[i].Id) + `" onclick="addliulan(` + strconv.Itoa(v[i].Id) + `)">
										<img class="listheadimg" src="` + v[i].AvatarPath + `" />
									</a>
									</div></div>
									<div class="col-sm-7 teamainmiddle">
									<a href="http://` + models.OnlineUrl + `/orange/Teacher/TeacherInformation/` + strconv.Itoa(v[i].Id) + `" onclick="addliulan(` + strconv.Itoa(v[i].Id) + `)">
										<span>` + v[i].UserName + onlinestr + `</span>
									</a>
									<p>` + v[i].SchoolName + ` | ` + v[i].DegreeName + ` | ` + strconv.Itoa(v[i].LevelYear) + `级 | ` + v[i].Professional + `</p>
									<p>主辅导课：` + v[i].CourseNameZhu + ` | 辅辅导课：` + v[i].CourseNameFu + `</p></div>
									<div class="col-sm-3"><div class="teamainright"><p>` + strconv.FormatFloat(v[i].UnitPrice, 'f', -1, 64) + `元/小时</p><br />` + btnstr + btnhou + `
									</div>
									</div>
									</div>`
				allteacherstr = allteacherstr + items
			}
			c.Data["allteacherstr"] = allteacherstr
			//fmt.Println(allteacherstr)
		}
	}

	c.TplNames = "teacher_Model.html" //
}

// 查看老师
// @Title TeacherInformation
// @Description TeacherInformation the TbUser
// @Param			"The id you want to TeacherInformation"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherInformation/:tid [get]
func (c *TeacherController) TeacherInformation() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":tid"]
	userid, _ := strconv.Atoi(idStr) //获取点击的老师主键id
	var vuser models.UserinformationTeacherModu
	vuser, _ = models.GetUserinformationTeacherModu(userid)
	c.Data["AvatarPath"] = vuser.AvatarPath
	c.Data["userid"] = "00000" + strconv.Itoa(vuser.Id)
	c.Data["Teacheruserid"] = vuser.Id
	c.Data["UserName"] = vuser.UserName
	c.Data["AllPerson"] = vuser.AllPerson
	c.Data["AllTime"] = vuser.AllTime
	c.Data["AllTimeMouth"] = vuser.AllTimeMouth
	c.Data["SchoolName"] = vuser.SchoolName
	c.Data["DegreeName"] = vuser.DegreeName
	c.Data["LevelYear"] = vuser.LevelYear
	c.Data["Professional"] = vuser.Professional
	c.Data["CourseName"] = vuser.CourseName
	c.Data["BriefIntroduction"] = vuser.BriefIntroduction
	c.Data["UserHobby"] = vuser.UserHobby
	c.Data["UnitPrice"] = vuser.UnitPrice

	c.Data["CourseName"] = vuser.CourseName
	zhucourse, fuerr := models.GetRemedialcoursesMain(userid, 0)
	fmt.Println(zhucourse)
	var fuzhu string = ""
	if zhucourse != nil && fuerr == nil {
		for i := 0; i < len(zhucourse); i++ {
			fuzhu += zhucourse[i].CourseName
			fuzhu += " "
		}
	}
	c.Data["CourseNameFu"] = fuzhu

	c.TplNames = "teacherlist.html" //跳到老师个人中心
}

// 给他留言
// @Title TeacherMessage
// @Description TeacherMessage the TbUser
// @Param			"The id you want to TeacherMessage"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherMessage/:tid [get]
func (c *TeacherController) TeacherMessage() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":tid"]
	userid, _ := strconv.Atoi(idStr)
	var vuser models.UserinformationTeacher
	vuser, _ = models.GetUserinformationTeacher(userid)
	fmt.Println(vuser)
	c.Data["AvatarPath"] = vuser.AvatarPath
	c.Data["userid"] = "00000" + strconv.Itoa(vuser.Id)
	c.Data["teacherid"] = vuser.Id
	c.Data["UserName"] = vuser.UserName
	c.Data["CourseName"] = vuser.CourseName

	c.TplNames = "message.html" //跳到给老师留言页面
}

// 查看一条精彩问答
// @Title ProblemModel
// @Description ProblemModel the TbUser
// @Param			"The id you want to ProblemModel"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /ProblemModel/:adkid [get]
func (c *TeacherController) ProblemModel() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":adkid"]
	askid, _ := strconv.Atoi(idStr)
	var vask models.QuestionaskJingCaiOne
	vask, _ = models.GetQuestionaskByJingCaiOne(askid)
	fmt.Println(vask)
	c.Data["Title"] = vask.Title
	c.Data["Contents"] = vask.Contents
	var timez string = vask.BadeTime.Format("2006-01-02")
	c.Data["BadeTime"] = timez
	c.Data["UserName"] = vask.UserName
	c.Data["Hname"] = vask.Hname
	c.Data["HuiDaContents"] = vask.HuiDaContents
	c.Data["AnsTime"] = vask.AnsTime
	c.Data["AvatarPath"] = vask.AvatarPath
	c.Data["HuiDaAvatarPath"] = vask.HuiDaAvatarPath
	c.Data["AnswerUserId"] = "00000" + strconv.Itoa(vask.AnswerUserId)
	c.Data["userid"] = vask.AnswerUserId
	c.Data["SchoolName"] = vask.SchoolName

	c.TplNames = "problem_main.html" //
}

// 老师回答一条问题
// @Title ProblemAnswer
// @Description ProblemAnswer the TbUser
// @Param			"The id you want to ProblemAnswer"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /ProblemAnswer/:adkid [get]
func (c *TeacherController) ProblemAnswer() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":adkid"]
	askid, _ := strconv.Atoi(idStr)
	var vask models.QuestionaskJingCaiOne
	vask, _ = models.GetQuestionaskByJingCaiOne(askid)
	fmt.Println(vask)
	c.Data["Title"] = vask.Title
	c.Data["Contents"] = vask.Contents
	c.Data["BadeTime"] = vask.BadeTime
	c.Data["UserName"] = vask.Hname
	c.Data["Hname"] = vask.UserName
	c.Data["HuiDaContents"] = vask.HuiDaContents
	c.Data["AnsTime"] = vask.AnsTime
	c.Data["AvatarPath"] = vask.HuiDaAvatarPath
	c.Data["HuiDaAvatarPath"] = vask.AvatarPath
	c.Data["AnswerUserId"] = "00000" + strconv.Itoa(vask.AskUserId)
	c.Data["userid"] = vask.AskUserId //提问用户id
	c.Data["SchoolName"] = vask.UserSchoolName
	c.Data["AnswerId"] = vask.AnswerId
	c.Data["QuestionId"] = vask.Id
	c.TplNames = "problem_answer.html" //
}

// 我要提问
// @Title UserAskQuestion
// @Description UserAskQuestion the TbUser
// @Param			"The id you want to UserAskQuestion"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /UserAskQuestion/ [get]
func (c *TeacherController) UserAskQuestion() {
	c.Data["Website"] = models.OnlineUrl

	c.TplNames = "problem.html" //
}

// 老师从个人中心进入编辑预约
// @Title TeacherSetMeet
// @Description TeacherSetMeet the TbUser
// @Param			"The id you want to TeacherSetMeet"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherSetMeet/ [get]
func (c *TeacherController) TeacherSetMeet() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "reservationteacher.html" //
}

// 学生点击老师，预约此老师课程
// @Title StudentSetTeacherMeet
// @Description StudentSetTeacherMeet the TbUser
// @Param			"The id you want to StudentSetTeacherMeet"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /StudentSetTeacherMeet/:tid [get]
func (c *TeacherController) StudentSetTeacherMeet() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":tid"]
	tid, _ := strconv.Atoi(idStr)
	c.Data["Teacherid"] = tid
	userinfo, usererr := models.GetUserinformationTeacher(tid)
	if usererr == nil {
		c.Data["AvatarPath"] = userinfo.AvatarPath
		c.Data["UserName"] = userinfo.UserName
		c.Data["GradeName"] = userinfo.GradeName   //所教年级
		c.Data["CourseName"] = userinfo.CourseName //主辅导课
		c.Data["UnitPrice"] = userinfo.UnitPrice   //课时费
	}
	c.TplNames = "reservation.html" //
}

// 老师跳页，跳到听课页面
// @Title TeacherOnlineClass
// @Description TeacherOnlineClass the TbUser
// @Param			"The id you want to TeacherOnlineClass"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherOnlineClass/:onlineid [get]
func (c *TeacherController) TeacherOnlineClass() {
	c.Data["Website"] = models.OnlineUrl
	onlineidStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(onlineidStr)
	c.Data["onlineid"] = onlineid
	c.Ctx.SetCookie("onlinebookid", onlineidStr) //当前老师进入试听信息主键
	c.TplNames = "tk.html"                       //
}

// 学生跳页，跳到听课页面
// @Title StudentOnlineClass
// @Description StudentOnlineClass the TbUser
// @Param			"The id you want to StudentOnlineClass"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /StudentOnlineClass/:onlineid [get]
func (c *TeacherController) StudentOnlineClass() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	c.Data["onlineid"] = onlineid
	c.Ctx.SetCookie("onlinebookid", strconv.Itoa(onlineid))
	c.TplNames = "tk_student.html" //
}

// 老师进入试听课程
// @Title TeacherTryListenClass
// @Description TeacherTryListenClass the TbUser
// @Param			"The id you want to TeacherTryListenClass"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /TeacherTryListenClass/:listenid [get]
func (c *TeacherController) TeacherTryListenClass() {
	c.Data["Website"] = models.OnlineUrl
	onlineidStr := c.Ctx.Input.Params[":listenid"]
	//listenid, _ := strconv.Atoi(onlineidStr)
	//c.Data["listenid"] = listenid
	c.Ctx.SetCookie("onlinelistenid", onlineidStr) //当前老师进入试听信息主键
	c.TplNames = "listenteacher.html"              //
}

// 学生进入试听
// @Title StudentTryListenClass
// @Description StudentTryListenClass the TbUser
// @Param			"The id you want to StudentTryListenClass"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /StudentTryListenClass/:listenid [get]
func (c *TeacherController) StudentTryListenClass() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":listenid"]
	listenid, _ := strconv.Atoi(idStr)
	c.Data["listenid"] = listenid
	c.TplNames = "listenstudent.html" //
}

// 在线课堂结束后进入的页面
// @Title ClassOverHtml
// @Description ClassOverHtml the TbUser
// @Param			"The id you want to ClassOverHtml"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /ClassOverHtml/ [get]
func (c *TeacherController) ClassOverHtml() {
	c.Data["Website"] = models.OnlineUrl

	c.TplNames = "classover.html" //
}

// 在线试听结束后进入的页面
// @Title ListenOverHtml
// @Description ListenOverHtml the TbUser
// @Param			"The id you want to ListenOverHtml"
// @Success 200 {object} models.TbUser
// @Failure 403
// @router /ListenOverHtml/ [get]
func (c *TeacherController) ListenOverHtml() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "listenover.html" //
}
