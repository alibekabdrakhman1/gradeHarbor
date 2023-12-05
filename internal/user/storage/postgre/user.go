package postgre

import (
	"context"
	"errors"
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
	result := r.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("unable to create user")
	}

	return user.ID, nil
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

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id uint) (*model.UserResponse, error) {
	var res model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&res).Error

	return &res, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error

	return &res, err
}

func (r *UserRepository) GetProfileById(ctx context.Context, id uint) (*model.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error) {
	var parent model.ParentResponse

	if err := r.DB.WithContext(ctx).Where("id = (SELECT parent_id FROM users WHERE id = ?)", id).First(&parent).Error; err != nil {
		return nil, err
	}
	var err error
	parent.Children, err = r.GetParentChildrenByID(ctx, parent.ID)
	if err != nil {
		return nil, err
	}

	return &parent, nil
}

func (r *UserRepository) GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	var children []*model.UserResponse

	if err := r.DB.WithContext(ctx).Where("parent_id = ?", id).Find(&children).Error; err != nil {
		return nil, err
	}

	return children, nil
}
