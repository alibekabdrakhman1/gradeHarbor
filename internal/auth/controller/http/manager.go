package http

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/labstack/echo/v4"
)

type Manager struct {
	UserToken IUserTokenHandler
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{NewUserTokenHandler(srv)}
}

type IUserTokenHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	RefreshToken(c echo.Context) error
	Verify(c echo.Context) error
}
