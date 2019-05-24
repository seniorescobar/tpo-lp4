package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IAuthRepo interface {
	Register(u *entities.User) error
	Signin(string, string) (*entities.User, error)
}

type AuthRepo struct {
	db *mgo.Database
}

func NewAuthRepo(db *mgo.Database) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) Register(u *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}

	uNew := &entities.User{
		Id:       bson.NewObjectId(),
		Email:    u.Email,
		Password: string(hashedPassword),
	}

	if err := r.db.C("user").Insert(uNew); err != nil {
		return err
	}

	return nil
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
