package controllers

import (
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	dts "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
	imp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
	ech "github.com/labstack/echo/v4"
)

type ClientController struct {
	Config   cfg.Config
	Database sql.DB
}

func (c ClientController) Repository() rep.ClientRepository {
	var clientDatasource dts.ClientDatasource = dts.ClientDatasourceImpl{Config: c.Config}
	return imp.ClientRepositoryImpl{Datasource: clientDatasource}
}

func (c ClientController) Create(ctx ech.Context) error {
	usecase := usc.CreateClient{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}
