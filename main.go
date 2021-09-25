package main

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/http"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

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
