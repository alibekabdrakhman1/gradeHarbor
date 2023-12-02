package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1/class", s.jwt.ValidateAuth)
	classes := v1.Group("/classes")
	classes.GET("", s.handler.Class.GetAllClasses)
	classes.GET("/{id}", s.handler.Class.GetClassByID)                                      //TODO returns my grades
	classes.POST("/{id}/grades", s.handler.Class.PutClassGradesByID, s.jwt.ValidateTeacher) //TODO post grades to class by class_id
	classes.GET("/{id}/students", s.handler.Class.GetClassStudentsByID)                     //TODO returns class students
	classes.GET("/{id}/grades", s.handler.Class.GetClassGradesByID)                         //TODO returns all grades by class_id
	classes.GET("/{id}/teacher", s.handler.Class.GetClassTeacherByID)

	admin := v1.Group("/admin", s.jwt.ValidateAdmin)
	adminClass := admin.Group("/classes")
	adminClass.POST("/create", s.handler.Admin.CreateClass)  //TODO create new class
	adminClass.GET("", s.handler.Admin.GetAllClasses)        //TODO getAll classes new class needs to do search and filter
	adminClass.GET("/{id}", s.handler.Admin.GetClassByID)    //TODO get class by ID new class needs to do search and filter
	adminClass.PUT("/{id}", s.handler.Admin.UpdateClassByID) //TODO update class
	adminClass.DELETE("/{id}", s.handler.Admin.DeleteClassByID)
	adminClass.GET("/{id}/students", s.handler.Class.GetClassStudentsByID) //TODO returns class students
	adminClass.GET("/{id}/grades", s.handler.Class.GetClassGradesByID)     //TODO returns all grades by class_id
	adminClass.GET("/{id}/teacher", s.handler.Class.GetClassTeacherByID)

}
