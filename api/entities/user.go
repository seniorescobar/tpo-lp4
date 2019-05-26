package entities

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id" json:"-"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Role     string        `bson:"role" json:"role"`
}
