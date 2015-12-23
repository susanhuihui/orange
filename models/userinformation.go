package models

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//原始用户信息
type Userinformation struct {
	Id                int       `orm:"column(PKId);auto"`
	UserName          string    `orm:"column(UserName);size(50);null"`
	IphoneNum         string    `orm:"column(IphoneNum);size(50);null"`
	LoginPassword     string    `orm:"column(LoginPassword);size(20);null"`
	PayPassword       string    `orm:"column(PayPassword);size(20);null"`
	UserSex           string    `orm:"column(UserSex);size(10);null"`
	SchoolId          int       `orm:"column(SchoolId);null"`
	SchoolName        string    `orm:"column(SchoolName);size(50);null"`
	SchoolAgeId       int       `orm:"column(SchoolAgeId);null"`
	UserDegree        int       `orm:"column(UserDegree);null"`
	UserLocation      int       `orm:"column(UserLocation);null"`
	SeniorLocation    int       `orm:"column(SeniorLocation);null"`
	HighSchool        string    `orm:"column(HighSchool);size(50);null"`
	UserLevelId       int       `orm:"column(UserLevelId);null"`
	LevelYear         int       `orm:"column(LevelYear);null"`
	IdentityId        int       `orm:"column(IdentityId);null"`
	GradeId           int       `orm:"column(GradeId);null"`
	SchoolAgeIdT      string    `orm:"column(SchoolAgeIdT);size(50);null"`
	UnitPrice         float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	Professional      string    `orm:"column(Professional);size(50);null"`
	Mailbox           string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox     string    `orm:"column(ParentMailbox);size(50);null"`
	StudyDifficult    string    `orm:"column(StudyDifficult);size(200);null"`
	AvatarPath        string    `orm:"column(AvatarPath);size(200);null"`
	UserHobby         string    `orm:"column(UserHobby);size(200);null"`
	Remarks           string    `orm:"column(Remarks);size(200);null"`
	BriefIntroduction string    `orm:"column(BriefIntroduction);null"`
	RegisteredTime    time.Time `orm:"column(RegisteredTime);type(datetime);null"`
}

//用户图片轮换
type UserinformationPic struct {
	Id         int    `orm:"column(PKId);auto"`
	UserName   string `orm:"column(UserName);size(50);null"`
	AvatarPath string `orm:"column(AvatarPath);size(200);null"`
	GradeName  string `orm:"column(GradeName);size(50);null"`
	CourseName string `orm:"column(CourseName);size(50);null"`
	Counts     int    `orm:"column(counts);null"`
}

//老师模块全部老师
type UserinformationModels struct {
	Id                int       `orm:"column(PKId);auto"`
	UserName          string    `orm:"column(UserName);size(50);null"`
	IphoneNum         string    `orm:"column(IphoneNum);size(50);null"`
	LoginPassword     string    `orm:"column(LoginPassword);size(20);null"`
	PayPassword       string    `orm:"column(PayPassword);size(20);null"`
	UserSex           string    `orm:"column(UserSex);size(10);null"`
	SchoolId          int       `orm:"column(SchoolId);null"`
	SchoolName        string    `orm:"column(SchoolName);size(50);null"`
	SchoolAgeId       int       `orm:"column(SchoolAgeId);null"`
	UserDegree        int       `orm:"column(UserDegree);null"`
	UserLocation      int       `orm:"column(UserLocation);null"`
	SeniorLocation    int       `orm:"column(SeniorLocation);null"`
	HighSchool        string    `orm:"column(HighSchool);size(50);null"`
	UserLevelId       int       `orm:"column(UserLevelId);null"`
	LevelYear         int       `orm:"column(LevelYear);null"`
	IdentityId        int       `orm:"column(IdentityId);null"`
	GradeId           int       `orm:"column(GradeId);null"`
	SchoolAgeIdT      string    `orm:"column(SchoolAgeIdT);size(50);null"`
	UnitPrice         float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	Professional      string    `orm:"column(Professional);size(50);null"`
	Mailbox           string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox     string    `orm:"column(ParentMailbox);size(50);null"`
	StudyDifficult    string    `orm:"column(StudyDifficult);size(200);null"`
	AvatarPath        string    `orm:"column(AvatarPath);size(200);null"`
	UserHobby         string    `orm:"column(UserHobby);size(200);null"`
	Remarks           string    `orm:"column(Remarks);size(200);null"`
	BriefIntroduction string    `orm:"column(BriefIntroduction);null"`
	RegisteredTime    time.Time `orm:"column(RegisteredTime);type(datetime);null"`
	DegreeName        string    `orm:"column(DegreeName);size(50);null"`
	CourseNameZhu     string    `orm:"column(CourseNameZhu);size(50);null"`
	CourseNameFu      string    `orm:"column(CourseNameFu);size(50);null"`
	SortCondition     int       `orm:"column(SortCondition);null"`
	OnlineState       int       `orm:"column(OnlineState);null"`
}

