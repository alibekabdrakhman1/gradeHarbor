package storage

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage/postgre"
)

type Repository struct {
	ClassRepository IClassRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	DB, err := postgre.Dial(ctx, dsn(*cfg))
	if err != nil {
		return nil, err
	}
	classRepository := postgre.NewClassRepository(DB)
	return &Repository{classRepository}, nil
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

type IClassRepository interface {
	GetClass(ctx context.Context, id uint) (model.ClassCreate, error)
	Create(ctx context.Context, class model.ClassCreate) (uint, error)
	DeleteClass(ctx context.Context, id uint) error
}
