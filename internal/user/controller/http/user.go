package http

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	Service *service.Service
}

func NewUserHandler(s *service.Service) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetById(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
func (h *UserHandler) GetByEmail(c echo.Context) error {
	user, err := h.Service.User.GetByEmail(c.Request().Context(), c.Param("email"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *UserHandler) Delete(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
