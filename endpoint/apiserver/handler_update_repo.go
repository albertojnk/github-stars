package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/golang-crud-spa/datasource"
)

// UpdateRepositoryTagsRequest is a struct based on the requested parameters of UpdateRepositoryTags
type UpdateRepositoryTagsRequest struct {
	Username     string   `json:"username"`
	RepositoryID int      `json:"repo_id"`
	Tags         []string `json:"tags"`
}

// UpdateRepositoryTags will update ONE repository of a given user
func UpdateRepositoryTags(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := datasource.Connect()

	// Decode the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	reqData := UpdateRepositoryTagsRequest{}

	// Unmarshaling the decoded request
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	err = db.C("users").Update(
		bson.M{"_id": reqData.Username, "repositories.id": reqData.RepositoryID},
		bson.M{
			"$set": bson.M{"repositories.$.tags": reqData.Tags},
		},
	)
	if err != nil {
		log.Printf("error updating repository, err: %s", err)
	}

	user := User{}

	err = db.C("users").Find(bson.M{"_id": reqData.Username}).One(&user)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
	}

	json.NewEncoder(rw).Encode(user.Repositories)
}
