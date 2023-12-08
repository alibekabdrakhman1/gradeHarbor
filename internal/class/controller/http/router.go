package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/alibekabdrakhman1/gradeHarbor/internal/class/docs"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1/class", s.jwt.ValidateAuth)
	userClass := v1.Group("/classes")
	userClass.GET("", s.handler.Class.GetAllClasses)
	userClass.GET("/:id", s.handler.Class.GetClassByID)
	userClass.POST("/:id/grades", s.handler.Class.PutClassGradesByID, s.jwt.ValidateTeacher)
	userClass.GET("/:id/students", s.handler.Class.GetClassStudentsByID)
	userClass.GET("/:id/grades", s.handler.Class.GetClassGradesByID)
	userClass.GET("/:id/teacher", s.handler.Class.GetClassTeacherByID)

	admin := v1.Group("/admin", s.jwt.ValidateAdmin)
	adminClass := admin.Group("/classes")
	adminClass.POST("/create", s.handler.Admin.CreateClass)
	adminClass.GET("", s.handler.Admin.GetAllClasses)
	adminClass.GET("/:id", s.handler.Admin.GetClassByID)
	adminClass.PUT("/:id", s.handler.Admin.UpdateClassByID)
	adminClass.DELETE("/:id", s.handler.Admin.DeleteClassByID)
	adminClass.GET("/:id/students", s.handler.Admin.GetClassStudentsByID)
	adminClass.GET("/:id/grades", s.handler.Admin.GetClassGradesByID)
	adminClass.GET("/:id/teacher", s.handler.Admin.GetClassTeacherByID)
	s.App.GET("/swagger/*any", echoSwagger.WrapHandler)
}
