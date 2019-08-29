package search

import (
	"context"
	"golang-crud-spa/backend/model"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// IndexBody is the body template for indexing ...
type IndexBody struct {
	Repository model.Repository `json:"repository"`
}

// CreateIndex ...
func CreateIndex(index string, data model.User) error {

	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Printf("Error creating new client: %s", err)
		return err
	}

	for _, repository := range data.Repositories {
		idx, err := Insert2Index(client, index, repository)
		if err != nil {
			log.Printf("Index %s created", idx.Index)
		}
	}

	return nil
}

// Insert2Index inserts your data on an index ...
func Insert2Index(client *elastic.Client, index string, data model.Repository) (*elastic.IndexResponse, error) {
	body := IndexBody{}
	body.construct(data)
	id := strconv.Itoa(data.ID)

	idx, err := client.Index().
		Index(index).
		Id(id).
		BodyJson(&body).
		Do(context.Background())

	if err != nil {
		log.Printf("Error indexing %s to index %s \n", idx.Id, idx.Index)
		return idx, err
	}

	log.Printf("Indexed %s to index %s \n", idx.Id, idx.Index)

	return idx, err
}

// GetDataByID get data from an index ...
func GetDataByID(client *elastic.Client, id string, index string) (*elastic.GetResult, error) {
	resp, err := client.Get().
		Index(index).
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
func GetDataByQuery(client *elastic.Client, index string, query string) ([]model.Repository, error) {
	results := make([]model.Repository, 0)
	user := IndexBody{}

	prefix := elastic.NewPrefixQuery("repository.tags", query)
	searchResult, err := client.Search().
		Index(index).
		Query(prefix).
		From(0).Size(200).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Printf("Error getting document, err: %s", err)
		return results, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(user)) {
		i := item.(IndexBody)
		results = append(results, model.Repository{
			ID:           i.Repository.ID,
			Name:         i.Repository.Name,
			Description:  i.Repository.Description,
			URL:          i.Repository.URL,
			Language:     i.Repository.Language,
			Tags:         i.Repository.Tags,
			TagSuggester: i.Repository.TagSuggester,
		})

	}
	return results, nil
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

func (b *IndexBody) construct(data model.Repository) error {
	b.Repository = data

	return nil
}
