package datamodel

import "github.com/go-pg/pg/v10/orm"

type UserWorkspace struct {
	UserID      int
	WorkspaceID int
}

func (uw *UserWorkspace) String() string {
	return "Workspace"
}

func (uw *UserWorkspace) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}
