package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// DeleteRepositoryTagsRequest is a struct based on the requested parameters of DeleteRepositoryTags
type DeleteRepositoryTagsRequest struct {
	Username     string   `json:"username"`
	RepositoryID int      `json:"repo_id"`
	Tags         []string `json:"tags"`
}

// DeleteRepositoryTags is the endpoint that will manage tags deletion
func DeleteRepositoryTags(rw http.ResponseWriter, r *http.Response) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	reqData := DeleteRepositoryTagsRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return
	}

	user, err := DeleteUserRepositoryTags(reqData.Username, reqData.RepositoryID)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	// Encode response into json
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)

}
