package http

import (
	"encoding/json"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type Manager struct {
	UserToken IUserTokenHandler
}

type UserTokenHandler struct {
	Service *service.Service
}

func NewUserTokenHandler(s *service.Service) *UserTokenHandler {
	return &UserTokenHandler{
		Service: s,
	}
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{NewUserTokenHandler(srv)}
}

func (h *UserTokenHandler) Login(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	var request model.Login

	err = json.Unmarshal(body, &request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	userToken, err := h.Service.UserToken.Login(c.Request().Context(), request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	return c.JSON(http.StatusCreated, response)
}
func (h *UserTokenHandler) Register(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request Body reading error")
	}

	var request model.Register
	err = json.Unmarshal(body, &request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Register model unmarshalling error")
	}

	userId, err := h.Service.UserToken.Register(c.Request().Context(), request)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, userId)
}

func (h *UserTokenHandler) RenewToken(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *UserTokenHandler) Verify(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

type IUserTokenHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	RenewToken(c echo.Context) error
	Verify(c echo.Context) error
}
