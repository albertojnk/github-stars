package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StarGithubResponse is the struct we will use to unmarshal the github response
type StarGithubResponse struct {
	ID          int      `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	URL         string   `json:"html_url" bson:"url"`
	Language    string   `json:"language" bson:"language"`
	Tags        []string `json:"tags" bson:"tags"`
}

// GetStarredRepositories gets (duh) the starred repositories from github
func GetStarredRepositories(username string) []StarGithubResponse {
	// Setting up a http request to the github API
	url := fmt.Sprintf("https://api.github.com/users/%s/starred", username)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.star+json")

	// Making the request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	// Decoding the request response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
	}

	respData := []StarGithubResponse{}

	// Unmarshaling the results
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
	}

	return respData
}
