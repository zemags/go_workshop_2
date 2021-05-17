package service

import (
	"github.com/zemags/go_workshop_2/pkg/repository"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type Authorization interface {
	// CreateUser return user id and error
	CreateUser(user workshop_2.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list workshop_2.TodoList) (int, error)
	GetAll(userID int) ([]workshop_2.TodoList, error)
	GetByID(userID, listID int) (workshop_2.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input workshop_2.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService - its a constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		TodoList:      NewTodoListService(repos),
	}
}
