package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
)

// SQLStore provides all functions to execute db queries and transactions.
type PostgresqlImpl struct {
	Config cfg.Config
}

func (s PostgresqlImpl) Initalize() *sql.DB {
	db, err := sql.Open(s.Config.DbDriver, s.Config.DbConfig)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Duration(5) * time.Second)
	return db
}

func (s PostgresqlImpl) Closes(db sql.DB) error {
	fmt.Println("Connection closed")
	return db.Close()
}
