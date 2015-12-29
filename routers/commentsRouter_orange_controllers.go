package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"Post",
			`/AddAccountfunds/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"GetAccountfundsByuid",
			`/GetAccountfundsByuid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetAccountfundsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllAccountfunds/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateAccountfundsById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccountfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:AccountfundsController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteAccountfunds/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccounttypesController"] = append(beego.GlobalControllerRouter["orange/controllers:AccounttypesController"],
		beego.ControllerComments{
			"Post",
			`/AddAccounttypes/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccounttypesController"] = append(beego.GlobalControllerRouter["orange/controllers:AccounttypesController"],
		beego.ControllerComments{
			"GetOne",
			`/GetAccounttypesById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccounttypesController"] = append(beego.GlobalControllerRouter["orange/controllers:AccounttypesController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllAccounttypes/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccounttypesController"] = append(beego.GlobalControllerRouter["orange/controllers:AccounttypesController"],
		beego.ControllerComments{
			"Put",
			`/UpdateAccounttypesById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AccounttypesController"] = append(beego.GlobalControllerRouter["orange/controllers:AccounttypesController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteAccounttypes/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"Post",
			`/AddAmountrecords/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"Poststu",
			`/AddAmountrecordsStudent/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetAmountrecordsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsByUserid",
			`/GetAmountrecordsByUserid/:recordtype/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsByUseridCount",
			`/GetAmountrecordsByUseridCount/:recordtype/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsTixianByUserid",
			`/GetAmountrecordsTixianByUserid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsTixianByUseridCount",
			`/GetAmountrecordsTixianByUseridCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsTMcountByUid",
			`/GetAmountrecordsTMcountByUid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsAllT",
			`/GetAmountrecordsAllT/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAmountrecordsAllTCount",
			`/GetAmountrecordsAllTCount/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllAmountrecords/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateAmountrecordsById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"FaFang",
			`/FaFang/:id/:identityid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:AmountrecordsController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteAmountrecords/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AnswersController"] = append(beego.GlobalControllerRouter["orange/controllers:AnswersController"],
		beego.ControllerComments{
			"Post",
			`/AddAnswers/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AnswersController"] = append(beego.GlobalControllerRouter["orange/controllers:AnswersController"],
		beego.ControllerComments{
			"GetOne",
			`/GetAnswersById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AnswersController"] = append(beego.GlobalControllerRouter["orange/controllers:AnswersController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllAnswers/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AnswersController"] = append(beego.GlobalControllerRouter["orange/controllers:AnswersController"],
		beego.ControllerComments{
			"Put",
			`/UpdateAnswersById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:AnswersController"] = append(beego.GlobalControllerRouter["orange/controllers:AnswersController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteAnswers/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"],
		beego.ControllerComments{
			"Post",
			`/AddBrowsecollection/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"],
		beego.ControllerComments{
			"GetOne",
			`/GetBrowsecollectionById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllBrowsecollection/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"],
		beego.ControllerComments{
			"Put",
			`/UpdateBrowsecollectionById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:BrowsecollectionController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteBrowsecollection/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"Post",
			`/AddCitys/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"GetOne",
			`/GetCitysById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"GetCitysByPid",
			`/GetCitysByPid/:pid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllCitys/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"Put",
			`/UpdateCitysById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CitysController"] = append(beego.GlobalControllerRouter["orange/controllers:CitysController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteCitys/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"Post",
			`/AddCountys/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"GetOne",
			`/GetCountysById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"GetCountysByCid",
			`/GetCountysByCid/:cid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllCountys/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"Put",
			`/UpdateCountysById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CountysController"] = append(beego.GlobalControllerRouter["orange/controllers:CountysController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteCountys/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CourseController"] = append(beego.GlobalControllerRouter["orange/controllers:CourseController"],
		beego.ControllerComments{
			"Post",
			`/AddCourse/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CourseController"] = append(beego.GlobalControllerRouter["orange/controllers:CourseController"],
		beego.ControllerComments{
			"GetOne",
			`/GetCourseById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CourseController"] = append(beego.GlobalControllerRouter["orange/controllers:CourseController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllCourse/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CourseController"] = append(beego.GlobalControllerRouter["orange/controllers:CourseController"],
		beego.ControllerComments{
			"Put",
			`/UpdateCourseById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CourseController"] = append(beego.GlobalControllerRouter["orange/controllers:CourseController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteCourse/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"Post",
			`/AddCourseware/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"AddCoursewareOnbook",
			`/AddCoursewareOnbook/:bookid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"GetOne",
			`/GetCoursewareById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"GetCoursewareByOCBID",
			`/GetCoursewareByOCBID/:ocbrid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllCourseware/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"Put",
			`/UpdateCoursewareById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:CoursewareController"] = append(beego.GlobalControllerRouter["orange/controllers:CoursewareController"],
		beego.ControllerComments{
			"DeleteCourseware",
			`/DeleteCourseware/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:DegreeController"] = append(beego.GlobalControllerRouter["orange/controllers:DegreeController"],
		beego.ControllerComments{
			"Post",
			`/AddDegree/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:DegreeController"] = append(beego.GlobalControllerRouter["orange/controllers:DegreeController"],
		beego.ControllerComments{
			"GetOne",
			`/GetDegreeById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:DegreeController"] = append(beego.GlobalControllerRouter["orange/controllers:DegreeController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllDegree/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:DegreeController"] = append(beego.GlobalControllerRouter["orange/controllers:DegreeController"],
		beego.ControllerComments{
			"Put",
			`/UpdateDegreeById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:DegreeController"] = append(beego.GlobalControllerRouter["orange/controllers:DegreeController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteDegree/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:EnclosuresController"] = append(beego.GlobalControllerRouter["orange/controllers:EnclosuresController"],
		beego.ControllerComments{
			"Post",
			`/AddEnclosures/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:EnclosuresController"] = append(beego.GlobalControllerRouter["orange/controllers:EnclosuresController"],
		beego.ControllerComments{
			"GetOne",
			`/GetEnclosuresById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:EnclosuresController"] = append(beego.GlobalControllerRouter["orange/controllers:EnclosuresController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllEnclosures/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:EnclosuresController"] = append(beego.GlobalControllerRouter["orange/controllers:EnclosuresController"],
		beego.ControllerComments{
			"Put",
			`/UpdateEnclosuresById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:EnclosuresController"] = append(beego.GlobalControllerRouter["orange/controllers:EnclosuresController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteEnclosures/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"Post",
			`/AddFrozenfunds/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"AddUserFrozenfunds",
			`/AddUserFrozenfunds/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"UpdateUserFrozenfundsById",
			`/UpdateUserFrozenfundsById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"UpdateUserFrozenfundsByAnswer",
			`/UpdateUserFrozenfundsByAnswer/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"UpdateUserFrozenfundsByOnline",
			`/UpdateUserFrozenfundsByOnline/:id/:tid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"UpdateTeacherFrozenfundsById",
			`/UpdateTeacherFrozenfundsById/:id/tid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetFrozenfundsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"GetFrozenFundsByUserid",
			`/GetFrozenFundsByUserid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"GetFrozenfundsByUidOnId",
			`/GetFrozenfundsByUidOnId/:userid/:typeid/:selid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllFrozenfunds/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateFrozenfundsById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"] = append(beego.GlobalControllerRouter["orange/controllers:FrozenfundsController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteFrozenfunds/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FunctionController"] = append(beego.GlobalControllerRouter["orange/controllers:FunctionController"],
		beego.ControllerComments{
			"Post",
			`/AddFunction/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FunctionController"] = append(beego.GlobalControllerRouter["orange/controllers:FunctionController"],
		beego.ControllerComments{
			"GetOne",
			`/GetFunctionById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FunctionController"] = append(beego.GlobalControllerRouter["orange/controllers:FunctionController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllFunction/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FunctionController"] = append(beego.GlobalControllerRouter["orange/controllers:FunctionController"],
		beego.ControllerComments{
			"Put",
			`/UpdateFunctionById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:FunctionController"] = append(beego.GlobalControllerRouter["orange/controllers:FunctionController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteFunction/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradeController"] = append(beego.GlobalControllerRouter["orange/controllers:GradeController"],
		beego.ControllerComments{
			"Post",
			`/AddGrade/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradeController"] = append(beego.GlobalControllerRouter["orange/controllers:GradeController"],
		beego.ControllerComments{
			"GetOne",
			`/GetGradeById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradeController"] = append(beego.GlobalControllerRouter["orange/controllers:GradeController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllGrade/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradeController"] = append(beego.GlobalControllerRouter["orange/controllers:GradeController"],
		beego.ControllerComments{
			"Put",
			`/UpdateGradeById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradeController"] = append(beego.GlobalControllerRouter["orange/controllers:GradeController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteGrade/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"] = append(beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"],
		beego.ControllerComments{
			"Post",
			`/AddGradecurriculum/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"] = append(beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"],
		beego.ControllerComments{
			"GetOne",
			`/GetGradecurriculumById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"] = append(beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllGradecurriculum/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"] = append(beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"],
		beego.ControllerComments{
			"Put",
			`/UpdateGradecurriculumById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"] = append(beego.GlobalControllerRouter["orange/controllers:GradecurriculumController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteGradecurriculum/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:IdentityController"] = append(beego.GlobalControllerRouter["orange/controllers:IdentityController"],
		beego.ControllerComments{
			"Post",
			`/AddIdentity/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:IdentityController"] = append(beego.GlobalControllerRouter["orange/controllers:IdentityController"],
		beego.ControllerComments{
			"GetOne",
			`/GetIdentityById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:IdentityController"] = append(beego.GlobalControllerRouter["orange/controllers:IdentityController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllIdentity/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:IdentityController"] = append(beego.GlobalControllerRouter["orange/controllers:IdentityController"],
		beego.ControllerComments{
			"Put",
			`/UpdateIdentityById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:IdentityController"] = append(beego.GlobalControllerRouter["orange/controllers:IdentityController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteIdentity/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"Logins",
			`/Logins/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"LoginUser",
			`/LoginUser/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"OutLogins",
			`/OutLogins/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"Registered",
			`/Registered/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"QuestionsCenter",
			`/QuestionsCenter/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UserTeacher",
			`/UserTeacher/:tapid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UserStudent",
			`/UserStudent/:tapid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"OwnerUser",
			`/OwnerUser/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetOnLineEvaluation",
			`/GetOnLineEvaluation/:evalid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetOnLineEvaluationTeacher",
			`/GetOnLineEvaluationTeacher/:evalid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"AddOnLineEvaluation",
			`/AddOnLineEvaluation/:classid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetOnlineCourseBooking",
			`/GetOnlineCourseBooking/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetOnlineCourseBookingByTeacher",
			`/GetOnlineCourseBookingByTeacher/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetUserMessageList",
			`/GetUserMessageList/:msgid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"GetUserMessageListTeacher",
			`/GetUserMessageListTeacher/:msgid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UpdateStudent",
			`/UpdateStudent/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UpdateStudent2",
			`/UpdateStudent/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UpdateTeacher",
			`/UpdateTeacher/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"UpdateTeacher2",
			`/UpdateTeacher/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"RetrievePassword",
			`/RetrievePassword/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"AboutMe",
			`/AboutMe/:tapid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"TechnoRegister",
			`/TechnoRegister/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"PayMentUser",
			`/PayMentUser/:money`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"PayEnd",
			`/PayEnd/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:MainController"] = append(beego.GlobalControllerRouter["orange/controllers:MainController"],
		beego.ControllerComments{
			"PayEndNotify",
			`/PayEndNotify/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"Post",
			`/AddOnlinecoursebooking/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOne",
			`/GetOnlinecoursebookingById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingByTid",
			`/GetOnlinecoursebookingByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingByTidCount",
			`/GetOnlinecoursebookingByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingByUid",
			`/GetOnlinecoursebookingByUid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingByUidCount",
			`/GetOnlinecoursebookingByUidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingBySidNotOn",
			`/GetOnlinecoursebookingBySidNotOn/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingBySidNotOnCount",
			`/GetOnlinecoursebookingBySidNotOnCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingBySTidTime",
			`/GetOnlinecoursebookingBySTidTime/:sid/:tid/:time1/:time2`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingByTidTime",
			`/GetOnlinecoursebookingByTidTime/:userid/:time1`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllOnlinecoursebooking/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"Put",
			`/UpdateOnlinecoursebookingById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"DeleteOnlinecoursebooking",
			`/DeleteOnlinecoursebooking/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"DeleteOnlinecoursebookingMeeting",
			`/DeleteOnlinecoursebookingMeeting/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOss",
			`/GetBHtecher/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOnlineClassTeacherurl",
			`/GetOnlineClassTeacherurl/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetOe",
			`/GetBstudent/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"ClassPay",
			`/ClassPay/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingController"],
		beego.ControllerComments{
			"GetALLtimeminute",
			`/GetALLtimeminute/:oid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"Post",
			`/AddOnlinecoursebookingrecord/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetOne",
			`/GetOnlinecoursebookingrecordById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingrecordByUid",
			`/GetOnlinecoursebookingrecordByUid/:userid/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingrecordByUid2",
			`/GetOnlinecoursebookingrecordByUid2/:userid/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingrecordByTwoid",
			`/GetOnlinecoursebookingrecordByTwoid/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetOnlinecoursebookingrecordBybookiduid",
			`/GetOnlinecoursebookingrecordBybookiduid/:userid/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllOnlinecoursebookingrecord/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"Put",
			`/UpdateOnlinecoursebookingrecordById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecoursebookingrecordController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteOnlinecoursebookingrecord/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"Post",
			`/AddOnlinecourseevaluation/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetOne",
			`/GetOnlinecourseevaluationById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetOnlinecourseevaluationByTid",
			`/GetOnlinecourseevaluationByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetOnlinecourseevaluationByTidCount",
			`/GetOnlinecourseevaluationByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetOnlineCourseEvaluationBySid",
			`/GetOnlineCourseEvaluationBySid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetOnlineCourseEvaluationBySidCount",
			`/GetOnlineCourseEvaluationBySidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllOnlinecourseevaluation/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"Put",
			`/UpdateOnlinecourseevaluationById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourseevaluationController"],
		beego.ControllerComments{
			"DeleteOnlinecourseevaluation",
			`/DeleteOnlinecourseevaluation/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"Post",
			`/AddOnlinecourserecord/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOne",
			`/GetOnlinecourserecordById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordByBookid",
			`/GetOnlinecourserecordByBookid/:bookid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordByTid",
			`/GetOnlinecourserecordByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordByTidCount",
			`/GetOnlinecourserecordByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordByUid",
			`/GetOnlinecourserecordByUid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordByUidCount",
			`/GetOnlinecourserecordByUidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordTeacherByUid",
			`/GetOnlinecourserecordTeacherByUid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetOnlinecourserecordTeacherByUCid",
			`/GetOnlinecourserecordTeacherByUCid/:userid/:classid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllOnlinecourserecord/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"Put",
			`/UpdateOnlinecourserecordById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinecourserecordController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteOnlinecourserecord/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"Post",
			`/AddOnlinetrylisten/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"AddOnlinetrylistenOnline",
			`/AddOnlinetrylistenOnline/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetOne",
			`/GetOnlinetrylistenById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"OnlineTryListenByTid",
			`/OnlineTryListenByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"OnlineTryListenByTidCount",
			`/OnlineTryListenByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"OnlineTryListenByTidSid",
			`/OnlineTryListenByTidSid/:tid/:sid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetOnlinetrylistenOneBysidLast",
			`/GetOnlinetrylistenOneBysidLast/:sid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetOnlinetrylistenOneByTid",
			`/GetOnlinetrylistenOneByTid/:tid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"SetOnlinetrylistenEndTime",
			`/SetOnlinetrylistenEndTime/:sid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllOnlinetrylisten/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"Put",
			`/UpdateOnlinetrylistenById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteOnlinetrylisten/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetOss",
			`/GetListenTecher/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetListenTecherUrl",
			`/GetListenTecherUrl/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetOe",
			`/GetListenStudent/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"] = append(beego.GlobalControllerRouter["orange/controllers:OnlinetrylistenController"],
		beego.ControllerComments{
			"GetAdd",
			`/GetListenStudentAdd/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:PermissionController"] = append(beego.GlobalControllerRouter["orange/controllers:PermissionController"],
		beego.ControllerComments{
			"Post",
			`/AddPermission/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:PermissionController"] = append(beego.GlobalControllerRouter["orange/controllers:PermissionController"],
		beego.ControllerComments{
			"GetOne",
			`/GetPermissionById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:PermissionController"] = append(beego.GlobalControllerRouter["orange/controllers:PermissionController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllPermission/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:PermissionController"] = append(beego.GlobalControllerRouter["orange/controllers:PermissionController"],
		beego.ControllerComments{
			"Put",
			`/UpdatePermissionById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:PermissionController"] = append(beego.GlobalControllerRouter["orange/controllers:PermissionController"],
		beego.ControllerComments{
			"Delete",
			`/DeletePermission/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:ProvinceController"] = append(beego.GlobalControllerRouter["orange/controllers:ProvinceController"],
		beego.ControllerComments{
			"Post",
			`/AddProvince/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:ProvinceController"] = append(beego.GlobalControllerRouter["orange/controllers:ProvinceController"],
		beego.ControllerComments{
			"GetOne",
			`/GetProvinceById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:ProvinceController"] = append(beego.GlobalControllerRouter["orange/controllers:ProvinceController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllProvince/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:ProvinceController"] = append(beego.GlobalControllerRouter["orange/controllers:ProvinceController"],
		beego.ControllerComments{
			"Put",
			`/UpdateProvinceById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:ProvinceController"] = append(beego.GlobalControllerRouter["orange/controllers:ProvinceController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteProvince/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"Post",
			`/AddQuestionask/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetOne",
			`/GetQuestionaskById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskByJingCaiOne",
			`/GetQuestionaskByJingCaiOne/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskByTid",
			`/GetQuestionaskByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskByTidCount",
			`/GetQuestionaskByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskBySid",
			`/GetQuestionaskBySid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskBySidCount",
			`/GetQuestionaskBySidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskByJingCai",
			`/GetQuestionaskByJingCai/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetQuestionaskByJingCaiCount",
			`/GetQuestionaskByJingCaiCount`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllQuestionask/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"Put",
			`/UpdateQuestionaskById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"DeleteQuestionask",
			`/DeleteQuestionask/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:QuestionaskController"] = append(beego.GlobalControllerRouter["orange/controllers:QuestionaskController"],
		beego.ControllerComments{
			"DeleteQuestionaskFStu",
			`/DeleteQuestionaskFStu/:id/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"] = append(beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"],
		beego.ControllerComments{
			"Post",
			`/AddRecommendteacher/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"] = append(beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"],
		beego.ControllerComments{
			"GetOne",
			`/GetRecommendteacherById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"] = append(beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllRecommendteacher/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"] = append(beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"],
		beego.ControllerComments{
			"Put",
			`/UpdateRecommendteacherById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"] = append(beego.GlobalControllerRouter["orange/controllers:RecommendteacherController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteRecommendteacher/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"Post",
			`/AddRelations/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetRelationsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByST",
			`/GetRelationsByST/:sid/:tid/:guanxi`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByTid",
			`/GetRelationsByTid/:userid/:guanxi/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByTidCount",
			`/GetRelationsByTidCount/:userid/:guanxi`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByUid",
			`/GetRelationsByUid/:userid/:guanxi/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByUidCount",
			`/GetRelationsByUidCount/:userid/:guanxi`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByUidSee",
			`/GetRelationsByUidSee/:userid/:guanxi/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetRelationsByUidSeeCount",
			`/GetRelationsByUidSeeCount/:userid/:guanxi`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"AddRelationsBySTGuanxi",
			`/AddRelationsBySTGuanxi/:sid/:tid/:guanxi`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllRelations/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateRelationsById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RelationsController"] = append(beego.GlobalControllerRouter["orange/controllers:RelationsController"],
		beego.ControllerComments{
			"DeleteRelations",
			`/DeleteRelations/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"Post",
			`/AddRemedialcourses/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"GetOne",
			`/GetRemedialcoursesById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"GetRemedialcoursesMain",
			`/GetRemedialcoursesMain/:userid/:ismain`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllRemedialcourses/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"Put",
			`/UpdateRemedialcoursesById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"DeleteRemedialcourses",
			`/DeleteRemedialcourses/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"UpdateStudentClass",
			`/UpdateStudentClass/:sid/:classidlist`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"] = append(beego.GlobalControllerRouter["orange/controllers:RemedialcoursesController"],
		beego.ControllerComments{
			"UpdateStudentClassTeacher",
			`/UpdateStudentClassTeacher/:sid/:classidlist/:mainid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolagesController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolagesController"],
		beego.ControllerComments{
			"Post",
			`/AddSchoolages/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolagesController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolagesController"],
		beego.ControllerComments{
			"GetOne",
			`/GetSchoolagesById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolagesController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolagesController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllSchoolages/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolagesController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolagesController"],
		beego.ControllerComments{
			"Put",
			`/UpdateSchoolagesById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolagesController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolagesController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteSchoolages/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"Post",
			`/AddSchools/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetSchoolsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"GetSchoolsByCity",
			`/GetSchoolsByCity/:cid/:typeid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllSchools/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateSchoolsById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:SchoolsController"] = append(beego.GlobalControllerRouter["orange/controllers:SchoolsController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteSchools/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherList",
			`/TeacherList/:seltype`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherInformation",
			`/TeacherInformation/:tid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherMessage",
			`/TeacherMessage/:tid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"ProblemModel",
			`/ProblemModel/:adkid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"ProblemAnswer",
			`/ProblemAnswer/:adkid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"UserAskQuestion",
			`/UserAskQuestion/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"UserAskQuestion2",
			`/UserAskQuestion/`,
			[]string{"Post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherSetMeet",
			`/TeacherSetMeet/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"StudentSetTeacherMeet",
			`/StudentSetTeacherMeet/:tid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherOnlineClass",
			`/TeacherOnlineClass/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"StudentOnlineClass",
			`/StudentOnlineClass/:onlineid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"TeacherTryListenClass",
			`/TeacherTryListenClass/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"StudentTryListenClass",
			`/StudentTryListenClass/:listenid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"ClassOverHtml",
			`/ClassOverHtml/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TeacherController"] = append(beego.GlobalControllerRouter["orange/controllers:TeacherController"],
		beego.ControllerComments{
			"ListenOverHtml",
			`/ListenOverHtml/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TradingwayController"] = append(beego.GlobalControllerRouter["orange/controllers:TradingwayController"],
		beego.ControllerComments{
			"Post",
			`/AddTradingway/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TradingwayController"] = append(beego.GlobalControllerRouter["orange/controllers:TradingwayController"],
		beego.ControllerComments{
			"GetOne",
			`/GetTradingwayById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TradingwayController"] = append(beego.GlobalControllerRouter["orange/controllers:TradingwayController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllTradingway/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TradingwayController"] = append(beego.GlobalControllerRouter["orange/controllers:TradingwayController"],
		beego.ControllerComments{
			"Put",
			`/UpdateTradingwayById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TradingwayController"] = append(beego.GlobalControllerRouter["orange/controllers:TradingwayController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteTradingway/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"Post",
			`/AddTransactionrecords/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetOne",
			`/GetTransactionrecordsById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetTransactionrecordsByTid",
			`/GetTransactionrecordsByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetTransactionrecordsByTidCount",
			`/GetTransactionrecordsByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetTransactionrecordsBySid",
			`/GetTransactionrecordsBySid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetTransactionrecordsBySidCount",
			`/GetTransactionrecordsBySidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllTransactionrecords/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"Put",
			`/UpdateTransactionrecordsById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"] = append(beego.GlobalControllerRouter["orange/controllers:TransactionrecordsController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteTransactionrecords/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TreatysController"] = append(beego.GlobalControllerRouter["orange/controllers:TreatysController"],
		beego.ControllerComments{
			"Post",
			`/AddTreatys/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TreatysController"] = append(beego.GlobalControllerRouter["orange/controllers:TreatysController"],
		beego.ControllerComments{
			"GetOne",
			`/GetTreatysById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TreatysController"] = append(beego.GlobalControllerRouter["orange/controllers:TreatysController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllTreatys/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TreatysController"] = append(beego.GlobalControllerRouter["orange/controllers:TreatysController"],
		beego.ControllerComments{
			"Put",
			`/UpdateTreatysById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:TreatysController"] = append(beego.GlobalControllerRouter["orange/controllers:TreatysController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteTreatys/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercollectionController"],
		beego.ControllerComments{
			"Post",
			`/AddUsercollection/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercollectionController"],
		beego.ControllerComments{
			"GetOne",
			`/GetUsercollectionById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercollectionController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllUsercollection/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercollectionController"],
		beego.ControllerComments{
			"Put",
			`/UpdateUsercollectionById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercollectionController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercollectionController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteUsercollection/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercourseController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercourseController"],
		beego.ControllerComments{
			"Post",
			`/AddUsercourse/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercourseController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercourseController"],
		beego.ControllerComments{
			"GetOne",
			`/GetUsercourseById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercourseController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercourseController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllUsercourse/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercourseController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercourseController"],
		beego.ControllerComments{
			"Put",
			`/UpdateUsercourseById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsercourseController"] = append(beego.GlobalControllerRouter["orange/controllers:UsercourseController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteUsercourse/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"Post",
			`/AddUserinformation/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetOne",
			`/GetUserinformationById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationByPhone",
			`/GetUserinformationByPhone/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationByUserName",
			`/GetUserinformationByUserName/:name`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationOneByName",
			`/GetUserinformationOneByName/:name`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationPhone",
			`/GetUserinformationPhone/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationByIdPass",
			`/GetUserinformationByIdPass/:id/:pass`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationPicMove",
			`/GetUserinformationPicMove/:count`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationAllTeacher",
			`/GetUserinformationAllTeacher/:seltype/:nianji/:kecheng/:jibie/:shengfen/:shiqu/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetUserinformationAllTeacherCount",
			`/GetUserinformationAllTeacherCount/:seltype/:nianji/:kecheng/:jibie/:shengfen/:shiqu`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllUserinformation/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"Put",
			`/UpdateUserinformationById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteUserinformation/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"UpdateUserimg",
			`/UpdateUserimg/:userid`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserinformationController"] = append(beego.GlobalControllerRouter["orange/controllers:UserinformationController"],
		beego.ControllerComments{
			"UpdateUserimg2",
			`/UpdateUserimg2/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserlevelController"] = append(beego.GlobalControllerRouter["orange/controllers:UserlevelController"],
		beego.ControllerComments{
			"Post",
			`/AddUserlevel/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserlevelController"] = append(beego.GlobalControllerRouter["orange/controllers:UserlevelController"],
		beego.ControllerComments{
			"GetOne",
			`/GetUserlevelById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserlevelController"] = append(beego.GlobalControllerRouter["orange/controllers:UserlevelController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllUserlevel/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserlevelController"] = append(beego.GlobalControllerRouter["orange/controllers:UserlevelController"],
		beego.ControllerComments{
			"Put",
			`/UpdateUserlevelById/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UserlevelController"] = append(beego.GlobalControllerRouter["orange/controllers:UserlevelController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteUserlevel/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"Post",
			`/AddUsermessage/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"AddUsermessageOther",
			`/AddUsermessageOther/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"AddUsermessageTeacherlist",
			`/AddUsermessageTeacherlist/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetOne",
			`/GetUsermessageById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetUsermessageBymuid",
			`/GetUsermessageBymuid/:mid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"UpdateUsermessageBypiduid",
			`/UpdateUsermessageBypiduid/:mid/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetUsermessageByTid",
			`/GetUsermessageByTid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetUsermessageByTidCount",
			`/GetUsermessageByTidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetUsermessageBySid",
			`/GetUsermessageBySid/:userid/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetUsermessageBySidCount",
			`/GetUsermessageBySidCount/:userid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllUsermessage/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"Put",
			`/UpdateUsermessageById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:UsermessageController"] = append(beego.GlobalControllerRouter["orange/controllers:UsermessageController"],
		beego.ControllerComments{
			"DeleteUsermessage",
			`/DeleteUsermessage/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"Post",
			`/AddVerification/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"GetOne",
			`/GetVerificationById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"GetVerificationByPhone",
			`/GetVerificationByPhone/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"GetVerificationListByPhone",
			`/GetVerificationListByPhone/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllVerification/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"Put",
			`/UpdateVerificationById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"DeleteVerification",
			`/DeleteVerification/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"DeleteVerificationByPhone",
			`/DeleteVerificationByPhone/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:VerificationController"] = append(beego.GlobalControllerRouter["orange/controllers:VerificationController"],
		beego.ControllerComments{
			"GetOnees",
			`/GetVerification/:phone`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"] = append(beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"],
		beego.ControllerComments{
			"Post",
			`/AddWanderfulqa/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"] = append(beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"],
		beego.ControllerComments{
			"GetOne",
			`/GetWanderfulqaById/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"] = append(beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"],
		beego.ControllerComments{
			"GetAll",
			`/GetAllWanderfulqa/:page/:size`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"] = append(beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"],
		beego.ControllerComments{
			"Put",
			`/UpdateWanderfulqaById/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"] = append(beego.GlobalControllerRouter["orange/controllers:WanderfulqaController"],
		beego.ControllerComments{
			"Delete",
			`/DeleteWanderfulqa/:id`,
			[]string{"delete"},
			nil})

}
