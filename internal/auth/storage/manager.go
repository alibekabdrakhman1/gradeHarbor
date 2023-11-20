package storage

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage/postgre"
	"go.uber.org/zap"
)

func dsn(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)
}

type Repository struct {
	UserToken IUserTokenRepository
}
type IUserTokenRepository interface {
	CreateUserToken(ctx context.Context, userToken model.UserToken) error
	UpdateUserToken(ctx context.Context, userToken model.UserToken) error
}

func NewRepository(ctx context.Context, cfg *config.Config, logger *zap.SugaredLogger) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	userToken := postgre.NewUserTokenRepository(DB, logger)
	return &Repository{UserToken: userToken}, nil
}
