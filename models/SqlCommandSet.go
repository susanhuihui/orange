package models

var limitSql string = ` limit ?,?`

//1.
//用途：首页图片轮换，查询老师头像，姓名，主辅导课程
//查询字段：姓名，头像路径，主辅导课程，总课时数
//查询条件：全部老师，按照总课时降序，身份为老师，头像图片不能为空
//Controller：userinformation
//调用方法名：GetUserinformationPicMove
//参数说明：count 显示条数
//2015-11-03
var SqlUserPicList string = `SELECT users.PKId,UserName,AvatarPath,
	(select GradeName from grade as gd where gd.PKId = users.GradeId) as GradeName,
	(select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=users.PKId and IsMain=1 limit 1)) as CourseName,
	ifnull((select sum(ClassNumber) from onlineeducation.onlinecourserecord as oc where oc.OCBId in 
	(select PkId from onlineeducation.onlinecoursebooking as ob where ob.UserIdPassive = users.PKId)),0) as counts 
	 FROM onlineeducation.userinformation as users 
	where identityid = (select PkId from identity where identityName = '老师') and users.AvatarPath  <> '' and users.AvatarPath is not null 
	 order by counts desc 
	 limit ? `

//2.
//用途：查询老师个人信息详情
//查询字段：用户信息表全部字段，身份，所在大学，总辅导数，总课时数，上课次数，主辅导课程
//查询条件：根据用户主键id查询
//Controller：
//调用方法名：
//参数说明：userid 用户主键id
//2015-11-03
var SqlUserTeacher string = `select * ,
	(select IdentityName from identity as iden where users.identityid = iden.PkId) as IdentityName,
    ifnull((select count(userfor.pkid) from userinformation userfor where userfor.pkid in (select useridactive from onlineeducation.onlinecourserecord as ob where ob.UserIdPassive = users.PKId group by useridactive)),0) as AllPerson,
    ifnull((select sum(ClassNumber) from onlineeducation.onlinecourserecord as oc where oc.UserIdPassive = users.pkid) ,0) as  AllDate,
    ifnull((select count(*) from onlineeducation.onlinecourserecord as oc where oc.UserIdPassive = users.pkid) ,0)as AllCount,
        (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=users.PKId and IsMain=1 limit 1)) as CourseName,
        (select cous.pkid from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=users.PKId and IsMain=1 limit 1)) as CourseNameId,
    (select DegreeName from degree as deg where users.UserDegree = deg.PkId) as DegreeName,
    (select gradename from grade where grade.pkid = users.gradeid) as GradeName
from onlineeducation.userinformation as users 
	 where users.pkid = ? `

//3.
//用途：查询用户的主辅导课程或辅辅导课程信息
//查询字段：用户id，课程id，主or辅，课程名称
//查询条件：根据用户主键或是主还是辅
//Controller：
//调用方法名：
//参数说明：userid 用户主键，ismain（0/1）是主是辅
//2015-11-03
var SqlUserMainCourse string = `select *,(select CourseName from course as cous where cous.PKId = rec.Coursesid) as CourseName 
	 from remedialcourses rec where userid = ? and ismain = ? `
var SqlUserMainCourse2 string = `select *,(select CourseName from course as cous where cous.PKId = rec.Coursesid) as CourseName 
	 from remedialcourses rec where userid = ? and ismain = ? 
     limit 1`

//4.
//用途：查询老师全部课程信息
//查询字段：在线课程表字段，课程类型，学生姓名，学生学龄段
//查询条件：用户主键，分页
//Controller：
//调用方法名：
//参数说明：第一个参数为用户主键id，第二个参数为从第几行开始不包括此行，第三个参数为获取几行
//2015-11-03
var SqlOnlineAllTeacher string = `select *, 
	 (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=onlinerecd.UserIdPassive and IsMain=1 limit 1)) as CourseName, 
	 (select username from userinformation as userinfo where userinfo.pkid = onlinerecd.UserIdActive) as UserName,
     (select AgeName from schoolages as scha where scha.pkid = (select schoolageid from userinformation as userin where userin.pkid = onlinerecd.useridactive)) as AgeName 
	 from onlinecourserecord as onlinerecd 
	 where onlinerecd.UserIdPassive = ? 
     order by endtime desc `

//5.
//用途：老师查询在线课程评价内容
//查询字段：在线课程全部字段，评价人主键，姓名，头像
//查询条件：根据在线课程主键
//Controller：
//调用方法名：
//参数说明：在线课程主键id
//2015-11-03
var SqlOnlineEvaluationByT string = `select online.*,userinfo.pkid as UserPkid,userinfo.UserName,userinfo.AvatarPath 
	 from onlinecourseevaluation as online join userinformation as userinfo 
	 on  online.userid = userinfo.pkid 
	 where online.OCRId = ? `

