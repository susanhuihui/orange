package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Accountfunds struct {
	Id            int       `orm:"column(PKId);auto"`
	UserId        int       `orm:"column(UserId)"`
	Balance       float64   `orm:"column(Balance);null;digits(10);decimals(2)"`
	FundState     int       `orm:"column(FundState)"`
	OpenTime      time.Time `orm:"column(OpenTime);type(datetime)"`
	AccountTypeId int       `orm:"column(AccountTypeId)"`
}

//
type AccountfundsStudent struct {
	Id            int       `orm:"column(PKId);auto"`
	UserId        int       `orm:"column(UserId)"`
	Balance       float64   `orm:"column(Balance);null;digits(10);decimals(2)"`
	FundState     int       `orm:"column(FundState)"`
	OpenTime      time.Time `orm:"column(OpenTime);type(datetime)"`
	AccountTypeId int       `orm:"column(AccountTypeId)"`
	FrozenMoney   float64   `orm:"column(FrozenMoney);null;digits(10);decimals(2)"`
}

func (t *Accountfunds) TableName() string {
	return "accountfunds"
}

func init() {
	orm.RegisterModel(new(Accountfunds))
}

//	13.查询用户资金表
//	2015-11-06
func GetAccountfundsByuid(userid int) (account Accountfunds, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter	
	rs = o.Raw(SqlAccountFundsByUserId,userid)
	qs := rs.QueryRow(&account)
	return account, qs
}

//	22.查询用户账户信息和冻结资金
//	2015-11-06
func GetAccountfundsBySID(userid int) (account AccountfundsStudent, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter	
	rs = o.Raw(SqlAccountFundsByS,userid)
	qs := rs.QueryRow(&account)
	return account, qs
}

// AddAccountfunds insert a new Accountfunds into database and returns
// last inserted Id on success.
func AddAccountfunds(m *Accountfunds) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAccountfundsById retrieves Accountfunds by Id. Returns error if
// Id doesn't exist
func GetAccountfundsById(id int) (v *Accountfunds, err error) {
	o := orm.NewOrm()
	v = &Accountfunds{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAccountfunds retrieves all Accountfunds matches certain condition. Returns empty list if
// no records exist
func GetAllAccountfunds(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Accountfunds))
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

	var l []Accountfunds
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

// UpdateAccountfunds updates Accountfunds by Id and returns error if
// the record to be updated doesn't exist
func UpdateAccountfundsById(m *Accountfunds) (err error) {
	o := orm.NewOrm()
	v := Accountfunds{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAccountfunds deletes Accountfunds by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAccountfunds(id int) (err error) {
	o := orm.NewOrm()
	v := Accountfunds{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Accountfunds{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
