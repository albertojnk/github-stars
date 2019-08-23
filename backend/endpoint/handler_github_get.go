package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StarredRepositories is the struct we will use to unmarshal the github response
type StarredRepositories struct {
	ID          int      `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	URL         string   `json:"html_url" bson:"url"`
	Language    string   `json:"language" bson:"language"`
	Tags        []string `json:"tags" bson:"tags"`
}

// GetStarredRepositories gets (duh) the starred repositories from github
func GetStarredRepositories(username string) ([]StarredRepositories, error) {
	// Setting up a http request to the github API
	url := fmt.Sprintf("https://api.github.com/users/%s/starred?per_page=200", username)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github.v3.star+json")
	req.Header.Set("Content-Type", "application/json")

	// Making the request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return nil, err
	}

	// Decoding the request response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return nil, err
	}

	respData := []StarredRepositories{}

	// Unmarshaling the results
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return nil, err
	}

	return respData, nil
}
