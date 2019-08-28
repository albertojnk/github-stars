package endpoint

import (
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/search"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateRepositoryRequest is the struct we will use to unmarshal the request
type CreateRepositoryRequest struct {
	Username string `json:"username" bson:"username"`
}

// CreateRepository is the handler that will create our repository in the database
func CreateRepository(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// getting the username from frontend and decoding it
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	reqData := CreateRepositoryRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return
	}

	respData, err := GetStarredRepositories(reqData.Username)

	err = datasource.CreateUserRepositories(reqData.Username, respData)
	if err != nil {
		log.Printf("error while creating, err: %s", err)
		return
	}

	user, err := datasource.ListUserRepositories(reqData.Username)
	if err != nil {
		log.Printf("error while listing, err: %s", err)
		return
	}

	err = search.CreateIndex(indexName, user)
	if err != nil {
		log.Printf("Error creating index: %s", err)
		return
	}

	// Encode response into json
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(user)
}
