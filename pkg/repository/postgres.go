package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	todoListTable   = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	conString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password,
	)
	db, err := sqlx.Open("postgres", conString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
