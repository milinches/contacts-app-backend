package routes

import (
	"github.com/gorilla/mux"
)

func Load() []Route {
	routes := allRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