//老师个人信息详情
type UserinformationTeacher struct {
	Id                int       `orm:"column(PKId);auto"`
	UserName          string    `orm:"column(UserName);size(50);null"`
	IphoneNum         string    `orm:"column(IphoneNum);size(50);null"`
	LoginPassword     string    `orm:"column(LoginPassword);size(20);null"`
	PayPassword       string    `orm:"column(PayPassword);size(20);null"`
	UserSex           string    `orm:"column(UserSex);size(10);null"`
	SchoolId          int       `orm:"column(SchoolId);null"`
	SchoolName        string    `orm:"column(SchoolName);size(50);null"`
	SchoolAgeId       int       `orm:"column(SchoolAgeId);null"`
	UserDegree        int       `orm:"column(UserDegree);null"`
	UserLocation      int       `orm:"column(UserLocation);null"`
	SeniorLocation    int       `orm:"column(SeniorLocation);null"`
	HighSchool        string    `orm:"column(HighSchool);size(50);null"`
	UserLevelId       int       `orm:"column(UserLevelId);null"`
	LevelYear         int       `orm:"column(LevelYear);null"`
	IdentityId        int       `orm:"column(IdentityId);null"`
	GradeId           int       `orm:"column(GradeId);null"`
	SchoolAgeIdT      string    `orm:"column(SchoolAgeIdT);size(50);null"`
	UnitPrice         float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	Professional      string    `orm:"column(Professional);size(50);null"`
	Mailbox           string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox     string    `orm:"column(ParentMailbox);size(50);null"`
	StudyDifficult    string    `orm:"column(StudyDifficult);size(200);null"`
	AvatarPath        string    `orm:"column(AvatarPath);size(200);null"`
	UserHobby         string    `orm:"column(UserHobby);size(200);null"`
	Remarks           string    `orm:"column(Remarks);size(200);null"`
	BriefIntroduction string    `orm:"column(BriefIntroduction);null"`
	RegisteredTime    time.Time `orm:"column(RegisteredTime);type(datetime);null"`
	IdentityName      string    `orm:"column(IdentityName);size(50);null"`
	AllPerson         int       `orm:"column(AllPerson);null"`
	AllDate           int       `orm:"column(AllDate);null"`
	AllCount          int       `orm:"column(AllCount);null"`
	CourseName        string    `orm:"column(CourseName);size(50);null"`
	CourseNameId      int       `orm:"column(CourseNameId);null"`
	DegreeName        string    `orm:"column(DegreeName);size(50);null"`
	GradeName         string    `orm:"column(GradeName);size(50);null"`
}

