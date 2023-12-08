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
	var res []*model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("parent_id = ?", id).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&res).Error

	return res, err
}
