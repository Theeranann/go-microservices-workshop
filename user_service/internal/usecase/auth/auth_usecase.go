package auth

import (
	"errors"
	"fmt"
	"user_service/internal/domain/entity"
	"user_service/pkg/configs"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	AuthRepo  entity.AuthRepository
	UsersRepo entity.UsersRepository
}

func NewAuthUsecase(authRepo entity.AuthRepository, usersRepo entity.UsersRepository) entity.AuthUsecase {
	return &authUsecase{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUsecase) Login(cfg *configs.Configs, req *entity.UsersCredentials) (*entity.UsersLoginRes, error) {
	user, err := u.UsersRepo.FindOneUser(req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error, password is invalid")
	}

	token, err := u.AuthRepo.SignUsersAccessToken(user)
	if err != nil {
		return nil, err
	}
	res := &entity.UsersLoginRes{
		AccessToken: token,
	}
	return res, nil
}
