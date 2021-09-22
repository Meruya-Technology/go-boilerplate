package http

import (
	"log"
	"net/http"
	"time"

	config "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
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
	Config := x.config
	Host := Config.Host
	Port := Config.Port
	Address := Host + ":" + Port
	srv := &http.Server{
		Handler:      &x.router,
		Addr:         Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
