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

func (r *AdminRepository) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) DeleteUserByID(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) PutParent(ctx context.Context, studentID uint, parentID uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) CreateAdmin(ctx context.Context, user model.User) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetUserByID(ctx context.Context, id uint) (*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetStudentTeachersByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetUserClassesByID(ctx context.Context, id uint) ([]*model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
