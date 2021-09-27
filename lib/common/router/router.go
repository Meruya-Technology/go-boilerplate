package router

import (
	_ "github.com/Meruya-Technology/go-boilerplate/docs"

	u "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct{}

func (router Router) Handle() *mux.Router {

	muxClient := mux.NewRouter()
	// Swagger
	muxClient.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	http := muxClient.PathPrefix("/api").Subrouter()

	/// "/user/"
	userProfile := new(u.RetrieveProfile).Execute
	http.HandleFunc("/profile", userProfile).Methods("GET")

	/// "/user/{key}/"
	// http.HandleFunc("/{key}/", getUser).Methods("GET")
	return muxClient
}
