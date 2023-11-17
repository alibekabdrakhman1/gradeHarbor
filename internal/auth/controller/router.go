package controller

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/api/auth/v1")
	auth.POST("/login", s.handler.UserToken.Login)
	auth.POST("/register", s.handler.UserToken.Register)
	auth.POST("/renew-token", s.handler.UserToken.RenewToken)
	auth.POST("/verify/:code", s.handler.UserToken.Verify)
}
