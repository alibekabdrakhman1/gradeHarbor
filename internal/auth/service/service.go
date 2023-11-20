package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/dto"
)

type Service struct {
	UserToken IUserTokenService
}

func NewManager(dto *dto.UserTokenServiceDTO) *Service {
	authService := NewUserTokenService(dto)
	return &Service{
		UserToken: authService,
	}
}
