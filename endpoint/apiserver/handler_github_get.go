package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StarHandlerRequest is the struct we will use to unmarshal the StarHandler.Get request
type StarHandlerRequest struct {
	Username string `json:"username" bson:"username"`
}

// StarGithubResponse is the struct we will use to unmarshal the github response
type StarGithubResponse struct {
	Repository
}

// Starred is a type map the response and unmarshal
type Starred struct {
}

// Repository is a type map the response and unmarshal
type Repository struct {
	ID          int    `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"html_url" bson:"url"`
	Language    string `json:"language" bson:"language"`
}

// GetStarred gets (duh) the starred repositories from github
func GetStarred(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	reqData := StarHandlerRequest{}

	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/starred", reqData.Username)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.star+json")

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	respData := []StarGithubResponse{}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	json.NewEncoder(rw).Encode(respData)
}
