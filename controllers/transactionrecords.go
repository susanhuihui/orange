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

// oprations for Transactionrecords
type TransactionrecordsController struct {
	beego.Controller
}

func (c *TransactionrecordsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Transactionrecords
// @Param	body		body 	models.Transactionrecords	true		"body for Transactionrecords content"
// @Success 200 {int} models.Transactionrecords.Id
// @Failure 403 body is empty
// @router /AddTransactionrecords/ [post]
func (c *TransactionrecordsController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Transactionrecords
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddTransactionrecords(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Transactionrecords by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Transactionrecords
// @Failure 403 :id is empty
// @router /GetTransactionrecordsById/:id [get]
func (c *TransactionrecordsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTransactionrecordsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//14查询用户交易记录
// @Title GetTransactionrecordsByTid
// @Description 查询用户交易记录
// @Param	userid		path 	string	true		用户主键id
// @Param	page		path 	string	true		获取页数
// @Param	size		path 	string	true		获取行数
// @Success 200 {object} models.TransactionrecordsUserList
// @Failure Error
// @router /GetTransactionrecordsByTid/:userid/:page/:size [get]
func (c *TransactionrecordsController) GetTransactionrecordsByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetTransactionrecordsByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//14查询用户交易记录
// @Title GetTransactionrecordsByTidCount
// @Description 查询用户交易记录总条数
// @Param	userid		path 	string	true		用户主键id
// @Success json {int} json
// @Failure Error
// @router /GetTransactionrecordsByTidCount/:userid [get]
func (c *TransactionrecordsController) GetTransactionrecordsByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetTransactionrecordsByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//23.查询学生消费记录
// @Title GetTransactionrecordsBySid
// @Description 查询学生消费记录
// @Param	userid		path 	string	true		用户主键id
// @Param	page		path 	string	true		获取页数
// @Param	size		path 	string	true		获取行数
// @Success 200 {object} models.TransactionrecordsUserList
// @Failure Error
// @router /GetTransactionrecordsBySid/:userid/:page/:size [get]
func (c *TransactionrecordsController) GetTransactionrecordsBySid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetTransactionrecordsBySid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//23.查询学生消费记录总条数
// @Title GetTransactionrecordsBySidCount
// @Description 查询学生消费记录总条数
// @Param	userid		path 	string	true		用户主键id
// @Success json {int} json
// @Failure Error
// @router /GetTransactionrecordsBySidCount/:userid [get]
func (c *TransactionrecordsController) GetTransactionrecordsBySidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetTransactionrecordsBySidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Transactionrecords
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Transactionrecords
// @Failure 403
// @router /GetAllTransactionrecords/:page/:size [get]
func (c *TransactionrecordsController) GetAll() {
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

	l, err := models.GetAllTransactionrecords(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Transactionrecords
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Transactionrecords	true		"body for Transactionrecords content"
// @Success 200 {object} models.Transactionrecords
// @Failure 403 :id is not int
// @router /UpdateTransactionrecordsById/:id [post]
func (c *TransactionrecordsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Transactionrecords{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateTransactionrecordsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Transactionrecords
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteTransactionrecords/:id [delete]
func (c *TransactionrecordsController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTransactionrecords(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
