package datamodel

import "github.com/go-pg/pg/v10/orm"

type Workspace struct {
	ID          int
	Name        string
	URL         string
	BM          string
	CompanyUnit int
	UserID      int
}

func (w *Workspace) String() string {
	return "Workspace"
}

func (w *Workspace) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}
