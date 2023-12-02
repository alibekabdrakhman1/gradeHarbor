package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"gorm.io/gorm"
)

func NewParentRepository(db *gorm.DB) *ParentRepository {
	return &ParentRepository{
		DB: db,
	}
}

type ParentRepository struct {
	DB *gorm.DB
}

func (r *ParentRepository) GetChildren(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
