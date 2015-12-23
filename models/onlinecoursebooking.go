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
	} else {
		fmt.Println("chaxuncuowu ")
		fmt.Println(err.Error())
	}
	return nil, err
}

// GetOnlinecoursebookingById retrieves Onlinecoursebooking by Id. Returns error if
// Id doesn't exist
func GetOnlinecoursebookingByIdModel(id int) (v Onlinecoursebooking, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	rs = o.Raw(SqlOnlineBooningbyid, id)
	qs := rs.QueryRow(&v)
	return v, qs
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

// DeleteOnlinecoursebookingMeeting deletes Onlinecoursebooking by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnlinecoursebookingMeeting(id int) (err error) {
	o := orm.NewOrm()
	//v := Onlinecoursebooking{Id: id}
	var onlineclass *Onlinecoursebooking
	onlineclass, err = GetOnlinecoursebookingById(id)
	//end白板信息，清楚预约信息中白板信息
	meetroom := bbb4go.MeetingRoom{} //释放白板
	meetroom.MeetingID_ = onlineclass.ClassroomId
	meetroom.ModeratorPW_ = onlineclass.TeacherInId
	meetroom.End()
	fmt.Println("课程结束后的meetroom:")
	fmt.Println(meetroom)
	fmt.Println(&onlineclass)

	var rs orm.RawSeter
	var roocount int
	//var upstr string = `update onlinecoursebooking set ClassroomId='' ,StudentInId='' ,TeacherInId='' where PKId=` + strconv.Itoa(id) + `; SELECT ROW_COUNT() as roocount;`
	var upstr string = `update onlinecoursebooking set ClassroomId='' ,StudentInId='' ,TeacherInId='' where PKId=` + strconv.Itoa(id) + `;`
	rs = o.Raw(upstr)
	rs.QueryRow(&roocount)
	fmt.Println(roocount)
	return
}

//获取随机数方法
func getcode(index string) (vcode string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000)) //自动生成随机数
	return vcode + index
}

