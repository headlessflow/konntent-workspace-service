package pg_rel_registration

import (
	"konntent-workspace-service/internal/app/datamodel"

	"github.com/go-pg/pg/v10/orm"
)

func Register() {
	orm.RegisterTable((*datamodel.UserWorkspace)(nil))
}
