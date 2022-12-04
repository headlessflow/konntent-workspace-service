package schema

import "errors"

const (
	DummyEventType = "dummy"
)

const (
	MobilisimEnglishMessageDecoder = "default"
	MobilisimUnicodeMessageDecoder = "unicode"
	MobilisimTurkishMessageDecoder = "tr"
)

var (
	UnexpectedEventType   = errors.New("event: unexpected event type")
	UnexpectedHandlerType = errors.New("handler: unexpected event type")
)
