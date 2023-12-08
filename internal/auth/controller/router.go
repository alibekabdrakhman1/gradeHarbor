package controller

import (
	"github.com/swaggo/echo-swagger"

	_ "github.com/alibekabdrakhman1/gradeHarbor/internal/auth/docs"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/auth/v1")
	v1.POST("/login", s.handler.UserToken.Login)
	v1.POST("/register", s.handler.UserToken.Register)
	v1.POST("/refresh-token", s.handler.UserToken.RefreshToken)
	v1.POST("/confirm", s.handler.UserToken.Confirm)
	s.App.GET("/swagger/*any", echoSwagger.WrapHandler)
}
