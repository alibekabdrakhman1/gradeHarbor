package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type TeacherHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func NewTeacherHandler(service *service.Service, logger *zap.SugaredLogger) *TeacherHandler {
	return &TeacherHandler{
		service: service,
		logger:  logger,
	}
}

func (h *TeacherHandler) GetStudents(c echo.Context) error {
	students, err := h.service.Teacher.GetStudents(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    students,
	})
}
