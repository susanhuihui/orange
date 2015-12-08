// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"orange/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/css", `views/css`)
	beego.SetStaticPath("/images", `views/images`)
	beego.SetStaticPath("/js", `views/js`)
	beego.SetStaticPath("/b/css", `views/b/css`)
	beego.SetStaticPath("/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/Registered/css", `views/css`)
	beego.SetStaticPath("/orange/Main/Registered/images", `views/images`)
	beego.SetStaticPath("/orange/Main/Registered/js", `views/js`)
	beego.SetStaticPath("/orange/Main/Registered/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/Registered/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/Logins/css", `views/css`)
	beego.SetStaticPath("/orange/Main/Logins/images", `views/images`)
	beego.SetStaticPath("/orange/Main/Logins/js", `views/js`)
	beego.SetStaticPath("/orange/Main/Logins/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/Logins/b/js", `views/b/js`)
	beego.SetStaticPath("/Teacher/css", `views/css`)
	beego.SetStaticPath("/Teacher/images", `views/images`)
	beego.SetStaticPath("/Teacher/js", `views/js`)
	beego.SetStaticPath("/Teacher/b/css", `views/b/css`)
	beego.SetStaticPath("/Teacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/QuestionsCenter/css", `views/css`)
	beego.SetStaticPath("/orange/Main/QuestionsCenter/images", `views/images`)
	beego.SetStaticPath("/orange/Main/QuestionsCenter/js", `views/js`)
	beego.SetStaticPath("/orange/Main/QuestionsCenter/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/QuestionsCenter/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/css", `views/css`)
	beego.SetStaticPath("/orange/Main/images", `views/images`)
	beego.SetStaticPath("/orange/Main/js", `views/js`)
	beego.SetStaticPath("/orange/Main/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/grade/AddGrade/css", `views/css`)
	beego.SetStaticPath("/orange/grade/AddGrade/images", `views/images`)
	beego.SetStaticPath("/orange/grade/AddGrade/js", `views/js`)
	beego.SetStaticPath("/orange/grade/AddGrade/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/grade/AddGrade/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherInformation/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherInformation/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherInformation/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherInformation/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherInformation/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherMessage/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherMessage/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherMessage/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherMessage/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherMessage/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/ProblemModel/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/ProblemModel/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/ProblemModel/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/ProblemModel/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/ProblemModel/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/UserAskQuestion/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/UserAskQuestion/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/UserAskQuestion/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/UserAskQuestion/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/UserAskQuestion/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/UserStudent/css", `views/css`)
	beego.SetStaticPath("/orange/Main/UserStudent/images", `views/images`)
	beego.SetStaticPath("/orange/Main/UserStudent/js", `views/js`)
	beego.SetStaticPath("/orange/Main/UserStudent/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/UserStudent/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/UserTeacher/css", `views/css`)
	beego.SetStaticPath("/orange/Main/UserTeacher/images", `views/images`)
	beego.SetStaticPath("/orange/Main/UserTeacher/js", `views/js`)
	beego.SetStaticPath("/orange/Main/UserTeacher/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/UserTeacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluation/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluation/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluation/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluation/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluation/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/AddOnLineEvaluation/css", `views/css`)
	beego.SetStaticPath("/orange/Main/AddOnLineEvaluation/images", `views/images`)
	beego.SetStaticPath("/orange/Main/AddOnLineEvaluation/js", `views/js`)
	beego.SetStaticPath("/orange/Main/AddOnLineEvaluation/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/AddOnLineEvaluation/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBooking/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBooking/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBooking/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBooking/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBooking/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetUserMessageList/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetUserMessageList/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetUserMessageList/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetUserMessageList/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetUserMessageList/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/UpdateStudent/css", `views/css`)
	beego.SetStaticPath("/orange/Main/UpdateStudent/images", `views/images`)
	beego.SetStaticPath("/orange/Main/UpdateStudent/js", `views/js`)
	beego.SetStaticPath("/orange/Main/UpdateStudent/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/UpdateStudent/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/UpdateTeacher/css", `views/css`)
	beego.SetStaticPath("/orange/Main/UpdateTeacher/images", `views/images`)
	beego.SetStaticPath("/orange/Main/UpdateTeacher/js", `views/js`)
	beego.SetStaticPath("/orange/Main/UpdateTeacher/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/UpdateTeacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBookingByTeacher/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBookingByTeacher/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBookingByTeacher/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBookingByTeacher/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetOnlineCourseBookingByTeacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/ProblemAnswer/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/ProblemAnswer/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/ProblemAnswer/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/ProblemAnswer/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/ProblemAnswer/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluationTeacher/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluationTeacher/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluationTeacher/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluationTeacher/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetOnLineEvaluationTeacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/GetUserMessageListTeacher/css", `views/css`)
	beego.SetStaticPath("/orange/Main/GetUserMessageListTeacher/images", `views/images`)
	beego.SetStaticPath("/orange/Main/GetUserMessageListTeacher/js", `views/js`)
	beego.SetStaticPath("/orange/Main/GetUserMessageListTeacher/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/GetUserMessageListTeacher/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherSetMeet/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherSetMeet/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherSetMeet/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherSetMeet/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherSetMeet/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/StudentSetTeacherMeet/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/StudentSetTeacherMeet/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/StudentSetTeacherMeet/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/StudentSetTeacherMeet/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/StudentSetTeacherMeet/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherOnlineClass/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherOnlineClass/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherOnlineClass/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherOnlineClass/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherOnlineClass/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/StudentOnlineClass/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/StudentOnlineClass/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/StudentOnlineClass/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/StudentOnlineClass/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/StudentOnlineClass/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/RetrievePassword/css", `views/css`)
	beego.SetStaticPath("/orange/Main/RetrievePassword/images", `views/images`)
	beego.SetStaticPath("/orange/Main/RetrievePassword/js", `views/js`)
	beego.SetStaticPath("/orange/Main/RetrievePassword/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/RetrievePassword/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherList/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherList/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherList/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherList/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherList/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherTryListenClass/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherTryListenClass/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/TeacherTryListenClass/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/TeacherTryListenClass/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/TeacherTryListenClass/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Teacher/StudentTryListenClass/css", `views/css`)
	beego.SetStaticPath("/orange/Teacher/StudentTryListenClass/images", `views/images`)
	beego.SetStaticPath("/orange/Teacher/StudentTryListenClass/js", `views/js`)
	beego.SetStaticPath("/orange/Teacher/StudentTryListenClass/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Teacher/StudentTryListenClass/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/AboutMe/css", `views/css`)
	beego.SetStaticPath("/orange/Main/AboutMe/images", `views/images`)
	beego.SetStaticPath("/orange/Main/AboutMe/js", `views/js`)
	beego.SetStaticPath("/orange/Main/AboutMe/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/AboutMe/b/js", `views/b/js`)
	beego.SetStaticPath("/orange/Main/TechnologicalProcess/css", `views/css`)
	beego.SetStaticPath("/orange/Main/TechnologicalProcess/images", `views/images`)
	beego.SetStaticPath("/orange/Main/TechnologicalProcess/js", `views/js`)
	beego.SetStaticPath("/orange/Main/TechnologicalProcess/b/css", `views/b/css`)
	beego.SetStaticPath("/orange/Main/TechnologicalProcess/b/js", `views/b/js`)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/Teacher", &controllers.TeacherController{})

	ns := beego.NewNamespace("/orange",

		beego.NSNamespace("/accountfunds",
			beego.NSInclude(
				&controllers.AccountfundsController{},
			),
		),

		beego.NSNamespace("/Main",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),

		beego.NSNamespace("/Teacher",
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),

		beego.NSNamespace("/accounttypes",
			beego.NSInclude(
				&controllers.AccounttypesController{},
			),
		),

		beego.NSNamespace("/amountrecords",
			beego.NSInclude(
				&controllers.AmountrecordsController{},
			),
		),

		beego.NSNamespace("/answers",
			beego.NSInclude(
				&controllers.AnswersController{},
			),
		),

		beego.NSNamespace("/browsecollection",
			beego.NSInclude(
				&controllers.BrowsecollectionController{},
			),
		),

		beego.NSNamespace("/citys",
			beego.NSInclude(
				&controllers.CitysController{},
			),
		),

		beego.NSNamespace("/countys",
			beego.NSInclude(
				&controllers.CountysController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/courseware",
			beego.NSInclude(
				&controllers.CoursewareController{},
			),
		),

		beego.NSNamespace("/degree",
			beego.NSInclude(
				&controllers.DegreeController{},
			),
		),

		beego.NSNamespace("/enclosures",
			beego.NSInclude(
				&controllers.EnclosuresController{},
			),
		),

		beego.NSNamespace("/frozenfunds",
			beego.NSInclude(
				&controllers.FrozenfundsController{},
			),
		),

		beego.NSNamespace("/function",
			beego.NSInclude(
				&controllers.FunctionController{},
			),
		),

		beego.NSNamespace("/grade",
			beego.NSInclude(
				&controllers.GradeController{},
			),
		),

		beego.NSNamespace("/gradecurriculum",
			beego.NSInclude(
				&controllers.GradecurriculumController{},
			),
		),

		beego.NSNamespace("/identity",
			beego.NSInclude(
				&controllers.IdentityController{},
			),
		),

		beego.NSNamespace("/onlinecoursebooking",
			beego.NSInclude(
				&controllers.OnlinecoursebookingController{},
			),
		),

		beego.NSNamespace("/onlinecourseevaluation",
			beego.NSInclude(
				&controllers.OnlinecourseevaluationController{},
			),
		),

		beego.NSNamespace("/onlinecourserecord",
			beego.NSInclude(
				&controllers.OnlinecourserecordController{},
			),
		),

		beego.NSNamespace("/onlinetrylisten",
			beego.NSInclude(
				&controllers.OnlinetrylistenController{},
			),
		),

		beego.NSNamespace("/permission",
			beego.NSInclude(
				&controllers.PermissionController{},
			),
		),

		beego.NSNamespace("/province",
			beego.NSInclude(
				&controllers.ProvinceController{},
			),
		),

		beego.NSNamespace("/questionask",
			beego.NSInclude(
				&controllers.QuestionaskController{},
			),
		),

		beego.NSNamespace("/recommendteacher",
			beego.NSInclude(
				&controllers.RecommendteacherController{},
			),
		),

		beego.NSNamespace("/relations",
			beego.NSInclude(
				&controllers.RelationsController{},
			),
		),

		beego.NSNamespace("/remedialcourses",
			beego.NSInclude(
				&controllers.RemedialcoursesController{},
			),
		),

		beego.NSNamespace("/schoolages",
			beego.NSInclude(
				&controllers.SchoolagesController{},
			),
		),

		beego.NSNamespace("/schools",
			beego.NSInclude(
				&controllers.SchoolsController{},
			),
		),

		beego.NSNamespace("/tradingway",
			beego.NSInclude(
				&controllers.TradingwayController{},
			),
		),

		beego.NSNamespace("/transactionrecords",
			beego.NSInclude(
				&controllers.TransactionrecordsController{},
			),
		),

		beego.NSNamespace("/treatys",
			beego.NSInclude(
				&controllers.TreatysController{},
			),
		),

		beego.NSNamespace("/usercollection",
			beego.NSInclude(
				&controllers.UsercollectionController{},
			),
		),

		beego.NSNamespace("/usercourse",
			beego.NSInclude(
				&controllers.UsercourseController{},
			),
		),

		beego.NSNamespace("/userinformation",
			beego.NSInclude(
				&controllers.UserinformationController{},
			),
		),

		beego.NSNamespace("/userlevel",
			beego.NSInclude(
				&controllers.UserlevelController{},
			),
		),

		beego.NSNamespace("/usermessage",
			beego.NSInclude(
				&controllers.UsermessageController{},
			),
		),

		beego.NSNamespace("/verification",
			beego.NSInclude(
				&controllers.VerificationController{},
			),
		),

		beego.NSNamespace("/wanderfulqa",
			beego.NSInclude(
				&controllers.WanderfulqaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
