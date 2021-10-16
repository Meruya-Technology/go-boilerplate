package https

import (
	"log"
	"net/http"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/gorilla/mux"
)

type HttpHandler struct {
	Router mux.Router
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
