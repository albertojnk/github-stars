package endpoint

import (
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/model"
	"golang-crud-spa/backend/search"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
)

// SearchHandlerRequest is the request params struct
type SearchHandlerRequest struct {
	ID     string `json:"id"`
	Search string `json:"search"`
}

// SearchHandler is the handler that will create our repository in the database
func SearchHandler(c echo.Context) error {
	r := c.Request()

	defer r.Body.Close()

	reqData, err := Decode(r.Body, "search")
	if err != nil {
		log.Printf("something went wrong decoding body, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	data := reqData.(SearchHandlerRequest)

	err = data.validate()
	if err != nil {
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	repositories := []model.Repository{}

	if data.Search == "" {
		// get the user repositories from DB
		users, err := datasource.ListUserRepositories(data.ID)
		if err != nil {
			log.Printf("error while accessing DB, err: %s", err)
			status, err := HandleErrors(err)
			c.JSON(status, err)
			return err
		}
		repositories = users.Repositories

	} else {

		client, err := search.NewClient()
		if err != nil {
			status, err := HandleErrors(err)
			c.JSON(status, err)
			return err
		}

		repositories, err = search.GetDataByQuery(client, indexName, data.ID, data.Search)
		if err != nil {
			status, err := HandleErrors(err)
			c.JSON(status, err)
			return err
		}
	}

	c.JSON(http.StatusOK, repositories)

	return nil
}

func (s SearchHandlerRequest) validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required),
	)
}
