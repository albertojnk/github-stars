package datasource

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

var session *mgo.Session

// Connect will connect us to mongoDB
func Connect() *mgo.Database {

	var err error
	var s *mgo.Session

	db := "golang-crud-spa"
	uri := "mongodb://localhost:27017"

	log.Println("Starting connection with MongoDB...")

	s, err = mgo.DialWithTimeout(uri, 30*time.Second)
	if err != nil {
		log.Fatalf("error while dialing mongodb, err: %s", err)
	}

	s.SetSocketTimeout(10 * time.Minute)

	session = s
	log.Println("MongoDB connected")

	return session.DB(db)
}
