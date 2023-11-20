package http

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Manager struct {
	UserToken IUserTokenHandler
}

func NewManager(srv *service.Service, logger *zap.SugaredLogger) *Manager {
	return &Manager{NewUserTokenHandler(srv, logger)}
}

type IUserTokenHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	RefreshToken(c echo.Context) error
	Verify(c echo.Context) error
}
