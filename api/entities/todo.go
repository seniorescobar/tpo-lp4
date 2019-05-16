package entities

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id          bson.ObjectId `bson:"_id" json:"_id""`
	Description string        `json:"description"`
}
