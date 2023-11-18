package http

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/labstack/echo/v4"
)

type Manager struct {
	User IUserHandler
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{NewUserHandler(srv)}
}

type IUserHandler interface {
	GetById(c echo.Context) error
	GetByEmail(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
