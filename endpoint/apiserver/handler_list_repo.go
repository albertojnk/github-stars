package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/golang-crud-spa/datasource"
)

// ListRepositories get our repositories from DB and return
func ListRepositories(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	db := datasource.Connect()
	users := []User{}

	values, ok := r.URL.Query()["username"]

	if !ok || len(values[0]) < 1 {
		log.Println("url parameter 'username' is missing")
		return
	}

	username := values[0]

	err := db.C("users").Find(bson.M{"_id": username}).All(&users)
	if err != nil {
		log.Printf("something went wrong while accessing the DB, err: ", err)
	}

	json.NewEncoder(rw).Encode(users)
}
