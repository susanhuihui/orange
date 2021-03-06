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

// oprations for Amountrecords
type AmountrecordsController struct {
	beego.Controller
}

func (c *AmountrecordsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Amountrecords
// @Param	body		body 	models.Amountrecords	true		"body for Amountrecords content"
// @Success 200 {int} models.Amountrecords.Id
// @Failure 403 body is empty
// @router /AddAmountrecords/ [post]
func (c *AmountrecordsController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	var v models.Amountrecords
	json.Unmarshal([]byte(jsonS), &v)
	if id, err := models.AddAmountrecords(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Post
// @Description 新增一条提现记录
// @Param Amountrecords form models.Amountrecords true 提现记录实体
// @Success 200 {int} models.Amountrecords.Id
// @Failure 403 body is empty
// @router /AddAmountrecordsStudent/ [post]
func (c *AmountrecordsController) Poststu() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	fmt.Println("申请提现：")
	fmt.Println(jsonS)
	var v models.Amountrecords
	json.Unmarshal([]byte(jsonS), &v)
	v.Balance = v.Balance - v.RecordMoney
	if id, err := models.AddAmountrecords(&v); err == nil {
		//添加冻结学生中户要提现的资金
		//修改学生账户资金余额
		var fontfonz models.Frozenfunds
		fontfonz.UserId = v.UserId
		fontfonz.FrozenMoney = v.RecordMoney
		fontfonz.FrozenType = 2 //提现
		var text = strconv.FormatInt(int64(id), 10)
		fontfonz.BusinessId, _ = strconv.Atoi(text)
		fontfonz.FrozenTime = time.Now()
		fontfonz.FrozenState = 1 //1是已冻结，0否未冻结
		addf, aerr := models.AddFrozenfunds(&fontfonz)
		fmt.Println(addf)
		if aerr == nil {
			useracc, _ := models.GetAccountfundsByuid(v.UserId) //用户账户信息
			useracc.Balance = useracc.Balance - v.RecordMoney
			upacerr := models.UpdateAccountfundsById(&useracc)
			fmt.Println(upacerr)
		}

		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Amountrecords by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Amountrecords
// @Failure 403 :id is empty
// @router /GetAmountrecordsById/:id [get]
func (c *AmountrecordsController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAmountrecordsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsByUserid
// @Description 查询用户（提现/充值）记录状态都为1的情况
// @Param	recordtype query int true 充值提现类型
// @Param	userid query int true 用户主键id
// @Param	page query int true 获取页码数
// @Param	size query int true 获取行数
// @Success 200 {object} models.AmountrecordsUserList
// @Failure 403 :id is empty
// @router /GetAmountrecordsByUserid/:recordtype/:userid/:page/:size [get]
func (c *AmountrecordsController) GetAmountrecordsByUserid() {
	recordtypestr := c.Ctx.Input.Params[":recordtype"]
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	recordtype, _ := strconv.Atoi(recordtypestr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetAmountrecordsByUserid(recordtype, userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsByUseridCount
// @Description 查询用户（提现/充值）记录总条数
// @Param	recordtype query int true 充值提现类型
// @Param	userid query int true 用户主键id
// @Success 200 {object} models.Amountrecords
// @Failure 403 :id is empty
// @router /GetAmountrecordsByUseridCount/:recordtype/:userid [get]
func (c *AmountrecordsController) GetAmountrecordsByUseridCount() {
	recordtypestr := c.Ctx.Input.Params[":recordtype"]
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	recordtype, _ := strconv.Atoi(recordtypestr)
	v, err := models.GetAmountrecordsByUseridCount(recordtype, userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsTixianByUserid
// @Description 查询用户（提现recordtype = 1）全部提现记录
// @Param	userid query int true 用户主键id
// @Param	page query int true 获取页码数
// @Param	size query int true 获取行数
// @Success 200 {object} models.AmountrecordsUserList
// @Failure 403 :id is empty
// @router /GetAmountrecordsTixianByUserid/:userid/:page/:size [get]
func (c *AmountrecordsController) GetAmountrecordsTixianByUserid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetAmountrecordsTixianByUserid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsTixianByUseridCount
// @Description 查询用户（提现recordtype = 1）全部提现记录
// @Param	userid query int true 用户主键id
// @Success 200 {object} models.Amountrecords
// @Failure 403 :id is empty
// @router /GetAmountrecordsTixianByUseridCount/:userid [get]
func (c *AmountrecordsController) GetAmountrecordsTixianByUseridCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetAmountrecordsTixianByUseridCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsTMcountByUid
// @Description 查询用户正在提现的全部金额，继续提现的时候根据此值判断是否可继续提现
// @Param	userid path int true 用户主键id
// @Success 200 {object} models.Amountrecords
// @Failure 403 :id is empty
// @router /GetAmountrecordsTMcountByUid/:userid [get]
func (c *AmountrecordsController) GetAmountrecordsTMcountByUid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetAmountrecordsTMcountByUid(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsAllT
// @Description 管理员查询全部用户正在申请的（提现recordtype = 1）全部提现记录
// @Param	page path int true 获取页码数
// @Param	size path int true 获取行数
// @Success 200 {object} models.AmountrecordsUserAllT
// @Failure 403 :id is empty
// @router /GetAmountrecordsAllT/:page/:size [get]
func (c *AmountrecordsController) GetAmountrecordsAllT() {
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetAmountrecordsAllT(offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//15.
// @Title GetAmountrecordsAllTCount
// @Description 管理员查询全部用户正在申请的（提现recordtype = 1）全部提现记录总条数
// @Success 200 {object} models.Amountrecords
// @Failure Error
// @router /GetAmountrecordsAllTCount/ [get]
func (c *AmountrecordsController) GetAmountrecordsAllTCount() {
	v, err := models.GetAmountrecordsAllTCount()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Amountrecords
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Amountrecords
// @Failure 403
// @router /GetAllAmountrecords/:page/:size [get]
func (c *AmountrecordsController) GetAll() {
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

	l, err := models.GetAllAmountrecords(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Amountrecords
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Amountrecords	true		"body for Amountrecords content"
// @Success 200 {object} models.Amountrecords
// @Failure 403 :id is not int
// @router /UpdateAmountrecordsById/:id [post]
func (c *AmountrecordsController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Amountrecords{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateAmountrecordsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Update
// @Description 给用户发放自己资金
// @Param	id		path 	int	true		提现主键id
// @Param	identityid		path 	int	true		用户身份id
// @Success OK {string} success!
// @Failure error
// @router /FaFang/:id/:identityid [post]
func (c *AmountrecordsController) FaFang() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr) //提现信息主键
	identityidStr := c.Ctx.Input.Params[":identityid"]
	identityid, _ := strconv.Atoi(identityidStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Amountrecords{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	//查询当前用户账户信息
	account, _ := models.GetAccountfundsByuid(v.UserId)

	fmt.Println("当前身份和金额")
	fmt.Println(identityid)
	fmt.Println(v.Balance)

	tixianmoney := v.RecordMoney //操作金额
	if identityid == 1 {         //老师发放需用当前账户余额减去提现金额，更新老师的余额展示，学生余额为冻结后的余额 不需要
		v.Balance = account.Balance - v.RecordMoney //更新提现记录余额
	}

	fmt.Println(v.Balance)
	if err := models.UpdateAmountrecordsById(&v); err == nil { //修改提现记录为
		if identityid == 1 { //给老师 发放金额：老师账户余额需要更新
			account.Balance = account.Balance - tixianmoney
			upacc := models.UpdateAccountfundsById(&account)
			fmt.Println(upacc)
		} else if identityid >= 2 && identityid <= 3 { //学生无需更新账户余额，金额已经从余额中扣除
			//解冻资金，
			fontfonz, _ := models.GetFrozenfundsByUidOnId(v.UserId, 2, v.Id)
			fontfonz.ThawingTime = time.Now()
			fontfonz.FrozenState = 0
			fonterr := models.UpdateFrozenfundsById(&fontfonz)
			fmt.Println(fonterr)
		}
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the Amountrecords
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteAmountrecords/:id [get]
func (c *AmountrecordsController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAmountrecords(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
