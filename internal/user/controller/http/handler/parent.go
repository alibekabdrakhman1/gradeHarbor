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

// GetChildren @Summary Get children
// @Description Retrieves the children associated with the parent.
// @ID parent-get-children
// @Tags parent
// @Security ApiKeyAuth
// @Success 200 {object} response.APIResponse "Successful children retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/parent/children [get]
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
