package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/dto"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
)

type UserTokenService struct {
	repository        *storage.Repository
	jwtSecretKey      string
	passwordSecretKey string
	userHttpTransport *transport.UserHttpTransport
	userGrpcTransport *transport.UserGrpcTransport
	logger            *zap.SugaredLogger
}

func NewUserTokenService(dto *dto.UserTokenServiceDTO) *UserTokenService {
	return &UserTokenService{
		repository:        dto.Repository,
		jwtSecretKey:      dto.JwtSecretKey,
		passwordSecretKey: dto.PasswordSecretKey,
		userHttpTransport: dto.UserHttpTransport,
		userGrpcTransport: dto.UserGrpcTransport,
		logger:            dto.Logger,
	}
}

func (s *UserTokenService) Login(ctx context.Context, login model.Login) (*model.TokenResponse, error) {
	user, err := s.userGrpcTransport.GetUserByEmail(ctx, &proto.GetUserByEmailRequest{Email: login.Email})
	if err != nil {
		s.logger.Errorf("GetUser request err: %v", err)
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}
	err = utils.CheckPassword(login.Password, user.GetUser().GetPassword())
	if err != nil {
		s.logger.Error("incorrect password")
		return nil, errors.New("incorrect password")
	}

	userClaim := model.UserClaim{
		UserID: uint(user.GetUser().GetId()),
		Role:   user.GetUser().GetRole(),
	}

	tokens, err := s.generateToken(ctx, userClaim)
	if err != nil {
		s.logger.Errorf("generating token err: %v", err)
		return nil, fmt.Errorf("generating token err: %w", err)
	}

	res := &model.TokenResponse{
		UserID:       uint(user.GetUser().GetId()),
		Email:        user.GetUser().GetEmail(),
		Role:         user.GetUser().GetRole(),
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	return res, nil
}

func (s *UserTokenService) Register(ctx context.Context, user model.Register) (uint, error) {
	if user.Role == "admin" {
		return 0, errors.New("not permitted")
	}
	s.logger.Info(user)
	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		s.logger.Error(err)
		return 0, err
	}
	req := &proto.CreateUserRequest{
		User: &proto.CreateUser{
			FullName: user.FullName,
			Email:    user.Email,
			Password: pass,
			Role:     user.Role,
		},
	}
	res, err := s.userGrpcTransport.CreateUser(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return 0, err
	}
	return uint(res.GetId()), nil
}

func (s *UserTokenService) Verify(ctx context.Context, email string, code string) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserTokenService) RefreshToken(ctx context.Context, refreshToken string) (*model.JwtTokens, error) {
	token, err := jwt.Parse(
		refreshToken,
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
				return nil, errors.New("expiration date validation error")
			}
		}

		s.logger.Error(err)
		return nil, fmt.Errorf("failed to parse jwt err: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		s.logger.Error(err)
		return nil, fmt.Errorf("unexpected type %T", claims)
	}
	user, err := s.userGrpcTransport.GetUserByEmail(ctx, &proto.GetUserByEmailRequest{Email: claims["email"].(string)})
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}

	userClaim := model.UserClaim{
		UserID: uint(user.GetUser().GetId()),
		Role:   user.GetUser().GetRole(),
	}
	tokens, err := s.generateToken(ctx, userClaim)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("generating token err: %w", err)
	}
	return tokens, nil
}

func (s *UserTokenService) generateToken(ctx context.Context, user model.UserClaim) (*model.JwtTokens, error) {
	accessTokenExpirationTime := time.Now().Add(time.Hour)
	refreshTokenExpirationTime := time.Now().Add(24 * time.Hour)

	accessTokenClaims := &model.JWTClaim{
		UserID: user.UserID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime.Unix(),
		},
	}

	secretKey := []byte(s.jwtSecretKey)
	accessClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("AccessToken: SignedStrign err: %v", err)
		return nil, fmt.Errorf("AccessToken: SignedString err: %w", err)
	}

	refreshTokenClaims := model.RefreshJWTClaim{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpirationTime.Unix(),
		},
	}

	refreshClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("RefreshToken: SignedString err: %v", err)
		return nil, fmt.Errorf("RefreshToken: SignedString err: %w", err)
	}

	userToken := model.UserToken{
		UserID:       user.UserID,
		Role:         user.Role,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	s.logger.Info(userToken)

	err = s.repository.UserToken.CreateUserToken(ctx, userToken)
	if err != nil {
		s.logger.Errorf("CreateUserToken err: %v", err)
		return nil, errors.New(fmt.Sprintf("CreateUserToken err: %v", err))
	}

	jwtToken := &model.JwtTokens{
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	return jwtToken, nil
}
