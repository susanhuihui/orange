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

type Onlinecoursebooking struct {
	Id               int       `orm:"column(PKId);auto"`
	UserIdActive     int       `orm:"column(UserIdActive)"`
	UserIdPassive    int       `orm:"column(UserIdPassive)"`
	ReservationState int       `orm:"column(ReservationState)"`
	Payment          int       `orm:"column(Payment)"`
	Leaming          int       `orm:"column(Leaming)"`
	PayWay           int       `orm:"column(PayWay)"`
	StartTime        time.Time `orm:"column(StartTime);type(datetime)"`
	EndTime          time.Time `orm:"column(EndTime);type(datetime)"`
	AppointTime      time.Time `orm:"column(AppointTime);type(datetime)"`
	AppointMessage   string    `orm:"column(AppointMessage);null"`
	ClassroomId      string    `orm:"column(ClassroomId);size(50);null"`
	StudentInId      string    `orm:"column(StudentInId);size(50);null"`
	TeacherInId      string    `orm:"column(TeacherInId);size(50);null"`
}

//查询老师/学生预约课程信息
type OnlinecoursebookingList struct {
	Id               int       `orm:"column(PKId);auto"`
	UserIdActive     int       `orm:"column(UserIdActive)"`
	UserIdPassive    int       `orm:"column(UserIdPassive)"`
	ReservationState int       `orm:"column(ReservationState)"`
	Payment          int       `orm:"column(Payment)"`
	Leaming          int       `orm:"column(Leaming)"`
	PayWay           int       `orm:"column(PayWay)"`
	StartTime        time.Time `orm:"column(StartTime);type(datetime)"`
	EndTime          time.Time `orm:"column(EndTime);type(datetime)"`
	AppointTime      time.Time `orm:"column(AppointTime);type(datetime)"`
	AppointMessage   string    `orm:"column(AppointMessage);null"`
	ClassroomId      string    `orm:"column(ClassroomId);size(50);null"`
	StudentInId      string    `orm:"column(StudentInId);size(50);null"`
	TeacherInId      string    `orm:"column(TeacherInId);size(50);null"`
	CourseName       string    `orm:"column(CourseName);size(50);null"`
	UserName         string    `orm:"column(UserName);size(50);null"`
	IphoneNum        string    `orm:"column(IphoneNum);size(50);null"`
}

func (t *Onlinecoursebooking) TableName() string {
	return "onlinecoursebooking"
}

func init() {
	orm.RegisterModel(new(Onlinecoursebooking))
}

//    6.根据老师主键查询预约课程信息
//    2015-11-06
func GetOnlinecoursebookingByTid(userid int, rows int, counts int) (onlinebooking []OnlinecoursebookingList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineTeacherbookingByT+limitSql, userid, rows, counts)
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebooking, qs
	}
	return
}

//    6.根据老师主键查询预约课程信息总条数
//    2015-11-06
func GetOnlinecoursebookingByTidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineTeacherbookingByT, userid)
	var onlinebooking []OnlinecoursebookingList
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinebooking), qs
	}
	return
}

//    10.查询老师预约信息
//    2015-11-06
func GetOnlinecoursebookingByTidSimp(userid int, times string) (onlinebook []Onlinecoursebooking, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBookingByT, userid, times)
	num, qs := rs.QueryRows(&onlinebook)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebook, qs
	}
	return
}

//    20.根据学生主键查询预约课程信息
//    2015-11-06
func GetOnlinecoursebookingByUid(userid int, rows int, counts int) (onlinebooking []OnlinecoursebookingList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw((SqlOnlineTeacherbookingByS + limitSql), userid, rows, counts)
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebooking, qs
	}
	return
}

//    20.根据学生主键查询预约课程信息总条数
//    2015-11-06
func GetOnlinecoursebookingByUidCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineTeacherbookingByS, userid)
	var onlinebooking []OnlinecoursebookingList
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinebooking), qs
	}
	return
}

