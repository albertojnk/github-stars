package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers route each handle registered in this function
func Handlers() http.Handler {
	router := mux.NewRouter()

	{
		handler := StarHandler{}

		router.Handle("/github", handler.Get()).Methods("GET")
	}

	return router
}
