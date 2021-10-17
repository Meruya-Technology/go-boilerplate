package datasources

import (
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	dtb "github.com/Meruya-Technology/go-boilerplate/lib/common/database"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	ech "github.com/labstack/echo/v4"
	zap "go.uber.org/zap"
)

type ClientDatasourceImpl struct {
	Config cfg.Config
}

func (i ClientDatasourceImpl) Create(ctx ech.Context, Name string, Secret string) (*mdl.ClientModel, error) {
	logger, _ := zap.NewProduction()

	/// Open connection
	postgresql := new(dtb.Postgresql)
	session, err := postgresql.ConnectAndGet(i.Config)
	jwtHandler := new(sec.JwtHandler)

	/// Error handler
	if err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	/// Functional code
	SecretKey := jwtHandler.Generate(Secret)
	var Id int
	const createClient = `INSERT INTO authentication.client (created_at, name, secret) VALUES ($1, $2, $3) RETURNING id`
	session.Begin()
	session.QueryRow(createClient, time.Now(), Name, SecretKey).Scan(&Id)

	/// Close connection
	postgresql.Close(*session)

	/// Compile result
	result := mdl.ClientModel{
		Id:     Id,
		Name:   Name,
		Secret: SecretKey,
	}
	return &result, nil
}
