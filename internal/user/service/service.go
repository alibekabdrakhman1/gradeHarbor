package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
)

type Service struct {
	User IUserService
}

func NewManager(repo *storage.Repository, config *config.Config) *Service {
	userService := NewUserService(repo, config)
	return &Service{
		User: userService,
	}
}
