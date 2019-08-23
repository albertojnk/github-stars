package main

import (
	"log"

	"github.com/golang-crud-spa/backend/endpoint"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	log.Println("Starting the application...")

	endpoint.Handlers()

}
