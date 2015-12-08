package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Frozenfunds struct {
	Id          int       `orm:"column(PKId);auto"`
	UserId      int       `orm:"column(UserId)"`
	FrozenMoney float64   `orm:"column(FrozenMoney);null;digits(10);decimals(2)"`
	FrozenType  int       `orm:"column(FrozenType);null"`
	BusinessId  int       `orm:"column(BusinessId)"`
	FrozenTime  time.Time `orm:"column(FrozenTime);type(datetime);null"`
	ThawingTime time.Time `orm:"column(ThawingTime);type(datetime);null"`
	FrozenState int       `orm:"column(FrozenState)"`
}
type FrozenfundsMoney struct {
	FrozenMoney float64 `orm:"column(FrozenMoney);null;digits(10);decimals(2)"`
}

func (t *Frozenfunds) TableName() string {
	return "frozenfunds"
}

func init() {
	orm.RegisterModel(new(Frozenfunds))
}

// AddFrozenfunds insert a new Frozenfunds into database and returns
// last inserted Id on success.
func AddFrozenfunds(m *Frozenfunds) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFrozenfundsById retrieves Frozenfunds by Id. Returns error if
// Id doesn't exist
func GetFrozenfundsById(id int) (v *Frozenfunds, err error) {
	o := orm.NewOrm()
	v = &Frozenfunds{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//    35.查询学生全部已经冻结的资金总和
//    2015-11-18
func GetFrozenFundsByUserid(userid int) (frozen FrozenfundsMoney, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlFrozenFundsByUserid, userid)
	qs := rs.QueryRow(&frozen)
	fmt.Println(frozen.FrozenMoney)
	return frozen, qs
}

//    38.查询学生预约课程信息相关的冻结信息，条件：用户主键id，是预约还是提问，预约id或提问id
//    2015-11-18
func GetFrozenfundsByUidOnId(userid int, selType int, selId int) (frozen Frozenfunds, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlFrozenfundsByUidOnId, userid, selType, selId)
	qs := rs.QueryRow(&frozen)
	fmt.Println(frozen.FrozenMoney)
	return frozen, qs
}

// GetAllFrozenfunds retrieves all Frozenfunds matches certain condition. Returns empty list if
// no records exist
func GetAllFrozenfunds(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Frozenfunds))
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

	var l []Frozenfunds
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

// UpdateFrozenfunds updates Frozenfunds by Id and returns error if
// the record to be updated doesn't exist
func UpdateFrozenfundsById(m *Frozenfunds) (err error) {
	o := orm.NewOrm()
	v := Frozenfunds{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFrozenfunds deletes Frozenfunds by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFrozenfunds(id int) (err error) {
	o := orm.NewOrm()
	v := Frozenfunds{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Frozenfunds{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
