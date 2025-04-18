package userevent

import (
	"game_service/internal/domain/event"
	"game_service/pkg/logs"
)

type userEventUsecase struct {
	userEventRepo event.UserEventRepository
}

func NewUserEventUse(userEventRepo event.UserEventRepository) event.UserEventUsecase {
	return &userEventUsecase{userEventRepo: userEventRepo}
}

func (u *userEventUsecase) HandleUserCreated(user event.UserCreatedEvent) error {
	logs.Infof("[Usecase/UserCreated] Handling creation for UserID: %d", user.UserID)
	err := u.userEventRepo.SaveUser(user)
	if err != nil {
		logs.Errorf("[Usecase/UserCreated] Error UserCreatedEvent: %v", err)
		return err
	}

	logs.Infof("[Usecase/UserCreated] User created with ID: %d", user.UserID)
	return nil
}

func (u *userEventUsecase) HandleUserDeleted(user event.UserDeletedEvent) error {
	logs.Infof("[Usecase/UserDeleted] Handling deleting for UserID: %d", user.UserID)
	err := u.userEventRepo.RemoveUser(user)
	if err != nil {
		logs.Errorf("[Usecase/UserDeleted] Error UserDeletedEvent: %v", err)
		return err
	}
	
	logs.Infof("[Usecase/UserDeleted] User deleted with ID: %d", user.UserID)
	return nil
}
