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
	GetStudentTeachersByID(ctx context.Context, userID uint) ([]*model.UserResponse, error)
	GetStudentParentByID(ctx context.Context, userID uint) (*model.ParentResponse, error)
	GetParentChildrenByID(ctx context.Context, userID uint) ([]*model.UserResponse, error)
}

type IAuthService interface {
	GetJwtUserID(jwtToken string) (*model.ContextUserID, error)
	GetJwtUserRole(jwtToken string) (*model.ContextUserRole, error)
}

type IAdminService interface {
	GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error)
	GetAllStudents(ctx context.Context) ([]*model.UserResponse, error)
	GetAllParents(ctx context.Context) ([]*model.ParentResponse, error)
	DeleteUserByID(ctx context.Context, id uint) error
	PutParent(ctx context.Context, studentID uint, parentID uint) error
	CreateAdmin(ctx context.Context, user model.User) (uint, error)
	GetUserByID(ctx context.Context, id uint) (*model.UserResponse, error)
	GetStudentTeachersByID(ctx context.Context, id uint) ([]*model.UserResponse, error)
	GetUserClassesByID(ctx context.Context, id uint) ([]*model.Class, error)
	//GetStudentGradesByID(ctx context.Context, id uint)
	GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error)
	GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error)
}

type IParentService interface {
	GetChildren(ctx context.Context) ([]*model.UserResponse, error)
}

type IStudentService interface {
	GetGroupmates(ctx context.Context) ([]*model.UserResponse, error)
	//GetGrades(ctx context.Context) error
	GetParent(ctx context.Context) ([]*model.ParentResponse, error)
	GetTeachers(ctx context.Context) ([]*model.UserResponse, error)
}

type ITeacherService interface {
	GetStudents(ctx context.Context) ([]*model.UserResponse, error)
}

//CreateClass(ctx context.Context, class model.Class) (uint, error)
//UpdateClass(ctx context.Context, class model.Class) (model.Class, error)
//DeleteClass(ctx context.Context, id uint) error
//GetAllClasses(ctx context.Context) ([]model.Class, error)
//GetClassByID(ctx context.Context, id uint) (model.Class, error)
