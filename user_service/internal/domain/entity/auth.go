package entity

import (
	"user_service/pkg/configs"

	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	// Create Access Token (JWT) when user succesfully login
	SignUsersAccessToken(req *UsersPassport) (string, error)
}

type AuthUsecase interface {
	Login(cfg *configs.Configs, req *UsersCredentials) (*UsersLoginRes, error)
}

type UsersCredentials struct {
	Username string `json:"username" db:"username" form:"username"`
	Password string `json:"password" db:"password" form:"password"`
}

type UsersPassport struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}

func (UsersPassport) TableName() string {
	return "users"
}
