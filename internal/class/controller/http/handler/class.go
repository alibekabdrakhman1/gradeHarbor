package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ClassHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func (h *ClassHandler) GetAllClasses(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassHandler) GetClassByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassHandler) GetClassStudentsByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassHandler) GetClassGradesByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassHandler) PutClassGradesByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *ClassHandler) GetClassTeacherByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewClassHandler(service *service.Service, logger *zap.SugaredLogger) *ClassHandler {
	return &ClassHandler{service: service, logger: logger}
}
