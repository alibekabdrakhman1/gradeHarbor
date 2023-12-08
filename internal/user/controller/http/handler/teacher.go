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

// GetStudents @Summary Get all students
// @Description Retrieves a list of all students.
// @ID teacher-get-students
// @Tags teacher
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful students retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/teacher/students [get]
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
