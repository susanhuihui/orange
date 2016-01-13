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

// oprations for Verification
type VerificationController struct {
	beego.Controller
}

func (c *VerificationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Verification
// @Param	body		body 	models.Verification	true		"body for Verification content"
// @Success 200 {int} models.Verification.Id
// @Failure 403 body is empty
// @router /AddVerification/ [post]
func (c *VerificationController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Verification
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddVerification(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Verification by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Verification
// @Failure 403 :id is empty
// @router /GetVerificationById/:id [get]
func (c *VerificationController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetVerificationById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

/**36.根据手机号码获取一条最新的验证码信息**/
// @Title GetVerificationByPhone
// @Description 根据手机号码获取一条最新的验证码信息
// @Param	phone		path 	string	true		手机号
// @Success 200 {object} models.Verification
// @Failure 403 Error
// @router /GetVerificationByPhone/:phone [get]
func (c *VerificationController) GetVerificationByPhone() {
	phoneStr := c.Ctx.Input.Params[":phone"]
	v, err := models.GetVerificationByPhone(phoneStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetVerificationListByPhone
// @Description 查询手机号下的所有验证码信息
// @Param	phone		path 	string	true		手机号码
// @Success 200 {object} models.Verification
// @Failure 403 Error
// @router /GetVerificationListByPhone/:phone [get]
func (c *VerificationController) GetVerificationListByPhone() {
	phoneStr := c.Ctx.Input.Params[":phone"] //查询手机号下的所有验证码信息
	v, err := models.GetVerificationListByPhone(phoneStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Verification
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Verification
// @Failure 403
// @router /GetAllVerification/:page/:size [get]
func (c *VerificationController) GetAll() {
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

	l, err := models.GetAllVerification(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Verification
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Verification	true		"body for Verification content"
// @Success 200 {object} models.Verification
// @Failure 403 :id is not int
// @router /UpdateVerificationById/:id [post]
func (c *VerificationController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Verification{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateVerificationById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Verification
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteVerification/:id [get]
func (c *VerificationController) DeleteVerification() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteVerification(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title DeleteVerificationByPhone
// @Description 根据手机号码删除其下的全部验证码信息
// @Param	phone		path 	string	true		手机号码
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteVerificationByPhone/:phone [get]
func (c *VerificationController) DeleteVerificationByPhone() {
	phoneStr := c.Ctx.Input.Params[":phone"]
	var dellisterr error
	phonever, geterr := models.GetVerificationListByPhone(phoneStr)
	if geterr == nil {
		for i := 0; i < len(phonever); i++ {
			delerr := models.DeleteVerification(phonever[i].Id)
			if delerr != nil {
				dellisterr = delerr
			}
		}
	}
	if dellisterr == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = dellisterr.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description 给手机号码发送验证码
// @Param	phone		path 	string	true		手机号
// @Success 200 {string} 0
// @Failure 403 1
// @router /GetVerification/:phone [get]
func (c *VerificationController) GetOnees() {
	idStr := c.Ctx.Input.Params[":phone"]

	err := models.GetVerifications(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"state": "0"}
	} else {
		c.Data["json"] = map[string]string{"state": "1"}
	}
	c.ServeJson()
}
