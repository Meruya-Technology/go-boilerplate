package datasources

import (
	ctx "context"

	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type AccessTokenDatasource interface {
	Create(ctx ctx.Context, UserId int) (*mdl.AccessTokenModel, error)
	Check(ctx ctx.Context, Token string) (*mdl.AccessTokenModel, error)
	GetById(ctx ctx.Context, AccessTokenId int) (*mdl.AccessTokenModel, error)
	Revoke(ctx ctx.Context, Token string)
}
