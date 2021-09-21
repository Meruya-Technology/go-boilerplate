package main

import (
	"github.com/Meruya-Technology/go-boilerplate/common/http"
	"github.com/Meruya-Technology/go-boilerplate/common/router"
)

func main() {
	router := new(router.Router)
	routeHandler := router.Handle()
	http := new(http.HttpHandler)
	http.Initialize(*routeHandler)
}
