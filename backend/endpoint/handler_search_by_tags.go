package endpoint

import (
	"encoding/json"
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	reqData := SearchHandlerRequest{}

	err = json.Unmarshal(body, &reqData)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return
	}

	client := search.NewClient()
	search.GetDataByQuery(client, indexName, reqData.Search)

}
