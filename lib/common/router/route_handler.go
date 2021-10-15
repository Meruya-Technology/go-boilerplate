package router

import (
	_ "github.com/Meruya-Technology/go-boilerplate/docs"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	usc "github.com/Meruya-Technology/go-boilerplate/lib/domain/usecases"
	ctr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/controllers"
	"github.com/gorilla/mux"
	hsw "github.com/swaggo/http-swagger"
)

type RouteHandler struct {
	Config cfg.Config
}

func (r RouteHandler) Handle() *mux.Router {
	muxClient := mux.NewRouter()
	// Swagger
	muxClient.PathPrefix("/swagger").Handler(hsw.WrapHandler)
	http := muxClient.PathPrefix("/api").Subrouter()

	//  PATH /client/create
	clientController := ctr.ClientController{Config: r.Config}
	http.HandleFunc("/client/create", clientController.Create).Methods("POST")

	// PATH /user/
	userProfile := new(usc.RetrieveProfile).Execute
	http.HandleFunc("/profile", userProfile).Methods("GET")

	return muxClient
}
