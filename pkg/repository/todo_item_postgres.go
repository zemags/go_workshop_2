package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listID int, item workshop_2.TodoItem) (int, error) {
	// Run transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQuery := fmt.Sprintf("isnert into %s (title, description) values ($1, $2) returning id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(itemID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("insert into %s (item_id, list_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, itemID, listID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemID, tx.Commit()

}
