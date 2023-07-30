package middlewares

import (
	"database/sql"
	"fmt"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	repositories_impl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
	ctr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/controllers"
	"github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
)

type ClientMiddleware struct {
	Config           cfg.Config
	Database         sql.DB
	ClientController ctr.ClientController
}

func (c ClientMiddleware) CheckClient(h ech.HandlerFunc) ech.HandlerFunc {
	return func(ctx ech.Context) error {
		repo := repositories_impl.ClientRepositoryImpl{Config: c.Config, Database: c.Database}
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return https.ErrorBadRequest(ctx, fmt.Errorf("Secret key is required"), nil)
		}
		request := requests.CheckClientRequest{
			Secret: token,
		}
		newContext := ctx.Request().Context()
		_, err := repo.Check(newContext, request)
		if err != nil {
			return https.ErrorInternalServerResponse(ctx, err, nil)
		}
		return h(ctx)
	}
}
