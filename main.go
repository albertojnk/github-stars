package main

import (
	"golang-crud-spa/backend/datasource"
	"golang-crud-spa/backend/endpoint"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	defer datasource.AfterStop()
	log.Println("Starting the application...")
	datasource.BeforeStart()
	endpoint.ConnectServer()
}
