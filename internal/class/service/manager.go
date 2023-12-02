package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
)

type IAdminService interface {
	CreateClass(ctx context.Context, class model.ClassRequest) (uint, error)
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error)
	DeleteClassByID(ctx context.Context, id uint) error
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
}

type IClassService interface {
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
	PutClassGradesByID(ctx context.Context, grades model.GradesRequest) error
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
}