//6.
//用途：根据老师主键查询预约课程信息
//查询字段：预约课程表字段，主修课程，学习姓名，学生电话
//查询条件：老师用户主键，根据预约开始时间升序，没有学习过的，分页
//Controller：
//调用方法名：
//参数说明：老师用户主键，从第几行开始，获取多少行
//2015-11-03
var SqlOnlineTeacherbookingByT string = `select online.*,
			(select CourseName from course as cous where cous.PKId = 
				(select CoursesId from remedialcourses as remec where remec.UserId=online.UserIdPassive and IsMain=1 limit 1)) as CourseName,
			userinfo.UserName,userinfo.IphoneNum
			from onlinecoursebooking as online join userinformation as userinfo on online.UserIdActive = userinfo.pkid
			where online.UserIdPassive = ?   and starttime >= CURDATE()`

//7.
//用途：老师查询我的留言列表,分页
//查询字段：留言表所有字段，留言人姓名，最新留言时间
//查询条件：老师主键，分页，根据最新留言时间降序排列
//Controller：
//调用方法名：
//参数说明：老师主键id，从第几行开始获取，获取几行
//2015-11-03
var SqlUserMessageTeacher string = `select usermsg.*,userinfo.UserName,(select mestime from usermessage as msg where msg.messageid = usermsg.pkid  order by mestime desc limit 1) as MesTimeNew ,
     (select count(*) from usermessage uu where (uu.ActiveUserId = usermsg.ActiveUserId and uu.States =0  and uu.MessageId = usermsg.PKId)or(uu.pkid = usermsg.pkid and uu.states=0)) as State
	 from usermessage as usermsg join userinformation as userinfo on usermsg.ActiveUserId = userinfo.pkid 
	 where usermsg.Messageid = 0 and usermsg.PassiveUserId = ?  
	 order by MesTime desc `

//8.
//用途：老师查看一条留言详情
//查询字段：
//查询条件：留言首条主键id
//Controller：
//调用方法名：
//参数说明：留言首条主键id,留言首条主键id
//2015-11-03
var SqlUserMessageTchOne string = `select usermsg.*, (select username from userinformation as users where users.pkid = usermsg.ActiveUserId) as ActiveName,
	(select username from userinformation as users where users.pkid = usermsg.PassiveUserId) as PassiveName 
	from usermessage as usermsg  
	where usermsg.pkid = ? or messageid = ?`

//9.
//用途：老师查看最近“浏览”过我的人（包括老师和学生）
//查询字段：关系表字段，姓名，头像，学龄段，学校，身份
//查询条件：老师主键，分页，浏览关系，根据浏览时间降序
//Controller：
//调用方法名：
//参数说明：老师主键id，关系，从第几行开始，获取几行
//2015-11-04
var SqlRelationByTeacher string = `select rel.*,userinfo.UserName,userinfo.AvatarPath,
	(select AgeName from schoolages as sa where userinfo.schoolageid = sa.pkid) as AgeName,userinfo.SchoolName,
	(select identityname from identity as ide where userinfo.IdentityId = ide.pkid) as IdentityName 
	from relations as rel join userinformation as userinfo on rel.afteruserid = userinfo.pkid 
	where FrontUserId = ? and sources = ?   
	order by setdate desc  `

//10.
//用途：查询老师预约信息
//查询字段：在线预约字段
//查询条件：老师主键，时间
//Controller：
//调用方法名：
//参数说明：老师主键id，哪个时间之后的所有预约信息
//2015-11-04
var SqlOnlineBookingByT string = `select *
	from onlinecoursebooking 
	where starttime > '?' and useridpassive = ?`

//11.
//用途：根据公告标题查询公告内容
//查询字段：公告表字段
//查询条件：公告标题
//Controller：
//调用方法名：
//参数说明：公告标题
//2015-11-04
var SqlTreatysByTitle string = `select * 
		from treatys 
		where treatytitle = '?'`

//12.
//用途：查询某个区县下的某类型学校信息
//查询字段：学校表字段
//查询条件：地区，类型
//Controller：
//调用方法名：
//参数说明：区县id，学校类型（0为中小学，1为大学）
//2015-11-04
var SqlSchoolByCityType string = `select * 
	from schools 
	where cityid = ? and schooltype = ?`

//13.
//用途：查询用户余额信息
//查询字段：账户表
//查询条件：根据用户查询
//Controller：
//调用方法名：
//参数说明：用户主键id
//2015-11-04
var SqlAccountFundsByUserId string = `select * 
	from accountfunds 
	where userid = ?`

