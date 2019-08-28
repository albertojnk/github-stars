package endpoint

import (
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/model"
	"golang-crud-spa/backend/search"
	"io/ioutil"
	"log"
	"net/http"
)

// SearchHandlerRequest is the request params struct
type SearchHandlerRequest struct {
	ID     string `json:"id"`
	Search string `json:"search"`
}

// SearchHandler is the handler that will create our repository in the database
func SearchHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	repositories := []model.Repository{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	reqData := SearchHandlerRequest{}

	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		status, err := HandleErrors(err)
		JSONResponse(rw, err, status)
		return
	}

	if reqData.Search == "" {
		// get the user repositories from DB
		users, err := datasource.ListUserRepositories(reqData.ID)
		if err != nil {
			log.Printf("error while accessing DB, err: %s", err)
			status, err := HandleErrors(err)
			JSONResponse(rw, err, status)
			return
		}
		repositories = users.Repositories
	} else {

		client := search.NewClient()
		repositories, err = search.GetDataByQuery(client, indexName, reqData.Search)
		if err != nil {
			status, err := HandleErrors(err)
			JSONResponse(rw, err, status)
			return
		}
	}

	JSONResponse(rw, repositories, http.StatusOK)
}
