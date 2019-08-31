package search_test

import (
	"context"
	"golang-crud-spa/backend/model"
	"golang-crud-spa/backend/search"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var indexName = "starred_repositories"

func TestElasticsearch(t *testing.T) {
	Convey("Given an index and a model.User{}", t, func() {
		user := model.User{
			ID: "test_elastic_user1",
			Repositories: []model.Repository{
				model.Repository{
					ID:           1,
					Name:         "test_name1",
					Description:  "test repository tags elastic 1",
					URL:          "http://google.com",
					Language:     "Gotest",
					Tags:         []string{"Go", "Python", "DOCKER", "NodeJS"},
					TagSuggester: "Gotest",
				},
				model.Repository{
					ID:           2,
					Name:         "test_name1",
					Description:  "test repository tags elastic 1",
					URL:          "http://google.com",
					Language:     "Gotest",
					Tags:         []string{"JavaScript", "nodejs", "java"},
					TagSuggester: "Gotest",
				},
			},
		}

		Convey("When calling function CreateIndex", func() {
			err := search.CreateIndex(indexName, user)
			So(err, ShouldBeNil)

			Convey("Then the index must exists", func() {
				client, _ := search.NewClient()

				exists, err := client.IndexExists(indexName).Do(context.Background())

				So(err, ShouldBeNil)
				So(exists, ShouldBeTrue)

				Convey("Then when calling function GetDataByQuery", func() {
					query := "no"
					client, _ := search.NewClient()
					repos, err := search.GetDataByQuery(client, indexName, user.ID, query)

					So(err, ShouldBeNil)
					So(repos, ShouldNotBeNil)

					ok := false
					for _, repo := range repos {
						for _, tag := range repo.Tags {
							if strings.Contains(tag, query) {
								ok = true
							}
						}
					}

					So(ok, ShouldBeTrue)
				})
			})

		})

	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
