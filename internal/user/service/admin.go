package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enum"
	"go.uber.org/zap"
)

func NewAdminService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *AdminService {
	return &AdminService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type AdminService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *AdminService) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	teachers, err := s.repository.Admin.GetAllTeachers(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all teachers error: %v", err)
	}
	return teachers, nil
}

func (s *AdminService) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	students, err := s.repository.Admin.GetAllStudents(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all students error: %v", err)
	}
	return students, nil
}

func (s *AdminService) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	parents, err := s.repository.Admin.GetAllParents(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all parents error: %v", err)
	}
	return parents, nil
}

func (s *AdminService) DeleteUserByID(ctx context.Context, id uint) error {
	err := s.repository.Admin.DeleteUserByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return fmt.Errorf("deleting user error: %v", err)
	}
	return nil
}

func (s *AdminService) PutParent(ctx context.Context, studentID uint, parentID uint) error {
	student, err := s.GetUserByID(ctx, studentID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting student by id error: %v", err))
		return err
	}
	if student.Role != enum.Student {
		return errors.New(fmt.Sprintf("user by %v is not student", studentID))
	}
	parent, err := s.GetUserByID(ctx, parentID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting parent by id error: %v", err))
		return err
	}
	if parent.Role != enum.Parent {
		return errors.New(fmt.Sprintf("user by %v is not parent", parentID))
	}
	err = s.repository.Admin.PutParent(ctx, studentID, parentID)
	if err != nil {
		s.logger.Error(err)
		return fmt.Errorf("putting parent for student error: %v", err)
	}
	return nil
}

func (s *AdminService) GetUserByID(ctx context.Context, id uint) (*model.UserResponse, error) {
	user, err := s.repository.Admin.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting user by id error: %v", err)
	}
	return user, nil
}

func (s *AdminService) GetStudentTeachersByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enum.Student {
		return nil, errors.New("this user is not a student")
	}
	teachers, err := s.repository.Admin.GetStudentTeachersByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting student teachers by id error: %v", err)
	}
	return teachers, nil
}

func (s *AdminService) GetUserClassesByID(ctx context.Context, id uint) ([]*model.Class, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role == enum.Parent {
		return nil, errors.New("this user is parent")
	}
	classes, err := s.repository.Admin.GetUserClassesByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting classes by id error: %v", err)
	}
	return classes, nil
}

func (s *AdminService) GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enum.Student {
		return nil, errors.New("this user is not a student")
	}
	parent, err := s.repository.Admin.GetStudentParentByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting student parent by id error: %v", err)
	}
	return parent, nil
}

func (s *AdminService) GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enum.Parent {
		return nil, errors.New("this user is not a parent")
	}
	children, err := s.repository.Admin.GetParentChildrenByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting parent children by id error: %v", err)
	}
	return children, nil
}

func (s *AdminService) CreateAdmin(ctx context.Context, user model.User) (uint, error) {
	user.IsConfirmed = true
	user.Role = "admin"
	return s.repository.User.Create(ctx, user)
}

//func (s *AdminService) CreateClass(ctx context.Context, class model.Class) (uint, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *AdminService) UpdateClass(ctx context.Context, class model.Class) (model.Class, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *AdminService) GetAllClasses(ctx context.Context) ([]model.Class, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *AdminService) GetClassByID(ctx context.Context, id uint) (model.Class, error) {
//	//TODO implement me
//	panic("implement me")
//}
//func (s *AdminService) DeleteClass(ctx context.Context, id uint) error {
//	//TODO implement me
//	panic("implement me")
//}
