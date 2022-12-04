package handler

import (
	"context"
	"konntent-workspace-service/internal/listener/consumer"
	"konntent-workspace-service/pkg/dummyclient/model"
	"konntent-workspace-service/pkg/event"
	"konntent-workspace-service/pkg/eventmanager"
	"konntent-workspace-service/pkg/utils"
)

type dummyConsumerHandler struct {
	service consumer.Service
	event   *event.DummyEvent
}

func NewDummyConsumerHandler(service consumer.Service, event *event.DummyEvent) eventmanager.EventHandler {
	return &dummyConsumerHandler{
		service: service,
		event:   event,
	}
}

func (o *dummyConsumerHandler) Handle(ctx context.Context) error {
	ctx = o.prepareHeaderMap(ctx)
	return o.service.DummyEvent(ctx, o.preparePayload())
}

func (o *dummyConsumerHandler) preparePayload() model.DummyRequest {
	return model.DummyRequest{
		Data: o.event.EventData,
	}
}

func (o *dummyConsumerHandler) prepareHeaderMap(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.HeaderMapCtx, o.event.EventHeaders)
}
