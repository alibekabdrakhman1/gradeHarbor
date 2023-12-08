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

// CreateClass creates a new class.
// @Summary Create a new class
// @Description Creates a new class based on the provided data.
// @ID create-class
// @Tags admin
// @Accept json
// @Produce json
// @Param request body model.ClassRequest true "Class creation request payload"
// @Success 200 {object} response.APIResponse "Successful class creation response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes [post]
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

// GetAllClasses retrieves all classes.
// @Summary Get all classes
// @Description Retrieves a list of all classes.
// @ID get-all-classes
// @Tags admin
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of classes response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes [get]
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

// GetClassByID retrieves a class by ID.
// @Summary Get class by ID
// @Description Retrieves a class based on the provided ID.
// @ID get-class-by-id
// @Tags admin
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} response.APIResponse "Successful retrieval of class by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id} [get]
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

// UpdateClassByID updates a class by ID.
// @Summary Update class by ID
// @Description Updates a class based on the provided ID and request data.
// @ID update-class-by-id
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Class ID"
// @Param request body model.ClassRequest true "Update class request payload"
// @Success 200 {object} response.APIResponse "Successful update of class by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id} [put]
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

// DeleteClassByID deletes a class by ID.
// @Summary Delete class by ID
// @Description Deletes a class based on the provided ID.
// @ID delete-class-by-id
// @Tags admin
// @Param id path int true "Class ID"
// @Success 200 {object} response.APIResponse "Successful deletion of class by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id} [delete]
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

// GetClassStudentsByID retrieves students of a class by ID.
// @Summary Get class students by ID
// @Description Retrieves students of a class based on the provided ID.
// @ID get-class-students-by-id
// @Tags admin
// @Param id path int true "Class ID"
// @Success 200 {object} response.APIResponse "Successful retrieval of class students by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id}/students [get]
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

// GetClassGradesByID retrieves grades of a class by ID.
// @Summary Get class grades by ID
// @Description Retrieves grades of a class based on the provided ID.
// @ID get-class-grades-by-id
// @Tags admin
// @Param id path int true "Class ID"
// @Success 200 {object} response.APIResponse "Successful retrieval of class grades by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id}/grades [get]
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

// GetClassTeacherByID retrieves the teacher of a class by ID.
// @Summary Get class teacher by ID
// @Description Retrieves the teacher of a class based on the provided ID.
// @ID get-class-teacher-by-id
// @Tags admin
// @Param id path int true "Class ID"
// @Success 200 {object} response.APIResponse "Successful retrieval of class teacher by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/admin/classes/{id}/teacher [get]
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
