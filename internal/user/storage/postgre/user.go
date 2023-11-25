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

func (r *UserRepository) Update(ctx context.Context, user model.User, userID uint) (*model.User, error) {
	return &model.User{}, nil
}

func (r *UserRepository) Delete(ctx context.Context, userID uint) error {
	return nil
}

func (r *UserRepository) GetById(ctx context.Context, userID uint) (*model.ResponseUser, error) {
	var res model.ResponseUser
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&res).Error

	return &res, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error

	return &res, err
}
