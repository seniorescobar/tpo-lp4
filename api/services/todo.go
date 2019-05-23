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

func (s *TodoService) List(email string) ([]entities.Todo, error) {
	return s.todoRepo.List(email)
}

func (s *TodoService) Add(email string, t *entities.Todo) (*entities.Todo, error) {
	return s.todoRepo.Add(email, t)
}

func (s *TodoService) Edit(email string, id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	return s.todoRepo.Edit(email, id, t)
}

func (s *TodoService) Remove(email string, id bson.ObjectId) error {
	return s.todoRepo.Remove(email, id)
}
