package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type StudentHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func NewStudentHandler(service *service.Service, logger *zap.SugaredLogger) *StudentHandler {
	return &StudentHandler{
		service: service,
		logger:  logger,
	}
}

// GetGroupmates @Summary Get groupmates
// @Description Retrieves a list of groupmates for the student.
// @ID student-get-groupmates
// @Tags student
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful groupmates retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/student/groupmates [get]
func (h *StudentHandler) GetGroupmates(c echo.Context) error {
	groupmates, err := h.service.Student.GetGroupmates(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    groupmates,
	})
}

// GetGrades @Summary Get grades
// @Description Retrieves the grades for the student.
// @ID student-get-grades
// @Tags student
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful grades retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/student/grades [get]
func (h *StudentHandler) GetGrades(c echo.Context) error {
	grades, err := h.service.Student.GetGrades(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    grades,
	})
}

// GetParent @Summary Get parent
// @Description Retrieves the parent associated with the student.
// @ID student-get-parent
// @Tags student
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful parent retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/student/parent [get]
func (h *StudentHandler) GetParent(c echo.Context) error {
	parent, err := h.service.Student.GetParent(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    parent,
	})
}

// GetTeachers @Summary Get teachers
// @Description Retrieves the teachers associated with the student.
// @ID student-get-teachers
// @Tags student
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful teachers retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/student/teachers [get]
func (h *StudentHandler) GetTeachers(c echo.Context) error {
	teachers, err := h.service.Student.GetTeachers(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    teachers,
	})
}
