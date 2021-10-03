package http

import (
	"log"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/gorilla/mux"
)

type HttpHandler struct {
	router mux.Router
	config config.Config
}

func (x HttpHandler) Initialize(z mux.Router, c config.Config) {
	x.router = z
	x.config = c
	x.StartAndListen()
}

func (x HttpHandler) StartAndListen() {
	/// Re initiate config
	Config := x.config

	/// Pack configs
	Host := Config.Host
	Port := Config.Port

	Address := Host + ":" + Port
	srv := &http.Server{
		Handler:      &x.router,
		Addr:         Address,
		WriteTimeout: Config.WriteTimeout,
		ReadTimeout:  Config.ReadTimeout,
	}
	log.Fatal(srv.ListenAndServe())
}
