package datasource

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

var session *mgo.Session
var database *mgo.Database

// Connect will connect us to mongoDB
func connect() {

	var err error
	var s *mgo.Session

	db := "golang-crud-spa"
	uri := "mongodb://mongodb:27018"

	log.Println("Starting connection with MongoDB...")

	s, err = mgo.DialWithTimeout(uri, 30*time.Second)
	if err != nil {
		log.Fatalf("error while dialing mongodb, err: %s", err)
	}

	s.SetSocketTimeout(10 * time.Minute)

	session = s
	log.Println("MongoDB connected")

	database = session.DB(db)
}

func disconnect() {
	if session != nil {
		session.Close()
	}
}

// BeforeStart will be called as the program starts ...
func BeforeStart() {
	connect()
}

// AfterStop will be called after program stops ...
func AfterStop() {
	disconnect()
}
