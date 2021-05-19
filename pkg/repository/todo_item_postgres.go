package repository

import (
	"fmt"
	"strings"

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

func (r *TodoItemPostgres) GetAll(userID, listID int) ([]workshop_2.TodoItem, error) {
	var items []workshop_2.TodoItem

	query := fmt.Sprintf(
		`select ti.title, ti.description, ti.done from %s ti
		inner join %s li
		inner join %s ul
		on li.id  = ul.list_id
		on ti.id = li.item_id
		where ul.user_id = $1
		and li.list_id = $2`,
		todoItemsTable, todoListTable, usersListsTable,
	)

	if err := r.db.Select(&items, query, userID, listID); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetByID(userID, itemID int) (workshop_2.TodoItem, error) {
	var item workshop_2.TodoItem

	query := fmt.Sprintf(
		`select ti.title, ti.description, ti.done from %s ti
		inner join %s li
		inner join %s ul
		on li.list_id  = ul.list_id
		on ti.id = li.item_id
		where ul.user_id = $1
		and ti.id = $2`,
		todoItemsTable, todoListTable, usersListsTable,
	)

	if err := r.db.Get(&item, query, userID, itemID); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Update(userID, itemID int, input workshop_2.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s ti set %s from %s li, %s ul where ti.id = li.item_id and li.list_id=ul.list_id and ul.user_id = $%d and ti.id = $%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argID, argID+1)
	args = append(args, userID, itemID)

	_, err := r.db.Exec(query, args...)
	return err
}
