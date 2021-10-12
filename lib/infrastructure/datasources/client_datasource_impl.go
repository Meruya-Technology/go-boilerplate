package datasources

import (
	"fmt"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	db "github.com/Meruya-Technology/go-boilerplate/lib/common/database"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	m "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	"go.uber.org/zap"
)

type ClientDatasourceImpl struct {
	Config cfg.Config
}

func (i ClientDatasourceImpl) Create(name string) (*m.CLientModel, error) {
	logger, _ := zap.NewProduction()

	/// Open connection
	postgresql := new(db.Postgresql)
	session, err := postgresql.ConnectAndGet(i.Config)
	jwtHandler := new(sec.JwtHandler)

	/// Error handler
	if err != nil {
		fmt.Print(err.Error())
		logger.Fatal(err.Error())
		return nil, err
	}

	/// Functional code
	Secret := jwtHandler.Generate(i.Config.Secret)
	var result m.CLientModel
	const createClient = `-- name: CreateClient :exec INSERT INTO client (created_at, name, secret) VALUES ($1, $2, $3)`
	session.QueryRow(createClient, time.Now(), name, Secret).Scan(result)
	logger.Info("Field", zap.Any("ID", result), zap.String("Secret", Secret))

	/// Close connection
	postgresql.Close(*session)

	/// Compile result
	return &result, nil
}

func (impl ClientDatasourceImpl) Check(secret string, clientId int) (bool, error) {
	///
	return false, nil
}

func (impl ClientDatasourceImpl) CheckLocalSecret(secret string) bool {
	return false
}
