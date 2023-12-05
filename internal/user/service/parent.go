package service

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
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
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	children, err := s.repository.Parent.GetChildren(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting parent children error: %v", err))
		return nil, fmt.Errorf("getting parent children error: %v", err)
	}

	return children, nil
}
