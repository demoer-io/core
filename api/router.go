package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Router main function
func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/github/code", GithubCode).Methods("POST")

	return router
}
