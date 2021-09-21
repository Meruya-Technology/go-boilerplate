package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct{}

func (router Router) Handle() *mux.Router {

	// var dir string
	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()
	// r := mux.NewRouter()

	// // This will serve files under http://localhost:8000/static/<filename>
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	muxClient := mux.NewRouter()
	http := muxClient.PathPrefix("/test").Subrouter()

	// / "/user/"
	http.HandleFunc("/", user).Methods("GET")

	// / "/user/{key}/"
	http.HandleFunc("/{key}/", getUser).Methods("GET")
	return muxClient
}

func user(http.ResponseWriter, *http.Request) {
	println("here is your user")

}

func getUser(http.ResponseWriter, *http.Request) {
	println("here is your get user")
}
