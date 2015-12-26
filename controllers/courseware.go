package controllers

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"

	"orange/command"
	"orange/models"
	"strconv"
	"strings"
	"time"
)

// oprations for Courseware
type CoursewareController struct {
	beego.Controller
}

func (c *CoursewareController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Courseware
// @Param	body		body 	models.Courseware	true		"body for Courseware content"
// @Success 200 {int} models.Courseware.Id
// @Failure 403 body is empty
// @router /AddCourseware/ [post]
func (c *CoursewareController) Post() {
	var v models.Courseware
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if id, err := models.AddCourseware(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title AddCoursewareOnbook
// @Description AddCoursewareOnbook the Courseware
// @Param	id		path 	string	true		"The id you want to AddCoursewareOnbook"
// @Param	body		body 	models.Courseware	true		"body for Courseware content"
// @Success 200 {object} models.Courseware
// @Failure 403 :id is not int
// @router /AddCoursewareOnbook/:bookid [post]
func (c *CoursewareController) AddCoursewareOnbook() {
	// 获取上传数据
	stuUserId, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	bookId, _ := strconv.Atoi(c.Ctx.Input.Params[":bookid"])

	mapFilePath, err := command.SaveUploadFiles(&c.Controller, "file_image")
	if nil != err {
		beego.Error("SAVE UPLOAD FILES: ", err.Error())
		c.Data["json"] = map[string]interface{}{"state": -1}
		c.ServeJson()
		return
	}

	// Insert data to DB
	for _, fileSavePath := range mapFilePath {
		addcour := &models.Courseware{
			OCBRId:      bookId,
			UserId:      stuUserId,
			CoursePath:  fileSavePath,
			CourseType:  0,
			AuditStatus: 0,
			UploadTime:  time.Now(),
		}

		if _, err := models.AddCourseware(addcour); nil != err {
			beego.Error("INSERT DATA: ", err.Error())
			c.Data["json"] = map[string]interface{}{"state": -1}
			c.ServeJson()
			return
		}
	}

	// c.Data["json"] = map[string]interface{}{"state": 1}
	c.Ctx.Redirect(302, "http://"+models.OnlineUrl+
		"/orange/Teacher/StudentSetTeacherMeet/1")

	//--------------------------------
	// 2015/12/27
	// 基本完成了预约部分多附件上传的功能, 细节需补充一下几点:
	// 1. 上传预览部分需要去除, 在上传按钮附近提示附件数量;
	// 2. 上传附件的文件类型未做限定, 需要做限定;
	// 3. 要修改重定向功能为模板渲染输出功能, 将错误信息展现在模板上;
	// 4. 为文件保存函数增加文件数量校验功能;
	// 5. 要去读懂页面的JS. 这个难度有些太大了......
}

// @Title Get
// @Description get Courseware by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Courseware
// @Failure 403 :id is empty
// @router /GetCoursewareById/:id [get]
func (c *CoursewareController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCoursewareById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//37.查询预约课程附件信息
// @Title GetCoursewareByOCBID
// @Description GetCoursewareByOCBID Courseware by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Courseware
// @Failure 403 :id is empty
// @router /GetCoursewareByOCBID/:ocbrid [get]
func (c *CoursewareController) GetCoursewareByOCBID() {
	ocbridStr := c.Ctx.Input.Params[":ocbrid"]
	ocbrid, _ := strconv.Atoi(ocbridStr)
	v, err := models.GetCoursewareByOCBID(ocbrid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Courseware
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Courseware
// @Failure 403
// @router /GetAllCourseware/:page/:size [get]
func (c *CoursewareController) GetAll() {
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

	l, err := models.GetAllCourseware(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Courseware
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Courseware	true		"body for Courseware content"
// @Success 200 {object} models.Courseware
// @Failure 403 :id is not int
// @router /UpdateCoursewareById/:id [put]
func (c *CoursewareController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Courseware{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCoursewareById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Courseware
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteCourseware/:id [get]
func (c *CoursewareController) DeleteCourseware() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCourseware(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
