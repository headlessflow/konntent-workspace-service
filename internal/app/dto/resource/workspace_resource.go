package resource

import (
    "konntent-workspace-service/internal/app/datamodel"
    "log"
)

type Workspace struct {
    Name string `json:"name"`
    URL string `json:"url"`
}

func NewWorkspaceResource(w *datamodel.Workspace) *Workspace {
    return &Workspace{}
}

func NewWorkspacesResource(ws []datamodel.Workspace) []Workspace {
    var res = make([]Workspace, 0, len(ws))

    for _, w := range ws {
        log.Print(w)
        res = append(res, Workspace{})
    }

    return res
}