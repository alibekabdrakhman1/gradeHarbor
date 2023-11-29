package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/jmoiron/sqlx"
)

func NewClassRepository(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{
		DB: db,
	}
}

type ClassRepository struct {
	DB *sqlx.DB
}

func (r *ClassRepository) GetClass(ctx context.Context, id uint) (model.ClassCreate, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) Create(ctx context.Context, class model.ClassCreate) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) DeleteClass(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
