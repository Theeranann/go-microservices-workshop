package consumer

import (
	"encoding/json"
	"game_service/internal/domain/event"
	"game_service/pkg/logs"
)

func (obj *userEventHandler) handleCreateUserEvent(data []byte) error {
	var e event.UserCreatedEvent
	if err := json.Unmarshal(data, &e); err != nil {
		logs.Errorf("[Consumer/UserCreated] Error decoding %v: %v | raw: %s", event.TopicUserCreated, err, string(data))
		return err
	}
	if e.UserID == 0 {
		logs.Info("Invalid UserID")
		return nil
	}

	logs.Infof("[Consumer/UserCreated] Received event for UserID: %v", e.UserID)
	if err := obj.userEvent.HandleUserCreated(e); err != nil {
		logs.Errorf("UserCreatedEvent failed: %v", err)
	}
	return nil
}

func (obj *userEventHandler) handleDeleteUserEvent(data []byte) error {
	var e event.UserDeletedEvent
	if err := json.Unmarshal(data, &e); err != nil {
		logs.Errorf("[Consumer/UserDeleted] Error decoding %v: %v | raw: %s", event.TopicUserDeleted, err, string(data))
		return err
	}
	if e.UserID == 0 {
		logs.Info("Invalid UserID")
		return nil
	}

	logs.Infof("[Consumer/UserDeleted] Received event for UserID: %v", e.UserID)
	if err := obj.userEvent.HandleUserDeleted(e); err != nil {
		logs.Errorf("UserDeletedEvent failed: %v", err)
	}
	return nil
}
