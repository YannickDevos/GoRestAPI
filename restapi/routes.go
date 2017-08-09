package restapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter1() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		// var handler http.Handler

		// handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Getall",
		"GET",
		"/get",
		Getall,
	},
	Route{
		"Getid",
		"GET",
		"/get/{Id}",
		Getid,
	},
}

// router.HandleFunc("/", Index)
// 	router.HandleFunc("/get/", Getall)
// 	router.HandleFunc("/get/{id}", Getid)
// 	router.HandleFunc("/post/{id}", Createid)
// 	router.HandleFunc("/delete/{id}", Deleteid)
