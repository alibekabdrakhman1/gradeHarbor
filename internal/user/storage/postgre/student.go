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

func (r *StudentRepository) GetGroupmates(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetParent(ctx context.Context, id uint) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
