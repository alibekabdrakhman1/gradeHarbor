package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1/user", s.jwt.ValidateAuth)
	v1.GET("/profile", s.handler.User.Me)
	v1.DELETE("/delete", s.handler.User.Delete)
	v1.PUT("/profile/update", s.handler.User.Update)
	v1.GET("/profile/{id}", s.handler.User.GetByID)
	v1.GET("/profile/{id}/teachers", s.handler.User.GetStudentTeachersByID)
	v1.GET("/profile/{id}/classes", s.handler.User.GetClassesByID)
	v1.GET("/profile/{id}/grades", s.handler.User.GetStudentGradesByID)
	v1.GET("/profile/{id}/parent", s.handler.User.GetStudentParentByID)
	v1.GET("/profile/{id}/children", s.handler.User.GetParentChildrenByID)

	parent := v1.Group("/parent", s.jwt.ValidateParent)
	parent.GET("/children", s.handler.Parent.GetChildren)

	student := v1.Group("/student", s.jwt.ValidateStudent)
	student.GET("/groupmates", s.handler.Student.GetGroupmates)
	student.GET("/grades", s.handler.Student.GetGrades)
	student.GET("/parent", s.handler.Student.GetParent)
	student.GET("/teachers", s.handler.Student.GetTeachers)

	teacher := v1.Group("/teacher", s.jwt.ValidateTeacher)
	teacher.GET("/students", s.handler.Teacher.GetStudents)

	admin := v1.Group("/admin", s.jwt.ValidateAdmin)
	admin.GET("/profile/{id}", s.handler.Admin.GetUserByID)
	admin.GET("/profile/{id}/teachers", s.handler.Admin.GetStudentTeachersByID)
	admin.GET("/profile/{id}/classes", s.handler.Admin.GetUserClassesByID)
	admin.GET("/profile/{id}/grades", s.handler.User.GetStudentGradesByID)
	admin.GET("/profile/{id}/parent", s.handler.User.GetStudentParentByID)
	admin.GET("/profile/{id}/children", s.handler.User.GetParentChildrenByID)
	admin.GET("/parents", s.handler.Admin.GetAllParents)
	admin.GET("/teachers", s.handler.Admin.GetAllTeachers)
	admin.GET("/students", s.handler.Admin.GetAllStudents)
	admin.PUT("/students/{id}/parent", s.handler.Admin.PutParent)
	admin.POST("/user/delete/{id}", s.handler.Admin.DeleteUser)
	admin.POST("/create/admin", s.handler.Admin.CreateAdmin)

}
