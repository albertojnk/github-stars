package datasource

import (
	"log"
	"os"
	"time"

	"github.com/globalsign/mgo"
)

var session *mgo.Session
var database *mgo.Database

// Connect will connect us to mongoDB
func Connect() {

	var err error
	var s *mgo.Session

	db := "github-stars"
	uri := os.Getenv("MONGO_URI")

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

// Disconnect session
func Disconnect() {
	if session != nil {
		session.Close()
	}
}

// BeforeStart will be called as the program starts ...
func BeforeStart() {
	Connect()
}

// AfterStop will be called after program stops ...
func AfterStop() {
	Disconnect()
}

// GetDatabase return a *mgo.Database an external pkg
func GetDatabase() *mgo.Database {
	return database
}
