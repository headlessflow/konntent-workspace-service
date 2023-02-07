//go:build wireinject
// +build wireinject

package konntent_workspace_service

import (
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/handler"
	"konntent-workspace-service/internal/app/orchestration"
	"konntent-workspace-service/internal/app/workspace"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/pg"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var serviceProviders = wire.NewSet(
	workspace.NewWorkspaceRepository,
)

var orchestratorProviders = wire.NewSet(
	orchestration.NewWorkspaceOrchestrator,
)

var handlerProviders = wire.NewSet(
	handler.NewWorkspaceHandler,
)

var repositoryProviders = wire.NewSet(
	workspace.NewWorkspaceService,
)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
	repositoryProviders,
)

func InitAll(
	l *zap.Logger,
	pgInstance pg.Instance,
	nrInstance nrclient.NewRelicInstance,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
