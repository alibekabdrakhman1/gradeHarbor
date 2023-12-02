package storage

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage/postgre"
)

type Repository struct {
	Class IClassRepository
	Admin IAdminRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	classRepository := postgre.NewClassRepository(DB)
	adminRepository := postgre.NewAdminRepository(DB)
	return &Repository{
		Class: classRepository,
		Admin: adminRepository,
	}, nil
}

func dsn(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)
}

type IAdminRepository interface {
	CreateClass(ctx context.Context, class model.ClassWithID) (uint, error)
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error)
	DeleteClassByID(ctx context.Context, id uint) error
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
}

type IClassRepository interface {
	GetAllClasses(ctx context.Context) ([]*model.Class, error)
	GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error)
	GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error)
	GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error)
	PutClassGradesByID(ctx context.Context, grades model.GradesRequest) error
	GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error)
}
