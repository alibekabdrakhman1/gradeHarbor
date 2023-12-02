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

func NewParentService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *ParentService {
	return &ParentService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type ParentService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *ParentService) GetChildren(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	children, err := s.repository.Parent.GetChildren(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting parent children error: %v", err))
		return nil, fmt.Errorf("getting parent children error: %v", err)
	}

	return children, nil
}
