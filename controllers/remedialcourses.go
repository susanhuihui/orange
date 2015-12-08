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

// oprations for Remedialcourses
type RemedialcoursesController struct {
	beego.Controller
}

func (c *RemedialcoursesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Remedialcourses
// @Param	body		body 	models.Remedialcourses	true		"body for Remedialcourses content"
// @Success 200 {int} models.Remedialcourses.Id
// @Failure 403 body is empty
// @router /AddRemedialcourses/ [post]
func (c *RemedialcoursesController) Post() {
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Remedialcourses
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddRemedialcourses(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Remedialcourses by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Remedialcourses
// @Failure 403 :id is empty
// @router /GetRemedialcoursesById/:id [get]
func (c *RemedialcoursesController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRemedialcoursesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//3.查询老师主辅导课程或辅辅导课程/学生的学习难点
// @Title GetRemedialcoursesMain
// @Description GetRemedialcoursesMain Remedialcourses by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Remedialcourses
// @Failure 403 :id is empty
// @router /GetRemedialcoursesMain/:userid/:ismain [get]
func (c *RemedialcoursesController) GetRemedialcoursesMain() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	ismainStr := c.Ctx.Input.Params[":ismain"]
	ismain, _ := strconv.Atoi(ismainStr)
	v, err := models.GetRemedialcoursesMain(userid, ismain)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Remedialcourses
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Remedialcourses
// @Failure 403
// @router /GetAllRemedialcourses/:page/:size [get]
func (c *RemedialcoursesController) GetAll() {
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

	l, err := models.GetAllRemedialcourses(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Remedialcourses
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Remedialcourses	true		"body for Remedialcourses content"
// @Success 200 {object} models.Remedialcourses
// @Failure 403 :id is not int
// @router /UpdateRemedialcoursesById/:id [post]
func (c *RemedialcoursesController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Remedialcourses{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateRemedialcoursesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Remedialcourses
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteRemedialcourses/:id [get]
func (c *RemedialcoursesController) DeleteRemedialcourses() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRemedialcourses(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title UpdateStudentClass
// @Description UpdateStudentClass the Remedialcourses
// @Param	id		path 	string	true		"The id you want to UpdateStudentClass"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /UpdateStudentClass/:sid/:classidlist [get]
func (c *RemedialcoursesController) UpdateStudentClass() {
	idStr := c.Ctx.Input.Params[":sid"]
	sid, _ := strconv.Atoi(idStr)
	classidlist := c.Ctx.Input.Params[":classidlist"]
	sellist := strings.Split(classidlist, ",") //总长度为8
	for i := 0; i < len(sellist); i++ {
		fmt.Println(sellist[i])
	}
	var err error = nil
	if sid > 0 && classidlist != "" {
		userclass, _ := models.GetRemedialcoursesMain(sid, 0)
		if userclass != nil {
			for a := 0; a < len(userclass); a++ { //循环旧的是否存在新的集合中，是不做操作，否删除
				var have bool = false
				for i := 0; i < len(sellist); i++ {
					selid, _ := strconv.Atoi(sellist[i])
					if userclass[a].CoursesId == selid {
						have = true
					}
				}
				if have == false {
					//shanchu
					delresult := models.DeleteRemedialcourses(userclass[a].Id) //删除没有勾选的项
					if delresult != nil {
						err = delresult
					}
				}
			}
			for j := 0; j < len(sellist); j++ { //循环新的是否存在旧的集合中，是不做操作，否添加
				var haveadd bool = false
				selids, _ := strconv.Atoi(sellist[j])
				for b := 0; b < len(userclass); b++ {
					if selids == userclass[b].CoursesId {
						haveadd = true
					}
				}
				if haveadd == false {
					var addrc models.Remedialcourses
					addrc.UserId = sid
					addrc.CoursesId = selids
					addrc.IsMain = 0
					addresult, adderr := models.AddRemedialcourses(&addrc)
					fmt.Println(addresult)
					if adderr != nil {
						err = adderr
					}
				}
			}
		}
	}
	if err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
