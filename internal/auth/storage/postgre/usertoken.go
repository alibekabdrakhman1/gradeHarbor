package postgre

import (
	"context"
	"errors"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserTokenRepository struct {
	DB     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserTokenRepository(db *gorm.DB, logger *zap.SugaredLogger) *UserTokenRepository {
	return &UserTokenRepository{
		DB:     db,
		logger: logger,
	}
}

func (r *UserTokenRepository) CreateUserMessage(ctx context.Context, message model.Message) error {
	err := r.DB.WithContext(ctx).Create(&message).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserTokenRepository) DeleteUserMessage(ctx context.Context, email string) error {
	var res model.Message
	if err := r.DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(ctx).Delete(&res).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserTokenRepository) GetUserMessage(ctx context.Context, email string) (string, error) {
	var res model.Message
	err := r.DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error

	return res.Code, err
}

func (r *UserTokenRepository) CreateUserToken(ctx context.Context, userToken model.UserToken) error {
	var existingToken model.UserToken
	result := r.DB.WithContext(ctx).Where("user_id = ?", userToken.UserID).First(&existingToken)

	if result.Error == nil {
		existingToken.AccessToken = userToken.AccessToken
		existingToken.RefreshToken = userToken.RefreshToken
		result = r.DB.WithContext(ctx).Save(&existingToken)
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := r.DB.WithContext(ctx).Create(&userToken).Error; err != nil {
			r.logger.Errorf("creating new user-token error: %v", err)
			return err
		}
	}
	return nil
}

func (r *UserTokenRepository) UpdateUserToken(ctx context.Context, userToken model.UserToken) error {
	if err := r.DB.WithContext(ctx).Save(&userToken).Error; err != nil {
		r.logger.Errorf("updating user-token error: %v", err)
		return err
	}
	return nil
}
