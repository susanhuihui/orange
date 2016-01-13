package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Onlinecoursebookingrecord struct {
	Id        int       `orm:"column(PKId);auto"`
	OCBId     int       `orm:"column(OCBId)"`
	UserId    int       `orm:"column(UserId)"`
	StartTime time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime   time.Time `orm:"column(EndTime);type(datetime);null"`
}

func (t *Onlinecoursebookingrecord) TableName() string {
	return "onlinecoursebookingrecord"
}

func init() {
	orm.RegisterModel(new(Onlinecoursebookingrecord))
}

// AddOnlinecoursebookingrecord insert a new Onlinecoursebookingrecord into database and returns
// last inserted Id on success.
func AddOnlinecoursebookingrecord(m *Onlinecoursebookingrecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnlinecoursebookingrecordById retrieves Onlinecoursebookingrecord by Id. Returns error if
// Id doesn't exist
func GetOnlinecoursebookingrecordById(id int) (v *Onlinecoursebookingrecord, err error) {
	o := orm.NewOrm()
	v = &Onlinecoursebookingrecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//	46.查询老师或学生一条课堂时间记录，一条时间最近且结束时间为null的记录 并记录结束时间
//	2015-12-19
func GetOnlinecoursebookingrecordByUid(userid int, ocbid int) (bookrecord Onlinecoursebookingrecord, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBookingRecord, userid, ocbid)
	qs := rs.QueryRow(&bookrecord)
	if qs == nil {
		var roocount int
		var timestrpm = time.Now().Format("2006-01-02 15:04:05 PM")
		fmt.Println("当前结束日期为")
		fmt.Println(timestrpm)
		var upstr string = `update onlinecoursebookingrecord set EndTime='` + timestrpm + `'  where PKId=` + strconv.Itoa(bookrecord.Id) + `; SELECT ROW_COUNT() as roocount;`
		rs = o.Raw(upstr)
		rs.QueryRow(&roocount)
		fmt.Println(roocount)
	}
	return bookrecord, qs
}

//	46.查询老师或学生一条课堂时间记录，一条时间最近且结束时间为null的记录 并记录结束时间
//	2015-12-19
func GetOnlinecoursebookingrecordByUid2(userid int, ocbid int) (bookrecord Onlinecoursebookingrecord, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBookingRecord, userid, ocbid)
	qs := rs.QueryRow(&bookrecord)
	if qs == nil {
		bookrecord.EndTime = time.Now()
		up := UpdateOnlinecoursebookingrecordById(&bookrecord)
		if up == nil {

		}
	}
	return bookrecord, qs
}

//	47.查询老师或学生关于某次课程的全部课程时间记录信息
//	2015-12-20
func GetOnlinecoursebookingrecordBybookiduid(userid int, bookid int) (onlinerecord []Onlinecoursebookingrecord, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBookingRecordBybookiduid, userid, bookid)
	num, qs := rs.QueryRows(&onlinerecord)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinerecord, qs
	}
	return
}

// GetAllOnlinecoursebookingrecord retrieves all Onlinecoursebookingrecord matches certain condition. Returns empty list if
// no records exist
func GetAllOnlinecoursebookingrecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Onlinecoursebookingrecord))
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

	var l []Onlinecoursebookingrecord
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

// UpdateOnlinecoursebookingrecord updates Onlinecoursebookingrecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnlinecoursebookingrecordById(m *Onlinecoursebookingrecord) (err error) {
	o := orm.NewOrm()
	v := Onlinecoursebookingrecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOnlinecoursebookingrecord deletes Onlinecoursebookingrecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinecoursebookingrecord(id int) (err error) {
	o := orm.NewOrm()
	v := Onlinecoursebookingrecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Onlinecoursebookingrecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
