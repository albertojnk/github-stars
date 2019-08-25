package search

import (
	"log"
	"os"
	"strings"

	"github.com/blevesearch/bleve"
)

// SearchDIR is the directory that blevesearch indexes should be created
var SearchDIR = "./backend/search/indexes/"

// NewBleveMapping will create a new IndexMapping and return an index and an error
func NewBleveMapping(path string) (bleve.Index, error) {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(SearchDIR+path, mapping)
	if err != nil && strings.Contains(err.Error(), "path already exists") {
		err = RemoveExistingMap(path)
		if err == nil {
			index, err = bleve.New(SearchDIR+path, mapping)
		}
	}
	if err != nil {
		log.Printf("error creating a new bleveMap, err: %s", err)
		return nil, err
	}
	return index, nil
}

// NewBleveIndex gets an bleve.Index and data and index it
func NewBleveIndex(index bleve.Index, data interface{}) error {
	// index some data
	err := index.Index("id", data)
	if err != nil {
		log.Printf("error indexing data, err: %s", err)
		return err
	}

	return nil
}

// RemoveExistingMap removes an existing index map
func RemoveExistingMap(path string) error {
	return os.RemoveAll(SearchDIR + path)
}

// search for some text
// query := bleve.NewMatchQuery("text")
// search := bleve.NewSearchRequest(query)
// searchResults, err := index.Search(search)
// if err != nil {
// 	fmt.Println(err)
// 	return nil, err
// }
// return searchResults, nil
