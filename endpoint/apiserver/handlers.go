package apiserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers route each handle registered in this function
func Handlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create", CreateRepository).Methods("POST")
	router.HandleFunc("/list", ListRepositories).Methods("GET")
	router.HandleFunc("/update", UpdateRepositoryTags).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8090", router))

	return router
}
