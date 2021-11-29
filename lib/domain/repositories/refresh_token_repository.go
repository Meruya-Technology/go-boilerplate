package repositories

import (
	ctx "context"

	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type RefreshTokenRepository interface {
	///
	Create(ctx ctx.Context, AccessTokenId int, RefreshToken string) (*ent.RefreshToken, error)
	///
	Revoke(ctx ctx.Context, id int) (*bool, error)
}
