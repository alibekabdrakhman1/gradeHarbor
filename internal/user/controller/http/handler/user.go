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

// Me retrieves information about the authenticated user.
// @Summary Get authenticated user information
// @Description Retrieves information about the authenticated user.
// @ID user-me
// @Tags user
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of authenticated user information response"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Router /v1/user/profile [get]
func (h *UserHandler) Me(c echo.Context) error {
	user, err := h.service.User.GetByContext(c.Request().Context())
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusUnauthorized, response.APIResponse{
			Message: err.Error(),
		})
	}
	h.logger.Info(user)
	return c.JSON(http.StatusOK, response.APIResponse{
		Message: "OK",
		Data:    user,
	})
}

// Update @Summary Update user profile
// @Description Updates the profile of the authenticated user.
// @ID user-update-profile
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body model.User true "User update request payload"
// @Success 200 {object} response.APIResponse "Successful update of user profile response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "User not found"
// @Router /v1/user/profile [put]
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

// Delete @Summary Delete user profile
// @Description Deletes the profile of the authenticated user.
// @ID user-delete-profile
// @Tags user
// @Security ApiKeyAuth
// @Success 204 {object} response.APIResponse "Successful deletion of user profile response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /v1/user/profile [delete]
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

// GetByID @Summary Get user by ID
// @Description Retrieves user details by the provided user ID.
// @ID user-get-by-id
// @Tags user
// @Security ApiKeyAuth
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful user retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "User not found"
// @Router /v1/user/users/{id} [get]
func (h *UserHandler) GetByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
		})
	}
	fmt.Println(id)

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

// GetStudentTeachersByID @Summary Get teachers of a student by ID
// @Description Retrieves the list of teachers associated with the provided student ID.
// @ID user-get-student-teachers-by-id
// @Tags student
// @Security ApiKeyAuth
// @Param id path int true "Student ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful teacher retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "Student not found"
// @Router /v1/user/students/{id}/teachers [get]
func (h *UserHandler) GetStudentTeachersByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
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

// GetStudentGradesByID @Summary Get grades of a student by ID
// @Description Retrieves the list of grades associated with the provided student ID.
// @ID user-get-student-grades-by-id
// @Tags student
// @Security ApiKeyAuth
// @Param id path int true "Student ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful grades retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "Student not found"
// @Router /v1/user/students/{id}/grades [get]
func (h *UserHandler) GetStudentGradesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
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

// GetStudentParentByID @Summary Get parent of a student by ID
// @Description Retrieves the parent associated with the provided student ID.
// @ID user-get-student-parent-by-id
// @Tags student
// @Security ApiKeyAuth
// @Param id path int true "Student ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful parent retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "Student not found"
// @Failure 400 {object} response.APIResponse "User is not a student"
// @Router /v1/user/students/{id}/parent [get]
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
	parent, err := h.service.User.GetStudentParentByID(c.Request().Context(), user.ID)
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

// GetParentChildrenByID @Summary Get children of a parent by ID
// @Description Retrieves the children associated with the provided parent ID.
// @ID user-get-parent-children-by-id
// @Tags parent
// @Security ApiKeyAuth
// @Param id path int true "Parent ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful children retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "Parent not found"
// @Failure 400 {object} response.APIResponse "User is not a parent"
// @Router /v1/user/parents/{id}/children [get]
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

// GetClassesByID @Summary Get classes of a user by ID
// @Description Retrieves the classes associated with the provided user ID.
// @ID user-get-classes-by-id
// @Tags user
// @Security ApiKeyAuth
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} response.APIResponse "Successful classes retrieval response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Failure 404 {object} response.APIResponse "User not found"
// @Router /v1/user/users/{id}/classes [get]
func (h *UserHandler) GetClassesByID(c echo.Context) error {
	id, err := utils.ConvertIdToUint(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		return c.JSON(http.StatusBadRequest, response.APIResponse{
			Message: err.Error(),
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
