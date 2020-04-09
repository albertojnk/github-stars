package endpoint

import (
	"github-stars/backend/datasource"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
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

	err = data.validate()
	if err != nil {
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	user, err := datasource.DeleteUserRepositoryTags(data.Username, data.RepositoryID, data.Tags)
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

// validate validates...
func (d DeleteRepositoryTagsRequest) validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Username, validation.Required),
		validation.Field(&d.RepositoryID, validation.Required),
		validation.Field(&d.Tags, validation.Required),
	)
}
