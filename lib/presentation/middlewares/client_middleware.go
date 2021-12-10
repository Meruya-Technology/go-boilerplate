package middlewares

import (
	"database/sql"
	"net/http"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	ctr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/controllers"
	ech "github.com/labstack/echo/v4"
)

type ClientMiddleware struct {
	Config           cfg.Config
	Database         sql.DB
	ClientController ctr.ClientController
}

func (c ClientMiddleware) CheckClient(h ech.HandlerFunc) ech.HandlerFunc {
	return func(ctx ech.Context) error {
		err := c.ClientController.Check(ctx)
		if res, ok := err.(*ech.HTTPError); ok {
			if res.Code == http.StatusOK {
				return h(ctx)
			}
		}
		return nil
	}
}
