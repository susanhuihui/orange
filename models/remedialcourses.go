package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Remedialcourses struct {
	Id        int `orm:"column(PKId);auto"`
	UserId    int `orm:"column(UserId)"`
	CoursesId int `orm:"column(CoursesId)"`
	IsMain    int `orm:"column(IsMain)"`
}

//查询主辅导课程
type RemedialcoursesMain struct {
	Id         int    `orm:"column(PKId);auto"`
	UserId     int    `orm:"column(UserId)"`
	CoursesId  int    `orm:"column(CoursesId)"`
	IsMain     int    `orm:"column(IsMain)"`
	CourseName string `orm:"column(CourseName);size(50);null"`
}

func (t *Remedialcourses) TableName() string {
	return "remedialcourses"
}

func init() {
	orm.RegisterModel(new(Remedialcourses))
}

//	3.查询老师主辅导课程或辅辅导课程/学生的学习难点
//	参数：ismain（0为否，1为是）(学生都为0，老师主为1辅为0）
//	2015-11-06
func GetRemedialcoursesMain(userid int, ismain int) (remedial []RemedialcoursesMain, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMainCourse, userid, ismain)
	num, qs := rs.QueryRows(&remedial)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return remedial, qs
	}
	return
}

// AddRemedialcourses insert a new Remedialcourses into database and returns
// last inserted Id on success.
func AddRemedialcourses(m *Remedialcourses) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRemedialcoursesById retrieves Remedialcourses by Id. Returns error if
// Id doesn't exist
func GetRemedialcoursesById(id int) (v *Remedialcourses, err error) {
	o := orm.NewOrm()
	v = &Remedialcourses{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRemedialcourses retrieves all Remedialcourses matches certain condition. Returns empty list if
// no records exist
func GetAllRemedialcourses(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Remedialcourses))
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

	var l []Remedialcourses
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

// UpdateRemedialcourses updates Remedialcourses by Id and returns error if
// the record to be updated doesn't exist
func UpdateRemedialcoursesById(m *Remedialcourses) (err error) {
	o := orm.NewOrm()
	v := Remedialcourses{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRemedialcourses deletes Remedialcourses by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRemedialcourses(id int) (err error) {
	o := orm.NewOrm()
	v := Remedialcourses{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Remedialcourses{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
