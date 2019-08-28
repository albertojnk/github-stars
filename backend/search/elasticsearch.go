package search

import (
	"context"
	"golang-crud-spa/backend/model"
	"log"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/olivere/elastic.v5"
)

// IndexBody is the body template for indexing ...
type IndexBody struct {
	ID           string `json:"id"`
	Repositories []model.Repository
}

// CreateIndex ...
func CreateIndex(index string, data model.User) error {

	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Printf("Error creating new client: %s", err)
		return err
	}

	// check if index exists
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		log.Printf("Error while checking if index exists: %s", err)
		return err
	}

	if !exists {
		idx, err := Insert2Index(client, index, data)
		if err != nil {
			log.Printf("Index %s created", idx.Index)
		}
	}

	return nil
}

// Insert2Index inserts your data on an index ...
func Insert2Index(client *elastic.Client, index string, data model.User) (*elastic.IndexResponse, error) {
	body := IndexBody{}
	body.construct(data)

	idx, err := client.Index().
		Index(index).
		Type("static").
		Id(data.ID).
		BodyJson(&body).
		Do(context.Background())

	if err != nil {
		log.Printf("Error indexing %s to index %s, type %s\n", idx.Id, idx.Index, idx.Type)
		return idx, err
	}

	log.Printf("Indexed %s to index %s, type %s\n", idx.Id, idx.Index, idx.Type)

	return idx, err
}

// GetDataByID get data from an index ...
func GetDataByID(client *elastic.Client, id string, index string) (*elastic.GetResult, error) {
	resp, err := client.Get().
		Index(index).
		Type("static").
		Id(id).
		Do(context.Background())

	if err != nil {
		log.Printf("Error getting document %s in version %d from index %s, type %s\n", resp.Id, resp.Version, resp.Index, resp.Type)
		return resp, err
	}

	log.Printf("Got document %s in version %d from index %s, type %s\n", resp.Id, resp.Version, resp.Index, resp.Type)

	return resp, err
}

// GetDataByQuery ...
func GetDataByQuery(client *elastic.Client, index string, query string) error {
	fuzzy := elastic.NewFuzzyQuery("tags", query)
	searchResult, err := client.Search().
		Index(index).
		Query(fuzzy).
		From(0).Size(200).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Printf("Error getting document, err: %s", err)
		return err
	}
	user := IndexBody{}
	for _, item := range searchResult.Each(reflect.TypeOf(user)) {
		i := item.(IndexBody)
		spew.Dump(i)
	}
	return nil
}

// NewClient returns a new *elastic.Client
func NewClient() *elastic.Client {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Printf("Error creating new client: %s", err)
	}

	return client
}

func (b *IndexBody) construct(data model.User) error {
	b.ID = data.ID
	b.Repositories = data.Repositories

	return nil
}
