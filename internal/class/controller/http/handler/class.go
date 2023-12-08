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

type ClassHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

// GetAllClasses retrieves all classes.
// @Summary Get all classes
// @Description Retrieves all classes.
// @ID user-get-all-classes
// @Tags user
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of all classes response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes [get]
func (h *ClassHandler) GetAllClasses(c echo.Context) error {
	classes, err := h.service.Class.GetAllClasses(c.Request().Context())
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
// @ID user-get-class-by-id
// @Tags user
// @Param id path int true "Class ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of class by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes/{id} [get]
func (h *ClassHandler) GetClassByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	class, err := h.service.Class.GetClassByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    class,
	})
}

// GetClassStudentsByID retrieves students of a class by ID.
// @Summary Get class students by ID
// @Description Retrieves students of a class based on the provided ID.
// @ID user-get-class-students-by-id
// @Tags user
// @Param id path int true "Class ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of class students by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes/{id}/students [get]
func (h *ClassHandler) GetClassStudentsByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	students, err := h.service.Class.GetClassStudentsByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
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
// @ID user-get-class-grades-by-id
// @Tags user
// @Param id path int true "Class ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of class grades by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes/{id}/grades [get]
func (h *ClassHandler) GetClassGradesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	grades, err := h.service.Class.GetClassGradesByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    grades,
	})
}

// PutClassGradesByID updates grades of a class by ID.
// @Summary Update class grades by ID
// @Description Updates grades of a class based on the provided ID and request data.
// @ID user-put-class-grades-by-id
// @Tags user
// @Param id path int true "Class ID"
// @Accept json
// @Success 200 {object} response.APIResponse "Successful update of class grades by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes/{id}/grades [put]
func (h *ClassHandler) PutClassGradesByID(c echo.Context) error {
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
			Message: err.Error(),
		})
	}

	var request model.GradesRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	err = h.service.Class.PutClassGradesByID(c.Request().Context(), id, request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
	})
}

// GetClassTeacherByID retrieves the teacher of a class by ID.
// @Summary Get class teacher by ID
// @Description Retrieves the teacher of a class based on the provided ID.
// @ID user-get-class-teacher-by-id
// @Tags user
// @Param id path int true "Class ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of class teacher by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/class/classes/{id}/teacher [get]
func (h *ClassHandler) GetClassTeacherByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	teacher, err := h.service.Class.GetClassTeacherByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    teacher,
	})
}

func NewClassHandler(service *service.Service, logger *zap.SugaredLogger) *ClassHandler {
	return &ClassHandler{service: service, logger: logger}
}
