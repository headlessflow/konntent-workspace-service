package request

import (
	"konntent-workspace-service/internal/app/datamodel"
	"log"
)

type GetWorkspaceRequest struct {
	WorkspaceID int `json:"workspaceId"`
	UserID      int `json:"-" validate:"required"`
}

type AddWorkspaceRequest struct {
	Name        string `json:"name" validate:"required"`
	URL         string `json:"url" validate:"required"`
	BM          string `json:"businessModel" validate:"required"`
	CompanyUnit int    `json:"companyUnit" validate:"required"`
	UserID      int    `json:"-" validate:"required"`
}

func (gwr GetWorkspaceRequest) ToDataModel() datamodel.UserWorkspace {
	return datamodel.UserWorkspace{
		WorkspaceID: gwr.WorkspaceID,
		UserID:      gwr.UserID,
	}
}

func (awr AddWorkspaceRequest) ToDataModel() *datamodel.Workspace {
	log.Println("uıser id", awr.UserID)
	return &datamodel.Workspace{
		Name:        awr.Name,
		URL:         awr.URL,
		BM:          awr.BM,
		CompanyUnit: awr.CompanyUnit,
		UserID:      awr.UserID,
	}
}
