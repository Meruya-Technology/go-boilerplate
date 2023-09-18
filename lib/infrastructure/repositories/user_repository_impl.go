package repositories_impl

import (
	ctx "context"
	"database/sql"
	"time"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	res "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/responses"
)

type UserRepositoryImpl struct {
	Config   cfg.Config
	Database *sql.DB
}

func (c UserRepositoryImpl) Login(ctx ctx.Context, Request req.LoginRequest) (*res.LoginResponse, error) {
	/// Initiate transaction
	dbHandler := utl.DbHandler{DB: c.Database}
	hashHandler := sec.HashHandler{}
	dbTx, err := dbHandler.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	/// Initialize datasource
	var userDatasource dts.UserDatasource = dts.UserDatasourceImpl{Config: c.Config, DBTransaction: dbTx, HashHandler: hashHandler}
	var accessTokenDatasource dts.AccessTokenDatasource = dts.AccessTokenDatasourceImpl{Config: c.Config, Database: c.Database, DBTransaction: dbTx}
	var refreshTokenDatasource dts.RefreshTokenDatasource = dts.RefreshTokenDatasourceImpl{Config: c.Config, Database: c.Database, DBTransaction: dbTx}

	/// User login
	user, err := userDatasource.Login(ctx, Request.Email, Request.Password)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Create access token
	accessToken, err := accessTokenDatasource.Create(ctx, user.Id)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Create access token
	refreshToken, err := refreshTokenDatasource.Create(ctx, accessToken.Id)
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
	expiredDate := accessToken.ExpiredAt
	expiresIn := int(today.Sub(expiredDate).Seconds())
	result := res.LoginResponse{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
	}
	return &result, err
}

func (c UserRepositoryImpl) Register(ctx ctx.Context, Request req.RegisterRequest) (*bool, error) {
	/// Initiate transaction
	dbHandler := utl.DbHandler{DB: c.Database}
	hashHandler := sec.HashHandler{}
	dbTx, err := dbHandler.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	/// Initialize datasource
	datasource := dts.UserDatasourceImpl{
		Config:        c.Config,
		Database:      c.Database,
		DBTransaction: dbTx,
		HashHandler:   hashHandler,
	}

	err = datasource.CheckPhone(ctx, Request.Phone)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	err = datasource.CheckEmail(ctx, Request.Email)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	/// Schema to model
	request := mdl.RegisterRequestModel{
		Firstname: Request.Firstname,
		Lastname:  Request.Lastname,
		Phone:     Request.Phone,
		Email:     Request.Email,
		Password:  Request.Password,
	}

	_, err = datasource.Register(ctx, request)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	err = dbHandler.CommitTx(dbTx)
	if err != nil {
		return nil, err
	}

	result := (err == nil)
	return &result, nil
}
