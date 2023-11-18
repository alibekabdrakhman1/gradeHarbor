package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
	"github.com/golang-jwt/jwt"
	"time"
)

type UserTokenService struct {
	repository        *storage.Repository
	jwtSecretKey      string
	passwordSecretKey string
	userHttpTransport *transport.UserHttpTransport
	userGrpcTransport *transport.UserGrpcTransport
}

func NewUserTokenService(repo *storage.Repository, authConfig config.Auth, userHttpTransport *transport.UserHttpTransport, userGrpcTransport *transport.UserGrpcTransport) *UserTokenService {
	return &UserTokenService{
		repository:        repo,
		jwtSecretKey:      authConfig.JwtSecretKey,
		passwordSecretKey: authConfig.PasswordSecretKey,
		userHttpTransport: userHttpTransport,
		userGrpcTransport: userGrpcTransport,
	}
}

func (s *UserTokenService) Login(ctx context.Context, login model.Login) (*model.TokenResponse, error) {
	user, err := s.userHttpTransport.GetUser(ctx, login.Email)
	if err != nil {
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}
	fmt.Println(user.Password, login.Password)
	generatedPassword := s.generatePassword(login.Password)
	if user.Password != generatedPassword {
		return nil, fmt.Errorf("password is wrong")
	}

	userClaim := model.UserClaim{
		UserId: user.Id,
		Email:  user.Email,
		Role:   user.Role,
	}

	tokens, err := s.generateToken(ctx, userClaim)
	if err != nil {
		return nil, fmt.Errorf("generating token err: %w", err)
	}

	res := &model.TokenResponse{
		UserId:       user.Id,
		Email:        user.Email,
		Role:         user.Role,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return res, nil
}

func (s *UserTokenService) Register(ctx context.Context, user model.Register) (uint, error) {
	req := &proto.CreateUserRequest{
		User: &proto.User{
			FullName: user.FullName,
			Email:    user.Email,
			Password: s.generatePassword(user.Password),
			Role:     user.Role,
		},
	}
	res, err := s.userGrpcTransport.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	return uint(res.GetId()), nil
}

func (s *UserTokenService) Verify(ctx context.Context, email string, code string) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserTokenService) RenewToken(ctx context.Context, refreshToken string) (*model.JwtTokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserTokenService) generateToken(ctx context.Context, user model.UserClaim) (*model.JwtTokens, error) {
	accessTokenExpirationTime := time.Now().Add(time.Hour)
	refreshTokenExpirationTime := time.Now().Add(24 * time.Hour)

	accessTokenClaims := &model.JWTClaim{
		UserId: user.UserId,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime.Unix(),
		},
	}

	secretKey := []byte(s.jwtSecretKey)
	accessClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	refreshTokenClaims := model.RefreshJWTClaim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpirationTime.Unix(),
		},
	}

	refreshClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	userToken := model.UserToken{
		UserId:       user.UserId,
		Email:        user.Email,
		Role:         user.Role,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	err = s.repository.UserToken.CreateUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("CreateUserToken err: %w", err)
	}

	jwtToken := &model.JwtTokens{
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	return jwtToken, nil
}
func (s *UserTokenService) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}