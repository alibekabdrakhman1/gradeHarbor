package dto

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
	"go.uber.org/zap"
)

type UserTokenServiceDTO struct {
	Repository        *storage.Repository
	JwtSecretKey      string
	PasswordSecretKey string
	UserHttpTransport *transport.UserHttpTransport
	UserGrpcTransport *transport.UserGrpcTransport
	Logger            *zap.SugaredLogger
}

type UserCode struct {
	UserID uint
	Code   string
}
