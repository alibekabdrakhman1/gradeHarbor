package postgre

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"gorm.io/gorm"
)

type UserTokenRepository struct {
	DB *gorm.DB
}

func NewUserTokenRepository(db *gorm.DB) *UserTokenRepository {
	return &UserTokenRepository{
		DB: db,
	}
}
func (r *UserTokenRepository) CreateUserToken(ctx context.Context, userToken model.UserToken) error {
	if err := r.DB.WithContext(ctx).Create(&userToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserTokenRepository) UpdateUserToken(ctx context.Context, userToken model.UserToken) error {
	if err := r.DB.WithContext(ctx).Save(&userToken).Error; err != nil {
		return err
	}
	return nil
}
