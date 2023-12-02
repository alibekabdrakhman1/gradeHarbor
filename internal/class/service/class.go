package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
	"go.uber.org/zap"
)

func NewClassService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *ClassService {
	return &ClassService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type ClassService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *ClassService) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassService) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassService) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassService) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClassService) PutClassGradesByID(ctx context.Context, grades model.GradesRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *ClassService) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
