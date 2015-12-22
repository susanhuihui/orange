package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Questionask struct {
	Id           int       `orm:"column(PKId);auto"`
	AskUserId    int       `orm:"column(AskUserId)"`
	AnswerUserId int       `orm:"column(AnswerUserId)"`
	GCId         int       `orm:"column(GCId)"`
	Title        string    `orm:"column(Title);size(50);null"`
	Contents     string    `orm:"column(Contents);null"`
	BadeTime     time.Time `orm:"column(BadeTime);type(datetime);null"`
	EndTime      time.Time `orm:"column(EndTime);type(datetime);null"`
	AmountMoney  float64   `orm:"column(AmountMoney);null;digits(10);decimals(2)"`
	IsSee        int       `orm:"column(IsSee)"`
}

//查询被提问者回答
type QuestionaskByUid struct {
	Id           int       `orm:"column(PKId);auto"`
	AskUserId    int       `orm:"column(AskUserId)"`
	AnswerUserId int       `orm:"column(AnswerUserId)"`
	GCId         int       `orm:"column(GCId)"`
	Title        string    `orm:"column(Title);size(50);null"`
	Contents     string    `orm:"column(Contents);null"`
	BadeTime     time.Time `orm:"column(BadeTime);type(datetime);null"`
	EndTime      time.Time `orm:"column(EndTime);type(datetime);null"`
	AmountMoney  float64   `orm:"column(AmountMoney);null;digits(10);decimals(2)"`
	IsSee        int       `orm:"column(IsSee)"`
	UserName     string    `orm:"column(UserName);size(50);null"`
	AnswerCount  int       `orm:"column(AnswerCount)"`
}

//查询精彩问答
type QuestionaskJingCai struct {
	Id       int       `orm:"column(PKId);auto"`
	Title    string    `orm:"column(Title);size(50);null"`
	Contents string    `orm:"column(Contents);null"`
	BadeTime time.Time `orm:"column(BadeTime);type(datetime);null"`
	Numbers  int       `orm:"column(Numbers)"`
	State    int       `orm:"column(State)"`
	Count    int       `orm:"column(Count)"`
}

//查询精彩问答详情
type QuestionaskJingCaiOne struct {
	Title           string    `orm:"column(Title);size(50);null"`
	Contents        string    `orm:"column(Contents);null"`
	BadeTime        time.Time `orm:"column(BadeTime);type(datetime);null"`
	AskUserId       int       `orm:"column(AskUserId)"`
	UserName        string    `orm:"column(UserName);size(50);null"`
	Hname           string    `orm:"column(Hname);size(50);null"`
	HuiDaContents   string    `orm:"column(HuiDaContents);size(50);null"`
	AnsTime         time.Time `orm:"column(AnsTime);type(datetime);null"`
	AvatarPath      string    `orm:"column(AvatarPath);size(50);null"`
	HuiDaAvatarPath string    `orm:"column(HuiDaAvatarPath);size(50);null"`
	AnswerUserId    int       `orm:"column(AnswerUserId)"`
	SchoolName      string    `orm:"column(SchoolName);size(50);null"`
	UserSchoolName  string    `orm:"column(UserSchoolName);size(50);null"`
	AnswerId        int       `orm:"column(AnswerId)"`
	Id              int       `orm:"column(PKId);auto"`
}

func (t *Questionask) TableName() string {
	return "questionask"
}

func init() {
	orm.RegisterModel(new(Questionask))
}

//    16.查询被提问者所有问答信息
//    2015-11-06
func GetQuestionaskByTid(userid int, rows int, counts int) (list []QuestionaskByUid, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskByTUserid+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return list, qs
	}
	return
}

//    16.查询被提问者所有问答信息总条数
//    2015-11-18
func GetQuestionaskByTidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskByTUserid, userid)
	var list []QuestionaskByUid
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(list), qs
	}
	return
}

//    24.查询提问者所有问答信息 学生查询自己的提问
//    2015-11-06
func GetQuestionaskBySid(userid int, rows int, counts int) (list []QuestionaskByUid, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskBySUserid+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return list, qs
	}
	return
}

//    24.查询提问者所有问答信息 学生查询自己的提问总条数
//    2015-11-06
func GetQuestionaskBySidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskBySUserid, userid)
	var list []QuestionaskByUid
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(list), qs
	}
	return
}

//    32.查询精彩问答
//    2015-11-06
func GetQuestionaskByJingCai(rows int, counts int) (list []QuestionaskJingCai, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskJingcai+limitSql, rows, counts)
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return list, qs
	}
	return
}

//    32.查询精彩问答总条数
//    2015-11-06
func GetQuestionaskByJingCaiCount() (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskJingcai)
	var list []QuestionaskJingCai
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(list), qs
	}
	return
}

//	33.精彩回答详情
//	2015-11-06
func GetQuestionaskByJingCaiOne(qaid int) (quest QuestionaskJingCaiOne, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlQuestionAskById, qaid)
	qs := rs.QueryRow(&quest)
	return quest, qs
}

// AddQuestionask insert a new Questionask into database and returns
// last inserted Id on success.
func AddQuestionask(m *Questionask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuestionaskById retrieves Questionask by Id. Returns error if
// Id doesn't exist
func GetQuestionaskById(id int) (v *Questionask, err error) {
	o := orm.NewOrm()
	v = &Questionask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQuestionask retrieves all Questionask matches certain condition. Returns empty list if
// no records exist
func GetAllQuestionask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Questionask))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Questionask
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateQuestionask updates Questionask by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuestionaskById(m *Questionask) (err error) {
	o := orm.NewOrm()
	v := Questionask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQuestionask deletes Questionask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuestionask(id int) (err error) {
	o := orm.NewOrm()
	v := Questionask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Questionask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
