package event

type DummyEvent struct {
	EventData    interface{}
	EventHeaders map[string]string
	EventType    string
}

func (de *DummyEvent) Type() string {
	return de.EventType
}

func (de *DummyEvent) Data() interface{} {
	return de.EventData
}
