package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IUserRepo interface {
	Register(u *entities.User) (*entities.User, error)
	Signin(string, string) (*entities.User, error)
	ChangePassword(email, oldPassword, newPassword string) error
}

type UserRepo struct {
	c *mgo.Collection
}

func NewUserRepo(db *mgo.Database) *UserRepo {
	return &UserRepo{db.C("user")}
}

func (r *UserRepo) Register(u *entities.User) (*entities.User, error) {
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	uNew := &entities.User{
		Id:       bson.NewObjectId(),
		Email:    u.Email,
		Password: hashedPassword,
		Role:     u.Role,
	}

	if err := r.c.Insert(uNew); err != nil {
		return nil, err
	}

	return uNew, nil
}

func (r *UserRepo) Signin(email, password string) (*entities.User, error) {
	var u entities.User
	if err := r.c.Find(bson.M{"email": email}).One(&u); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepo) ChangePassword(email, oldPassword, newPassword string) error {
	if _, err := r.Signin(email, oldPassword); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := r.c.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": hashedPassword}}); err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", nil
	}

	return string(hp), nil
}
