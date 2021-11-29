package datasources

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	ctx "context"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	"github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type RefreshTokenDatasourceImpl struct {
	Config        cfg.Config
	Database      sql.DB
	DBTransaction *sql.Tx
}

func (i RefreshTokenDatasourceImpl) Create(ctx ctx.Context, AccessTokenId int) (*models.RefreshTokenModel, error) {
	/// Initiate transaction
	dbTx := i.DBTransaction
	config := i.Config
	jwtHandler := new(sec.JwtHandler)

	/// Functional code
	SecretKey := jwtHandler.Generate(config.Secret)
	var localId int
	const createClient = `INSERT INTO authentication.refresh_token (created_at, access_token_id, token, expired_at) VALUES ($1, $2, $3, $4) RETURNING id`

	stmt, err := dbTx.PrepareContext(ctx, createClient)
	if err != nil {
		return nil, err
	}

	expiredTime := time.Now().Add(time.Hour * 168)
	stmtContext, err := dbTx.StmtContext(ctx, stmt).QueryContext(ctx, time.Now(), AccessTokenId, SecretKey, expiredTime)
	if err != nil {
		return nil, err
	}

	insertStatus := stmtContext.Next()
	if !insertStatus {
		return nil, errors.New("Failed to create refresh token")
	}

	err = stmtContext.Scan(&localId)
	if err != nil {
		return nil, err
	}

	result := models.RefreshTokenModel{
		Id:    localId,
		Token: SecretKey,
	}
	return &result, nil
}

func (i RefreshTokenDatasourceImpl) Check(ctx ctx.Context, Token string) (*int, error) {
	/// Initialize transaction
	db := i.Database
	var accessTokenId int
	var expiredAt time.Time
	const checkRefreshToken = `SELECT access_token_id, expired_at FROM authentication.refresh_token WHERE Token = $1`

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
		return nil, fmt.Errorf("Refresh token is expired")
	}
	return &accessTokenId, nil
}

func (i RefreshTokenDatasourceImpl) Revoke(ctx ctx.Context, Token string) (bool, error) {
	return true, nil
}
