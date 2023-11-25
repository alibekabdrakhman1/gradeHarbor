package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1", s.jwt.ValidateAuth)
	v1.GET("/profile", s.handler.User.Me)       //+
	v1.DELETE("/delete", s.handler.User.Delete) //+
	v1.PUT("/profile/update", s.handler.User.Update)

	v1.GET("/students", s.handler.User.GetById)
	v1.GET("/students/{id}", s.handler.User.GetById)

	v1.GET("/parents", s.handler.User.GetById)
	v1.GET("/parents/{id}", s.handler.User.GetById)

	student := v1.Group("/student")
	student.GET("/grades", s.handler.User.GetById)   //TODO returns my grades
	student.GET("/teachers", s.handler.User.GetById) //TODO returns my teachers

	parent := v1.Group("/parent/students")
	parent.GET("/", s.handler.User.GetById) //TODO возвращает моих ребенков
	parent.GET("/{id}", s.handler.User.GetById)
	parent.GET("/{id}/grades", s.handler.User.GetById)   //TODO возвращает оценки ребенка
	parent.GET("/{id}/teachers", s.handler.User.GetById) //TODO возвращает преподов ребенка

	admin := v1.Group("/admin")
	admin.POST("/user/delete/{id}", s.handler.User.DeleteByID) //+
	admin.POST("/create", s.handler.User.GetById)              //TODO create new admin
	admin.POST("/parent", s.handler.User.GetById)              //TODO set parent

}

//v1.GET("/classes", s.handler.User.GetById)
//v1.GET("/classes/{id}/students", s.handler.User.GetByEmail) //TODO returns class students
//v1.GET("/classes/{id}/students", s.handler.User.GetByEmail) //TODO returns class students
//parent.GET("/{id}/classes", s.handler.User.GetById)      //TODO возвращает классы ребенка
//parent.GET("/{id}/classes/{id}", s.handler.User.GetById) //TODO возвращает класс по айди ребенка
//admin.POST("/classes", s.handler.User.GetById)                     //TODO create new class
//admin.PUT("/classes/{id}/add/{studentId}", s.handler.User.GetById) //TODO adding new student in class
//admin.GET("/classes", s.handler.User.GetById)                      //TODO getAll classes new class needs to do search and filter
//admin.POST("/classes/{id}", s.handler.User.GetById) //TODO create new class
//admin.GET("/students", s.handler.User.GetById)                     //TODO returns all students //TODO needs to do search, filter and per_page
//admin.GET("/parents", s.handler.User.GetById)                      //TODO returns all parents
//admin.GET("/teachers", s.handler.User.GetById)                     //TODO returns all teachers
//teacher := v1.Group("/teacher")
//teacherClass := teacher.Group("/classes")
//teacherClass.GET("", s.handler.User.GetById)               //TODO returns own classes
//teacherClass.GET("/{id}", s.handler.User.GetById)          //TODO returns class by id
//teacherClass.GET("/{id}/grades", s.handler.User.GetById)   //TODO returns all grades by class_id
//teacherClass.GET("/{id}/students", s.handler.User.GetById) //TODO returns all students in class by class_id
//teacherClass.POST("/{id}/grades", s.handler.User.GetById)  //TODO post grades to class by class_id
