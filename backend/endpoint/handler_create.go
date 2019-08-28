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
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	reqData := CreateRepositoryRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	respData, err := GetStarredRepositories(reqData.Username)

	err = datasource.CreateUserRepositories(reqData.Username, respData)
	if err != nil {
		log.Printf("error while creating, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	user, err := datasource.ListUserRepositories(reqData.Username)
	if err != nil {
		log.Printf("error while listing, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	err = search.CreateIndex(indexName, user)
	if err != nil {
		log.Printf("Error creating index: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	// Encode response into json
	JSONResponse(rw, user, http.StatusCreated)
}
