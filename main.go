package main

import (
	"github.com/golang-crud-spa/endpoint/apiserver"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	apiserver.Handlers()

}
