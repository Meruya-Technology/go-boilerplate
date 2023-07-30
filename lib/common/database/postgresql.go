package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgresql interface {
	Initalize() *sql.DB
	Closes(db sql.DB) error
}
