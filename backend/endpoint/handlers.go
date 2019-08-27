package endpoint

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

// Handlers route each handle registered in this function
func Handlers() http.Handler {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "XMLHttpRequest", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PATCH"})
	router.HandleFunc("/create", CreateRepository).Methods("POST")
	router.HandleFunc("/list", ListRepositories).Methods("GET")
	router.HandleFunc("/update", UpdateRepositoryTags).Methods("PATCH")
	router.HandleFunc("/search", HandleSearch).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(origins, methods, headers)(router)))

	return router
}
