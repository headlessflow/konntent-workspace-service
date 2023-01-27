//go:build wireinject
// +build wireinject

package konntent_service_template

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/dummy"
	"konntent-workspace-service/internal/app/handler"
	"konntent-workspace-service/internal/app/orchestration"
	"konntent-workspace-service/internal/listener/consumer"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/pg"
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
	l *zap.Logger,
	pgInstance pg.Instance,
	nrInstance nrclient.NewRelicInstance,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
