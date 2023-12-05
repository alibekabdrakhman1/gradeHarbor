package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
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

func NewAdminHandler(service *service.Service, logger *zap.SugaredLogger) *AdminHandler {
	return &AdminHandler{service: service, logger: logger}
}

func (h *AdminHandler) GetUserByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.Admin.GetUserByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting user by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *AdminHandler) GetStudentTeachersByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	teachers, err := h.service.Admin.GetStudentTeachersByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting teachers by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    teachers,
	})
}

func (h *AdminHandler) GetUserClassesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	classes, err := h.service.Admin.GetUserClassesByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting classes by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    classes,
	})
}

func (h *AdminHandler) GetStudentGradesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.Admin.GetUserByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	if user.Role != enums.Student {
		h.logger.Error("user is not a student")
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("user with id %v is not a student", user.ID),
		})
	}
	grades, err := h.service.Admin.GetStudentGradesByID(c.Request().Context(), id)
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

func (h *AdminHandler) GetStudentParentByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	parent, err := h.service.Admin.GetStudentParentByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting parent by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    parent,
	})
}

func (h *AdminHandler) GetParentChildrenByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	children, err := h.service.Admin.GetParentChildrenByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error getting children by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    children,
	})
}

func (h *AdminHandler) GetAllStudents(c echo.Context) error {
	students, err := h.service.Admin.GetAllStudents(c.Request().Context())
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

func (h *AdminHandler) GetAllTeachers(c echo.Context) error {
	teachers, err := h.service.Admin.GetAllStudents(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    teachers,
	})
}

func (h *AdminHandler) GetAllParents(c echo.Context) error {
	parents, err := h.service.Admin.GetAllStudents(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    parents,
	})
}

func (h *AdminHandler) DeleteUser(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
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

func (h *AdminHandler) PutParent(c echo.Context) error {
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
	var request model.ParentIDReq

	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: "Parent model unmarshalling error",
		})
	}
	err = h.service.Admin.PutParent(c.Request().Context(), id, request.ID)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("error putting parent by id: %v", err),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
	})
}
