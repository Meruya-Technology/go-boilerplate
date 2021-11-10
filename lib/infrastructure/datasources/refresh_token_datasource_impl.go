package datasources

import (
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
)

type RefreshTokenDatasourceImpl struct {
	Config        cfg.Config
	DBTransaction *sql.Tx
}