//更新白板信息方法
//根据预约信息主键id更新白板meetingid 老师进入密码，学生进入密码
func Updatebookings(id int, classid string, sid string, tid string) (num int, err error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	fmt.Println("参数为：")
	fmt.Println(id)
	fmt.Println(classid + sid + tid)
	//var upstr string = `update onlinecoursebooking set ClassroomId='` + classid + `' ,StudentInId='` + sid + `' ,TeacherInId='` + tid + `' where PKId=` + strconv.Itoa(id) + `; SELECT ROW_COUNT() as roocount;`
	var upstr string = `update onlinecoursebooking set ClassroomId='` + classid + `' ,StudentInId='` + sid + `' ,TeacherInId='` + tid + `' where PKId=` + strconv.Itoa(id) + `;`
	fmt.Println(upstr)
	rs = o.Raw(upstr)
	qs := rs.QueryRow(&num)
	//fmt.Println(qs)
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
				fmt.Println("判断是否有人：")
				//判断白板教室内是否有人
				meetroom := bbb4go.MeetingRoom{}
				meetroom.MeetingID_ = onlinenow.ClassroomId
				meetroom.ModeratorPW_ = onlinenow.TeacherInId
				meetroom.GetMeetingInfo()
				var pcount = meetroom.MeetingInfo.ParticipantCount
				fmt.Println("判断是否有人：")
				fmt.Println(pcount)
				if pcount >= 2 { //进入重复课堂
					//重新建立课堂并保存到数据库
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

// 重写老师创建课堂方法，查询预约信息，是否存在meetingid，存在- 判断此课堂是否存在，存在-判断里面是否有人
//																     不存在-根据已有meetid建立课堂（此种可能几乎没有）
//												 不存在- 生成meetingid 建立课堂，
// onlineid:当前预约课程主键id
func Getecherlession3(onlineid int) (urlse string, err error) {
	//获取预约信息
	onlinenow, onlineerr := GetOnlinecoursebookingById(onlineid)
	timenow := time.Now()
	var overTimeMinute = TotalMinute
	if timenow.Hour() == onlinenow.EndTime.Hour() {
		overTimeMinute = 60 - AdvanceMinutes - timenow.Minute()
	} else {
		overTimeMinute = (60 - timenow.Minute()) + TotalMinute
	}

	if onlineerr == nil && onlinenow.Id > 0 {
		if onlinenow.ClassroomId != "" { //meetingid存在
			//判断此meetingid的会议室是否存在
			mettroom := bbb4go.MeetingRoom{}
			mettroom.MeetingID_ = onlinenow.ClassroomId
			istorf := mettroom.IsMeetingRunning()
			fmt.Println("课堂是否存在：")
			fmt.Println(istorf)
			if istorf { //如果会议室存在，判断会议室是否有人
				mettroom.ModeratorPW_ = onlinenow.TeacherInId
				mettroom.GetMeetingInfo()
				var pcount = mettroom.MeetingInfo.ParticipantCount
				if pcount <= 1 { //没有人在
					urlse = strconv.Itoa(onlinenow.Id) //老师可以跳页进入课堂
				} else {
					urlse = "-2" //会议室已存在一个人以上，老师不得进入
				}
			} else { //如果会议室不存在，建立会议室
				meetingroom := bbb4go.MeetingRoom{}
				meetingroom.Name_ = "泛鲲教育第一教室"
				meetingroom.MeetingID_ = onlinenow.ClassroomId
				meetingroom.AttendeePW_ = onlinenow.StudentInId
				meetingroom.ModeratorPW_ = onlinenow.TeacherInId
				meetingroom.Welcome = "欢迎来到泛鲲教育第一教室"
				meetingroom.LogoutURL = "http://" + OnlineUrl + "/orange/Teacher/ClassOverHtml/" //试听结束进入首页+/orange/Teacher/ClassOverHtml/
				meetingroom.Duration = overTimeMinute
				meetingroom.AllowStartStopRecording = false
				meetingroom.CreateMeeting()
				fmt.Println("课堂是否建立成功：")
				fmt.Println(meetingroom.CreateMeetingResponse.Returncode)
				if meetingroom.CreateMeetingResponse.Returncode == "SUCCESS" {
					urlse = strconv.Itoa(onlinenow.Id) //会议室创建成功老师可以跳页进入课堂
				}

			}
		} else { //meetingid不存在
			var meetingID string = getcode("1")
			var attendeePW string = getcode("2")  //学生进入密码
			var moderatorPW string = getcode("3") //老师进入密码
			meetingroom := bbb4go.MeetingRoom{}
			meetingroom.Name_ = "泛鲲教育第一教室"
			meetingroom.MeetingID_ = meetingID
			meetingroom.AttendeePW_ = attendeePW
			meetingroom.ModeratorPW_ = moderatorPW
			meetingroom.Welcome = "欢迎来到泛鲲教育第一教室"
			meetingroom.LogoutURL = "http://" + OnlineUrl + "/orange/Teacher/ClassOverHtml/" //试听结束进入首页+/orange/Teacher/ClassOverHtml/
			meetingroom.Duration = overTimeMinute
			meetingroom.AllowStartStopRecording = false
			//将白板信息保存到数据库试听信息中
			onlinenow.StartTime = time.Now()
			onlinenow.ClassroomId = meetingID
			onlinenow.StudentInId = attendeePW
			onlinenow.TeacherInId = moderatorPW
			_, uperr := Updatebookings(onlineid, meetingID, attendeePW, moderatorPW) //更新数据库的白板信息
			fmt.Println(uperr)
			//if uperr == nil {
			meetingroom.CreateMeeting()
			fmt.Println("课堂是否建立成功：")
			fmt.Println(meetingroom.CreateMeetingResponse.Returncode)
			if meetingroom.CreateMeetingResponse.Returncode == "SUCCESS" {
				urlse = strconv.Itoa(onlinenow.Id) //会议室创建成功老师可以跳页进入课堂
			}
			//}
		}
	}
	return
}

//获取老师进入课堂url
func GetOnlineClassTeacherurl(bookid int) (urlse string, err error) {
	//onlinetrylisten, geterr := GetOnlinetrylistenById(listenid)
	onlineclass, geterr := GetOnlinecoursebookingByIdModel(bookid)
	fmt.Println("获取老师进入结构：")
	fmt.Println(onlineclass)
	if geterr == nil && onlineclass.Id > 0 {
		//获取老师信息
		userinfo, gerr := GetUserinformationTeacher(onlineclass.UserIdPassive)
		fmt.Println(userinfo)
		if gerr == nil {
			urlse = "0"
			pardTeacher := bbb4go.Participants{}
			pardTeacher.IsAdmin_ = 0
			pardTeacher.FullName_ = userinfo.UserName
			pardTeacher.MeetingID_ = onlineclass.ClassroomId //教室id
			pardTeacher.Password_ = onlineclass.TeacherInId  //老师进入密码
			//pardTeacher.CreateTime = time.Now().Format("2006-01-02 03:04:05 PM")
			pardTeacher.UserID = strconv.Itoa(userinfo.Id)
			pardTeacher.AvatarURL = "http://" + OnlineUrl + "/" + userinfo.AvatarPath

			fmt.Println("老师进入结构：")
			fmt.Println(pardTeacher)
			pardTeacher.GetJoinURL()
			urlse = pardTeacher.JoinURL

			fmt.Println(urlse)
		}
	}
	return
}

//重新学生进入课堂的方法---------学生查询在线白板的密码并进入课堂
func Getstudentlession2(onlineid int) (urlse string, err error) {
	//获取预约课程信息
	onlinenow, onerr := GetOnlinecoursebookingById(onlineid)
	//获取老师信息
	userinfo, usererr := GetUserinformationStudent(onlinenow.UserIdActive)
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

////结算课程
////传入参数：onlineid:预约信息主键id
////传出参数：resultmsg:返回结果 err:错误信息
////        resultmsg---  >0:结算成功，返回总分钟数
////        				-1:查询不到预约信息
////        				-2:没有课程时间记录（几乎没有这种可能）
////        				-3:没有查到冻结资金
////        				-4:没有查到用户账户信息
////        				-5:给用户返还钱失败
////        				-6:解冻资金失败
////        				-7:新增交易记录失败
////        				-8:更新预约信息状态失败
////        				-9:添加在线课程记录失败
////2015-12-19
////思路：1.查询预约信息 2.查询此次预约课程所有时间记录信息，计算总时间 3.根据学生主键找到此次预约冻结资金，解冻，根据冻结时间分配金额，4.钱打给老师，剩余退还学生，
////     5.新增一条老师打钱给学生的交易记录，6.将预约信息状态修改为已学习，已支付，7.新增一条课程记录信息
//func SetUserClassPay(onlineid int) (resultmsg string, err error) {
//	//根据预约信息主键id查询 一条预约信息
//	onlineclass, geterr := GetOnlinecoursebookingById(onlineid)
//	fmt.Println("获取预约信息")
//	fmt.Println(onlineid)
//	fmt.Println(onlineclass)
//	fmt.Println("1")
//	if geterr == nil && onlineclass.Id > 0 {
//		allminutes := GetALLtimeminute(onlineid)
//		if allminutes > 0 {
//			//查询此次预约的冻结信息
//			fonze, ferr := GetFrozenfundsByUidOnId(onlineclass.UserIdActive, 0, onlineid)
//			fmt.Println("2")
//			if ferr == nil && fonze.Id > 0 {
//				//计算此次课程总费用
//				allm, _ := strconv.ParseFloat(strconv.Itoa(allminutes), 64)
//				alltm, _ := strconv.ParseFloat(strconv.Itoa(TotalMinute), 64)
//				teachermoney := (allm / alltm) * fonze.FrozenMoney //（上课分钟/总分钟）*总钱数 = 应给老师多少钱
//				returnmoney := fonze.FrozenMoney - teachermoney    //返还学生的钱
//				resultmsg = SetUserMoney(onlineclass.UserIdPassive, teachermoney)
//				resultmsg = SetUserMoney(onlineclass.UserIdActive, returnmoney)
//				fmt.Println("3")
//				if resultmsg == "1" { //钱各自打成功，解冻冻结资金信息
//					fonze.FrozenState = 0
//					upfonzerr := UpdateFrozenfundsById(&fonze)
//					fmt.Println("4")
//					if upfonzerr == nil { //解冻成功新增一条交易记录
//						resultmsg = AddUserTransactionRecords(onlineclass.UserIdActive, onlineclass.UserIdPassive, teachermoney)
//						fmt.Println("5")
//						if resultmsg == "1" { //添加交易记录完成
//							onlineclass.Payment = 1
//							onlineclass.Leaming = 1
//							uponerr := UpdateOnlinecoursebookingById(onlineclass)
//							fmt.Println("6")
//							if uponerr == nil {
//								//新增一条课程记录信息
//								var onlinerecord Onlinecourserecord
//								onlinerecord.OCBId = onlineclass.Id
//								onlinerecord.UserIdActive = onlineclass.UserIdActive
//								onlinerecord.UserIdPassive = onlineclass.UserIdPassive
//								onlinerecord.CourseContent = onlineclass.AppointMessage
//								onlinerecord.StartTime = onlineclass.StartTime
//								onlinerecord.EndTime = onlineclass.EndTime
//								onlinerecord.UnitPrice = (allm / alltm)
//								onlinerecord.TotalPrice = teachermoney
//								onlinerecord.ClassNumber = allminutes
//								addrecordid, recorderr := AddOnlinecourserecord(&onlinerecord)
//								fmt.Println("7")
//								if recorderr == nil && addrecordid > 0 {
//									resultmsg = strconv.Itoa(allminutes) //返回总分钟数
//									fmt.Println("8")
//								} else {
//									resultmsg = "-9"
//								}
//							} else {
//								resultmsg = "-8"
//							}
//						} else {
//							resultmsg = "-7"
//						}
//					} else {
//						resultmsg = "-6"
//					}
//				}
//			} else {
//				resultmsg = "-3"
//			}
//		} else {
//			resultmsg = "-2" //没有课程时间记录
//		}
//	} else {
//		resultmsg = "-1" //查询不到预约信息
//	}
//	return
//}

//根据预约信息主键计算两人共同在线时间
func GetALLtimeminute(onlineid int) (allminute int) {
	classnow, _ := GetOnlinecoursebookingById(onlineid)
	teacherlistrecord, _ := GetOnlinecoursebookingrecordBybookiduid(classnow.UserIdPassive, onlineid)
	studentlistrecord, _ := GetOnlinecoursebookingrecordBybookiduid(classnow.UserIdActive, onlineid)
	fmt.Println(teacherlistrecord)
	fmt.Println(studentlistrecord)
	var tint [50]int //老师在线时间数组
	var sint [50]int //学生在线时间数组

	loc, _ := time.LoadLocation("Local")
	//	var starttime string = ""
	//	starttime = strconv.Itoa(classnow.StartTime.Year()) + "-" + GetMonth(classnow.StartTime.Month()) + "-" + strconv.Itoa(classnow.StartTime.Day()) + " 24:00:00"
	//	dd, _ := time.ParseInLocation("2006-01-02 15:04:05", starttime, loc) //将字符串转换为时间
	//外层循环50数组，循环每个时间点
	//内层循环老师所有时间记录信息，比对外层时间点是否存在于这些数组之中，存在即修改数组值为1，不存在值为0
	for i := 0; i < len(tint); i++ {
		tint[i] = 0
		minute := strconv.Itoa(i + 1)
		if i < 9 {
			minute = "0" + strconv.Itoa(i+1)
		}
		var itemtime string = strconv.Itoa(classnow.StartTime.Year()) + "-" + GetMonth(classnow.StartTime.Month()) + "-" + strconv.Itoa(classnow.StartTime.Day()) + " "
		itemtime = itemtime + strconv.Itoa(classnow.StartTime.Hour()) + ":" + minute + ":00"
		timecomp, _ := time.ParseInLocation("2006-01-02 15:04:05", itemtime, loc) //将字符串转换为时间
		fmt.Println(timecomp)
		for j := 0; j < len(teacherlistrecord); j++ {
			var endtime time.Time
			if teacherlistrecord[j].EndTime.Year() <= 1 {
				endtime = time.Now()
			} else {
				endtime = teacherlistrecord[j].EndTime
			}
			if teacherlistrecord[j].StartTime.Before(timecomp) && timecomp.Before(endtime) {
				tint[i] = 1
				fmt.Println(teacherlistrecord[j].StartTime)
			}
		}

	}
	//外层循环50数组，循环每个时间点
	//内层循环学生所有时间记录信息，比对外层时间点是否存在于这些数组之中，存在即修改数组值为1，不存在值为0
	for i := 0; i < len(sint); i++ {
		sint[i] = 0
		minute := strconv.Itoa(i + 1)
		if i < 9 {
			minute = "0" + strconv.Itoa(i+1)
		}
		var itemtime string = strconv.Itoa(classnow.StartTime.Year()) + "-" + GetMonth(classnow.StartTime.Month()) + "-" + strconv.Itoa(classnow.StartTime.Day()) + " "
		itemtime = itemtime + strconv.Itoa(classnow.StartTime.Hour()) + ":" + minute + ":00"
		timecomp, _ := time.ParseInLocation("2006-01-02 15:04:05", itemtime, loc) //将字符串转换为时间
		for j := 0; j < len(studentlistrecord); j++ {
			var endtime time.Time
			if studentlistrecord[j].EndTime.Year() <= 1 {
				endtime = time.Now()
			} else {
				endtime = studentlistrecord[j].EndTime
			}
			if studentlistrecord[j].StartTime.Before(timecomp) && timecomp.Before(endtime) {
				sint[i] = 1
				fmt.Println(studentlistrecord[j].StartTime)
			}
		}

	}
	//循环两个数组计算共同在线时间
	allminute = 0
	for i := 0; i < len(tint); i++ {
		if tint[i] == 1 && sint[i] == 1 {
			allminute = allminute + 1
		}
	}
	fmt.Println(tint)
	fmt.Println(sint)
	return
}

func GetMonth(month time.Month) (yue string) {
	if month.String() == "January" {
		yue = "1"
	} else if month.String() == "February" {
		yue = "2"
	} else if month.String() == "March" {
		yue = "3"
	} else if month.String() == "April" {
		yue = "4"
	} else if month.String() == "May" {
		yue = "5"
	} else if month.String() == "June" {
		yue = "6"
	} else if month.String() == "July" {
		yue = "7"
	} else if month.String() == "August" {
		yue = "8"
	} else if month.String() == "September" {
		yue = "9"
	} else if month.String() == "October" {
		yue = "10"
	} else if month.String() == "November" {
		yue = "11"
	} else if month.String() == "December" {
		yue = "12"
	}
	return
}

//给此用户打钱
func SetUserMoney(userid int, money float64) (result string) {
	useraccount, accerr := GetAccountfundsByuid(userid)
	if accerr == nil && useraccount.Id > 0 {
		useraccount.Balance = useraccount.Balance + money
		uperr := UpdateAccountfundsById(&useraccount)
		if uperr == nil {
			result = "1"
		} else {
			result = "-5"
		}
	} else {
		result = "-4"
	}
	return
}

//新增一条交易记录
func AddUserTransactionRecords(sid int, tid int, money float64) (result string) {
	var addtrecord Transactionrecords
	addtrecord.SendUserId = sid
	addtrecord.CollectUserId = tid
	addtrecord.RecordMoney = money
	addtrecord.TradingWayId = TradingWayId
	addtrecord.RecordTime = time.Now()
	addid, adderr := AddTransactionrecords(&addtrecord)
	if adderr == nil && addid > 0 {
		result = "1"
	} else {
		result = "-7"
	}
	return
}
