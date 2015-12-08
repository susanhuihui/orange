package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Citys struct {
	Id       int    `orm:"column(PKId);auto"`
	ProId    int    `orm:"column(ProId)"`
	CityCode string `orm:"column(CityCode);size(50);null"`
	CityName string `orm:"column(CityName);size(50);null"`
}

func (t *Citys) TableName() string {
	return "citys"
}

func init() {
	orm.RegisterModel(new(Citys))
}

// AddCitys insert a new Citys into database and returns
// last inserted Id on success.
func AddCitys(m *Citys) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCitysById retrieves Citys by Id. Returns error if
// Id doesn't exist
func GetCitysById(id int) (v *Citys, err error) {
	o := orm.NewOrm()
	v = &Citys{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//	根据省份主键id查询市信息
//	2015-11-17
func GetCitysByPid(pid int) (v []Citys, err error) {
	o := orm.NewOrm()
	var citylist []Citys
	qs,err := o.QueryTable("Citys").Filter("ProId",pid).All(&citylist)
	v = citylist	
	if err == nil && qs != 0{
		return v, nil
	}
	return nil, err
}

// GetAllCitys retrieves all Citys matches certain condition. Returns empty list if
// no records exist
func GetAllCitys(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Citys))
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

	var l []Citys
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

// UpdateCitys updates Citys by Id and returns error if
// the record to be updated doesn't exist
func UpdateCitysById(m *Citys) (err error) {
	o := orm.NewOrm()
	v := Citys{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCitys deletes Citys by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCitys(id int) (err error) {
	o := orm.NewOrm()
	v := Citys{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Citys{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
