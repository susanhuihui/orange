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

// oprations for Questionask
type QuestionaskController struct {
	beego.Controller
}

func (c *QuestionaskController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Questionask
// @Param	body		body 	models.Questionask	true		"body for Questionask content"
// @Success 200 {int} models.Questionask.Id
// @Failure 403 body is empty
// @router /AddQuestionask/ [post]
func (c *QuestionaskController) Post() {
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Println(k)
		fmt.Println(v)
		jsonS = jsonS + k
	}
	var v models.Questionask
	gerr := json.Unmarshal([]byte(jsonS), &v)
	fmt.Println(v)
	if gerr != nil {
		fmt.Println(gerr.Error())
	}
	if id, err := models.AddQuestionask(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get Questionask by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Questionask
// @Failure 403 :id is empty
// @router /GetQuestionaskById/:id [get]
func (c *QuestionaskController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetQuestionaskById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title GetQuestionaskByJingCaiOne
// @Description 获取精彩回答详情
// @Param	id		path 	string	true		问答主键id
// @Success 200 {object} models.QuestionaskJingCaiOne
// @Failure Error
// @router /GetQuestionaskByJingCaiOne/:id [get]
func (c *QuestionaskController) GetQuestionaskByJingCaiOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetQuestionaskByJingCaiOne(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//16.查询被提问者所有问答信息
// @Title GetQuestionaskByTid
// @Description 查询被提问者所有问答信息
// @Param	userid		path 	string	true		用户信息主键id
// @Param	page		path 	string	true		获取第几页
// @Param	size		path 	string	true		获取第几行
// @Success 200 {object} models.QuestionaskByUid
// @Failure Error
// @router /GetQuestionaskByTid/:userid/:page/:size [get]
func (c *QuestionaskController) GetQuestionaskByTid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetQuestionaskByTid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//16.查询被提问者所有问答信息总条数
// @Title GetQuestionaskByTidCount
// @Description 查询被提问者所有问答信息总条数
// @Param	userid		path 	string	true		用户主键id
// @Success json {int} json
// @Failure Error
// @router /GetQuestionaskByTidCount/:userid [get]
func (c *QuestionaskController) GetQuestionaskByTidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetQuestionaskByTidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//24.学生查询自己的提问
// @Title GetQuestionaskBySid
// @Description 学生查询自己的提问
// @Param	userid		path 	string	true		用户信息主键id
// @Param	page		path 	string	true		获取第几页
// @Param	size		path 	string	true		获取第几行
// @Success 200 {object} models.Questionask
// @Failure Error
// @router /GetQuestionaskBySid/:userid/:page/:size [get]
func (c *QuestionaskController) GetQuestionaskBySid() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetQuestionaskBySid(userid, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//24.学生查询自己的提问总条数
// @Title GetQuestionaskBySidCount
// @Description 学生查询自己的提问总条数
// @Param	userid		path 	string	true		用户信息主键id
// @Success json {int} json
// @Failure Error
// @router /GetQuestionaskBySidCount/:userid [get]
func (c *QuestionaskController) GetQuestionaskBySidCount() {
	idStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(idStr)
	v, err := models.GetQuestionaskBySidCount(userid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//32.查询精彩问答
// @Title GetQuestionaskByJingCai
// @Description 查询全部精彩问答
// @Param	page		path 	string	true		获取第几页
// @Param	size		path 	string	true		获取第几行
// @Success 200 {object} models.QuestionaskJingCai
// @Failure Error
// @router /GetQuestionaskByJingCai/:page/:size [get]
func (c *QuestionaskController) GetQuestionaskByJingCai() {
	page := c.Ctx.Input.Param(":page") //获取页数	//新加--------开始--------
	size := c.Ctx.Input.Param(":size") //获取每页显示条数 //SAdd 20151027
	pages, _ := strconv.Atoi(page)     //传来的页数
	rows, _ := strconv.Atoi(size)      //传来的显示行数
	truepages := (pages - 1) * rows    //计算舍弃多少行
	limit := rows                      //显示行数
	offset := truepages                //舍弃行数	//新加--------结束--------
	v, err := models.GetQuestionaskByJingCai(offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

//32.查询精彩问答总条数
// @Title GetQuestionaskByJingCaiCount
// @Description 查询精彩问答总条数
// @Success json {int} json
// @Failure Error
// @router /GetQuestionaskByJingCaiCount [get]
func (c *QuestionaskController) GetQuestionaskByJingCaiCount() {
	v, err := models.GetQuestionaskByJingCaiCount()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get Questionask
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Questionask
// @Failure 403
// @router /GetAllQuestionask/:page/:size [get]
func (c *QuestionaskController) GetAll() {
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

	l, err := models.GetAllQuestionask(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the Questionask
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Questionask	true		"body for Questionask content"
// @Success 200 {object} models.Questionask
// @Failure 403 :id is not int
// @router /UpdateQuestionaskById/:id [post]
func (c *QuestionaskController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	var jsonS string
	for k, v := range c.Ctx.Request.Form {
		fmt.Printf("k=%v, v=%v\n", k, v)
		jsonS = k
	}
	v := models.Questionask{Id: id}
	json.Unmarshal([]byte(jsonS), &v)
	if err := models.UpdateQuestionaskById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title DeleteQuestionask
// @Description DeleteQuestionask the Questionask
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteQuestionask/:id [get]
func (c *QuestionaskController) DeleteQuestionask() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteQuestionask(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

//学生使用的删除提问
//逻辑：判断此学生的问题冻结的资金是否解冻，如果解冻直接操作，如果没有解冻，解冻资金退还账户在执行删除
// @Title DeleteQuestionaskFStu
// @Description 学生使用的删除提问 判断此学生的问题冻结的资金是否解冻，如果解冻直接操作，如果没有解冻，解冻资金退还账户在执行删除
// @Param	id		path 	string	true		提问主键id
// @Param	userid		path 	string	true		用户主键id
// @Success OK {string} success
// @Failure Error
// @router /DeleteQuestionaskFStu/:id/:userid [get]
func (c *QuestionaskController) DeleteQuestionaskFStu() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	useridStr := c.Ctx.Input.Params[":userid"]
	userid, _ := strconv.Atoi(useridStr)
	//根据提问主键查询一条冻结资金信息
	fontfonz, geterr := models.GetFrozenfundsByUidOnId(userid, 1, id)
	if geterr == nil {
		if fontfonz.FrozenState == 1 { //资金正在冻结
			fontfonz.FrozenState = 0
			upf := models.UpdateFrozenfundsById(&fontfonz)
			if upf == nil {
				account, accerr := models.GetAccountfundsByuid(userid)
				if accerr == nil {
					account.Balance = account.Balance + fontfonz.FrozenMoney
					upacc := models.UpdateAccountfundsById(&account)
					if upacc == nil {

					}
				}
			}
		}
	}
	if err := models.DeleteQuestionask(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