//14.
//用途：根据被支付方查询交易记录（老师）
//查询字段：交易记录字段，支付方，支付方式
//查询条件：被支付方主键，分页，排序
//Controller：
//调用方法名：
//参数说明：被支付方主键id，从第几行开始，查询几行
//2015-11-04
var SqlTransactionRecordByT string = `select tr.*, (select UserName from userinformation as userinfo where tr.senduserid = userinfo.pkid) as UserName,
	(select tradingname from tradingway as tw where tr.TradingWayId = tw.pkid) as TradingName 
	from transactionrecords as tr 
	where CollectUserId = ? 
	order by tr.RecordTime desc  `

//15.
//用途：查询用户的提现/充值信息记录
//查询字段：记录字段，操作方式
//查询条件：用户主键，充值还是提现，按操作时间降序，分页
//Controller：
//调用方法名：
//参数说明：recordtype：0为充值，1为提现；userid用户主键id；从第几行开始，获取几行
//2015-11-04
var SqlAccountRecordByUidType string = `select amt.*,(select tradingname from tradingway as tw where amt.TradingWayId = tw.pkid) as TradingName 
	from amountrecords as amt 
	where recordtype = ? and userid = ?  and IsComplete=1
	order by recordtime desc  `

//查询用户（提现recordtype = 1）全部提现记录
//
//参数说明：userid用户主键id；从第几行开始，获取几行
var SqlAccountRecordTixianByUid string = `select amt.*,(select tradingname from tradingway as tw where amt.TradingWayId = tw.pkid) as TradingName
	from amountrecords as amt 
	where recordtype = 1 and userid = ? 
	order by recordtime desc `

//查询用户正在提现的全部金额，继续提现的时候根据此值判断是否可继续提现
//参数：用户id，返回正在申请提现的金额总值
var SqlAccountRecordTMcountByUID string = `select sum(amt.RecordMoney) as RecordMoney
	from amountrecords as amt
	where recordtype = 1 and IsComplete=0 and userid = ?
	order by recordtime desc
	limit 0,10`

//管理员查询全部用户正在申请的（提现recordtype = 1）全部提现记录
//无参数，分页
var SqlAccountRecordAll string = `select amt.*,(select username from userinformation as userinfo where userinfo.pkid = amt.userid) as UserName
	,(select IphoneNum from userinformation as userinfo where userinfo.pkid = amt.userid) as IphoneNum
,(select identityname from identity as iden where iden.pkid = (select identityid from userinformation as userinfo where userinfo.pkid = amt.userid)) as IdentityName,
 (select identityid from userinformation as userinfo where userinfo.pkid = amt.userid) as IdentityId
	from amountrecords as amt
	where recordtype = 1 and iscomplete = 0
	order by recordtime desc `

//16.
//用途：查询被提问者的问题信息
//查询字段：问题表字段，提问人姓名
//查询条件：被提问者，分页，根据提问时间降序
//Controller：
//调用方法名：
//参数说明：被提问者主键id，从第几行开始，获取几行
//2015-11-04
var SqlQuestionAskByTUserid string = `select qa.PKId,qa.AskUserId,qa.AnswerUserId,qa.GCId,qa.Title,qa.BadeTime,qa.EndTime,qa.AmountMoney,qa.IsSee,
	(select username from userinformation as userinfo where qa.askuserid = userinfo.pkid) as UserName ,
	(select count(*) from answers as ans where ans.qaid = qa.pkid) as AnswerCount
	from questionask as qa 
	where qa.answeruserid = ? 
	order by qa.badetime desc `

//17.
//用途：查询学生详细信息
//查询字段：用户信息字段，身份，学校，辅导老师总数，总课时，提问总数，学龄段
//查询条件：根据用户主键
//Controller：
//调用方法名：
//参数说明：用户主键id
//2015-11-03
var SqlUserInformationByS string = `select * ,
		(select IdentityName from identity as iden where users.identityid = iden.PkId) as IdentityName,
	    (select count(userfor.pkid) from userinformation userfor where userfor.pkid in 
			(select UserIdPassive from onlineeducation.onlinecourserecord as ob where ob.UserIdActive = users.pkid group by UserIdPassive)) as AllPerson,
	    ifnull((select sum(ClassNumber) from onlineeducation.onlinecourserecord as oc where oc.OCBId in 
			(select PkId from onlineeducation.onlinecoursebooking as ob where ob.UserIdActive = users.PKId)),0) as AllDate,
	    ifnull((select count(*) from onlineeducation.questionask as qa where qa.askuserid = users.pkid),0) as AllCount,
	        (select agename from schoolages as sas where users.schoolageid = sas.pkid) as AgeName
	from onlineeducation.userinformation as users
	where users.pkid = ?`

