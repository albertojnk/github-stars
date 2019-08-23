package main

import (
	"log"

	"github.com/golang-crud-spa/endpoint/apiserver"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	log.Println("Starting the application...")

	apiserver.Handlers()

}
