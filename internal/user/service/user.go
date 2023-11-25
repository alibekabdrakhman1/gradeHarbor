package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
)

type UserService struct {
	repository *storage.Repository
	config     *config.Config
}

func NewUserService(r *storage.Repository, cfg *config.Config) *UserService {
	return &UserService{
		repository: r,
		config:     cfg,
	}
}

func (s *UserService) Create(ctx context.Context, user model.User) (uint, error) {
	user.IsConfirmed = false
	return s.repository.User.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, userID uint) (*model.ResponseUser, error) {
	return s.repository.User.GetById(ctx, userID)
}

func (s *UserService) GetByContext(ctx context.Context) (*model.ResponseUser, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		//s..Logger(ctx).Error("not valid context username")
		return nil, errors.New("not valid context userID")
	}

	user, err := s.GetByID(ctx, id.ID)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repository.User.GetByEmail(ctx, email)
}

func (s *UserService) Update(ctx context.Context, user model.User, userID uint) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Delete(ctx context.Context) error {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		//s..Logger(ctx).Error("not valid context username")
		return errors.New("not valid context userID")
	}

	err := s.repository.User.Delete(ctx, id.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteByID(ctx context.Context, userID uint) error {
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		//s..Logger(ctx).Error("not valid context username")
		return errors.New("not valid context userRole")
	}
	if role.Role != "admin" {
		return errors.New("not permitted")
	}

	err := s.repository.User.Delete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
