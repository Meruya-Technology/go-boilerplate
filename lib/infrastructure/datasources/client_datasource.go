package datasources

import (
	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
	ech "github.com/labstack/echo/v4"
)

type ClientDatasource interface {
	Create(ctx ech.Context, Name string, Secret string) (*mdl.ClientModel, error)
}
