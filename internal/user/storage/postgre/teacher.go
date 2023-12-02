package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"gorm.io/gorm"
)

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{
		DB: db,
	}
}

type TeacherRepository struct {
	DB *gorm.DB
}

func (r *TeacherRepository) GetStudents(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
