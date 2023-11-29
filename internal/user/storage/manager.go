package storage

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage/postgre"
)

type Repository struct {
	User    IUserRepository
	Admin   IAdminRepository
	Parent  IClientRepository
	Student IClientRepository
	Teacher IClientRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	userRepository := postgre.NewUserRepository(DB)
	studentRepository := postgre.NewStudentRepository(DB)
	teacherRepository := postgre.NewTeacherRepository(DB)
	adminRepository := postgre.NewAdminRepository(DB)
	parentRepository := postgre.NewParentRepository(DB)
	return &Repository{
		User:    userRepository,
		Admin:   adminRepository,
		Parent:  parentRepository,
		Student: studentRepository,
		Teacher: teacherRepository,
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

type IUserRepository interface {
	Create(ctx context.Context, user model.User) (uint, error)
	Delete(ctx context.Context, userID uint) error
	Update(ctx context.Context, user model.User, userID uint) (*model.User, error)
	GetById(ctx context.Context, userID uint) (*model.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type IClientRepository interface {
	GetAllParents(ctx context.Context, id uint) ([]*model.ParentResponse, error)
	GetParentByID(ctx context.Context, id uint, parentID uint) (*model.ParentResponse, error)
	GetAllTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error)
	GetTeacherByID(ctx context.Context, id uint, teacherID uint) (*model.TeacherResponse, error)
	GetAllStudents(ctx context.Context, id uint) ([]*model.UserResponse, error)
	GetStudentByID(ctx context.Context, id uint, studentID uint) (*model.StudentResponse, error)
}

type IAdminRepository interface {
	GetAllParents(ctx context.Context) ([]*model.ParentResponse, error)
	GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error)
	GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error)
	GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error)
	GetAllStudents(ctx context.Context) ([]*model.UserResponse, error)
	GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error)
}
