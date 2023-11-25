package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type UserToken struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint      `gorm:"unique;not null" json:"full_name"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Role         string    `gorm:"not null" json:"role"`
	AccessToken  string    `gorm:"unique;not null" json:"access_token"`
	RefreshToken string    `gorm:"unique;not null" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type TokenResponse struct {
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaim struct {
	UserID         uint   `json:"user_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	StandardClaims jwt.StandardClaims
}
type RefreshJWTClaim struct {
	UserID         uint `json:"user_id"`
	StandardClaims jwt.StandardClaims
}

type UserClaim struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}

func (r RefreshJWTClaim) Valid() error {
	return nil
}

func (J JWTClaim) Valid() error {
	return nil
}
