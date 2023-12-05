package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"gorm.io/gorm"
)

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

type StudentRepository struct {
	DB *gorm.DB
}

func (r *StudentRepository) GetParent(ctx context.Context, id uint) (*model.ParentResponse, error) {
	var parent model.ParentResponse

	if err := r.DB.WithContext(ctx).Where("id = (SELECT parent_id FROM users WHERE id = ?)", id).First(&parent).Error; err != nil {
		return nil, err
	}
	var children []*model.UserResponse

	if err := r.DB.WithContext(ctx).Where("parent_id = ?", id).Find(&children).Error; err != nil {
		return nil, err
	}
	parent.Children = children
	return &parent, nil
}
