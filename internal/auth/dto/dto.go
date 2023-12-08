package dto

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/kafka"
	"go.uber.org/zap"
)

type UserTokenServiceDTO struct {
	Repository               *storage.Repository
	JwtSecretKey             string
	PasswordSecretKey        string
	UserHttpTransport        *transport.UserHttpTransport
	UserGrpcTransport        *transport.UserGrpcTransport
	Logger                   *zap.SugaredLogger
	UserVerificationProducer *kafka.Producer
}

type UserCode struct {
	Email string
	Code  string
}
