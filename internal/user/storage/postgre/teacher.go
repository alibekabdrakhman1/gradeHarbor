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

func (r *TeacherRepository) GetAllParents(ctx context.Context, id uint) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TeacherRepository) GetParentByID(ctx context.Context, id uint, parentID uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TeacherRepository) GetAllTeachers(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TeacherRepository) GetTeacherByID(ctx context.Context, id uint, teacherID uint) (*model.TeacherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TeacherRepository) GetAllStudents(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TeacherRepository) GetStudentByID(ctx context.Context, id uint, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}
