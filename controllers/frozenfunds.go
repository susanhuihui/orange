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

// oprations for Frozenfunds
type FrozenfundsController struct {
	beego.Controller
}

func (c *FrozenfundsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Frozenfunds
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {int} models.Frozenfunds.Id
// @Failure 403 body is empty
// @router /AddFrozenfunds/ [post]
func (c *FrozenfundsController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Frozenfunds
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddFrozenfunds(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

//冻结账户资金同时，从账户余额中扣除冻结资金
// @Title Post
// @Description create Frozenfunds
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {int} models.Frozenfunds.Id
// @Failure 403 body is empty
// @router /AddUserFrozenfunds/ [post]
func (c *FrozenfundsController) AddUserFrozenfunds() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Frozenfunds
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddFrozenfunds(&v); err == nil {
		useraccount, _ := models.GetAccountfundsByuid(v.UserId)
		useraccount.Balance = useraccount.Balance - v.FrozenMoney
		err := models.UpdateAccountfundsById(&useraccount)
		if err == nil {
			c.Data["json"] = map[string]int64{"id": id}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Frozenfunds
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is not int
// @router /UpdateUserFrozenfundsById/:id [post]
func (c *FrozenfundsController) UpdateUserFrozenfundsById() {
	idStr := c.Ctx.Input.Params[":id"] //取消预约时调用的方法
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println([]byte(jsonS))
	fmt.Println(id)
	v := models.Frozenfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println("修改为")
	fmt.Println(&v)
	if err := models.UpdateFrozenfundsById(&v); err == nil { //修改冻结资金信息（解冻资金）
		useraccount, _ := models.GetAccountfundsByuid(v.UserId)
		useraccount.Balance = useraccount.Balance + v.FrozenMoney //退还预约人的资金
		err := models.UpdateAccountfundsById(&useraccount)        //保存预约人账户信息
		if err == nil {
			c.Data["json"] = "OK"
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Frozenfunds
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is not int
// @router /UpdateUserFrozenfundsByAnswer/:id [post]
func (c *FrozenfundsController) UpdateUserFrozenfundsByAnswer() {
	idStr := c.Ctx.Input.Params[":id"] //回答问题后调用的方法
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println([]byte(jsonS))
	fmt.Println(id)
	v := models.Frozenfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println("修改为")
	fmt.Println(&v)
	if err := models.UpdateFrozenfundsById(&v); err == nil { //修改冻结资金信息（解冻资金）
		teacherid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))   //获取老师主键
		useraccount, _ := models.GetAccountfundsByuid(teacherid)  //获取老师账户信息
		useraccount.Balance = useraccount.Balance + v.FrozenMoney //将解冻钱打给老师
		err := models.UpdateAccountfundsById(&useraccount)        //保存老师账户信息
		var jiaoyijilu models.Transactionrecords
		jiaoyijilu.SendUserId = v.UserId
		jiaoyijilu.CollectUserId = teacherid
		jiaoyijilu.RecordMoney = v.FrozenMoney
		jiaoyijilu.TradingWayId = 1
		jiaoyijilu.RecordTime = time.Now()
		addid, _ := models.AddTransactionrecords(&jiaoyijilu) //添加一条交易记录
		if err == nil && addid > 0 {
			c.Data["json"] = "OK"
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title UpdateUserFrozenfundsByOnline
// @Description UpdateUserFrozenfundsByOnline the Frozenfunds
// @Param	id		path 	string	true		"The id you want to UpdateUserFrozenfundsByOnline"
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is not int
// @router /UpdateUserFrozenfundsByOnline/:id/:tid [post]
func (c *FrozenfundsController) UpdateUserFrozenfundsByOnline() {
	idStr := c.Ctx.Input.Params[":id"] //上完课后调用的方法
	id, _ := strconv.Atoi(idStr)
	tidStr := c.Ctx.Input.Params[":tid"] //老师id
	tid, _ := strconv.Atoi(tidStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println([]byte(jsonS))
	fmt.Println(id)
	v := models.Frozenfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println("修改为")
	fmt.Println(&v)
	if err := models.UpdateFrozenfundsById(&v); err == nil { //修改冻结资金信息（解冻资金）
		//teacherid, _ := strconv.Atoi(c.Ctx.GetCookie("userid"))   //获取老师主键
		useraccount, _ := models.GetAccountfundsByuid(tid)        //获取老师账户信息
		useraccount.Balance = useraccount.Balance + v.FrozenMoney //将解冻钱打给老师
		err := models.UpdateAccountfundsById(&useraccount)        //保存老师账户信息
		var jiaoyijilu models.Transactionrecords
		jiaoyijilu.SendUserId = v.UserId
		jiaoyijilu.CollectUserId = tid
		jiaoyijilu.RecordMoney = v.FrozenMoney
		jiaoyijilu.TradingWayId = 1
		jiaoyijilu.RecordTime = time.Now()
		addid, _ := models.AddTransactionrecords(&jiaoyijilu) //添加一条交易记录
		if err == nil && addid > 0 {
			c.Data["json"] = "OK"
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Frozenfunds
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is not int
// @router /UpdateTeacherFrozenfundsById/:id/tid [post]
func (c *FrozenfundsController) UpdateTeacherFrozenfundsById() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	tidStr := c.Ctx.Input.Params[":tid"]
	tid, _ := strconv.Atoi(tidStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println([]byte(jsonS))
	fmt.Println(id)
	v := models.Frozenfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	fmt.Println("修改为")
	fmt.Println(&v)
	if err := models.UpdateFrozenfundsById(&v); err == nil {
		useraccount, _ := models.GetAccountfundsByuid(tid)
		useraccount.Balance = useraccount.Balance + v.FrozenMoney
		err := models.UpdateAccountfundsById(&useraccount)
		if err == nil {
			c.Data["json"] = "OK"
		}
		var jiaoyijilu models.Transactionrecords
		jiaoyijilu.SendUserId = v.UserId
		jiaoyijilu.CollectUserId = tid
		jiaoyijilu.RecordMoney = v.FrozenMoney
		jiaoyijilu.TradingWayId = 1
		jiaoyijilu.RecordTime = time.Now()
		id, adderr := models.AddTransactionrecords(&jiaoyijilu)
		if id >= 0 && adderr == nil {
			c.Data["json"] = "OK"
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Frozenfunds by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is empty
// @router /GetFrozenfundsById/:id [get]
func (c *FrozenfundsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetFrozenfundsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetFrozenFundsByUserid
// @Description GetFrozenFundsByUserid Frozenfunds by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is empty
// @router /GetFrozenFundsByUserid/:userid [get]
func (c *FrozenfundsController) GetFrozenFundsByUserid() {
	idStr := c.Ctx.Input.Params[":userid"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetFrozenFundsByUserid(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetFrozenfundsByUidOnId
// @Description GetFrozenfundsByUidOnId Frozenfunds by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is empty
// @router /GetFrozenfundsByUidOnId/:userid/:typeid/:selid [get]
func (c *FrozenfundsController) GetFrozenfundsByUidOnId() {
	idStr := c.Ctx.Input.Params[":userid"]
	id, _ := strconv.Atoi(idStr)
	typeidStr := c.Ctx.Input.Params[":typeid"]
	typeid, _ := strconv.Atoi(typeidStr)
	selIdStr := c.Ctx.Input.Params[":selid"]
	selid, _ := strconv.Atoi(selIdStr)
	v, err := models.GetFrozenfundsByUidOnId(id, typeid, selid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Frozenfunds
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Frozenfunds
// @Failure 403
// @router /GetAllFrozenfunds/:page/:size [get]
func (c *FrozenfundsController) GetAll() {
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

	l, err := models.GetAllFrozenfunds(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Frozenfunds
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Frozenfunds	true		"body for Frozenfunds content"
// @Success 200 {object} models.Frozenfunds
// @Failure 403 :id is not int
// @router /UpdateFrozenfundsById/:id [put]
func (c *FrozenfundsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Frozenfunds{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateFrozenfundsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Frozenfunds
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteFrozenfunds/:id [delete]
func (c *FrozenfundsController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteFrozenfunds(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
