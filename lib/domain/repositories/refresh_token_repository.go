package repositories

import (
	ctx "context"

	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	res "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/responses"
)

type RefreshTokenRepository interface {
	Refresh(ctx ctx.Context, Request req.RefreshTokenRequest) (*res.LoginResponse, error)
}
