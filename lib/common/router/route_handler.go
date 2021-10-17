package router

import (
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	ctr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/controllers"
	ech "github.com/labstack/echo/v4"
	ecs "github.com/swaggo/echo-swagger"
)

type RouteHandler struct {
	Config cfg.Config
}

func (r RouteHandler) Handle() *ech.Echo {
	echoServer := ech.New()

	/// Health check
	echoServer.GET("/healtz", func(c ech.Context) error {
		return c.String(http.StatusOK, "Server is Healthy")
	})

	/// Swagger
	echoServer.GET("/swagger/*", ecs.WrapHandler)

	/// V1
	routerV1 := echoServer.Group("/api")

	//  PATH /client/create
	clientController := ctr.ClientController{Config: r.Config}
	routerV1.POST("/client/create", clientController.Create)

	return echoServer
}

/// Legacy
// muxClient := mux.NewRouter()
// muxClient.PathPrefix("/swagger").Handler(hsw.WrapHandler)
// return muxClient
// PATH /user/
// userProfile := new(usc.RetrieveProfile).Execute
// http.HandleFunc("/profile", userProfile).Methods("GET")