//老师模块的，老师个人信息详情
type UserinformationTeacherModu struct {
	Id                int       `orm:"column(PKId);auto"`
	UserName          string    `orm:"column(UserName);size(50);null"`
	IphoneNum         string    `orm:"column(IphoneNum);size(50);null"`
	LoginPassword     string    `orm:"column(LoginPassword);size(20);null"`
	PayPassword       string    `orm:"column(PayPassword);size(20);null"`
	UserSex           string    `orm:"column(UserSex);size(10);null"`
	SchoolId          int       `orm:"column(SchoolId);null"`
	SchoolName        string    `orm:"column(SchoolName);size(50);null"`
	SchoolAgeId       int       `orm:"column(SchoolAgeId);null"`
	UserDegree        int       `orm:"column(UserDegree);null"`
	UserLocation      int       `orm:"column(UserLocation);null"`
	SeniorLocation    int       `orm:"column(SeniorLocation);null"`
	HighSchool        string    `orm:"column(HighSchool);size(50);null"`
	UserLevelId       int       `orm:"column(UserLevelId);null"`
	LevelYear         int       `orm:"column(LevelYear);null"`
	IdentityId        int       `orm:"column(IdentityId);null"`
	GradeId           int       `orm:"column(GradeId);null"`
	SchoolAgeIdT      string    `orm:"column(SchoolAgeIdT);size(50);null"`
	UnitPrice         float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	Professional      string    `orm:"column(Professional);size(50);null"`
	Mailbox           string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox     string    `orm:"column(ParentMailbox);size(50);null"`
	StudyDifficult    string    `orm:"column(StudyDifficult);size(200);null"`
	AvatarPath        string    `orm:"column(AvatarPath);size(200);null"`
	UserHobby         string    `orm:"column(UserHobby);size(200);null"`
	Remarks           string    `orm:"column(Remarks);size(200);null"`
	BriefIntroduction string    `orm:"column(BriefIntroduction);null"`
	RegisteredTime    time.Time `orm:"column(RegisteredTime);type(datetime);null"`
	AllPerson         int       `orm:"column(AllPerson);null"`
	AllTime           int       `orm:"column(AllTime);null"`
	AllTimeMouth      int       `orm:"column(AllTimeMouth);null"`
	DegreeName        string    `orm:"column(DegreeName);size(50);null"`
	CourseName        string    `orm:"column(CourseName);size(50);null"`
	CourseNameFu      string    `orm:"column(CourseNameFu);size(50);null"`
}

//学生个人信息详情
type UserinformationStudent struct {
	Id                int       `orm:"column(PKId);auto"`
	UserName          string    `orm:"column(UserName);size(50);null"`
	IphoneNum         string    `orm:"column(IphoneNum);size(50);null"`
	LoginPassword     string    `orm:"column(LoginPassword);size(20);null"`
	PayPassword       string    `orm:"column(PayPassword);size(20);null"`
	UserSex           string    `orm:"column(UserSex);size(10);null"`
	SchoolId          int       `orm:"column(SchoolId);null"`
	SchoolName        string    `orm:"column(SchoolName);size(50);null"`
	SchoolAgeId       int       `orm:"column(SchoolAgeId);null"`
	UserDegree        int       `orm:"column(UserDegree);null"`
	UserLocation      int       `orm:"column(UserLocation);null"`
	SeniorLocation    int       `orm:"column(SeniorLocation);null"`
	HighSchool        string    `orm:"column(HighSchool);size(50);null"`
	UserLevelId       int       `orm:"column(UserLevelId);null"`
	LevelYear         int       `orm:"column(LevelYear);null"`
	IdentityId        int       `orm:"column(IdentityId);null"`
	GradeId           int       `orm:"column(GradeId);null"`
	SchoolAgeIdT      string    `orm:"column(SchoolAgeIdT);size(50);null"`
	UnitPrice         float64   `orm:"column(UnitPrice);null;digits(10);decimals(2)"`
	Professional      string    `orm:"column(Professional);size(50);null"`
	Mailbox           string    `orm:"column(Mailbox);size(50);null"`
	ParentMailbox     string    `orm:"column(ParentMailbox);size(50);null"`
	StudyDifficult    string    `orm:"column(StudyDifficult);size(200);null"`
	AvatarPath        string    `orm:"column(AvatarPath);size(200);null"`
	UserHobby         string    `orm:"column(UserHobby);size(200);null"`
	Remarks           string    `orm:"column(Remarks);size(200);null"`
	BriefIntroduction string    `orm:"column(BriefIntroduction);null"`
	RegisteredTime    time.Time `orm:"column(RegisteredTime);type(datetime);null"`
	IdentityName      string    `orm:"column(IdentityName);size(50);null"`
	AllPerson         int       `orm:"column(AllPerson);null"`
	AllDate           int       `orm:"column(AllDate);null"`
	AllCount          int       `orm:"column(AllCount);null"`
	AgeName           string    `orm:"column(AgeName);size(50);null"`
}

func (t *Userinformation) TableName() string {
	return "userinformation"
}

func init() {
	orm.RegisterModel(new(Userinformation))
}

