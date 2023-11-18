package http

func (s *Server) SetupRoutes() {
	user := s.App.Group("/api/user/v1")

	usersGroup := user.Group("/user")
	usersGroup.GET("/getById/:id", s.handler.User.GetById)
	usersGroup.GET("/getByEmail/:email", s.handler.User.GetByEmail)
	usersGroup.DELETE("/", s.handler.User.Delete)
	usersGroup.PUT("/renew-token", s.handler.User.Update)
}
