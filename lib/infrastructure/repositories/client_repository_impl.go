package repositories_impl

import (
	"database/sql"

	ctx "context"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/mappers"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
)

type ClientRepositoryImpl struct {
	Config   cfg.Config
	Database sql.DB
}

func (c ClientRepositoryImpl) Create(ctx ctx.Context, Request req.CreateClientRequest) (*ent.Client, error) {
	/// Initialize transaction
	dbHandler := utl.DbHandler{DB: &c.Database}
	dbTx, err := dbHandler.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	/// Initialize datasource
	datasource := dts.ClientDatasourceImpl{Config: c.Config, DBTransaction: dbTx}
	result, err := datasource.Create(ctx, Request.Name, Request.Secret)

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
	mapper := mp.ClientMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}

func (c ClientRepositoryImpl) Check(ctx ctx.Context, Request req.CheckClientRequest) (*ent.Client, error) {

	/// Initiate
	datasource := dts.ClientDatasourceImpl{Config: c.Config, Database: c.Database}
	result, err := datasource.Check(ctx, Request.Secret)

	/// Error handling
	if err != nil {
		return nil, err
	}

	/// Compile Result
	mapper := mp.ClientMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}
