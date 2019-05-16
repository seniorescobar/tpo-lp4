package repositories

import (
	"log"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"github.com/stretchr/testify/mock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ITodoRepo interface {
	List() ([]entities.TodoWithId, error)
	Add(*entities.Todo) (*entities.Todo, error)
	Edit(bson.ObjectId, *entities.Todo) (*entities.Todo, error)
	Delete(int) error
}

type TodoRepo struct {
	db *mgo.Database
}

func NewTodoRepo(db *mgo.Database) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) List() ([]entities.TodoWithId, error) {
	log.Println("list")
	return nil, nil
}

func (r *TodoRepo) Add(t *entities.Todo) (*entities.Todo, error) {
	tNew := &entities.Todo{
		Id:          bson.NewObjectId(),
		Description: t.Description,
	}

	if err := r.db.C("todo").Insert(tNew); err != nil {
		return nil, err
	}

	return tNew, nil
}

func (r *TodoRepo) Edit(id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	tNew := &entities.Todo{
		Id:          id,
		Description: t.Description,
	}

	if err := r.db.C("todo").UpdateId(id, tNew); err != nil {
		return nil, err
	}

	return tNew, nil
}

func (r *TodoRepo) Delete(id int) error {
	log.Println("delete", id)
	return nil
}

// MOCK

type TodoRepoMock struct {
	mock.Mock
}

func NewTodoRepoMock() *TodoRepoMock {
	return new(TodoRepoMock)
}

func (m *TodoRepoMock) List() ([]entities.TodoWithId, error) {
	args := m.Called()
	return args.Get(0).([]entities.TodoWithId), args.Error(1)
}

func (m *TodoRepoMock) Add(t *entities.Todo) (*entities.Todo, error) {
	args := m.Called(t)
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Edit(id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	args := m.Called(id, t)
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
