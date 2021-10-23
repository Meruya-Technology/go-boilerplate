package repositories

import (
	ctx "context"

	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
)

type ClientRepository interface {
	Create(ctx ctx.Context, Request req.CreateClientRequest) (*ent.Client, error)
}
