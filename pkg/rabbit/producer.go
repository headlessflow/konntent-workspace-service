package rabbit

import (
	"context"
	"konntent-workspace-service/pkg/constants"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Client interface {
	ConnectToBroker(connectionString string) error
	PublishOnQueue(msg []byte, eventType string, ttl int64) error
	Close()

	Sync
}

type Sync interface {
	Consume(ctx context.Context, handler ConsumerGroupHandler) error
}

type MessagingClient struct {
	logger      *logrus.Logger
	client      ClientConnector
	manager     ClientManager
	preProducer PreProducer
}

func NewMessagingClient(l *logrus.Logger, client ClientConnector, preProducer PreProducer) Client {
	return &MessagingClient{
		logger:      l,
		client:      client,
		preProducer: preProducer,
	}
}

func (m *MessagingClient) ConnectToBroker(connectionString string) error {
	var err error

	m.manager, err = m.client.Connect(connectionString)
	if err != nil {
		return err
	}

	err = m.SetupQueues()
	if err != nil {
		return err
	}

	return nil
}

// PublishOnQueue Publishes a message onto the queue.
func (m *MessagingClient) PublishOnQueue(body []byte, eventType string, ttl int64) error {
	closer, err := m.manager.Publish(PublishValues{
		ExchangeName: m.preProducer.Config().ExchangeName,
		RoutingKey:   m.preProducer.Config().RoutingKey + ".#",
		Mandatory:    true,
		Immediate:    false,
		Publishing: &amqp.Publishing{
			ContentType: "application/json",
			Type:        eventType,
			Timestamp:   time.Now(),
			Body:        body,
			Headers: map[string]interface{}{
				"x-delay": ttl,
			},
		},
	})
	defer closer.Close()

	return err
}

// Consume Consumes messages from the queue.
func (m *MessagingClient) Consume(ctx context.Context, handler ConsumerGroupHandler) error {
	deliveries, _, err := m.manager.Consume(ConsumeValues{
		QueueName:    m.preProducer.Config().QueueName,
		ConsumerName: constants.AppConsumerName,
	})

	for {
		err = handler.ConsumeClaim(ctx, deliveries)
		if err != nil {
			log.Println(err)
		}
	}
}

func (m *MessagingClient) SetupQueues() error {
	return m.preProducer.SetupQueues(m.manager.Processor())
}

func (m *MessagingClient) Close() {
	err := m.client.Close()
	if err != nil {
		m.logger.Error(err)
	}
}
