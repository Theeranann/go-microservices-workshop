package gameevent

import (
	"user_service/internal/domain/event"
	"user_service/pkg/logs"
)

type gameEventUsecase struct {
	gameEventRepo event.GameEventRepository
}

func NewGameEventUse(gameEventrepo event.GameEventRepository) event.GameEventUsecase {
	return &gameEventUsecase{gameEventRepo: gameEventrepo}
}

func (u *gameEventUsecase) HandleUserRead(req event.ReadedEvent) error {
	err := u.gameEventRepo.SaveRead(req)
	if err != nil {
		logs.Errorf("[Usecase/UserRead] Error UserDeletedEvent: %v", err)
		return err
	} 

	logs.Infof("[Usecase/UserRead] User read created with | userID: %v | game: %v", req.UserID, req.Title)
	return nil
}
