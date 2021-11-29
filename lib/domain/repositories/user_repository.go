package repositories

import (
	ctx "context"

	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	res "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/responses"
)

type UserRepository interface {
	Login(ctx ctx.Context, Request req.LoginRequest) (*res.LoginResponse, error)
	Register(ctx ctx.Context, Request req.RegisterRequest) (*bool, error)
}
