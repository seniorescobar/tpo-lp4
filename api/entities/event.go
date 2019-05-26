package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Date        time.Time     `bson:"date" json:"date"`
	Organizer   string        `bson:"organizer" json:"organizer"`
	Description string        `bson:"description" json:"description"`
}
