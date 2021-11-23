package repositories_impl

import (
	ctx "context"
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
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

	/// Initiate
	dbTx, err := c.Database.Begin()
	if err != nil {
		return nil, err
	}
	datasource := dts.UserDatasourceImpl{Config: c.Config, DBTransaction: dbTx}
	result, err := datasource.Login(ctx, Request.Username, Request.Password)

	/// Error handling
	if err != nil {
		dbTx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = dbTx.Commit()
	if err != nil {
		return nil, err
	}

	/// Compile Result
	mapper := mpr.UserMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}

func (c UserRepositoryImpl) Register(ctx ctx.Context, Request req.RegisterRequest) (*ent.User, error) {

	/// Initiate
	dbTx, err := c.Database.Begin()
	if err != nil {
		return nil, err
	}
	datasource := dts.UserDatasourceImpl{Config: c.Config, DBTransaction: dbTx}
	request := mdl.RegisterRequestModel{
		Firstname: Request.Firstname,
		Lastname:  Request.Lastname,
		Phone:     Request.Phone,
		Email:     Request.Email,
		Password:  Request.Password,
	}
	result, err := datasource.Register(ctx, request)

	/// Error handling
	if err != nil {
		dbTx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = dbTx.Commit()
	if err != nil {
		return nil, err
	}

	/// Compile Result
	mapper := mpr.UserMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}
