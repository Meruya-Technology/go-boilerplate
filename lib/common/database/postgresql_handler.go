package database

import (
	"database/sql"
	"fmt"

	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	_ "github.com/lib/pq" // <------------ here
)

type PostgresqlHandler struct{}

func (p PostgresqlHandler) ConnectAndGet(cfg config.Config) (*sql.DB, error) {
	fmt.Println(`Connection Opened : ` + cfg.DbDriver)
	return sql.Open(cfg.DbDriver, cfg.DbName)
}

func (p PostgresqlHandler) Close(db sql.DB) {
	fmt.Println("Connection closed")
	defer db.Close()
}
