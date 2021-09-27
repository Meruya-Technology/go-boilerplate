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
// @host http://localhost:8080
// @BasePath /api

func main() {
	/// Initialize router
	router := new(router.Router)

	/// Handle router
	routeHandler := router.Handle()

	/// Initialize config
	config := new(config.Config)

	/// Initialize Http handler
	http := new(http.HttpHandler)

	/// Bind router to HttpHandler
	http.Initialize(*routeHandler, config.LoadConfig())

}
