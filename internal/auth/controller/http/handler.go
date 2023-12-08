package http

import (
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type UserTokenHandler struct {
	Service *service.Service
	logger  *zap.SugaredLogger
}

func NewUserTokenHandler(s *service.Service, logger *zap.SugaredLogger) *UserTokenHandler {
	return &UserTokenHandler{
		Service: s,
		logger:  logger,
	}
}

// Login @Summary Login
// @Description Logs in a user and returns an access token and refresh token.
// @Tags auth
// @ID user-login
// @Accept json
// @Produce json
// @Param request body model.Login true "Login request"
// @Success 201 {object} response.APIResponse "Successful login response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /login [post]
func (h *UserTokenHandler) Login(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	var request model.Login

	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	fmt.Println(request)

	userToken, err := h.Service.UserToken.Login(c.Request().Context(), request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	res := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	h.logger.Info(res)
	return c.JSON(http.StatusCreated, response.APIResponse{
		Message: "OK",
		Data:    res,
	})
}

// Register @Summary User registration
// @Description Registers a new user and returns the user ID.
// @ID user-register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.Register true "Registration request payload"
// @Success 201 {object} response.APIResponse "Successful registration response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /register [post]
func (h *UserTokenHandler) Register(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Request Body reading error",
		})
	}
	fmt.Println(body)
	var request model.Register
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Register model unmarshalling error",
		})
	}

	userId, err := h.Service.UserToken.Register(c.Request().Context(), request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	h.logger.Info(userId)
	return c.JSON(http.StatusCreated, response.APIResponse{
		Message: "OK",
		Data: response.IDResponse{
			ID: userId,
		},
	})
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken @Summary Refresh user token
// @Description Refreshes a user's access token using the provided refresh token.
// @ID user-refresh-token
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshRequest body refreshRequest true "Refresh token request payload"
// @Success 201 {object} response.APIResponse "Successful token refresh response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /refresh-token [post]
func (h *UserTokenHandler) RefreshToken(c echo.Context) error {
	var r refreshRequest
	err := c.Bind(&r)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	tokens, err := h.Service.UserToken.RefreshToken(c.Request().Context(), r.RefreshToken)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	h.logger.Info(tokens)
	return c.JSON(http.StatusCreated, response.APIResponse{
		Message: "OK",
		Data:    tokens,
	})
}

// Confirm @Summary Confirm user registration
// @Description Confirms user registration by providing email and confirmation code.
// @ID user-confirm
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.MessageRequest true "Confirmation request payload"
// @Success 200 {object} response.APIResponse "Successful confirmation response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /confirm [post]
func (h *UserTokenHandler) Confirm(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Request Body reading error",
		})
	}
	var request model.MessageRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Message model unmarshalling error",
		})
	}
	err = h.Service.UserToken.Confirm(c.Request().Context(), request.Email, request.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("user confirm error: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
	})
}
