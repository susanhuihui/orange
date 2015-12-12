package models

import (
	"bytes"
	"math/rand"

	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Verification struct {
	Id         int       `orm:"column(PKId);auto"`
	UserPhone  string    `orm:"column(UserPhone);size(50);null"`
	VerCode    string    `orm:"column(VerCode);size(50);null"`
	CodePath   string    `orm:"column(CodePath);size(200);null"`
	CreateTime time.Time `orm:"column(CreateTime);type(datetime);null"`
}

func (t *Verification) TableName() string {
	return "verification"
}

func init() {
	orm.RegisterModel(new(Verification))
}

// AddVerification insert a new Verification into database and returns
// last inserted Id on success.
func AddVerification(m *Verification) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVerificationById retrieves Verification by Id. Returns error if
// Id doesn't exist
func GetVerificationById(id int) (v *Verification, err error) {
	o := orm.NewOrm()
	v = &Verification{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//    36.根据手机号码获取一条最新的验证码信息
//    2015-11-19
func GetVerificationByPhone(phone string) (ver Verification, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlVerificationByPhone, phone)
	qs := rs.QueryRow(&ver)
	return ver, qs
}

//	查询手机号下的所有验证码信息
//	2015-12-12
func GetVerificationListByPhone(phone string) (verif []Verification, err error) {
	o := orm.NewOrm()
	var allver []Verification
	_, errget := o.QueryTable("verification").Filter("UserPhone", phone).All(&allver)
	verif = allver
	if errget == nil {
		return verif, nil
	}
	return nil, err
}

// GetAllVerification retrieves all Verification matches certain condition. Returns empty list if
// no records exist
func GetAllVerification(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Verification))
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

	var l []Verification
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

// UpdateVerification updates Verification by Id and returns error if
// the record to be updated doesn't exist
func UpdateVerificationById(m *Verification) (err error) {
	o := orm.NewOrm()
	v := Verification{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVerification deletes Verification by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVerification(id int) (err error) {
	o := orm.NewOrm()
	v := Verification{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Verification{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取验证码
func GetVerifications(id string) (err error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	apikey := "1df965aee8d6301a7748add78a9637e3"
	phone := id
	text := "您的验证码是" + vcode + ",欢迎使用【泛鲲教育】"

	uisss := "apikey=" + apikey + "&text=" + text + "&mobile=" + phone
	body := bytes.NewBuffer([]byte(uisss))
	var verif Verification
	verif.CreateTime = time.Now()
	verif.UserPhone = id
	verif.VerCode = vcode
	_, err = AddVerification(&verif)
	if err == nil {
		resp, errs := http.Post("http://yunpian.com/v1/sms/send.json", "application/x-www-form-urlencoded", body)
		if errs != nil {
			log.Fatal(errs)
			return
		}
		_, errf := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if errf != nil {
			fmt.Println(errf.Error())
			log.Fatal(errf)
			return
		}
	}
	return err
}