//18.
//用途：查询学生全部课程信息
//查询字段：在线课程表字段，课程类型，老师姓名
//查询条件：用户主键，分页
//Controller：
//调用方法名：
//参数说明：第一个参数为用户主键id，第二个参数为从第几行开始不包括此行，第三个参数为获取几行
//2015-11-04
var SqlOnlineAllStudent string = `select * , 
	 (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=onlinerecd.UserIdPassive and IsMain=1 limit 1)) as CourseName, 
	 (select username from userinformation as userinfo where userinfo.pkid = onlinerecd.UserIdPassive) as UserName ,
     ifnull((select recone.pkid from onlinecourseevaluation as recone where recone.ocrid = onlinerecd.pkid limit 1),0) as RecordId
	 from onlinecourserecord as onlinerecd 
	 where onlinerecd.UserIdActive = ? 
     order by endtime desc  `

//19.
//用途：学生查询在线课程评价内容
//查询字段：在线课程全部字段，被评价人主键，姓名，头像，主辅导课程
//查询条件：根据在线课程主键
//Controller：
//调用方法名：
//参数说明：在线课程主键id
//2015-11-03
var SqlOnlineEvaluationByS string = `select online.*,userinfo.pkid as UserPkid,userinfo.username,userinfo.AvatarPath,
	(select coursename from course as cou where cou.pkid = (select coursesid from remedialcourses as recs where ismain=1 and recs.userid = 
		(select useridPassive from onlinecourserecord as onrec where onrec.pkid = online.ocrid limit 1) limit 1)) as CourseName 
	 from onlinecourseevaluation as online join userinformation as userinfo  
	 on  userinfo.pkid =  online.UserId 
	 where online.OCRId = ?`

//20.
//用途：根据学生主键查询预约课程信息
//查询字段：预约课程表字段
//查询条件：学生用户主键，根据预约开始时间升序，没有学习过的，分页
//Controller：
//调用方法名：
//参数说明：学生用户主键，从第几行开始，获取多少行
//2015-11-03
var SqlOnlineTeacherbookingByS string = `select online.*, 
	 (select CourseName from course as cous where cous.PKId =  
	 (select CoursesId from remedialcourses as remec where remec.UserId=online.UserIdPassive and IsMain=1 limit 1)) as CourseName, 
	 userinfo.UserName,userinfo.IphoneNum 
	 from onlinecoursebooking as online join userinformation as userinfo on online.UserIdPassive = userinfo.pkid 
	 where online.UserIdActive = ? and starttime >= CURDATE()   
     order by starttime asc  `

//查询学生没有上过的预约课程
var SqlOnlinebookingBySidNotOn string = ` select online.*, 
	 (select CourseName from course as cous where cous.PKId =  
	 (select CoursesId from remedialcourses as remec where remec.UserId=online.UserIdPassive and IsMain=1 limit 1)) as CourseName, 
     (select UserName from userinformation as uf where uf.pkid = online.useridpassive) as UserName,
     (select IphoneNum from userinformation as uf where uf.pkid = online.useridpassive) as IphoneNum 
	 from onlinecoursebooking as online   
	 where online.UserIdActive = ? and starttime < CURDATE() and leaming=0 and online.pkid not in(select ocbid from onlinecourserecord) 
     order by starttime asc`

//查询学生预约某个老师某天预约了几次课程
var SqlOnlinebookingbySTidTime string = ` select online.*, 
	 (select CourseName from course as cous where cous.PKId =  
	 (select CoursesId from remedialcourses as remec where remec.UserId=online.UserIdPassive and IsMain=1 limit 1)) as CourseName, 
     (select UserName from userinformation as uf where uf.pkid = online.useridpassive) as UserName,
     (select IphoneNum from userinformation as uf where uf.pkid = online.useridpassive) as IphoneNum 
	 from onlinecoursebooking as online  
	 where online.UserIdActive = ? and online.useridpassive = ? and leaming=0 and online.pkid not in(select ocbid from onlinecourserecord) 
	 and starttime > ? and endtime < ?
     order by starttime asc`

//21.
//用途：查询学生全部关注的老师
//查询字段：关系表字段，姓名，老师主辅导课程，学习名称，总课时
//查询条件：学生主键，关系关键词，分页，排序
//Controller：
//调用方法名：
//参数说明：学生用户主键，关系关键词（关注），从第几行开始，获取多少行
//2015-11-04
var SqlRelationsBySGuanZhuT string = `select rela.*,(select users.username from userinformation as users where users.pkid = rela.frontuserid) as UserName,
		(select CourseName from course as cous where cous.PKId =  
	    (select CoursesId from remedialcourses as remec where remec.UserId=rela.FrontUserId and IsMain=1 limit 1)) as CourseName,
     	(select SchoolName from userinformation as userin where userin.pkid = rela.FrontUserId) as SchoolName,
	     ifnull((select sum(ClassNumber) from onlineeducation.onlinecourserecord as oc where oc.OCBId in 
			(select PkId from onlineeducation.onlinecoursebooking as ob where ob.UserIdPassive = rela.FrontUserId)),0) as AllDate
	from relations  as rela 
	where rela.afteruserid = ? and sources = ?  
	order by AllDate desc  `