// AddUserinformation insert a new Userinformation into database and returns
// last inserted Id on success.
func AddUserinformation(m *Userinformation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserinformationById retrieves Userinformation by Id. Returns error if
// Id doesn't exist
func GetUserinformationById(id int) (v *Userinformation, err error) {
	o := orm.NewOrm()
	v = &Userinformation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//	根据用户名密码查询一条用户信息
//	2015-11-02
func GetUserinformationLogin(username string, password string) (v *Userinformation, err error) {
	o := orm.NewOrm()
	var user Userinformation
	err = o.QueryTable("userinformation").Filter("UserName", username).Filter("LoginPassword", password).One(&user)
	v = &user
	if err == nil {
		return v, nil
	}
	return nil, err
}

//	根据手机号密码查询一条用户信息
//	2015-11-27
func GetUserinformationLoginPhone(phone string, password string) (v *Userinformation, err error) {
	o := orm.NewOrm()
	var user Userinformation
	err = o.QueryTable("userinformation").Filter("IphoneNum", phone).Filter("LoginPassword", password).One(&user)
	v = &user
	if err == nil {
		return v, nil
	}
	return nil, err
}

//	根据手机号查询一条用户信息
//	2015-11-27
func GetUserinformationByPhone(phone string) (v *Userinformation, err error) {
	o := orm.NewOrm()
	var user Userinformation
	err = o.QueryTable("userinformation").Filter("IphoneNum", phone).One(&user)
	v = &user
	if err == nil {
		return v, nil
	}
	return nil, err
}

//	根据姓名昵称查询一条用户信息
//	2015-12-09
func GetUserinformationByUserName(username string) (v *Userinformation, err error) {
	o := orm.NewOrm()
	var user Userinformation
	err = o.QueryTable("userinformation").Filter("UserName", username).One(&user)
	v = &user
	if err == nil {
		return v, nil
	}
	return nil, err
}

//	1.查询首页老师图片轮换
//	2015-11-03
func GetUserinformationPicMove(count int) (userlist []UserinformationPic, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserPicList, count)
	num, qs := rs.QueryRows(&userlist)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return userlist, qs
	}
	return
}

//	2.查询老师个人信息详情
//	2015-11-06
func GetUserinformationTeacher(userid int) (teacher UserinformationTeacher, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserTeacher, userid)
	qs := rs.QueryRow(&teacher)
	return teacher, qs
}

//	17.查询学生个人详细信息
//	2015-11-06
func GetUserinformationStudent(userid int) (student UserinformationStudent, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserInformationByS, userid)
	qs := rs.QueryRow(&student)
	return student, qs
}

