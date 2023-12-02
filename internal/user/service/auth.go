package service

import (
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/golang-jwt/jwt"
	"strconv"
)

var ErrExpiredToken = errors.New("expiration date validation error")

type AuthService struct {
	jwtSecretKey string
}

func NewAuthService(config config.Auth) *AuthService {
	return &AuthService{
		jwtSecretKey: config.JwtSecretKey,
	}
}

func (s *AuthService) GetJwtUserID(jwtToken string) (*model.ContextUserID, error) {
	token, err := jwt.Parse(
		jwtToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.jwtSecretKey), nil
		},
	)

	if err != nil {
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.Errors&jwt.ValidationErrorExpired > 0 {
				return nil, ErrExpiredToken
			}
		}

		return nil, fmt.Errorf("failed to parse jwt err: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T", claims)
	}

	user, err := s.getUserIDFromJwt(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from jwt err: %w", err)
	}
	return user, nil
}

func (s *AuthService) GetJwtUserRole(jwtToken string) (*model.ContextUserRole, error) {
	token, err := jwt.Parse(
		jwtToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.jwtSecretKey), nil
		},
	)

	if err != nil {
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.Errors&jwt.ValidationErrorExpired > 0 {
				return nil, ErrExpiredToken
			}
		}

		return nil, fmt.Errorf("failed to parse jwt err: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T", claims)
	}

	user, err := s.getUserRoleFromJwt(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from jwt err: %w", err)
	}
	return user, nil
}

func (s *AuthService) getUserIDFromJwt(claims jwt.MapClaims) (*model.ContextUserID, error) {
	user := &model.ContextUserID{}
	userId, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("user is not exists in jwt")
	}

	userId = fmt.Sprintf("%.0f", userId)

	parsedUserId, err := strconv.ParseUint(userId.(string), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("unexpected in userID value: %T", userId)
	}

	user.ID = uint(parsedUserId)

	return user, nil
}

func (s *AuthService) getUserRoleFromJwt(claims jwt.MapClaims) (*model.ContextUserRole, error) {
	user := &model.ContextUserRole{}
	role, ok := claims["role"]
	if !ok {
		return nil, fmt.Errorf("user is not exists in jwt")
	}

	user.Role = role.(string)

	return user, nil
}
