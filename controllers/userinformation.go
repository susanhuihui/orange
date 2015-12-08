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

// oprations for Userinformation
type UserinformationController struct {
	beego.Controller
}

func (c *UserinformationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Userinformation
// @Param	body		body 	models.Userinformation	true		"body for Userinformation content"
// @Success 200 {int} models.Userinformation.Id
// @Failure 403 body is empty
// @router /AddUserinformation/ [post]
func (c *UserinformationController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Userinformation
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println(v)
	if id, err := models.AddUserinformation(&v); err == nil {
		//添加账户信息
		var userid string = strconv.FormatInt(int64(id), 10)
		addid, _ := strconv.Atoi(userid)
		var useraccount models.Accountfunds
		useraccount.UserId = addid

		useraccount.Balance = 0
		useraccount.FundState = 0
		useraccount.OpenTime = time.Now()
		useraccount.AccountTypeId = 0
		idaccount, errac := models.AddAccountfunds(&useraccount)
		if errac == nil {
			c.Data["json"] = map[string]int64{"id": idaccount}
		}
		//c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = map[string]int64{"id": 0}
	}
	c.ServeJson()
}

// @Title Get
// @Description get Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationById/:id [get]
func (c *UserinformationController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserinformationById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetUserinformationByPhone
// @Description GetUserinformationByPhone Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationByPhone/:phone [get]
func (c *UserinformationController) GetUserinformationByPhone() {
	phoneStr := c.Ctx.Input.Params[":phone"]
	fmt.Println(phoneStr)
	v, err := models.GetUserinformationByPhone(phoneStr)
	if err != nil && v == nil {
		c.Data["json"] = map[string]int64{"result": 0}
	} else {
		c.Data["json"] = map[string]int64{"result": 1}
	}
	c.ServeJson()
}

// @Title GetUserinformationPhone
// @Description GetUserinformationPhone Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationPhone/:phone [get]
func (c *UserinformationController) GetUserinformationPhone() {
	phoneStr := c.Ctx.Input.Params[":phone"]
	fmt.Println(phoneStr)
	v, err := models.GetUserinformationByPhone(phoneStr)
	if err != nil && v == nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//根据用户验证密码是否正确(仅用于修改密码时调用，登录时不可调用)
// @Title GetUserinformationByIdPass
// @Description GetUserinformationByIdPass Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationByIdPass/:id/:pass [get]
func (c *UserinformationController) GetUserinformationByIdPass() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	passStr := c.Ctx.Input.Params[":pass"]
	v, err := models.GetUserinformationById(id)
	var reslut string = ""
	if v.LoginPassword == passStr {
		reslut = "true"
	} else {
		reslut = "false"
	}
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = reslut
	}
	c.ServeJson()
}

// @Title Get
// @Description get Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationPicMove/:count [get]
func (c *UserinformationController) GetUserinformationPicMove() {
	idStr := c.Ctx.Input.Params[":count"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserinformationPicMove(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetUserinformationAllTeacher
// @Description GetUserinformationAllTeacher Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationAllTeacher/:seltype/:nianji/:kecheng/:jibie/:shengfen/:shiqu/:page/:size [get]
func (c *UserinformationController) GetUserinformationAllTeacher() {
	seltypestr := c.Ctx.Input.Params[":seltype"]
	newnianji := ""
	newkecheng := ""
	newjibie := ""
	newshengfen := ""
	newshiqu := ""
	nianji := c.Ctx.Input.Params[":nianji"]
	fmt.Println(nianji)
	if nianji != "" {
		newnianji = nianji
	} else if nianji == "" {
		newnianji = `%%`
	}
	kecheng := c.Ctx.Input.Params[":kecheng"]
	if kecheng != "" {
		newkecheng = kecheng
	} else if kecheng == "" {
		newkecheng = `%%`
	}
	jibie := c.Ctx.Input.Params[":jibie"]
	if jibie != "" {
		newjibie = jibie
	} else if jibie == "" {
		newjibie = `%%`
	}
	shengfen := c.Ctx.Input.Params[":shengfen"]
	if shengfen != "" {
		newshengfen = shengfen
	} else if shengfen == "" {
		newshengfen = `%%`
	}
	shiqu := c.Ctx.Input.Params[":shiqu"]
	if shiqu != "" {
		newshiqu = shiqu
	} else if shiqu == "" {
		newshiqu = `%%`
	}
	seltype, _ := strconv.Atoi(seltypestr)
	fmt.Println(newjibie + "," + newkecheng + "," + newnianji + "," + newshengfen + "," + newshiqu)

	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetUserinformationAllTeacher(seltype, newnianji, newkecheng, newjibie, newshengfen, newshiqu, offset, limit)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetUserinformationAllTeacherCount
// @Description GetUserinformationAllTeacherCount Userinformation by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is empty
// @router /GetUserinformationAllTeacherCount/:seltype/:nianji/:kecheng/:jibie/:shengfen/:shiqu [get]
func (c *UserinformationController) GetUserinformationAllTeacherCount() {
	seltypestr := c.Ctx.Input.Params[":seltype"]
	newnianji := ""
	newkecheng := ""
	newjibie := ""
	newshengfen := ""
	newshiqu := ""
	nianji := c.Ctx.Input.Params[":nianji"]
	if nianji != "" {
		newnianji = nianji
	} else if nianji == "" {
		newnianji = `%%`
	}
	kecheng := c.Ctx.Input.Params[":kecheng"]
	if kecheng != "" {
		newkecheng = kecheng
	} else if kecheng == "" {
		newkecheng = `%%`
	}
	jibie := c.Ctx.Input.Params[":jibie"]
	if jibie != "" {
		newjibie = jibie
	} else if jibie == "" {
		newjibie = `%%`
	}
	shengfen := c.Ctx.Input.Params[":shengfen"]
	if shengfen != "" {
		newshengfen = shengfen
	} else if shengfen == "" {
		newshengfen = `%%`
	}
	shiqu := c.Ctx.Input.Params[":shiqu"]
	if shiqu != "" {
		newshiqu = shiqu
	} else if shiqu == "" {
		newshiqu = `%%`
	}
	seltype, _ := strconv.Atoi(seltypestr)
	fmt.Println(newjibie + "," + newkecheng + "," + newnianji + "," + newshengfen + "," + newshiqu)

	v, err := models.GetUserinformationAllTeacherCount(seltype, newnianji, newkecheng, newjibie, newshengfen, newshiqu)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Userinformation
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Userinformation
// @Failure 403
// @router /GetAllUserinformation/:page/:size [get]
func (c *UserinformationController) GetAll() {
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

	l, err := models.GetAllUserinformation(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Userinformation
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Userinformation	true		"body for Userinformation content"
// @Success 200 {object} models.Userinformation
// @Failure 403 :id is not int
// @router /UpdateUserinformationById/:id [post]
func (c *UserinformationController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)

	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Userinformation{Id: id}
	json.Unmarshal([]byte(jsonS), &v)

	if err := models.UpdateUserinformationById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Userinformation
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteUserinformation/:id [delete]
func (c *UserinformationController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUserinformation(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Update
// @Description update the TbUser
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TbUser	true		"body for TbUser content"
// @Success 200 {object} models.TbUser
// @Failure 403 :id is not int
// @router /UpdateUserimg/:userid [post]
func (c *UserinformationController) UpdateUserimg() {
	request := c.Ctx.Request
	fmt.Println("是否调到")
	jsons, imgstr := models.GetImganddata2(request, Headurl)
	var v models.Userinformation
	json.Unmarshal([]byte(jsons), &v)
	fmt.Println(jsons)
	//保存用户头像
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	fmt.Println(imgstr)
	if imgstr != "" {
		updateuser, _ := models.GetUserinformationById(userid)
		updateuser.AvatarPath = imgstr
		upresulterr := models.UpdateUserinformationById(updateuser)
		if upresulterr == nil {
			c.Data["json"] = map[string]interface{}{"state": 1} //修改成功
		} else {
			c.Data["json"] = map[string]interface{}{"state": 0} //修改失败
		}
	} else {
		c.Data["json"] = map[string]interface{}{"state": -1} //上传失败
	}
	c.ServeJson()
}
