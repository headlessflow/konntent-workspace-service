package request

import (
	"konntent-workspace-service/pkg/event"
	"konntent-workspace-service/pkg/event/schema"
)

type DummyRequest struct {
	Data interface{} `json:"data" validate:"required"`
}

func (dr *DummyRequest) ToEvent() *event.DummyEvent {
	return &event.DummyEvent{EventData: dr.Data, EventType: schema.DummyEventType}
}
