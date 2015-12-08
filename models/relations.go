package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Relations struct {
	Id          int       `orm:"column(PKId);auto"`
	FrontUserId int       `orm:"column(FrontUserId)"`
	AfterUserId int       `orm:"column(AfterUserId)"`
	SetDate     time.Time `orm:"column(SetDate);type(datetime)"`
	Sources     string    `orm:"column(Sources);size(50)"`
}

//老师查看谁看过我
type RelationsByTid struct {
	Id           int       `orm:"column(PKId);auto"`
	FrontUserId  int       `orm:"column(FrontUserId)"`
	AfterUserId  int       `orm:"column(AfterUserId)"`
	SetDate      time.Time `orm:"column(SetDate);type(datetime)"`
	Sources      string    `orm:"column(Sources);size(50)"`
	UserName     string    `orm:"column(UserName);size(50);null"`
	AvatarPath   string    `orm:"column(AvatarPath);size(200);null"`
	AgeName      string    `orm:"column(AgeName);size(50);null"`
	SchoolName   string    `orm:"column(SchoolName);size(50);null"`
	IdentityName string    `orm:"column(IdentityName);size(50);null"`
}

//学生查看关注的老师
type RelationsByUid struct {
	Id          int       `orm:"column(PKId);auto"`
	FrontUserId int       `orm:"column(FrontUserId)"`
	AfterUserId int       `orm:"column(AfterUserId)"`
	SetDate     time.Time `orm:"column(SetDate);type(datetime)"`
	Sources     string    `orm:"column(Sources);size(50)"`
	UserName    string    `orm:"column(UserName);size(50);null"`
	CourseName  string    `orm:"column(CourseName);size(50);null"`
	SchoolName  string    `orm:"column(SchoolName);size(50);null"`
	AllDate     string    `orm:"column(AllDate)"`
}

//学生查看我浏览过谁
type RelationsByUidSee struct {
	Id          int       `orm:"column(PKId);auto"`
	FrontUserId int       `orm:"column(FrontUserId)"`
	AfterUserId int       `orm:"column(AfterUserId)"`
	SetDate     time.Time `orm:"column(SetDate);type(datetime)"`
	Sources     string    `orm:"column(Sources);size(50)"`
	UserName    string    `orm:"column(UserName);size(50);null"`
	AvatarPath  string    `orm:"column(AvatarPath);size(200);null"`
	UnitPrice   float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	GradeName   string    `orm:"column(GradeName);size(50);null"`
	CourseName  string    `orm:"column(CourseName);size(50);null"`
}

func (t *Relations) TableName() string {
	return "relations"
}

func init() {
	orm.RegisterModel(new(Relations))
}

//	查询一条师生某个关系信息
//	2015-11-25
func GetRelationsByST(sid int, tid int, guanxi string) (v *Relations, err error) {
	o := orm.NewOrm()
	var rela Relations
	err = o.QueryTable("relations").Filter("FrontUserId", tid).Filter("AfterUserId", sid).Filter("Sources", guanxi).One(&rela)
	v = &rela
	if err == nil {
		return v, nil
	}
	return nil, err
}

//    9.老师查看谁看过我
//    2015-11-06
func GetRelationsByTid(userid int, guanxi string, rows int, counts int) (relationsd []RelationsByTid, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationByTeacher+limitSql, userid, guanxi, rows, counts)
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return relationsd, qs
	}
	return
}

//    9.老师查看谁看过我总条数
//    2015-11-06
func GetRelationsByTidCount(userid int, guanxi string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationByTeacher, userid, guanxi)
	var relationsd []RelationsByTid
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(relationsd), qs
	}
	return
}

//    21.查询学生全部关注的老师
//    2015-11-06
func GetRelationsByUid(userid int, guanxi string, rows int, counts int) (relationsd []RelationsByUid, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationsBySGuanZhuT+limitSql, userid, guanxi, rows, counts)
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return relationsd, qs
	}
	return
}

//    21.查询学生全部关注的老师总条数
//    2015-11-06
func GetRelationsByUidCount(userid int, guanxi string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationsBySGuanZhuT, userid, guanxi)
	var relationsd []RelationsByUid
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(relationsd), qs
	}
	return
}

//    31.学生查看我浏览过谁
//    2015-11-06
func GetRelationsByUidSee(userid int, guanxi string, rows int, counts int) (relationsd []RelationsByUidSee, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationsByUidSee+limitSql, userid, guanxi, rows, counts)
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return relationsd, qs
	}
	return
}

//    31.学生查看我浏览过谁总条数
//    2015-11-06
func GetRelationsByUidSeeCount(userid int, guanxi string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlRelationsByUidSee, userid, guanxi)
	var relationsd []RelationsByUidSee
	num, qs := rs.QueryRows(&relationsd)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(relationsd), qs
	}
	return
}

// AddRelations insert a new Relations into database and returns
// last inserted Id on success.
func AddRelations(m *Relations) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRelationsById retrieves Relations by Id. Returns error if
// Id doesn't exist
func GetRelationsById(id int) (v *Relations, err error) {
	o := orm.NewOrm()
	v = &Relations{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRelations retrieves all Relations matches certain condition. Returns empty list if
// no records exist
func GetAllRelations(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Relations))
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

	var l []Relations
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

// UpdateRelations updates Relations by Id and returns error if
// the record to be updated doesn't exist
func UpdateRelationsById(m *Relations) (err error) {
	o := orm.NewOrm()
	v := Relations{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRelations deletes Relations by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRelations(id int) (err error) {
	o := orm.NewOrm()
	v := Relations{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Relations{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
