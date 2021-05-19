package repository

import (
	"github.com/jmoiron/sqlx"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type Authorization interface {
	CreateUser(user workshop_2.User) (int, error)
	GetUser(username, password string) (workshop_2.User, error)
}

type TodoList interface {
	Create(userID int, list workshop_2.TodoList) (int, error)
	GetAll(userID int) ([]workshop_2.TodoList, error)
	GetByID(userID, listID int) (workshop_2.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input workshop_2.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item workshop_2.TodoItem) (int, error)
	GetAll(userID, listID int) ([]workshop_2.TodoItem, error)
	GetByID(userID, itemID int) (workshop_2.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository its a constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		// initialize
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
