package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
)

type IUserTokenService interface {
	Login(ctx context.Context, login model.Login) (*model.TokenResponse, error)
	Register(ctx context.Context, user model.Register) (uint, error)
	Confirm(ctx context.Context, email string, code string) error
	RefreshToken(ctx context.Context, refreshToken string) (*model.JwtTokens, error)
}
