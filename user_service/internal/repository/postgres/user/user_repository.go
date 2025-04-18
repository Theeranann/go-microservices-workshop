package user

import (
	"errors"
	"fmt"
	"strconv"
	"user_service/internal/domain/entity"
	"user_service/internal/domain/event"
	"user_service/pkg/logs"

	"gorm.io/gorm"
)

type usersRepositoryDB struct {
	Db *gorm.DB
}

func NewUsersRepositoryDB(db *gorm.DB) entity.UsersRepository {
	db.AutoMigrate(&entity.Users{})
	return &usersRepositoryDB{Db: db}
}

func (r *usersRepositoryDB) Register(req *entity.UsersRegisterReq) (*entity.UsersRegisterRes, error) {
	user := &entity.Users{
		Username: req.Username,
		Password: req.Password,
	}

	result := r.Db.Create(user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, result.Error
	}

	return &entity.UsersRegisterRes{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (r *usersRepositoryDB) FindOneUser(username string) (*entity.UsersPassport, error) {
	var user entity.UsersPassport

	result := r.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, errors.New("error, user not found")
	}

	return &user, nil
}

func (r *usersRepositoryDB) GetUserById(userID uint) (*entity.Users, error) {
	user := new(entity.Users)
	err := r.Db.First(user, userID).Error
	if err == gorm.ErrRecordNotFound {
		logs.Errorf("User with ID %d not found", userID)
		return nil, nil
	}

	return user, err
}

func (r *usersRepositoryDB) UpdateUser(user *entity.Users) (*entity.Users, error) {
	// Start a Gorm transaction
	tx := r.Db.Begin()
	if tx.Error != nil {
		logs.Errorf("Error starting transaction: %v", tx.Error)
		return nil, tx.Error
	}

	// Defer a function to handle the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			logs.Infof("Recovered from panic: %v", r)
		}
		tx.Rollback()
	}()

	// Attempt to update the user within the transaction
	err := tx.Save(&user).Error
	if err != nil {
		logs.Errorf("Error updating user: %v", err)
		return nil, err
	}

	// Commit the transaction if everything succeeded
	err = tx.Commit().Error
	if err != nil {
		logs.Errorf("Error committing transaction: %v", err)
		return nil, err
	}

	return user, nil
}

func (r *usersRepositoryDB) DeleteUser(userID uint) error {
	// Start a Gorm transaction
	tx := r.Db.Begin()
	if tx.Error != nil {
		logs.Errorf("Error starting transaction: %v", tx.Error)
		return tx.Error
	}
	// Defer a function to handle the rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			logs.Infof("Recovered from panic: %v", r)
		}
		tx.Rollback()
	}()

	// Attempt to delete the user within the transaction
	result := tx.Delete(&entity.Users{}, userID)
	if result.Error != nil {
		logs.Errorf("Error deleting user by ID %d: %v", userID, result.Error)
		return result.Error
	}

	// Commit the transaction
	err := tx.Commit().Error
	if err != nil {
		logs.Errorf("Error committing transaction: %v", err)
		return err
	}

	return nil
}

func (r *usersRepositoryDB) GetReadHistory(userID uint) ([]*event.ReadedEvent, error) {
	userIDString := strconv.FormatUint(uint64(userID), 10)

	var userRead []*event.ReadedEvent
	err := r.Db.Find(&userRead, userIDString).Error
	if err == gorm.ErrRecordNotFound {
		logs.Errorf("No readed history found for UserID: %d", userID)
		return nil, nil
	}

	return userRead, err
}