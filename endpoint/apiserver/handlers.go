package apiserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers route each handle registered in this function
func Handlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/github", CreateRepository).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", router))

	return router
}
