package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const welcomeMessage = "Welcome to RestAPI"

//NewRouter api router++
func NewRouter(mainRouter *mux.Router) http.Handler {

	//Root route
	mainRouter.HandleFunc("/", index).Methods("GET")
	mainRouter.Handle("/metrics", prometheus.Handler()).Methods("GET")

	apiRouter := mainRouter.PathPrefix("/").Subrouter().StrictSlash(true)

	// init repository
	repository := restapi.NewRepository(database.ProductDb, database.ProductDBLegacy)
	// init manager
	manager := restapi.NewManager(repository)
	restapi.NewRouter(apiRouter, restapi.NewController(manager))
	return apiRouter
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, welcomeMessage)
}