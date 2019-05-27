package container

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	DB *mgo.Database
)

func initDB() {
	mongodb_uri := os.Getenv("MONGODB_URI")

	session, err := mgo.Dial(mongodb_uri)
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = session.DB("")
}
