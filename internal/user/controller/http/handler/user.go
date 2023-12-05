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

type UserHandler struct {
	service *service.Service
	logger  *zap.SugaredLogger
}

func NewUserHandler(s *service.Service, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		service: s,
		logger:  logger,
	}
}

func (h *UserHandler) Me(c echo.Context) error {
	user, err := h.service.User.GetByContext(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	h.logger.Info(user)
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *UserHandler) Update(c echo.Context) error {
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
			Message: "Update model unmarshalling error",
		})
	}

	user, err := h.service.User.Update(c.Request().Context(), request)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusNotFound, response.APIResponse{
			Message: fmt.Sprintf("update err: %v", err),
		})
	}
	h.logger.Info(user)
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *UserHandler) Delete(c echo.Context) error {
	err := h.service.User.Delete(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, response.APIResponse{
		Message: "deleted",
	})
}

func (h *UserHandler) GetByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	user, err := h.service.User.GetByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *UserHandler) GetStudentTeachersByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.User.GetByID(c.Request().Context(), id)
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
	teachers, err := h.service.User.GetStudentTeachersByID(c.Request().Context(), id)
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

func (h *UserHandler) GetStudentGradesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.User.GetByID(c.Request().Context(), id)
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
	grades, err := h.service.User.GetStudentGradesByID(c.Request().Context(), id)
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

func (h *UserHandler) GetStudentParentByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.User.GetByID(c.Request().Context(), id)
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
	parent, err := h.service.User.GetStudentParentByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    parent,
	})
}

func (h *UserHandler) GetParentChildrenByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.User.GetByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	if user.Role != enums.Parent {
		h.logger.Error("user is not a parent")
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("user with id %v is not a parent", user.ID),
		})
	}
	children, err := h.service.User.GetParentChildrenByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    children,
	})
}

func (h *UserHandler) GetClassesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	user, err := h.service.User.GetByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	if user.Role == enums.Parent {
		h.logger.Error("user is not student or teacher")
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: fmt.Sprintf("user with id %v is not student or teacher", user.ID),
		})
	}
	classes, err := h.service.User.GetUserClassesByID(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    classes,
	})
}
