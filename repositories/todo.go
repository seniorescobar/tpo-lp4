package repositories

import (
	"context"
	"log"

	"bitbucket.org/aj5110/tpo-lp4/entities"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITodoRepo interface {
	List() ([]entities.TodoWithId, error)
	Add(context.Context, *entities.Todo) (*entities.TodoWithId, error)
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

func (r *TodoRepo) Add(ctx context.Context, t *entities.Todo) (*entities.TodoWithId, error) {
	res, err := r.db.Collection("todo").InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}

	log.Println("res", res)

	tNew := entities.TodoWithId{1, entities.Todo{"abc"}}
	// if err := r.db.Collection("todo").FindOne(ctx, entities.TodoWithId{Id: res.ID}); err != nil {
	// 	return nil, err
	// }

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

func (m *TodoRepoMock) Add(ctx context.Context, t *entities.Todo) (*entities.TodoWithId, error) {
	args := m.Called(ctx, t)
	return args.Get(0).(*entities.TodoWithId), args.Error(1)
}

func (m *TodoRepoMock) Edit(id int, t *entities.Todo) (*entities.TodoWithId, error) {
	args := m.Called(id, t)
	return args.Get(0).(*entities.TodoWithId), args.Error(1)
}

func (m *TodoRepoMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
