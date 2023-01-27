package request

import "konntent-workspace-service/internal/app/datamodel"

type GetWorkspaceRequest struct {
	WorkspaceID int `json:"workspaceId"`
	UserID      int `json:"userId"`
}

type AddWorkspaceRequest struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	BM          string `json:"businessModel"`
	CompanyUnit int    `json:"companyUnit"`
	UserID      int    `json:"-"`
}

func (gwr GetWorkspaceRequest) ToDataModel() datamodel.UserWorkspace {
	return datamodel.UserWorkspace{
        WorkspaceID: gwr.WorkspaceID,
        UserID:      gwr.UserID,
	}
}

func (awr AddWorkspaceRequest) ToDataModel() *datamodel.Workspace {
    return &datamodel.Workspace{
        ID:          0,
        Name:        awr.Name,
        URL:         awr.URL,
        BM:          awr.BM,
        CompanyUnit: awr.CompanyUnit,
        UserID:      awr.UserID,
    }
}