package resource

import (
	"konntent-workspace-service/internal/app/datamodel"
)

type Workspace struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name" json:"name,omitempty"`
	URL  string `json:"url" json:"url,omitempty"`
}

func NewWorkspaceResource(w *datamodel.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
		URL:  w.URL,
	}
}

func NewWorkspacesResource(ws []datamodel.Workspace) []Workspace {
	var res = make([]Workspace, 0, len(ws))

	for _, w := range ws {
		res = append(res, Workspace{
			ID:   w.ID,
			Name: w.Name,
			URL:  w.URL,
		})
	}

	return res
}
