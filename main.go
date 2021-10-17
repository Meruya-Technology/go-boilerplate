package main

import (
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
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
// @BasePath /api

func main() {
	/// Initialize config
	configHandler := cfg.ConfigHandler{}
	config := configHandler.LoadConfig()

	/// Initialize router
	routeHandler := rtr.RouteHandler{
		Config: config,
	}
	router := routeHandler.Handle()

	/// Initialize Http handler
	http := htt.HttpHandler{
		Router: *router,
		Config: config,
	}
	http.StartAndListen()

}
