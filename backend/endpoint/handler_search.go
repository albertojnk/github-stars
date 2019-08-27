package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/blevesearch/bleve"
	"github.com/golang-crud-spa/backend/search"
)

// HandleSearchRequest accepts any string
type HandleSearchRequest struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// HandleSearch will handle the search using a string or substring
func HandleSearch(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Decode the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("something went wrong, err: %s", err)
		return
	}

	value := HandleSearchRequest{}

	// Unmarshaling the decoded request
	err = json.Unmarshal(body, &value)
	if err != nil {
		log.Printf("error while unmarshaling, err: %s", err)
		return
	}

	index, err := bleve.Open(search.SearchDIR + value.ID)

	// QUERIES
	match := bleve.NewMatchQuery(value.Value)
	prefix := bleve.NewPrefixQuery(value.Value)
	fuzzy := bleve.NewFuzzyQuery(value.Value)
	fuzzy.SetFuzziness(2)
	phrase := bleve.NewMatchPhraseQuery(value.Value)
	str := bleve.NewQueryStringQuery(value.Value)

	disjunction := bleve.NewDisjunctionQuery(match, prefix, fuzzy, phrase, str)
	searchRequest := bleve.NewSearchRequestOptions(disjunction, 5573, 0, false)
	searchResult, _ := index.Search(searchRequest)
	docsFromSearch := search.GetOriginalDocsFromSearchResults(searchResult, index)

	spew.Dump(docsFromSearch)
	// for _, value := range docsFromSearch {
	// 	spew.Dump(value)
	// }

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(searchResult)
}
