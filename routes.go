package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}
type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = Logger(handler,route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}



var routes = Routes{
	Route{
	"RomanNumber",
	"GET",
	"/roman/{number}",
	Index,
	},
	Route{
	"Todos",
	"GET",
	"/todos",
	TodoIndex,
	},
}