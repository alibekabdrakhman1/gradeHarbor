package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1/user", s.jwt.ValidateAuth)
	v1.GET("/profile", s.handler.User.Me)            //+
	v1.DELETE("/delete", s.handler.User.Delete)      //+
	v1.PUT("/profile/update", s.handler.User.Update) //+
	v1.GET("/grades", s.handler.User.GetByID)

	v1.GET("/parents", s.handler.User.GetAllParents)      //TODO
	v1.GET("/parents/{id}", s.handler.User.GetParentByID) //TODO

	v1.GET("/teachers", s.handler.User.GetAllTeachers)      //TODO returns my teachers
	v1.GET("/teachers/{id}", s.handler.User.GetAllTeachers) //TODO returns my teachers

	students := v1.Group("/students")
	students.GET("", s.handler.User.GetAllStudents)                       //TODO возвращает моих ребенков
	students.GET("/{id}", s.handler.User.GetStudentByID)                  //TODO
	students.GET("/{id}/teachers", s.handler.User.GetStudentTeachersByID) //TODO возвращает преподов ребенка
	students.GET("/{id}/classes", s.handler.User.GetStudentTeachersByID)  //TODO возвращает преподов ребенка
	students.GET("/{id}/grades", s.handler.User.GetStudentTeachersByID)   //TODO возвращает преподов ребенка

	admin := v1.Group("/admin", s.jwt.ValidateAdmin)
	admin.GET("/parents", s.handler.Admin.GetAllParents)
	admin.GET("/parents/{id}", s.handler.Admin.GetParentByID)
	admin.GET("/teachers", s.handler.Admin.GetAllTeachers)
	admin.GET("/teachers/{id}", s.handler.Admin.GetTeacherByID)
	admin.GET("/students", s.handler.Admin.GetAllStudents)
	admin.GET("/students/{id}", s.handler.Admin.GetStudentByID)
	admin.GET("/students/{id}/grades", s.handler.Admin.GetTeacherByID) //TODO

	admin.POST("/user/delete/{id}", s.handler.Admin.DeleteUser) //+
	admin.POST("/create", s.handler.Admin.CreateAdmin)          //+

	admin.POST("/class", s.handler.Admin.CreateClass)        //TODO create new class
	admin.GET("/classes", s.handler.Admin.GetAllClasses)     //TODO getAll classes new class needs to do search and filter
	admin.GET("/classes/{id}", s.handler.Admin.GetClassByID) //TODO get class by ID new class needs to do search and filter
	admin.PUT("/classes/{id}", s.handler.Admin.UpdateClass)  //TODO update class
	//TODO DELETE
}
