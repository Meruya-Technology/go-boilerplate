package repositories

import (
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
)

type ClientRepository interface {
	Create(ctx ech.Context, Request req.CreateClientRequest) (*ent.Client, error)
}
