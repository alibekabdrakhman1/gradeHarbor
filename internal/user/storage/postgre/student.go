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

func (r *StudentRepository) GetAllParents(ctx context.Context, id uint) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetParentByID(ctx context.Context, id uint, parentID uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetAllTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetTeacherByID(ctx context.Context, id uint, teacherID uint) (*model.TeacherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetAllStudents(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StudentRepository) GetStudentByID(ctx context.Context, id uint, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}
