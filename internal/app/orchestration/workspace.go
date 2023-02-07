package orchestration

import (
	"context"
	"konntent-workspace-service/internal/app/dto/request"
	"konntent-workspace-service/internal/app/workspace"

	"go.uber.org/zap"
)

type WorkspaceOrchestrator interface {
	GetWorkspaces(c context.Context, uid int) (interface{}, error)
	GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (interface{}, error)
	AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error
}

type workspaceOrchestrator struct {
	workspaceService workspace.Service
	l                *zap.Logger
}

func NewWorkspaceOrchestrator(l *zap.Logger, ws workspace.Service) WorkspaceOrchestrator {
	return &workspaceOrchestrator{workspaceService: ws, l: l}
}

func (w *workspaceOrchestrator) GetWorkspaces(c context.Context, uid int) (interface{}, error) {
	return w.workspaceService.GetWorkspaces(c, uid)
}

func (w *workspaceOrchestrator) GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (interface{}, error) {
	return w.workspaceService.GetWorkspace(c, req)
}

func (w *workspaceOrchestrator) AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error {
	return w.workspaceService.AddWorkspace(c, req)
}
