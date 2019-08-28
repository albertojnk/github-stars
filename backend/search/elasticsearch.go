package search

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

// StartDefaultClient creates a new client with default options
func StartDefaultClient() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
}
