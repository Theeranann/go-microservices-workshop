package entity

import "user_service/internal/domain/event"

type UsersRepository interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	FindOneUser(username string) (*UsersPassport, error)
	GetUserById(userID uint) (*Users, error)
	UpdateUser(user *Users) (*Users, error)
	DeleteUser(userID uint) error
	GetReadHistory(userID uint) ([]*event.ReadedEvent, error)
}

type UsersUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	GetUserByID(userID uint) (*UsersRes, error)
	UpdateUser(userId uint, req *UsersUpdateReq) (*UsersRes, error)
	DeleteUser(userID uint) error
	GetReadHistory(userID uint) ([]*event.ReadedEvent, error)
}

type Users struct {
	Id       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsersUpdateReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
