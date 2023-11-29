package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func (h *AdminHandler) GetAllStudents(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetStudentByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetAllTeachers(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetTeacherByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetAllParents(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetParentByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) CreateClass(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) UpdateClass(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetAllClasses(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetClassByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) DeleteUser(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	err = h.service.User.DeleteByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error deleting by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
	})
}

func (h *AdminHandler) CreateAdmin(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Request Body reading error",
		})
	}
	var request model.User

	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Create model unmarshalling error",
		})
	}
	id, err := h.service.Admin.CreateAdmin(c.Request().Context(), request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("Creating new admin error: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    id,
	})
}

func NewAdminHandler(service *service.Service, logger *zap.SugaredLogger) *AdminHandler {
	return &AdminHandler{service: service, logger: logger}
}
func (h *AdminHandler) convertIdToUint(in string) (uint, error) {
	id, err := strconv.ParseUint(in, 10, 32)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("converting id to uint error: %v", err))
	}

	return uint(id), err
}
