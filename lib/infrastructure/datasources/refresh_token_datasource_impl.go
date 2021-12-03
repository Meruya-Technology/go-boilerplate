package datasources

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	ctx "context"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type RefreshTokenDatasourceImpl struct {
	Config        cfg.Config
	Database      sql.DB
	DBTransaction *sql.Tx
}

func (i RefreshTokenDatasourceImpl) Create(ctx ctx.Context, AccessTokenId int) (*mdl.RefreshTokenModel, error) {
	/// Initiate transaction
	dbTx := i.DBTransaction
	config := i.Config
	jwtHandler := new(sec.JwtHandler)
	expiredTime := time.Now().Add(time.Hour * 168)

	/// Functional code
	SecretKey := jwtHandler.Generate(config.Secret, expiredTime.String())
	var localId int
	const createRefreshToken = `INSERT INTO authentication.refresh_token (created_at, access_token_id, token, expired_at) VALUES ($1, $2, $3, $4) RETURNING id`

	stmt, err := dbTx.PrepareContext(ctx, createRefreshToken)
	if err != nil {
		return nil, err
	}

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

	err = stmtContext.Close()
	if err != nil {
		return nil, err
	}

	result := mdl.RefreshTokenModel{
		Id:    localId,
		Token: SecretKey,
	}
	return &result, nil
}

func (i RefreshTokenDatasourceImpl) Check(ctx ctx.Context, Token string) (*mdl.RefreshTokenModel, error) {
	/// Initialize transaction
	db := i.Database
	const checkRefreshToken = `SELECT access_token_id, expired_at FROM authentication.refresh_token WHERE Token = $1`

	stmt, err := db.PrepareContext(ctx, checkRefreshToken)
	if err != nil {
		return nil, err
	}

	sqlRow := stmt.QueryRowContext(ctx, Token)
	if sqlRow == nil {
		return nil, err
	}

	result := mdl.RefreshTokenModel{}
	err = sqlRow.Scan(&result.Id, &result.CreatedAt, &result.AccessTokenId, &result.Token, &result.ExpiredAt, &result.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	/// Logical block
	today := time.Now()
	isTokenGranted := today.Before(result.ExpiredAt)
	if !isTokenGranted {
		return nil, fmt.Errorf("Refresh token is expired")
	}
	return &result, nil
}

func (i RefreshTokenDatasourceImpl) Revoke(ctx ctx.Context, Token string) (bool, error) {
	return true, nil
}
