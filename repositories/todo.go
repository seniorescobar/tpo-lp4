package repositories

import (
	"database/sql"
	"log"

	"bitbucket.org/aj5110/tpo-lp4/entities"
	"github.com/stretchr/testify/mock"
)

type ITodoRepo interface {
	List() ([]entities.TodoWithId, error)
	Add(*entities.Todo) error
	Edit(int, *entities.Todo) (*entities.TodoWithId, error)
	Delete(int) error
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) List() ([]entities.TodoWithId, error) {
	log.Println("list")
	return nil, nil
}

func (r *TodoRepo) Add(t *entities.Todo) error {
	log.Println("add", *t)
	return nil
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

func (m *TodoRepoMock) Add(t *entities.Todo) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *TodoRepoMock) Edit(id int, t *entities.Todo) (*entities.TodoWithId, error) {
	args := m.Called(id, t)
	return args.Get(0).(*entities.TodoWithId), args.Error(1)
}

func (m *TodoRepoMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
