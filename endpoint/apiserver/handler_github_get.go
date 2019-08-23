package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/golang-crud-spa/datasource"
)

// StarHandlerRequest is the struct we will use to unmarshal the StarHandler.Get request
type StarHandlerRequest struct {
	Username string `json:"username" bson:"username"`
}

// StarGithubResponse is the struct we will use to unmarshal the github response
type StarGithubResponse struct {
	ID          int      `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	URL         string   `json:"html_url" bson:"url"`
	Language    string   `json:"language" bson:"language"`
	Tags        []string `json:"tags" bson:"tags"`
}

// GetStarred gets (duh) the starred repositories from github
func GetStarred(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := datasource.Connect()

	// getting the username from frontend and decoding it
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	reqData := StarHandlerRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	// Setting up a http request to the github API
	url := fmt.Sprintf("https://api.github.com/users/%s/starred", reqData.Username)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.star+json")

	// Making the request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	// Decoding the request response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	respData := []StarGithubResponse{}

	// Unmarshaling the results
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	// Writing down the results on the database
	db.C("repository").Upsert(
		bson.M{"_id": reqData.Username},
		bson.M{
			"$set": bson.M{"repositories": respData},
		},
	)
	// Encoding the response back to frontend
	json.NewEncoder(rw).Encode(respData)

}