//	27.检索老师全部信息
//	2015-11-03
func GetUserinformationAllTeacher(seltype int, nianji string, kecheng string, jibie string, shengfen string, shiqu string, rows int, counts int) (userlist []UserinformationModels, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	fmt.Println(shiqu)
	if seltype == 1 { //在线
		rs = o.Raw(SqlUserinformationAllTeacherByOnline+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	} else if seltype == 2 { //人气
		rs = o.Raw(SqlUserinformationAllTeacherByPerson+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	} else if seltype == 3 { //授课经验
		rs = o.Raw(SqlUserinformationAllTeacherByTime+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	}
	num, qs := rs.QueryRows(&userlist)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return userlist, qs
	}
	return
}

//	27.检索老师全部信息
//	2015-11-03
func GetUserinformationAllTeacher2(seltype int, nianji string, kecheng string, jibie string, shengfen string, shiqu string, rows int, counts int) (userlist []UserinformationModels, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	fmt.Println(shiqu)
	var SqlSchoolAge = ` and userinfo.schoolageidt like `
	var SqlClass = ` and (select coursesid from remedialcourses recs where recs.userid = userinfo.pkid and ismain = 1 limit 1) in (select cou.pkid from course as cou where cou.coursename like `
	var SqlLevel = ` and userinfo.UserLevelId in (select ulv.pkid from userlevel as ulv where ulv.levelname like `
	var SqlProvince = ` and userinfo.SeniorLocation in 
	    (select cit.pkid from citys as cit where cit.proid in (select prov.pkid from province as prov where prov.proname like `
	var selstr string = ""
	if seltype == 1 { //在线
		selstr = SqlUserinformationAllTeacherByPerson1
	} else if seltype == 2 { //人气
		selstr = SqlUserinformationAllTeacherByTime1
		//rs = o.Raw(SqlUserinformationAllTeacherByPerson+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	} else if seltype == 3 { //授课经验
		selstr = SqlUserinformationAllTeacherByOnline1
		//rs = o.Raw(SqlUserinformationAllTeacherByTime+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	}
	if nianji != `%%` {
		selstr = selstr + SqlSchoolAge + `'` + nianji + `' `
	}
	if kecheng != `%%` {
		selstr = selstr + SqlClass + `'` + kecheng + `'` + `) `
	}
	if jibie != `%%` {
		selstr = selstr + SqlLevel + `'` + jibie + `'` + `) `
	}
	if shengfen != `%%` {
		selstr = selstr + SqlProvince + `'` + shengfen + `'` + `))	`
	}
	selstr = selstr + SqlUserOver
	fmt.Println("最终查询语句为：")
	fmt.Println(selstr + limitSql)
	rs = o.Raw((selstr + limitSql), rows, counts)
	num, qs := rs.QueryRows(&userlist)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return userlist, qs
	}
	return
}

//	27.检索老师全部信息总条数
//	2015-11-03
func GetUserinformationAllTeacherCount(seltype int, nianji string, kecheng string, jibie string, shengfen string, shiqu string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	if seltype == 1 {
		rs = o.Raw(SqlUserinformationAllTeacherByOnline, nianji, kecheng, jibie, shengfen)
	} else if seltype == 2 {
		rs = o.Raw(SqlUserinformationAllTeacherByPerson, nianji, kecheng, jibie, shengfen)
	} else if seltype == 3 {
		rs = o.Raw(SqlUserinformationAllTeacherByTime, nianji, kecheng, jibie, shengfen)
	}
	var userlist []UserinformationModels
	num, qs := rs.QueryRows(&userlist)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(userlist), qs
	}
	return
}

//	27.检索老师全部信息总条数
//	2015-11-03
func GetUserinformationAllTeacherCount2(seltype int, nianji string, kecheng string, jibie string, shengfen string, shiqu string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	var SqlSchoolAge = ` and userinfo.schoolageidt like `
	var SqlClass = ` and (select coursesid from remedialcourses recs where recs.userid = userinfo.pkid and ismain = 1 limit 1) in (select cou.pkid from course as cou where cou.coursename like `
	var SqlLevel = ` and userinfo.UserLevelId in (select ulv.pkid from userlevel as ulv where ulv.levelname like `
	var SqlProvince = ` and userinfo.SeniorLocation in 
	    (select cit.pkid from citys as cit where cit.proid in (select prov.pkid from province as prov where prov.proname like `
	var selstr string = ""
	if seltype == 1 { //在线
		selstr = SqlUserinformationAllTeacherByPerson1
	} else if seltype == 2 { //人气
		selstr = SqlUserinformationAllTeacherByTime1
		//rs = o.Raw(SqlUserinformationAllTeacherByPerson+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	} else if seltype == 3 { //授课经验
		selstr = SqlUserinformationAllTeacherByOnline1
		//rs = o.Raw(SqlUserinformationAllTeacherByTime+limitSql, nianji, kecheng, jibie, shengfen, rows, counts)
	}
	if nianji != `%%` {
		selstr = selstr + SqlSchoolAge + `'` + nianji + `' `
	}
	if kecheng != `%%` {
		selstr = selstr + SqlClass + `'` + kecheng + `'` + `) `
	}
	if jibie != `%%` {
		selstr = selstr + SqlLevel + `'` + jibie + `'` + `) `
	}
	if shengfen != `%%` {
		selstr = selstr + SqlProvince + `'` + shengfen + `'` + `))	`
	}
	selstr = selstr + SqlUserOver
	rs = o.Raw(selstr)
	var userlist []UserinformationModels
	num, qs := rs.QueryRows(&userlist)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(userlist), qs
	}
	return
}

//	28.老师模块：查看老师详情
//	2015-11-06
func GetUserinformationTeacherModu(userid int) (student UserinformationTeacherModu, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlUserinformationByTid, userid)
	qs := rs.QueryRow(&student)
	return student, qs
}

// GetAllUserinformation retrieves all Userinformation matches certain condition. Returns empty list if
// no records exist
func GetAllUserinformation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Userinformation))
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

	var l []Userinformation
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

