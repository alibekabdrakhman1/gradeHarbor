package postgre

import (
	"context"
	"errors"
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

func (r *ParentRepository) GetAllParents(ctx context.Context, id uint) ([]*model.ParentResponse, error) {
	return nil, errors.New("you are parent")
}

func (r *ParentRepository) GetParentByID(ctx context.Context, id uint, parentID uint) (*model.ParentResponse, error) {
	return nil, errors.New("you are parent")
}

func (r *ParentRepository) GetAllTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ParentRepository) GetTeacherByID(ctx context.Context, id uint, teacherID uint) (*model.TeacherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ParentRepository) GetAllStudents(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ParentRepository) GetStudentByID(ctx context.Context, id uint, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}