//    20.查询学生没有上过的预约课程
//    2015-12-04
func GetOnlinecoursebookingBySidNotOn(userid int, rows int, counts int) (onlinebooking []OnlinecoursebookingList, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw((SqlOnlinebookingBySidNotOn + limitSql), userid, rows, counts)
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebooking, qs
	}
	return
}

//    20.查询学生没有上过的预约课程总条数
//    2015-12-04
func GetOnlinecoursebookingBySidNotOnCount(userid int) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw((SqlOnlinebookingBySidNotOn), userid)
	var onlinebooking []OnlinecoursebookingList
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinebooking), qs
	}
	return
}

//    20.查询学生预约某个老师某天预约了几次课程
//    2015-12-05
func GetOnlinecoursebookingBySTidTime(sid int, tid int, time1 string, time2 string) (allcount int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlinebookingbySTidTime, sid, tid, time1, time2)
	var onlinebooking []OnlinecoursebookingList
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return 0, qs
	} else {
		return len(onlinebooking), qs
	}
	return
}

//    30.查询老师一段时间内的授课情况
//    2015-11-06
func GetOnlinecoursebookingByTimeTid(starttime string, endtime string, userid int) (onlinebooking []Onlinecoursebooking, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineCourseBookingTimeByTid, starttime, endtime, starttime, endtime, userid)
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebooking, qs
	}
	return
}

//	38.根据老师主键id，和时间段查询此时间段预约课程信息
//	2015-11-22
func GetOnlinecoursebookingByTidTime(userid int, time1 string) (onlinebooking []Onlinecoursebooking, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBookingByTidTime, userid, time1)
	num, qs := rs.QueryRows(&onlinebooking)
	if qs != nil {
		fmt.Printf("num", num)
		return nil, qs
	} else {
		return onlinebooking, qs
	}
	return
}

// AddOnlinecoursebooking insert a new Onlinecoursebooking into database and returns
// last inserted Id on success.
func AddOnlinecoursebooking(m *Onlinecoursebooking) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnlinecoursebookingById retrieves Onlinecoursebooking by Id. Returns error if
// Id doesn't exist
func GetOnlinecoursebookingById(id int) (v *Onlinecoursebooking, err error) {
	o := orm.NewOrm()
	v = &Onlinecoursebooking{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOnlinecoursebooking retrieves all Onlinecoursebooking matches certain condition. Returns empty list if
// no records exist
func GetAllOnlinecoursebooking(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Onlinecoursebooking))
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

	var l []Onlinecoursebooking
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

// UpdateOnlinecoursebooking updates Onlinecoursebooking by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnlinecoursebookingById(m *Onlinecoursebooking) (err error) {
	o := orm.NewOrm()
	fmt.Println(m)
	v := Onlinecoursebooking{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		fmt.Println("时间")
		fmt.Println(m.StartTime)
		fmt.Println(time.Now())
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
			fmt.Println(m)
		}
	}
	return
}

// DeleteOnlinecoursebooking deletes Onlinecoursebooking by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinecoursebooking(id int) (err error) {
	o := orm.NewOrm()
	v := Onlinecoursebooking{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Onlinecoursebooking{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取随机数方法
func getcode(index string) (vcode string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000)) //自动生成随机数
	return vcode + index
}

func Updatebookings(id int, classid string, sid string, tid string) (num int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	fmt.Println("参数为：")
	fmt.Println(id)
	fmt.Println(classid + sid + tid)
	var upstr string = `update onlinecoursebooking set ClassroomId='` + classid + `' ,StudentInId='` + sid + `' ,TeacherInId='` + tid + `' where PKId=` + strconv.Itoa(id) + `; SELECT ROW_COUNT() as roocount;`
	fmt.Println(upstr)
	rs = o.Raw(upstr)
	qs := rs.QueryRow(&num)
	fmt.Println(qs)
	return num, qs
}

