package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// StarHandlerRequest is the struct we will use to unmarshal the StarHandler.Get request
type StarHandlerRequest struct {
	Username string `json:"username" bson:"username"`
}

// StarGithubResponse is the struct we will use to unmarshal the github response
type StarGithubResponse struct {
	Starred []Starred
}

// Starred is a type map the response and unmarshal
type Starred struct {
	Repo Repository `json:"repo"`
}

// Repository is a type map the response and unmarshal
type Repository struct {
	ID          int    `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"html_url" bson:"url"`
	Language    string `json:"language" bson:"language"`
}

// Get gets (duh) the repositories starred from github
func (h StarHandler) Get() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(fmt.Errorf("something went wrong, err: %s", err))
			return
		}

		reqData := StarHandlerRequest{}

		err = json.Unmarshal(body, &reqData)
		if err != nil {
			fmt.Println(fmt.Errorf("error while unmarshaling, err: %s", err))
			return
		}

		url := fmt.Sprintf("https://api.github.com/users/%s/starred", reqData.Username)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(fmt.Errorf("something went wrong, err: %s", err))
			return
		}

		respData := StarGithubResponse{}

		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			fmt.Println(fmt.Errorf("error while decoding, err: %s", err))
			return
		}

	})
}
