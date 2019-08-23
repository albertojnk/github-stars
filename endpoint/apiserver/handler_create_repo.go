package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/golang-crud-spa/datasource"
)

// CreateRepositoryRequest is the struct we will use to unmarshal the request
type CreateRepositoryRequest struct {
	Username string `json:"username" bson:"username"`
}

// CreateRepository is the handler that will create our repository in the database
func CreateRepository(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := datasource.Connect()

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

	respData := GetStarredRepositories(reqData.Username)

	// Writing down the results on the database
	db.C("repository").Upsert(
		bson.M{"_id": reqData.Username},
		bson.M{
			"$set": bson.M{"repositories": respData},
		},
	)

	// Encode the results and send back to requester
	json.NewEncoder(rw).Encode(respData)
}
