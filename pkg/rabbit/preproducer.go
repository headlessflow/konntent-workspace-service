package rabbit

import (
	"konntent-workspace-service/configs/app"

	"github.com/sirupsen/logrus"
)

type PreProducer interface {
	SetupQueues(manager Processor) error
	Config() *app.RabbitQueueSettings
}

type PreProducerController struct {
	logger  *logrus.Logger
	conf    app.RabbitQueueSettings
	manager Processor
}

func NewPreProducer(l *logrus.Logger, conf app.RabbitQueueSettings) PreProducer {
	return &PreProducerController{
		logger: l,
		conf:   conf,
	}
}

func (pp *PreProducerController) Config() *app.RabbitQueueSettings {
	return &pp.conf
}

func (pp *PreProducerController) SetupQueues(manager Processor) error {
	pp.manager = manager

	if primaryErr := pp.SetupPrimaryQueue(); primaryErr != nil {
		return primaryErr
	}

	return pp.SetupDlqQueue()
}

// SetupPrimaryQueue Declares primary queue
func (pp *PreProducerController) SetupPrimaryQueue() error {
	err := pp.manager.QueueDeclare(QueueDeclarationValues{
		QueueName: pp.conf.QueueName,
		Durable:   true,
		Args: map[string]interface{}{
			"x-dead-letter-exchange":    pp.conf.Dlx,
			"x-dead-letter-routing-key": pp.conf.Dlrk + ".dlq1",
		},
	})
	if err != nil {
		pp.logger.Fatal("An error occurred while declaring primary queue. Error is ", err)
		return err
	}

	exErr := pp.manager.ExchangeDeclare(ExchangeDeclarationValues{
		ExchangeName: pp.conf.ExchangeName,
		ExchangeKind: "x-delayed-message",
		Durable:      true,
		Args: map[string]interface{}{
			"x-delayed-type": "topic",
		},
	})
	if exErr != nil {
		pp.logger.Fatal("An error occurred while declaring primary queue exchange. Error is ", exErr)
		return exErr
	}

	qbErr := pp.manager.QueueBind(QueueBindValues{
		QueueName:    pp.conf.QueueName,
		RoutingKey:   pp.conf.RoutingKey + ".#",
		ExchangeName: pp.conf.ExchangeName,
	})
	if qbErr != nil {
		pp.logger.Fatal("An error occurred while binding primary queue. Error is ", qbErr)
		return qbErr
	}

	return nil
}

// SetupDlqQueue Dead Letter Handling
func (pp *PreProducerController) SetupDlqQueue() error {
	dlqErr := pp.manager.QueueDeclare(QueueDeclarationValues{
		QueueName: pp.conf.Dlq,
		Durable:   true,
		Args: map[string]interface{}{
			"x-message-ttl":             pp.conf.DlqTTL,
			"x-dead-letter-exchange":    pp.conf.ExchangeName,
			"x-dead-letter-routing-key": pp.conf.RoutingKey + ".consume1",
		},
	})
	if dlqErr != nil {
		pp.logger.Fatal("An error occurred while declaring dlq queue. Error is ", dlqErr)
		return dlqErr
	}

	dlxErr := pp.manager.ExchangeDeclare(ExchangeDeclarationValues{
		ExchangeName: pp.conf.Dlx,
		ExchangeKind: ExchangeKindTopic,
		Durable:      true,
	})
	if dlxErr != nil {
		pp.logger.Fatal("An error occurred while declaring dlq queue exchange. Error is ", dlxErr)
		return dlxErr
	}

	dlbErr := pp.manager.QueueBind(QueueBindValues{
		QueueName:    pp.conf.Dlq,
		RoutingKey:   pp.conf.Dlrk + ".#",
		ExchangeName: pp.conf.Dlx,
	})
	if dlbErr != nil {
		pp.logger.Fatal("An error occurred while binding dlq queue. Error is", dlbErr)
		return dlbErr
	}

	return nil
}
