package repositories_impl

import (
	"fmt"

	en "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	ds "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	mp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/mappers"
	"go.uber.org/zap"
)

type ClientRepositoryImpl struct {
	Datasource ds.ClientDatasource
}

func (repo ClientRepositoryImpl) Create() *en.Client {

	/// Initiate
	logger, _ := zap.NewProduction()
	datasource := repo.Datasource
	result, err := datasource.Create()

	/// Error handling
	if err != nil {
		fmt.Print(err.Error())
		logger.Fatal(err.Error())
		return nil
	}

	/// Compile Result
	mapper := mp.ClientMapper{}
	entity := mapper.ToEntity(result)
	return entity
}
