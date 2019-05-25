package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	UserId      bson.ObjectId `bson:"user_id" json:"-"`
	Name        string        `json:"name"`
	StartDt     time.Time     `json:"start_dt"`
	EndDt       time.Time     `json:"end_dt"`
	Description string        `json:"description"`
}
