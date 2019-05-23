package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IAuthRepo interface {
	Signin(string, string) (*entities.User, error)
}

type AuthRepo struct {
	db *mgo.Database
}

func NewAuthRepo(db *mgo.Database) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) Signin(email, password string) (*entities.User, error) {
	var u entities.User
	if err := r.db.C("user").Find(bson.M{"email": email}).One(&u); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &u, nil
}
