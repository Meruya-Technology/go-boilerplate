package datasources

import (
	ctx "context"

	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type RefreshTokenDatasource interface {
	/// This method will create new refresh token, by supplied access token id
	Create(ctx ctx.Context, AccessTokenId int) (*mdl.RefreshTokenModel, error)

	/// This method will check existing refresh token, and returning a access token id
	Check(ctx ctx.Context, Token string) (*mdl.RefreshTokenModel, error)

	/// This method will revoke existing refresh token
	Revoke(ctx ctx.Context, Token string) (bool, error)
}
