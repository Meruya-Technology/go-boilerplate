package main

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/http"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/router"
)

// @title OAuth2 API
// @version 1.0
// @description Authentications api with OAuth2

// @contact.name Meruya Technology Founder
// @contact.email dwikurnianto.mulyadien@gmail.com

// @host localhost:8000
// @BasePath api/v1
// @query.collection.format multi

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
