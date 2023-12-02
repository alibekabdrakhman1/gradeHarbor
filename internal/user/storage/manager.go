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
	Parent  IParentRepository
	Student IStudentRepository
	Teacher ITeacherRepository
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
	GetProfileById(ctx context.Context, id uint, userID uint) (*model.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetStudentTeachersByID(ctx context.Context, id uint, userID uint) ([]*model.UserResponse, error)
	GetStudentParentByID(ctx context.Context, id uint, userID uint) (*model.ParentResponse, error)
	GetParentChildrenByID(ctx context.Context, id uint, userID uint) ([]*model.UserResponse, error)
}

type IAdminRepository interface {
	GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error)
	GetAllStudents(ctx context.Context) ([]*model.UserResponse, error)
	GetAllParents(ctx context.Context) ([]*model.ParentResponse, error)
	// GetStudentByID CreateClass(ctx context.Context, class model.Class) (uint, error)
	//UpdateClass(ctx context.Context, class model.Class) (model.Class, error)
	//DeleteClass(ctx context.Context, id uint) error
	//GetAllClasses(ctx context.Context) ([]model.Class, error)
	//GetClassByID(ctx context.Context, id uint) (model.Class, error)
	GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error)
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

type IParentRepository interface {
	GetChildren(ctx context.Context, id uint) ([]*model.UserResponse, error)
}

type IStudentRepository interface {
	GetGroupmates(ctx context.Context, id uint) ([]*model.UserResponse, error)
	//GetGrades(ctx context.Context) error
	GetParent(ctx context.Context, id uint) ([]*model.ParentResponse, error)
	GetTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error)
}

type ITeacherRepository interface {
	GetStudents(ctx context.Context, id uint) ([]*model.UserResponse, error)
}
