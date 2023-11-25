package storage

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage/postgre"
)

type Repository struct {
	User IUserRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	userRepository := postgre.NewUserRepository(DB)
	return &Repository{User: userRepository}, nil
}

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

type IUserRepository interface {
	Create(ctx context.Context, user model.User) (uint, error)
	Delete(ctx context.Context, userID uint) error
	Update(ctx context.Context, user model.User, userID uint) (*model.User, error)
	GetById(ctx context.Context, userID uint) (*model.ResponseUser, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
