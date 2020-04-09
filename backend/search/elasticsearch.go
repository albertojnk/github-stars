package search

import (
	"context"
	"github-stars/backend/model"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// IndexBody is the body template for indexing ...
type IndexBody struct {
	ID         string           `json:"id"`
	Repository model.Repository `json:"repository"`
}

// CreateIndex ...
func CreateIndex(index string, data model.User) error {

	client, err := NewClient()
	if err != nil {
		return err
	}

	_, err = Insert2Index(client, index, data.ID, data.Repositories)
	if err != nil {
		log.Printf("error creating index %s", err)
		return err
	}

	return nil
}

// Insert2Index inserts your data on an index ...
func Insert2Index(client *elastic.Client, index string, id string, data []model.Repository) (idx *elastic.IndexResponse, err error) {

	for _, repository := range data {
		body := IndexBody{}
		body.construct(id, repository)
		repoID := strconv.Itoa(repository.ID)

		idx, err = client.Index().
			Index(index).
			Id(repoID).
			BodyJson(&body).
			Do(context.Background())

		if err != nil {
			log.Printf("Error indexing %s to index %s \n", idx.Id, idx.Index)
			return idx, err
		}

		log.Printf("Indexed %s to index %s \n", idx.Id, idx.Index)
	}

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
func GetDataByQuery(client *elastic.Client, index string, id string, query string) ([]model.Repository, error) {
	results := make([]model.Repository, 0)
	user := IndexBody{}

	match := elastic.NewMatchQuery("id", id)
	prefix := elastic.NewPrefixQuery("repository.tags", query)
	bq := elastic.NewBoolQuery().Must(match, prefix)
	searchResult, err := client.Search().
		Index(index).
		Query(bq).
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
func NewClient() (*elastic.Client, error) {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ES_URL")), elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Printf("Error creating new client: %s", err)
		return client, err
	}

	return client, err
}

func (b *IndexBody) construct(id string, data model.Repository) error {
	b.ID = id
	b.Repository = data

	return nil
}
