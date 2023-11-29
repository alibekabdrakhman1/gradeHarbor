package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/v1/class", s.jwt.ValidateAuth)
	v1.GET("/classes", s.handler.User.GetByID)
	v1.GET("/classes/{id}", s.handler.User.GetStudentGrades)    //TODO returns my grades
	v1.POST("/classes/{id}/grades", s.handler.User.GetByID)     //TODO post grades to class by class_id
	v1.GET("/classes/{id}/students", s.handler.User.GetByEmail) //TODO returns class students
	v1.GET("/classes/{id}/grades", s.handler.User.GetByID)      //TODO returns all grades by class_id
}
