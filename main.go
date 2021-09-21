package main

import (
	"github.com/Meruya-Technology/go-boilerplate/common/http"
	"github.com/Meruya-Technology/go-boilerplate/common/router"
)

func main() {
	/// Initialize router
	router := new(router.Router)

	/// Handle router
	routeHandler := router.Handle()

	/// Initialize Http handler
	http := new(http.HttpHandler)

	/// Bind router to HttpHandler
	http.Initialize(*routeHandler)
}
