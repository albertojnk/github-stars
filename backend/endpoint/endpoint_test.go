package endpoint_test

import (
	"bytes"
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/endpoint"
	"golang-crud-spa/backend/model"
	"golang-crud-spa/backend/search"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/globalsign/mgo"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateRepository(t *testing.T) {
	Convey("Given a POST request to /create", t, func() {
		newURL, _ := url.Parse("/create")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

		clearDB(datasource.GetDatabase())

		Convey("With a valid body", func() {
			data := endpoint.CreateRepositoryRequest{
				Username: "albertojnk",
			}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					So(resp.Body, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusCreated)
				})
			})
		})

		Convey("With an invalid body", func() {
			data := endpoint.CreateRepositoryRequest{}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					So(resp.Body, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
		})
	})
}

func TestDeleteRepositoryTags(t *testing.T) {
	Convey("Given a DELETE request to /delete", t, func() {
		newURL, _ := url.Parse("/delete")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

		db := datasource.GetDatabase()
		clearDB(db)

		testRepo := []model.Repository{
			model.Repository{
				ID:           12345,
				Name:         "test_name",
				Description:  "test repository tags deletion",
				URL:          "http://google.com",
				Language:     "Gotest",
				Tags:         []string{"Test1", "Test2", "Test3"},
				TagSuggester: "Gotest",
			},
		}

		db.C("users").Upsert(
			bson.M{"_id": "test_delete_user1"},
			bson.M{
				"$set": bson.M{"repositories": testRepo},
			},
		)

		Convey("With a valid body", func() {
			data := endpoint.DeleteRepositoryTagsRequest{
				Username:     "test_delete_user1",
				RepositoryID: 12345,
				Tags:         []string{"Test1", "Test3"},
			}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("DELETE", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := model.User{}
					json.Unmarshal(resp.Body.Bytes(), &response)
					for _, rep := range response.Repositories {
						if rep.ID == 12345 {
							So(len(rep.Tags), ShouldEqual, 1)
							So(rep.Tags[0], ShouldEqual, "Test2")
							break
						}
					}
					So(resp.Code, ShouldEqual, http.StatusOK)
				})
			})
		})

		Convey("With an invalid body", func() {
			data := endpoint.DeleteRepositoryTagsRequest{}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("DELETE", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					So(resp.Body, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
		})
	})
}

func TestListRepositoryTags(t *testing.T) {
	Convey("Given a GET request to /list", t, func() {
		newURLValidParam, _ := url.Parse("/list?username=test_list_user1")
		newURLInvalidParam, _ := url.Parse("/list")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

		db := datasource.GetDatabase()
		clearDB(db)

		testUsers := []model.User{
			model.User{
				ID: "test_list_user1",
				Repositories: []model.Repository{
					model.Repository{
						ID:           1,
						Name:         "test_name1",
						Description:  "test repository tags list 1",
						URL:          "http://google.com",
						Language:     "Gotest",
						Tags:         []string{"Test1", "Test2", "Test3"},
						TagSuggester: "Gotest",
					},
					model.Repository{
						ID:           2,
						Name:         "test_name1",
						Description:  "test repository tags list 1",
						URL:          "http://google.com",
						Language:     "Gotest",
						Tags:         []string{"Test1", "Test2", "Test3"},
						TagSuggester: "Gotest",
					},
				},
			},
			model.User{
				ID: "test_list_user2",
				Repositories: []model.Repository{
					model.Repository{
						ID:           3,
						Name:         "test_name2",
						Description:  "test repository tags list 2",
						URL:          "http://google.com",
						Language:     "Gotest",
						Tags:         []string{"Test1", "Test2", "Test3"},
						TagSuggester: "Gotest",
					},
				},
			},
		}

		bulk := db.C("users").Bulk()
		for _, user := range testUsers {
			bulk.Insert(user)
		}

		bulk.Run()

		Convey("With a valid url parameter", func() {
			req := httptest.NewRequest("GET", server.URL+newURLValidParam.String(), nil)
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := model.User{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(resp.Body, ShouldNotBeNil)
					So(response.ID, ShouldEqual, "test_list_user1")
					So(response.Repositories, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusOK)
				})
			})
		})

		Convey("With an invalid url parameter or missing parameter", func() {
			req := httptest.NewRequest("GET", server.URL+newURLInvalidParam.String(), nil)
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					So(resp.Body, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
		})
	})
}

