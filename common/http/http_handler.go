package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpHandler struct {
	router mux.Router
}

func (x HttpHandler) Initialize(z mux.Router) {
	x.router = z
	x.StartAndListen()
}

func (x HttpHandler) StartAndListen() {
	srv := &http.Server{
		Handler:      &x.router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
