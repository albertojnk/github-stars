package main

import (
	"log"
	"os"

	"github.com/golang-crud-spa/backend/search"

	"github.com/golang-crud-spa/backend/endpoint"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	defer os.RemoveAll(search.SearchDIR)
	log.Println("Starting the application...")
	endpoint.Handlers()
}
