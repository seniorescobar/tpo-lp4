package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"github.com/stretchr/testify/mock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ITodoRepo interface {
	List(email string) ([]entities.Todo, error)
	Add(email string, t *entities.Todo) (*entities.Todo, error)
	Edit(email string, id bson.ObjectId, t *entities.Todo) (*entities.Todo, error)
	Remove(email string, id bson.ObjectId) error
}

type TodoRepo struct {
	db *mgo.Database
}

func NewTodoRepo(db *mgo.Database) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) List(email string) ([]entities.Todo, error) {
	var todos []entities.Todo
	if err := r.db.C("todo").Find(bson.M{"email": email}).All(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepo) Add(email string, t *entities.Todo) (*entities.Todo, error) {
	tNew := &entities.Todo{
		Id:          bson.NewObjectId(),
		Email:       email,
		Description: t.Description,
	}

	if err := r.db.C("todo").Insert(tNew); err != nil {
		return nil, err
	}

	return tNew, nil
}

func (r *TodoRepo) Edit(email string, id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	if err := r.db.C("todo").Update(bson.M{"_id": id, "email": email}, bson.M{"$set": bson.M{"description": t.Description}}); err != nil {
		return nil, err
	}

	var tNew entities.Todo
	if err := r.db.C("todo").Find(bson.M{"_id": id, "email": email}).One(&tNew); err != nil {
		return nil, err
	}

	return &tNew, nil
}

func (r *TodoRepo) Remove(email string, id bson.ObjectId) error {
	if err := r.db.C("todo").Remove(bson.M{"_id": id, "email": email}); err != nil {
		return err
	}

	return nil
}

// MOCK

type TodoRepoMock struct {
	mock.Mock
}

func NewTodoRepoMock() *TodoRepoMock {
	return new(TodoRepoMock)
}

func (m *TodoRepoMock) List() ([]entities.Todo, error) {
	args := m.Called()
	return args.Get(0).([]entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Add(t *entities.Todo) (*entities.Todo, error) {
	args := m.Called(t)
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Edit(id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	args := m.Called(id, t)
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Remove(id bson.ObjectId) error {
	args := m.Called(id)
	return args.Error(0)
}
