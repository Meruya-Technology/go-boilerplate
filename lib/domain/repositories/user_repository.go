package repositories

import (
	ctx "context"

	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
)

type UserRepository interface {
	Login(ctx ctx.Context, Request req.LoginRequest) (*ent.User, error)
}
