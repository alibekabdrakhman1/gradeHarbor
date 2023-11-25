package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) (uint, error)
	GetByID(ctx context.Context, userID uint) (*model.ResponseUser, error)
	GetByContext(ctx context.Context) (*model.ResponseUser, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.User, userID uint) (*model.User, error)
	Delete(ctx context.Context) error
	DeleteByID(ctx context.Context, userID uint) error
}

type IAuthService interface {
	GetJwtUserID(jwtToken string) (*model.ContextUserID, error)
	GetJwtUserRole(jwtToken string) (*model.ContextUserRole, error)
}
