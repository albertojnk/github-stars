package endpoint

import (
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/model"
	"golang-crud-spa/backend/search"
	"io/ioutil"
	"log"
	"net/http"

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

	repositories := []model.Repository{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	reqData := SearchHandlerRequest{}

	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	if reqData.Search == "" {
		// get the user repositories from DB
		users, err := datasource.ListUserRepositories(reqData.ID)
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

		repositories, err = search.GetDataByQuery(client, indexName, reqData.ID, reqData.Search)
		if err != nil {
			status, err := HandleErrors(err)
			c.JSON(status, err)
			return err
		}
	}

	c.JSON(http.StatusOK, repositories)

	return nil
}
