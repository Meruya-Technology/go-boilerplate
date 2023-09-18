package controllers

import (
	"database/sql"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	imp "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
	ech "github.com/labstack/echo/v4"
)

type AccessController struct {
	Config   cfg.Config
	Database *sql.DB
}

func (c AccessController) Repository() rep.AccessTokenRepository {
	return imp.AccessTokenRepositoryImpl{
		Config:   c.Config,
		Database: c.Database,
	}
}

func (c AccessController) Refresh(ctx ech.Context) error {
	usecase := usc.RefreshToken{Repository: c.Repository(), Config: c.Config}
	return usecase.Execute(ctx)
}
