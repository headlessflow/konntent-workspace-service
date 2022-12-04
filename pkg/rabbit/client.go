package rabbit

import (
	"crypto/tls"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ClientConnector interface {
	Connect(connectionString string) (ClientManager, error)
	Close() error
}

type ClientManager interface {
	Channel() (*amqp.Channel, Closer)
	Publish(values PublishValues) (Closer, error)
	Consume(values ConsumeValues) (<-chan amqp.Delivery, Closer, error)

	Processor() Processor
}

type client struct {
	l    *logrus.Logger
	conn *amqp.Connection
}

func NewClientConnector(l *logrus.Logger) ClientConnector {
	return &client{
		l: l,
	}
}

func (c *client) Connect(connectionString string) (ClientManager, error) {
	var err error

	c.conn, err = amqp.DialTLS(connectionString, &tls.Config{})
	if err != nil {
		c.l.Fatal("Failed to connect to AMQP compatible broker at: " + connectionString)
		return nil, err
	}

	return &clientManager{conn: c.conn}, nil
}

func (c *client) Close() error {
	if c.conn != nil && !c.conn.IsClosed() {
		return c.conn.Close()
	}
	return nil
}
