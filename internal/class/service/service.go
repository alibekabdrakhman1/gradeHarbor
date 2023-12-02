package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
	"go.uber.org/zap"
)

type Service struct {
	Admin        IAdminService
	ClassService IClassService
}

func NewManager(repository *storage.Repository, config *config.Config, logger *zap.SugaredLogger) *Service {
	adminService := NewAdminService(repository, config, logger)
	classService := NewClassService(repository, config, logger)
	return &Service{
		Admin:        adminService,
		ClassService: classService,
	}
}
