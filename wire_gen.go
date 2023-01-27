// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package konntent_workspace_service

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/handler"
	"konntent-workspace-service/internal/app/orchestration"
	"konntent-workspace-service/internal/app/workspace"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/pg"
)

// Injectors from wire.go:

func InitAll(l *zap.Logger, pgInstance pg.Instance, nrInstance nrclient.NewRelicInstance) app.Router {
	repository := workspace.NewWorkspaceRepository(pgInstance)
	service := workspace.NewWorkspaceService(repository)
	workspaceOrchestrator := orchestration.NewWorkspaceOrchestrator(l, service)
	workspaceHandler := handler.NewWorkspaceHandler(workspaceOrchestrator)
	router := app.NewRoute(workspaceHandler)
	return router
}

// wire.go:

var serviceProviders = wire.NewSet(workspace.NewWorkspaceRepository)

var orchestratorProviders = wire.NewSet(orchestration.NewWorkspaceOrchestrator)

var handlerProviders = wire.NewSet(handler.NewWorkspaceHandler)

var repositoryProviders = wire.NewSet(workspace.NewWorkspaceService)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
	repositoryProviders,
)
