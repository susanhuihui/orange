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

type Usermessage struct {
	Id            int       `orm:"column(PKId);auto"`
	ActiveUserId  int       `orm:"column(ActiveUserId)"`
	PassiveUserId int       `orm:"column(PassiveUserId)"`
	MessageId     int       `orm:"column(MessageId)"`
	Contents      string    `orm:"column(Contents);size(100);null"`
	States        int       `orm:"column(States);null"`
	MesTime       time.Time `orm:"column(MesTime);type(datetime);null"`
}

//老师看学生留言
type UsermessageFStu struct {
	Id            int       `orm:"column(PKId);auto"`
	ActiveUserId  int       `orm:"column(ActiveUserId)"`
	PassiveUserId int       `orm:"column(PassiveUserId)"`
	MessageId     int       `orm:"column(MessageId)"`
	Contents      string    `orm:"column(Contents);size(100);null"`
	States        int       `orm:"column(States);null"`
	MesTime       time.Time `orm:"column(MesTime);type(datetime);null"`
	UserName      string    `orm:"column(UserName);size(50);null"`
	MesTimeNew    time.Time `orm:"column(MesTimeNew);type(datetime);null"`
	State         int       `orm:"column(State);null"`
}

//老师/学生看一条留言
type UsermessageOneList struct {
	Id            int       `orm:"column(PKId);auto"`
	ActiveUserId  int       `orm:"column(ActiveUserId)"`
	PassiveUserId int       `orm:"column(PassiveUserId)"`
	MessageId     int       `orm:"column(MessageId)"`
	Contents      string    `orm:"column(Contents);size(100);null"`
	States        int       `orm:"column(States);null"`
	MesTime       time.Time `orm:"column(MesTime);type(datetime);null"`
	ActiveName    string    `orm:"column(ActiveName);size(50);null"`
	PassiveName   string    `orm:"column(PassiveName);size(50);null"`
}

func (t *Usermessage) TableName() string {
	return "usermessage"
}

func init() {
	orm.RegisterModel(new(Usermessage))
}

//    7.老师看学生留言
//    2015-11-06
func GetUsermessageByTid(userid int, rows int, counts int) (usermess []UsermessageFStu, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessageTeacher+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return usermess, qs
	}
	return
}

//    7.老师看学生留言总条数
//    2015-11-06
func GetUsermessageByTidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessageTeacher, userid)
	var usermess []UsermessageFStu
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(usermess), qs
	}
	return
}

//	8.老师/学生看一条留言
//   2015-11-06
func GetUsermessageByMessageId(usermesid int, usermesidt int) (usermess []UsermessageOneList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessageTchOne, usermesid, usermesidt)
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return usermess, qs
	}
	return
}

//    25.学生查看自己的全部留言信息
//    2015-11-06
func GetUsermessageBySid(userid int, rows int, counts int) (usermess []UsermessageFStu, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessageBySid+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return usermess, qs
	}
	return
}

//    25.学生查看自己的全部留言信息总条数
//    2015-11-06
func GetUsermessageBySidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessageBySid, userid)
	var usermess []UsermessageFStu
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(usermess), qs
	}
	return
}

//
//
func GetUsermessageBymuid(mid int, userid int) (usermess []Usermessage, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserMessagebymuid, mid, userid)
	num, qs := rs.QueryRows(&usermess)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return usermess, qs
	}
	return
}

//
//
//func GetUsermessageBymuidft(mid int, userid int) (usermess []Usermessage, err error) {
//	o := orm.NewOrm()
//	var rs orm.RawSeter
//	rs = o.Raw(SqlUserMessagebymuidft, mid, userid)
//	num, qs := rs.QueryRows(&usermess)
//	if qs != nil {
//		fmt.Printf("num", num)
//		return nil, qs
//	} else {
//		return usermess, qs
//	}
//	return
//}

//更改一条留言下的所有回复为已读(错误方法)
func UpdateUsermessageBypiduid(mid int, userid int) (num int, err error) {
	o := orm.NewOrm()
	usermesg, _ := GetUsermessageBymuid(mid, userid)
	fmt.Println(usermesg)
	var rs orm.RawSeter
	for i := 0; i < len(usermesg); i++ {
		var upstr string = `update usermessage set states='1' where PKId=` + strconv.Itoa(usermesg[i].Id) + `; SELECT ROW_COUNT() as roocount;`
		fmt.Println(upstr)
		rs = o.Raw(upstr)
		err = rs.QueryRow(&num)
		fmt.Println(err)
	}
	return num, err
}

// AddUsermessage insert a new Usermessage into database and returns
// last inserted Id on success.
func AddUsermessage(m *Usermessage) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsermessageById retrieves Usermessage by Id. Returns error if
// Id doesn't exist
func GetUsermessageById(id int) (v *Usermessage, err error) {
	o := orm.NewOrm()
	v = &Usermessage{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsermessage retrieves all Usermessage matches certain condition. Returns empty list if
// no records exist
func GetAllUsermessage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Usermessage))
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

	var l []Usermessage
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

// UpdateUsermessage updates Usermessage by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsermessageById(m *Usermessage) (err error) {
	o := orm.NewOrm()
	v := Usermessage{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsermessage deletes Usermessage by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsermessage(id int) (err error) {
	o := orm.NewOrm()
	v := Usermessage{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Usermessage{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
