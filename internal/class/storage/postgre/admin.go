package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/jmoiron/sqlx"
)

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

type AdminRepository struct {
	DB *sqlx.DB
}

func (r *AdminRepository) CreateClass(ctx context.Context, class model.ClassWithID) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) DeleteClassByID(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdminRepository) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
