package pg_migration

import (
	"konntent-workspace-service/internal/app/datamodel"
)

var MigrationModels = []interface{}{
	(*datamodel.Workspace)(nil),
	(*datamodel.UserWorkspace)(nil),
}
