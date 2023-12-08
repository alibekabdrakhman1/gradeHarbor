package postgre

import (
	"context"
	"errors"

	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
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
	var res []*model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("role = ?", enums.Teacher).Select("id, full_name, email, is_confirmed, role").Scan(&res).Error

	return res, err
}

func (r *AdminRepository) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	var res []*model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("role = ?", enums.Student).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&res).Error

	return res, err
}

func (r *AdminRepository) GetAllParents(ctx context.Context) ([]*model.UserResponse, error) {
	var parents []*model.UserResponse
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("role = ?", enums.Parent).Select("id, full_name, email, is_confirmed, parent_id, role").Scan(&parents).Error

	// for _, parent := range parents {
	//	children, err := r.GetParentChildrenByID(ctx, parent.ID)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	parent.Children = children
	//}

	return parents, err
}

func (r *AdminRepository) DeleteUserByID(ctx context.Context, id uint) error {
	if err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) PutParent(ctx context.Context, studentID uint, parentID uint) error {
	result := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", studentID).Update("parent_id", parentID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("student not found or unable to update parent")
	}

	return nil
}

func (r *AdminRepository) GetUserByID(ctx context.Context, id uint) (*model.UserResponse, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &model.UserResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		IsConfirmed: user.IsConfirmed,
		ParentID:    user.ParentID,
		Role:        user.Role,
	}, err
}

func (r *AdminRepository) GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where("id = (SELECT parent_id FROM users WHERE id = ?)", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	parent := model.ParentResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		IsConfirmed: user.IsConfirmed,
	}
	parent.Children, err = r.GetParentChildrenByID(ctx, parent.ID)
	if err != nil {
		return nil, err
	}

	return &parent, nil
}

func (r *AdminRepository) GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	var children []*model.UserResponse

	if err := r.DB.WithContext(ctx).Model(&model.User{}).Where("parent_id = ?", id).Find(&children).Error; err != nil {
		return nil, err
	}

	return children, nil
}