// 老师创建房间
// onlineid:当前预约课程主键id
func Getecherlession(onlineid int) (urlse string, err error) {
	//	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	//	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000)) //自动生成随机数

	slt := CreateStrVerify
	var name string
	var meetingID string
	var attendeePW string
	var moderatorPW string
	//var welcome string = "welcome online class ,I am your teacher!" //欢迎语

	//获取预约课程信息
	onlinenow, _ := GetOnlinecoursebookingById(onlineid)
	fmt.Println("预约信息为1")
	fmt.Println(onlinenow)
	//获取老师信息
	userinfo, _ := GetUserinformationTeacher(onlinenow.UserIdPassive)
	fmt.Println("老师信息")
	fmt.Println(userinfo)
	//name = userinfo.UserName //老师姓名
	name = "teacher"
	if onlinenow.ClassroomId+"" != "" {
		meetingID = onlinenow.ClassroomId
		attendeePW = onlinenow.StudentInId
		moderatorPW = onlinenow.TeacherInId
	} else {
		meetingID = getcode("1")   //房间主键随机数id
		attendeePW = getcode("2")  //学生进入课程密码
		moderatorPW = getcode("3") //老师进入课程密码
		//修改预约课程信息数据

		onlinenow.ClassroomId = meetingID
		onlinenow.StudentInId = attendeePW
		onlinenow.TeacherInId = moderatorPW
		num, _ := Updatebookings(onlineid, meetingID, attendeePW, moderatorPW)
		fmt.Println(num)
		//err := UpdateOnlinecoursebookingById(onlinenow) //预约信息中添加密码
		fmt.Println(err)
	}
	Duration := "60"       //课程总时长
	logoutURL := OnlineUrl //退出时返回到路径
	url := OnlineClassUrl  //请求创建的路径
	fmt.Println("请求创建的路径:")
	fmt.Println(OnlineClassUrl)
	//参数处理
	sturl := "name=" + name + "&meetingID=" + meetingID + "&attendeePW=" + attendeePW + "&moderatorPW=" + moderatorPW + "&duration=" + Duration + "&logoutURL=" + logoutURL
	//生成验证码
	pd := b("create" + sturl + slt)
	fmt.Println(url + "?" + sturl + "&checksum=" + pd)
	request, err := http.Get(url + "?" + sturl + "&checksum=" + pd) //get请求，请求创建课程
	fmt.Println("错误信息")
	if err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Println("没有错误")
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
		pds := b("join" + sturls + slt)                                                            //生成加入验证码
		fmt.Println(urlt + "?" + sturls + "&checksum=" + pds)
		urlse = urlt + "?" + sturls + "&checksum=" + pds //

	}
	defer request.Body.Close()
	return
}

//学生查询在线白板的密码并进入课堂
func Getstudentlession(onlineid int) (urlse string, err error) {
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	//vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//	attendeePW := "123"

	slt := CreateStrVerify
	var name string      //学生姓名
	var meetingID string //房间主键
	var pwd string       //学生进入密码

	//获取预约课程信息
	onlinenow, _ := GetOnlinecoursebookingById(onlineid)
	//获取老师信息
	userinfo, _ := GetUserinformationStudent(onlinenow.UserIdPassive)
	fmt.Println(userinfo)
	//name = userinfo.UserName
	name = "Student"
	meetingID = onlinenow.ClassroomId
	pwd = onlinenow.StudentInId

	url := OnlineInClassUrl
	sturl := "meetingID=" + meetingID + "&fullName=" + name + "&password=" + pwd

	pd := b("join" + sturl + slt)

	request, _ := http.Get(url + "?" + sturl + "&checksum=" + pd)
	urlse = url + "?" + sturl + "&checksum=" + pd
	fmt.Println("lujingshi")
	fmt.Println(urlse)
	defer request.Body.Close()
	return
}

