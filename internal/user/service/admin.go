package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"go.uber.org/zap"
)

func NewAdminService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *AdminService {
	return &AdminService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type AdminService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *AdminService) CreateClass(ctx context.Context, class model.Class) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetClassByID(ctx context.Context, id uint) (model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) UpdateClass(ctx context.Context, id uint) (model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) CreateAdmin(ctx context.Context, user model.User) (uint, error) {
	user.IsConfirmed = true
	user.Role = "admin"
	return s.repository.User.Create(ctx, user)
}
