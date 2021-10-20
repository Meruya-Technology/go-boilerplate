package datasources

import (
	"database/sql"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	ech "github.com/labstack/echo/v4"
)

type ClientDatasourceImpl struct {
	Config   cfg.Config
	Database sql.DB
}

func (i ClientDatasourceImpl) Create(ctx ech.Context, Name string, Secret string) (*mdl.ClientModel, error) {

	/// Open connection
	session := i.Database
	jwtHandler := new(sec.JwtHandler)

	/// Functional code
	SecretKey := jwtHandler.Generate(Secret)
	var Id int
	const createClient = `INSERT INTO authentication.client (created_at, name, secret) VALUES ($1, $2, $3) RETURNING id`
	session.Begin()
	session.QueryRow(createClient, time.Now(), Name, SecretKey).Scan(&Id)

	/// Close connection
	session.Close()

	/// Compile result
	result := mdl.ClientModel{
		Id:     Id,
		Name:   Name,
		Secret: SecretKey,
	}
	return &result, nil
}