// UpdateUserinformation updates Userinformation by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserinformationById(m *Userinformation) (err error) {
	o := orm.NewOrm()
	v := Userinformation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserinformation deletes Userinformation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserinformation(id int) (err error) {
	o := orm.NewOrm()
	v := Userinformation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Userinformation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//上传图片
func GetImganddata(request *http.Request, picpath string) (json string, imgstr string) {
	request.ParseMultipartForm(32 << 20)
	multipartForm := request.MultipartForm

	if multipartForm != nil {
		//获取图片文件
		files := multipartForm.File
		//获取表单数据
		values := multipartForm.Value
		fmt.Println(values)
		fmt.Println(files)
		if len(values["jsondata"]) > 0 {
			json = values["jsondata"][0]
		}
		if 0 == len(files) {
			log.Println("NO FILE.")

			return
		}

		for k := 1; k <= len(files); k++ {

			fileimg := files["file"+strconv.Itoa(k)]

			if fileimg != nil {
				for i := 0; i < len(fileimg); i++ {

					filename := fileimg[i].Filename
					str := strings.Split(filename, ".")
					name := strconv.Itoa(time.Now().Nanosecond()) + strconv.Itoa(i) + "." + str[1]
					if imgstr == "" {
						imgstr = picpath + name //"images/PersonHeadImg/"
					} else {
						imgstr = imgstr + "," + picpath + name
					}
					fileo, _ := fileimg[i].Open()
					defer fileo.Close()
					filePath, _ := exec.LookPath(os.Args[0])
					path := strings.Replace(filePath, "orange.exe", "", 1)
					outputFilePath := path + "views/" + picpath + name
					fileWriter, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE,
						0666)

					if nil != err {
						log.Println("FILE CREATE ERROR: ", err.Error())
						return
					}

					io.Copy(fileWriter, fileo)
					fileWriter.Close()
				}
			}
		}

	}
	return json, imgstr
}

//上传图片
func GetImganddata2(request *http.Request, headpath string) (json string, imgstr string) {
	request.ParseMultipartForm(32 << 20)

	multipartForm := request.MultipartForm

	if multipartForm != nil {
		//获取图片文件
		files := multipartForm.File
		//获取表单数据
		values := multipartForm.Value
		fmt.Println(values)
		fmt.Println(files)
		if len(values["jsondata"]) > 0 {
			json = values["jsondata"][0]
		}
		if 0 == len(files) {
			log.Println("NO FILE.")

			return
		} else if 4 < len(files) {
			log.Println("MAX FILE NUMBER.")

			return
		}
		fileimg := files["file"]

		for i := 0; i < len(fileimg); i++ {

			filename := fileimg[i].Filename
			str := strings.Split(filename, ".")
			name := strconv.Itoa(time.Now().Nanosecond()) + strconv.Itoa(i) + "." + str[1]
			if i == 0 {
				imgstr = headpath + name
			} else {
				imgstr = imgstr + "," + headpath + name
			}
			fileo, _ := fileimg[i].Open()
			defer fileo.Close()
			filePath, _ := exec.LookPath(os.Args[0])
			bendipath := strings.Replace(filePath, "orange.exe", "", 1)
			bendipath = "../orange"
			rs := []rune(bendipath)
			rl := len(rs)
			overs := string(rs[rl-1])
			fmt.Println(overs)
			var outputFilePath string
			if overs == `/` || overs == `\` {
				outputFilePath = bendipath + "views/" + headpath + name
			} else {
				outputFilePath = bendipath + "/views/" + headpath + name
			}
			fmt.Println("保存图片路径为：")
			fmt.Println(outputFilePath)
			fileWriter, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE,
				0666)

			if nil != err {
				log.Println("FILE CREATE ERROR: ", err.Error())
				return
			}

			io.Copy(fileWriter, fileo)
			fileWriter.Close()
		}
	}
	return json, imgstr
}

////更新头像和昵称
//func UpdateUserimg(m *Userinformation, imgstr string) (rownums int, err error) {
//	var img string
//	if imgstr != "" {
//		str := strings.Split(imgstr, ",")
//		img = str[0]
//	}
//	o := orm.NewOrm()
//	var rs orm.RawSeter
//	rs = o.Raw("call proc_UpdateUserPortraitByUserid(?,?)", m.Id, img)
//	qs := rs.QueryRow(&rownums)
//	return rownums, qs
//}
