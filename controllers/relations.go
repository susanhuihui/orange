package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"orange/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// oprations for Relations
type RelationsController struct {
	beego.Controller
}

func (c *RelationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Relations
// @Param	body		body 	models.Relations	true		"body for Relations content"
// @Success 200 {int} models.Relations.Id
// @Failure 403 body is empty
// @router /AddRelations/ [post]
func (c *RelationsController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Relations
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddRelations(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Relations by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Relations
// @Failure 403 :id is empty
// @router /GetRelationsById/:id [get]
func (c *RelationsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRelationsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetRelationsByST
// @Description 查询一条师生某个关系信息
// @Param	sid		path 	string	true	学生主键id
// @Param	tid		path 	string	true		老师主键id
// @Param	guanxi		path 	string	true		关系关键词
// @Success 200 {object} models.Relations
// @Failure Error
// @router /GetRelationsByST/:sid/:tid/:guanxi [get]
func (c *RelationsController) GetRelationsByST() {
	sidStr := c.Ctx.Input.Params[":sid"]
	sid, _ := strconv.Atoi(sidStr)
	tidStr := c.Ctx.Input.Params[":tid"]
	tid, _ := strconv.Atoi(tidStr)
	guanxi := c.Ctx.Input.Params[":guanxi"]
	v, err := models.GetRelationsByST(sid, tid, guanxi)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//9老师查看谁看过我
// @Title GetRelationsByTid
// @Description 老师查看谁看过我
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Param	page		path 	string	true 获取第几页
// @Param	size		path 	string	true 获取多少行
// @Success 200 {object} models.Relations
// @Failure Error
// @router /GetRelationsByTid/:userid/:guanxi/:page/:size [get]
func (c *RelationsController) GetRelationsByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetRelationsByTid(userid, guanxi, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//9老师查看谁看过我总条数
// @Title GetRelationsByTidCount
// @Description 老师查看谁看过我总条数
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Success json {int} json
// @Failure Error
// @router /GetRelationsByTidCount/:userid/:guanxi [get]
func (c *RelationsController) GetRelationsByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetRelationsByTidCount(userid, guanxi)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//21.查询学生全部关注的老师
// @Title GetRelationsByUid
// @Description 查询学生全部关注的老师
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Param	page		path 	string	true 获取第几页
// @Param	size		path 	string	true 获取多少行
// @Success 200 {object} models.Relations
// @Failure Error
// @router /GetRelationsByUid/:userid/:guanxi/:page/:size [get]
func (c *RelationsController) GetRelationsByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetRelationsByUid(userid, guanxi, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//21.查询学生全部关注的老师总条数
// @Title GetRelationsByUidCount
// @Description 查询学生全部关注的老师总条数
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Success json {int} json
// @Failure Error
// @router /GetRelationsByUidCount/:userid/:guanxi [get]
func (c *RelationsController) GetRelationsByUidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetRelationsByUidCount(userid, guanxi)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//31.学生查看我浏览过谁
// @Title GetRelationsByUidSee
// @Description 学生查看我浏览过谁
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Param	page		path 	string	true 获取第几页
// @Param	size		path 	string	true 获取多少行
// @Success 200 {object} models.Relations
// @Failure Error
// @router /GetRelationsByUidSee/:userid/:guanxi/:page/:size [get]
func (c *RelationsController) GetRelationsByUidSee() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetRelationsByUidSee(userid, guanxi, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//31.学生查看我浏览过谁总条数
// @Title GetRelationsByUidSeeCount
// @Description 学生查看我浏览过谁总条数
// @Param	userid		path 	string	true 用户主键id
// @Param	guanxi		path 	string	true 关系关键词
// @Success json {int} json
// @Failure Error
// @router /GetRelationsByUidSeeCount/:userid/:guanxi [get]
func (c *RelationsController) GetRelationsByUidSeeCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	guanxi := c.Ctx.Input.Params[":guanxi"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetRelationsByUidSeeCount(userid, guanxi)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// 新增一条师生某个关系
// @Title GetRelationsByUidSeeCount
// @Description 新增一条师生某个关系
// @Param	sid		path 	string	true		学生主键id
// @Param	tid		path 	string	true		老师主键id
// @Param	guanxi		path 	string	true		关系关键词
// @Success id {int} id
// @Failure 0 添加失败
// @Failure -1 已存在
// @router /AddRelationsBySTGuanxi/:sid/:tid/:guanxi [get]
func (c *RelationsController) AddRelationsBySTGuanxi() {
	sidStr := c.Ctx.Input.Params[":sid"]
	sid, _ := strconv.Atoi(sidStr)
	tidStr := c.Ctx.Input.Params[":tid"]
	tid, _ := strconv.Atoi(tidStr)
	guanxi := c.Ctx.Input.Params[":guanxi"]
	v, _ := models.GetRelationsByST(sid, tid, guanxi)
	if v == nil {
		var addre models.Relations
		addre.FrontUserId = tid
		addre.AfterUserId = sid
		addre.SetDate = time.Now()
		addre.Sources = guanxi
		id, err := models.AddRelations(&addre)
		if id > 0 && err == nil {
			c.Data["json"] = map[string]int64{"id": id} //sucess
		} else {
			c.Data["json"] = map[string]int64{"id": 0} //shibai
		}
	} else {
		c.Data["json"] = map[string]int64{"id": -1} //cunzai
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Relations
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Relations
// @Failure 403
// @router /GetAllRelations/:page/:size [get]
func (c *RelationsController) GetAll() {
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

	l, err := models.GetAllRelations(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Relations
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Relations	true		"body for Relations content"
// @Success 200 {object} models.Relations
// @Failure 403 :id is not int
// @router /UpdateRelationsById/:id [post]
func (c *RelationsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Relations{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateRelationsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Relations
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteRelations/:id [get]
func (c *RelationsController) DeleteRelations() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRelations(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
