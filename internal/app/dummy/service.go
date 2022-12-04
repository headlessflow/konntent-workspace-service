package dummy

import (
	"context"
	"konntent-workspace-service/internal/app/dto/resource"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/event"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Handle(ctx context.Context, dto *event.DummyEvent) (resource.DummyResource, error)
}

type dummyService struct {
	logger *logrus.Logger
	client dummyclient.Client
}

func NewDummyService(l *logrus.Logger, c dummyclient.Client) Service {
	return &dummyService{
		logger: l,
		client: c,
	}
}

func (s *dummyService) Handle(ctx context.Context, dto *event.DummyEvent) (resource.DummyResource, error) {
	return resource.DummyResource{}, nil
}
