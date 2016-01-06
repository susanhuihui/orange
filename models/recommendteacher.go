package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Recommendteacher struct {
	Id            int       `orm:"column(PKId);auto"`
	UserId        int       `orm:"column(UserId)"`
	GradeId       int       `orm:"column(GradeId);null"`
	ClassId       int       `orm:"column(ClassId);null"`
	StartPrice    float64   `orm:"column(StartPrice);null;digits(10);decimals(2)"`
	EndPreice     float64   `orm:"column(EndPreice);null;digits(10);decimals(2)"`
	CityId        int       `orm:"column(CityId);null"`
	RecommendTime time.Time `orm:"column(RecommendTime);type(datetime);null"`
	MyName        string    `orm:"column(MyName);size(50);null"`
	MyPhone       string    `orm:"column(MyPhone);size(50);null"`
	Remarks       string    `orm:"column(Remarks);size(50);null"`
}

type RecommendteacherAll struct {
	Id            int       `orm:"column(PKId);auto"`
	UserId        int       `orm:"column(UserId)"`
	GradeId       int       `orm:"column(GradeId);null"`
	ClassId       int       `orm:"column(ClassId);null"`
	StartPrice    float64   `orm:"column(StartPrice);null;digits(10);decimals(2)"`
	EndPreice     float64   `orm:"column(EndPreice);null;digits(10);decimals(2)"`
	CityId        int       `orm:"column(CityId);null"`
	RecommendTime time.Time `orm:"column(RecommendTime);type(datetime);null"`
	MyName        string    `orm:"column(MyName);size(50);null"`
	MyPhone       string    `orm:"column(MyPhone);size(50);null"`
	Remarks       string    `orm:"column(Remarks);size(50);null"`
	IphoneNum     string    `orm:"column(IphoneNum);size(50);null"`
	UserName      string    `orm:"column(UserName);size(50);null"`
	Mailbox       string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox string    `orm:"column(ParentMailbox);size(50);null"`
	CourseName    string    `orm:"column(CourseName);size(50);null"`
	GradeName     string    `orm:"column(GradeName);size(50);null"`
	CityName      string    `orm:"column(CityName);size(50);null"`
}

func (t *Recommendteacher) TableName() string {
	return "recommendteacher"
}

func init() {
	orm.RegisterModel(new(Recommendteacher))
}

//	51.查询全部学生推荐信息
//	2016-01-06
func GetRecommendteacherAll(rows int, counts int) (recomteacher []RecommendteacherAll, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRecommendTeacherAll+limitSql, rows, counts)
	num, qs := rs.QueryRows(&recomteacher)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return recomteacher, qs
	}
	return
}

//	51.查询全部学生推荐信息
//	2016-01-06
func GetRecommendteacherAllCount() (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRecommendTeacherAll)
	var recomteacher []RecommendteacherAll
	num, qs := rs.QueryRows(&recomteacher)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(recomteacher), qs
	}
	return
}

// AddRecommendteacher insert a new Recommendteacher into database and returns
// last inserted Id on success.
func AddRecommendteacher(m *Recommendteacher) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRecommendteacherById retrieves Recommendteacher by Id. Returns error if
// Id doesn't exist
func GetRecommendteacherById(id int) (v *Recommendteacher, err error) {
	o := orm.NewOrm()
	v = &Recommendteacher{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecommendteacher retrieves all Recommendteacher matches certain condition. Returns empty list if
// no records exist
func GetAllRecommendteacher(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Recommendteacher))
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

	var l []Recommendteacher
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

// UpdateRecommendteacher updates Recommendteacher by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecommendteacherById(m *Recommendteacher) (err error) {
	o := orm.NewOrm()
	v := Recommendteacher{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRecommendteacher deletes Recommendteacher by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecommendteacher(id int) (err error) {
	o := orm.NewOrm()
	v := Recommendteacher{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Recommendteacher{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
