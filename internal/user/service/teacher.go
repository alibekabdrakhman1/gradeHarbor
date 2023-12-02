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

func NewTeacherService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *TeacherService {
	return &TeacherService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type TeacherService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *TeacherService) GetStudents(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	students, err := s.repository.Teacher.GetStudents(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting teacher students error: %v", err))
		return nil, fmt.Errorf("getting teacher students error: %v", err)
	}

	return students, nil
}