func b(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 老师开始讲课重写，如果房间已建立直接进入，否则重建
// onlineid:当前预约课程主键id
func Getecherlession2(onlineid int) (urlse string, err error) {
	//获取预约信息
	onlinenow, onlineerr := GetOnlinecoursebookingById(onlineid)
	//获取老师信息
	userinfo, useerr := GetUserinformationTeacher(onlinenow.UserIdPassive)
	if onlineerr == nil {
		if useerr == nil && userinfo.Id > 0 {
			if onlinenow.ClassroomId+"" != "" {
				//判断白板教室内是否有人
				meetroom := bbb4go.MeetingRoom{}
				meetroom.MeetingID_ = onlinenow.ClassroomId
				meetroom.ModeratorPW_ = onlinenow.TeacherInId
				meetroom.GetMeetingInfo()
				var pcount = meetroom.MeetingInfo.ParticipantCount
				if pcount >= 2 { //进入重复课堂
					//重新建立课堂并保存到数据库
					urlse = CreateMeeting(onlineid, userinfo)
				} else if pcount <= 1 {
					//创建白板成功后创建一个老师并生成老师进入课堂的URL
					partTeacher := bbb4go.Participants{}
					partTeacher.IsAdmin_ = 1
					partTeacher.FullName_ = userinfo.UserName + "老师"
					partTeacher.MeetingID_ = onlinenow.ClassroomId //教室id
					partTeacher.Password_ = onlinenow.TeacherInId  //老师进入密码
					//partTeacher.CreateTime = time.Now().Format("2006-01-02 03:04:05 PM") //
					partTeacher.UserID = strconv.Itoa(userinfo.Id)
					partTeacher.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath //http://www.fankunedu.com/images/PersonHeadImg/yanyan.png
					partTeacher.GetJoinURL()
					urlse = partTeacher.JoinURL
				}

			} else {
				//新建课堂并进入
				var meetingID string = getcode("1")
				var attendeePW string = getcode("2")  //学生进入密码
				var moderatorPW string = getcode("3") //老师进入密码
				meetingroom := bbb4go.MeetingRoom{}
				meetingroom.Name_ = "泛鲲教育第一教室"
				meetingroom.MeetingID_ = meetingID
				meetingroom.AttendeePW_ = attendeePW
				meetingroom.ModeratorPW_ = moderatorPW
				meetingroom.Welcome = "欢迎来到泛鲲教育第一教室"
				meetingroom.LogoutURL = "http://" + OnlineUrl + "/orange/Teacher/ClassOverHtml/" //试听结束进入课程结束页面
				meetingroom.Duration = 50
				meetingroom.AllowStartStopRecording = false
				num, _ := Updatebookings(onlineid, meetingID, attendeePW, moderatorPW) //更新数据库的白板信息
				fmt.Println(num)
				meetingroom.CreateMeeting()
				if meetingroom.CreateMeetingResponse.Returncode == "SUCCESS" {
					fmt.Println("新建课堂成功:")
					//创建白板成功后创建一个老师并生成老师进入课堂的URL
					partTeacher := bbb4go.Participants{}
					partTeacher.IsAdmin_ = 1
					partTeacher.FullName_ = userinfo.UserName + "老师"
					partTeacher.MeetingID_ = meetingID                                    //教室id
					partTeacher.Password_ = moderatorPW                                   //老师进入密码
					partTeacher.CreateTime = meetingroom.CreateMeetingResponse.CreateTime //与创建教室时时间一致
					partTeacher.UserID = strconv.Itoa(userinfo.Id)
					partTeacher.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath
					partTeacher.GetJoinURL()
					urlse = partTeacher.JoinURL
				} else {
					urlse = "-3" //创建课堂失败
				}
				//新建课堂
				//urlse = CreateMeeting(onlineid, userinfo)
			}
		} else {
			urlse = "-2" //用户不存在
		}
	} else {
		urlse = "-1" //预约信息不存在
	}
	return
}

//新建课堂方法
func CreateMeeting(onlineid int, userinfo UserinformationTeacher) (urlse string) {
	var meetingID string = getcode("1")
	var attendeePW string = getcode("2")  //学生进入密码
	var moderatorPW string = getcode("3") //老师进入密码
	meetingroom := bbb4go.MeetingRoom{}
	meetingroom.Name_ = "泛鲲教育第一教室"
	meetingroom.MeetingID_ = meetingID
	meetingroom.AttendeePW_ = attendeePW
	meetingroom.ModeratorPW_ = moderatorPW
	meetingroom.Welcome = "欢迎来到泛鲲教育第一教室"
	meetingroom.LogoutURL = "http://" + OnlineUrl + "/orange/Teacher/ClassOverHtml/" //试听结束进入课程结束页面
	meetingroom.Duration = 50
	meetingroom.AllowStartStopRecording = false
	num, _ := Updatebookings(onlineid, meetingID, attendeePW, moderatorPW) //更新数据库的白板信息
	fmt.Println(num)
	meetingroom.CreateMeeting()
	if meetingroom.CreateMeetingResponse.Returncode == "SUCCESS" {
		//创建白板成功后创建一个老师并生成老师进入课堂的URL
		partTeacher := bbb4go.Participants{}
		partTeacher.IsAdmin_ = 1
		partTeacher.FullName_ = userinfo.UserName + "老师"
		partTeacher.MeetingID_ = meetingID                                    //教室id
		partTeacher.Password_ = moderatorPW                                   //老师进入密码
		partTeacher.CreateTime = meetingroom.CreateMeetingResponse.CreateTime //与创建教室时时间一致
		partTeacher.UserID = strconv.Itoa(userinfo.Id)
		partTeacher.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath
		partTeacher.GetJoinURL()
		urlse = partTeacher.JoinURL
	} else {
		urlse = "-3" //创建课堂失败
	}
	return
}

//重新学生进入课堂的方法---------学生查询在线白板的密码并进入课堂
func Getstudentlession2(onlineid int) (urlse string, err error) {
	//获取预约课程信息
	onlinenow, onerr := GetOnlinecoursebookingById(onlineid)
	//获取老师信息
	userinfo, usererr := GetUserinformationStudent(onlinenow.UserIdPassive)
	if onerr == nil && onlinenow.Id > 0 {
		if usererr == nil && userinfo.Id > 0 {
			pardStudent := bbb4go.Participants{}
			pardStudent.IsAdmin_ = 0
			pardStudent.FullName_ = userinfo.UserName
			pardStudent.MeetingID_ = onlinenow.ClassroomId //教室id
			pardStudent.Password_ = onlinenow.StudentInId  //学生进入密码
			//pardStudent.CreateTime = time.Now().Format("2006-01-02 03:04:05 PM")
			pardStudent.UserID = strconv.Itoa(userinfo.Id)
			pardStudent.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath
			//判断此教室现在是否有老师
			meetroom := bbb4go.MeetingRoom{}
			meetroom.MeetingID_ = onlinenow.ClassroomId
			meetroom.ModeratorPW_ = onlinenow.TeacherInId
			meetroom.GetMeetingInfo()
			var pcount = meetroom.MeetingInfo.ParticipantCount
			if pcount == 0 { //老师不在
				urlse = "-5"
			} else if pcount >= 2 { //已有人在上课
				urlse = "1"
			} else if pcount == 1 {
				attent := meetroom.MeetingInfo.Attendees
				onlinetid := attent.Attendees[0].UserID
				if onlinetid == strconv.Itoa(onlinenow.UserIdPassive) {
					//进入前添加学生进入记录
					pardStudent.GetJoinURL()
					urlse = pardStudent.JoinURL
				} else {
					urlse = "-4" //不是自己的老师
				}
			}
		} else {
			urlse = "-2"
		}
	} else {
		urlse = "-1"
	}
	return
}
