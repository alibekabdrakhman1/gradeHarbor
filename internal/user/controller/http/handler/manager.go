package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Manager struct {
	User  IUserHandler
	Admin IAdminHandler
}

func NewManager(srv *service.Service, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		User:  NewUserHandler(srv, logger),
		Admin: NewAdminHandler(srv, logger),
	}
}

type IUserHandler interface {
	Me(c echo.Context) error
	GetByID(c echo.Context) error
	GetAllStudents(c echo.Context) error
	GetStudentByID(c echo.Context) error
	GetStudentTeachersByID(c echo.Context) error
	GetAllParents(c echo.Context) error
	GetParentByID(c echo.Context) error
	GetAllTeachers(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	DeleteByID(c echo.Context) error
}

type IAdminHandler interface {
	GetAllStudents(c echo.Context) error
	GetStudentByID(c echo.Context) error
	GetAllTeachers(c echo.Context) error
	GetTeacherByID(c echo.Context) error
	GetAllParents(c echo.Context) error
	GetParentByID(c echo.Context) error
	CreateClass(c echo.Context) error
	UpdateClass(c echo.Context) error
	GetAllClasses(c echo.Context) error
	GetClassByID(c echo.Context) error
	DeleteUser(c echo.Context) error
	CreateAdmin(c echo.Context) error
}
