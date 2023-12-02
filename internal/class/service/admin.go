package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
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

func (s *AdminService) CreateClass(ctx context.Context, class model.ClassRequest) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) DeleteClassByID(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AdminService) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
