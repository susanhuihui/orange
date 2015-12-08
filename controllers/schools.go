package controllers

import (
	"encoding/json"
	"errors"
	"orange/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Schools
type SchoolsController struct {
	beego.Controller
}

func (c *SchoolsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Schools
// @Param	body		body 	models.Schools	true		"body for Schools content"
// @Success 200 {int} models.Schools.Id
// @Failure 403 body is empty
// @router /AddSchools/ [post]
func (c *SchoolsController) Post() {
	var v models.Schools
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if id, err := models.AddSchools(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Schools by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Schools
// @Failure 403 :id is empty
// @router /GetSchoolsById/:id [get]
func (c *SchoolsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSchoolsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetSchoolsByCity
// @Description GetSchoolsByCity Countys by pid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Countys
// @Failure 403 :id is empty
// @router /GetSchoolsByCity/:cid/:typeid [get]
func (c *SchoolsController) GetSchoolsByCity() {
	idStr := c.Ctx.Input.Params[":cid"]
	cid, _ := strconv.Atoi(idStr)
	typeidStr := c.Ctx.Input.Params[":typeid"]
	typeid, _ := strconv.Atoi(typeidStr)
	v, err := models.GetSchoolsByCity(cid,typeid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Schools
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Schools
// @Failure 403
// @router /GetAllSchools/:page/:size [get]
func (c *SchoolsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)

	page := c.Ctx.Input.Param(":page")		//获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size")		//获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.ParseInt(page, 0, 0)//传来的页数
	rows, _ := strconv.ParseInt(size, 0, 0) //传来的显示行数
	truepages := (pages - 1) * rows         //计算舍弃多少行
	limit := rows                           //显示行数
	offset := truepages                     //舍弃行数	//新加--------结束--------

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

	l, err := models.GetAllSchools(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Schools
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Schools	true		"body for Schools content"
// @Success 200 {object} models.Schools
// @Failure 403 :id is not int
// @router /UpdateSchoolsById/:id [put]
func (c *SchoolsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Schools{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateSchoolsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Schools
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteSchools/:id [delete]
func (c *SchoolsController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSchools(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