//22.
//用途：查询用户账户信息和冻结资金
//查询字段：账户表字段，冻结资金
//查询条件：用户主键,此用户已冻结的所有资金
//Controller：
//调用方法名：
//参数说明：用户主键id
//2015-11-04
var SqlAccountFundsByS string = `select af.*,
		ifnull((select sum(frozenmoney) from frozenfunds as ffs where ffs.afid = af.userid and FrozenState=0),0) as FrozenMoney 
	from accountfunds as af  
	where userid = ？ and FundState = 0 
	limit 1`

//23.
//用途：查询学生消费记录
//查询字段：交易记录字段，被支付方，交易方式
//查询条件：学生主键，分页，排序
//Controller：
//调用方法名：
//参数说明：学生主键id，从第几行开始，获取多少行
//2015-11-04
var SqlTranscationRecordsByUserid string = `select trds.*,(select username from userinformation as userinfo where userinfo.pkid = trds.collectuserid) as UserName,
	(select TradingName from tradingway as tw where tw.pkid = trds.tradingwayid) as TradingName
	from transactionrecords as trds 
	where trds.senduserid = ? 
	order by recordtime desc `

//24.
//用途：学生查询自己的提问
//查询字段：提问表字段，被提问人姓名
//查询条件：学生主键，分页，排序
//Controller：
//调用方法名：
//参数说明：学生主键id，从第行开始，获取多少行
//2015-11-05
var SqlQuestionAskBySUserid string = `select qa.PKId,qa.AskUserId,qa.AnswerUserId,qa.GCId,qa.Title,qa.BadeTime,qa.EndTime,qa.AmountMoney,qa.IsSee,
	(select username from userinformation as userinfo where qa.AnswerUserId = userinfo.pkid) as UserName 
	from questionask as qa 
	where AskUserId = ? 
	order by badetime desc `

//25.
//用途：学生查看全部留言信息
//查询字段：留言表全部字段，老师姓名，最新留言时间
//查询条件：学生主键，分页，根据最新留言时间降序排序
//Controller：
//调用方法名：
//参数说明：学生主键id，从第几行开始，获取多少行
//2015-11-05
var SqlUserMessageBySid string = `select usermsg.*,userinfo.UserName,(select mestime from usermessage as msg where msg.messageid = usermsg.pkid  order by mestime desc limit 1) as MesTimeNew, 	 
     (select count(*) from usermessage uu where uu.ActiveUserId = usermsg.PassiveUserId and uu.States =0  and uu.MessageId = usermsg.PKId) as State
	from usermessage as usermsg join userinformation as userinfo on usermsg.PassiveUserId = userinfo.pkid 
	 where usermsg.Messageid = 0 and usermsg.ActiveUserId = ? 
	 order by MesTime desc `

//26.
//用途：
//查询字段：
//查询条件：
//Controller：
//调用方法名：
//参数说明：
//2015-11-05
var SqlUser string = ``

//27.
//用途：检索老师全部信息
//查询字段：用户信息表字段，学校名称，专业，主辅导，辅辅导，排序字段（评价人数，课时总数，评价人数加课时总数和）
//查询条件：身份，年级，课程，级别，省市，（人气，总课时，和）排序，分页
//Controller：
//调用方法名：
//参数说明：年级名称，课程名称，级别名称，省份名称，市区名称，从第几行开始，获取多少行
//2015-11-05

//根据人气查询
var SqlUserinformationAllTeacherByPerson string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
		 ifnull((select count(*) from onlinecourseevaluation as onev where onev.ocrid in (select onrec.pkid from onlinecourserecord as onrec where onrec.useridpassive = userinfo.pkid)),0) as SortCondition,
         ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState 
	from userinformation as userinfo 
	where userinfo.identityid = (select pkid from identity where identityname='老师')
		and userinfo.schoolageidt like ? 
		and (select coursesid from remedialcourses recs where recs.userid = userinfo.pkid and ismain = 1 limit 1) in (select cou.pkid from course as cou where cou.coursename like ?)
		and userinfo.UserLevelId in (select ulv.pkid from userlevel as ulv where ulv.levelname like ?)
	    and userinfo.SeniorLocation in 
	    (select cit.pkid from citys as cit where cit.proid in (select prov.pkid from province as prov where prov.proname like ?))		
	order by SortCondition desc  `

var SqlUserinformationAllTeacherByPerson1 string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
		 ifnull((select count(*) from onlinecourseevaluation as onev where onev.ocrid in (select onrec.pkid from onlinecourserecord as onrec where onrec.useridpassive = userinfo.pkid)),0) as SortCondition,
         ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState 
	from userinformation as userinfo
	where userinfo.identityid = (select pkid from identity where identityname='老师') `

var SqlUserOver string = `	order by SortCondition desc  `
var SqlUserOveronline string = `	order by OnlineState desc  `

