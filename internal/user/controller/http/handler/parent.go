package handler

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type ParentHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func NewParentHandler(service *service.Service, logger *zap.SugaredLogger) *ParentHandler {
	return &ParentHandler{
		service: service,
		logger:  logger,
	}
}

func (h *ParentHandler) GetChildren(c echo.Context) error {
	children, err := h.service.Parent.GetChildren(c.Request().Context())

	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    children,
	})
}
