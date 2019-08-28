package main

import (
	"log"

	"github.com/albertojnk/golang-crud-spa/backend/endpoint"
	"github.com/albertojnk/golang-crud-spa/backend/search"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	log.Println("Starting the application...")
	endpoint.Handlers()
	search.StartDefaultClient()
}
