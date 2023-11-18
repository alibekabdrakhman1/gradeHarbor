package http

import (
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type UserTokenHandler struct {
	Service *service.Service
}

func NewUserTokenHandler(s *service.Service) *UserTokenHandler {
	return &UserTokenHandler{
		Service: s,
	}
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
	fmt.Println(request)

	userToken, err := h.Service.UserToken.Login(c.Request().Context(), request)
	if err != nil {
		fmt.Println(err)
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
	fmt.Println(body)
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
