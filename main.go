package main

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/http"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/router"
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
	/// Initialize router
	routeHandler := router.RouteHandler{}
	router := routeHandler.Handle()

	/// Initialize config
	configHandler := config.ConfigHandler{}
	config := configHandler.LoadConfig()

	/// Initialize Http handler
	http := http.HttpHandler{
		Router: *router,
		Config: config,
	}
	http.StartAndListen()

}
