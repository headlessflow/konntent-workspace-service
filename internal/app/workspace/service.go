package workspace

import (
	"context"
	"konntent-workspace-service/internal/app/dto/request"
)

type Service interface {
	GetWorkspaces(c context.Context, uid int) (interface{}, error)
	GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (interface{}, error)
	AddWorkspace(c context.Context, uid int) (interface{}, error)
}

type service struct {
	repo Repository
}

func NewWorkspaceService(r Repository) Service {
	return &service{repo: r}
}

func (s service) GetWorkspaces(c context.Context, uid int) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) AddWorkspace(c context.Context, uid int) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
