package consumer

import (
	"game_service/internal/domain/event"
	"game_service/pkg/kafka"
	"game_service/pkg/logs"
)

type EventHandlerFunc func(data []byte) error

type userEventHandler struct {
	userEvent       event.UserEventUsecase
	handlersByTopic map[string]EventHandlerFunc
}

func NewUserEventHandler(userEvent event.UserEventUsecase) kafka.EventHandler {
	handler := &userEventHandler{
		userEvent: userEvent,
	}
	handler.registerHandlers()
	return handler
}

func (obj *userEventHandler) registerHandlers() {
	obj.handlersByTopic = map[string]EventHandlerFunc{
		event.TopicUserCreated: obj.handleCreateUserEvent,
		event.TopicUserDeleted: obj.handleDeleteUserEvent,
	}
}

func (obj *userEventHandler) Handle(topic string, eventBytes []byte) {
	if handler, ok := obj.handlersByTopic[topic]; ok {
		err := handler(eventBytes)
		if err != nil {
			logs.Errorf("Handler for topic %s failed: %v", topic, err)
		}
	} else {
		logs.Infof("No handler registered for topic: %s", topic)
	}
}
