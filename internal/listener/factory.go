package listener

import (
	"konntent-workspace-service/internal/listener/consumer"
	"konntent-workspace-service/internal/listener/handler"
	"konntent-workspace-service/pkg/event"
	"konntent-workspace-service/pkg/event/schema"
	"konntent-workspace-service/pkg/eventmanager"
)

type eventHandlerFactory struct {
	service consumer.Service
}

func NewEventHandlerFactory(cs consumer.Service) eventmanager.IEventHandlerFactory {
	return &eventHandlerFactory{
		service: cs,
	}
}

func (eh *eventHandlerFactory) Make(e event.Factory) (eventmanager.EventHandler, error) {
	if e.Type() == schema.DummyEventType {
		return handler.NewDummyConsumerHandler(eh.service, e.Data().(*event.DummyEvent)), nil
	}

	return nil, schema.UnexpectedHandlerType
}
