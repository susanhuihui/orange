package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Transactionrecords struct {
	Id            int       `orm:"column(PKId);auto"`
	SendUserId    int       `orm:"column(SendUserId)"`
	CollectUserId int       `orm:"column(CollectUserId)"`
	RecordMoney   float64   `orm:"column(RecordMoney);digits(10);decimals(2)"`
	TradingWayId  int       `orm:"column(TradingWayId)"`
	RecordTime    time.Time `orm:"column(RecordTime);type(datetime)"`
}

//查询用户交易记录
type TransactionrecordsUserList struct {
	Id            int       `orm:"column(PKId);auto"`
	SendUserId    int       `orm:"column(SendUserId)"`
	CollectUserId int       `orm:"column(CollectUserId)"`
	RecordMoney   float64   `orm:"column(RecordMoney);digits(10);decimals(2)"`
	TradingWayId  int       `orm:"column(TradingWayId)"`
	RecordTime    time.Time `orm:"column(RecordTime);type(datetime)"`
	UserName      string    `orm:"column(UserName);size(50);null"`
	TradingName   string    `orm:"column(TradingName);size(50);null"`
}

func (t *Transactionrecords) TableName() string {
	return "transactionrecords"
}

func init() {
	orm.RegisterModel(new(Transactionrecords))
}

//    14.查询老师用户交易记录
//    2015-11-06
func GetTransactionrecordsByTid (userid int,rows int,counts int) (list []TransactionrecordsUserList, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlTransactionRecordByT+limitSql,userid,rows,counts)
    num, qs := rs.QueryRows(&list)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return list, qs
    }
    return
}

//    14.查询老师用户交易记录总条数
//    2015-11-06
func GetTransactionrecordsByTidCount (userid int) (allcount int, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlTransactionRecordByT,userid)
	var list []TransactionrecordsUserList
    num, qs := rs.QueryRows(&list)
    if qs != nil {
        fmt.Printf("num", num)
        return 0, qs
    } else {
        return len(list), qs
    }
    return
}

//    23.查询学生用户交易记录
//    2015-11-06
func GetTransactionrecordsBySid (userid int,rows int,counts int) (list []TransactionrecordsUserList, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlTranscationRecordsByUserid+limitSql,userid,rows,counts)
    num, qs := rs.QueryRows(&list)
    if qs != nil {
        fmt.Printf("num", num)
        return nil, qs
    } else {
        return list, qs
    }
    return
}

//    23.查询学生用户交易记录总条数
//    2015-11-06
func GetTransactionrecordsBySidCount (userid int) (allcount int, err error) {
    o := orm.NewOrm()
    var rs orm.RawSeter    
    rs = o.Raw(SqlTranscationRecordsByUserid,userid)
	var list []TransactionrecordsUserList
    num, qs := rs.QueryRows(&list)
    if qs != nil {
        fmt.Printf("num", num)
        return 0, qs
    } else {
        return len(list), qs
    }
    return
}

// AddTransactionrecords insert a new Transactionrecords into database and returns
// last inserted Id on success.
func AddTransactionrecords(m *Transactionrecords) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTransactionrecordsById retrieves Transactionrecords by Id. Returns error if
// Id doesn't exist
func GetTransactionrecordsById(id int) (v *Transactionrecords, err error) {
	o := orm.NewOrm()
	v = &Transactionrecords{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTransactionrecords retrieves all Transactionrecords matches certain condition. Returns empty list if
// no records exist
func GetAllTransactionrecords(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Transactionrecords))
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

	var l []Transactionrecords
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

// UpdateTransactionrecords updates Transactionrecords by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransactionrecordsById(m *Transactionrecords) (err error) {
	o := orm.NewOrm()
	v := Transactionrecords{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTransactionrecords deletes Transactionrecords by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTransactionrecords(id int) (err error) {
	o := orm.NewOrm()
	v := Transactionrecords{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Transactionrecords{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