func TestUpdateRepositoryTags(t *testing.T) {
	Convey("Given a PATCH request to /update", t, func() {
		newURL, _ := url.Parse("/update")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

		db := datasource.GetDatabase()
		clearDB(db)

		testUser := model.User{
			ID: "test_update_user1",
			Repositories: []model.Repository{
				model.Repository{
					ID:           1,
					Name:         "test_name1",
					Description:  "test repository tags update 1",
					URL:          "http://google.com",
					Language:     "Gotest",
					Tags:         []string{"Test1", "Test2", "Test3"},
					TagSuggester: "Gotest",
				},
				model.Repository{
					ID:           2,
					Name:         "test_name1",
					Description:  "test repository tags update 2",
					URL:          "http://uol.com",
					Language:     "Gotest",
					Tags:         []string{"Test1", "Test2", "Test3"},
					TagSuggester: "Gotest",
				},
			},
		}

		db.C("users").Insert(testUser)

		Convey("With a valid body", func() {
			tags := []string{"Test1", "Test2", "Test3", "Test4", "Test5"}
			data := endpoint.UpdateRepositoryTagsRequest{
				Username:     "test_update_user1",
				RepositoryID: 1,
				Tags:         tags,
			}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("PATCH", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := []model.Repository{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(resp.Body, ShouldNotBeNil)
					So(response, ShouldNotBeNil)

					ok := true
					for _, repo := range response {
						if repo.ID == 1 {
							for _, tag := range repo.Tags {
								if !contains(tags, tag) {
									ok = false
								}
							}
						}
					}

					So(ok, ShouldEqual, true)
					So(resp.Code, ShouldEqual, http.StatusCreated)
				})
			})
		})

		Convey("With an invalid body", func() {
			data := endpoint.UpdateRepositoryTagsRequest{}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("PATCH", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := []model.Repository{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(response, ShouldBeEmpty)
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
		})
	})
}

func TestSearchByTags(t *testing.T) {
	Convey("Given a POST request to /search", t, func() {
		newURL, _ := url.Parse("/search")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

		db := datasource.GetDatabase()
		clearDB(db)

		client, _ := search.NewClient()
		client.DeleteIndex("starred_repositories")

		username := "test_search_user1"
		testUser := model.User{
			ID: username,
			Repositories: []model.Repository{
				model.Repository{
					ID:           123,
					Name:         "test_name1",
					Description:  "test repository tags search 1",
					URL:          "http://google.com",
					Language:     "Gotest",
					Tags:         []string{"Test1", "Go", "VueJS"},
					TagSuggester: "Gotest",
				},
				model.Repository{
					ID:           456,
					Name:         "test_name1",
					Description:  "test repository tags search 2",
					URL:          "http://uol.com",
					Language:     "Gotest",
					Tags:         []string{"Test1", "Node", "Docker", "Google"},
					TagSuggester: "Gotest",
				},
				model.Repository{
					ID:           789,
					Name:         "test_name1",
					Description:  "test repository tags search 2",
					URL:          "http://uol.com",
					Language:     "Gotest",
					Tags:         []string{"Documentation", "Test2", "Golang"},
					TagSuggester: "Gotest",
				},
			},
		}

		search.CreateIndex("starred_repositories", testUser)
		db.C("users").Insert(testUser)

		Convey("With a valid body searching for 'doc' ", func() {
			search := "doc"
			data := endpoint.SearchHandlerRequest{
				ID:     "test_search_user1",
				Search: search,
			}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := []model.Repository{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(resp.Body, ShouldNotBeNil)
					So(response, ShouldNotBeNil)

					ok := false
					for _, repo := range response {
						for _, tag := range repo.Tags {
							if strings.Contains(strings.ToLower(tag), strings.ToLower(search)) {
								ok = true
							}
						}
					}

					So(ok, ShouldEqual, true)
					So(resp.Code, ShouldEqual, http.StatusOK)
				})
			})
		})

		Convey("With an valid body with tag search empty", func() {
			search := ""
			data := endpoint.SearchHandlerRequest{
				ID:     "test_search_user1",
				Search: search,
			}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := []model.Repository{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(resp.Body, ShouldNotBeNil)
					So(response, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusOK)
				})
			})
		})

		Convey("With an invalid body", func() {
			data := endpoint.SearchHandlerRequest{}

			bodyReq, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", server.URL+newURL.String(), bytes.NewBuffer(bodyReq))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			Convey("When the request is handled by the router", func() {
				handler.ServeHTTP(resp, req)
				Convey("The response should be valid", func() {
					response := []model.Repository{}
					json.Unmarshal(resp.Body.Bytes(), &response)

					So(resp.Body, ShouldNotBeNil)
					So(response, ShouldNotBeNil)
					So(resp.Code, ShouldEqual, http.StatusBadRequest)
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

func clearDB(db *mgo.Database) {
	db.C("users").DropCollection()
}
