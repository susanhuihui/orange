package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"orange/models"
	"strconv"
	"strings"
	"time"
)

// oprations for Usermessage
type UsermessageController struct {
	beego.Controller
}

func (c *UsermessageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Usermessage
// @Param	body		body 	models.Usermessage	true		"body for Usermessage content"
// @Success 200 {int} models.Usermessage.Id
// @Failure 403 body is empty
// @router /AddUsermessage/ [post]
func (c *UsermessageController) Post() {
	userMessage := models.Usermessage{}
	if err := c.ParseForm(&userMessage); nil != err {
		beego.Error("PARSE FORM: ", err.Error())
		c.Data["json"] = err.Error()
		c.ServeJson()
	}

	userMessage.MessageId = 0
	userMessage.States = 0
	userMessage.MesTime = time.Now()

	if _, err := models.AddUsermessage(&userMessage); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJson()
	}

	// 留言成功刷新页面
	c.Redirect("http://"+models.OnlineUrl+"/orange/Teacher/TeacherMessage/1", 302)
}

// @Title Get
// @Description get Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageById/:id [get]
func (c *UsermessageController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsermessageById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetUsermessageBymuid
// @Description GetUsermessageBymuid Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageBymuid/:mid/:userid [get]
func (c *UsermessageController) GetUsermessageBymuid() {
	midStr := c.Ctx.Input.Params[":mid"]
	mid, _ := strconv.Atoi(midStr)
	useridStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(useridStr)
	v, err := models.GetUsermessageBymuid(mid, userid)
	if err == nil && len(v) > 0 {
		for i := 0; i < len(v); i++ {
			if v[i].States != 1 {
				v[i].States = 1
				err = models.UpdateUsermessageById(&v[i])
			}
		}
	}
	if err != nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = "NO"
	}
	c.ServeJson()
}

//老师使用
// @Title GetUsermessageBymuidft
// @Description GetUsermessageBymuidft Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageBymuidft/:mid/:userid [get]
//func (c *UsermessageController) GetUsermessageBymuidft() {
//	midStr := c.Ctx.Input.Params[":mid"]
//	mid, _ := strconv.Atoi(midStr)
//	useridStr := c.Ctx.Input.Params[":userid"]
//	userid, _ := strconv.Atoi(useridStr)
//	v, err := models.GetUsermessageBymuidft(mid, userid)
//	if err != nil {
//		c.Data["json"] = err.Error()
//	} else {
//		c.Data["json"] = v
//	}
//	c.ServeJson()
//}

// @Title UpdateUsermessageBypiduid
// @Description UpdateUsermessageBypiduid Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /UpdateUsermessageBypiduid/:mid/:userid [get]
func (c *UsermessageController) UpdateUsermessageBypiduid() {
	midStr := c.Ctx.Input.Params[":mid"]
	mid, _ := strconv.Atoi(midStr)
	useridStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(useridStr)
	v, err := models.UpdateUsermessageBypiduid(mid, userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//7老师看学生留言
// @Title GetUsermessageByTid
// @Description GetUsermessageByTid Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageByTid/:userid/:page/:size [get]
func (c *UsermessageController) GetUsermessageByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetUsermessageByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//7老师看学生留言总条数
// @Title GetUsermessageByTidCount
// @Description GetUsermessageByTidCount Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageByTidCount/:userid [get]
func (c *UsermessageController) GetUsermessageByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetUsermessageByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//25.学生查看自己的全部留言信息
// @Title GetUsermessageBySid
// @Description GetUsermessageBySid Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageBySid/:userid/:page/:size [get]
func (c *UsermessageController) GetUsermessageBySid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetUsermessageBySid(userid, offset, limit)

	beego.Debug(v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//25.学生查看自己的全部留言信息总条数
// @Title GetUsermessageBySidCount
// @Description GetUsermessageBySidCount Usermessage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is empty
// @router /GetUsermessageBySidCount/:userid [get]
func (c *UsermessageController) GetUsermessageBySidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetUsermessageBySidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Usermessage
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usermessage
// @Failure 403
// @router /GetAllUsermessage/:page/:size [get]
func (c *UsermessageController) GetAll() {
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

	l, err := models.GetAllUsermessage(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Usermessage
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usermessage	true		"body for Usermessage content"
// @Success 200 {object} models.Usermessage
// @Failure 403 :id is not int
// @router /UpdateUsermessageById/:id [post]
func (c *UsermessageController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println("更新留言变已看")
	fmt.Println(jsonS)
	v := models.Usermessage{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println(v)
	if err := models.UpdateUsermessageById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Usermessage
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteUsermessage/:id [get]
func (c *UsermessageController) DeleteUsermessage() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsermessage(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
