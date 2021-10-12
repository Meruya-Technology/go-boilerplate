package database

import (
	"database/sql"
	"fmt"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Postgresql struct{}

func (p Postgresql) ConnectAndGet(config cfg.Config) (*sql.DB, error) {

	db, err := sql.Open(config.DbDriver, config.DbConfig)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Duration(5) * time.Second)
	return db, nil
}

func (p Postgresql) Close(db sql.DB) {
	fmt.Println("Connection closed")
	defer db.Close()
}
