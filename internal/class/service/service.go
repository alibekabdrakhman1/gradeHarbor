package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
	"go.uber.org/zap"
)

type Service struct {
	Admin IAdminService
	Class IClassService
	Auth  IAuthService
}

func NewManager(repository *storage.Repository, config *config.Config, logger *zap.SugaredLogger) *Service {
	adminService := NewAdminService(repository, config, logger)
	classService := NewClassService(repository, config, logger)
	authService := NewAuthService(config.Auth)
	return &Service{
		Admin: adminService,
		Class: classService,
		Auth:  authService,
	}
}
