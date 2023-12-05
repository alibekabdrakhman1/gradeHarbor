package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/response"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type AdminHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func (h *AdminHandler) CreateClass(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Request Body reading error",
		})
	}
	var request model.ClassRequest

	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Create class model unmarshalling error",
		})
	}
	id, err := h.service.Admin.CreateClass(c.Request().Context(), request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("Creating new class error: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    id,
	})
}

func (h *AdminHandler) GetAllClasses(c echo.Context) error {
	classes, err := h.service.Admin.GetAllClasses(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("Getting all classes error: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    classes,
	})
}

func (h *AdminHandler) GetClassByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	class, err := h.service.Admin.GetClassByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting class by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    class,
	})
}

func (h *AdminHandler) UpdateClassByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Request Body reading error",
		})
	}

	var request model.ClassRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Update model unmarshalling error",
		})
	}

	class, err := h.service.Admin.UpdateClassByID(c.Request().Context(), id, request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error updating class by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    class,
	})

}

func (h *AdminHandler) DeleteClassByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	err = h.service.Admin.DeleteClassByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error deleting class by id: %v", err),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
	})
}

func (h *AdminHandler) GetClassStudentsByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	students, err := h.service.Admin.GetClassStudentsByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting class students by id: %v", err),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    students,
	})
}

func (h *AdminHandler) GetClassGradesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	grades, err := h.service.Admin.GetClassGradesByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting class grades by id: %v", err),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    grades,
	})
}

func (h *AdminHandler) GetClassTeacherByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	teacher, err := h.service.Admin.GetClassTeacherByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting class teacher by id: %v", err),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    teacher,
	})
}

func NewAdminHandler(service *service.Service, logger *zap.SugaredLogger) *AdminHandler {
	return &AdminHandler{service: service, logger: logger}
}
