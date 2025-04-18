package kafka

import (
	"encoding/json"
	"user_service/internal/domain/event"

	"github.com/IBM/sarama"
)

type EventProducer interface {
	Produce(event event.Event, topic string) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return eventProducer{producer}
}

func (obj eventProducer) Produce(event event.Event, topic string) error {
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = obj.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil
}
