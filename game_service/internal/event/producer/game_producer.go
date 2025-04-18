package producer

import (
	"game_service/internal/domain/event"
	"game_service/pkg/kafka"
	"game_service/pkg/logs"
)

type gameProducer struct {
	eventProducer kafka.EventProducer
}

func NewGameProducer(eventProducer kafka.EventProducer) event.GamesProducer {
	return gameProducer{eventProducer: eventProducer}
}

func (obj gameProducer) UserReaded(command event.ReadedEvent) error {
	err := obj.eventProducer.Produce(command, event.TopicUserReaded)
	if err != nil {
		logs.Errorf("Error producing user.read event: %v", err)
		return err
	}

	return nil
}
