package service

import (
	"context"
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

func (s *UserService) Delete(ctx context.Context, userId uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Update(ctx context.Context, user model.User, userId uint) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) GetById(ctx context.Context, userId uint) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
func (s *UserService) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return s.repository.User.GetByEmail(ctx, email)
}
