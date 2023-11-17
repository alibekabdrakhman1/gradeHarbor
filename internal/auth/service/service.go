package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
)

type Service struct {
	UserToken IUserTokenService
}

func NewManager(repo *storage.Repository, authConfig config.Auth, userHttpTransport *transport.UserHttpTransport, userGrpcTransport *transport.UserGrpcTransport) *Service {
	authService := NewUserTokenService(repo, authConfig, userHttpTransport, userGrpcTransport)
	return &Service{
		UserToken: authService,
	}
}
