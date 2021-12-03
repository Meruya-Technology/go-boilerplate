package repositories_impl

import (
	ctx "context"
	"database/sql"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	res "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/responses"
)

type AccessTokenRepositoryImpl struct {
	Config   cfg.Config
	Database sql.DB
}

func (i AccessTokenRepositoryImpl) Refresh(ctx ctx.Context, Request req.RefreshTokenRequest) (*res.LoginResponse, error) {
	/// Initiate transaction
	dbHandler := utl.DbHandler{DB: &i.Database}
	dbTx, err := dbHandler.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	/// Initialize datasource
	var accessTokenDatasource dts.AccessTokenDatasource = dts.AccessTokenDatasourceImpl{Config: i.Config, Database: i.Database, DBTransaction: dbTx}
	var refreshTokenDatasource dts.RefreshTokenDatasource = dts.RefreshTokenDatasourceImpl{Config: i.Config, Database: i.Database, DBTransaction: dbTx}

	/// Check refresh token expirity
	refreshToken, err := refreshTokenDatasource.Check(ctx, Request.RefreshToken)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Create access refresh token
	accessToken, err := accessTokenDatasource.GetById(ctx, refreshToken.AccessTokenId)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Create access token
	createAccessToken, err := accessTokenDatasource.Create(ctx, accessToken.UserId)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Create refresh token
	createRefreshToken, err := refreshTokenDatasource.Create(ctx, createAccessToken.Id)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Finishing block
	err = dbHandler.CommitTx(dbTx)
	if err != nil {
		return nil, err
	}

	/// Compile Result
	today := time.Now()
	expiredDate := createAccessToken.ExpiredAt
	expiresIn := int(today.Sub(expiredDate).Seconds())
	result := res.LoginResponse{
		AccessToken:  createAccessToken.Token,
		RefreshToken: createRefreshToken.Token,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
	}
	return &result, err
}
