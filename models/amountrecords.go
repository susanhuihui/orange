package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Amountrecords struct {
	Id           int       `orm:"column(PKId);auto"`
	UserId       int       `orm:"column(UserId)"`
	RecordMoney  float64   `orm:"column(RecordMoney);digits(10);decimals(2)"`
	Balance      float64   `orm:"column(Balance);digits(10);decimals(2)"`
	RecordType   int       `orm:"column(RecordType)"`
	RecordTime   time.Time `orm:"column(RecordTime);type(datetime)"`
	TradingWayId int       `orm:"column(TradingWayId)"`
	IsComplete   int       `orm:"column(IsComplete)"`
}

//查询用户充值提现记录
type AmountrecordsUserList struct {
	Id           int       `orm:"column(PKId);auto"`
	UserId       int       `orm:"column(UserId)"`
	RecordMoney  float64   `orm:"column(RecordMoney);digits(10);decimals(2)"`
	Balance      float64   `orm:"column(Balance);digits(10);decimals(2)"`
	RecordType   int       `orm:"column(RecordType)"`
	RecordTime   time.Time `orm:"column(RecordTime);type(datetime)"`
	TradingWayId int       `orm:"column(TradingWayId)"`
	IsComplete   int       `orm:"column(IsComplete)"`
	TradingName  string    `orm:"column(TradingName);size(50);null"`
}

func (t *Amountrecords) TableName() string {
	return "amountrecords"
}

func init() {
	orm.RegisterModel(new(Amountrecords))
}

//    15.查询用户（提现/充值）记录
//    2015-11-06
func GetAmountrecordsByUserid(recordtype int, userid int, rows int, counts int) (list []AmountrecordsUserList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlAccountRecordByUidType+limitSql, recordtype, userid, rows, counts)
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return list, qs
	}
	return
}

//    15.查询用户（提现/充值）记录总条数
//    2015-11-06
func GetAmountrecordsByUseridCount(recordtype int, userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlAccountRecordByUidType, recordtype, userid)
	var list []AmountrecordsUserList
	num, qs := rs.QueryRows(&list)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(list), qs
	}
	return
}

// AddAmountrecords insert a new Amountrecords into database and returns
// last inserted Id on success.
func AddAmountrecords(m *Amountrecords) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAmountrecordsById retrieves Amountrecords by Id. Returns error if
// Id doesn't exist
func GetAmountrecordsById(id int) (v *Amountrecords, err error) {
	o := orm.NewOrm()
	v = &Amountrecords{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAmountrecords retrieves all Amountrecords matches certain condition. Returns empty list if
// no records exist
func GetAllAmountrecords(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Amountrecords))
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

	var l []Amountrecords
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

// UpdateAmountrecords updates Amountrecords by Id and returns error if
// the record to be updated doesn't exist
func UpdateAmountrecordsById(m *Amountrecords) (err error) {
	o := orm.NewOrm()
	v := Amountrecords{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAmountrecords deletes Amountrecords by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAmountrecords(id int) (err error) {
	o := orm.NewOrm()
	v := Amountrecords{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Amountrecords{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
