package workspace

import (
	"context"
	"konntent-workspace-service/internal/app/datamodel"
	"konntent-workspace-service/internal/app/dto/request"
	"konntent-workspace-service/internal/app/dto/resource"
)

type Service interface {
	GetWorkspaces(c context.Context, uid int) ([]resource.Workspace, error)
	GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (*resource.Workspace, error)
	AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error
}

type service struct {
	repo Repository
}

func NewWorkspaceService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetWorkspaces(c context.Context, uid int) ([]resource.Workspace, error) {
	var res, err = s.repo.GetUserWorkspaces(c, datamodel.UserWorkspace{UserID: uid})
	if err != nil {
		return nil, err
	}

	return resource.NewWorkspacesResource(res), nil
}

func (s *service) GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (*resource.Workspace, error) {
	var res, err = s.repo.GetUserWorkspace(c, req.ToDataModel())
	if err != nil {
		return nil, err
	}

	return resource.NewWorkspaceResource(res), nil
}

func (s *service) AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error {
	var _, err = s.repo.AddWorkspace(c, req.ToDataModel())
	if err != nil {
		return err
	}

	return nil
}
