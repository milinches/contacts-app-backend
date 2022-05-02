package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Route struct {
		URI     string
		Method  string
		Handler func(http.ResponseWriter, *http.Request)
	}
)

func Load() []Route {
	routes := userRoutes
	routes = append(routes, postRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
