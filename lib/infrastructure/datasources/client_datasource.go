package datasources

import (
	ctx "context"

	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type ClientDatasource interface {
	Create(ctx ctx.Context, Name string, Secret string) (*mdl.ClientModel, error)
}
