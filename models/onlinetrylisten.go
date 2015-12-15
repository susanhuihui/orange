package models

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/MsloveDl/bbb4go"
	"github.com/astaxie/beego/orm"
)

type Onlinetrylisten struct {
	Id           int       `orm:"column(PKId);auto"`
	Tid          int       `orm:"column(Tid)"`
	Sid          int       `orm:"column(Sid);null"`
	StartTime    time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime      time.Time `orm:"column(EndTime);type(datetime);null"`
	StuStartTime time.Time `orm:"column(StuStartTime);type(datetime);null"`
	StuEndTime   time.Time `orm:"column(StuEndTime);type(datetime);null"`
	ClassroomId  string    `orm:"column(ClassroomId);size(50);null"`
	StudentInId  string    `orm:"column(StudentInId);size(50);null"`
	TeacherInId  string    `orm:"column(TeacherInId);size(50);null"`
}

type OnlinetrylistenList struct {
	Id           int       `orm:"column(PKId);auto"`
	Tid          int       `orm:"column(Tid)"`
	Sid          int       `orm:"column(Sid);null"`
	StartTime    time.Time `orm:"column(StartTime);type(datetime);null"`
	EndTime      time.Time `orm:"column(EndTime);type(datetime);null"`
	StuStartTime time.Time `orm:"column(StuStartTime);type(datetime);null"`
	StuEndTime   time.Time `orm:"column(StuEndTime);type(datetime);null"`
	ClassroomId  string    `orm:"column(ClassroomId);size(50);null"`
	StudentInId  string    `orm:"column(StudentInId);size(50);null"`
	TeacherInId  string    `orm:"column(TeacherInId);size(50);null"`
	UserName     string    `orm:"column(UserName);size(50);null"`
}

func (t *Onlinetrylisten) TableName() string {
	return "onlinetrylisten"
}

func init() {
	orm.RegisterModel(new(Onlinetrylisten))
}

// AddOnlinetrylisten insert a new Onlinetrylisten into database and returns
// last inserted Id on success.
func AddOnlinetrylisten(m *Onlinetrylisten) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnlinetrylistenById retrieves Onlinetrylisten by Id. Returns error if
// Id doesn't exist
func GetOnlinetrylistenById(id int) (v *Onlinetrylisten, err error) {
	o := orm.NewOrm()
	v = &Onlinetrylisten{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//    41.查询老师的试听信息
//    2015-12-01
func OnlineTryListenByTid(userid int, rows int, counts int) (onlinetry []OnlinetrylistenList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw((SqlOnlineTryListenByTid + limitSql), userid, rows, counts)
	num, qs := rs.QueryRows(&onlinetry)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinetry, qs
	}
	return
}

//    41.查询老师的试听信息总条数
//    2015-12-01
func OnlineTryListenByTidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw((SqlOnlineTryListenByTid), userid)
	var onlinetry []OnlinetrylistenList
	num, qs := rs.QueryRows(&onlinetry)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinetry), qs
	}
	return
}

//	根据老师id查询一条信息
//	2015-12-01
func GetOnlinetrylistenOneByTid(tid int) (trylisten Onlinetrylisten, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlinetrylistenOn, tid)
	qs := rs.QueryRow(&trylisten)
	return trylisten, qs
}

// GetAllOnlinetrylisten retrieves all Onlinetrylisten matches certain condition. Returns empty list if
// no records exist
func GetAllOnlinetrylisten(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Onlinetrylisten))
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

	var l []Onlinetrylisten
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

