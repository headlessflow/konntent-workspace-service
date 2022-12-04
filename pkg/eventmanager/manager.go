package eventmanager

import (
	"context"
	"fmt"
	"konntent-workspace-service/pkg/event"
	"konntent-workspace-service/pkg/rabbit"

	"github.com/streadway/amqp"
)

type EventManager interface {
	Handle(ctx context.Context, delivery amqp.Delivery) (*rabbit.MessageAttribute, error)
	HandleException(eventType string, err error) error
}

type eventManager struct {
	handlerFactory IEventHandlerFactory
	eventCreation  event.Creation
}

func NewEventManager(hFactory IEventHandlerFactory, tCreator event.Creation) EventManager {
	return &eventManager{
		handlerFactory: hFactory,
		eventCreation:  tCreator,
	}
}

func (em *eventManager) Handle(ctx context.Context, delivery amqp.Delivery) (*rabbit.MessageAttribute, error) {
	attr := rabbit.GetEventAttributes(&delivery)

	e, err := em.eventCreation.Make(attr.Type, delivery.Body)
	if err != nil {
		return attr, em.HandleException(attr.Type, err)
	}

	eh, ehErr := em.handlerFactory.Make(e)
	if ehErr != nil {
		return attr, em.HandleException(attr.Type, ehErr)
	}

	if handleErr := eh.Handle(ctx); handleErr != nil {
		return attr, em.HandleException(attr.Type, handleErr)
	}

	return attr, nil
}

func (em *eventManager) HandleException(eventType string, err error) error {
	return fmt.Errorf("event type: %s - err: %v", eventType, err)
}
