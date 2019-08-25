package endpoint

import (
	"encoding/json"
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
	}

	reqData := CreateRepositoryRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	respData, err := GetStarredRepositories(reqData.Username)

	err = CreateUserRepositories(reqData.Username, respData)
	if err != nil {
		log.Printf("error while creating, err: %s", err)
	}

	user, err := ListUserRepositories(reqData.Username)
	if err != nil {
		log.Printf("error while listing, err: %s", err)
	}

	// Encode response into json
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}
