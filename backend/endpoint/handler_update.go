package endpoint

import (
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/search"
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
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	reqData := UpdateRepositoryTagsRequest{}

	// Unmarshaling the decoded request
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	tags := removeDuplicates(reqData.Tags)

	// update tags on DB
	user, err := datasource.UpdateUserRepositoryTags(reqData.Username, reqData.RepositoryID, tags)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	err = search.CreateIndex(indexName, user)
	if err != nil {
		log.Printf("Error creating index: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	// Encode response into json
	JSONResponse(rw, user.Repositories, http.StatusOK)
}

// removeDuplicates will remove all duplicates from a given string slice
func removeDuplicates(tags []string) []string {
	tagsMap := make(map[string]interface{})
	results := make([]string, 0)

	for _, tag := range tags {
		tagsMap[tag] = nil
	}

	for key := range tagsMap {
		results = append(results, key)
	}

	return results
}