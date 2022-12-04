package event

import (
	"encoding/json"
	"konntent-workspace-service/pkg/event/schema"
)

type Factory interface {
	Type() string
	Data() interface{}
}

type Creation interface {
	Make(t string, data []byte) (Factory, error)
}

type eventCreator struct{}

var eventMap = map[string]Factory{
	schema.DummyEventType: &DummyEvent{},
}

func NewEventCreator() Creation {
	return &eventCreator{}
}

func (ec *eventCreator) Make(t string, data []byte) (Factory, error) {
	event := ec.getEventByType(t)
	if event == nil {
		return nil, schema.UnexpectedEventType
	}

	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}

	return event, nil
}

func (ec *eventCreator) getEventByType(t string) Factory {
	if event, ok := eventMap[t]; ok {
		return event
	}

	return nil
}
