package endpoint_test

import (
	"bytes"
	"encoding/json"
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/endpoint"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRepository(t *testing.T) {
	Convey("Given a POST request to /create", t, func() {
		newURL, _ := url.Parse("/create")
		handler := endpoint.NewEndpoint()
		server := httptest.NewServer(handler)
		datasource.Connect()
		defer server.Close()

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
	})
}