var SqlUserinformationAllTeacherByTime string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
	     ifnull((select sum(classnumber) from onlinecourserecord as ocr where ocr.useridpassive = userinfo.pkid),0) as SortCondition,
         ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState
	from userinformation as userinfo 
	where userinfo.identityid = (select pkid from identity where identityname='老师')
		and userinfo.schoolageidt like ? 
		and (select coursesid from remedialcourses recs where recs.userid = userinfo.pkid and ismain = 1 limit 1) in (select cou.pkid from course as cou where cou.coursename like ?)
		and userinfo.UserLevelId in (select ulv.pkid from userlevel as ulv where ulv.levelname like ?)
	    and userinfo.SeniorLocation in 
	    (select cit.pkid from citys as cit where cit.proid in (select prov.pkid from province as prov where prov.proname like ?))		
	order by SortCondition desc  `

var SqlUserinformationAllTeacherByTime1 string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
	     ifnull((select sum(classnumber) from onlinecourserecord as ocr where ocr.useridpassive = userinfo.pkid),0) as SortCondition,
         ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState
	from userinformation as userinfo 
	where userinfo.identityid = (select pkid from identity where identityname='老师')`

var SqlUserinformationAllTeacherByOnline string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
	     ifnull(
				((select count(*) from onlinecourseevaluation as onev where onev.ocrid in (select onrec.pkid from onlinecourserecord as onrec where onrec.useridpassive = userinfo.pkid))+
				 (select sum(classnumber) from onlinecourserecord as ocr where ocr.useridpassive = userinfo.pkid))
				  ,0) as SortCondition,
     	 ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState
		from userinformation as userinfo
		where userinfo.identityid = (select pkid from identity where identityname='老师')
		and userinfo.schoolageidt like ? 
			and (select coursesid from remedialcourses recs where recs.userid = userinfo.pkid and ismain = 1 limit 1) in (select cou.pkid from course as cou where cou.coursename like ?)
			and userinfo.UserLevelId in (select ulv.pkid from userlevel as ulv where ulv.levelname like ?)
		    and userinfo.SeniorLocation in 
	    (select cit.pkid from citys as cit where cit.proid in (select prov.pkid from province as prov where prov.proname like ?))		
	order by SortCondition desc  `

var SqlUserinformationAllTeacherByOnline1 string = `select userinfo.*,
	     (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseNameZhu,
	     (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu,
	     ifnull(
				((select count(*) from onlinecourseevaluation as onev where onev.ocrid in (select onrec.pkid from onlinecourserecord as onrec where onrec.useridpassive = userinfo.pkid))+
				 (select sum(classnumber) from onlinecourserecord as ocr where ocr.useridpassive = userinfo.pkid))
				  ,0) as SortCondition,
     	 ifnull((select count(*)from onlinetrylisten as ontry where ontry.tid = userinfo.pkid and (ontry.sid=0 or ontry.sid is null)),0) as OnlineState
		from userinformation as userinfo
		where userinfo.identityid = (select pkid from identity where identityname='老师')`

//28.
//用途：老师模块：点击老师姓名查看老师详情
//查询字段：用户信息表字段，学生总数，课时总数，本月课时总数，学历，主辅导课程
//查询条件：时间，老师主键
//Controller：
//调用方法名：
//参数说明：本月开始时间，本月结束时间，老师主键id
//2015-11-05
//var SqlUserinformationByTid string = `select userinfo.* ,ifnull((select count(*) from userinformation as uuf where uuf.pkid in
//							(select useridactive from onlineeducation.onlinecourserecord as ocb where ocb.UserIdPassive = userinfo.pkid group by useridactive)),0) as AllPerson,
//					   ifnull((select sum(oc.ClassNumber) from onlineeducation.onlinecourserecord as oc where oc.useridpassive = userinfo.pkid),0) as AllTime,
//	                   ifnull((select sum(ocr.ClassNumber) from onlineeducation.onlinecourserecord as ocr where ocr.useridpassive = userinfo.pkid and ocr.starttime between ? and ?),0) as AllTimeMouth,
//	                   (select DegreeName from degree as deg where userinfo.UserDegree = deg.pkid) as DegreeName,
//					   (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseName,
//				   (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=0 limit 1)) as CourseNameFu
//	from userinformation as userinfo
//	where userinfo.pkid = ?`

//29.
//用途：老师模块：查询老师的所有在线课程评价
//查询字段：在线课程表字段，学生主键，学生姓名，头像路径，两个星级和
//查询条件：老师主键，分页
//Controller：
//调用方法名：
//参数说明：老师用户主键id，从第几行开始，获取多少行
//2015-11-05
var SqlOnlineCourseEvalAllByTid string = `select online.*,userinfo.pkid,userinfo.UserName,userinfo.AvatarPath ,ifnull((online.startclear+online.startefficiency),0) as AllStart
	 from onlinecourseevaluation as online join userinformation as userinfo 
	 on  online.userid = userinfo.pkid 
	 where online.OCRId in (select ocr.pkid from onlinecourserecord as ocr where useridpassive = ?) 
     order by reviewtime desc  `

//30.
//用途：老师模块：查询老师一段时间内的预约课程
//查询字段：在线预约表字段
//查询条件：时间，老师主键
//Controller：
//调用方法名：
//参数说明：开始时间，结束时间，老师主键id
//2015-11-05
var SqlOnlineCourseBookingTimeByTid string = `select ocb.* 
	from onlinecoursebooking as ocb 
	where (starttime between '?' and '?' or endtime between '?' and '?') 
	      and ocb.useridpassive = ?`

//31.
//用途：学生登录时，老师模块，我浏览过谁
//查询字段：师生关系表字段，老师姓名，头像，课时价格，所教年级，主辅导课程
//查询条件：关键词，学生主键，分页，排序
//Controller：
//调用方法名：
//参数说明：学生主键id，关键词：浏览，从第第行开始，获取多少行
//2015-11-05
var SqlRelationsByUidSee string = `select rel.*,userinfo.UserName,userinfo.AvatarPath,userinfo.UnitPrice,
		(select grd.GradeName from grade as grd where grd.pkid = userinfo.gradeid) as GradeName,
	    (select CourseName from course as cous where cous.PKId = (select CoursesId from remedialcourses as remec where remec.UserId=userinfo.PKId and IsMain=1 limit 1)) as CourseName 
		from relations as rel join userinformation as userinfo on rel.FrontUserId = userinfo.pkid 
		where afteruserid = ? and sources like ? 
		order by setdate desc  `

//32.
//用途：精彩问答模块，获取所有精彩问答
//查询字段：问答表主键，标题，内容，时间，浏览人数，我是否收藏，收藏总数
//查询条件：精彩，分页，排序
//Controller：
//调用方法名：
//参数说明：
//2015-11-05
var SqlQuestionAskJingcai string = `select questionask.PKId,questionask.Title,questionask.Contents,questionask.BadeTime,
		(select count(*) from browsecollection where browsecollection.QAId= questionask.PKId ) as Numbers,
		ifnull((select browsecollection.KeepState from browsecollection where UserId=3 && QAId= questionask.PKId),0) as State,
		(select count(*) from browsecollection where browsecollection.QAId= questionask.PKId && browsecollection.KeepState=1 ) as Count 
	from questionask left join wanderfulqa  on questionask.PKId= wanderfulqa.QAId 
	where wanderfulqa.IsWonderful=1 
	order by questionask.BadeTime desc    `

//33.
//用途：查询一条精彩回答详情
//查询字段：问答表字段，回答表字段
//查询条件：问答主键
//Controller：
//调用方法名：
//参数说明：问答信息主键id
//2015-11-05
var SqlQuestionAskById string = `select questionask.Title,questionask.Contents,questionask.BadeTime,questionask.AskUserId,userinformation.UserName,userinformation.AvatarPath,questionask.AnswerUserId,
		(select userinformation.UserName from userinformation where userinformation.PKId= questionask.AnswerUserId) as Hname,
		(select userinformation.AvatarPath from userinformation where userinformation.PKId= questionask.AnswerUserId) as HuiDaAvatarPath,
		(select userinformation.SchoolName from userinformation where userinformation.PKId= questionask.AskUserId) as UserSchoolName,
		(select userinformation.SchoolName from userinformation where userinformation.PKId= questionask.AnswerUserId) as SchoolName,
		answers.Contents as HuiDaContents,answers.AnsTime ,answers.pkid as AnswerId,questionask.PKId
		from questionask left join userinformation on questionask.AskUserId= userinformation.PKId 
		left join answers on answers.QAId= questionask.PKId 
		where questionask.PKId=?`

/**34.学生查看自己所评价的老师们**/
var SqlOnlineCourseEvaluationBySid string = `
			select online.*,(select userinfo.UserName from userinformation as userinfo where userinfo.pkid=
				(select onc.UserIdPassive from onlinecourserecord as onc where onc.pkid = online.ocrid)) as UserName,
                (select userinfo.pkid from userinformation as userinfo where userinfo.pkid=
				(select onc.UserIdPassive from onlinecourserecord as onc where onc.pkid = online.ocrid)) as UserPkid,
                (select userinfo.AvatarPath from userinformation as userinfo where userinfo.pkid=
				(select onc.UserIdPassive from onlinecourserecord as onc where onc.pkid = online.ocrid)) as AvatarPath
			from onlinecourseevaluation as online 
			where online.userid = ? `

/**35.查询学生全部已经冻结的资金总和**/
var SqlFrozenFundsByUserid string = `select sum(frozenmoney) as FrozenMoney 
			from frozenfunds 
			where frozenstate = 1 and userid =?`

/**36.根据手机号码获取一条最新的验证码信息**/
var SqlVerificationByPhone string = `select * 
			from verification
			where userphone = ?
			order by createtime desc
			limit 1`

/**37.查询预约课程附件信息**/
var SqlCoursewareByOCBID string = `select * from courseware where ocbrid = ?`

/**38.根据老师主键id，和时间段查询此时间段预约课程信息**/
var SqlOnlineBookingByTidTime string = `
						select * 
						from onlinecoursebooking
						where UserIdPassive = ? and StartTime > ?`

/**39.查询学生预约课程信息相关的冻结信息，条件：用户主键id，是预约0还是提问1，预约id或提问id**/
var SqlFrozenfundsByUidOnId string = `select * 
						from frozenfunds 
						where userid =? and FrozenType=? and businessid= ? and frozenstate=1 
						limit 1`

/**40.查询给我上过课的老师们**/
var SqlOnlineCourseRecordTByUid string = `select oncour.*,(select username from userinformation as userinfo where userinfo.pkid = oncour.useridpassive)as UserName  
						from onlinecourserecord as oncour 
						where oncour.useridactive = ? 
						group by oncour.useridpassive `

/**41.查询老师的试听信息**/
var SqlOnlineTryListenByTid string = `select * ,(select username from userinformation as user where user.pkid = ontry.sid) as UserName 
						from onlinetrylisten as ontry 
						where tid = ?  and sid is not null  and sid >0 `

/*42.查询老师一条在线信息*/
var SqlOnlinetrylistenOn string = `select * 
						from onlinetrylisten as ontry 
						where tid = ? and (sid is null or sid = 0)
						limit 1 `

/**43.查询给我上过课的某个学科的老师们(参数：学生id,课程id)**/
var SqlOnlineCourseRecordTByUCid string = `select oncour.*,(select username from userinformation as userinfo where userinfo.pkid = oncour.useridpassive)as UserName
						from onlinecourserecord as oncour
						where oncour.useridactive = ? and (select CoursesId from remedialcourses as remec where remec.UserId=oncour.useridpassive and IsMain=1 limit 1)=?
						group by oncour.useridpassive`

/**44.查询这个学生试听过这个老师几次课程**/
var SqlOnlineTrylistenBtidsid string = `select ontry.*
						from onlinetrylisten as ontry
						where ontry.tid = ? and sid= ?`

/**45.查询学生最后一条试听信息，学生试听结束时记录结束时间到此条信息中**/
var SqlOnlineTrylistenBysidLast string = `SELECT * FROM onlineeducation.onlinetrylisten
						where sid = ? 
						order by StuStartTime desc
						limit 1`

/**46.查询老师或学生一条课堂时间记录，一条时间最近且结束时间为null的记录**/
var SqlOnlineBookingRecord string = `select * 
						from onlinecoursebookingrecord
						where userid = ? and ocbid = ? and EndTime is null
						order by starttime desc
						limit 1`

/**47.查询老师或学生关于某次课程的全部课程时间记录信息**/
var SqlOnlineBookingRecordBybookiduid string = `select * 
						from onlinecoursebookingrecord
						where userid = ? and ocbid = ?`

/**48.**/
var SqlUserMessagebymuid string = `SELECT * FROM onlineeducation.usermessage where (messageid = ? and activeuserid = ?) or pkid=?`

/**49.**/
var SqlOnlineBooningbyid string = `SELECT * FROM onlineeducation.onlinecoursebooking where pkid=?`

/**50.查询一条学生试听结束时间为空的最后一条信息**/
var SqlOnlinetrylistenbysid = `select * 
							from onlinetrylisten as ontry
							where sid=? and stuendtime is null
							order by stustarttime desc
							limit 1`

/**51.查询全部学生推荐信息**/
var SqlRecommendTeacherAll = `select rt.*,userinfo.UserName,userinfo.Mailbox,userinfo.ParentMailbox,userinfo.IphoneNum,
							(select CourseName from course as cou where cou.pkid = rt.classid) as CourseName,
							(select GradeName from grade as gr where gr.pkid = rt.gradeid) as GradeName,
							(select CityName from citys as ci where ci.pkid = rt.cityid) as CityName
							from  recommendteacher as rt JOIN userinformation as userinfo on rt.userid = userinfo.pkid`

/**52.管理员查询全部老师信息**/
var SqlUserinformationAllByAdmin = `select userinfo.* ,(select levelname from userlevel as ul where ul.pkid = userinfo.userlevelid) as LevelName,
							(select DegreeName from degree as de where de.pkid = userinfo.userdegree) as DegreeName
							from userinformation as userinfo
							where IdentityId = (select id.pkid from identity as id where IdentityName='老师')`
