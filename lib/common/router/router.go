package router

import (
	_ "github.com/Meruya-Technology/go-boilerplate/docs"

	"github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct{}

func (router Router) Handle() *mux.Router {

	muxClient := mux.NewRouter()
	// Swagger
	muxClient.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	http := muxClient.PathPrefix("/api/v1").Subrouter()

	/// "/user/"
	userProfile := new(usecases.RetrieveUser).Execute
	http.HandleFunc("/profile", userProfile).Methods("GET")

	/// "/user/{key}/"
	// http.HandleFunc("/{key}/", getUser).Methods("GET")
	return muxClient
}
