package rabbit

import (
	"errors"

	"github.com/streadway/amqp"
)

type Closer interface {
	Close() error
}

type Processor interface {
	QueueDeclare(values QueueDeclarationValues) error
	ExchangeDeclare(values ExchangeDeclarationValues) error
	QueueBind(values QueueBindValues) error
}

type PublishValues struct {
	ExchangeName string
	RoutingKey   string
	Mandatory    bool
	Immediate    bool
	Publishing   *amqp.Publishing
}

type ConsumeValues struct {
	QueueName    string
	ConsumerName string
	AutoAck      bool
	Exclusive    bool
	NoLocal      bool
	NoWait       bool
	Args         amqp.Table
}

type clientManager struct {
	conn *amqp.Connection
}

/*
Channel opens a unique, concurrent server channel to process the bulk of AMQP
messages.  Any error from methods on this receiver will render the receiver
invalid and a new Channel should be opened.
*/
func (c *clientManager) Channel() (*amqp.Channel, Closer) {
	if c.conn == nil {
		return nil, nil
	}

	ch, _ := c.conn.Channel()
	return ch, ch
}

/*
Publish sends a Publishing from the client to an exchange on the server.

When you want a single message to be delivered to a single queue, you can
publish to the default exchange with the routingKey of the queue name.  This is
because every declared queue gets an implicit route to the default exchange.

Since publishings are asynchronous, any undeliverable message will get returned
by the server.  Add a listener with Channel.NotifyReturn to handle any
undeliverable message when calling publish with either the mandatory or
immediate parameters as true.

Publishings can be undeliverable when the mandatory flag is true and no queue is
bound that matches the routing key, or when the immediate flag is true and no
consumer on the matched queue is ready to accept the delivery.

This can return an error when the channel, connection or socket is closed.  The
error or lack of an error does not indicate whether the server has received this
publishing.

It is possible for publishing to not reach the broker if the underlying socket
is shut down without pending publishing packets being flushed from the kernel
buffers.  The easy way of making it probable that all publishings reach the
server is to always call Connection.Close before terminating your publishing
application.  The way to ensure that all publishings reach the server is to add
a listener to Channel.NotifyPublish and put the channel in confirm mode with
Channel.Confirm.  Publishing delivery tags and their corresponding
confirmations start at 1.  Exit when all publishings are confirmed.

When Publish does not return an error and the channel is in confirm mode, the
internal counter for DeliveryTags with the first confirmation starts at 1.
*/
func (c *clientManager) Publish(values PublishValues) (Closer, error) {
	if values.Publishing == nil {
		return nil, errors.New("you cannot publish empty message")
	}

	ch, _ := c.Channel()

	return ch, ch.Publish(
		values.ExchangeName,
		values.RoutingKey,
		values.Mandatory,
		values.Immediate,
		*values.Publishing)
}

/*
Consume immediately starts delivering queued messages.

Begin receiving on the returned chan Delivery before any other operation on the
Connection or Channel.

Continues deliveries to the returned chan Delivery until Channel.Cancel,
Connection.Close, Channel.Close, or an AMQP exception occurs.  Consumers must
range over the chan to ensure all deliveries are received.  Unreceived
deliveries will block all methods on the same connection.

All deliveries in AMQP must be acknowledged.  It is expected of the consumer to
call Delivery.Ack after it has successfully processed the delivery.  If the
consumer is cancelled or the channel or connection is closed any unacknowledged
deliveries will be requeued at the end of the same queue.

The consumer is identified by a string that is unique and scoped for all
consumers on this channel.  If you wish to eventually cancel the consumer, use
the same non-empty identifier in Channel.Cancel.  An empty string will cause
the library to generate a unique identity.  The consumer identity will be
included in every Delivery in the ConsumerTag field

When autoAck (also known as noAck) is true, the server will acknowledge
deliveries to this consumer prior to writing the delivery to the network.  When
autoAck is true, the consumer should not call Delivery.Ack. Automatically
acknowledging deliveries means that some deliveries may get lost if the
consumer is unable to process them after the server delivers them.
See http://www.rabbitmq.com/confirms.html for more details.

When exclusive is true, the server will ensure that this is the sole consumer
from this queue. When exclusive is false, the server will fairly distribute
deliveries across multiple consumers.

The noLocal flag is not supported by RabbitMQ.

It's advisable to use separate connections for
Channel.Publish and Channel.Consume so not to have TCP pushback on publishing
affect the ability to consume messages, so this parameter is here mostly for
completeness.

When noWait is true, do not wait for the server to confirm the request and
immediately begin deliveries.  If it is not possible to consume, a channel
exception will be raised and the channel will be closed.

Optional arguments can be provided that have specific semantics for the queue
or server.

Inflight messages, limited by Channel.Qos will be buffered until received from
the returned chan.

When the Channel or Connection is closed, all buffered and inflight messages will
be dropped.

When the consumer tag is cancelled, all inflight messages will be delivered until
the returned chan is closed.
*/
func (c *clientManager) Consume(values ConsumeValues) (<-chan amqp.Delivery, Closer, error) {
	ch, _ := c.Channel()

	deliveries, cErr := ch.Consume(
		values.QueueName,
		values.ConsumerName,
		values.AutoAck,
		values.Exclusive,
		values.NoLocal,
		values.NoWait,
		values.Args,
	)
	if cErr != nil {
		return nil, nil, cErr
	}

	return deliveries, ch, nil
}

/*
Processor returns a container which holds a Channel to maintain the organization of
queue and exchange declarations. Where all the operations done just by single-pass
Channel.

If you want to manage all your work by independently, you need to create N number of
Processor.
*/
func (c *clientManager) Processor() Processor {
	ch, _ := c.Channel()
	if ch == nil {
		return nil
	}
	return &processor{ch: ch}
}
