package endpoint

import (
	"encoding/json"
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
	users, err := ListUserRepositories(values[0])
	if err != nil {
		log.Printf("error while accessing DB, err: %s", err)
		return
	}

	// Endoce response into json
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(users)
}
