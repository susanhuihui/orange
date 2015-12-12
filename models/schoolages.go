package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Schoolages struct {
	Id      int    `orm:"column(PKId);auto"`
	AgeName string `orm:"column(AgeName);size(50);null"`
}

func (t *Schoolages) TableName() string {
	return "schoolages"
}

func init() {
	orm.RegisterModel(new(Schoolages))
}

// AddSchoolages insert a new Schoolages into database and returns
// last inserted Id on success.
func AddSchoolages(m *Schoolages) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSchoolagesById retrieves Schoolages by Id. Returns error if
// Id doesn't exist
func GetSchoolagesById(id int) (v *Schoolages, err error) {
	o := orm.NewOrm()
	v = &Schoolages{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//	根据学龄段名称查询一条学龄段信息
//	2015-12-12
func GetSchoolagesByName(agename string) (schoolages *Schoolages, err error) {
	o := orm.NewOrm()
	var ages Schoolages
	err = o.QueryTable("schoolages").Filter("AgeName", agename).One(&ages)
	schoolages = &ages
	if err == nil {
		return schoolages, nil
	}
	return nil, err
}

// GetAllSchoolages retrieves all Schoolages matches certain condition. Returns empty list if
// no records exist
func GetAllSchoolages(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Schoolages))
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

	var l []Schoolages
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

// UpdateSchoolages updates Schoolages by Id and returns error if
// the record to be updated doesn't exist
func UpdateSchoolagesById(m *Schoolages) (err error) {
	o := orm.NewOrm()
	v := Schoolages{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSchoolages deletes Schoolages by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSchoolages(id int) (err error) {
	o := orm.NewOrm()
	v := Schoolages{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Schoolages{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
