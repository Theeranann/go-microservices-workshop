package consumer

import (
	"encoding/json"
	"user_service/internal/domain/event"
	"user_service/pkg/logs"
)

func (obj *userEventHandler) handleReadedEvent(data []byte) error {
	var e event.ReadedEvent
	if err := json.Unmarshal(data, &e); err != nil {
		logs.Errorf("[Consumer/UserRead] Error decoding %v: %v | raw: %s", event.TopicUserReaded, err, string(data))
		return err
	}

	logs.Infof("[Consumer/UserRead] Received %v event for UserID: %v", event.TopicUserReaded, e.UserID)
	if err := obj.gameEvent.HandleUserRead(e); err != nil {
		logs.Errorf("[Consumer/UserRead] ReadedEvent failed: %v", err)
	}
	return nil
}
