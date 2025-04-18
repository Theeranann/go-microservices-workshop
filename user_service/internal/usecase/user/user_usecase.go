package user

import (
	"errors"
	"fmt"
	"user_service/internal/domain/entity"
	"user_service/internal/domain/event"
	"user_service/pkg/logs"

	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	UsersRepo    entity.UsersRepository
	userProducer event.UserProducer
}

func NewUsersUseCase(usersRepo entity.UsersRepository, producer event.UserProducer) entity.UsersUsecase {
	return &usersUsecase{
		UsersRepo:    usersRepo,
		userProducer: producer,
	}
}

func (u *usersUsecase) Register(req *entity.UsersRegisterReq) (*entity.UsersRegisterRes, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}

	// Produce the CreateUserEvent
	command := event.UserEventCommand{
		UserID: uint(user.Id),
	}

	err = u.userProducer.UserCreate(command)
	if err != nil {
		logs.Errorf("Error producing CreateUserEvent: %v", err)
		return nil, err
	}

	logs.Infof("[POST users/] %s event produced successfully", event.TopicUserCreated)
	return user, nil
}

func (u *usersUsecase) GetUserByID(userID uint) (*entity.UsersRes, error) {
	user, err := u.UsersRepo.GetUserById(userID)
	if err != nil {
		logs.Errorf("Error getting user by ID %d: %v", userID, err)
		return nil, err
	}

	logs.Infof("[GET users/:id] User retrieved successfully | ID: %d | Username: %s", user.Id, user.Username)
	resp := entity.UsersRes{
		Id:       user.Id,
		Username: user.Username,
	}
	return &resp, nil
}

func (u *usersUsecase) UpdateUser(userId uint, req *entity.UsersUpdateReq) (*entity.UsersRes, error) {
	// Check if the req parameter is nil
	if req == nil {
		logs.Info("Update request is nil")
		return nil, errors.New("update request is nil")
	}

	// Check if the user is in this service
	user, err := u.UsersRepo.GetUserById(userId)
	if err != nil {
		logs.Infof("Error getting user by ID: %v", err)
		return nil, err
	}

	// Update fields from the request if provided
	if req.Username != "" {
		user.Username = req.Username
		logs.Infof("[PUT users/:id] Updated username: %s", user.Username)
	}
	if req.Password != "" {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			logs.Infof("Error hashing password: %v", err)
			return nil, err
		}
		user.Password = string(hashedPassword)
		logs.Info("[PUT users/:id] Password updated")
	}

	// Save the updated user
	updatedUser, err := u.UsersRepo.UpdateUser(user)
	if err != nil {
		logs.Infof("Error updating user: %v", err)
		return nil, err
	}

	logs.Infof("[PUT users/:id] User updated successfully | ID: %d | Username: %s", updatedUser.Id, updatedUser.Username)
	resp := entity.UsersRes{
		Id:       updatedUser.Id,
		Username: updatedUser.Username,
	}

	return &resp, nil
}

func (u *usersUsecase) DeleteUser(userID uint) error {
	logs.Infof("[DELETE users/:id] Deleting user with ID: %d", userID)
	err := u.UsersRepo.DeleteUser(userID)
	if err != nil {
		logs.Errorf("Error deleting user with ID %d: %v", userID, err)
		return err
	}

	// Produce the DeleteUserEvent
	command := event.UserEventCommand{
		UserID: userID,
	}

	err = u.userProducer.UserDelete(command)
	if err != nil {
		logs.Errorf("Error producing DeleteUserEvent: %v", err)
		return err
	}
	logs.Infof("[DELETE users/:id] %s event produced successfully", event.TopicUserDeleted)
	logs.Infof("[DELETE users/:id] User with ID %d deleted successfully", userID)
	return nil
}

func (u *usersUsecase) GetReadHistory(userID uint) ([]*event.ReadedEvent, error) {
	logs.Infof("[GET users/:id/read] Retrieving readed events for user with ID: %d", userID)

	readeds, err := u.UsersRepo.GetReadHistory(userID)
	if err != nil {
		logs.Errorf("Error retrieving readed events for user with ID %d: %v", userID, err)
		return nil, err
	}

	logs.Infof("[GET users/:id/read] Readed events retrieved successfully for user with ID %d", userID)
	return readeds, nil
}
