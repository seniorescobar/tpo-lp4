package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	Id       bson.ObjectId `bson:"_id" json:"_id"`
	UserId   bson.ObjectId `bson:"user_id" json:"-"`
	Name     string        `bson:"name" json:"name"`
	Duration int           `bson:"duration" json:"duration"`
	Color    string        `bson:"color" json:"color"`
}
