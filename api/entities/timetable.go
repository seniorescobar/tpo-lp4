package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	Id       bson.ObjectId `bson:"_id" json:"_id""`
	Name     string        `json:"name"`
	Duration int           `json:"duration"`
	Color    string        `json:"color"`
}
