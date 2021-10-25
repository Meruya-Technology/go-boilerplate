package datasources

import (
	ctx "context"
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
)

type AccessTokenDatasourceImpl struct {
	Config        cfg.Config
	DBTransaction *sql.Tx
}

func (a AccessTokenDatasourceImpl) Create(ctx ctx.Context) {

}
