package service

import (
	"github.com/zemags/go_workshop_2/pkg/repository"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID, listID int, item workshop_2.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list doesnt exist or doesnt belongs to user
		return 0, err
	}

	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]workshop_2.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}
