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

func NewAdminHandler(service *service.Service, logger *zap.SugaredLogger) *AdminHandler {
	return &AdminHandler{service: service, logger: logger}
}

func (h *AdminHandler) GetUserByID(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
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
	id, err := h.convertIdToUint(c.Param("id"))
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
	id, err := h.convertIdToUint(c.Param("id"))
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
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetStudentParentByID(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
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
	id, err := h.convertIdToUint(c.Param("id"))
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
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetAllTeachers(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *AdminHandler) GetAllParents(c echo.Context) error {
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

func (h *AdminHandler) PutParent(c echo.Context) error {
	id, err := h.convertIdToUint(c.Param("id"))
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

//func (h *AdminHandler) CreateClass(c echo.Context) error {
//	body, err := ioutil.ReadAll(c.Request().Body)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: "Request Body reading error",
//		})
//	}
//	var request model.Class
//
//	err = json.Unmarshal(body, &request)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: "Class model unmarshalling error",
//		})
//	}
//	id, err := h.service.Admin.CreateClass(c.Request().Context(), request)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: fmt.Sprintf("error creating class: %v", err),
//		})
//	}
//	return c.JSON(http.StatusCreated, response.APIResponse{
//		Message: "OK",
//		Data:    id,
//	})
//}
//
//func (h *AdminHandler) UpdateClass(c echo.Context) error {
//	body, err := ioutil.ReadAll(c.Request().Body)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: "Request Body reading error",
//		})
//	}
//	var request model.Class
//
//	err = json.Unmarshal(body, &request)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: "Class model unmarshalling error",
//		})
//	}
//	class, err := h.service.Admin.UpdateClass(c.Request().Context(), request)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: fmt.Sprintf("error updating class: %v", err),
//		})
//	}
//	return c.JSON(http.StatusCreated, response.APIResponse{
//		Message: "OK",
//		Data:    class,
//	})
//}
//
//func (h *AdminHandler) GetAllClasses(c echo.Context) error {
//	classes, err := h.service.Admin.GetAllClasses(c.Request().Context())
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: fmt.Sprintf("error getting all classes: %v", err),
//		})
//	}
//
//	return c.JSON(http.StatusOK, response.APIResponse{
//		Message: "OK",
//		Data: classes,
//	})
//}
//
//func (h *AdminHandler) GetClassByID(c echo.Context) error {
//	id, err := h.convertIdToUint(c.Param("id"))
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: err.Error(),
//		})
//	}
//	class, err := h.service.Admin.GetClassByID(c.Request().Context(), id)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: fmt.Sprintf("error getting class by id: %v", err),
//		})
//	}
//
//	return c.JSON(http.StatusOK, response.APIResponse{
//		Message: "OK",
//		Data: class,
//	})
//}
//
//func (h *AdminHandler) DeleteClass(c echo.Context) error {
//	id, err := h.convertIdToUint(c.Param("id"))
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: err.Error(),
//		})
//	}
//	err = h.service.Admin.DeleteClass(c.Request().Context(), id)
//	if err != nil {
//		h.logger.Error(err)
//		return c.JSON(http.StatusBadRequest, response.APIResponse{
//			Message: fmt.Sprintf("error deleting class by id: %v", err),
//		})
//	}
//
//	return c.JSON(http.StatusOK, response.APIResponse{
//		Message: "OK",
//	})
//}

func (h *AdminHandler) convertIdToUint(in string) (uint, error) {
	id, err := strconv.ParseUint(in, 10, 32)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("converting id to uint error: %v", err))
	}

	return uint(id), err
}
