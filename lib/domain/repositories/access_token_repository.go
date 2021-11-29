package repositories

import (
	ctx "context"

	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type AccessTokenRepositoryImpl interface {
	/// Create / issues an access token, this will also create a refresh token
	Create(ctx ctx.Context, UserId int, DeviceId string) (*ent.AccessToken, error)
	/// Refresh access token with refresh token which will be refreshed also
	Refresh(ctx ctx.Context, RefreshToken string) (*ent.User, error)
	/// Revoke access token, this will revoke refresh token also
	Revoke(ctx ctx.Context, Token string) (*bool, error)
}
