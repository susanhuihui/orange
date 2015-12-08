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

// oprations for Accountfunds
type AccountfundsController struct {
	beego.Controller
}

func (c *AccountfundsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Accountfunds
// @Param	body		body 	models.Accountfunds	true		"body for Accountfunds content"
// @Success 200 {int} models.Accountfunds.Id
// @Failure 403 body is empty
// @router /AddAccountfunds/ [post]
func (c *AccountfundsController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Accountfunds
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddAccountfunds(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title GetAccountfundsByuid
// @Description GetAccountfundsByuid Accountfunds by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accountfunds
// @Failure 403 :id is empty
// @router /GetAccountfundsByuid/:userid [get]
func (c *AccountfundsController) GetAccountfundsByuid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetAccountfundsByuid(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get
// @Description get Accountfunds by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accountfunds
// @Failure 403 :id is empty
// @router /GetAccountfundsById/:id [get]
func (c *AccountfundsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAccountfundsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Accountfunds
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Accountfunds
// @Failure 403
// @router /GetAllAccountfunds/:page/:size [get]
func (c *AccountfundsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	//var limit int64 = 10
	//var offset int64 = 0

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
	vtext := "OpenTime"
	// sortby: col1,col2
	//if v := c.GetString("sortby"); v != "" {
	sortby = strings.Split(vtext, ",")
	//}
	vse := "desc"
	// order: desc,asc
	//if v := c.GetString("order"); v != "" {
	order = strings.Split(vse, ",")
	//}
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

	l, err := models.GetAllAccountfunds(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Accountfunds
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Accountfunds	true		"body for Accountfunds content"
// @Success 200 {object} models.Accountfunds
// @Failure 403 :id is not int
// @router /UpdateAccountfundsById/:id [post]
func (c *AccountfundsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Accountfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateAccountfundsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Accountfunds
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteAccountfunds/:id [get]
func (c *AccountfundsController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAccountfunds(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
