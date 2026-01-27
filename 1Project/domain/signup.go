package domain

import (
	"context"
)

type SignuRequest struct {
	Name     string `form:"name" binding:"required"`
	Email     string `form:"email" binding:"required, email"`
	Password     string `form:"name" binding:"required"`
}

type SignupResponse struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUseCase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}