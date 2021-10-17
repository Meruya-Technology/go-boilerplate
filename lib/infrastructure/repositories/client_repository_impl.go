package repositories_impl

import (
	"fmt"

	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	ds "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/mappers"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ClientRepositoryImpl struct {
	Datasource ds.ClientDatasource
}

func (c ClientRepositoryImpl) Create(ctx ech.Context, Request req.CreateClientRequest) (*ent.Client, error) {

	/// Initiate
	logger, _ := zap.NewProduction()
	datasource := c.Datasource
	result, err := datasource.Create(ctx, Request.Name, Request.Secret)

	/// Error handling
	if err != nil {
		fmt.Print(err.Error())
		logger.Fatal(err.Error())
		return nil, err
	}

	/// Compile Result
	mapper := mp.ClientMapper{}
	entity := mapper.ToEntity(*result)
	return entity, err
}
