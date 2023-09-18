package controllers

import (
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	imp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
	ech "github.com/labstack/echo/v4"
)

type UserController struct {
	Config   cfg.Config
	Database *sql.DB
}

func (c UserController) Repository() rep.UserRepository {
	return imp.UserRepositoryImpl{
		Config:   c.Config,
		Database: c.Database,
	}
}

func (c UserController) Login(ctx ech.Context) error {
	usecase := usc.Login{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}

func (c UserController) Register(ctx ech.Context) error {
	usecase := usc.Register{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}
