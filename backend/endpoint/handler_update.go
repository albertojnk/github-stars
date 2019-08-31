package endpoint

import (
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/search"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
)

// UpdateRepositoryTagsRequest is a struct based on the requested parameters of UpdateRepositoryTags
type UpdateRepositoryTagsRequest struct {
	Username     string   `json:"username"`
	RepositoryID int      `json:"repo_id"`
	Tags         []string `json:"tags"`
}

// UpdateRepositoryTags will update ONE repository of a given user
func UpdateRepositoryTags(c echo.Context) error {
	r := c.Request()

	defer r.Body.Close()

	reqData, err := Decode(r.Body, "update")
	if err != nil {
		log.Printf("something went wrong decoding body, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	data := reqData.(UpdateRepositoryTagsRequest)

	err = data.validate()
	if err != nil {
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	tags := removeDuplicates(data.Tags)

	// update tags on DB
	user, err := datasource.UpdateUserRepositoryTags(data.Username, data.RepositoryID, tags)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	err = search.CreateIndex(indexName, user)
	if err != nil {
		log.Printf("Error creating index: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	// Encode response into json
	c.JSON(http.StatusCreated, user.Repositories)

	return nil
}

// validate UpdateRepositoryTagsRequest
func (u UpdateRepositoryTagsRequest) validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.RepositoryID, validation.Required),
		validation.Field(&u.Tags, validation.Required),
	)
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
