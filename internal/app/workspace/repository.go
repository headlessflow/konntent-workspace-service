package workspace

import (
    "context"
    "konntent-workspace-service/internal/app/datamodel"
    "konntent-workspace-service/pkg/pg"
    pg_conditions "konntent-workspace-service/pkg/pg-conditions"
)

type Repository interface {
    GetUserWorkspaces(c context.Context, m datamodel.UserWorkspace) ([]datamodel.Workspace, error)
    GetUserWorkspace(c context.Context, m datamodel.UserWorkspace) (*datamodel.Workspace, error)
    AddWorkspace(c context.Context, m *datamodel.Workspace) (*datamodel.Workspace, error)
}

type repository struct {
    pg pg.Instance
}

func NewWorkspaceRepository(pg pg.Instance) Repository {
    return &repository{pg: pg}
}

func (r *repository) GetUserWorkspaces(c context.Context, m datamodel.UserWorkspace) ([]datamodel.Workspace, error) {
    var workspaces []datamodel.Workspace
    err := r.pg.Open().
        ModelContext(c, (*datamodel.Workspace)(nil)).
        Where(pg_conditions.WhereUserID, m.UserID).
        Select(&workspaces)

    if err != nil {
        return nil, err
    }

    return workspaces, nil
}

func (r *repository) GetUserWorkspace(c context.Context, m datamodel.UserWorkspace) (*datamodel.Workspace, error) {
    var workspace datamodel.Workspace
    err := r.pg.Open().
        ModelContext(c, &workspace).
        Where(pg_conditions.WhereID, m.WorkspaceID).
        Where(pg_conditions.WhereUserID, m.UserID).
        First()

    if err != nil {
        return nil, err
    }

    return &workspace, nil
}

func (r *repository) AddWorkspace(c context.Context, m *datamodel.Workspace) (*datamodel.Workspace, error) {
    _, err := r.pg.Open().
        ModelContext(c, m).
        Insert()

    if err != nil {
        return nil, err
    }

    return m, nil
}
