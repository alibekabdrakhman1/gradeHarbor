package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Manager struct {
	Admin IAdminHandler
	Class IClassHandler
}

func NewManager(srv *service.Service, logger *zap.SugaredLogger) *Manager {
	return &Manager{}
}

type IAdminHandler interface {
	CreateClass(c echo.Context) error
	GetAllClasses(c echo.Context) error
	GetClassByID(c echo.Context) error
	UpdateClassByID(c echo.Context) error
	DeleteClassByID(c echo.Context) error
	GetClassStudentsByID(c echo.Context) error
	GetClassGradesByID(c echo.Context) error
	GetClassTeacherByID(c echo.Context) error
}

type IClassHandler interface {
	GetAllClasses(c echo.Context) error
	GetClassByID(c echo.Context) error
	GetClassStudentsByID(c echo.Context) error
	GetClassGradesByID(c echo.Context) error
	PutClassGradesByID(c echo.Context) error
	GetClassTeacherByID(c echo.Context) error
}
