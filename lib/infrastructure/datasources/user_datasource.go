package datasources

import (
	ctx "context"

	mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type UserDatasource interface {
	User() (string, error)
	Login(ctx ctx.Context, Email string, Password string) (*mdl.UserModel, error)
	Register(ctx ctx.Context, Request mdl.RegisterRequestModel) (*int, error)
	CheckPhone(ctx ctx.Context, Phone string) error
	CheckEmail(ctx ctx.Context, Email string) error
}
