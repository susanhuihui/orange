package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Schools struct {
	Id         int    `orm:"column(PKId);auto"`
	SchoolName string `orm:"column(SchoolName);size(50);null"`
	CityId     int    `orm:"column(CityId);null"`
	SchoolType int    `orm:"column(SchoolType);null"`
}

func (t *Schools) TableName() string {
	return "schools"
}

func init() {
	orm.RegisterModel(new(Schools))
}

//    12.根据类型查询相应类型的学校（0小学，2初中，3高中，大学1）
//    2015-11-06
func GetSchoolsByCity(cityid int, schooltype int) (list []Schools, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlSchoolByCityType, cityid, schooltype)
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return list, qs
	}
	return
}

// AddSchools insert a new Schools into database and returns
// last inserted Id on success.
func AddSchools(m *Schools) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSchoolsById retrieves Schools by Id. Returns error if
// Id doesn't exist
func GetSchoolsById(id int) (v *Schools, err error) {
	o := orm.NewOrm()
	v = &Schools{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSchools retrieves all Schools matches certain condition. Returns empty list if
// no records exist
func GetAllSchools(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Schools))
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

	var l []Schools
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

// UpdateSchools updates Schools by Id and returns error if
// the record to be updated doesn't exist
func UpdateSchoolsById(m *Schools) (err error) {
	o := orm.NewOrm()
	v := Schools{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSchools deletes Schools by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSchools(id int) (err error) {
	o := orm.NewOrm()
	v := Schools{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Schools{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
