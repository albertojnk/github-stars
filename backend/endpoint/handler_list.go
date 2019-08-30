package endpoint

import (
	"errors"
	"golang-crud-spa/backend/datasource"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// ListRepositories get our repositories from DB and return
func ListRepositories(c echo.Context) error {
	r := c.Request()

	defer r.Body.Close()

	values, ok := r.URL.Query()["username"]

	if !ok || len(values[0]) < 1 {
		log.Println("url parameter 'username' is missing")
		c.JSON(500, errors.New("url parameter 'username' is missing"))
		return errors.New("url parameter 'username' is missing")
	}

	// get the user repositories from DB
	users, err := datasource.ListUserRepositories(values[0])
	if err != nil {
		log.Printf("error while accessing DB, err: %s", err)
		status, err := HandleErrors(err)
		c.JSON(status, err)
		return err
	}

	// Endoce response into json
	c.JSON(http.StatusOK, users)

	return nil
}
