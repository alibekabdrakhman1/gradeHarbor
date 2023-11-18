package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) (uint, error)
	Delete(ctx context.Context, userId uint) error
	Update(ctx context.Context, user model.User, userId uint) (model.User, error)
	GetById(ctx context.Context, userId uint) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}