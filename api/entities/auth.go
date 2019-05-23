package entities

import "gopkg.in/mgo.v2/bson"

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id       bson.ObjectId `bson:"_id" json:"_id""`
	Email    string        `json:"email"`
	Password string        `json:"-"`
}
