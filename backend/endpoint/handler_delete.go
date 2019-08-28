package endpoint

import (
	"encoding/json"
	"golang-crud-spa/backend/datasource"
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
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	reqData := DeleteRepositoryTagsRequest{}

	// Unmarshaling the decoded username
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	user, err := datasource.DeleteUserRepositoryTags(reqData.Username, reqData.RepositoryID)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	// Encode response into json
	JSONResponse(rw, user, http.StatusOK)

}
