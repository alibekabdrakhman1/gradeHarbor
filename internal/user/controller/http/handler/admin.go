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

// GetUserByID retrieves a user by ID.
// @Summary Get user by ID
// @Description Retrieves a user based on the provided ID.
// @ID admin-get-user-by-id
// @Tags admin
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of user by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/users/{id} [get]
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

// GetStudentTeachersByID retrieves teachers of a student by ID.
// @Summary Get student teachers by ID
// @Description Retrieves teachers of a student based on the provided ID.
// @ID admin-get-student-teachers-by-id
// @Tags admin
// @Param id path int true "Student ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of student teachers by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/students/{id}/teachers [get]
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

// GetUserClassesByID retrieves classes of a user by ID.
// @Summary Get user classes by ID
// @Description Retrieves classes of a user based on the provided ID.
// @ID admin-get-user-classes-by-id
// @Tags admin
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of user classes by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/users/{id}/classes [get]
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

// GetStudentGradesByID retrieves grades of a student by ID.
// @Summary Get student grades by ID
// @Description Retrieves grades of a student based on the provided ID.
// @ID admin-get-student-grades-by-id
// @Tags admin
// @Param id path int true "Student ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of student grades by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/students/{id}/grades [get]
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

// GetStudentParentByID retrieves parent of a student by ID.
// @Summary Get student parent by ID
// @Description Retrieves parent of a student based on the provided ID.
// @ID admin-get-student-parent-by-id
// @Tags admin
// @Param id path int true "Student ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of student parent by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/students/{id}/parent [get]
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

// GetParentChildrenByID retrieves children of a parent by ID.
// @Summary Get parent children by ID
// @Description Retrieves children of a parent based on the provided ID.
// @ID admin-get-parent-children-by-id
// @Tags admin
// @Param id path int true "Parent ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of parent children by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/parents/{id}/children [get]
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

// GetAllStudents retrieves all students.
// @Summary Get all students
// @Description Retrieves all students.
// @ID admin-get-all-students
// @Tags admin
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of all students response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/students [get]
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

// GetAllTeachers retrieves all teachers.
// @Summary Get all teachers
// @Description Retrieves all teachers.
// @ID admin-get-all-teachers
// @Tags admin
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of all teachers response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/teachers [get]
func (h *AdminHandler) GetAllTeachers(c echo.Context) error {
	teachers, err := h.service.Admin.GetAllTeachers(c.Request().Context())
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

// GetAllParents retrieves all parents.
// @Summary Get all parents
// @Description Retrieves all parents.
// @ID admin-get-all-parents
// @Tags admin
// @Produce json
// @Success 200 {object} response.APIResponse "Successful retrieval of all parents response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/parents [get]
func (h *AdminHandler) GetAllParents(c echo.Context) error {
	parents, err := h.service.Admin.GetAllParents(c.Request().Context())
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

// DeleteUser deletes a user by ID.
// @Summary Delete user by ID
// @Description Deletes a user based on the provided ID.
// @ID admin-delete-user
// @Tags admin
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} response.APIResponse "Successful deletion of user by ID response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/users/{id} [delete]
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

// CreateAdmin creates a new admin user.
// @Summary Create a new admin
// @Description Creates a new admin user based on the provided data.
// @ID admin-create-admin
// @Tags admin
// @Accept json
// @Produce json
// @Param request body model.User true "Admin creation request payload"
// @Success 200 {object} response.APIResponse "Successful creation of new admin response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/admins [post]
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

// PutParent updates the parent of a user by ID.
// @Summary Update user parent by ID
// @Description Updates the parent of a user based on the provided ID and request data.
// @ID admin-put-parent
// @Tags admin
// @Param id path int true "User ID"
// @Accept json
// @Produce json
// @Param request body model.ParentIDReq true "Parent ID request payload"
// @Success 200 {object} response.APIResponse "Successful update of user parent response"
// @Failure 400 {object} response.APIResponse "Bad Request"
// @Router /admin/students/{id}/parent [put]
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
