package orchestration

import (
	"context"
	"konntent-workspace-service/internal/app/dto/request"
	"konntent-workspace-service/internal/app/dto/resource"
	"konntent-workspace-service/internal/app/dummy"
)

type DummyOrchestrator interface {
	Handle(ctx context.Context, req request.DummyRequest) (resource.DummyResource, error)
}

type dummyOrchestrator struct {
	dummyService dummy.Service
}

func NewDummyOrchestrator(ds dummy.Service) DummyOrchestrator {
	return &dummyOrchestrator{dummyService: ds}
}

func (o *dummyOrchestrator) Handle(ctx context.Context, req request.DummyRequest) (resource.DummyResource, error) {
	return o.dummyService.Handle(ctx)
}
