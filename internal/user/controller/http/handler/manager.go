package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Manager struct {
	User    IUserHandler
	Admin   IAdminHandler
	Teacher ITeacherHandler
	Parent  IParentHandler
	Student IStudentHandler
}

func NewManager(srv *service.Service, logger *zap.SugaredLogger) *Manager {
	return &Manager{
		User:    NewUserHandler(srv, logger),
		Admin:   NewAdminHandler(srv, logger),
		Teacher: NewTeacherHandler(srv, logger),
		Parent:  NewParentHandler(srv, logger),
		Student: NewStudentHandler(srv, logger),
	}
}

type IUserHandler interface {
	Me(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetByID(c echo.Context) error
	GetStudentTeachersByID(c echo.Context) error
	GetClassesByID(c echo.Context) error
	GetStudentGradesByID(c echo.Context) error
	GetStudentParentByID(c echo.Context) error
	GetParentChildrenByID(c echo.Context) error
}

type IParentHandler interface {
	GetChildren(c echo.Context) error
}

type IStudentHandler interface {
	GetGroupmates(c echo.Context) error
	GetGrades(c echo.Context) error
	GetParent(c echo.Context) error
	GetTeachers(c echo.Context) error
}

type ITeacherHandler interface {
	GetStudents(c echo.Context) error
}

type IAdminHandler interface {
	GetAllStudents(c echo.Context) error
	GetAllTeachers(c echo.Context) error
	GetAllParents(c echo.Context) error
	DeleteUser(c echo.Context) error
	PutParent(c echo.Context) error
	CreateAdmin(c echo.Context) error
	GetUserByID(c echo.Context) error
	GetStudentTeachersByID(c echo.Context) error
	GetUserClassesByID(c echo.Context) error
	GetStudentGradesByID(c echo.Context) error
	GetStudentParentByID(c echo.Context) error
	GetParentChildrenByID(c echo.Context) error
}
