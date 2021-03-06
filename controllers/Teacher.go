package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"orange/models"
	"strconv"
	"strings"
	"time"
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

// @Title TeacherList
// @Description 展示找老师页面
// @Param seltype query int true 查询类型，排序类型
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
				var userheadurl string = ""
				if v[i].AvatarPath+"" == "" {
					userheadurl = "images/PersonHeadImg/moren.png"
				} else {
					userheadurl = v[i].AvatarPath
				}
				//fmt.Println(v[i].UserName)
				var items string = `<div class="row" style="border-bottom: 1px solid #F1F1F1">
									<div class=" col-sm-2">
									<div class="teamainleft">
									<a onclick="javascript:window.open(` + `'` + `http://` + models.OnlineUrl + `/orange/Teacher/TeacherInformation/` + strconv.Itoa(v[i].Id) + `'` + `);addliulan(` + strconv.Itoa(v[i].Id) + `)">
										<img class="listheadimg" src="` + userheadurl + `" />
									</a>
									</div></div>
									<div class="col-sm-7 teamainmiddle">
									<a onclick="javascript:window.open(` + `'` + `http://` + models.OnlineUrl + `/orange/Teacher/TeacherInformation/` + strconv.Itoa(v[i].Id) + `'` + `);addliulan(` + strconv.Itoa(v[i].Id) + `)">
										<span>` + v[i].UserName + onlinestr + `</span>
									</a>
									<p>` + v[i].SchoolName + ` | ` + v[i].DegreeName + ` | ` + strconv.Itoa(v[i].LevelYear) + `级 | ` + v[i].Professional + `</p>
									<p>主辅导课：` + v[i].CourseNameZhu + ` </p></div>
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

// @Title TeacherInformation
// @Description 跳页到查看老师
// @Param tid query int true 老师主键id
// @router /TeacherInformation/:tid [get]
func (c *TeacherController) TeacherInformation() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":tid"]
	userid, _ := strconv.Atoi(idStr) //获取点击的老师主键id
	var vuser models.UserinformationTeacherModu
	//获取当前时间的月初和月末

	vuser, _ = models.GetUserinformationTeacherModu(userid)
	fmt.Println("展示老师的信息为：")
	fmt.Println(vuser)
	c.Data["AvatarPath"] = vuser.AvatarPath
	c.Data["userid"] = "00000" + strconv.Itoa(vuser.Id)
	c.Data["Teacheruserid"] = vuser.Id
	c.Data["UserName"] = vuser.UserName
	c.Data["AllPerson"] = vuser.AllPerson
	//计算课时
	fa, _ := strconv.ParseFloat(strconv.Itoa(vuser.AllTime), 64)
	allhour := fmt.Sprintf("%.1f", fa/60)
	//计算课时
	fa2, _ := strconv.ParseFloat(strconv.Itoa(vuser.AllTimeMouth), 64)
	allhour2 := fmt.Sprintf("%.1f", fa2/60)
	c.Data["AllTime"] = allhour
	c.Data["AllTimeMouth"] = allhour2
	c.Data["SchoolName"] = vuser.SchoolName
	c.Data["DegreeName"] = vuser.DegreeName
	c.Data["LevelYear"] = vuser.LevelYear
	c.Data["Professional"] = vuser.Professional
	c.Data["CourseName"] = vuser.CourseName
	c.Data["BriefIntroduction"] = vuser.BriefIntroduction
	c.Data["UserHobby"] = vuser.UserHobby
	c.Data["UnitPrice"] = vuser.UnitPrice

	c.Data["CourseName"] = vuser.CourseName

	fmt.Println("月总课时")
	fmt.Println(vuser.AllTimeMouth)
	zhucourse, fuerr := models.GetRemedialcoursesMain(userid, 0)
	fmt.Println(zhucourse)
	var fuzhu string = ""
	if zhucourse != nil && fuerr == nil {
		for i := 0; i < len(zhucourse); i++ {
			fuzhu += zhucourse[i].CourseName
			fuzhu += " "
		}
	}
	allclasscount, _ := models.GetOnlinecourserecordByTidCount(userid)
	allpingluncount, _ := models.GetOnlinecourseevaluationByTidCount(userid)
	c.Data["nowallclasscount"] = allclasscount
	c.Data["nowallPingjia"] = allpingluncount

	c.Data["ceshi"] = time.Now()
	c.Data["CourseNameFu"] = fuzhu
	c.TplNames = "teacherlist.html" //跳到老师个人中心
}

// @Title TeacherMessage
// @Description 跳页到给老师留言页面
// @Param tid query int true 老师主键id
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

// @Title ProblemModel
// @Description 跳页到查看一条问答
// @Param adkid query int true 问答主键id
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

// @Title ProblemModel
// @Description 跳页到查看一条精彩问答值展示问答
// @Param adkid query int true 问答主键id
// @router /ProblemModelShow/:adkid [get]
func (c *TeacherController) ProblemModelshow() {
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

	c.TplNames = "problem_show.html" //
}

// @Title ProblemAnswer
// @Description 跳页到老师回答一条问题
// @Param adkid query int true 问答主键id
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

// @Title UserAskQuestion
// @Description 跳页到我要提问页面
// @router /UserAskQuestion/ [get]
func (c *TeacherController) UserAskQuestion() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "problem.html" //
}

// @Title UserAskQuestion
// @Description 添加一条提问信息
// @Param Questionask form models.Questionask true 一条提问实体
// @router /UserAskQuestion/ [Post]
func (c *TeacherController) UserAskQuestion2() {
	c.Data["Website"] = models.OnlineUrl
	userid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	var class []string = c.Ctx.Input.Request.Form["selClass"]       //
	var teacherid []string = c.Ctx.Input.Request.Form["selteacher"] //
	var times []string = c.Ctx.Input.Request.Form["txtdate"]        //
	var title []string = c.Ctx.Input.Request.Form["txtTitle"]
	var content []string = c.Ctx.Input.Request.Form["hiddencontent"]
	var money []string = c.Ctx.Input.Request.Form["selmoney"]
	var question models.Questionask
	question.AskUserId = userid
	question.AnswerUserId, _ = strconv.Atoi(teacherid[0])
	question.GCId, _ = strconv.Atoi(class[0])
	question.Title = title[0]
	question.Contents = content[0]
	question.BadeTime = time.Now()
	fmt.Println("当前时间：")
	fmt.Println(question.BadeTime)
	question.AmountMoney, _ = strconv.ParseFloat(money[0], 64)
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02", times[0], loc) //time类型
	question.EndTime = t1
	question.IsSee = 0
	addint, aderr := models.AddQuestionask(&question)
	if aderr == nil && addint > 0 {
		var fontsize models.Frozenfunds
		fontsize.UserId = userid
		fontsize.FrozenMoney, _ = strconv.ParseFloat(money[0], 64)
		fontsize.FrozenType = 1
		addintstr := strconv.FormatInt(int64(addint), 10)
		fontsize.BusinessId, _ = strconv.Atoi(addintstr)
		fontsize.FrozenTime = time.Now()
		fontsize.FrozenState = 1
		addfontid, fonterr := models.AddFrozenfunds(&fontsize)
		if fonterr == nil && addfontid > 0 {
			useraccount, _ := models.GetAccountfundsByuid(userid)
			useraccount.Balance = useraccount.Balance - fontsize.FrozenMoney
			err := models.UpdateAccountfundsById(&useraccount)
			if err == nil {

			}
		}
	}

	c.Redirect("/orange/Main/UserStudent/3", 302)
}

// @Title TeacherSetMeet
// @Description 老师从个人中心进入编辑预约
// @router /TeacherSetMeet/ [get]
func (c *TeacherController) TeacherSetMeet() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "reservationteacher.html" //
}

// @Title StudentSetTeacherMeet
// @Description 学生点击老师，预约此老师课程
// @Param tid query int true 老师主键id
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

// @Title TeacherOnlineClass
// @Description 老师跳页，跳到听课页面
// @Param onlineid query int true 预约信息主键id
// @router /TeacherOnlineClass/:onlineid [get]
func (c *TeacherController) TeacherOnlineClass() {
	c.Data["Website"] = models.OnlineUrl
	onlineidStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(onlineidStr)
	c.Data["onlineid"] = onlineid
	c.Ctx.SetCookie("onlinebookid", onlineidStr) //当前老师进入试听信息主键
	c.TplNames = "tk.html"                       //
}

// @Title StudentOnlineClass
// @Description 学生跳页，跳到听课页面
// @Param onlineid query int true 预约信息主键id
// @router /StudentOnlineClass/:onlineid [get]
func (c *TeacherController) StudentOnlineClass() {
	c.Data["Website"] = models.OnlineUrl
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	c.Data["onlineid"] = onlineid
	c.Ctx.SetCookie("onlinebookid", strconv.Itoa(onlineid))
	c.TplNames = "tk_student.html" //
}

// @Title TeacherTryListenClass
// @Description 老师进入试听课程
// @Param listenid query int true 试听信息主键id
// @router /TeacherTryListenClass/:listenid [get]
func (c *TeacherController) TeacherTryListenClass() {
	c.Data["Website"] = models.OnlineUrl
	onlineidStr := c.Ctx.Input.Params[":listenid"]
	//listenid, _ := strconv.Atoi(onlineidStr)
	//c.Data["listenid"] = listenid
	c.Ctx.SetCookie("onlinelistenid", onlineidStr) //当前老师进入试听信息主键
	c.TplNames = "listenteacher.html"              //
}

// @Title StudentTryListenClass
// @Description 学生进入试听
// @Param listenid query int true 试听信息主键id
// @router /StudentTryListenClass/:listenid [get]
func (c *TeacherController) StudentTryListenClass() {
	c.Data["Website"] = models.OnlineUrl
	onlineidStr := c.Ctx.Input.Params[":listenid"]
	//listenid, _ := strconv.Atoi(idStr)
	c.Ctx.SetCookie("onlinelistenid", onlineidStr) //当前老师进入试听信息主键
	c.TplNames = "listenstudent.html"              //
}

// @Title ClassOverHtml
// @Description 在线课堂结束后进入的页面
// @router /ClassOverHtml/ [get]
func (c *TeacherController) ClassOverHtml() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "classover.html" //
}

// @Title ListenOverHtml
// @Description 在线试听结束后进入的页面
// @router /ListenOverHtml/ [get]
func (c *TeacherController) ListenOverHtml() {
	c.Data["Website"] = models.OnlineUrl
	c.TplNames = "listenover.html" //
}
