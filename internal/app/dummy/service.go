package dummy

import (
	"context"
	"konntent-workspace-service/internal/app/dto/resource"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Handle(ctx context.Context) (resource.DummyResource, error)
}

type dummyService struct {
	logger *logrus.Logger
}

func NewDummyService(l *logrus.Logger) Service {
	return &dummyService{
		logger: l,
	}
}

func (s *dummyService) Handle(ctx context.Context) (resource.DummyResource, error) {
	return resource.DummyResource{}, nil
}
