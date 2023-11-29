package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"gorm.io/gorm"
)

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

type AdminRepository struct {
	DB *gorm.DB
}

func (r *AdminRepository) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}
