package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) (uint, error)
	GetByID(ctx context.Context, userID uint) (*model.UserResponse, error)
	GetByContext(ctx context.Context) (*model.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.User) (*model.User, error)
	Delete(ctx context.Context) error
	DeleteByID(ctx context.Context, userID uint) error
}

type IAuthService interface {
	GetJwtUserID(jwtToken string) (*model.ContextUserID, error)
	GetJwtUserRole(jwtToken string) (*model.ContextUserRole, error)
}

type IAdminService interface {
	CreateAdmin(ctx context.Context, user model.User) (uint, error)
	GetAllParents(ctx context.Context) ([]*model.ParentResponse, error)
	GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error)
	GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error)
	GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error)
	GetAllStudents(ctx context.Context) ([]*model.UserResponse, error)
	GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error)
	CreateClass(ctx context.Context, class model.Class) (uint, error)
	GetAllClasses(ctx context.Context) ([]model.Class, error)
	GetClassByID(ctx context.Context, id uint) (model.Class, error)
	UpdateClass(ctx context.Context, id uint) (model.Class, error)
}

type IClientService interface {
	GetAllParents(ctx context.Context) ([]*model.ParentResponse, error)
	GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error)
	GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error)
	GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error)
	GetAllStudents(ctx context.Context) ([]*model.UserResponse, error)
	GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error)
}
