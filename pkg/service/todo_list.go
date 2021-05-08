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