// UpdateOnlinetrylisten updates Onlinetrylisten by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnlinetrylistenById(m *Onlinetrylisten) (err error) {
	o := orm.NewOrm()
	v := Onlinetrylisten{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOnlinetrylisten deletes Onlinetrylisten by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinetrylisten(id int) (err error) {
	o := orm.NewOrm()
	v := Onlinetrylisten{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Onlinetrylisten{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 老师为试听创建房间
// trylistenid:当前试听课程主键id
func GeListentecherlession(trylistenid int) (urlse string, err error) {
	slt := CreateStrVerify
	var name string
	var meetingID string
	var attendeePW string
	var moderatorPW string
	//获取预约课程信息
	trylistenmodel, _ := GetOnlinetrylistenById(trylistenid)
	fmt.Println("临时课程：")
	fmt.Println(trylistenmodel)
	//获取老师信息
	userinfo, _ := GetUserinformationTeacher(trylistenmodel.Tid)
	fmt.Println("老师信息")
	fmt.Println(userinfo)
	//name = userinfo.UserName //老师姓名
	name = "teacher"
	if trylistenmodel.ClassroomId+"" != "" {
		fmt.Println("课堂密码：" + trylistenmodel.ClassroomId)
		meetingID = trylistenmodel.ClassroomId
		attendeePW = trylistenmodel.StudentInId
		moderatorPW = trylistenmodel.TeacherInId
	} else {
		fmt.Println(getcodes("1"))
		meetingID = getcodes("1")   //房间主键随机数id
		attendeePW = getcodes("2")  //学生进入课程密码
		moderatorPW = getcodes("3") //老师进入课程密码
		//修改临时课程信息数据
		trylistenmodel.ClassroomId = meetingID
		trylistenmodel.StudentInId = attendeePW
		trylistenmodel.TeacherInId = moderatorPW
		err := UpdateOnlinetrylistenById(trylistenmodel) //预约信息中添加密码
		fmt.Println(err)
	}
	Duration := "180"      //课程总时长
	logoutURL := OnlineUrl //退出时返回到路径
	url := OnlineClassUrl  //请求创建的路径
	//参数处理
	sturl := "name=" + name + "&meetingID=" + meetingID + "&attendeePW=" + attendeePW + "&moderatorPW=" + moderatorPW + "&duration=" + Duration + "&logoutURL=" + logoutURL
	//生成验证码
	pd := bs("create" + sturl + slt)
	fmt.Println(url + "?" + sturl + "&checksum=" + pd)
	request, err := http.Get(url + "?" + sturl + "&checksum=" + pd) //get请求，请求创建课程

	if err != nil {
		fmt.Println(err.Error())
	} else {
		urls, errs := ioutil.ReadAll(request.Body)
		urlses := string(urls)
		if errs != nil {
			fmt.Println("1")
			fmt.Println(err.Error())
		} else {
			fmt.Println("2")
			fmt.Println(urlses)
		}
		//老师加入课堂路径
		urlt := OnlineInClassUrl
		sturls := "meetingID=" + meetingID + "&fullName=" + "Teacher" + "&password=" + moderatorPW //参数
		pds := bs("join" + sturls + slt)                                                           //生成加入验证码
		fmt.Println(urlt + "?" + sturls + "&checksum=" + pds)
		urlse = urlt + "?" + sturls + "&checksum=" + pds //
	}
	defer request.Body.Close()
	return
}

//学生查询在线白板的密码并进入课堂
func GetListenStudentlession(listenid int) (urlse string, err error) {
	slt := CreateStrVerify
	var name string      //学生姓名
	var meetingID string //房间主键
	var pwd string       //学生进入密码

	//获取预约课程信息
	onlinetrylisten, _ := GetOnlinetrylistenById(listenid)

	name = "Student"
	meetingID = onlinetrylisten.ClassroomId
	pwd = onlinetrylisten.StudentInId

	url := OnlineInClassUrl
	sturl := "meetingID=" + meetingID + "&fullName=" + name + "&password=" + pwd

	pd := bs("join" + sturl + slt)

	request, _ := http.Get(url + "?" + sturl + "&checksum=" + pd)
	urlse = url + "?" + sturl + "&checksum=" + pd
	fmt.Println("lujingshi")
	fmt.Println(urlse)
	defer request.Body.Close()
	return
}

//获取随机数方法
func getcodes(index string) (vcode string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000)) //自动生成随机数
	return vcode + index
}
func bs(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 重新编写--------------------------------------老师为试听创建房间--------------------------------------
// trylistenid:当前试听课程主键id
func GeListentecherlession2(trylistenid int) (urlse string, err error) {
	//生成三个随机数
	var meetingID string = getcode("1")
	var attendeePW string = getcode("2")  //学生进入密码
	var moderatorPW string = getcode("3") //老师进入密码
	//获取试听信息
	onlinetryclass, tryerr := GetOnlinetrylistenById(trylistenid)

	if onlinetryclass != nil && tryerr == nil {
		//每次老师开启试听进入课堂时都要建立一个新的课堂
		meetingroom := bbb4go.MeetingRoom{}
		meetingroom.Name_ = "泛鲲教育第一教室"
		meetingroom.MeetingID_ = meetingID
		meetingroom.AttendeePW_ = attendeePW
		meetingroom.ModeratorPW_ = moderatorPW
		meetingroom.Welcome = "欢迎来到泛鲲教育第一教室"
		meetingroom.LogoutURL = "http://" + OnlineUrl //试听结束进入首页+/orange/Teacher/ClassOverHtml/
		meetingroom.Duration = 0
		meetingroom.AllowStartStopRecording = false
		//将白板信息保存到数据库试听信息中
		onlinetryclass.StartTime = time.Now()
		onlinetryclass.ClassroomId = meetingID
		onlinetryclass.StudentInId = attendeePW
		onlinetryclass.TeacherInId = moderatorPW
		uperr := UpdateOnlinetrylistenById(onlinetryclass) //预约信息中添加密码
		if uperr == nil {
			meetingroom.CreateMeeting()
			if meetingroom.CreateMeetingResponse.Returncode == "SUCCESS" {
				//获取老师信息
				userinfo, gerr := GetUserinformationTeacher(onlinetryclass.Tid)
				if gerr == nil && userinfo.Id > 0 {
					//创建白板成功后创建一个老师并生成老师进入课堂的URL
					partTeacher := bbb4go.Participants{}
					partTeacher.IsAdmin_ = 1
					partTeacher.FullName_ = userinfo.UserName + "老师"
					partTeacher.MeetingID_ = meetingID                                    //教室id
					partTeacher.Password_ = moderatorPW                                   //老师进入密码
					partTeacher.CreateTime = meetingroom.CreateMeetingResponse.CreateTime //与创建教室时时间一致
					partTeacher.UserID = strconv.Itoa(userinfo.Id)
					partTeacher.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath //http://www.fankunedu.com/images/PersonHeadImg/yanyan.png
					partTeacher.GetJoinURL()
					urlse = partTeacher.JoinURL
				}
			}
		}
	}
	return
}

//重写学生进入课堂--------------学生查询在线白板的密码并进入课堂
func GetListenStudentlession2(listenid int, sid int) (urlse string, err error) {
	onlinetrylisten, geterr := GetOnlinetrylistenById(listenid)
	if geterr == nil && onlinetrylisten.Id > 0 {
		//获取学生信息
		userinfo, gerr := GetUserinformationTeacher(sid)
		fmt.Println(userinfo)
		if gerr == nil {
			urlse = "0"
			pardStudent := bbb4go.Participants{}
			pardStudent.IsAdmin_ = 0
			pardStudent.FullName_ = userinfo.UserName
			pardStudent.MeetingID_ = onlinetrylisten.ClassroomId //教室id
			pardStudent.Password_ = onlinetrylisten.StudentInId  //学生进入密码
			//pardStudent.CreateTime = time.Now().Format("2006-01-02 03:04:05 PM")
			pardStudent.UserID = strconv.Itoa(userinfo.Id)
			pardStudent.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath
			//判断此教室现在是否有老师
			meetroom := bbb4go.MeetingRoom{}
			meetroom.MeetingID_ = onlinetrylisten.ClassroomId
			meetroom.ModeratorPW_ = onlinetrylisten.TeacherInId
			meetroom.GetMeetingInfo()
			var pcount = meetroom.MeetingInfo.ParticipantCount
			fmt.Println("查看有几个人在线：")
			fmt.Println(pcount)
			if pcount == 0 { //老师不在
				urlse = "0"
			} else if pcount >= 2 { //已有人在试听
				urlse = "1"
			} else if pcount == 1 {
				//获取当前在线人的主键id
				attent := meetroom.MeetingInfo.Attendees
				onlinetid := attent.Attendees[0].UserID
				if onlinetid == strconv.Itoa(onlinetrylisten.Tid) {
					//添加学生一条试听信息
					var addonlinetry Onlinetrylisten
					addonlinetry.Tid = onlinetrylisten.Tid
					addonlinetry.Sid = userinfo.Id
					addonlinetry.StartTime = onlinetrylisten.StartTime
					addonlinetry.StuStartTime = time.Now()
					addonlinetry.ClassroomId = onlinetrylisten.ClassroomId
					addonlinetry.StudentInId = onlinetrylisten.StudentInId
					addid, adderr := AddOnlinetrylisten(&addonlinetry)
					if addid > 0 && adderr == nil {
						pardStudent.GetJoinURL()
						urlse = pardStudent.JoinURL
					}
				}
			}
		}
	}
	return
}
