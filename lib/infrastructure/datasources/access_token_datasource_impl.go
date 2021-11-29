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

type AccessTokenDatasourceImpl struct {
	Config        cfg.Config
	Database      sql.DB
	DBTransaction *sql.Tx
}

func (i AccessTokenDatasourceImpl) Create(ctx ctx.Context, UserId int) (*mdl.AccessTokenModel, error) {
	/// Initiate transaction
	dbTx := i.DBTransaction
	config := i.Config
	jwtHandler := new(sec.JwtHandler)
	expiredTime := time.Now().Add(time.Hour * 24)

	/// Functional code
	SecretKey := jwtHandler.Generate(config.Secret, expiredTime.String())
	var localId int
	const createClient = `INSERT INTO authentication.access_token (created_at, user_id, client_id, token, expired_at, device_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	stmt, err := dbTx.PrepareContext(ctx, createClient)
	if err != nil {
		return nil, err
	}

	stmtContext, err := dbTx.StmtContext(ctx, stmt).QueryContext(ctx, time.Now(), UserId, 1, SecretKey, expiredTime, "testDeviceId")
	if err != nil {
		return nil, err
	}

	insertStatus := stmtContext.Next()
	if !insertStatus {
		return nil, errors.New("Failed to create access token")
	}

	err = stmtContext.Scan(&localId)
	if err != nil {
		return nil, err
	}

	err = stmtContext.Close()
	if err != nil {
		return nil, err
	}

	result := mdl.AccessTokenModel{
		Id:    localId,
		Token: SecretKey,
	}

	return &result, nil
}

func (i AccessTokenDatasourceImpl) Check(ctx ctx.Context, UserId string, Token string) (*int, error) {
	/// Initialize transaction
	db := i.Database
	var accessTokenId int
	var expiredAt time.Time
	const checkRefreshToken = `SELECT id, expired_at FROM authentication.access_token WHERE token = $1 AND user_id = $2`

	stmt, err := db.PrepareContext(ctx, checkRefreshToken)
	if err != nil {
		return nil, err
	}

	sqlRow := stmt.QueryRowContext(ctx, Token)
	if sqlRow == nil {
		return nil, err
	}

	err = sqlRow.Scan(&accessTokenId, &expiredAt)
	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	/// Logical block
	today := time.Now()
	isTokenGranted := today.Before(expiredAt)
	if !isTokenGranted {
		return nil, fmt.Errorf("Access token is expired")
	}
	return &accessTokenId, nil
}
