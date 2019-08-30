package endpoint

import (
	"golang-crud-spa/backend/datasource"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// DeleteRepositoryTagsRequest is a struct based on the requested parameters of DeleteRepositoryTags
type DeleteRepositoryTagsRequest struct {
	Username     string   `json:"username"`
	RepositoryID int      `json:"repo_id"`
	Tags         []string `json:"tags"`
}

// DeleteRepositoryTags is the endpoint that will manage tags deletion
func DeleteRepositoryTags(c echo.Context) error {
	r := c.Request()
	defer r.Body.Close()

	reqData, err := Decode(r.Body, "delete")
	if err != nil {
		log.Printf("something went wrong decoding body, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	data := reqData.(DeleteRepositoryTagsRequest)

	user, err := datasource.DeleteUserRepositoryTags(data.Username, data.RepositoryID)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	// Encode response into json
	c.JSON(http.StatusOK, user)

	return nil
}
