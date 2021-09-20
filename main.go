package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newDatasource() *datasources.ExampleDatasourceImpl {
	test := datasources.ExampleDatasourceImpl{datasources.ExampleDatasource{}}
	test.ExampleDatasource.Example = test
	return &test
}

func main() {
	c := newDatasource()
	c.GetSomething()
	handleRequest()
}
