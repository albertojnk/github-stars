package endpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

var indexName = "starred_repositories"

// Handlers route each handler registered in this function
func Handlers() http.Handler {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "XMLHttpRequest", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PATCH"})
	router.HandleFunc("/create", CreateRepository).Methods("POST")
	router.HandleFunc("/list", ListRepositories).Methods("GET")
	router.HandleFunc("/update", UpdateRepositoryTags).Methods("PATCH")
	router.HandleFunc("/search", SearchHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(origins, methods, headers)(router)))

	return router
}

// ConnectServer ...
func ConnectServer() {
	log.Println("Server is up")
	Handlers()
}

// JSONResponse handles http responses
func JSONResponse(rw http.ResponseWriter, body interface{}, code int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)

	if nil != body {
		json.NewEncoder(rw).Encode(body)
	}
}
