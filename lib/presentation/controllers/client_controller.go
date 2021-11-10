package controllers

import (
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	imp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
	ech "github.com/labstack/echo/v4"
)

type ClientController struct {
	Config   cfg.Config
	Database sql.DB
}

func (c ClientController) Repository() rep.ClientRepository {
	return imp.ClientRepositoryImpl{
		Config:   c.Config,
		Database: c.Database,
	}
}

func (c ClientController) Create(ctx ech.Context) error {
	usecase := usc.CreateClient{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}

func (c ClientController) Check(ctx ech.Context) error {
	usecase := usc.CheckClient{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}
