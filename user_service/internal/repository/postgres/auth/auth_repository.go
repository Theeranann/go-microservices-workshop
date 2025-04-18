package auth

import (
	"fmt"
	"os"
	"time"
	"user_service/internal/domain/entity"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepositoryDB struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) entity.AuthRepository {
	return &authRepositoryDB{
		Db: db,
	}
}

func (r *authRepositoryDB) SignUsersAccessToken(req *entity.UsersPassport) (string, error) {
	claims := entity.UsersClaims{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return ss, nil
}
