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
