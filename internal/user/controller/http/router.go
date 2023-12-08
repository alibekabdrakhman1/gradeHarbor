package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/alibekabdrakhman1/gradeHarbor/internal/user/docs"
)

func (s *Server) SetupRoutes() {
	s.App.GET("/swagger/*any", echoSwagger.WrapHandler)
	v1 := s.App.Group("/v1/user", s.jwt.ValidateAuth)
	v1.GET("/profile", s.handler.User.Me)
	v1.DELETE("/profile", s.handler.User.Delete)
	v1.PUT("/profile", s.handler.User.Update)
	v1.GET("/users/:id", s.handler.User.GetByID)
	v1.GET("/users/:id/classes", s.handler.User.GetClassesByID)
	v1.GET("/students/:id/teachers", s.handler.User.GetStudentTeachersByID)
	v1.GET("/students/:id/grades", s.handler.User.GetStudentGradesByID, s.jwt.ValidateParent)
	v1.GET("/students/:id/parent", s.handler.User.GetStudentParentByID, s.jwt.ValidateTeacher)
	v1.GET("/parents/:id/children", s.handler.User.GetParentChildrenByID, s.jwt.ValidateTeacher)

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
	admin.GET("/users/:id", s.handler.Admin.GetUserByID)
	admin.GET("/users/:id/classes", s.handler.Admin.GetUserClassesByID)
	admin.DELETE("/users/:id", s.handler.Admin.DeleteUser)
	admin.GET("/parents", s.handler.Admin.GetAllParents)
	admin.GET("/parents/:id/children", s.handler.Admin.GetParentChildrenByID)
	admin.GET("/teachers", s.handler.Admin.GetAllTeachers)
	admin.GET("/students", s.handler.Admin.GetAllStudents)
	admin.PUT("/students/:id/parent", s.handler.Admin.PutParent)
	admin.GET("/students/:id/parent", s.handler.Admin.GetStudentParentByID)
	admin.GET("/students/:id/teachers", s.handler.Admin.GetStudentTeachersByID)
	admin.GET("/students/:id/grades", s.handler.Admin.GetStudentGradesByID)
	admin.POST("/admins", s.handler.Admin.CreateAdmin)

}
