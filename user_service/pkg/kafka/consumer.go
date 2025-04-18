package kafka

import (
	"user_service/pkg/logs"

	"github.com/IBM/sarama"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler}
}

func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	logs.Infof("Starting to consume topic: %s, partition: %d", claim.Topic(), claim.Partition())

	for msg := range claim.Messages() {
		logs.Debugf("[Consumer] message received | topic: %s | offset: %d", msg.Topic, msg.Offset)

		// ส่งต่อไปให้ handler ภายนอกจัดการ (decouple)
		obj.eventHandler.Handle(msg.Topic, msg.Value)

		session.MarkMessage(msg, "")
	}

	return nil
}
