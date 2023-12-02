package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"go.uber.org/zap"
)

func NewStudentService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *StudentService {
	return &StudentService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type StudentService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *StudentService) GetGroupmates(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	gm, err := s.repository.Student.GetGroupmates(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student groupmates error: %v", err))
		return nil, fmt.Errorf("getting student groupmates error: %v", err)
	}

	return gm, nil
}

func (s *StudentService) GetParent(ctx context.Context) ([]*model.ParentResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	parent, err := s.repository.Student.GetParent(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student parent error: %v", err))
		return nil, fmt.Errorf("getting student parent error: %v", err)
	}

	return parent, nil
}

func (s *StudentService) GetTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	teachers, err := s.repository.Student.GetTeachers(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student teachers error: %v", err))
		return nil, fmt.Errorf("getting student teachers error: %v", err)
	}

	return teachers, nil
}
