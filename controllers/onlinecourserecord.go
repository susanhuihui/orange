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

// oprations for Onlinecourserecord
type OnlinecourserecordController struct {
	beego.Controller
}

func (c *OnlinecourserecordController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Onlinecourserecord
// @Param	body		body 	models.Onlinecourserecord	true		"body for Onlinecourserecord content"
// @Success 200 {int} models.Onlinecourserecord.Id
// @Failure 403 body is empty
// @router /AddOnlinecourserecord/ [post]
func (c *OnlinecourserecordController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Onlinecourserecord
	fmt.Println(jsonS)
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddOnlinecourserecord(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordById/:id [get]
func (c *OnlinecourserecordController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecourserecordById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetOnlinecourserecordByBookid
// @Description GetOnlinecourserecordByBookid Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordByBookid/:bookid [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordByBookid() {
	idStr := c.Ctx.Input.Params[":bookid"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecourserecordByBookid(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//4查询老师全部课程信息
// @Title GetOnlinecourserecordByTid
// @Description GetOnlinecourserecordByTid Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordByTid/:userid/:page/:size [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetOnlinecourserecordByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//4查询老师全部课程信息总条数
// @Title GetOnlinecourserecordByTidCount
// @Description GetOnlinecourserecordByTidCount Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordByTidCount/:userid [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecourserecordByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//18.查询学生全部课程
// @Title GetOnlinecourserecordByUid
// @Description GetOnlinecourserecordByUid Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordByUid/:userid/:page/:size [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetOnlinecourserecordByUid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//18.查询学生全部课程
// @Title GetOnlinecourserecordByUidCount
// @Description GetOnlinecourserecordByUidCount Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordByUidCount/:userid [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordByUidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecourserecordByUidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//40.查询给我上过课的老师们
// @Title GetOnlinecourserecordTeacherByUid
// @Description GetOnlinecourserecordTeacherByUid Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordTeacherByUid/:userid [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordTeacherByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecourserecordTeacherByUid(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//40.查询给我上过课的老师们
// @Title GetOnlinecourserecordTeacherByUCid
// @Description GetOnlinecourserecordTeacherByUCid Onlinecourserecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GetOnlinecourserecordTeacherByUCid
// @Failure 403 :id is empty
// @router /GetOnlinecourserecordTeacherByUCid/:userid/:classid [get]
func (c *OnlinecourserecordController) GetOnlinecourserecordTeacherByUCid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	classidStr := c.Ctx.Input.Params[":classid"]
	classid, _ := strconv.Atoi(classidStr)
	v, err := models.GetOnlinecourserecordTeacherByUCid(userid, classid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Onlinecourserecord
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403
// @router /GetAllOnlinecourserecord/:page/:size [get]
func (c *OnlinecourserecordController) GetAll() {
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

	l, err := models.GetAllOnlinecourserecord(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Onlinecourserecord
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Onlinecourserecord	true		"body for Onlinecourserecord content"
// @Success 200 {object} models.Onlinecourserecord
// @Failure 403 :id is not int
// @router /UpdateOnlinecourserecordById/:id [post]
func (c *OnlinecourserecordController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Onlinecourserecord{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateOnlinecourserecordById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Onlinecourserecord
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteOnlinecourserecord/:id [delete]
func (c *OnlinecourserecordController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOnlinecourserecord(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
