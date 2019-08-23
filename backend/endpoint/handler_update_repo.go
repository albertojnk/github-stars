package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	// Decode the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	reqData := UpdateRepositoryTagsRequest{}

	// Unmarshaling the decoded request
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return
	}

	// update tags on DB
	user, err := UpdateUserRepositoryTags(reqData.Username, reqData.RepositoryID, reqData.Tags)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		return
	}

	// Encode response into json
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user.Repositories)
}
