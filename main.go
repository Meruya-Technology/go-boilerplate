package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/common/configs"
	"github.com/Meruya-Technology/go-boilerplate/common/logger"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var config *configs.Config

func main() {
	logger.InitLogger()
	config = configs.Get()
	logger.SetLogLevel(config)
	// Wire everything up
	http := InitializeService()

	// // Run server
	http.SetupAndServe()
}
