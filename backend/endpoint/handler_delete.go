package endpoint

import (
	"golang-crud-spa/backend/datasource"
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

	reqData, err := Decode(r.Body, "delete")
	if err != nil {
		log.Printf("something went wrong decoding body, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
	}

	data := reqData.(DeleteRepositoryTagsRequest)

	user, err := datasource.DeleteUserRepositoryTags(data.Username, data.RepositoryID)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	// Encode response into json
	JSONResponse(rw, user, http.StatusOK)

}
