package main

import (
	"context"
	"os"
	"os/signal"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	dtb "github.com/Meruya-Technology/go-boilerplate/lib/common/database"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rtr "github.com/Meruya-Technology/go-boilerplate/lib/common/router"
)

// @title Oauth2 API Documentation
// @version 1.0
// @description Go boiler plate with Oauth2 implementation, documented with Swagger
// @contact.name Meruya Technology
// @contact.url https://blog.meruyatechnology.com
// @contact.email support@meruyatechnology.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/auth

// @securitydefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @securitydefinitions.apikey  ClientSecret
// @in                          header
// @name                        Authorization

func main() {
	/// Initialize config
	configHandler := cfg.ConfigHandler{}
	config := configHandler.LoadConfig()

	/// Initialize db
	var dbHandler dtb.Postgresql = dtb.PostgresqlImpl{Config: config}
	database := dbHandler.Initalize()

	/// Initialize router
	routeHandler := rtr.RouteHandler{
		Config:   config,
		Database: *database,
	}
	router := routeHandler.Handle()

	/// Initialize Http handler
	http := htt.HttpHandler{
		Router: *router,
		Config: config,
	}
	http.StartAndListen()

	// * code for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), config.WriteTimeout)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}
}
