package main

import (
	"app/helper"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter : mux.router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// CORS preflight
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.WriteHeader(http.StatusAccepted)
		})
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = helper.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
