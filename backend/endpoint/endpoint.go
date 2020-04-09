package endpoint

import (
	"encoding/json"
	"github-stars/backend/model"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var indexName = "starred_repositories"

// NewEndpoint route each handler registered in this function
func NewEndpoint() *echo.Echo {
	e := echo.New()

	svc := e.Group("")

	//add middleware
	svc.Use(middleware.Recover())
	svc.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		Skipper:      middleware.DefaultSkipper,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// SERVICES LAYER
	svc.POST("/create", CreateRepository)
	svc.POST("/search", SearchHandler)

	svc.GET("/list", ListRepositories)
	svc.PATCH("/update", UpdateRepositoryTags)
	svc.DELETE("/delete", DeleteRepositoryTags)

	// WEB LAYER
	webPath := os.Getenv("WEB_PATH")
	if webPath == "" {
		webPath = "frontend/app/dist"
	}

	sts := e.Group("")
	sts.Static("/", webPath)
	sts.Static("/repositories", webPath)

	return e
}

// Decode body
func Decode(body io.ReadCloser, bodyType string) (interface{}, error) {
	switch bodyType {
	case "create":
		reqData := CreateRepositoryRequest{}
		// getting the username from frontend and decoding it
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("something went wrong, err: %s", err)
			return reqData, err
		}

		// Unmarshaling the decoded username
		err = json.Unmarshal(b, &reqData)
		if err != nil {
			log.Printf("error while unmarshaling, err: %s", err)
			return reqData, err
		}
		return reqData, err

	case "update":
		reqData := UpdateRepositoryTagsRequest{}
		// getting the username from frontend and decoding it
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("something went wrong, err: %s", err)
			return reqData, err
		}

		// Unmarshaling the decoded username
		err = json.Unmarshal(b, &reqData)
		if err != nil {
			log.Printf("error while unmarshaling, err: %s", err)
			return reqData, err
		}
		return reqData, err

	case "repository":
		reqData := model.Repository{}
		// getting the username from frontend and decoding it
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("something went wrong, err: %s", err)
			return reqData, err
		}

		// Unmarshaling the decoded username
		err = json.Unmarshal(b, &reqData)
		if err != nil {
			log.Printf("error while unmarshaling, err: %s", err)
			return reqData, err
		}
		return reqData, err

	case "delete":
		reqData := DeleteRepositoryTagsRequest{}
		// getting the username from frontend and decoding it
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("something went wrong, err: %s", err)
			return reqData, err
		}

		// Unmarshaling the decoded username
		err = json.Unmarshal(b, &reqData)
		if err != nil {
			log.Printf("error while unmarshaling, err: %s", err)
			return reqData, err
		}
		return reqData, err

	case "search":
		reqData := SearchHandlerRequest{}
		// getting the username from frontend and decoding it
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("something went wrong, err: %s", err)
			return reqData, err
		}

		// Unmarshaling the decoded username
		err = json.Unmarshal(b, &reqData)
		if err != nil {
			log.Printf("error while unmarshaling, err: %s", err)
			return reqData, err
		}
		return reqData, err

	}
	return nil, nil
}
