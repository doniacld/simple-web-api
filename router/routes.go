package router

import (
	"net/http"

	"github.com/doniacld/simple-web-api/handlers"
	"github.com/doniacld/simple-web-api/logger"

	"github.com/gorilla/mux"
)

// NewRouter creates a router with all the methods
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

// Route holds information about a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoCreate",
		http.MethodPost,
		"/todos",
		handlers.TodoCreate,
	},
	Route{
		"Index",
		http.MethodGet,
		"/",
		handlers.Index,
	},
	Route{
		"TodoIndex",
		http.MethodGet,
		"/todos",
		handlers.TodosRetrieve,
	},
	Route{
		"TodoShow",
		http.MethodGet,
		"/todos/{todoID}",
		handlers.TodoShow,
	},
	Route{
		"TodoDelete",
		http.MethodDelete,
		"/todos/{todoID}",
		handlers.TodoDelete,
	},
}
