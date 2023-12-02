package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/jmoiron/sqlx"
)

func NewClassRepository(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{
		DB: db,
	}
}

type ClassRepository struct {
	DB *sqlx.DB
}

func (r *ClassRepository) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) PutClassGradesByID(ctx context.Context, grades model.GradesRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepository) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
