package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	model2 "github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
)

type IAdminService interface {
	CreateClass(ctx context.Context, class model.ClassRequest) (uint, error)
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error)
	DeleteClassByID(ctx context.Context, id uint) error
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
	GetStudentGradesByID(ctx context.Context, studentID uint) ([]*model.Grade, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
}

type IClassService interface {
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
	PutClassGradesByID(ctx context.Context, id uint, grades model.GradesRequest) error
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
	GetStudentGradesByID(ctx context.Context, studentID uint) ([]*model.Grade, error)
	GetMyStudents(ctx context.Context, userID uint, role string) ([]uint, error)
	GetMyTeachers(ctx context.Context, userID uint) ([]uint, error)
	GetClassesByID(ctx context.Context, userID uint, role string) ([]*model.Class, error)
}

type IAuthService interface {
	GetJwtUserID(jwtToken string) (*model2.ContextUserID, error)
	GetJwtUserRole(jwtToken string) (*model2.ContextUserRole, error)
}
