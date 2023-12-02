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
	var oldUser model.User
	if err := r.DB.WithContext(ctx).First(&oldUser, userID).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&oldUser).Updates(user).Error; err != nil {
		return nil, err
	}

	return &oldUser, nil
}

func (r *UserRepository) Delete(ctx context.Context, userID uint) error {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, userID uint) (*model.UserResponse, error) {
	var res model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&res).Error

	return &res, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error

	return &res, err
}

func (r *UserRepository) GetProfileById(ctx context.Context, id uint, userID uint) (*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetStudentTeachersByID(ctx context.Context, id uint, userID uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetStudentParentByID(ctx context.Context, id uint, userID uint) (*model.ParentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetParentChildrenByID(ctx context.Context, id uint, userID uint) ([]*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
