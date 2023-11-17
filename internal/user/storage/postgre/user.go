package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(ctx context.Context, user model.User) (uint, error) {
	err := r.DB.WithContext(ctx).Create(&user).Error
	return user.ID, err
}

func (r *UserRepository) Update(ctx context.Context, user model.User, userId uint) (model.User, error) {
	return model.User{}, nil
}

func (r *UserRepository) Delete(ctx context.Context, userId uint) error {
	return nil
}

func (r *UserRepository) Get(ctx context.Context, userId uint) (model.User, error) {
	return model.User{}, nil
}
