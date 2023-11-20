package controller

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/auth/v1")
	v1.POST("/login", s.handler.UserToken.Login)
	v1.POST("/register", s.handler.UserToken.Register)
	v1.POST("/refresh-token", s.handler.UserToken.RefreshToken)
	v1.POST("/verify/:code", s.handler.UserToken.Verify)
}
