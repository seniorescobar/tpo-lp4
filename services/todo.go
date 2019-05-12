package services

import (
	"bitbucket.org/aj5110/tpo-lp4/entities"
	"bitbucket.org/aj5110/tpo-lp4/repositories"
)

type TodoService struct {
	todoRepo repositories.ITodoRepo
}

func NewTodoService(todoRepo repositories.ITodoRepo) *TodoService {
	return &TodoService{todoRepo}
}

func (s *TodoService) Add(t *entities.Todo) error {
	return s.todoRepo.Add(t)
}
