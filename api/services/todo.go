package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"gopkg.in/mgo.v2/bson"
)

type TodoService struct {
	todoRepo repositories.ITodoRepo
}

func NewTodoService(todoRepo repositories.ITodoRepo) *TodoService {
	return &TodoService{todoRepo}
}

func (s *TodoService) List() ([]entities.Todo, error) {
	return s.todoRepo.List()
}

func (s *TodoService) Add(t *entities.Todo) (*entities.Todo, error) {
	return s.todoRepo.Add(t)
}

func (s *TodoService) Edit(id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	return s.todoRepo.Edit(id, t)
}

func (s *TodoService) Remove(id bson.ObjectId) error {
	return s.todoRepo.Remove(id)
}
