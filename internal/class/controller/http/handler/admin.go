package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AdminHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func (s *AdminHandler) CreateClass(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) GetAllClasses(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) GetClassByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) UpdateClassByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) DeleteClassByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) GetClassStudentsByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) GetClassGradesByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminHandler) GetClassTeacherByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewAdminHandler(service *service.Service, logger *zap.SugaredLogger) *AdminHandler {
	return &AdminHandler{service: service, logger: logger}
}
