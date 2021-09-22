package main

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/http"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/router"
)

// @title Example API
// @version 1.0
// @description This is a sample service for managing
// @termsOfService http://swagger.io/terms/
// @contact.name Meruya Technology Founder
// @contact.email dwikurnianto.mulyadien@gmail.com
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

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
