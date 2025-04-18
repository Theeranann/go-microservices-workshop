package producer

import (
	"errors"
	"user_service/internal/domain/event"
	"user_service/pkg/kafka"
	"user_service/pkg/logs"
)

type userProducer struct {
	eventProducer kafka.EventProducer
}

func NewUserProducer(eventProducer kafka.EventProducer) event.UserProducer {
	return userProducer{eventProducer: eventProducer}
}

func (obj userProducer) UserCreate(command event.UserEventCommand) error {
	if command.UserID == 0 {
		logs.Infof("UserCreateEvent bad request: UserID is missing")
		return errors.New("UserCreateEvent bad request")
	}

	// Create CreateUserEvent from the command
	payload := event.CreateUserEvent{
		UserID: uint64(command.UserID),
	}

	// Produce the CreateUserEvent
	err := obj.eventProducer.Produce(payload, event.TopicUserCreated)
	if err != nil {
		logs.Errorf("Error producing CreateUserEvent: %v", err)
		return err
	}
	return nil
}

func (obj userProducer) UserDelete(command event.UserEventCommand) error {
	if command.UserID == 0 {
		logs.Infof("UserDeleteEvent bad request: UserID is missing")
		return errors.New("UserDeleteEvent bad request")
	}

	// Create DeleteUserEvent from the command
	payload := event.DeleteUserEvent{
		UserID: uint64(command.UserID),
	}

	// Produce the DeleteUserEvent
	err := obj.eventProducer.Produce(payload, event.TopicUserDeleted)
	if err != nil {
		logs.Errorf("Error producing DeleteUserEvent: %v", err)
		return err
	}

	return nil
}
