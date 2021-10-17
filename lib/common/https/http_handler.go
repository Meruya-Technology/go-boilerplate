package https

import (
	"log"
	"net/http"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	ech "github.com/labstack/echo/v4"
)

type HttpHandler struct {
	Router ech.Echo
	Config cfg.Config
}

func (x HttpHandler) StartAndListen() {
	/// Re initiate config
	Config := x.Config

	/// Pack configs
	Host := Config.Host
	Port := Config.Port

	Address := Host + ":" + Port
	srv := &http.Server{
		Handler:      &x.Router,
		Addr:         Address,
		WriteTimeout: Config.WriteTimeout,
		ReadTimeout:  Config.ReadTimeout,
	}
	log.Fatal(srv.ListenAndServe())
}
