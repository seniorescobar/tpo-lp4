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

func (s *TodoService) List(uid bson.ObjectId) ([]entities.Todo, error) {
	return s.todoRepo.List(uid)
}

func (s *TodoService) Add(uid bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	return s.todoRepo.Add(uid, t)
}

func (s *TodoService) Edit(t *entities.Todo) error {
	return s.todoRepo.Edit(t)
}

func (s *TodoService) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	return s.todoRepo.Remove(uid, id)
}
