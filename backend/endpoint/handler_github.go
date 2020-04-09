package endpoint

import (
	"encoding/json"
	"fmt"
	"github-stars/backend/model"
	"io/ioutil"
	"log"
	"net/http"
)

// GetStarredRepositories get starred repositories from github
func GetStarredRepositories(username string) ([]model.Repository, error) {
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

	respData := []model.Repository{}

	// Unmarshaling the results
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return nil, err
	}

	for i, repo := range respData {
		if repo.Language == "" {
			respData[i].TagSuggester = fmt.Sprint("Documentation")
		} else {
			respData[i].TagSuggester = fmt.Sprint(repo.Language)
		}
	}

	return respData, nil
}
