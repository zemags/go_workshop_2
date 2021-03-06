package service

import (
	"github.com/zemags/go_workshop_2/pkg/repository"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create list and return list id and error
func (s *TodoListService) Create(userID int, list workshop_2.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]workshop_2.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (workshop_2.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input workshop_2.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, listID, input)
}
