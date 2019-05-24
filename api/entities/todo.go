package entities

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id          bson.ObjectId `bson:"_id" json:"_id""`
	UserId      bson.ObjectId `bson:"user_id" json:"-"`
	Description string        `bson:"description" json:"description"`
}
