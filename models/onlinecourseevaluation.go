package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Onlinecourseevaluation struct {
	Id              int       `orm:"column(PKId);auto"`
	OCRId           int       `orm:"column(OCRId)"`
	UserId          int       `orm:"column(UserId)"`
	StartClear      int       `orm:"column(StartClear);null"`
	StartEfficiency int       `orm:"column(StartEfficiency);null"`
	ReviewContent   string    `orm:"column(ReviewContent);size(100);null"`
	ReviewTime      time.Time `orm:"column(ReviewTime);type(datetime);null"`
}

//查询老师在线课程评价
type OnlinecourseevaluationByT struct {
	Id              int       `orm:"column(PKId);auto"`
	OCRId           int       `orm:"column(OCRId)"`
	UserId          int       `orm:"column(UserId)"`
	StartClear      int       `orm:"column(StartClear);null"`
	StartEfficiency int       `orm:"column(StartEfficiency);null"`
	ReviewContent   string    `orm:"column(ReviewContent);size(100);null"`
	ReviewTime      time.Time `orm:"column(ReviewTime);type(datetime);null"`
	UserPkid        int       `orm:"column(UserPkid)"`
}

//查询学生在线课程评价
type OnlinecourseevaluationByS struct {
	Id              int       `orm:"column(PKId);auto"`
	OCRId           int       `orm:"column(OCRId)"`
	UserId          int       `orm:"column(UserId)"`
	StartClear      int       `orm:"column(StartClear);null"`
	StartEfficiency int       `orm:"column(StartEfficiency);null"`
	ReviewContent   string    `orm:"column(ReviewContent);size(100);null"`
	ReviewTime      time.Time `orm:"column(ReviewTime);type(datetime);null"`
	UserPkid        int       `orm:"column(UserPkid)"`
	UserName        string    `orm:"column(UserName);size(50);null"`
	AvatarPath      string    `orm:"column(AvatarPath);size(50);null"`
}

//查询老师所有在线评价
type OnlinecourseevaluationByTid struct {
	Id              int       `orm:"column(PKId);auto"`
	OCRId           int       `orm:"column(OCRId)"`
	UserId          int       `orm:"column(UserId)"`
	StartClear      int       `orm:"column(StartClear);null"`
	StartEfficiency int       `orm:"column(StartEfficiency);null"`
	ReviewContent   string    `orm:"column(ReviewContent);size(100);null"`
	ReviewTime      time.Time `orm:"column(ReviewTime);type(datetime);null"`
	UserPkid        int       `orm:"column(UserPkid)"`
	UserName        string    `orm:"column(UserName);size(50);null"`
	AvatarPath      string    `orm:"column(AvatarPath);size(50);null"`
	AllStart        int       `orm:"column(AllStart)"`
}

func (t *Onlinecourseevaluation) TableName() string {
	return "onlinecourseevaluation"
}

func init() {
	orm.RegisterModel(new(Onlinecourseevaluation))
}

//    5.老师查询在线课程评价内容
//    2015-11-06
func GetOnlinecourseevaluationByOcridT (ocrid int) (onlierecourseval []OnlinecourseevaluationByT, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineEvaluationByT,ocrid)
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return onlierecourseval, qs
    }
    return
}

//    19.学生查询在线课程评价内容
//    2015-11-06
func GetOnlinecourseevaluationByOcridS (ocrid int) (onlierecourseval []OnlinecourseevaluationByS, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineEvaluationByS,ocrid)
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return onlierecourseval, qs
    }
    return
}

//    29.查询老师的所有在线课程评价
//    2015-11-06
func GetOnlinecourseevaluationByTid (userid int,rows int,counts int) (onlierecourseval []OnlinecourseevaluationByTid, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineCourseEvalAllByTid+limitSql,userid,rows,counts)
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return onlierecourseval, qs
    }
    return
}

//    29.查询老师的所有在线课程评价总条数
//    2015-11-06
func GetOnlinecourseevaluationByTidCount (userid int) (allcount int, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineCourseEvalAllByTid,userid)
	var onlierecourseval []OnlinecourseevaluationByTid
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return 0, qs
    } else {
        return len(onlierecourseval), qs
    }
    return
}



//    34.学生查看自己所评价的老师们
//    2015-11-18
func GetOnlineCourseEvaluationBySid (userid int,rows int,counts int) (onlierecourseval []OnlinecourseevaluationByS, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineCourseEvaluationBySid+limitSql,userid,rows,counts)
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return onlierecourseval, qs
    }
    return
}

//    34.学生查看自己所评价的老师们总条数
//    2015-11-18
func GetOnlineCourseEvaluationBySidCount (userid int) (allcount int, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlOnlineCourseEvaluationBySid,userid)
	var onlierecourseval []OnlinecourseevaluationByS
    num, qs := rs.QueryRows(&onlierecourseval)
    if qs != nil {
        fmt.Printf("num", num)
        return 0, qs
    } else {
        return len(onlierecourseval), qs
    }
    return
}

// AddOnlinecourseevaluation insert a new Onlinecourseevaluation into database and returns
// last inserted Id on success.
func AddOnlinecourseevaluation(m *Onlinecourseevaluation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnlinecourseevaluationById retrieves Onlinecourseevaluation by Id. Returns error if
// Id doesn't exist
func GetOnlinecourseevaluationById(id int) (v *Onlinecourseevaluation, err error) {
	o := orm.NewOrm()
	v = &Onlinecourseevaluation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOnlinecourseevaluation retrieves all Onlinecourseevaluation matches certain condition. Returns empty list if
// no records exist
func GetAllOnlinecourseevaluation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Onlinecourseevaluation))
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

	var l []Onlinecourseevaluation
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

// UpdateOnlinecourseevaluation updates Onlinecourseevaluation by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnlinecourseevaluationById(m *Onlinecourseevaluation) (err error) {
	o := orm.NewOrm()
	v := Onlinecourseevaluation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOnlinecourseevaluation deletes Onlinecourseevaluation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinecourseevaluation(id int) (err error) {
	o := orm.NewOrm()
	v := Onlinecourseevaluation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Onlinecourseevaluation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
