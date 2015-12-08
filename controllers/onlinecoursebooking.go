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
// @Failure 403 :id is empty
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

//6根据老师主键查询预约课程信息
// @Title GetOnlinecoursebookingByTid
// @Description GetOnlinecoursebookingByTid Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//6根据老师主键查询预约课程信息总条数
// @Title GetOnlinecoursebookingByTidCount
// @Description GetOnlinecoursebookingByTidCount Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//20根据学生主键查询预约课程信息
// @Title GetOnlinecoursebookingByUid
// @Description GetOnlinecoursebookingByUid Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//20根据学生主键查询预约课程信息
// @Title GetOnlinecoursebookingByUidCount
// @Description GetOnlinecoursebookingByUidCount Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//20查询学生没有上过的预约课程
// @Title GetOnlinecoursebookingBySidNotOn
// @Description GetOnlinecoursebookingBySidNotOn Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//20查询学生没有上过的预约课程总条数
// @Title GetOnlinecoursebookingBySidNotOnCount
// @Description GetOnlinecoursebookingBySidNotOnCount Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

//查询学生预约某个老师某天预约了几次课程
// @Title GetOnlinecoursebookingBySTidTime
// @Description GetOnlinecoursebookingBySTidTime Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

// 38.根据老师主键id，和时间段查询此时间段预约课程信息
// @Title GetOnlinecoursebookingByTidTime
// @Description GetOnlinecoursebookingByTidTime Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
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

// @Title Get
// @Description get Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
// @router /GetBHtecher/:onlineid [get]
func (c *OnlinecoursebookingController) GetOss() {
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	fmt.Println(onlineid)
	l, err := models.Getecherlession(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(l)
		c.Data["json"] = map[string]string{"url": l}
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinecoursebooking by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebooking
// @Failure 403 :id is empty
// @router /GetBstudent/:onlineid [get]
func (c *OnlinecoursebookingController) GetOe() {
	idStr := c.Ctx.Input.Params[":onlineid"]
	onlineid, _ := strconv.Atoi(idStr)
	l, err := models.Getstudentlession(onlineid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(l)
		c.Data["json"] = l
	}
	c.ServeJson()
}
