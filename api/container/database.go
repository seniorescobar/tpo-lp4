package container

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var (
	DB *mgo.Database
)

func initDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = session.DB("tpo")
}
