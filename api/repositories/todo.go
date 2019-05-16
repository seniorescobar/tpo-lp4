package repositories

import (
	"context"
	"log"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITodoRepo interface {
	List() ([]entities.TodoWithId, error)
	Add(context.Context, *entities.Todo) (*entities.Todo, error)
	Edit(int, *entities.Todo) (*entities.TodoWithId, error)
	Delete(int) error
}

type TodoRepo struct {
	db *mongo.Database
}

func NewTodoRepo(db *mongo.Database) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) List() ([]entities.TodoWithId, error) {
	log.Println("list")
	return nil, nil
}

func (r *TodoRepo) Add(ctx context.Context, t *entities.Todo) (*entities.Todo, error) {
	todoC := r.db.Collection("todo")

	res, err := todoC.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID

	var tNew entities.Todo
	if err := todoC.FindOne(ctx, bson.M{"_id": id}).Decode(&tNew); err != nil {
		return nil, err
	}

	// FIXME hack; find out how FindOne can return full doc (including _id)
	tNew.Id = id

	return &tNew, nil
}

func (r *TodoRepo) Edit(id int, t *entities.Todo) (*entities.TodoWithId, error) {
	log.Println("edit", id, *t)
	return nil, nil
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

func (m *TodoRepoMock) Add(ctx context.Context, t *entities.Todo) (*entities.Todo, error) {
	args := m.Called(ctx, t)
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *TodoRepoMock) Edit(id int, t *entities.Todo) (*entities.TodoWithId, error) {
	args := m.Called(id, t)
	return args.Get(0).(*entities.TodoWithId), args.Error(1)
}

func (m *TodoRepoMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
