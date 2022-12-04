package consumer

import (
	"context"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/dummyclient/model"

	"github.com/sirupsen/logrus"
)

type Service interface {
	DummyEvent(ctx context.Context, req model.DummyRequest) error
}

type consumerService struct {
	logger      *logrus.Logger
	dummyClient dummyclient.Client
}

func NewDummyConsumerService(l *logrus.Logger, mc dummyclient.Client) Service {
	return &consumerService{
		logger:      l,
		dummyClient: mc,
	}
}

func (cs *consumerService) DummyEvent(ctx context.Context, req model.DummyRequest) error {
	cs.logger.WithField("req", req).Info("Incoming service request...")
	return cs.dummyClient.Dummy(ctx, req)
}
