package rabbit

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ConsumerGroupHandler interface {
	Status() chan bool
	Close()

	ConsumeClaim(ctx context.Context, queue <-chan amqp.Delivery) error
}

type CustomHandler interface {
	Do(ctx context.Context, delivery amqp.Delivery) error
}

type ConsumerManager interface {
	Process(ctx context.Context, delivery amqp.Delivery) error
}

type consumerManager struct {
	logger        *logrus.Logger
	customHandler CustomHandler
}

func NewConsumerManager(l *logrus.Logger, ch CustomHandler) ConsumerManager {
	return &consumerManager{
		logger:        l,
		customHandler: ch,
	}
}

func (cm *consumerManager) Process(ctx context.Context, delivery amqp.Delivery) error {
	if err := cm.customHandler.Do(ctx, delivery); err != nil {
		cm.logger.WithField("event", delivery.Type).WithError(err).Error("processing error")
		return err
	}

	delivery.Ack(false)
	return nil
}
