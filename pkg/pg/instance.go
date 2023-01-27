package pg

import (
	"context"
	"github.com/go-pg/pg/extra/pgsegment/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/pgjson"
	"go.uber.org/zap"
	"konntent-workspace-service/configs/app"
	"konntent-workspace-service/pkg/pg/hooks"
)

type Instance interface {
	Open() *pg.DB
}

type instance struct {
	db *pg.DB
}

func init() {
	pgjson.SetProvider(pgsegment.NewJSONProvider())
}

func NewPGInstance(l *zap.Logger, conf app.PGSettings) (Instance, error) {
	var i = &instance{
		db: pg.Connect(&pg.Options{
			User:     conf.User,
			Password: conf.Password,
			Database: "konntent-auth",
		}),
	}

	err := i.db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	if conf.Debug {
		i.db.AddQueryHook(hooks.NewDebugHook(l))
	}

	return i, nil
}

func (i *instance) Open() *pg.DB {
	return i.db
}
