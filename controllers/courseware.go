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
	request := c.Ctx.Request

	jsons, imgstr := models.GetImganddata2(request, Fujianurl)
	var v models.Userinformation
	json.Unmarshal([]byte(jsons), &v)
	fmt.Println(jsons)
	//新增图片到数据库
	stuuserid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))
	idStr := c.Ctx.Input.Params[":bookid"]
	bookid, _ := strconv.Atoi(idStr) //预约课程主键id
	fmt.Println(imgstr)              //所有图片路径集合
	var adderr error
	if imgstr != "" {
		// sellist = 文件路径
		sellist := strings.Split(imgstr, ",")
		for i := 0; i < len(sellist); i++ {
			fmt.Println(sellist[i])
			var addcour models.Courseware
			addcour.OCBRId = bookid
			addcour.UserId = stuuserid
			addcour.CoursePath = sellist[i]
			addcour.CourseType = 0
			addcour.AuditStatus = 0
			addcour.UploadTime = time.Now()
			addid, err := models.AddCourseware(&addcour)
			if err != nil && addid > 0 {
				adderr = err
			}
		}
		if adderr == nil {
			c.Data["json"] = map[string]interface{}{"state": 1} //添加成功
		} else {
			c.Data["json"] = map[string]interface{}{"state": 0} //添加失败
		}
	} else {
		c.Data["json"] = map[string]interface{}{"state": -1} //上传失败
	}
	c.ServeJson()
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
