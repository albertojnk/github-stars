package endpoint

import (
	"golang-crud-spa/backend/datasource"
	"log"
	"net/http"
)

// ListRepositories get our repositories from DB and return
func ListRepositories(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	values, ok := r.URL.Query()["username"]

	if !ok || len(values[0]) < 1 {
		log.Println("url parameter 'username' is missing")
		return
	}

	// get the user repositories from DB
	users, err := datasource.ListUserRepositories(values[0])
	if err != nil {
		log.Printf("error while accessing DB, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	// Endoce response into json
	JSONResponse(rw, users, http.StatusOK)
}
