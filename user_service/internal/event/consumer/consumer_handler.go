package consumer

import (
	"user_service/internal/domain/event"
	"user_service/pkg/kafka"
	"user_service/pkg/logs"
)

type EventHandlerFunc func(data []byte) error

type userEventHandler struct {
	gameEvent     event.GameEventUsecase
	topicEventMap map[string]EventHandlerFunc
}

func NewUserEventHandler(gameEvent event.GameEventUsecase) kafka.EventHandler {
	handler := &userEventHandler{
		gameEvent: gameEvent,
	}
	handler.registerHandlers()
	return handler
}

func (obj *userEventHandler) registerHandlers() {
	obj.topicEventMap = map[string]EventHandlerFunc{
		event.TopicUserReaded: obj.handleReadedEvent,
	}
}

func (obj userEventHandler) Handle(topic string, eventBytes []byte) {
	if handler, ok := obj.topicEventMap[topic]; ok {
		err := handler(eventBytes)
		if err != nil {
			logs.Errorf("Handler for topic %s failed: %v", topic, err)
		}
	} else {
		logs.Infof("No handler registered for topic: %s", topic)
	}
}
