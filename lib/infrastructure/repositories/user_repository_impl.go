package repositories_impl

import (
	ctx "context"
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	sec "github.com/Meruya-Technology/go-boilerplate/lib/common/security"
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mpr "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/mappers"
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
)

type UserRepositoryImpl struct {
	Config   cfg.Config
	Database sql.DB
}

func (c UserRepositoryImpl) Login(ctx ctx.Context, Request req.LoginRequest) (*ent.User, error) {
	/// Initiate transaction
	dbHandler := utl.DbHandler{DB: &c.Database}
	hashHandler := sec.HashHandler{}
	dbTx, err := dbHandler.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	datasource := dts.UserDatasourceImpl{Config: c.Config, DBTransaction: dbTx, HashHandler: hashHandler}
	result, err := datasource.Login(ctx, Request.Email, Request.Password)

	/// Error handling
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	err = dbHandler.CommitTx(dbTx)
	if err != nil {
		return nil, err
	}

	/// Compile Result
	mapper := mpr.UserMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}

func (c UserRepositoryImpl) Register(ctx ctx.Context, Request req.RegisterRequest) (*ent.User, error) {
	/// Initiate transaction
	dbHandler := utl.DbHandler{DB: &c.Database}
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

	result, err := datasource.Register(ctx, request)
	if err != nil {
		dbHandler.RollbackTx(dbTx)
		return nil, err
	}

	err = dbHandler.CommitTx(dbTx)
	if err != nil {
		return nil, err
	}

	entity := ent.User{
		Id: *result,
	}
	return &entity, err
}
