package main

import (
	"github-stars/backend/datasource"
	"github-stars/backend/endpoint"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Service serves a router
type Service struct {
	router *echo.Echo
}

func main() {
	defer datasource.AfterStop()
	log.Println("Starting the application...")
	datasource.BeforeStart()

	service, err := NewService()
	if err != nil {
		panic("failed to start")
	}
	service.Start()
}

// NewService returns a server service
func NewService() (*Service, error) {
	svc := Service{
		router: endpoint.NewEndpoint(),
	}

	return &svc, nil
}

// Start endpoint
func (s Service) Start() {
	log.Println("HTTP Listening on port 8090")
	log.Fatal(http.ListenAndServe(":8090", s.router))
}
