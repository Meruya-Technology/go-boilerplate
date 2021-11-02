package repositories_impl

import (
	ctx "context"
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/mappers"
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
	mapper := mp.UserMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}
