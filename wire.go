//go:build wireinject
// +build wireinject

package konntent_service_template

import (
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/dummy"
	"konntent-workspace-service/internal/app/handler"
	"konntent-workspace-service/internal/app/orchestration"
	"konntent-workspace-service/internal/listener/consumer"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/rabbit"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var serviceProviders = wire.NewSet(
	dummy.NewDummyService,
	consumer.NewDummyConsumerService,
)

var orchestratorProviders = wire.NewSet(
	orchestration.NewDummyOrchestrator,
)

var handlerProviders = wire.NewSet(
	handler.NewDummyHandler,
)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
)

func InitAll(
	l *logrus.Logger,
	mc dummyclient.Client,
	mqp rabbit.Client,
	jwtInstance claimer.Claimer,
	nrInstance nrclient.NewRelicInstance,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
