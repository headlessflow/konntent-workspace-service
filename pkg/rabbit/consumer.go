package rabbit

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ConsumerInstance interface {
	Handler() ConsumerGroupHandler
	Consume(ctx context.Context)

	Close()
}

type consumerInstance struct {
	logger  *logrus.Logger
	client  Client
	handler ConsumerGroupHandler
}

func NewConsumerInstance(l *logrus.Logger, c Client, h ConsumerGroupHandler) ConsumerInstance {
	return &consumerInstance{
		logger:  l,
		client:  c,
		handler: h,
	}
}

func (k *consumerInstance) Handler() ConsumerGroupHandler {
	return k.handler
}

func (k *consumerInstance) Close() {
	k.client.Close()
}

func (k *consumerInstance) Consume(ctx context.Context) {
	for {
		if err := k.client.Consume(ctx, k.handler); err != nil {
			ctx = context.WithValue(ctx, "error", err)
			k.logger.Fatalf("consume: %v", err)
		}

		if ctx.Err() != nil {
			return
		}
	}
}
