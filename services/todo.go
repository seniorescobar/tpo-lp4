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

func (s *TodoService) List() ([]entities.TodoWithId, error) {
	return s.todoRepo.List()
}

func (s *TodoService) Add(t *entities.Todo) error {
	return s.todoRepo.Add(t)
}

func (s *TodoService) Edit(id int, t *entities.Todo) (*entities.TodoWithId, error) {
	return s.todoRepo.Edit(id, t)
}
