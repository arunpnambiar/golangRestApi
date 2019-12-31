package Routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	url     string
	method  string
	handler func(w http.ResponseWriter, r *http.Request)
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, Route := range load() {
		r.HandleFunc(Route.url, Route.handler).Methods(Route.method)
	}
	return r
}

func load() []Route {
	routers := userRouters
	return routers
}
