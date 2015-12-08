package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Courseware struct {
	Id          int       `orm:"column(PKId);auto"`
	OCBRId      int       `orm:"column(OCBRId)"`
	UserId      int       `orm:"column(UserId)"`
	CourseName  string    `orm:"column(CourseName);size(50);null"`
	Content     string    `orm:"column(Content);size(50);null"`
	FileType    string    `orm:"column(FileType);size(50);null"`
	CoursePath  string    `orm:"column(CoursePath);size(200);null"`
	CourseType  int       `orm:"column(CourseType)"`
	AuditStatus int       `orm:"column(AuditStatus)"`
	UploadTime  time.Time `orm:"column(UploadTime);type(datetime);null"`
}

func (t *Courseware) TableName() string {
	return "courseware"
}

func init() {
	orm.RegisterModel(new(Courseware))
}

// AddCourseware insert a new Courseware into database and returns
// last inserted Id on success.
func AddCourseware(m *Courseware) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCoursewareById retrieves Courseware by Id. Returns error if
// Id doesn't exist
func GetCoursewareById(id int) (v *Courseware, err error) {
	o := orm.NewOrm()
	v = &Courseware{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//    37.查询预约课程附件信息
//    2015-11-19
func GetCoursewareByOCBID (ocbrid int) (coursew []Courseware, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlCoursewareByOCBID,ocbrid)
    num, qs := rs.QueryRows(&coursew)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return coursew, qs
    }
    return
}

// GetAllCourseware retrieves all Courseware matches certain condition. Returns empty list if
// no records exist
func GetAllCourseware(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Courseware))
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

	var l []Courseware
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

// UpdateCourseware updates Courseware by Id and returns error if
// the record to be updated doesn't exist
func UpdateCoursewareById(m *Courseware) (err error) {
	o := orm.NewOrm()
	v := Courseware{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourseware deletes Courseware by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourseware(id int) (err error) {
	o := orm.NewOrm()
	v := Courseware{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Courseware{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
