package request

type GetWorkspaceRequest struct {
	WorkspaceID int `json:"workspaceId"`
	UserID      int `json:"userId"`
}
