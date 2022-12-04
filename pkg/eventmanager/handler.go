package eventmanager

import (
	"context"
	"konntent-workspace-service/pkg/event"
)

type EventHandler interface {
	Handle(ctx context.Context) error
}

type IEventHandlerFactory interface {
	Make(e event.Factory) (EventHandler, error)
}
