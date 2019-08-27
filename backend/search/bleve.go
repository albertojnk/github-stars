package search

import (
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"

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
	spew.Dump("teste")
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

func GetBleveDocsFromSearchResults(results *bleve.SearchResult, index bleve.Index) []map[string]string {
	docs := make([]map[string]string, 0)

	for _, val := range results.Hits {
		id := val.ID
		doc, _ := index.Document(id)

		rv := struct {
			ID     string            `json:"id"`
			Fields map[string]string `json:"fields"`
		}{
			ID:     id,
			Fields: map[string]string{},
		}
		for _, field := range doc.Fields {
			var newval string
			newval = string(field.Value())
			rv.Fields[field.Name()] = newval
		}
		docs = append(docs, rv.Fields)
	}

	return docs
}

func GetOriginalDocsFromSearchResults(results *bleve.SearchResult, index bleve.Index) [][]byte {
	docs := make([][]byte, 0)

	for _, val := range results.Hits {
		id := val.ID
		raw, err := index.GetInternal([]byte(id))
		if err != nil {
			log.Fatal("Trouble getting internal doc:", err)
		}
		docs = append(docs, raw)
	}
	return docs
}
