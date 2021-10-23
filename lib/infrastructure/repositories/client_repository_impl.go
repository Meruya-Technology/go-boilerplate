package repositories_impl

import (
	"database/sql"

	ctx "context"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
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

	/// Initiate

	dbTx, err := c.Database.Begin()
	if err != nil {
		return nil, err
	}
	datasource := dts.ClientDatasourceImpl{Config: c.Config, DBTransaction: dbTx}
	result, err := datasource.Create(ctx, Request.Name, Request.Secret)

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
	mapper := mp.ClientMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}
