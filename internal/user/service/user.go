package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"go.uber.org/zap"
)

type UserService struct {
	repository *storage.Repository
	logger     *zap.SugaredLogger
}

func NewUserService(r *storage.Repository, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		repository: r,
		logger:     logger,
	}
}

func (s *UserService) Create(ctx context.Context, user model.User) (uint, error) {
	user.IsConfirmed = false
	id, err := s.repository.User.Create(ctx, user)
	if err != nil {
		s.logger.Error(fmt.Sprintf("creating new user error: %v", err))
		return 0, fmt.Errorf("creating new user error: %v", err)
	}
	return id, nil
}

func (s *UserService) GetByID(ctx context.Context, userID uint) (*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	user, err := s.repository.User.GetProfileById(ctx, id.ID, userID)
	if err != nil {
		s.logger.Error(fmt.Errorf("GetProfileByID error: %v", err))
		return nil, fmt.Errorf("GetProfileByID error: %v", err)
	}
	return user, nil
}

func (s *UserService) GetByContext(ctx context.Context) (*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}

	user, err := s.repository.User.GetById(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repository.User.GetByEmail(ctx, email)
}

func (s *UserService) Update(ctx context.Context, user model.User) (*model.User, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	oldUser, err := s.repository.User.GetById(ctx, id.ID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}

	if user.Email != oldUser.Email {
		return nil, errors.New("can not change email")
	}

	return s.repository.User.Update(ctx, user, id.ID)
}

func (s *UserService) Delete(ctx context.Context) error {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return errors.New("not valid context userID")
	}

	err := s.repository.User.Delete(ctx, id.ID)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *UserService) DeleteByID(ctx context.Context, userID uint) error {
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return errors.New("not valid context userRole")
	}

	if role.Role != "admin" {
		return errors.New("not permitted")
	}

	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return err
	}
	if user.Role == "admin" {
		s.logger.Error("cannot delete admin")
		return errors.New("cannot delete admin")
	}

	err = s.repository.User.Delete(ctx, userID)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *UserService) GetStudentTeachersByID(ctx context.Context, userID uint) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != "student" {
		s.logger.Error(errors.New("user is not a student"))
		return nil, errors.New("user is not a student")
	}

	teachers, err := s.repository.User.GetStudentTeachersByID(ctx, id.ID, userID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student teachers error: %v", err))
		return nil, fmt.Errorf("getting student teachers error: %v", err)
	}

	return teachers, nil
}

func (s *UserService) GetStudentParentByID(ctx context.Context, userID uint) (*model.ParentResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != "student" {
		s.logger.Error(errors.New("user is not a student"))
		return nil, errors.New("user is not a student")
	}

	parent, err := s.repository.User.GetStudentParentByID(ctx, id.ID, userID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student parent error: %v", err))
		return nil, fmt.Errorf("getting student parent error: %v", err)
	}

	return parent, nil
}

func (s *UserService) GetParentChildrenByID(ctx context.Context, userID uint) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != "parent" {
		s.logger.Error(errors.New("user is not a parent"))
		return nil, errors.New("user is not a parent")
	}

	children, err := s.repository.User.GetParentChildrenByID(ctx, id.ID, userID)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting parent children error: %v", err))
		return nil, fmt.Errorf("getting parent children error: %v", err)
	}

	return children, nil
}
