package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
)

type Service struct {
	User IUserService
	Auth IAuthService
}

func NewManager(repo *storage.Repository, config *config.Config) *Service {
	userService := NewUserService(repo, config)
	authService := NewAuthService(config.Auth)
	return &Service{
		User: userService,
		Auth: authService,
	}
}
