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

// oprations for Onlinetrylisten
type OnlinetrylistenController struct {
	beego.Controller
}

func (c *OnlinetrylistenController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Onlinetrylisten
// @Param	body		body 	models.Onlinetrylisten	true		"body for Onlinetrylisten content"
// @Success 200 {int} models.Onlinetrylisten.Id
// @Failure 403 body is empty
// @router /AddOnlinetrylisten/ [post]
func (c *OnlinetrylistenController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Onlinetrylisten
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddOnlinetrylisten(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Post
// @Description create Onlinetrylisten
// @Param	body		body 	models.Onlinetrylisten	true		"body for Onlinetrylisten content"
// @Success 200 {int} models.Onlinetrylisten.Id
// @Failure 403 body is empty
// @router /AddOnlinetrylistenOnline/ [post]
func (c *OnlinetrylistenController) AddOnlinetrylistenOnline() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Onlinetrylisten
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddOnlinetrylisten(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /GetOnlinetrylistenById/:id [get]
func (c *OnlinetrylistenController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinetrylistenById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//41.查询老师的试听信息
// @Title OnlineTryListenByTid
// @Description OnlineTryListenByTid Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /OnlineTryListenByTid/:userid/:page/:size [get]
func (c *OnlinetrylistenController) OnlineTryListenByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.OnlineTryListenByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//41.查询老师的试听信息总条数
// @Title OnlineTryListenByTidCount
// @Description OnlineTryListenByTidCount Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /OnlineTryListenByTidCount/:userid [get]
func (c *OnlinetrylistenController) OnlineTryListenByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.OnlineTryListenByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetOnlinetrylistenOneByTid
// @Description GetOnlinetrylistenOneByTid Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /GetOnlinetrylistenOneByTid/:tid [get]
func (c *OnlinetrylistenController) GetOnlinetrylistenOneByTid() {
	idStr := c.Ctx.Input.Params[":tid"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOnlinetrylistenOneByTid(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Onlinetrylisten
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403
// @router /GetAllOnlinetrylisten/:page/:size [get]
func (c *OnlinetrylistenController) GetAll() {
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

	l, err := models.GetAllOnlinetrylisten(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Onlinetrylisten
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Onlinetrylisten	true		"body for Onlinetrylisten content"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is not int
// @router /UpdateOnlinetrylistenById/:id [get]
func (c *OnlinetrylistenController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Onlinetrylisten{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateOnlinetrylistenById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Onlinetrylisten
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteOnlinetrylisten/:id [get]
func (c *OnlinetrylistenController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOnlinetrylisten(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /GetListenTecher/:listenid [get]
func (c *OnlinetrylistenController) GetOss() {
	idStr := c.Ctx.Input.Params[":listenid"]
	listenid, _ := strconv.Atoi(idStr)
	c.Ctx.SetCookie("onlinelistenid", strconv.Itoa(listenid)) //当前老师进入试听信息主键
	joinurl, err := models.GeListentecherlession2(listenid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Println(joinurl)
		c.Data["json"] = map[string]string{"url": joinurl}
	}
	c.ServeJson()
}

// @Title Get
// @Description get Onlinetrylisten by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Onlinetrylisten
// @Failure 403 :id is empty
// @router /GetListenStudent/:listenid [get]
func (c *OnlinetrylistenController) GetOe() {
	idStr := c.Ctx.Input.Params[":listenid"]
	listenid, _ := strconv.Atoi(idStr)
	c.Ctx.SetCookie("onlinelistenid", strconv.Itoa(listenid))
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	joinurl, err := models.GetListenStudentlession2(listenid, stuuserid) //当前学生进入试听信息主键
	fmt.Println("路径是：")
	fmt.Println(joinurl)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = joinurl
	}
	c.ServeJson()
}
