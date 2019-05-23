package entities

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id          bson.ObjectId `bson:"_id" json:"_id""`
	Email       string        `json:"-"`
	Description string        `json:"description"`
}
