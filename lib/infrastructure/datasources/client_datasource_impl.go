package datasources

import (
	ctx "context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type ClientDatasourceImpl struct {
	Config        cfg.Config
	Database      sql.DB
	DBTransaction *sql.Tx
}

func (i ClientDatasourceImpl) Create(ctx ctx.Context, Name string, Secret string) (*mdl.ClientModel, error) {
	/// Initiate transaction
	dbTx := i.DBTransaction
	jwtHandler := new(sec.JwtHandler)

	/// Functional code
	SecretKey := jwtHandler.Generate(Secret)
	var Id int
	const createClient = `INSERT INTO authentication.client (created_at, name, secret) VALUES ($1, $2, $3) RETURNING id`

	stmt, err := dbTx.PrepareContext(ctx, createClient)
	if err != nil {
		return nil, err
	}

	fmt.Println(Name, SecretKey)
	stmtContext, err := dbTx.StmtContext(ctx, stmt).QueryContext(ctx, time.Now(), Name, SecretKey)
	if err != nil {
		return nil, err
	}

	inserStatus := stmtContext.Next()
	if !inserStatus {
		return nil, errors.New("Failed to insert client")
	}

	err = stmtContext.Scan(&Id)
	if err != nil {
		return nil, err
	}

	result := mdl.ClientModel{
		Id:     int(Id),
		Name:   Name,
		Secret: SecretKey,
	}
	return &result, nil
}

func (i ClientDatasourceImpl) Check(ctx ctx.Context, Secret string) (*mdl.ClientModel, error) {
	/// Initialize transaction
	db := i.Database
	var localId int
	const checkClient = `SELECT ID FROM authentication.client WHERE SECRET = $1 AND IS_ACTIVE = true`

	stmt, err := db.PrepareContext(ctx, checkClient)
	if err != nil {
		return nil, err
	}

	sqlRow := stmt.QueryRowContext(ctx, Secret)
	if sqlRow == nil {
		return nil, err
	}

	err = sqlRow.Scan(&localId)
	if err != nil {
		return nil, fmt.Errorf("Invalid client")
	}

	result := mdl.ClientModel{
		Id: int(localId),
	}
	return &result, nil
}
