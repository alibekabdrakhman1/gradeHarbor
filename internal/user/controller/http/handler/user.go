package handler

import (
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *service.Service
}

func NewUserHandler(s *service.Service) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) Me(c echo.Context) error {
	user, err := h.Service.User.GetByContext(c.Request().Context())
	fmt.Println(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetById(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.Service.User.GetByID(c.Request().Context(), id)
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
	return h.Service.User.Delete(c.Request().Context())
}

func (h *UserHandler) DeleteByID(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = h.Service.User.DeleteByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, "deleted")
}

func (h *UserHandler) convertIdToUint(in string) (uint, error) {
	id, err := strconv.ParseUint(in, 10, 32)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("converting id to uint error: %v", err))
	}

	return uint(id), err
}
