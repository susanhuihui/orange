package controllers

import (
	"encoding/json"
	"errors"
	"orange/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Onlinecoursebookingrecord
type OnlinecoursebookingrecordController struct {
	beego.Controller
}

func (c *OnlinecoursebookingrecordController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Onlinecoursebookingrecord
// @Param	body		body 	models.Onlinecoursebookingrecord	true		"body for Onlinecoursebookingrecord content"
// @Success 200 {int} models.Onlinecoursebookingrecord.Id
// @Failure 403 body is empty
// @router /AddOnlinecoursebookingrecord/ [post]
func (c *OnlinecoursebookingrecordController) Post() {
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Onlinecoursebookingrecord
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddOnlinecoursebookingrecord(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinecoursebookingrecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403 :id is empty
// @router /GetOnlinecoursebookingrecordById/:id [get]
func (c *OnlinecoursebookingrecordController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinecoursebookingrecordById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetOnlinecoursebookingrecordByUid
// @Description GetOnlinecoursebookingrecordByUid Onlinecoursebookingrecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403 :id is empty
// @router /GetOnlinecoursebookingrecordByUid/:userid/:bookid [get]
func (c *OnlinecoursebookingrecordController) GetOnlinecoursebookingrecordByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)
	v, err := models.GetOnlinecoursebookingrecordByUid(userid, bookid)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetOnlinecoursebookingrecordByTwoid
// @Description GetOnlinecoursebookingrecordByTwoid Onlinecoursebookingrecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403 :id is empty
// @router /GetOnlinecoursebookingrecordByTwoid/:bookid [get]
func (c *OnlinecoursebookingrecordController) GetOnlinecoursebookingrecordByTwoid() {
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)
	v, err := models.GetOnlinecoursebookingrecordByTwoid(bookid)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetOnlinecoursebookingrecordBybookiduid
// @Description GetOnlinecoursebookingrecordBybookiduid Onlinecoursebookingrecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403 :id is empty
// @router /GetOnlinecoursebookingrecordBybookiduid/:userid/:bookid [get]
func (c *OnlinecoursebookingrecordController) GetOnlinecoursebookingrecordBybookiduid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	bookidStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(bookidStr)
	v, err := models.GetOnlinecoursebookingrecordBybookiduid(userid, bookid)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Onlinecoursebookingrecord
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403
// @router /GetAllOnlinecoursebookingrecord/:page/:size [get]
func (c *OnlinecoursebookingrecordController) GetAll() {
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

	l, err := models.GetAllOnlinecoursebookingrecord(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Onlinecoursebookingrecord
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Onlinecoursebookingrecord	true		"body for Onlinecoursebookingrecord content"
// @Success 200 {object} models.Onlinecoursebookingrecord
// @Failure 403 :id is not int
// @router /UpdateOnlinecoursebookingrecordById/:id [post]
func (c *OnlinecoursebookingrecordController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, _ := range c.Ctx.Request.Form {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Onlinecoursebookingrecord{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateOnlinecoursebookingrecordById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Onlinecoursebookingrecord
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteOnlinecoursebookingrecord/:id [get]
func (c *OnlinecoursebookingrecordController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOnlinecoursebookingrecord(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
