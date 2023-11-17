package model

import "github.com/golang-jwt/jwt"

type UserToken struct {
	Id           int    `db:"id"`
	UserId       uint   `db:"user_id"`
	Email        string `db:"email"`
	Role         string `db:"role"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
}

type TokenResponse struct {
	UserId       uint   `json:"user_id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaim struct {
	UserId         uint   `json:"user_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	StandardClaims jwt.StandardClaims
}
type RefreshJWTClaim struct {
	Email          string `json:"email"`
	StandardClaims jwt.StandardClaims
}

type UserClaim struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func (r RefreshJWTClaim) Valid() error {
	return nil
}

func (J JWTClaim) Valid() error {
	return nil
}
