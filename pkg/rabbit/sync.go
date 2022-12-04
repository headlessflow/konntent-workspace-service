package rabbit

import (
	"context"
	"konntent-workspace-service/pkg/rabbit/utils"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type syncHandler struct {
	logger          *logrus.Logger
	consumerManager ConsumerManager

	maxRetries int
	ready      chan bool
}

func NewSyncHandler(l *logrus.Logger, cm ConsumerManager, maxRetries int) ConsumerGroupHandler {
	return &syncHandler{
		logger:          l,
		consumerManager: cm,
		maxRetries:      maxRetries,
		ready:           make(chan bool),
	}
}

func (s *syncHandler) ConsumeClaim(ctx context.Context, queue <-chan amqp.Delivery) error {
	errCh := make(chan error)
	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Context done...")
			return nil
		case delivery := <-queue:
			deathCount := utils.GetXDeathCount(delivery.Headers)
			s.logger.Info("Got a delivery... ")

			if deathCount >= s.maxRetries-1 {
				delivery.Ack(false)
				s.logger.Info("death count has been exceed. Message will be acked...")
				return nil
			}

			if len(delivery.Body) == 0 {
				return nil
			}

			go func() {
				errCh <- s.consumerManager.Process(ctx, delivery)
			}()

			select {
			case err := <-errCh:
				if err != nil {
					delivery.Reject(false)
				}
			}
		}
	}
}

func (s *syncHandler) Close() {
	close(s.ready)
}

func (s *syncHandler) Status() chan bool {
	return s.ready
}
