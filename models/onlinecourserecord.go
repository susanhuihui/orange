package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//原始在线课程结构
type Onlinecourserecord struct {
	Id            int       `orm:"column(PKId);auto"`
	OCBId         int       `orm:"column(OCBId)"`
	UserIdActive  int       `orm:"column(UserIdActive)"`
	UserIdPassive int       `orm:"column(UserIdPassive)"`
	CourseTitle   string    `orm:"column(CourseTitle);size(50);null"`
	CourseContent string    `orm:"column(CourseContent);null"`
	AttributeName string    `orm:"column(AttributeName);size(50);null"`
	StartTime     time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime       time.Time `orm:"column(EndTime);type(datetime);null"`
	UnitPrice     float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	TotalPrice    float64   `orm:"column(TotalPrice);null;digits(10);decimals(2)"`
	ClassNumber   int       `orm:"column(ClassNumber);null"`
	ReviewPath    string    `orm:"column(ReviewPath);size(200);null"`
}

//查询老师全部课程
type OnlinecourserecordByT struct {
	Id            int       `orm:"column(PKId);auto"`
	OCBId         int       `orm:"column(OCBId)"`
	UserIdActive  int       `orm:"column(UserIdActive)"`
	UserIdPassive int       `orm:"column(UserIdPassive)"`
	CourseTitle   string    `orm:"column(CourseTitle);size(50);null"`
	CourseContent string    `orm:"column(CourseContent);null"`
	AttributeName string    `orm:"column(AttributeName);size(50);null"`
	StartTime     time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime       time.Time `orm:"column(EndTime);type(datetime);null"`
	UnitPrice     float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	TotalPrice    float64   `orm:"column(TotalPrice);null;digits(10);decimals(2)"`
	ClassNumber   int       `orm:"column(ClassNumber);null"`
	ReviewPath    string    `orm:"column(ReviewPath);size(200);null"`
	CourseName    string    `orm:"column(CourseName);size(50);null"`
	UserName      string    `orm:"column(UserName);size(50);null"`
	AgeName       string    `orm:"column(AgeName);size(50);null"`
	RecordId      int       `orm:"column(RecordId)"`
}

//查询学生所有辅导过的老师们
type OnlinecourserecordByU struct {
	Id            int       `orm:"column(PKId);auto"`
	OCBId         int       `orm:"column(OCBId)"`
	UserIdActive  int       `orm:"column(UserIdActive)"`
	UserIdPassive int       `orm:"column(UserIdPassive)"`
	CourseTitle   string    `orm:"column(CourseTitle);size(50);null"`
	CourseContent string    `orm:"column(CourseContent);null"`
	AttributeName string    `orm:"column(AttributeName);size(50);null"`
	StartTime     time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime       time.Time `orm:"column(EndTime);type(datetime);null"`
	UnitPrice     float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	TotalPrice    float64   `orm:"column(TotalPrice);null;digits(10);decimals(2)"`
	ClassNumber   int       `orm:"column(ClassNumber);null"`
	ReviewPath    string    `orm:"column(ReviewPath);size(200);null"`
	UserName      string    `orm:"column(UserName);size(50);null"`
}

func (t *Onlinecourserecord) TableName() string {
	return "onlinecourserecord"
}

func init() {
	orm.RegisterModel(new(Onlinecourserecord))
}

//	根据预约课程助教查询一条课程信息
//	2015-11-25
func GetOnlinecourserecordByBookid(bookid int) (v *Onlinecourserecord, err error) {
	o := orm.NewOrm()
	var oncord Onlinecourserecord
	err = o.QueryTable("onlinecourserecord").Filter("OCBId", bookid).One(&oncord)
	v = &oncord
	if err == nil {
		return v, nil
	}
	return nil, err
}

//    4.查询老师全部课程信息
//    2015-11-06
func GetOnlinecourserecordByTid(userid int, rows int, counts int) (onlinecourse []OnlinecourserecordByT, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineAllTeacher+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinecourse, qs
	}
	return
}

//    4.查询老师全部课程信息总条数
//    2015-11-06
func GetOnlinecourserecordByTidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineAllTeacher, userid)
	var onlinecourse []OnlinecourserecordByT
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinecourse), qs
	}
	return
}

//   18.查询学生全部课程
//   2015-11-06
func GetOnlinecourserecordByUid(userid int, rows int, counts int) (onlinecourse []OnlinecourserecordByT, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineAllStudent+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinecourse, qs
	}
	return
}

//   18.查询学生全部课程总条数
//   2015-11-06
func GetOnlinecourserecordByUidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineAllStudent, userid)
	var onlinecourse []OnlinecourserecordByT
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinecourse), qs
	}
	return
}

//   40.查询给我上过课的老师们
//   2015-11-06
func GetOnlinecourserecordTeacherByUid(userid int) (onlinecourse []OnlinecourserecordByU, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineCourseRecordTByUid, userid)
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinecourse, qs
	}
	return
}

//   43.查询给我上过课的某个学科的老师们
//   2015-12-11

// GetOnlinecourserecordTeacherByUCid retrieves Onlinecourserecord by userid classid. Returns error if
// Id doesn't exist
func GetOnlinecourserecordTeacherByUCid(userid int, classid int) (onlinecourse []OnlinecourserecordByU, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineCourseRecordTByUCid, userid, classid)
	num, qs := rs.QueryRows(&onlinecourse)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinecourse, qs
	}
	return
}

// AddOnlinecourserecord insert a new Onlinecourserecord into database and returns
// last inserted Id on success.
func AddOnlinecourserecord(m *Onlinecourserecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnlinecourserecordById retrieves Onlinecourserecord by Id. Returns error if
// Id doesn't exist
func GetOnlinecourserecordById(id int) (v *Onlinecourserecord, err error) {
	o := orm.NewOrm()
	v = &Onlinecourserecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOnlinecourserecord retrieves all Onlinecourserecord matches certain condition. Returns empty list if
// no records exist
func GetAllOnlinecourserecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Onlinecourserecord))
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

	var l []Onlinecourserecord
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

// UpdateOnlinecourserecord updates Onlinecourserecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnlinecourserecordById(m *Onlinecourserecord) (err error) {
	o := orm.NewOrm()
	v := Onlinecourserecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOnlinecourserecord deletes Onlinecourserecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinecourserecord(id int) (err error) {
	o := orm.NewOrm()
	v := Onlinecourserecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Onlinecourserecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
