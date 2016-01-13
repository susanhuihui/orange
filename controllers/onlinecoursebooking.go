package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"orange/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Onlinecoursebooking
type OnlinecoursebookingController struct {
	beego.Controller
}

func (c *OnlinecoursebookingController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Onlinecoursebooking
// @Param	body		body 	models.Onlinecoursebooking	true		"body for Onlinecoursebooking content"
// @Success 200 {int} models.Onlinecoursebooking.Id
// @Failure 403 body is empty
// @router /AddOnlinecoursebooking/ [post]
func (c *OnlinecoursebookingController) Post() {
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Onlinecoursebooking
	fmt.Println(jsonS)
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddOnlinecoursebooking(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure Error
// @router /GetOnlinecoursebookingById/:id [get]
func (c *OnlinecoursebookingController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecoursebookingById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//6
// @Title GetOnlinecoursebookingByTid
// @Description 根据老师主键查询预约课程信息
// @Param	userid path int true 用户信息主键id
// @Param	page path int true 获取第几页
// @Param	size path int true 获取多少行
// @Success 200 {object} models.OnlinecoursebookingList
// @Failure Error
// @router /GetOnlinecoursebookingByTid/:userid/:page/:size [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetOnlinecoursebookingByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//6
// @Title GetOnlinecoursebookingByTidCount
// @Description 根据老师主键查询预约课程信息总条数
// @Param	userid path int true 用户信息主键id
// @Success 200 {int} json
// @Failure Error
// @router /GetOnlinecoursebookingByTidCount/:userid [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecoursebookingByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//20
// @Title GetOnlinecoursebookingByUid
// @Description 根据学生主键查询预约课程信息
// @Param	userid path int true 用户信息主键id
// @Param	page path int true 获取第几页
// @Param	size path int true 获取多少行
// @Success 200 {object} models.OnlinecoursebookingList
// @Failure Error
// @router /GetOnlinecoursebookingByUid/:userid/:page/:size [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetOnlinecoursebookingByUid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//20
// @Title GetOnlinecoursebookingByUidCount
// @Description 根据学生主键查询预约课程信息总条数
// @Param	userid path int true 用户信息主键id
// @Success 200 {int} json
// @Failure Error
// @router /GetOnlinecoursebookingByUidCount/:userid [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingByUidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecoursebookingByUidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//20
// @Title GetOnlinecoursebookingBySidNotOn
// @Description 查询学生没有上过的预约课程
// @Param	userid path int true 用户信息主键id
// @Param	page path int true 获取第几页
// @Param	size path int true 获取多少行
// @Success 200 {object} models.OnlinecoursebookingList
// @Failure Error
// @router /GetOnlinecoursebookingBySidNotOn/:userid/:page/:size [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingBySidNotOn() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetOnlinecoursebookingBySidNotOn(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//20
// @Title GetOnlinecoursebookingBySidNotOnCount
// @Description 查询学生没有上过的预约课程总条数
// @Param	userid path int true 用户信息主键id
// @Success 200 {int} json
// @Failure Error
// @router /GetOnlinecoursebookingBySidNotOnCount/:userid [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingBySidNotOnCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecoursebookingBySidNotOnCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//
// @Title GetOnlinecoursebookingBySTidTime
// @Description 查询学生预约某个老师某天预约了几次课程
// @Param	sid path int true 学生用户主键id
// @Param	tid path int true 老师用户主键id
// @Param	time1 path string true 开始时间
// @Param	time2 path string true 结束时间
// @Success 200 {int} json
// @Failure Error
// @router /GetOnlinecoursebookingBySTidTime/:sid/:tid/:time1/:time2 [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingBySTidTime() {
	sidStr := c.Ctx.Input.Params[":sid"]
	sid, _ := strconv.Atoi(sidStr)
	tidStr := c.Ctx.Input.Params[":tid"]
	tid, _ := strconv.Atoi(tidStr)
	time1Str := c.Ctx.Input.Params[":time1"]
	time2Str := c.Ctx.Input.Params[":time2"]
	v, err := models.GetOnlinecoursebookingBySTidTime(sid, tid, time1Str, time2Str)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// 38.
// @Title GetOnlinecoursebookingByTidTime
// @Description 根据老师主键id，和时间段查询此时间段预约课程信息
// @Param	userid path int true 老师用户主键id
// @Param	time1 path string true 开始时间
// @Success 200 {object} models.Onlinecoursebooking
// @Failure Error
// @router /GetOnlinecoursebookingByTidTime/:userid/:time1 [get]
func (c *OnlinecoursebookingController) GetOnlinecoursebookingByTidTime() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	time1 := c.Ctx.Input.Param(":time1")
	v, err := models.GetOnlinecoursebookingByTidTime(userid, time1)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Onlinecoursebooking
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403
// @router /GetAllOnlinecoursebooking/:page/:size [get]
func (c *OnlinecoursebookingController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)

	page := c.Ctx.Input.Param(":page")       //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size")       //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.ParseInt(page, 0, 0) //传来的页数
	rows, _ := strconv.ParseInt(size, 0, 0)  //传来的显示行数
	truepages := (pages - 1) * rows          //计算舍弃多少行
	limit := rows                            //显示行数
	offset := truepages                      //舍弃行数	//新加--------结束--------

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllOnlinecoursebooking(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Onlinecoursebooking
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Onlinecoursebooking	true		"body for Onlinecoursebooking content"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is not int
// @router /UpdateOnlinecoursebookingById/:id [post]
func (c *OnlinecoursebookingController) Put() {
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Onlinecoursebooking{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateOnlinecoursebookingById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Onlinecoursebooking
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteOnlinecoursebooking/:id [get]
func (c *OnlinecoursebookingController) DeleteOnlinecoursebooking() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOnlinecoursebooking(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title DeleteOnlinecoursebookingMeeting
// @Description DeleteOnlinecoursebookingMeeting the Onlinecoursebooking
// @Param	id		path 	string	true		"The id you want to DeleteOnlinecoursebookingMeeting"
// @Success 200 {string} DeleteOnlinecoursebookingMeeting success!
// @Failure 403 id is empty
// @router /DeleteOnlinecoursebookingMeeting/:id [get]
func (c *OnlinecoursebookingController) DeleteOnlinecoursebookingMeeting() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOnlinecoursebookingMeeting(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description 根据预约信息主键id查询此信息老师是否可以进入课堂
// @Param	onlineid		path 	string	true		预约信息主键id
// @Success url {string} url
// @Failure Error
// @router /GetBHtecher/:onlineid [get]
func (c *OnlinecoursebookingController) GetOss() {
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	if onlineid > 0 {
		c.Ctx.SetCookie("onlinebookid", strconv.Itoa(onlineid))
	}
	fmt.Println(onlineid)
	teacherinurl, err := models.Getecherlession3(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(teacherinurl)
		c.Data["json"] = map[string]string{"url": teacherinurl} // >0可以进入，-2会议室已存在一个人以上，老师不得进入
	}
	c.ServeJson()
}

// @Title GetOnlineClassTeacherurl
// @Description 获取老师进入白板的路径
// @Param	onlineid		path 	string	true		预约信息主键
// @Success url {string} url
// @Failure Error
// @router /GetOnlineClassTeacherurl/:onlineid [get]
func (c *OnlinecoursebookingController) GetOnlineClassTeacherurl() {
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	teacherinurl, err := models.GetOnlineClassTeacherurl(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(teacherinurl)
		c.Data["json"] = map[string]string{"url": teacherinurl} //-1:预约信息不存在 -2:用户不存在 -3:创建课堂失败
	}
	c.ServeJson()
}

// @Title Get
// @Description 获取学生进入白板的url
// @Param	onlineid		path 	string	true		预约信息主键id
// @Success json {string} url
// @Failure Error
// @router /GetBstudent/:onlineid [get]
func (c *OnlinecoursebookingController) GetOe() {
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	if onlineid > 0 {
		c.Ctx.SetCookie("onlinebookid", strconv.Itoa(onlineid))
	}
	l, err := models.Getstudentlession2(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(l)
		c.Data["json"] = l
	}
	c.ServeJson()
}

//
// @Title ClassPay
// @Description 根据预约表主键id结算此次课程
// @Param	onlineid		path 	string	true		预约信息主键id
// @Success json {string} resultshow
// @Failure Error
// @router /ClassPay/:onlineid [get]
func (c *OnlinecoursebookingController) ClassPay() {
	idStr := c.Ctx.Input.Params[":onlineid"] //预约表主键id
	onlineid, _ := strconv.Atoi(idStr)
	//	result, _ := models.GetOnlinecoursebookingrecordByTwoid(onlineid)
	//	fmt.Println(result)
	resultshow, err := SetUserClassPay(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(resultshow)
		c.Data["json"] = resultshow
	}
	c.ServeJson()
}

//结算课程
//传入参数：onlineid:预约信息主键id
//传出参数：resultmsg:返回结果 err:错误信息
//        resultmsg---  >0:结算成功，返回总分钟数
//        				0:已结算不操作
//        				-1:查询不到预约信息
//        				-2:没有课程时间记录（几乎没有这种可能）
//        				-3:没有查到冻结资金
//        				-4:没有查到用户账户信息
//        				-5:给用户返还钱失败
//        				-6:解冻资金失败
//        				-7:新增交易记录失败
//        				-8:更新预约信息状态失败
//        				-9:添加在线课程记录失败
//2015-12-19
//思路：1.查询预约信息 2.查询此次预约课程所有时间记录信息，计算总时间 3.根据学生主键找到此次预约冻结资金，解冻，根据冻结时间分配金额，4.钱打给老师，剩余退还学生，
//     5.新增一条老师打钱给学生的交易记录，6.将预约信息状态修改为已学习，已支付，7.新增一条课程记录信息
func SetUserClassPay(onlineid int) (resultmsg string, err error) {
	//根据预约信息主键id查询 一条预约信息
	onlineclass, geterr := models.GetOnlinecoursebookingById(onlineid)
	fmt.Println("获取预约信息")
	fmt.Println(onlineid)
	fmt.Println(onlineclass)
	fmt.Println("1")
	if geterr == nil && onlineclass.Id > 0 {
		overstu, _ := models.GetOnlinecoursebookingrecordByUid2(onlineclass.UserIdActive, onlineid)  //记录终止时间
		overtea, _ := models.GetOnlinecoursebookingrecordByUid2(onlineclass.UserIdPassive, onlineid) //记录终止时间
		fmt.Println(overstu)
		fmt.Println(overtea)
		allminutes := models.GetALLtimeminute(onlineid) //计算总分钟数
		if allminutes > 0 {
			if onlineclass.Leaming == 0 && onlineclass.Payment == 0 {
				//查询此次预约的冻结信息
				fonze, ferr := models.GetFrozenfundsByUidOnId(onlineclass.UserIdActive, 0, onlineid)
				fmt.Println("2")
				if ferr == nil && fonze.Id > 0 {
					//计算此次课程总费用
					allm, _ := strconv.ParseFloat(strconv.Itoa(allminutes), 64)
					alltm, _ := strconv.ParseFloat(strconv.Itoa(models.TotalMinute), 64)
					teachermoney := (allm / alltm) * fonze.FrozenMoney //（上课分钟/总分钟）*总钱数 = 应给老师多少钱
					returnmoney := fonze.FrozenMoney - teachermoney    //返还学生的钱
					resultmsg = models.SetUserMoney(onlineclass.UserIdPassive, teachermoney)
					resultmsg = models.SetUserMoney(onlineclass.UserIdActive, returnmoney)
					fmt.Println("3")
					if resultmsg == "1" { //钱各自打成功，解冻冻结资金信息
						fonze.FrozenState = 0
						upfonzerr := models.UpdateFrozenfundsById(&fonze)
						fmt.Println("4")
						if upfonzerr == nil { //解冻成功新增一条交易记录
							resultmsg = models.AddUserTransactionRecords(onlineclass.UserIdActive, onlineclass.UserIdPassive, teachermoney)
							fmt.Println("5")
							if resultmsg == "1" { //添加交易记录完成
								onlineclass.Payment = 1
								onlineclass.Leaming = 1
								uponerr := models.UpdateOnlinecoursebookingById(onlineclass)
								fmt.Println("6")
								if uponerr == nil {
									//新增一条课程记录信息
									var onlinerecord models.Onlinecourserecord
									onlinerecord.OCBId = onlineclass.Id
									onlinerecord.UserIdActive = onlineclass.UserIdActive
									onlinerecord.UserIdPassive = onlineclass.UserIdPassive
									onlinerecord.CourseContent = onlineclass.AppointMessage
									onlinerecord.StartTime = onlineclass.StartTime
									onlinerecord.EndTime = onlineclass.EndTime
									onlinerecord.UnitPrice = (allm / alltm)
									onlinerecord.TotalPrice = teachermoney
									onlinerecord.ClassNumber = allminutes
									addrecordid, recorderr := models.AddOnlinecourserecord(&onlinerecord)
									fmt.Println("7")
									if recorderr == nil && addrecordid > 0 {
										resultmsg = strconv.Itoa(allminutes) //返回总分钟数
										fmt.Println("8")
									} else {
										resultmsg = "-9"
									}
								} else {
									resultmsg = "-8"
								}
							} else {
								resultmsg = "-7"
							}
						} else {
							resultmsg = "-6"
						}
					}
				} else {
					resultmsg = "-3"
				}
			} else {
				resultmsg = "0"
			}
		} else {
			resultmsg = "-2" //没有课程时间记录
		}
	} else {
		resultmsg = "-1" //查询不到预约信息
	}
	return
}

//
// @Title ClassPay
// @Description 根据预约主键id计算此次预约课程共计上了多少分钟课程
// @Param	oid		path 	string	true		预约信息主键id
// @Success allminutes {int} resultshow
// @router /GetALLtimeminute/:oid [get]
func (c *OnlinecoursebookingController) GetALLtimeminute() {
	oidStr := c.Ctx.Input.Params[":oid"] //预约表主键id
	oid, _ := strconv.Atoi(oidStr)
	resultshow := models.GetALLtimeminute(oid)
	c.Data["json"] = resultshow
	c.ServeJson()
}
